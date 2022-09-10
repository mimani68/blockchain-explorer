package networkscaning

import (
	"fmt"
	"time"

	"app.io/config"
	"app.io/internal/data/domain"
	"app.io/internal/data/dto"
	"app.io/internal/data/repository"
	"app.io/pkg/logHandler"
	tatumNetworkExporlorer "app.io/pkg/tatum"
)

func SyncBlocks(startBlock, endBlock int, cfg config.Config, blockRepo *repository.BlockRepository, transactionRepo *repository.TransactionRepository) ([]dto.CompleteBlock, error) {
	logHandler.Log(logHandler.INFO, "Start background process")
	result := []dto.CompleteBlock{}
	networkList := []string{cfg.Server.NetworkTitle}
	for _, network := range networkList {
		if endBlock == 0 {
			//
			// Current block number
			//
			currentBlockNumber, err := tatumNetworkExporlorer.GetCurrentBlockNumber(network, cfg.Server.TatumApiToken)
			if err != nil || currentBlockNumber <= 0 {
				return nil, err
			}
			endBlock = currentBlockNumber
			logHandler.Log(logHandler.INFO, fmt.Sprintf("latest block %d", currentBlockNumber))
		}
		if startBlock == 0 {
			startBlock = endBlock - cfg.Server.NumberOfBlockForCapturing
		}
		//
		// Get "cfg.Server.NumberOfBlockForCapturing" last blocks of network
		//
		for index := endBlock; index > endBlock-startBlock; index-- {
			time.Sleep(5 * time.Second)
			block, trxList := tatumNetworkExporlorer.GetBlockData(network, index, cfg.Server.TatumApiToken)
			blockInstance := domain.Block{
				TxCount: int64(len(trxList)),
				Hash:    block.Hash,
				Number:  int64(block.Number),
			}
			completeBlockInstance := dto.CompleteBlock{
				TxCount: int64(len(trxList)),
				Hash:    block.Hash,
				Number:  int64(block.Number),
			}
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
				completeBlockInstance.Transaction = append(completeBlockInstance.Transaction, trx)

				// Store transactions
				transactionRepo.CreateTransaction(trxInstance)
			}
			result = append(result, completeBlockInstance)
		}
	}
	logHandler.Log(logHandler.INFO, "End of background process")
	return result, nil
}
