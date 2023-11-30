package main

import "fmt"

const portugueseHelloPrefix = "Olá, "

func Hello(name string) string {
	return portugueseHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
