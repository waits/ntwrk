package main

import "testing"

type Example struct {
	Bytes   int
	Seconds float64
	Output  string
}

var examples = [...]Example{
	{0, 1, "0.00 B/s"},
	{1, 100, "0.01 B/s"},
	{1, 1, "1.00 B/s"},
	{512, 1, "512.00 B/s"},
	{1024, 1, "1.00 kB/s"},
	{3072, 1.5, "2.00 kB/s"},
	{4096, 1.5, "2.67 kB/s"},
	{1024, 0.001, "1000.00 kB/s"},
	{1048565, 1, "1023.99 kB/s"},
	{524288000, 1, "500.00 MB/s"},
	{4831838208, 1, "4.50 GB/s"},
}

func TestFormat(t *testing.T) {
	for _, ex := range examples {
		s := format(ex.Bytes, ex.Seconds)
		if s != ex.Output {
			t.Errorf("incorrect output: got %s want %s", s, ex.Output)
		}
	}
}
