package main

import (
	"fmt"
	"strings"
)

type animal interface {
	Eat() string
	Move() string
	Speak() string
}

type cow struct {
	food, locomotion, noise string
}

func (act cow) Eat() string {
	return act.food
}

func (act cow) Move() string {
	return act.locomotion
}

func (act cow) Speak() string {
	return act.noise
}

type bird struct {
	food, locomotion, noise string
}

func (act bird) Eat() string {
	return act.food
}

func (act bird) Move() string {
	return act.locomotion
}

func (act bird) Speak() string {
	return act.noise
}

type snake struct {
	food, locomotion, noise string
}

func (act snake) Eat() string {
	return act.food
}

func (act snake) Move() string {
	return act.locomotion
}

func (act snake) Speak() string {
	return act.noise
}

func main() {
	var idMap map[string]animal
	idMap = make(map[string]animal)

	var com, obj, meth string

	cow_data := cow{"grass", "walk", "moo"}
	bird_data := bird{"worms", "fly", "peep"}
	snake_data := snake{"mice", "slither", "hsss"}

	fmt.Println("Enter \"Q\" to quit")
	for {
		fmt.Print(">")

		fmt.Scan(&com)
		com = strings.ToLower(com)
		if com == "q" {
			break
		} else if com == "newanimal" {
			fmt.Scan(&obj)
			fmt.Scan(&meth)
			switch strings.ToLower(meth) {
			case "cow":
				idMap[obj] = &cow_data
			case "bird":
				idMap[obj] = &bird_data
			case "snake":
				idMap[obj] = &snake_data
			}
			fmt.Println("Created it!")
		} else if com == "query" {
			fmt.Scan(&obj)
			fmt.Scan(&meth)
			data, flag := idMap[obj]
			if flag {
				switch strings.ToLower(meth) {
				case "eat":
					fmt.Println(data.Eat())
				case "move":
					fmt.Println(data.Move())
				case "speak":
					fmt.Println(data.Speak())
				}
			} else {
				fmt.Println("There is no animal with such name, try again.")
			}
		}
	}
}
