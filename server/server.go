package server

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

func (e englishBot) getGreeting() string {
	// fmt.Println("hola")
	return "Helo"
}

func printGreetingEnglish(e englishBot) {
	fmt.Println(e.getGreeting())
}

func M() string {
	return "welcome to Go"
}
