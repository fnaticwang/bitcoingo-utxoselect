package BitcoinUtxo

import (
	"BitcoinUtxo/common"
)

func Blackjack(utxos []*common.TxIn, outputs []*common.TxOut, feeRate uint64) ([]*common.TxIn, []*common.TxOut, uint64) {
	var bytesAccum = common.TransactionBytes(nil, outputs)

	var inAccum uint64 = 0
	var inputs []*common.TxIn
	var outAccum = common.OutputsValue(outputs)
	var threshold = common.DustThreshold(nil, feeRate)

	for i := 0; i < len(utxos); i++ {
		var input = utxos[i]
		var inputBytes = common.InputBytes(input)
		var fee = feeRate * (bytesAccum + inputBytes)
		var inputValue = input.Value

		// would it waste value?
		if (inAccum + inputValue) > (outAccum + fee + threshold) {
			continue
		}

		bytesAccum += inputBytes
		inAccum += inputValue
		inputs = append(inputs, input)

		// go again?
		if inAccum < outAccum+fee {
			continue
		}

		return common.Finalize(inputs, outputs, feeRate)
	}

	return nil, nil, feeRate * bytesAccum
}
