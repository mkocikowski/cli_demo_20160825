package cli_demo_20160825

import (
	"bufio"
	"fmt"
	"io"
)

func Echo1(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Fprintln(out, scanner.Text())
	}
}

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

func Echo4(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		b := scanner.Bytes()      // beware! see documentation
		out.Write(append(b, 0xA)) // append newline
	}
}

//func main() {
//	Echo1(os.Stdin, os.Stdout)
//}
