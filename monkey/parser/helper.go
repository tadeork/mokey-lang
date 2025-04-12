package parser

import (
	"fmt"
	"os"
)

func debug(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "[DEBUG] "+format+"\n", args...)
}
