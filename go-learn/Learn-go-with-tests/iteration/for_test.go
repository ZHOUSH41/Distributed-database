package iteration

import "testing"

func TestRepeated(t *testing.T)  {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeated(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Repeat("a")
	}
}
