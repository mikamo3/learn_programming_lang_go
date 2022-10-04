package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type IntSet struct{}

func (*IntSet) String() string {
	return "a"
}

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}
type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
	Resolution() (x, y int)
}

func main() {
	var w io.Writer
	fmt.Printf("%T", w)
	fmt.Println(w == nil)
	w = os.Stdout
	w.Write([]byte("hello"))
	fmt.Printf("%T", w)
	w = new(bytes.Buffer)
	fmt.Printf("%T", w)
	var buf *bytes.Buffer
	f(buf)
	w = nil
	f(w)
}
func f(out io.Writer) {
	fmt.Println(out != nil)
}
