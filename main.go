package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const dutch = "Dutch"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const dutchHelloPrefix = "Hallo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
    case dutch:
		prefix = dutchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Dane", "Spanish"))
}
