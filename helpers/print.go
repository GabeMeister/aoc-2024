package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func Pause(vals ...interface{}) {
	// return

	fmt.Println()
	if len(vals) > 0 {
		fmt.Println(vals...)
	}
	fmt.Println()

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
