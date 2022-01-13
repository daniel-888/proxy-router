module example.com/block_header

go 1.16

replace example.com/wire => ../wire

require (
	example.com/chainhash v0.0.0-00010101000000-000000000000
	example.com/wire v0.0.0-00010101000000-000000000000
)

replace example.com/chainhash => ../chainhash

replace example.com/hashing => ../hashing

replace example.com/blockchain => ../blockchain
