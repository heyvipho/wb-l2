package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

func main() {
	city := new(City)

	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})

	result := city.Accept(&People{})

	fmt.Println(result)
}

type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

type Place interface {
	Accept(v Visitor) string
}

type People struct {
}

func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

type City struct {
	places []Place
}

func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

type SushiBar struct {
}

func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

type Pizzeria struct {
}

func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

type BurgerBar struct {
}

func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}
