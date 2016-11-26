// Page 8
// Echo speed test.
// Tests each of echo1, echo2, and echo3 in the ch1 folder against a really long string.
package main

import (
	"fmt"
	"github.com/ncruz8991/gopl.io/ch1/echo1/echo1"
	"github.com/ncruz8991/gopl.io/ch1/echo2/echo2"
	"github.com/ncruz8991/gopl.io/ch1/echo3/echo3"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Runs the function echoTest and returns the time.
func benchmark(s string, echoTest func()) float64 {
	fmt.Printf("Running %s test...", s)
	start := time.Now()
	echoTest()
	totalTime := time.Since(start).Seconds()
	fmt.Println("Done.")
	return totalTime
}

// Checks if an error exists. If it does, panic.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// The book mentions that there are more thorough and systematic speed tests to write.
// Revisit this in Section 11.4.
// TODO Rewrite this as a real benchmark speed test (Section 11.4).
func main() {
	// To make it easier to test, we are creating our own os.Args.
	argsData, err := ioutil.ReadFile("ch1/exercises/e1.3/args.txt")
	check(err)

	// Make a gigantic array.
	argsString := strings.Fields(string(argsData))
	for i := 0; i < 40; i++ {
		os.Args = append(os.Args, argsString...)
	}

	// Start the tests.
	echo1Time := benchmark("echo1", echo1.Run)
	echo2Time := benchmark("echo2", echo2.Run)
	echo3Time := benchmark("echo3", echo3.Run)

	// Create an output file and check for an error.
	f, err := os.Create("ch1/exercises/e1.3/output.txt")
	check(err)
	defer f.Close()

	// Write the elapsed times to the file.
	f.WriteString("Finished Benchmarking:")
	f.WriteString(fmt.Sprintf("\n\t%s\t %.2fs", "echo1", echo1Time))
	f.WriteString(fmt.Sprintf("\n\t%s\t %.2fs", "echo2", echo2Time))
	f.WriteString(fmt.Sprintf("\n\t%s\t %.2fs", "echo3", echo3Time))

	// Flush writes to stable storage.
	f.Sync()
}
