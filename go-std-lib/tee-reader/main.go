package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz
/*
r: download the file
w: progress counter: measure in real time the number of mb downloaded
w: write the file into our local fs
w: write the file into our archive

*/

var sourceFile = "http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz"

const (
	ONE_MEGABYTE = 1024 * 1024
)

type counter struct {
	total uint64
}

func (c *counter) Write(b []byte) (int, error) {

	amountOfByte := uint64(len(b))
	c.total += amountOfByte
	progress := float64(c.total) / (ONE_MEGABYTE)
	fmt.Printf("\rDownloading %f MB...", progress)
	return int(amountOfByte), nil
}

func main() {

	res, err := http.Get(sourceFile)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	local, err := os.OpenFile("download-5gram.gz", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		panic(err)
	}

	defer local.Close()

	if _, err := io.Copy(local, io.TeeReader(res.Body, &counter{})); err != nil {
		panic(err)
	}

}
