# gollection
A simple library to manage collections of things in go

## Installation

```bash
go get github.com/rodriez/gollection
```

## Usage

```go
import "github.com/rodriez/gollection"

collection := gollection.NewStringCollection("Barber", "Alba", "Connan")

total := collection.Map(func(e gollection.Element) gollection.Element {
    return gollection.NewString(fmt.Sprintf("%s;", e.String()))
}).Reduce(func(acc, e gollection.Element) gollection.Element {
    return gollection.NewString(fmt.Sprintf("%s%s", acc.String(), e.String()))
}, gollection.NewString(""))

fmt.Println(total)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

[Hey 👋 buy me a beer! ](https://www.buymeacoffee.com/rodriez)