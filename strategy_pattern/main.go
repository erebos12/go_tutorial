package main

import "fmt"

type GreeterStrategy interface {
	Greet(name string) string
}

type FriendlyStrategy struct{}

func (s *FriendlyStrategy) Greet(name string) string {
	return "Hey, " + name + "! Schön, dich zu sehen!"
}

type FormalStrategy struct{}

func (s *FormalStrategy) Greet(name string) string {
	return "Guten Tag, " + name + ". Es ist mir eine Ehre."
}

// Greeter context
type Greeter struct {
	Strategy GreeterStrategy
	Name     string
}

// binds context struct to interface method
func (g *Greeter) Greet() {
	fmt.Println(g.Strategy.Greet(g.Name))
}

func (g *Greeter) SetStrategy(strategy GreeterStrategy) {
	g.Strategy = strategy
}

func main() {
	greeter := Greeter{Name: "Alex"}

	friendly := &FriendlyStrategy{}
	greeter.SetStrategy(friendly)
	greeter.Greet()

	formal := &FormalStrategy{}
	greeter.SetStrategy(formal)
	greeter.Greet()
}
