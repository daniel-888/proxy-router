package blockHeader

import "testing"
import "fmt"

//each test serializes the information in the blockheader with the provided nonce
//then the result is double hashed Sha256(Sha256(serialization))
//then the result is compared to the blockHash result

func TestBlock1(t *testing.T) {
	blockHash := "000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506"

	nonce := "274148111"
	time := "4d1b2237"

	blockHeader := BlockHeader{
		Version:           "00000001",
		MerkleRoot:        "f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766",
		Time:              "4d1b2237",
		Difficulty:        "1b04864c",
		PreviousBlockHash: "000000000002d01c1fccc21636b607dfd930d31d01c3a62104612a1719011250",
	}

	resultHash := blockHeader.HashInput(nonce, time)

	if fmt.Sprintf("%x", resultHash) != blockHash {
		t.Error(fmt.Printf("expected blockhash: %s\nCalculated blockhash: %x\n", blockHash, resultHash))
	}
}

func TestBlock2(t *testing.T) {
	blockHash := "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf"
	nonce := "4158183488"
	time := "505d96e7"
	blockHeader := BlockHeader{
		Version:           "00000002",
		MerkleRoot:        "a08f8101f50fd9c9b3e5252aff4c1c1bd668f878fffaf3d0dbddeb029c307e88",
		Time:              "1348310759",
		Difficulty:        "1a05db8b",
		PreviousBlockHash: "00000000000003a20def7a05a77361b9657ff954b2f2080e135ea6f5970da215",
	}

	resultHash := blockHeader.HashInput(nonce, time)

	if fmt.Sprintf("%x", resultHash) != blockHash {
		t.Error(fmt.Printf("expected blockhash: %s\nCalculated blockhash: %s\n", blockHash, resultHash))
	}
}

func TestBlock3(t *testing.T) {
	blockHash := "000000000000000082ccf8f1557c5d40b21edabb18d2d691cfbf87118bac7254"
	nonce := "222771801"
	time := "536dc802"
	blockHeader := BlockHeader{
		Version:           "00000002",
		MerkleRoot:        "915c887a2d9ec3f566a648bedcf4ed30d0988e22268cfe43ab5b0cf8638999d3",
		Time:              "1399703554",
		Difficulty:        "1900896c",
		PreviousBlockHash: "000000000000000067ecc744b5ae34eebbde14d21ca4db51652e4d67e155f07e",
	}

	resultHash := blockHeader.HashInput(nonce, time)

	if fmt.Sprintf("%x", resultHash) != blockHash {
		t.Error(fmt.Printf("expected blockhash: %s\nCalculated blockhash: %x\n", blockHash, resultHash))
	}
}

func TestBlock4(t *testing.T) {
	blockHash := "000000000000000004ec466ce4732fe6f1ed1cddc2ed4b328fff5224276e3f6f"
	nonce := "657220870"
	time := "56cf2acc"
	blockHeader := BlockHeader{
		Version:           "00000004",
		MerkleRoot:        "b0e8f88d4fb7cbc49ab49a3a43c368550e22a8e9e3e04b15e34240306a53aeec",
		Time:              "1456417484",
		Difficulty:        "1806b99f",
		PreviousBlockHash: "0000000000000000030034b661aed920a9bdf6bbfa6d2e7a021f78481882fa39",
	}

	resultHash := blockHeader.HashInput(nonce, time)

	if fmt.Sprintf("%x", resultHash) != blockHash {
		t.Error(fmt.Printf("expected blockhash: %s\nCalculated blockhash: %x\n", blockHash, resultHash))
	}
}

//values copied from https://en.bitcoin.it/wiki/Block_hashing_algorithm but with prev-block hash and merkle little endian, while all other values are big endian
/*
func TestBlock5(t *testing.T) {
	blockHash := "00000000000000001e8d6829a8a21adc5d38d0a473b144b6765798e61f98bd1d"
	  nonce := "2504433986" //int
	time := "4dd7f5c7" //big-endian
	blockHeader := BlockHeader {
		//bitcoin-cli returns all values as big-endian
		Version: "00000001", //big-endian
		MerkleRoot: "e320b6c2fffc8d750423db8b1eb942ae710e951ed797f7affc8892b0f1fc122b", //little-endian
		Time: "4dd7f5c7", //big-endian
		Difficulty: "1a44b9f2", //big-endian
		PreviousBlockHash: "81cd02ab7e569e8bcd9317e2fe99f2de44d49ab2b8851ba4a308000000000000", //little-endian
	}

	resultHash := blockHeader.HashInput(nonce, time)

	if fmt.Sprintf("%x", resultHash) != blockHash {
		t.Error(fmt.Printf("expected blockhash: %s\nCalculated blockhash: %x\n", blockHash, resultHash))
	}
}
*/
