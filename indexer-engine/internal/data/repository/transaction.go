package repository

import (
	"errors"
	"time"

	"app.io/config"
	"app.io/internal/data/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	cfg config.Config
	db  gorm.DB
}

func CreateTransactionRepository(cfg config.Config, db gorm.DB) *TransactionRepository {
	return &TransactionRepository{cfg, db}
}

func (r *TransactionRepository) GetTransactions(skip int64, limit int64, status string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	result := r.db.Model(&domain.Transaction{}).Find(&transactions, "status = ?", status).Limit(int(limit)).Offset(int(skip)).Distinct()
	if result.Error != nil {
		return nil, errors.New("transaction list is empty")
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTransactionById(id string) (*domain.Transaction, error) {
	result := &domain.Transaction{}
	query := domain.Transaction{
		Base: domain.Base{
			ID: id,
		},
	}
	r.db.First(result, query)
	if result.ID == "" {
		return nil, errors.New("such transaction dose not exits")
	}
	return result, nil
}

func (r *TransactionRepository) GetTransactionsByBlockNumber(blockNumber int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	result := r.db.Model(&domain.Transaction{}).Find(&transactions, "block_number = ?", blockNumber).Distinct()
	if result.Error != nil {
		return nil, errors.New("transaction list is empty")
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTransactionsSum(startBlockNumber int, endBlockNumber int) (int, error) {
	var trxStats int
	result := r.db.Table("transaction").Select("sum(amount) as result").Where("block_number BETWEEN ? AND ? ", startBlockNumber, endBlockNumber).Distinct().Scan(&trxStats)
	if result.Error != nil {
		return 0, errors.New("transaction stats return invalid value")
	}
	return trxStats, nil
}

func (r *TransactionRepository) GetTransactionsNumber(startBlockNumber int, endBlockNumber int) (int, error) {
	var trxStats int
	result := r.db.Table("transaction").Select("count(amount) as result").Where("block_number BETWEEN ? AND ? ", startBlockNumber, endBlockNumber).Distinct().Scan(&trxStats)
	if result.Error != nil {
		return 0, errors.New("transaction stats return invalid value")
	}
	return trxStats, nil
}

func (r *TransactionRepository) CreateTransaction(data domain.Transaction) (*domain.Transaction, error) {
	result := domain.Transaction{
		Base: domain.Base{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		Hash:        data.Hash,
		BlockNumber: data.BlockNumber,
		Amount:      data.Amount,
		Nonce:       data.Nonce,
		From:        data.From,
		To:          data.To,
	}
	dbTrxOperation := r.db.Create(&result)
	if dbTrxOperation.Error != nil {
		return nil, errors.New("unable to create new transaction")
	}
	return &result, nil
}

func (r *TransactionRepository) UpdateTransaction(id string, updatedTransaction domain.Transaction) (*domain.Transaction, error) {
	transaction := domain.Transaction{
		Base: domain.Base{
			ID: id,
		}}
	// exclude
	updatedTransaction.ID = ""
	dbTrxOperation := r.db.Model(&transaction).Updates(updatedTransaction).Scan(&transaction)
	if dbTrxOperation.Error != nil {
		return nil, errors.New("transaction didn't updated")
	}
	return &transaction, nil
}

func (r *TransactionRepository) SoftDeleteTransaction(id string) (*domain.Transaction, error) {
	transaction := domain.Transaction{
		Base: domain.Base{
			ID: id,
		}}
	dbTrxOperation := r.db.Model(&transaction).Updates(
		map[string]interface{}{
			"deletedAt": time.Now().Format(time.RFC3339),
		},
	).Scan(&transaction)
	if dbTrxOperation.Error != nil {
		return nil, errors.New("transaction didn't deleted")
	}
	return &transaction, nil
}

func (r *TransactionRepository) DeleteTransaction(id string) (*domain.Transaction, error) {
	transaction := domain.Transaction{
		Base: domain.Base{
			ID: id,
		}}
	dbTrxOperation := r.db.Delete(transaction).Scan(&transaction)
	if dbTrxOperation.Error != nil {
		return nil, errors.New("transaction didn't delete")
	}
	return &transaction, nil
}
