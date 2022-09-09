package blockchairNetworkExporlorer

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"app.io/pkg/httpRequest"
)

func GetCurtentBlockNumber(netwrok string) (currentNumber int, err error) {
	// curl https://api.blockchair.com/bitcoin/stats
	url := "https://api.blockchair.com/" + netwrok + "/stats"
	result := httpRequest.Get(url, map[string]string{})
	regexPattern := regexp.MustCompile(`(?m)\"blocks\":\s{0,10}[a-zA-Z0-9]{2,100}`)
	currentNumberString := regexPattern.FindString(result)
	currentNumberString = strings.ReplaceAll(currentNumberString, "\"blocks\":", "")
	if len(currentNumberString) > 1 {
		currentNumber, err = strconv.Atoi(currentNumberString)
		return
	}
	return 0, errors.New("unable to find block number")
}

func GetBlockData(netwrok string, blockNumber int) (result string) {
	// curl "https://api.blockchair.com/bitcoin/dashboards/block/753319"
	url := fmt.Sprintf("https://api.blockchair.com/%s/dashboards/block/%d", netwrok, blockNumber)
	result = httpRequest.Get(url, map[string]string{})
	return
}

func GetTranactionData(netwrok string, hashList []string) (result string) {
	// curl https://api.blockchair.com/bitcoin/dashboards/transaction/{:hash}₀,..,{:hash}n
	hashListString := strings.Join(hashList, ",")
	url := fmt.Sprintf("https://api.blockchair.com/%s/dashboards/transactions/%s", netwrok, hashListString)
	result = httpRequest.Get(url, map[string]string{})
	return
}
