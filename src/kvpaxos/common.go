package kvpaxos

import (
	"crypto/rand"
	"hash/fnv"
	"math/big"
	"time"
)

const (
	OK       = "OK"
	ErrNoKey = "ErrNoKey"

	PUT     = "PUT"
	PUTHASH = "PUTHASH"
	GET     = "GET"
)

type Err string

type PutArgs struct {
	// You'll have to add definitions here.
	Key    string
	Value  string
	DoHash bool // For PutHash
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
	Id       int64
	ClientId int64
}

type PutReply struct {
	Err           Err
	PreviousValue string // For PutHash
}

type GetArgs struct {
	Key string
	// You'll have to add definitions here.
	Id       int64
	ClientId int64
}

type GetReply struct {
	Err   Err
	Value string
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func nrand() int64 {
	max := big.NewInt(int64(1) << 62)
	bigx, _ := rand.Int(rand.Reader, max)
	x := bigx.Int64()
	return x
}

func currTime() int64 {
	// https://golangbyexample.com/current-timestamp-in-golang/
	return time.Now().UnixNano()
}
