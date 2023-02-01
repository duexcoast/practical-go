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
		&i1,
		&p1,
		&i2,
	}
	moveAll(ms, 0, 0)

	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k: ", k)

	p1.FoundKey(1)
	fmt.Println(p1.Keys)
}

const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal (not exported)
)

type Key byte

// implement fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

/* Exercise
- Add a "Keys" field to a Player which is a slice of Key
- Add a "FoundKey(k Key)" method to player which will add k to Key if it's not there
	- Err if k is not one of the known keys
*/

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(int, int)
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
	Keys []Key
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}

	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
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

func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
