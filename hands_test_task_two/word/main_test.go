package word

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	totalWords := Count("We are testing data in this")

	if totalWords != 6 {
		t.Errorf("Error occured, expected %d, got %d", 6, totalWords)

	}
}

func TestUseCount(t *testing.T) {
	totalMapCount := UseCount("We are testing data in this data")

	compare := make(map[string]int)

	compare["We"] = 1
	compare["are"] = 1
	compare["testing"] = 1
	compare["data"] = 2
	compare["in"] = 1
	compare["this"] = 1

	if !reflect.DeepEqual(totalMapCount, compare) {
		t.Errorf("Error occured, expected %v, got %v", compare, totalMapCount)

	}
}
func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("We are testing data in this")

	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("We are testing data in this")

	}
}
