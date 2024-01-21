

## [Tutorial-URL](https://blog.logrocket.com/build-blockchain-with-go/)

- `Mining` is the process of adding a new block to the blockchain

- This involves generating a block hash that starts with a desired number of `zeros (the number of zeros is called the mining difficulty)`

- This means if the mining difficulty is three, you have to generate a block hash that starts with `"000"` like, `"0009a1bfb506…"`

- Because we are deriving a block’s hash from its `data` and `PoW`, we need to keep changing the `PoW` value of the current block until we get a hash that satisfies our mining condition `(starting zeros > difficulty)`

- Here, we created an addBlock method to the Blockchain type that does the following:

```sh
Collects the details of a transaction (sender, receiver, and transfer amount)
Creates a new block with the transaction details
Mines the new block with the previous block hash, current block data, and generated PoW
Adds the newly created block to the blockchain
```

- The * (asterisk) in Go is used to denote a pointer receiver. In Go, methods with pointer receivers can modify the value they're called on, while methods without pointer receivers work on a copy of the value.

- Let's break it down:

```golang

func (b Block) calculateHash() string {
	data, _ 		:= json.Marshal(b.data)
	blockData 		:= b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash 		:= sha256.Sum256([]byte(blockData))

	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}
```

- func (b Block) calculateHash() string: This method calculateHash() has a receiver of type Block. It operates on a copy of the Block value. If you call this method on a Block instance, it won't modify the original Block value but will perform calculations and return a hash based on the copied value.

- func (b *Block) mine(difficulty int): This method mine() has a receiver of type *Block, which is a pointer to a Block. When you call this method on a *Block, it operates directly on the original Block value that the pointer refers to. It can modify the fields of the original Block instance, in this case, updating the pow and hash fields until a condition is met.

- The impact of using a pointer receiver (*Block) in the mine() method is that it allows the method to modify the actual fields of the Block instance it's called on. Without the pointer receiver, the method would operate on a copy of the Block, and any modifications wouldn't affect the original Block outside of the method scope.

- The calculateHash() method, on the other hand, is designed to perform calculations on the data without altering the original Block value. Therefore, it uses a value receiver (Block) to work on a copy of the Block without affecting the original instance.
```