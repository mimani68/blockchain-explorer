package job

import (
	"fmt"
	"time"

	"app.io/config"
	"app.io/internal/data/db"
	"app.io/pkg/logHandler"
	tatumNetworkExporlorer "app.io/pkg/tatum"
)

func SyncNetwork(cfg config.Config, db db.Database) {

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
		logHandler.Log(logHandler.INFO, fmt.Sprintf("latest block %d \n", blockState))

		//
		// Get 5 block of a network
		//
		for index := currentBlockNumber; index > currentBlockNumber-2; index-- {
			time.Sleep(5 * time.Second)
			logHandler.Log(logHandler.INFO, fmt.Sprintf("block state %d \n", blockState))
			blockState = index
			block, trxList := tatumNetworkExporlorer.GetBlockData(network, blockState, cfg.Server.TatumApiToken)
			logHandler.Log(logHandler.INFO, fmt.Sprintf("block %v \n", block))
			logHandler.Log(logHandler.INFO, fmt.Sprintf("trx list %v \n", trxList))
			// create new payload => { number, hash, txCount }
			// store in db
			// sleep 2 sec
		}
	}
}
