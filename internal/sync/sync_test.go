package sync

import (
	"context"
	"testing"

	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/carbonable/carbonable-launchpad-backend/ent"
	"github.com/carbonable/carbonable-launchpad-backend/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestSynchronize(t *testing.T) {
	// Prepare
	ctx := context.Background()
	db := ent.NewTestClient(t)
	client, err := rpc.NewProvider("https://free-rpc.nethermind.io/sepolia-juno")
	if err != nil {
		t.Errorf("failed to dial in rpc provider : %s", err)
	}

	// Run sync code
	err = Synchronize(ctx, db, client)
	if err != nil {
		t.Errorf("failed to sync : %s", err)
	}

	// Assert
	project, err := db.Project.Query().All(context.Background())
	if err != nil {
		t.Errorf("failed to query projects : %s", err)
	}

	assert.Equal(t, 4, len(project))
}

func TestSplitSlug(t *testing.T) {
	slotUri := model.SlotUri{ExternalUrl: "https://app.carbonable.io/launchpad/forest-regeneration-banegas-farm-costa-rica"}

	slug := slotUri.Slug()
	assert.Equal(t, "forest-regeneration-banegas-farm-costa-rica", slug)
}
