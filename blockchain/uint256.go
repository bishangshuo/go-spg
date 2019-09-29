package blockchain

const (
	BITS uint32 = 256
	WIDTH uint32 = BITS/8;
)

type base_blob struct{
	data [WIDTH] uint8
}

