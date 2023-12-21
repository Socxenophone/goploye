package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

var employee struct {
	name     string
	age      int
	deployed bool
}

func main() {
	fmt.Printf("Management System\tver0.1")
	f, err := os.Open("foo.txt")
	if err != nil {
		switch {
		case errors.Is(err, fs.ErrNotExist):
			{
				fmt.Printf("\nFile does not exist!\n")
			}
		default:
			{
				log.Fatal(err)
			}
		}

	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

}
