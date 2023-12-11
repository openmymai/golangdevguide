package main

import "fmt"

type bot interface {
	getGreeting() string
}
// Rules of interfaces
// type bot interface {
// 	getGreeting(string, int) (string, error)
//	function		arg types			return types				
// }

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// no need many printGreeting functions
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}