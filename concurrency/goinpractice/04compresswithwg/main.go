package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	// A WaitGroup doesn't need to be initialized.
	var wg sync.WaitGroup

	// reference i outside the for loop
	var i int = -1
	var file string
	// no need : in the for loop
	for i, file = range os.Args[1:] {
		wg.Add(1)
		// This func calls compress and
		// then notifies the wait group
		// that it's done.
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()

	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
