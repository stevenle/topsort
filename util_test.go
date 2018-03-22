package topsort

import "testing"

func compareOrError(a, b []string, t *testing.T) {
	if len(a) != len(b) {
		t.Error("Slices has different sizes.")
	}
	for i, av := range a {
		if b[i] != av {
			t.Errorf("a[%[1]v]=%[2]q not eq b[%[1]v]=%[3]q.", i, av, b[i])
		}
	}
}

func TestReverseEmpty(t *testing.T) {
	values := []string{}
	Reverse(values)
	compareOrError([]string{}, values, t)
}

func TestReverseOneValue(t *testing.T) {
	values := []string{"a"}
	Reverse(values)
	compareOrError([]string{"a"}, values, t)
}

func TestReverseTwoValue(t *testing.T) {
	values := []string{"a", "b"}
	Reverse(values)
	compareOrError([]string{"b", "a"}, values, t)
}

func TestReverseManyValues(t *testing.T) {
	values := []string{"a", "b", "c", "d"}
	Reverse(values)
	compareOrError([]string{"d", "c", "b", "a"}, values, t)
}