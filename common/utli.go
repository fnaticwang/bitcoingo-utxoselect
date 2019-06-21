package common

// baseline estimates, used to improve performance
var TX_EMPTY_SIZE uint64 = 4 + 1 + 1 + 4
var TX_INPUT_BASE uint64 = 32 + 4 + 1 + 4
var TX_INPUT_PUBKEYHASH uint64 = 107
var TX_OUTPUT_BASE uint64 = 8 + 1
var TX_OUTPUT_PUBKEYHASH uint64 = 25

func InputBytes(input *TxIn) uint64 {
	if input == nil {
		return TX_INPUT_BASE
	}
	if len(input.Script) <= 0 {
		return TX_INPUT_BASE + (uint64)(len(input.Script))
	} else {
		return TX_INPUT_BASE + TX_INPUT_PUBKEYHASH
	}
}

func OutputBytes(output *TxOut) uint64 {
	if output == nil {
		return TX_OUTPUT_BASE
	}
	return TX_OUTPUT_BASE + TX_OUTPUT_PUBKEYHASH
}

func DustThreshold(output *TxIn, feeRate uint64) uint64 {
	/* ... classify the output for input estimate  */
	return InputBytes(nil) * feeRate
}

func TransactionBytes(inputs []*TxIn, outputs []*TxOut) uint64 {
	var input_bytes uint64 = 0
	var output_bytes uint64 = 0
	for _, input := range inputs {
		input_bytes += InputBytes(input)
	}
	for _, output := range outputs {
		output_bytes += OutputBytes(output)
	}
	return TX_EMPTY_SIZE + input_bytes + output_bytes
}

func InputsValue(inputs []*TxIn) uint64 {
	var value uint64 = 0
	for _, input := range inputs {
		value += input.Value
	}
	return value
}

func OutputsValue(outputs []*TxOut) uint64 {
	var value uint64 = 0
	for _, output := range outputs {
		value += output.Value
	}
	return value
}

var BLANK_OUTPUT = OutputBytes(nil)

func Finalize(inputs []*TxIn, outputs []*TxOut, feeRate uint64) ([]*TxIn, []*TxOut, uint64) {
	var bytesAccum = TransactionBytes(inputs, outputs)
	var feeAfterExtraOutput uint64 = feeRate * (bytesAccum + BLANK_OUTPUT)
	var remainderAfterExtraOutput = InputsValue(inputs) - (OutputsValue(outputs) + feeAfterExtraOutput)

	// is it worth a change output?
	if remainderAfterExtraOutput > DustThreshold(nil, feeRate) {
		output := &TxOut{Address: "left", Value: remainderAfterExtraOutput}
		outputs = append(outputs, output)
	}

	var fee = InputsValue(inputs) - OutputsValue(outputs)
	if fee < 0 {
		return nil, nil, feeRate * bytesAccum
	}

	return inputs, outputs, fee
}
