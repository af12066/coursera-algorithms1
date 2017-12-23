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

func (n QuickFindUF) Union(p int, q int) error {
	pid := id[p]
	qid := id[q]
	for i := 0; i < len(id); i++ {
		if id[i] == pid {
			id[i] = qid
		}
	}
	return nil
}

func (n QuickFindUF) Connected(p int, q int) (bool, error) {
	return id[p] == id[q], nil
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

		isConnected, err := uf.Connected(p, q)
		if !isConnected {
			uf.Union(p, q)
			fmt.Printf("%d %d\n", p, q)
			fmt.Printf("Array: %v\n", id)
		}
	}
}
