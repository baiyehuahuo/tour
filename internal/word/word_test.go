package word

import "testing"

func TestToUpper(t *testing.T) {
	origin := "Hello World"
	target := "HELLO WORLD"
	if ans := ToUpper(origin); ans != target {
		t.Errorf("got %q, want %q", ans, target)
	}
}

func TestToLower(t *testing.T) {
	origin := "Hello World"
	target := "hello world"
	if ans := ToLower(origin); ans != target {
		t.Errorf("got %q, want %q", ans, target)
	}
}

func TestUnderscoreToUpperCameCase(t *testing.T) {
	origin := "hello_world"
	target := "HelloWorld"
	if ans := UnderscoreToUpperCameCase(origin); ans != target {
		t.Errorf("got %q, want %q", ans, target)
	}
}

func TestUnderscoreToLowerCameCase(t *testing.T) {
	origin := "hello_world"
	target := "helloWorld"
	if ans := UnderscoreToLowerCameCase(origin); ans != target {
		t.Errorf("got %q, want %q", ans, target)
	}
}

func TestCamelCaseToUnderscore(t *testing.T) {
	origin := "helloWorld"
	target := "hello_world"
	if ans := CamelCaseToUnderscore(origin); ans != target {
		t.Errorf("got %q, want %q", ans, target)
	}
	origin = "HelloWorld"
	if ans := CamelCaseToUnderscore(origin); ans != target {
		t.Errorf("got %q, want %q", ans, target)
	}
}
