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
	fmt.Printf("Management System\tver0.1\n")
	// Take file name as user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter file name:")
	name, _ := reader.ReadString('\n')
	f, err := os.Open(name)
	//Handle errors
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
	//loop and print
	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

}
