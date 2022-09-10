package job

import (
	"fmt"
	"time"

	"app.io/config"
	"app.io/internal/data/domain"
	"app.io/internal/data/repository"
	"app.io/pkg/logHandler"
	tatumNetworkExporlorer "app.io/pkg/tatum"
)

func SyncNetwork(cfg config.Config, blockRepo *repository.BlockRepository, transactionRepo *repository.TransactionRepository) {
	logHandler.Log(logHandler.INFO, "Start background process")
	blockState := 0
	networkList := []string{cfg.Server.NetworkTitle}
	for _, network := range networkList {
		//
		// Current block number
		//
		currentBlockNumber, err := tatumNetworkExporlorer.GetCurtentBlockNumber(network, cfg.Server.TatumApiToken)
		if err != nil {
			return
		}
		blockState = currentBlockNumber
		logHandler.Log(logHandler.INFO, fmt.Sprintf("latest block %d", blockState))

		//
		// Get 5 block of a network
		//
		for index := currentBlockNumber; index > currentBlockNumber-2; index-- {
			time.Sleep(5 * time.Second)
			// logHandler.Log(logHandler.INFO, fmt.Sprintf("block state %d", blockState))
			blockState = index
			block, trxList := tatumNetworkExporlorer.GetBlockData(network, blockState, cfg.Server.TatumApiToken)
			// logHandler.Log(logHandler.INFO, fmt.Sprintf("block %v", block))
			// logHandler.Log(logHandler.INFO, fmt.Sprintf("trx list %v", trxList))
			blockInstance := domain.Block{
				TxCount: int64(len(trxList)),
				Hash:    block.Hash,
				Number:  int64(block.Number),
			}
			blockRepo.CreateBlock(blockInstance)
			for _, trx := range trxList {
				trxInstance := domain.Transaction{
					Hash:        trx.Hash,
					BlockNumber: int64(trx.BlockNumber),
					From:        trx.From,
					To:          trx.To,
					Amount:      int64(trx.Amount),
					Nonce:       int64(trx.Nonce),
				}
				transactionRepo.CreateTransaction(trxInstance)
			}
		}
	}
	logHandler.Log(logHandler.INFO, "End of background process")
}
