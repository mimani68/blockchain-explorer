package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"app.io/data/entity"
	"app.io/data/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionController_Create(t *testing.T) {
	transactionRepository.DeleteAll()
	createTransactionRequest := model.CreateTransactionRequest{
		Name:     "Test Transaction",
		Price:    1000,
		Quantity: 1000,
	}
	requestBody, _ := json.Marshal(createTransactionRequest)

	request := httptest.NewRequest("POST", "/api/transactions", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createTransactionResponse := model.CreateTransactionResponse{}
	json.Unmarshal(jsonData, &createTransactionResponse)
	assert.NotNil(t, createTransactionResponse.Id)
	assert.Equal(t, createTransactionRequest.Name, createTransactionResponse.Name)
	assert.Equal(t, createTransactionRequest.Price, createTransactionResponse.Price)
	assert.Equal(t, createTransactionRequest.Quantity, createTransactionResponse.Quantity)
}

func TestTransactionController_List(t *testing.T) {
	transactionRepository.DeleteAll()
	transaction := entity.Transaction{
		Id:       uuid.New().String(),
		Name:     "Sample Transaction",
		Price:    1000,
		Quantity: 1000,
	}
	transactionRepository.Insert(transaction)

	request := httptest.NewRequest("GET", "/api/transactions", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.([]interface{})
	containsTransaction := false

	for _, data := range list {
		jsonData, _ := json.Marshal(data)
		getTransactionResponse := model.GetTransactionResponse{}
		json.Unmarshal(jsonData, &getTransactionResponse)
		if getTransactionResponse.Id == transaction.Id {
			containsTransaction = true
		}
	}

	assert.True(t, containsTransaction)
}
