package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-version"
)

func main() {
	var in string
	if _, err := fmt.Scan(&in); err != nil {
		log.Panicln(err)
	}
	v, err := version.NewVersion(in)
	if err != nil {
		log.Panicf("invalid version: %q: %+v\n", in, err)
	}
	seg := v.Core().Segments()
	fmt.Printf("v%d.%d.%d", seg[0], seg[1], seg[2]+1)
}
