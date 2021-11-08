package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

func main() {
	var cpuCount = runtime.NumCPU()
	fmt.Printf("Running with %d cores\n", cpuCount) // not sure why `Println` does not take %d-style specifiers
	fmt.Println("Calling the function count")
	fmt.Printf("The result of calling count is [%d]\n", count(os.Stdin))
	fmt.Println("Done!")
}

func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	wc := 0
	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}
	return wc
}
