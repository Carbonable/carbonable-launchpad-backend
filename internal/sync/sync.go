package sync

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/carbonable-labs/indexer.sdk/sdk"
	"github.com/carbonable/carbonable-launchpad-backend/config"
	"github.com/carbonable/carbonable-launchpad-backend/ent"
	"github.com/carbonable/carbonable-launchpad-backend/ent/project"
	"github.com/carbonable/carbonable-launchpad-backend/ent/schema"
	"github.com/carbonable/carbonable-launchpad-backend/internal/model"
)

// Sync contracts with onchain data
func Synchronize(ctx context.Context, db *ent.Client, rpc rpc.RpcProvider) error {
	// Update indexer config
	c := config.GetContracts()
	project := c.FilterByName("project")
	for _, c := range project.Contracts {
		sc, err := c.Call(ctx, rpc, "slot_count")
		if err != nil {
			return err
		}
		slotCount := sc[0].Uint64()
		vd := getValueDecimals(ctx, rpc, c)
		var wg sync.WaitGroup
		for i := uint64(1); i < slotCount+1; i++ {
			p, err := syncProject(ctx, db, rpc, c, i, vd)
			if err != nil {
				slog.Error("failed to sync project", "err", err)
				continue
			}
			wg.Add(2)
			go syncMint(ctx, &wg, db, rpc, c, i, p)
			go syncLaunchpad(ctx, &wg, db, rpc, c, i, p)
			wg.Wait()
		}
	}

	return nil
}

func syncProject(ctx context.Context, db *ent.Client, rpc rpc.RpcProvider, c sdk.Contract, i uint64, vd uint64) (*ent.Project, error) {
	var slot felt.Felt
	slot.SetUint64(i)
	if !slotIsSetup(ctx, rpc, c, &slot) {
		return nil, fmt.Errorf("slot is not setup")
	}

	slotUri, err := getSlotUri(ctx, rpc, c, &slot)
	if err != nil {
		slog.Error("faield to get slot uri", err)
		return nil, err
	}
	minterAddr, err := getMinterAddress(ctx, rpc, c, &slot)
	if err != nil {
		slog.Error("faield to get minterAddr", err)
		return nil, err
	}
	tv := getTotalValue(ctx, rpc, c, &slot)
	if err != nil {
		slog.Error("faield to marshal total value", err)
		return nil, err
	}

	paymentToken, err := getPaymentToken(ctx, rpc, minterAddr)
	if err != nil {
		slog.Error("failed to get payment token", err)
		return nil, err
	}
	metadata, err := MetadataForProject(slotUri.Name)
	if err != nil {
		slog.Error("metadata not created", err)
		return nil, err
	}

	err = db.Project.Create().
		SetAddress(c.Address).
		SetSlot(int(i)).
		SetName(slotUri.Name).
		SetSlug(slotUri.Slug()).
		SetValueDecimal(int(vd)).
		SetForecastedApr("").
		SetTotalValue(tv.String()).
		SetPaymentToken(paymentToken).
		SetMetadata(metadata).
		OnConflict(sql.ConflictColumns("address", "slot")).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		slog.Error("faield to create project", err)
		return nil, err
	}
	p, _ := db.Project.Query().Where(project.AddressEQ(c.Address), project.SlotEQ(int(i))).Only(ctx)
	return p, nil
}

func syncMint(ctx context.Context, wg *sync.WaitGroup, db *ent.Client, rpc rpc.RpcProvider, c sdk.Contract, i uint64, p *ent.Project) {
	defer wg.Done()

	var slot felt.Felt
	slot.SetUint64(i)
	if !slotIsSetup(ctx, rpc, c, &slot) {
		return
	}

	minterAddr, err := getMinterAddress(ctx, rpc, c, &slot)
	if err != nil {
		slog.Error("faield to get minterAddr", err)
		return
	}
	minPerTx, err := callContract(ctx, rpc, minterAddr, "get_min_value_per_tx")
	if err != nil {
		slog.Error("faield to get min_value_per_tx", err)
		return
	}
	maxPerTx, err := callContract(ctx, rpc, minterAddr, "get_max_value_per_tx")
	if err != nil {
		slog.Error("faield to get max_value_per_tx", err)
		return
	}

	err = db.Mint.Create().
		SetMinterAddress(minterAddr).
		SetMinValuePerTx(minPerTx[0].String()).
		SetMaxValuePerTx(maxPerTx[0].String()).
		SetProject(p).
		OnConflict(sql.ConflictColumns("project_mint")).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		slog.Error("faield to create mint", err)
		return
	}
}

