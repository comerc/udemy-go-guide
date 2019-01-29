package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
	dummy := map[interface{}]interface{}{
		"one":   1,
		2:       "bad",
		"three": true,
	}
	fmt.Println(dummy)
	// var colors2 = map[string]string
	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"
	fmt.Println(colors2)
	colors3 := make(map[int]string)
	colors3[10] = "#ffffff"
	delete(colors3, 10)
	fmt.Println(colors3)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
