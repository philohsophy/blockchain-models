# dummy-blockchain-models

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/philohsophy/dummy-blockchain-models)

Part of [Dummy-Blockchain project](https://github.com/users/philohsophy/projects/1)

## Outline

Shared models for Dummy-Blockchain project

## Models

```Golang
// models.Transaction Example
{
    Id: "823ba770-7569-4b0c-8959-33d40446a8af"
    RecipientAddress: { Name: "Foo", Street: "FooStreet", HouseNumber: "1", Town: "FooTown" }
    SenderAddress: { Name: "Bar", Street: "BarStreet", HouseNumber: "1", Town: "BarTown" }
    Value: 100.21
}

// models.Address Example
{
    Name: "Foo",
    Street: "FooStreet",
    HouseNumber: "1",
    Town: "FooTown"
}
```

All models include:

- JSON-representation
- IsValid() receiver function for checking validity of model (check for zero values)
