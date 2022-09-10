package repository

import (
	"errors"
	"time"

	"app.io/config"
	"app.io/internal/data/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlockRepository struct {
	cfg config.Config
	db  gorm.DB
}

func CreateBlockRepository(cfg config.Config, db gorm.DB) *BlockRepository {
	return &BlockRepository{cfg, db}
}

func (r *BlockRepository) GetBlocks(skip int64, limit int64, status string) ([]domain.Block, error) {
	var blocks []domain.Block
	result := r.db.Model(&domain.Block{}).Find(&blocks, "status = ?", status).Limit(int(limit)).Offset(int(skip))
	if result.Error != nil {
		return nil, errors.New("block list is empty")
	}
	return blocks, nil
}

func (r *BlockRepository) GetBlockById(id string) (*domain.Block, error) {
	result := &domain.Block{}
	query := domain.Block{
		Base: domain.Base{
			ID: id,
		},
	}
	r.db.First(result, query)
	if result.ID == "" {
		return nil, errors.New("such block dose not exits")
	}
	return result, nil
}

func (r *BlockRepository) CreateBlock(data domain.Block) (*domain.Block, error) {
	result := domain.Block{
		Base: domain.Base{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		Hash:    data.Hash,
		Number:  data.Number,
		TxCount: data.TxCount,
	}
	dbTrxOperation := r.db.Create(&result)
	if dbTrxOperation.Error != nil {
		return nil, errors.New("unable to create new block")
	}
	return &result, nil
}

func (r *BlockRepository) UpdateBlock(id string, updatedBlock domain.Block) (*domain.Block, error) {
	block := domain.Block{
		Base: domain.Base{
			ID: id,
		}}
	// exclude
	updatedBlock.ID = ""
	dbTrxOperation := r.db.Model(&block).Updates(updatedBlock).Scan(&block)
	if dbTrxOperation.Error != nil {
		return nil, errors.New("block didn't updated")
	}
	return &block, nil
}

func (r *BlockRepository) SoftDeleteBlock(id string) (*domain.Block, error) {
	block := domain.Block{
		Base: domain.Base{
			ID: id,
		}}
	dbTrxOperation := r.db.Model(&block).Updates(
		map[string]interface{}{
			"deletedAt": time.Now().Format(time.RFC3339),
		},
	).Scan(&block)
	if dbTrxOperation.Error != nil {
		return nil, errors.New("block didn't deleted")
	}
	return &block, nil
}

func (r *BlockRepository) DeleteBlock(id string) (*domain.Block, error) {
	block := domain.Block{
		Base: domain.Base{
			ID: id,
		}}
	dbTrxOperation := r.db.Delete(block).Scan(&block)
	if dbTrxOperation.Error != nil {
		return nil, errors.New("block didn't delete")
	}
	return &block, nil
}
