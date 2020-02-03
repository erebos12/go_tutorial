package main

import "fmt"

/*
When you are type with exactly this function 'getGreeting() string'
then you are a honorary member of type bot. So you can now call the
function 'func printGreeting(b bot)'.

In Java you say "xyz implements interface" to describe xyz is an implementor of an interface.
In Go you just implement exactly the same function for a specific receiver.
Theres is no need to link together 'bot' and 'englishBot/spanishBot' because interfaces are implicit.
Implicitly, they are connected with each other.
*/
type bot interface {
	// just specify the types
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
