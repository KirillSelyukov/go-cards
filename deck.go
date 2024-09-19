package main

import "fmt"

type deck []string
type card string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)

	}
}

func (c card) print() {
	fmt.Println(c)

}
