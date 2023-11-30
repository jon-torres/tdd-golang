package main

import "fmt"

const portugueseHelloPrefix = "Olá, "

func Hello(name string) string {
	if name == "" {
		name = "Mundo"
	}
	return portugueseHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
