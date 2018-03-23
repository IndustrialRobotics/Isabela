package main

import "fmt"

func main() {

	colors := map[string]string{
		"ff0042": "Orange",
		"45dfr9": "White",
		"7de541": "Black",
		"25das9": "Green",
		"43fgrt": "turquoise",
	}

	printMap(colors)

}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Hex code for", k, "is color", v)
	}
}
