package util

import (
	"fmt"
	"os"
)

func DD(v interface{}) {
	fmt.Printf("%+v\n", v)
	os.Exit(1)
}
