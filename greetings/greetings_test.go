package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(name)

	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(
			`Hello(%v) = %q, %v, want match for %#q, nil`,
			name,
			msg,
			err,
			want,
		)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""

	msg, err := Hello(name)

	if msg != "" || err == nil {
		t.Fatalf(
			`Hello(%v) = %q, %v, want match for nil, error`,
			name,
			msg,
			err,
		)
	}
}
