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
	w = os.Stdout
	w = new(bytes.Buffer)
	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	w = rwc
	var s IntSet
	var _ fmt.Stringer = &s
	var any interface{}
	any = true
	any = 1234
}