func syncLaunchpad(ctx context.Context, wg *sync.WaitGroup, db *ent.Client, rpc rpc.RpcProvider, c sdk.Contract, i uint64, p *ent.Project) {
	defer wg.Done()

	var slot felt.Felt
	slot.SetUint64(i)
	if !slotIsSetup(ctx, rpc, c, &slot) {
		return
	}

	abi, err := getAbi(ctx, rpc, c.Address)
	if err != nil {
		slog.Error("faield to get abi", err)
		return
	}
	minterAddr, err := getMinterAddress(ctx, rpc, c, &slot)
	if err != nil {
		slog.Error("faield to get minterAddr", err)
		return
	}
	isPreSaleOpen, err := toBool(callContract, ctx, rpc, minterAddr, "is_pre_sale_open")
	if err != nil {
		slog.Error("faield to get is_pre_sale_open", err)
		return
	}
	isPublicSaleOpen, err := toBool(callContract, ctx, rpc, minterAddr, "is_public_sale_open")
	if err != nil {
		slog.Error("faield to get is_public_sale_open", err)
		return
	}
	isSoldOut, err := toBool(callContract, ctx, rpc, minterAddr, "is_sold_out")
	if err != nil {
		slog.Error("faield to get is_sold_out", err)
		return
	}
	isCanceled, err := toBool(callContract, ctx, rpc, minterAddr, "is_canceled")
	if err != nil {
		isCanceled = false
	}

	err = db.Launchpad.Create().
		SetMinterContract(schema.MinterContract{
			Address: minterAddr,
			Abi:     abi,
		}).
		SetWhitelistedSaleOpen(isPreSaleOpen).
		SetPublicSaleOpen(isPublicSaleOpen).
		SetIsSoldOut(isSoldOut).
		SetIsSoldOut(isCanceled).
		SetProject(p).
		OnConflict(sql.ConflictColumns("project_launchpad")).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		slog.Error("faield to create launchpad", err)
		return
	}
}

// Get sloturi from contract
func getSlotUri(ctx context.Context, rpc rpc.RpcProvider, c sdk.Contract, slot *felt.Felt) (model.SlotUri, error) {
	var slotUri model.SlotUri
	uri, err := c.Call(ctx, rpc, "slot_uri", slot, &felt.Zero)
	if err != nil {
		return slotUri, err
	}

	strVal := feltArrToBytesArr(uri[2:])

	err = json.Unmarshal(strVal, &slotUri)
	if err != nil {
		return slotUri, err
	}

	return slotUri, nil
}

// Check if contract slot is setup in db
func slotIsSetup(ctx context.Context, rpc rpc.RpcProvider, c sdk.Contract, i *felt.Felt) bool {
	impl, err := c.Call(ctx, rpc, "is_setup", i, &felt.Zero)
	if err != nil {
		return false
	}

	return impl[0].Uint64() == uint64(1)
}

// Get class abi from rpc using contract address
func getAbi(ctx context.Context, r rpc.RpcProvider, a string) (json.RawMessage, error) {
	addr, err := utils.HexToFelt(a)
	if err != nil {
		return nil, err
	}
	out, rpcErr := r.ClassAt(ctx, rpc.BlockID{Tag: "latest"}, addr)
	if rpcErr != nil {
		return nil, err
	}
	class := out.(*rpc.ContractClass)
	return json.RawMessage(class.ABI), nil
}

// Get value decimals from contract
func getValueDecimals(ctx context.Context, r rpc.RpcProvider, c sdk.Contract) uint64 {
	vd, err := c.Call(ctx, r, "value_decimals")
	if err != nil {
		return uint64(6)
	}

	return vd[0].Uint64()
}

// Get value decimals from contract
func getTotalValue(ctx context.Context, r rpc.RpcProvider, c sdk.Contract, slot *felt.Felt) *felt.Felt {
	tv, err := c.Call(ctx, r, "total_value", slot, &felt.Zero)
	if err != nil {
		return &felt.Zero
	}
	return tv[0]
}

// Get minter address from contract get_minters list
// last minter inserted is the minter in use (just reverse the list to get the first one)
func getMinterAddress(ctx context.Context, r rpc.RpcProvider, c sdk.Contract, slot *felt.Felt) (string, error) {
	uri, err := c.Call(ctx, r, "get_minters", slot, &felt.Zero)
	if err != nil {
		return "", err
	}

	for i := len(uri[1:]); i > 0; i-- {
		addr := uri[i].String()
		callResp, err := callContract(ctx, r, addr, "get_carbonable_project_address")
		if err != nil {
			continue
		}

		feltAddr, _ := utils.HexToFelt(c.Address)
		if callResp[0].String() == feltAddr.String() {
			return addr, nil
		}
	}
	return "", nil
}

