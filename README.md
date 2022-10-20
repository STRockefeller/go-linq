# Go-Linq

![license](https://img.shields.io/github/license/STRockefeller/go-linq?style=plastic)![code size](https://img.shields.io/github/languages/code-size/STRockefeller/go-linq?style=plastic)![open issues](https://img.shields.io/github/issues/STRockefeller/go-linq?style=plastic)![closed issues](https://img.shields.io/github/issues-closed/STRockefeller/go-linq?style=plastic)![go version](https://img.shields.io/github/go-mod/go-version/STRockefeller/go-linq?style=plastic)![latest version](https://img.shields.io/github/v/tag/STRockefeller/go-linq?style=plastic)

[![GitHub Super-Linter](https://github.com/STRockefeller/go-linq/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/marketplace/actions/super-linter)[![Go Report Card](https://goreportcard.com/badge/github.com/STRockefeller/go-linq)](https://goreportcard.com/report/github.com/STRockefeller/go-linq)

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

## Benchmark

compare with another linq package.

```go
package main

import (
 "fmt"
 "testing"

 STRLinq "github.com/STRockefeller/go-linq"
 AHMLinq "github.com/ahmetb/go-linq/v3"
)

var testStrings = []string{"Oh,", "mister", "ocean", "fish", "!"}
var testIntegers = []int{7, 5, 3, 9, 5, 1, 7, 4, 1, 0, 3, 6, 9}

func BenchmarkSTRockefeller_linq(b *testing.B) {
 for i := 0; i < b.N; i++ {
  STRLinq.Linq[string](testStrings).Where(func(s string) bool { return len(s) >= 3 }).Skip(1).Contains("mister")

  mySlice := STRLinq.Linq[int](testIntegers).Distinct().Where(func(i int) bool { return i > 3 }).ToSlice()

  if false {
   fmt.Print(mySlice)
  }
 }
}

func BenchmarkAhmetb_linq(b *testing.B) {
 for i := 0; i < b.N; i++ {
  AHMLinq.From(testStrings).Where(func(i interface{}) bool { return len(i.(string)) >= 3 }).Skip(1).Contains("mister")

  var mySlice []int
  AHMLinq.From(testIntegers).Distinct().Where(func(i interface{}) bool { return i.(int) > 3 }).ToSlice(&mySlice)

  if false {
   fmt.Print(mySlice)
  }
 }
}
```

result

```bash
goos: windows
goarch: amd64
pkg: test
cpu: Intel(R) Core(TM) i7-9700 CPU @ 3.00GHz
BenchmarkSTRockefeller_linq-8     2529393        467.3 ns/op      400 B/op       12 allocs/op
BenchmarkAhmetb_linq-8             521704       2121 ns/op     1080 B/op       45 allocs/op
PASS
coverage: [no statements]
ok   test 3.004s
```
