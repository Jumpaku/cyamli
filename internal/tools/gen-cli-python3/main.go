package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Jumpaku/cyamli/python3"
	"github.com/Jumpaku/cyamli/schema"
)

func main() {
	s, err := schema.Load(os.Stdin)
	if err != nil {
		log.Panicln(err)
	}

	buf := bytes.NewBuffer(nil)
	if err := python3.Generate(s, buf); err != nil {
		log.Panicln(err)
	}
	if err := format(buf, os.Stdout); err != nil {
		log.Panicln(err)

	}
}

func format(in io.Reader, out io.Writer) error {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanLines)
	lines := 0
	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			lines++
			continue
		}

		switch lines {
		case 0:
			line = "\n" + line
		case 1:
			line = "\n" + line
		case 2:
			line = "\n" + line
		default:
			line = "\n\n" + line
		}
		lines = 0

		_, err := out.Write([]byte(line))
		if err != nil {
			return fmt.Errorf("fail to format output: %w", err)
		}
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("fail to format output: %w", err)
	}
	return nil
}
