package googl_test

import (
	"testing"
	. "googl"
)

type URLTest struct {
	long string
	short string
}

var URLTests = []URLTest{
	{"http://www.google.com/", "http://goo.gl/fbsS"},
	{"http://www.twiword.com/", "http://goo.gl/Yd3lX"},
}

// use none key for test case
var googl = NewGoogl()

func TestShorten(t *testing.T) {
	for _, test := range URLTests {
		short, ok := googl.Shorten(test.long)
		if nil!=ok || short!=test.short {
			t.Errorf("Shorten(%s) = %v; want %v", test.long, short, test.short)
		}
	}
}

func TestExpand(t *testing.T) {
	for _, test := range URLTests {
		long, ok := googl.Expand(test.short)
		if nil!=ok || long!=test.long {
			t.Errorf("Expand(%s) = %v; want %v", test.short, long, test.long)
		}
	}
}
