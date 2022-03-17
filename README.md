# golang-sandbox

Bits and snippets to get started with golang

### Variable declaration

```
var varName type = value
```

or

```
varName := value
```

- variables cannot be used before declaration
- compliler shows a warning if the declared variable is never used
- while using ":=" syntax, the types are inferred. Here the card inferrs strigng type
- there are 4 basic variable types
  - string
  - int
  - boolean
  - float64

### Function declaration

```
func funcName(param paramType) returnType {...}

func getCard() string {
    return "Ace of Hearts"
}
```

### Arrays and Slices

- Arrays have <b>restricted length=</b> they cannot grow. You <b>cannot</b> add or delete elements
- Slices are </b>flexible.</b>It provides you methods to add and delete elements
- Both Array and Slices allow a single element type

```
<!-- Slice -->
varName := []type{val, val}
```

### Loops

```
for index,value := range slice {...}

for i, card := range cards {
    fmt.Println(i, card)
}
```

### Custom Types

```
type newType extendsType

type deck []string
```

### Functions with Receivers

Methods can be added to these custom types by passing the receiver to the function

```
func (ins, customType) funcName(param paramType ) {...}

func (d, deck) print{
    for i, card := range d {
        print(i, card)
    }
}


cards := deck {"Ace of Diamonds"}
cards = cards.append(cards, "Six of Hearts")
cards.print()
```

Here deck is the receiver, which gets all the variables with type deck to perform the function operation.

d is the instance of the varible on which the function will be executed.
