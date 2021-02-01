package iter

import (
	"reflect"
	"testing"
)

func TestRangeFold(t *testing.T) {
	testCases := map[*Iterator]int{
		Range(0, 0, 0):   0,
		Range(0, 1, 1):   0,
		Range(0, 4, 1):   6,
		Range(0, 5, 2):   6,
		Range(0, -4, -1): -6,
		Range(0, -5, -2): -6,
	}

	for iter, want := range testCases {
		got := iter.Fold(0, func(acc, item interface{}) interface{} {
			return acc.(int) + item.(int)
		}).(int)

		if got != want {
			t.Errorf("case: %s;got: %d; expected: %d", iter, got, want)
		}
	}
}

func TestRangeFoldFirst(t *testing.T) {
	testCases := map[*Iterator]Option{
		Range(0, 0, 0):   None,
		Range(0, 1, 1):   Some(0),
		Range(0, 4, 1):   Some(6),
		Range(0, 5, 2):   Some(6),
		Range(0, -4, -1): Some(-6),
		Range(0, -5, -2): Some(-6),
	}

	for iter, want := range testCases {
		got := iter.FoldFirst(func(acc, item interface{}) interface{} {
			return acc.(int) + item.(int)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeForEach(t *testing.T) {
	testCases := map[*Iterator]int{
		Range(0, 0, 0):   0,
		Range(0, 1, 1):   0,
		Range(0, 4, 1):   6,
		Range(0, 5, 2):   6,
		Range(0, -4, -1): -6,
		Range(0, -5, -2): -6,
	}

	for iter, want := range testCases {
		got := 0
		iter.ForEach(func(item interface{}) {
			got += item.(int)
		})

		if got != want {
			t.Errorf("case: %s;got: %d; expected: %d", iter, got, want)
		}
	}
}

func TestRangeMapAndCollect(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {0},
		Range(0, 4, 1):   {0, 1, 4, 9},
		Range(0, 5, 2):   {0, 4, 16},
		Range(0, -4, -1): {0, 1, 4, 9},
		Range(0, -5, -2): {0, 4, 16},
	}

	for iter, want := range testCases {
		got := iter.Map(func(item interface{}) interface{} {
			return item.(int) * item.(int)
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeCount(t *testing.T) {
	testCases := map[*Iterator]uint{
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
	testCases := map[*Iterator]Option{
		Range(0, 0, 0):   None,
		Range(0, 1, 1):   Some(0),
		Range(0, 4, 1):   Some(3),
		Range(0, 5, 2):   Some(4),
		Range(0, -4, -1): Some(-3),
		Range(0, -5, -2): Some(-4),
	}

	for iter, want := range testCases {
		got := iter.Last()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeNth(t *testing.T) {
	testCases := map[*Iterator]Option{
		Range(0, 0, 0):   None,
		Range(0, 1, 1):   None,
		Range(0, 4, 1):   Some(3),
		Range(0, 5, 2):   None,
		Range(0, -4, -1): Some(-3),
		Range(0, -5, -2): None,
	}

	for iter, want := range testCases {
		got := iter.Nth(3)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeAll(t *testing.T) {
	testCases := map[*Iterator]bool{
		Range(0, 0, 0):   true,
		Range(0, 1, 1):   true,
		Range(0, 4, 1):   true,
		Range(0, 5, 2):   true,
		Range(0, -4, -1): false,
		Range(0, -5, -2): false,
	}

	for iter, want := range testCases {
		got := iter.All(func(item interface{}) bool {
			return item.(int) >= 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeAny(t *testing.T) {
	testCases := map[*Iterator]bool{
		Range(0, 0, 0):   false,
		Range(0, 1, 1):   false,
		Range(0, 4, 1):   false,
		Range(0, 5, 2):   false,
		Range(0, -4, -1): true,
		Range(0, -5, -2): true,
	}

	for iter, want := range testCases {
		got := iter.Any(func(item interface{}) bool {
			return item.(int) < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeFind(t *testing.T) {
	testCases := map[*Iterator]Option{
		Range(0, 0, 0):   None,
		Range(0, 1, 1):   None,
		Range(0, 4, 1):   None,
		Range(0, 5, 2):   None,
		Range(0, -4, -1): Some(-1),
		Range(0, -5, -2): Some(-2),
	}
	for iter, want := range testCases {
		got := iter.Find(func(item interface{}) bool {
			return item.(int) < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangePosition(t *testing.T) {
	testCases := map[*Iterator]Option{
		Range(0, 0, 0):   None,
		Range(0, 1, 1):   None,
		Range(0, 4, 1):   None,
		Range(0, 5, 2):   None,
		Range(0, -4, -1): Some(uint(1)),
		Range(0, -5, -2): Some(uint(1)),
	}

	for iter, want := range testCases {
		got := iter.Position(func(item interface{}) bool {
			return item.(int) < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeSkipWhile(t *testing.T) {
	testCases := map[*Iterator]Option{
		Range(0, 0, 0):   None,
		Range(0, 1, 1):   None,
		Range(0, 4, 1):   None,
		Range(0, 5, 2):   None,
		Range(0, -4, -1): Some(-2),
		Range(0, -5, -2): Some(-4),
	}

	for iter, want := range testCases {
		got := iter.SkipWhile(func(item interface{}) bool {
			return item.(int) < 0
		}).Next()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeSkip(t *testing.T) {
	testCases := map[*Iterator]Option{
		Range(0, 0, 0):   None,
		Range(0, 1, 1):   None,
		Range(0, 4, 1):   Some(2),
		Range(0, 5, 2):   Some(4),
		Range(0, -4, -1): Some(-2),
		Range(0, -5, -2): Some(-4),
	}

	for iter, want := range testCases {
		got := iter.Skip(2).Next()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeFilter(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {},
		Range(0, 4, 1):   {},
		Range(0, 5, 2):   {},
		Range(0, -4, -1): {-1, -2, -3},
		Range(0, -5, -2): {-2, -4},
	}

	for iter, want := range testCases {
		got := iter.Filter(func(item interface{}) bool {
			return item.(int) < 0
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeEnumerate(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {Enumeration{0, 0}},
		Range(0, 4, 1):   {Enumeration{0, 0}, Enumeration{1, 1}, Enumeration{2, 2}, Enumeration{3, 3}},
		Range(0, 5, 2):   {Enumeration{0, 0}, Enumeration{1, 2}, Enumeration{2, 4}},
		Range(0, -4, -1): {Enumeration{0, 0}, Enumeration{1, -1}, Enumeration{2, -2}, Enumeration{3, -3}},
		Range(0, -5, -2): {Enumeration{0, 0}, Enumeration{1, -2}, Enumeration{2, -4}},
	}

	for iter, want := range testCases {
		got := iter.Enumerate().Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeTakeWhile(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {0},
		Range(0, 4, 1):   {0, 1, 2, 3},
		Range(0, 5, 2):   {0, 2, 4},
		Range(0, -4, -1): {0},
		Range(0, -5, -2): {0},
	}

	for iter, want := range testCases {
		got := iter.TakeWhile(func(item interface{}) bool {
			return item.(int) >= 0
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestRangeTake(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
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
	testCases := map[*Iterator][]interface{}{
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

func TestRangeZip(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Range(0, 0, 0):   {},
		Range(0, 1, 1):   {Pair{-1, 0}},
		Range(0, 4, 1):   {Pair{-1, 0}, Pair{0, 1}},
		Range(0, 5, 2):   {Pair{-1, 0}, Pair{0, 2}},
		Range(0, -4, -1): {Pair{-1, 0}, Pair{0, -1}},
		Range(0, -5, -2): {Pair{-1, 0}, Pair{0, -2}},
	}

	for iter, want := range testCases {
		got := Range(-1, 1, 1).Zip(iter).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s;got: %v; expected: %v", iter, got, want)
		}
	}
}
