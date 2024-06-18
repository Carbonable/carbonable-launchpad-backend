package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/carbonable/carbonable-launchpad-backend/internal/sync"
	"github.com/carbonable/carbonable-launchpad-backend/internal/utils"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()
	db, err := utils.OpenDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("failed opening connection to database", err)
		return
	}

	rpcClient, err := rpc.NewProvider(os.Getenv("RPC_URL"), ethrpc.WithHeader("x-apikey", os.Getenv("RPC_API_KEY")))
	if err != nil {
		slog.Error("failed dialing into rpc provider", err)
		return
	}

	// Update base project data
	err = sync.Synchronize(ctx, db, rpcClient)
	if err != nil {
		slog.Error("failed to sync contracts", err)
	}

	// Gracefully shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	<-done
}
