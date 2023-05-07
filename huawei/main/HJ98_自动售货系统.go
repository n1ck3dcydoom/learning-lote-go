package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	init := strings.Split(scanner.Text(), " ")

}

type Good struct {
	Price int
	Count int
}

type Coin struct {
	Price int
	Count int
}
