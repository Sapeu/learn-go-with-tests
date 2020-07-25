package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectSum := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectSum(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertCorrectSum(t, got, want, numbers)
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		// 尝试比较 slice 和 string。这显然是不合理的，但是却通过了编译！所以使用 reflect.DeepEqual 比较简洁但是在使用时需多加小心
		// want := "bob"
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		// 尝试比较 slice 和 string。这显然是不合理的，但是却通过了编译！所以使用 reflect.DeepEqual 比较简洁但是在使用时需多加小心
		// want := "bob"
		checkSums(t, got, want)
	})
}
