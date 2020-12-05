package fileinput

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func OpenInputFile() (*os.File, error) {
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		return nil, errors.New("input file name is missing")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed opening file: %w", err)
	}

	return file, nil
}
