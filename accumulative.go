package BitcoinUtxo

import (
	"bitcoingo-utxoselect/common"
)

// add inputs until we reach or surpass the target value (or deplete)
// worst-case: O(n)
func Accumulative(utxos []*common.TxIn, outputs []*common.TxOut, feeRate uint64) ([]*common.TxIn, []*common.TxOut, uint64) {
	var bytesAccum = common.TransactionBytes(nil, outputs)

	var inAccum uint64 = 0
	var inputs []*common.TxIn
	var outAccum = common.OutputsValue(outputs)

	for i := 0; i < len(utxos); i++ {
		var utxo = utxos[i]
		var utxoBytes = common.InputBytes(utxo)
		var utxoFee = feeRate * utxoBytes
		var utxoValue = (utxo.Value)

		// skip detrimental input
		if utxoFee > utxo.Value {
			if i == len(utxos)-1 {
				return nil, nil, feeRate * (bytesAccum + utxoBytes)
			}
			continue
		}

		bytesAccum += utxoBytes
		inAccum += utxoValue
		inputs = append(inputs, utxo)

		var fee = feeRate * bytesAccum

		// go again?
		if inAccum < outAccum+fee {
			continue
		}

		return common.Finalize(inputs, outputs, feeRate)
	}

	return nil, nil, feeRate * bytesAccum
}
