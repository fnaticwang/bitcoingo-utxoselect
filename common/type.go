package common

type TxIn struct {
	Hash    string
	Index   uint32
	Script  string
	PrivKey string
	Value   uint64 //0.00000001BTC(satoshi)
}

type TxOut struct {
	Address string
	Value   uint64 //0.00000001BTC(satoshi)
}
