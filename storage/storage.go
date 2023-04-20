package storage

import (
	"github.com/petrick-ribeiro/go-todo/types"
)

type Storage interface {
	GetAll() ([]*types.Todo, error)
	Get(uint64) (*types.Todo, error)
	Insert(*types.Todo) error
	Update(*types.Todo, uint64) (*types.Todo, error)
	Delete(uint64) (*types.Todo, error)
}
