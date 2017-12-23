package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UF struct {
	n int
}

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

func (n UF) Union(p int, q int) error {
	return nil
}

func (n UF) Connected(p int, q int) (bool, error) {
	return false, nil
}

func main() {
	n, err := IntStdin()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uf := UF{n}

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
		}
	}
}
