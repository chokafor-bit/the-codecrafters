package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Base Converter (Type 'quit' to exit)")
	for {
		fmt.Print("> ")
		var in, base string
		if _, err := fmt.Scan(&in); err != nil {
			continue
		}
		if in == "quit" {
			return
		}
		if _, err := fmt.Scan(&base); err != nil {
			continue
		}
		b := 10
		if base == "hex" {
			b = 16
		} else if base == "bin" {
			b = 2
		}
		val, err := strconv.ParseInt(in, b, 64)
		if err != nil || (base != "hex" && base != "bin" && base != "dec") {
			fmt.Printf("Error: '%s' is not valid %s\n", in, base)
			continue
		}
		if base == "dec" {
			fmt.Printf(" ✦ Binary: %b\n ✦ Hex:    %X\n", val, val)
		} else {
			fmt.Printf(" ✦ Decimal: %d\n", val)
		}
	}
}
