package main

import (
	"crypto/sha256"
	"encoding/hex"
	out "fmt"
	"time"
)

// Block : Simple part for a chain of blocks called a blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      User
	Hash      string
	PrevHash  string
}

// Permission : Permission struct holding informations about the permission
type Permission struct {
	ID     int
	Access string
}

// User : User struct holding informations about the user
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Permissions []Permission
}

func mains() {
	var Blockchain []Block

	s := User{
		0,
		"Kevin Jens",
		"Schaefer",
		[]Permission{
			Permission{
				0,
				"ACCESS"},
			Permission{
				1,
				"R&D",
			},
			Permission{
				2,
				"BOARD",
			},
			Permission{
				13,
				"UNICORN",
			},
			Permission{
				17,
				"SVP",
			},
			Permission{
				42,
				"SUPERUSER",
			},
		},
	}

	t := time.Now()
	genesisBlock := Block{0, t.String(), s, "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)
	genesisBlock.PrevHash = "Nil"
	Blockchain = append(Blockchain, genesisBlock)

	josef := User{
		1,
		"Josef",
		"Büttgen",
		[]Permission{
			Permission{
				0,
				"ACCESS"},
			Permission{
				1,
				"R&D",
			},
		},
	}

	secondBlock, _ := generateBlock(genesisBlock, josef)
	Blockchain = append(Blockchain, secondBlock)

	constantin := User{
		2,
		"Constantin Tiberius",
		"Tóth",
		[]Permission{
			Permission{
				0,
				"ACCESS"},
		},
	}

	thirdBlock, _ := generateBlock(secondBlock, constantin)
	Blockchain = append(Blockchain, thirdBlock)

	out.Println(Blockchain)
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data.FirstName + block.Data.LastName + string(block.Data.ID) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Data User) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
