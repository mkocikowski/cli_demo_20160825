package echo

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type echoFunc func(io.Reader, io.Writer)

func TestEcho(t *testing.T) {
	in := "foo\nbar\n"
	functions := []echoFunc{Echo1, Echo2, Echo3, Echo4}
	for _, f := range functions {
		out := new(bytes.Buffer)
		f(bytes.NewBufferString(in), out)
		b := out.Bytes()
		if !bytes.Equal([]byte(in), b) {
			t.Errorf("expected %q got %q", []byte(in), b)
		}
	}
}

func BenchmarkEcho1(b *testing.B) {
	in := "foo\nbar\n"
	for i := 0; i < b.N; i++ {
		out := new(bytes.Buffer)
		Echo1(bytes.NewBufferString(in), out)
	}
}

func BenchmarkEcho4(b *testing.B) {
	in := "foo\nbar\n"
	for i := 0; i < b.N; i++ {
		out := new(bytes.Buffer)
		Echo4(bytes.NewBufferString(in), out)
	}
}

func ExampleEcho1() {
	in := bytes.NewBufferString("hello\nworld\n") // "stdin"
	out := new(bytes.Buffer)                      // "stdout"
	Echo1(in, out)
	fmt.Print(out.String())
	// Output:
	// hello
	// world
}
