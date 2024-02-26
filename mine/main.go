package main

import (
	"fmt"
	"os"
)

func main() {
	getArgs := os.Args[1:] // get Arguments entered in Command Line
	if len(getArgs) != 9 { // Check the lenght of all arg (string) passed in command line
		fmt.Printf("error")
		return
	}
	Table2D := make([][]byte, 9) // Create a 2D Table using Make
	for i := range Table2D {     // make an
		if len(getArgs[i]) != 9 {
			fmt.Print("Error")
			return
		}
		fmt.Print("-->", Table2D[i], " <--\n")

		for _, checkChar := range getArgs[i] {
			if checkChar != '.' && (checkChar < '1' || checkChar > '9') {
				fmt.Print("Error")
				return
			}
		}

		Table2D[i] = []byte(getArgs[i])

	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", Table2D[i][j])
		}
		fmt.Println()
	}
}

/*




*/
