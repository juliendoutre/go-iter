package iter

import (
	"reflect"
	"testing"
)

func TestRangeFold(t *testing.T) {
	testCases := map[*IteratorForInt]int{
		Range(0, 0, 0):   0,
		Range(0, 1, 1):   0,
		Range(0, 4, 1):   6,
		Range(0, 5, 2):   6,
		Range(0, -4, -1): -6,
		Range(0, -5, -2): -6,
	}

	for iter, want := range testCases {
		got := iter.FoldForInt(0, func(acc, item int) int {
			return acc + item
		})

		if got != want {
			t.Errorf("case: %s;got: %d; expected: %d", iter, got, want)
		}
	}
}

func TestRangeFoldFirst(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		Range(0, 0, 0):   NoneInt(),
		Range(0, 1, 1):   SomeInt(0),
		Range(0, 4, 1):   SomeInt(6),
		Range(0, 5, 2):   SomeInt(6),
		Range(0, -4, -1): SomeInt(-6),
		Range(0, -5, -2): SomeInt(-6),
	}

	for iter, want := range testCases {
		got := iter.FoldFirst(func(acc, item int) int {
			return acc + item
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeForEach(t *testing.T) {
	testCases := map[*IteratorForInt]int{
		Range(0, 0, 0):   0,
		Range(0, 1, 1):   0,
		Range(0, 4, 1):   6,
		Range(0, 5, 2):   6,
		Range(0, -4, -1): -6,
		Range(0, -5, -2): -6,
	}

	for iter, want := range testCases {
		got := 0
		iter.ForEach(func(item int) {
			got += item
		})

		if got != want {
			t.Errorf("case: %s;got: %d; expected: %d", iter, got, want)
		}
	}
}

func TestRangeMapAndCollect(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {0},
		Range(0, 4, 1):   {0, 1, 4, 9},
		Range(0, 5, 2):   {0, 4, 16},
		Range(0, -4, -1): {0, 1, 4, 9},
		Range(0, -5, -2): {0, 4, 16},
	}

	for iter, want := range testCases {
		got := iter.Map(func(item int) int {
			return item * item
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeCount(t *testing.T) {
	testCases := map[*IteratorForInt]uint{
		Range(0, 0, 0):   0,
		Range(0, 1, 1):   1,
		Range(0, 4, 1):   4,
		Range(0, 5, 2):   3,
		Range(0, -4, -1): 4,
		Range(0, -5, -2): 3,
	}

	for iter, want := range testCases {
		got := iter.Count()

		if got != want {
			t.Errorf("case: %s;got: %d; expected: %d", iter, got, want)
		}
	}
}

func TestRangeLast(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		Range(0, 0, 0):   NoneInt(),
		Range(0, 1, 1):   SomeInt(0),
		Range(0, 4, 1):   SomeInt(3),
		Range(0, 5, 2):   SomeInt(4),
		Range(0, -4, -1): SomeInt(-3),
		Range(0, -5, -2): SomeInt(-4),
	}

	for iter, want := range testCases {
		got := iter.Last()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeNth(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		Range(0, 0, 0):   NoneInt(),
		Range(0, 1, 1):   NoneInt(),
		Range(0, 4, 1):   SomeInt(3),
		Range(0, 5, 2):   NoneInt(),
		Range(0, -4, -1): SomeInt(-3),
		Range(0, -5, -2): NoneInt(),
	}

	for iter, want := range testCases {
		got := iter.Nth(3)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeAll(t *testing.T) {
	testCases := map[*IteratorForInt]bool{
		Range(0, 0, 0):   true,
		Range(0, 1, 1):   true,
		Range(0, 4, 1):   true,
		Range(0, 5, 2):   true,
		Range(0, -4, -1): false,
		Range(0, -5, -2): false,
	}

	for iter, want := range testCases {
		got := iter.All(func(item int) bool {
			return item >= 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeAny(t *testing.T) {
	testCases := map[*IteratorForInt]bool{
		Range(0, 0, 0):   false,
		Range(0, 1, 1):   false,
		Range(0, 4, 1):   false,
		Range(0, 5, 2):   false,
		Range(0, -4, -1): true,
		Range(0, -5, -2): true,
	}

	for iter, want := range testCases {
		got := iter.Any(func(item int) bool {
			return item < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeFind(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		Range(0, 0, 0):   NoneInt(),
		Range(0, 1, 1):   NoneInt(),
		Range(0, 4, 1):   NoneInt(),
		Range(0, 5, 2):   NoneInt(),
		Range(0, -4, -1): SomeInt(-1),
		Range(0, -5, -2): SomeInt(-2),
	}
	for iter, want := range testCases {
		got := iter.Find(func(item int) bool {
			return item < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangePosition(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForUint{
		Range(0, 0, 0):   NoneUint(),
		Range(0, 1, 1):   NoneUint(),
		Range(0, 4, 1):   NoneUint(),
		Range(0, 5, 2):   NoneUint(),
		Range(0, -4, -1): SomeUint(1),
		Range(0, -5, -2): SomeUint(1),
	}

	for iter, want := range testCases {
		got := iter.Position(func(item int) bool {
			return item < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeSkipWhile(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		Range(0, 0, 0):   NoneInt(),
		Range(0, 1, 1):   NoneInt(),
		Range(0, 4, 1):   NoneInt(),
		Range(0, 5, 2):   NoneInt(),
		Range(0, -4, -1): SomeInt(-2),
		Range(0, -5, -2): SomeInt(-4),
	}

	for iter, want := range testCases {
		got := iter.SkipWhile(func(item int) bool {
			return item < 0
		}).Next()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeSkip(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		Range(0, 0, 0):   NoneInt(),
		Range(0, 1, 1):   NoneInt(),
		Range(0, 4, 1):   SomeInt(2),
		Range(0, 5, 2):   SomeInt(4),
		Range(0, -4, -1): SomeInt(-2),
		Range(0, -5, -2): SomeInt(-4),
	}

	for iter, want := range testCases {
		got := iter.Skip(2).Next()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeFilter(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {},
		Range(0, 4, 1):   {},
		Range(0, 5, 2):   {},
		Range(0, -4, -1): {-1, -2, -3},
		Range(0, -5, -2): {-2, -4},
	}

	for iter, want := range testCases {
		got := iter.Filter(func(item int) bool {
			return item < 0
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeTakeWhile(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {0},
		Range(0, 4, 1):   {0, 1, 2, 3},
		Range(0, 5, 2):   {0, 2, 4},
		Range(0, -4, -1): {0},
		Range(0, -5, -2): {0},
	}

	for iter, want := range testCases {
		got := iter.TakeWhile(func(item int) bool {
			return item >= 0
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeTake(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {0},
		Range(0, 4, 1):   {0, 1, 2},
		Range(0, 5, 2):   {0, 2, 4},
		Range(0, -4, -1): {0, -1, -2},
		Range(0, -5, -2): {0, -2, -4},
	}

	for iter, want := range testCases {
		got := iter.Take(3).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeChain(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		Range(0, 0, 0):   {-1, 0},
		Range(0, 1, 1):   {-1, 0, 0},
		Range(0, 4, 1):   {-1, 0, 0, 1, 2, 3},
		Range(0, 5, 2):   {-1, 0, 0, 2, 4},
		Range(0, -4, -1): {-1, 0, 0, -1, -2, -3},
		Range(0, -5, -2): {-1, 0, 0, -2, -4},
	}

	for iter, want := range testCases {
		got := Range(-1, 1, 1).Chain(iter).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func BenchmarkRangeDivisorsSearch(b *testing.B) {
	b.Run("with a loop", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			func() {
				results := []int{}
				for i := 0; i < 1_000_000; i++ {
					if i%14 == 0 {
						results = append(results, i)
					}
				}
			}()
		}
	})

	b.Run("with a range", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			Range(0, 1_000_000, 1).Filter(func(item int) bool {
				return item%14 == 0
			}).Collect()
		}
	})

}
