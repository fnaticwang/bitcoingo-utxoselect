package main

import (
	"bitcoingo-utxoselect"
	"bitcoingo-utxoselect/common"
	"fmt"
)

func main() {
	txIn := make([]*common.TxIn, 4)
	tx_in1 := &common.TxIn{
		Hash:    "42f86d97eb322b6d8e18f7ffe69badf646f256baed87138529bd8ddb47d4f2a0",
		Index:   1,
		Script:  "76a91446ecb23cfe0c1ab11cfb55292686b7d23e88e27688ac",
		PrivKey: "93EaG3dMcLXkSVLLvnRCEKJWRfJz6HvHK5XzTCUYTYQbmhMau1S",
		Value:   15000000,
	}

	tx_in2 := &common.TxIn{
		Hash:    "42f86d97eb322b6d8e18f7ffe69badf646f256baed87138529bd8ddb47d4f2a0",
		Index:   2,
		Script:  "76a91426b15e19fcf76b32e4db4e7f859634e631ecb50088ac",
		PrivKey: "93QEUhGD8BNc5jL4AwF46cdUXDhcmdVqr5vc4kWipfNJCwrFfBE",
		Value:   25000000,
	}

	tx_in3 := &common.TxIn{
		Hash:    "42f86d97eb322b6d8e18f7ffe69badf646f256baed87138529bd8ddb47d4f2a0",
		Index:   3,
		Script:  "76a91426b15e19fcf76b32e4db4e7f859634e631ecb50088ac",
		PrivKey: "93QEUhGD8BNc5jL4AwF46cdUXDhcmdVqr5vc4kWipfNJCwrFfBE",
		Value:   35000000,
	}

	tx_in4 := &common.TxIn{
		Hash:    "42f86d97eb322b6d8e18f7ffe69badf646f256baed87138529bd8ddb47d4f2a0",
		Index:   4,
		Script:  "76a91426b15e19fcf76b32e4db4e7f859634e631ecb50088ac",
		PrivKey: "93QEUhGD8BNc5jL4AwF46cdUXDhcmdVqr5vc4kWipfNJCwrFfBE",
		Value:   45000000,
	}

	txIn[0] = tx_in1
	txIn[1] = tx_in2
	txIn[2] = tx_in3
	txIn[3] = tx_in4

	txOut := make([]*common.TxOut, 2)
	tx_out1 := &common.TxOut{
		Address: "mp1jRsp67ys7VXLFsgF2yr5MRs9pytidms",
		Value:   25000000,
	}

	tx_out2 := &common.TxOut{
		Address: "mke87z1EYAwWpvNS7YdTqFxYvyaNHKUvGK",
		Value:   35000000,
	}
	txOut[0] = tx_out1
	txOut[1] = tx_out2
	//feeRate, detail:https://bitcoinfees.earn.com/api
	var feeRate uint64 = 77
	in, out, fee := BitcoinUtxo.CoinSelect(txIn, txOut, feeRate)

	for i := range in {
		fmt.Println(in[i].Hash)
		fmt.Println(in[i].Index)
		fmt.Println(in[i].Script)
		fmt.Println(in[i].Value)
		fmt.Println("--------------------------")
	}
	fmt.Println("*************************************")
	for i := range out {
		fmt.Println(out[i].Address)
		fmt.Println(out[i].Value)
		fmt.Println("--------------------------")
	}
	fmt.Println("*************************************")
	fmt.Println(fee)
}
