package networkscaning

import (
	"fmt"
	"time"

	"app.io/config"
	"app.io/internal/data/domain"
	"app.io/internal/data/repository"
	"app.io/pkg/logHandler"
	tatumNetworkExporlorer "app.io/pkg/tatum"
)

func SyncNetwork(cfg config.Config, blockState chan int, blockRepo *repository.BlockRepository, transactionRepo *repository.TransactionRepository) ([]domain.Block, error) {
	logHandler.Log(logHandler.INFO, "Start background process")
	tmpBlockList := []domain.Block{}
	networkList := []string{cfg.Server.NetworkTitle}
	for _, network := range networkList {
		//
		// Current block number
		//
		currentBlockNumber, err := tatumNetworkExporlorer.GetCurrentBlockNumber(network, cfg.Server.TatumApiToken)
		if err != nil {
			return nil, err
		}
		blockState <- currentBlockNumber
		logHandler.Log(logHandler.INFO, fmt.Sprintf("latest block %d", blockState))

		//
		// Get "cfg.Server.NumberOfBlockForCapturing" last blocks of network
		//
		for index := currentBlockNumber; index > currentBlockNumber-cfg.Server.NumberOfBlockForCapturing; index-- {
			time.Sleep(5 * time.Second)
			// logHandler.Log(logHandler.INFO, fmt.Sprintf("block state %d", blockState))
			blockState <- index
			block, trxList := tatumNetworkExporlorer.GetBlockData(network, index, cfg.Server.TatumApiToken)
			// logHandler.Log(logHandler.INFO, fmt.Sprintf("block %v", block))
			// logHandler.Log(logHandler.INFO, fmt.Sprintf("trx list %v", trxList))
			blockInstance := domain.Block{
				TxCount: int64(len(trxList)),
				Hash:    block.Hash,
				Number:  int64(block.Number),
			}
			tmpBlockList = append(tmpBlockList, block)
			// Store blocks data
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
				// Store transactions
				transactionRepo.CreateTransaction(trxInstance)
			}
		}
	}
	logHandler.Log(logHandler.INFO, "End of background process")
	return tmpBlockList, nil
}
