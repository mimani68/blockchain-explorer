package tatumNetworkExporlorer

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"app.io/internal/data/domain"
	"app.io/internal/data/dto"
	"app.io/pkg/httpRequest"
)

type blockPayloadWrapper struct {
	Transactions []dto.TatumTransaction `json:"transactions"`
}

func GetCurrentBlockNumber(network, vendorToken string) (currentNumber int, err error) {
	// curl https://api-eu1.tatum.io/v3/ethereum/block/current
	url := fmt.Sprintf("https://api-eu1.tatum.io/v3/%s/block/current", network)
	header := map[string]string{
		"token": vendorToken,
	}
	tmp := httpRequest.Get(url, header)
	if len(tmp) <= 0 {
		err = errors.New("Current number is empty")
	}
	currentNumber, err = strconv.Atoi(tmp)
	return
}

func GetBlockData(network string, blockNumber int, vendorToken string) (block domain.Block, trx []domain.Transaction) {
	// curl "https://api-eu1.tatum.io/v3/bitcoin/block/{ blockNumber }"
	url := fmt.Sprintf("https://api-eu1.tatum.io/v3/%s/block/%d", network, blockNumber)
	header := map[string]string{
		"token": vendorToken,
	}
	blockResponse := httpRequest.Get(url, header)
	block = domain.Block{}
	json.Unmarshal([]byte(blockResponse), &block)
	tmp := blockPayloadWrapper{}
	json.Unmarshal([]byte(blockResponse), &tmp)
	for _, item := range tmp.Transactions {
		p := domain.Transaction{
			BlockNumber: int64(item.BlockNumber),
			Hash:        item.Hash,
			From:        item.From,
			To:          item.To,
			Nonce:       int64(item.Nonce),
		}
		amount, _ := strconv.Atoi(item.Amount)
		p.Amount = int64(amount)
		trx = append(trx, p)
	}

	return
}
