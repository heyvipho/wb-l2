package main

import (
	"fmt"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

func main() {
	ctx := new(Context)

	data1 := "example"
	ctx.Algorithm(&FirstPrintAlg{})
	ctx.Print(data1)

	data2 := "example"
	ctx.Algorithm(&SecondPrintAlg{})
	ctx.Print(data2)
}

type StrategyPrint interface {
	Print(interface{})
}

type FirstPrintAlg struct {
}

func (s *FirstPrintAlg) Print(a interface{}) {
	fmt.Printf("[FirstPrintAlg] %v\n", a)
}

type SecondPrintAlg struct {
}

func (s *SecondPrintAlg) Print(a interface{}) {
	fmt.Printf("[SecondPrintAlg] %v\n", a)
}

type Context struct {
	strategy StrategyPrint
}

func (c *Context) Algorithm(a StrategyPrint) {
	c.strategy = a
}

func (c *Context) Print(a interface{}) {
	c.strategy.Print(a)
}
