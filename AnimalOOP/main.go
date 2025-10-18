package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

type Animal struct {
    FoodEaten string
    Locomotion string
    Sound string
}

func (a Animal) Eat() string {
    return a.FoodEaten
}

func (a Animal) Move() string {
    return a.Locomotion
}

func (a Animal) Speak() string {
    return a.Sound
}

func main() {
    animals := map[string]Animal{
        "cow":   {"grass", "walk", "moo"},
        "bird":  {"worms", "fly", "peep"},
        "snake": {"mice", "slither", "hsss"},
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }

        words := strings.Fields(strings.ToLower(strings.TrimSpace(input)))
        if len(words) != 2 {
			fmt.Println("Bad format, should be: <cow/bird/snake> <eat/move/speak>")
            continue
        }

        animalName := words[0]
        action := words[1]

        animal, exists := animals[animalName]
        if !exists {
            fmt.Println("Invalid animal:", animalName)
            continue
        }

        switch action {
        case "eat":
            fmt.Println(animal.Eat())
        case "move":
            fmt.Println(animal.Move())
        case "speak":
            fmt.Println(animal.Speak())
        default:
            fmt.Println("Invalid action:", action)
        }
    }
}
