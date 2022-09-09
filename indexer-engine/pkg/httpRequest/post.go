package httpRequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Post(address string, values map[string]string, header map[string]string) string {
	jsonData, err := json.Marshal(values)

	req, err := http.NewRequest(http.MethodPost, address, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()

	if header["param_token"] != "" {
		q.Add("token", header["token"])
	}

	req.Header.Set("x-api-key", header["token"])

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return ""
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseBody)
}
