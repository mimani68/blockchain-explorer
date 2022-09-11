package service

import "app.io/data/entity"

type BlockService interface {
	LastBlock() (response entity.Block, err error)
	GetBlockByNumber(blockNumber int) (response entity.Block, err error)
}
