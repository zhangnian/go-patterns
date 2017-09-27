package main

import (
	"fmt"
	"io"
)

const (
	ST_MEMORY = 1 << iota
	ST_FILE
)

// 实现了该接口的类型都拥有相同的能力
type IStore interface {
	Open(string) (io.ReadWriteCloser, error)
}

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (this *MemoryStore) Open(path string) (io.ReadWriteCloser, error) {
	// 实现了IStore接口
	return nil, nil
}

type FileStore struct{}

func NewFileStore() *FileStore {
	return &FileStore{}
}

func (this *FileStore) Open(path string) (io.ReadWriteCloser, error) {
	// 实现了IStore接口
	return nil, nil
}

func CreateStore(storeType int) IStore {
	switch storeType {
	case ST_MEMORY:
		return NewMemoryStore()
	case ST_FILE:
		return NewFileStore()
	default:
		return nil
	}
}

func main() {
	store := CreateStore(ST_FILE)
	f, err := store.Open("/test")
	fmt.Println(f, err)
}
