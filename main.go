package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mlbright/humanbytes/humanize"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		out := humanize.HumanizeNumber(scanner.Text())
		fmt.Println(string(out))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner error:", err)
		os.Exit(1)
	}
}
