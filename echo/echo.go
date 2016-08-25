/*
Shows reading lines from stdin and writing them to stdout.

My instinct, coming from Python, was to deal with strings. But most of the
time there is not need to understand what runes ("characters") a given
sequence of butes corresponds to; there is no need to think in terms of
"strings". All I want to do is split on newline (or 0xA).
*/
package echo

import (
	"bufio"
	"fmt"
	"io"
)

// Reading strings with bufio.NewScanner
func Echo1(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Fprintln(out, scanner.Text())
	}
}

// Reading strings with bufio.NewReader
func Echo2(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Fprint(out, s)
	}
}

// Reading []byte with bufio.NewReader
func Echo3(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		out.Write(b)
	}
}

// Reading []byte with bufio.NewScanner
func Echo4(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		b := scanner.Bytes()      // beware! see bufio documentation
		out.Write(append(b, 0xA)) // append newline
	}
}
