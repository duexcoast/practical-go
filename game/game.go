package main

import "fmt"

func main() {
	var i1 Item
	fmt.Printf("%#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("%#v\n", i2)

	i3 := Item{
		Y: 10,
	}
	fmt.Printf("%#v\n", i3)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))
}

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	i := Item{
		X: x,
		Y: y,
	}
	return &i, nil
}

const (
	maxX = 1000
	maxY = 600
)

// Item is an item in the game
type Item struct {
	X int
	Y int
}
