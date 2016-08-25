package replace

import "testing"

type replaceFunc func(string) string

func TestReplace(t *testing.T) {
	tests := []struct {
		function replaceFunc
		name     string
		expect   string
		inputs   []string
	}{
		// "\u4e16oo" == "世oo"
		{Replace1, "Replace1", "Xoo", []string{"foo", "\u4e16oo"}},
		{Replace2, "Replace2", "Xoo", []string{"foo", "\u4e16oo"}},
		{Replace3, "Replace3", "Xoo", []string{"foo", "\u4e16oo"}},
	}
	for _, test := range tests {
		for _, input := range test.inputs {
			out := test.function(input)
			if out != test.expect {
				t.Errorf("function %q input %q expected %q got %q",
					test.name, input, test.expect, out)
			}
		}
	}
}

func BenchmarkReplace1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Replace1("foo")
	}
}

func BenchmarkReplace2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Replace2("foo")
	}
}

func BenchmarkReplace3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Replace3("foo")
	}
}
