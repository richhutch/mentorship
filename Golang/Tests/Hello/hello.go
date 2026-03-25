package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case "Spanish":
		return "Hola, " + name
	case "French":
		return "Bonjour, " + name
	default:
		return englishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
