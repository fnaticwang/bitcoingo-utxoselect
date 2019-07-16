package BitcoinUtxo

import (
	"bitcoingo-utxoselect/common"
	"sort"
)

type TxInSlice []*common.TxIn

func (a TxInSlice) Len() int {
	return len(a)
}
func (a TxInSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a TxInSlice) Less(i, j int) bool {
	return a[j].Value < a[i].Value
}

func CoinSelect(utxos []*common.TxIn, outputs []*common.TxOut, feeRate uint64) ([]*common.TxIn, []*common.TxOut, uint64) {
	sort.Sort(TxInSlice(utxos))

	// attempt to use the blackjack strategy first (no change output)
	in, out, fee := Blackjack(utxos, outputs, feeRate)
	if in != nil {
		return in, out, fee
	}
	// else, try the accumulative strategy
	return Accumulative(utxos, outputs, feeRate)
}
