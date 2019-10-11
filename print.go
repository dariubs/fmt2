package fmt2

import (
	"fmt"
	"os"
)

// Printlnf formats according to a format specifier and writes to standard output and end string with a \n.
// It returns the number of bytes written and any write error encountered.
func Printlnf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format+"\n", a...)
}
