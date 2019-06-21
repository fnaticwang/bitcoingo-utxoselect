# bitcoin_utxo_select
bitcoin utxos selection algorithm of golang

Since bitcoin adopts UTXO(Unspent Transaction Output) model, its BTC is scattered in different UTXO even if it is the same address.
So when a user spends a bitcoin, his wallet or client needs to find the right UTXO Set and piece together the right transaction.
This algorithm for finding UTXO is called Coin select algorithm.
This program is implemented the Coin select strategy adopted by Bitcoin core.For more discussion, please refer to [An Evaluation of Coin Selection Strategies](http://murch.one/wp-content/uploads/2016/11/erhardt2016coinselection.pdf).
In this 2016 paper, Mark Erhardt USES more than 60 pages to explore different strategies for choosing UTXO based on different needs.

## License [MIT](LICENSE)
