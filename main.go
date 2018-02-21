package main

import (
	"fmt"
	"os"

	"golang.org/x/image/tiff"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s file.tiff\n", os.Args[0])
		os.Exit(1)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	cfg, err := tiff.DecodeConfig(f)
	if err != nil {
		return err
	}
	fmt.Println(cfg.Width, cfg.Height)
	return nil
}
