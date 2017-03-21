package ssdeep

import (
	"bytes"
	"fmt"
	"io"
)

var rollingWindow uint32 = 7
var blockSizeMin = 3

const spamSumLength = 64

var b64String = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var b64 = []byte(b64String)
var hashPrime uint32 = 0x01000193
var hashIinit uint32 = 0x28021967

type rollingState struct {
	window []byte
	h1     uint32
	h2     uint32
	h3     uint32
	n      uint32
}

// FuzzyHash struct for comparison
type FuzzyHash struct {
	blockSize   int
	hashString1 string
	hashString2 string
}

// SSDEEP state struct
type SSDEEP struct {
	rollingState rollingState
	blockSize    int
	hashString1  string
	hashString2  string
	blockHash1   uint32
	blockHash2   uint32
}

// rollHash based on Adler checksum
func (sdeep *SSDEEP) rollHash(c byte) uint32 {
	rs := &sdeep.rollingState
	rs.h2 -= rs.h1
	rs.h2 += rollingWindow * uint32(c)
	rs.h1 += uint32(c)
	rs.h1 -= uint32(rs.window[rs.n])
	rs.window[rs.n] = c
	rs.n++
	if rs.n == rollingWindow {
		rs.n = 0
	}
	rs.h3 = rs.h3 << 5
	rs.h3 ^= uint32(c)
	return rs.h1 + rs.h2 + rs.h3
}

func (sdeep *SSDEEP) processByte(b byte) {
	sdeep.blockHash1 = sumHash(b, sdeep.blockHash1)
	sdeep.blockHash2 = sumHash(b, sdeep.blockHash2)
	rh := int(sdeep.rollHash(b))
	if rh%sdeep.blockSize == (sdeep.blockSize - 1) {
		if len(sdeep.hashString1) < spamSumLength-1 {
			sdeep.hashString1 += string(b64[sdeep.blockHash1%64])
			sdeep.blockHash1 = hashIinit
		}
	}
	if rh%(sdeep.blockSize*2) == ((sdeep.blockSize * 2) - 1) {
		if len(sdeep.hashString2) < spamSumLength/2-1 {
			sdeep.hashString2 += string(b64[sdeep.blockHash2%64])
			sdeep.blockHash2 = hashIinit
		}
	}
}

// initSSDEEP creates a new SSDEEP hash
func initSSDEEP(blockSize int) *SSDEEP {
	return &SSDEEP{
		blockHash1: hashIinit,
		blockHash2: hashIinit,
		rollingState: rollingState{
			window: make([]byte, rollingWindow),
		},
		hashString1: "",
		hashString2: "",
		blockSize:   blockSize,
	}
}

// sumHash based on FNV hash
func sumHash(c byte, h uint32) uint32 {
	return (h * hashPrime) ^ uint32(c)
}

func initBlockSize(n int) int {
	blockSize := blockSizeMin
	for blockSize*spamSumLength < n {
		blockSize = blockSize * 2
	}
	return blockSize
}

func calcSpamSum(buf *bytes.Buffer, blockSize int) *FuzzyHash {
	sdeep := initSSDEEP(blockSize)
	for {
		b := buf.Bytes()
		for i := range b {
			sdeep.processByte(b[i])
		}
		// Finalize the hash string with the remaining data
		sdeep.hashString1 += string(b64[sdeep.blockHash1%64])
		sdeep.hashString2 += string(b64[sdeep.blockHash2%64])

		// wrong blockSize
		if (len(sdeep.hashString1) < spamSumLength/2) && (sdeep.blockSize > blockSizeMin) {
			sdeep = initSSDEEP(sdeep.blockSize / 2) // divide blocksize, and initialize ssdeep again
		} else {
			break
		}
	}

	return &FuzzyHash{
		sdeep.blockSize,
		sdeep.hashString1,
		sdeep.hashString2,
	}
}

// Fuzzy hash of a provided reader
func Fuzzy(r io.Reader) (*FuzzyHash, error) {
	buf := &bytes.Buffer{}
	n, err := io.Copy(buf, r)
	if err != nil {
		return nil, err
	}
	fuzzy := calcSpamSum(buf, initBlockSize(int(n)))

	// fmt.Printf("%d:%s:%s\n", fuzzy.blockSize, fuzzy.hashString1, fuzzy.hashString2)

	return fuzzy, nil
}

func (h FuzzyHash) String() string {
	return fmt.Sprintf("%d:%s:%s", h.blockSize, h.hashString1, h.hashString2)
}
