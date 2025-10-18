package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

type Animal interface {
	Speak()
	Move()
	Eat()
}

// All the animals and methods
type Cow struct {}
type Bird struct {}
type Snake struct {}


func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }


func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }

        words := strings.Fields(strings.ToLower(strings.TrimSpace(input)))
        if len(words) != 3 {
			fmt.Print("Bad format, should be:\nnewanimal <name> <cow/bird/snake>\nor\nquery <name> <eat/move/speak>")
            continue
        }

		command, name, arg := words[0], words[1], words[2]

		switch command {
		case "newanimal":
			var animal Animal
			switch arg {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Printf("Unknown animal type '%s'. Use cow, bird, or snake.\n", arg)
				continue
			}
			animals[name] = animal
			fmt.Println("Created it!")

		case "query":
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch arg {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown query. Use eat, move, or speak.")
			}

		default:
			fmt.Println("Unknown command. Use newanimal or query.")
		}
	}
}
