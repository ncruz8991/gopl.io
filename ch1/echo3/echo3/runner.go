package echo3

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
