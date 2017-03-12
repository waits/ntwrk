package main

import "testing"

type Example struct {
	Bytes   int64
	Seconds float64
	Output  string
}

var examples = [...]Example{
	{0, 1, "0.00 b/s"},
	{1, 800, "0.01 b/s"},
	{1, 8, "1.00 b/s"},
	{64, 1, "512.00 b/s"},
	{125, 1, "1.00 kb/s"},
	{375, 1.5, "2.00 kb/s"},
	{500, 1.5, "2.67 kb/s"},
	{999, 0.008, "999.00 kb/s"},
	{125000, 1, "1.00 Mb/s"},
	{62500000, 1, "500.00 Mb/s"},
	{562500000, 1, "4.50 Gb/s"},
}

func TestFormat(t *testing.T) {
	for _, ex := range examples {
		s := format(ex.Bytes, ex.Seconds)
		if s != ex.Output {
			t.Errorf("incorrect output: got %s want %s", s, ex.Output)
		}
	}
}
