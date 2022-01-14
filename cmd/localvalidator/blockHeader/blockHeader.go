package blockHeader

import (
	"encoding/binary"
	"encoding/json"
	"example.com/chainhash"
	"example.com/wire"
	"github.com/btcsuite/btcd/blockchain"
	chainhashOnline "github.com/btcsuite/btcd/chaincfg/chainhash"
	"fmt"
	"strconv"
	"math/big"
)

type BlockHeader struct {
	Version           string
	PreviousBlockHash string
	MerkleRoot        string
	Time              string
	Difficulty        string
}

//expects a string of the form `"Version": "001"`... etc to parse as a JSON
func ConvertToBlockHeader(message string) BlockHeader {
	// string will look like a JSON object
	// first convert the string into a map
	// return a BlockHeader object using map values
	var bi map[string]string             //create an empty map to put string variables into
	json.Unmarshal([]byte(message), &bi) //unmarshal string and put into bi (block info) map
	return BlockHeader{
		Version:           bi["Version"],
		PreviousBlockHash: bi["PreviousBlockHash"],
		MerkleRoot:        bi["MerkleRoot"],
		Time:              bi["Time"],
		Difficulty:        bi["Difficulty"],
	}

}

func ConvertBlockHeaderToString(h BlockHeader) string {
	return fmt.Sprintf(`{\"Version\":\"%s\",\"PreviousBlockHash\":\"%s\",\"MerkleRoot\":\"%s\",\"Time\":\"%s\",\"Difficulty\":\"%s\"}`, h.Version, h.PreviousBlockHash, h.MerkleRoot, h.Time, h.Difficulty)
}

func uintToLittleEndian(x string) string {
	u, _ := strconv.ParseUint(x, 10, 64) //convert string to uint64
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, u)
	return fmt.Sprintf("%x", buf[:4])
}

func reverseHexNumber(x string) [32]byte {
	newNum := ""
	for i := 0; i < len(x); i = i + 2 {
		newNum = x[i:i+2] + newNum
	}
	//pass newNum to NewHashFromString
	res := chainhash.NewHashFromStr(newNum)
	return res
}

// takes a given nonce, and serialize block header into form used for hashing
func (bh *BlockHeader) HashInput(nonce string, time string) [32]byte {
	sVersion, _ := strconv.ParseInt(bh.Version, 16, 32)
	sTime, _ := strconv.ParseInt(time, 16, 32)
	sDifficulty, _ := strconv.ParseInt(bh.Difficulty, 16, 32)
	sNonce, _ := strconv.Atoi(nonce)

	//PrevBlock and MerkleRoot need to be little-endian
	newBlockHash := wire.BlockHeader{
		Version:    int32(sVersion),
		PrevBlock:  chainhash.NewHashFromStr(bh.PreviousBlockHash),
		MerkleRoot: chainhash.NewHashFromStr(bh.MerkleRoot),
		Timestamp: int32(sTime),
		Bits:      uint32(sDifficulty),
		Nonce:     uint32(sNonce),
	}
	hash := newBlockHash.BlockHash() //little-endian

	//converting the resulting hash to big-endian format
	for i, j := 0, len(hash)-1; i < j; i, j = i+1, j-1 {
		hash[i], hash[j] = hash[j], hash[i]
	}
	return hash

}

func (bh *BlockHeader) UpdateHeaderInformation(_version string, _previousBlockHash string, _merkleRoot string, _time string, _difficulty string) {
	bh.Version = _version
	bh.PreviousBlockHash = _previousBlockHash
	bh.MerkleRoot = _merkleRoot
	bh.Time = _time
	bh.Difficulty = _difficulty
}

func BlockHashToBigInt(hash [32]byte) *big.Int {
	//convert input to chainhash.Hash
	chash, _ := chainhashOnline.NewHash(hash[:])
	return blockchain.HashToBig(chash)
}
//going to be getting the same message as described in stratum
