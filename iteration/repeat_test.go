package iteration

import (
	"fmt"
	"testing"
)

func TestIfContainsReturnString(t *testing.T) {
	t.Run("not contains", func(t *testing.T) {
		got := IfContainsReturnString("a", "b")
		want := ""
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("contains", func(t *testing.T) {
		got := IfContainsReturnString("a", "a")
		want := "a"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: bbb
}
