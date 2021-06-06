# blockchain-models

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/philohsophy/blockchain-models)

Part of [Blockchain-Demo project](https://github.com/philohsophy/blockchain-demo)

## Outline

Shared models for Blockchain-Demo project

## Models

### Address

```Golang
// models.Address Example
address := models.Address{
    Name: "Foo",
    Street: "FooStreet",
    HouseNumber: "1",
    Town: "FooTown",
}
```

### Transaction

```Golang
// models.Transaction Example
transaction := models.Transaction{
    Id: "823ba770-7569-4b0c-8959-33d40446a8af",
    RecipientAddress: address,
    SenderAddress: models.Address{ Name: "Bar", Street: "BarStreet", HouseNumber: "1", Town: "BarTown" },
    Value: 100.21,
}
```

### Block

```Golang
// models.Block Example
block := models.Block{
    PreviousBlockHash: [125 76 136 128 69 110 124 47 206 6 39 244 101 127 212 90 192 139 89 163 95 106 41 127 159 133 86 107 82 33 243 67],
    MerkleRootHash: [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], // calculated automatically
    Timestamp: 1622995001073060100,
    NBits: 1,
    Nonce: "bdc59445d268cdabeb6fdf54607b0cf7",
    Transactions: [
        transaction,
    ]
}

// Calculate hash of block
hash := block.GetHash()
// --> i.e. [192 7 217 229 91 139 203 249 86 116 56 83 8 159 148 47 9 110 9 23 246 23 196 78 147 197 195 233 119 59 119 125]
```

All models include:

- JSON-representation
- IsValid() receiver function for checking validity of model (check for zero values)
