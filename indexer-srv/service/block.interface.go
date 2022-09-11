package service

import "app.io/data/entity"

type BlockService interface {
	ForceScan(startBlock, endBlock int) (response []entity.Block, err error)
	LastBlock() (response entity.Block, err error)
	GetBlockByNumber(blockNumber int) (response entity.Block, err error)
}
