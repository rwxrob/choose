package choose

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var DefaultPrompt = `#? `

func readline() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

type Chooser[T any] interface {
	Choose() (T, error)
}

type Choices[T any] []T

// Choose prompts terminal user to pick from one of the choices.
func (c Choices[T]) Choose() (T, error) {
	var empty T
	for i, v := range c {
		fmt.Printf("%v. %v\n", i+1, v)
	}
	for {
		fmt.Print(DefaultPrompt)
		resp := readline()
		if resp == "q" {
			return empty, nil
		}
		n, _ := strconv.Atoi(resp)
		if 0 < n && n < len(c)+1 {
			return c[n-1], nil
		}
	}
}

// From prompts terminal user to pick from one of the choices using the
// DefaultChooser.
func From[T any](choices []T) (T, error) {
	return Choices[T](choices).Choose()
}
