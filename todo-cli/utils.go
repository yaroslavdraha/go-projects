package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toBool(i int) bool {
	return i != 0
}

func readInput(text string) string {
	fmt.Println(text)

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return name
}
