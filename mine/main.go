package main

import (
	"fmt"
	"os"
)

func main() {
	getArgs := os.Args[1:]
	if len(getArgs) != 9 {
		fmt.Printf("error")
		return
	}
	Table2D := make([][]int, 9)
	for get := range Table2D {
		if len(getArgs[get]) != 9 {
			fmt.Print("Error")
			return
		}
	}
}

/*


	getArgs := os.Args[1:]

	deuxDArray := make([][]byte, 9)
	for i := range deuxDArray {
		if len(getArgs) != 9 || len(getArgs[i]) != 9 {
			fmt.Print("Error")
			return
		}
	}
	fmt.Println(deuxDArray)
	for i := 0; i < len(getArgs); i++ {
		fmt.Print(getArgs[i])
		fmt.Print("\n")
	}
	deuxDArray: This is a variable declaration. It declares a new variable named deuxDArray.

	:=: This is the short variable declaration syntax in Go. It declares and initializes the deuxDArray variable in a single line.
	 The := operator is used to infer the type of the variable based on the value being assigned to it.

	make([][]byte, N): This is a built-in function in Go used to create slices, maps, and channels. In this case,
	 it's used to create a slice of byte slices (a 2D slice of bytes). The first argument specifies the type of the slice ([][]byte),
	  and the second argument specifies the length of the slice (N).

	[]byte: This specifies the type of elements in the slice. In this case, it's a byte slice, indicating that each element of deuxDArray
	will be a slice of bytes.

	N: This is a constant representing the size of the Sudoku deuxDArray. It specifies the length of the slice deuxDArray, indicating that deuxDArray
	will have N elements, each of which is a byte slice.

	Putting it all together, deuxDArray := make([][]byte, N) creates a 2D slice of bytes (deuxDArray) with N rows, where each row is a byte slice.
	 This is commonly used in Go to represent a 2D array or matrix.
*/
