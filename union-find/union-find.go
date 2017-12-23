package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QuickFindUF struct {
	n int
}

var id []int

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
	return stringInput
}

func IntStdin() (int, error) {
	stringInput := StrStdin()
	return strconv.Atoi(strings.TrimSpace(stringInput))
}

func root(i int) int {
	for {
		if i == id[i] {
			break
		}
		i = id[i]
	}
	return i
}

func (n QuickFindUF) Union(p int, q int) error {
	i := root(p)
	j := root(q)
	id[i] = j
	return nil
}

func (n QuickFindUF) Connected(p int, q int) bool {
	return root(p) == root(q)
}

func main() {
	n, err := IntStdin()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uf := QuickFindUF{n}
	for i := 0; i < uf.n; i++ {
		id = append(id, i)
	}

	fmt.Printf("Array: %v\n", id)

	for {
		stdin := StrStdin()
		if stdin == "" {
			break
		}
		p, err := strconv.Atoi(strings.TrimSpace(stdin))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		stdin = StrStdin()
		if stdin == "" {
			break
		}
		q, err := strconv.Atoi(strings.TrimSpace(stdin))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if !uf.Connected(p, q) {
			uf.Union(p, q)
			fmt.Printf("%d %d\n", p, q)
			fmt.Printf("Array: %v\n", id)
		}
	}
}
