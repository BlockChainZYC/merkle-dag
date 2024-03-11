package merkledag

import (
	"hash"
)

type Link struct {
	Name string
	Hash []byte
	Size int
}

type Object struct {
	Links []Link
	Data  []byte
}

func Add(store KVStore, node Node, h hash.Hash) ([]byte, error) {
	// 获取节点数据
	data := node.Data()

	// 计算数据的哈希值
	h.Write(data)
	merkleRoot := h.Sum(nil)

	// 将数据和哈希值存储到KVStore中
	err := store.Put(merkleRoot, data)
	if err != nil {
		return nil, err
	}

	// 返回Merkle Root
	return merkleRoot, nil
}
