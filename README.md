# Go-Linq

![license](https://img.shields.io/github/license/STRockefeller/go-linq?style=plastic)![code size](https://img.shields.io/github/languages/code-size/STRockefeller/go-linq?style=plastic)![open issues](https://img.shields.io/github/issues/STRockefeller/go-linq?style=plastic)![closed issues](https://img.shields.io/github/issues-closed/STRockefeller/go-linq?style=plastic)![go version](https://img.shields.io/github/go-mod/go-version/STRockefeller/go-linq?style=plastic)![latest version](https://img.shields.io/github/v/tag/STRockefeller/go-linq?style=plastic)

C# `System.Linq` [Enumerable](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable?view=net-6.0) methods and `System.Collections.Generic` [List](https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1?view=net-6.0) realization by go generic.

## Notice

golang version must be greater than v1.18

for previous go versions (<1.18), you can try [this one](https://github.com/STRockefeller/linqable).

## Usage

install package

```shell
go get github.com/STRockefeller/go-linq
```

import package

```go
import "github.com/STRockefeller/go-linq"
```

enjoy

```go
func linqTest() {
 type user struct {
  name string
  age  int
 }
 users := linq.Linq[user]{}
 users.Add(user{
  name: "Rockefeller",
  age:  27,
 })
 users.Prepend(user{
  name: "newUser",
  age:  18,
 })
 adultsCount := users.Where(func(u user) bool { return u.age >= 18 }).Count(func(u user) bool {return true})
 fmt.Println("there are ",adultsCount,"adults in users.")
}
```

See [pkg.go.dev document](https://pkg.go.dev/github.com/STRockefeller/go-linq) for details
