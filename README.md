# lib

ascii generation library for the iskre cli tool

## Features

### Generic

- [x] Boxes

### Diagrams

- [ ] pie chart
- [ ] ven diagram
- [ ] control flow diagrams
  - [x] sequence diagram
  - [ ] condition diagram
  - [x] repetition diagram
- [ ] flow chart
- [x] tables
- [ ] mind map
- [ ] state diagram
- [ ] gannt diagram

## Showcase

### Generic

```text
┌────────────┐
│  title     │
│  subtitle  │
└────────────┘
```

### Control flow

```text
┌───────────────────────────────┐
│  While loop conditiontest     │
│  ┌────────────────────────────┤
│  │  test                      │
│  │   test                     │
│  │    test                    │
│  │  test                      │
└──┴────────────────────────────┘

┌─────────────────────────┐
│  First                  │
├─────────────────────────┤
│  Second                 │
├─────────────────────────┤
│  Third                  │
├─────────────────────────┤
│  massive long Sequence  │
└─────────────────────────┘
```

### Table

```text
┌─────────────┬──────────────┬─────────────┐
│    First    │    Second    │    Third    │
├─────────────┼──────────────┼─────────────┤
│ Test1       │ Test2        │ Test3       │
│ Test1       │ Test2        │ Test3       │
└─────────────┴──────────────┴─────────────┘
```

## Usage

> for up to date examples take a look into [main.go](./main.go)

```go
package main

import (
    // import the library
    "github.com/iskre/lib/core"
)

func main(){
    // all methods in this library are defined on the Iskre struct
    i := core.Iskre{}

    // all methods return strings
    diagram := i.Repetition("While loop conditiontest", "test\n test\n  test\ntest")

    fmt.Println(diagram)
}
```

## Setup

```bash
git clone https://github.com/iskre/lib.git
cd lib
go run .
```

## Tests

```bash
go test ./...
```

Made with ❤️ by [intevel](https://github.com/intevel) and [xnacly](https://github.com/xnacly)
