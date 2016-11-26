// See page 8.
package main

import (
	"fmt"
	"os"
	"strings"
)

/*
This is a modified version of echo3 that prints out three things:
 	1. The full program path
	2. The program name that launched this go program
	3. Any command line arguments

Example output:
$ go run exercises/e1.1/main.go hey what\'s up
C:\Users\Nick\AppData\Local\Temp\go-build537770626\command-line-arguments\_obj\exe\main.exe

Program Name:
main.exe

Command line arguments:
hey what's up
*/
func main() {
	fullProgramPath := os.Args[0]
	fmt.Println("Program Path:")
	fmt.Println(fullProgramPath)

	fullProgramPathSplit := strings.Split(fullProgramPath, "\\")
	fmt.Println()
	fmt.Println("Program Name:")
	fmt.Println(fullProgramPathSplit[len(fullProgramPathSplit)-1])

	if len(os.Args) > 1 {
		fmt.Println()
		fmt.Println("Command line arguments:")
		fmt.Println(strings.Join(os.Args[1:], " "))
	}
}
