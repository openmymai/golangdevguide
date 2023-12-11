package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://go.dev")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999) // second arg: number of elements
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body) // same result
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
// type Response struct {
// 	Status			string	//	e.g. "200 OK"
// 	StatusCode	int			//	e.g. 200
// 	Proto				string  //	e.g. "HTTP/1.0"
// 	ProtoMajor	int			//	e.g. 1
// 	ProtoMinor	int			//	e.g. 0

// 	...
// 	...
// 	Body	io.ReadCloser
// }

// type ReadCloser interface {
// 	Reader
// 	Closer
// }

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }