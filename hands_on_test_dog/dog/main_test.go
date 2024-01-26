package dog

import "testing"

func TestYears(t *testing.T) {
	total := Years(10)

	if total != 70 {
		t.Errorf("Error occured, expected %d, got %v", 70, total)
	}
}

func TestYearsTwo(t *testing.T) {
	total := YearsTwo(20)
	if total != 140 {
		t.Errorf("Error occured, expected %d, got %d", 140, total)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
