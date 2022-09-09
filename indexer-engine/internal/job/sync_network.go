package job

import (
	"fmt"
	"time"

	"app.io/config"
	tatumNetworkExporlorer "app.io/pkg/tatum"
)

func SyncNetwork(cfg config.Config) {

	// blockRespoitory{}
	// TtransctionRepostioy{}

	// networkList := []string{"bitcoin", "ethereum", "xrp", "ltc"}
	blockState := 0
	networkList := []string{"ethereum"}
	// networkList := []string{"bitcoin"}
	for _, network := range networkList {
		//
		// Current block number
		//
		currentBlockNumber, err := tatumNetworkExporlorer.GetCurtentBlockNumber(network, cfg.Server.TatumApiToken)
		if err != nil {
			return
		}
		blockState = currentBlockNumber
		fmt.Printf("latest block %d \n", blockState)

		//
		// Get 5 block of a network
		//
		for index := currentBlockNumber; index > currentBlockNumber-2; index-- {
			time.Sleep(5 * time.Second)
			fmt.Printf("block state %d \n", blockState)
			blockState = index
			block, trxList := tatumNetworkExporlorer.GetBlockData(network, blockState, cfg.Server.TatumApiToken)
			fmt.Printf("block %v \n", block)
			fmt.Printf("trx list %v \n", trxList)
			// create new payload => { number, hash, txCount }
			// store in db
			// sleep 2 sec
		}
	}
}
