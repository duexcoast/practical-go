package main

import "fmt"

func main() {
	i3 := Item{
		Y: 10,
	}
	i1 := Item{
		X: 45,
		Y: 10,
	}
	i2 := Item{
		X: 200,
		Y: 50,
	}
	fmt.Printf("i3: %#v\n", i3)

	(&i3).Move(5, 50)
	fmt.Printf("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "Cornelius",
		Item: Item{500, 500},
	}
	fmt.Printf("p1: %#v\n", p1)

	p1.Move(20, 20)
	fmt.Printf("p1 (move): %#v\n", p1)

	fmt.Printf("p1.Item.X;%#v\n", p1.Item.X)

	ms := []mover{
		i1,
		p1,
		i2,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
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

type Player struct {
	Name string
	Item // embed Item
	T
}

type T struct {
	X int
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

// i is called "the receiver"
// if you want to mutate, use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}