// PaymentToken informations from project minter address
func getPaymentToken(ctx context.Context, r rpc.RpcProvider, c string) (schema.PaymentToken, error) {
	paymentTokenAddress, err := callContract(ctx, r, c, "get_payment_token_address")
	if err != nil {
		return schema.PaymentToken{}, err
	}

	symbol, err := callContract(ctx, r, paymentTokenAddress[0].String(), "symbol")
	if err != nil {
		return schema.PaymentToken{}, err
	}
	s, err := hex.DecodeString(strings.Replace(symbol[0].String(), "0x", "", 1))
	if err != nil {
		return schema.PaymentToken{
			Symbol:  symbol[0].String(),
			Address: paymentTokenAddress[0].String(),
		}, err
	}

	return schema.PaymentToken{
		Symbol:  string(s),
		Address: paymentTokenAddress[0].String(),
	}, nil
}

func callContract(ctx context.Context, r rpc.RpcProvider, a string, f string, args ...*felt.Felt) ([]*felt.Felt, error) {
	addr, err := utils.HexToFelt(a)
	if err != nil {
		return nil, err
	}
	tx := rpc.FunctionCall{
		ContractAddress:    addr,
		EntryPointSelector: utils.GetSelectorFromNameFelt(f),
	}
	callResp, rpcErr := r.Call(ctx, tx, rpc.BlockID{Tag: "latest"})
	if rpcErr != nil {
		return nil, rpcErr
	}

	return callResp, nil
}

func toBool(cb func(context.Context, rpc.RpcProvider, string, string, ...*felt.Felt) ([]*felt.Felt, error), ctx context.Context, r rpc.RpcProvider, a string, f string, args ...*felt.Felt) (bool, error) {
	resp, err := cb(ctx, r, a, f, args...)
	if err != nil {
		return false, err
	}
	return resp[0].Uint64() == 1, nil
}

// Convert cairo felt array to byte array
func feltArrToBytesArr(feltArr []*felt.Felt) []byte {
	var bArr []byte
	for _, f := range feltArr {
		b := f.Marshal()
		bArr = append(bArr, bytes.Trim(b[0:], "\x00")...)
	}
	return bArr
}

func MetadataForProject(name string) (schema.Metadata, error) {
	m := map[string]schema.Metadata{
		"Banegas Farm": {
			Rating:   "A",
			TonPrice: "11.19$",
			Milestones: []schema.Milestone{
				{
					Boost:    "0",
					Ceil:     17600,
					Area:     "",
					Id:       0,
					TonPrice: "",
				},
			},
		},
		"Las Delicias": {
			Rating:   "A",
			TonPrice: "10.99$",
			Milestones: []schema.Milestone{
				{
					Boost:    "0",
					Ceil:     39600,
					Area:     "",
					Id:       0,
					TonPrice: "",
				},
			},
		},
		"Manjarisoa": {
			Rating:   "A",
			TonPrice: "15.12$",
			Milestones: []schema.Milestone{
				{
					Boost:    "3",
					Ceil:     121099,
					Area:     "",
					Id:       0,
					TonPrice: "",
				},
			},
		},
		"Karathuru": {
			Rating:   "AA",
			TonPrice: "5.212$",
			Milestones: []schema.Milestone{
				{
					Boost:    "3",
					Ceil:     50000,
					Area:     "12",
					Id:       0,
					TonPrice: "9593",
				},
				{
					Boost:    "2",
					Ceil:     150000,
					Area:     "35",
					Id:       1,
					TonPrice: "28780",
				},
				{
					Boost:    "1.5",
					Ceil:     300000,
					Area:     "70",
					Id:       2,
					TonPrice: "57559",
				},
				{
					Boost:    "1.2",
					Ceil:     500000,
					Area:     "117",
					Id:       3,
					TonPrice: "95932",
				},
				{
					Boost:    "1.1",
					Ceil:     700000,
					Area:     "164",
					Id:       4,
					TonPrice: "134305",
				},
				{
					Boost:    "",
					Ceil:     1000000,
					Area:     "234",
					Id:       5,
					TonPrice: "191865",
				},
				{
					Boost:    "",
					Ceil:     1300000,
					Area:     "304",
					Id:       6,
					TonPrice: "249424",
				},
				{
					Boost:    "",
					Ceil:     1600000,
					Area:     "374",
					Id:       7,
					TonPrice: "306984",
				},
				{
					Boost:    "",
					Ceil:     1900000,
					Area:     "444",
					Id:       8,
					TonPrice: "364543",
				},
				{
					Boost:    "",
					Ceil:     2200000,
					Area:     "500",
					Id:       9,
					TonPrice: "410400",
				},
			},
		},
	}
	if v, ok := m[name]; ok {
		return v, nil
	}
	return schema.Metadata{}, fmt.Errorf("metadata not found")
}
