package iter

import (
	"reflect"
	"testing"
)

func TestVectorFold(t *testing.T) {
	testCases := map[*IteratorForInt]int{
		VectorOfInt([]int{}):            0,
		VectorOfInt([]int{0}):           0,
		VectorOfInt([]int{0, 1, 2, 3}):  6,
		VectorOfInt([]int{0, 1, -2, 3}): 2,
	}

	for iter, want := range testCases {
		got := iter.FoldForInt(0, func(acc, item int) int {
			return acc + item
		})

		if got != want {
			t.Errorf("case: %s; got:%d; expected: %d", iter, got, want)
		}
	}
}

func TestVectorFoldFirst(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		VectorOfInt([]int{}):            NoneInt(),
		VectorOfInt([]int{0}):           SomeInt(0),
		VectorOfInt([]int{0, 1, 2, 3}):  SomeInt(6),
		VectorOfInt([]int{0, 1, -2, 3}): SomeInt(2),
	}

	for iter, want := range testCases {
		got := iter.FoldFirst(func(acc, item int) int {
			return acc + item
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorForEach(t *testing.T) {
	testCases := map[*IteratorForInt]int{
		VectorOfInt([]int{}):            0,
		VectorOfInt([]int{0}):           0,
		VectorOfInt([]int{0, 1, 2, 3}):  6,
		VectorOfInt([]int{0, 1, -2, 3}): 2,
	}

	for iter, want := range testCases {
		got := 0
		iter.ForEach(func(item int) {
			got += item
		})

		if got != want {
			t.Errorf("case: %s; got:%d; expected: %d", iter, got, want)
		}
	}
}

func TestVectorMapAndCollect(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		VectorOfInt([]int{}):            {},
		VectorOfInt([]int{0}):           {0},
		VectorOfInt([]int{0, 1, 2, 3}):  {0, 1, 4, 9},
		VectorOfInt([]int{0, 1, -2, 3}): {0, 1, 4, 9},
	}

	for iter, want := range testCases {
		got := iter.Map(func(item int) int {
			return item * item
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorCount(t *testing.T) {
	testCases := map[*IteratorForInt]uint{
		VectorOfInt([]int{}):            0,
		VectorOfInt([]int{0}):           1,
		VectorOfInt([]int{0, 1, 2, 3}):  4,
		VectorOfInt([]int{0, 1, -2, 3}): 4,
	}

	for iter, want := range testCases {
		got := iter.Count()

		if got != want {
			t.Errorf("case: %s; got:%d; expected: %d", iter, got, want)
		}
	}
}

func TestVectorLast(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		VectorOfInt([]int{}):            NoneInt(),
		VectorOfInt([]int{0}):           SomeInt(0),
		VectorOfInt([]int{0, 1, 2, 3}):  SomeInt(3),
		VectorOfInt([]int{0, 1, -2, 3}): SomeInt(3),
	}

	for iter, want := range testCases {
		got := iter.Last()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorNth(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		VectorOfInt([]int{}):            NoneInt(),
		VectorOfInt([]int{0}):           NoneInt(),
		VectorOfInt([]int{0, 1, 2, 3}):  SomeInt(3),
		VectorOfInt([]int{0, 1, -2, 3}): SomeInt(3),
	}

	for iter, want := range testCases {
		got := iter.Nth(3)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorAll(t *testing.T) {
	testCases := map[*IteratorForInt]bool{
		VectorOfInt([]int{}):            true,
		VectorOfInt([]int{0}):           true,
		VectorOfInt([]int{0, 1, 2, 3}):  true,
		VectorOfInt([]int{0, 1, -2, 3}): false,
	}

	for iter, want := range testCases {
		got := iter.All(func(item int) bool {
			return item >= 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorAny(t *testing.T) {
	testCases := map[*IteratorForInt]bool{
		VectorOfInt([]int{}):            false,
		VectorOfInt([]int{0}):           false,
		VectorOfInt([]int{0, 1, 2, 3}):  false,
		VectorOfInt([]int{0, 1, -2, 3}): true,
	}

	for iter, want := range testCases {
		got := iter.Any(func(item int) bool {
			return item < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorFind(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		VectorOfInt([]int{}):            NoneInt(),
		VectorOfInt([]int{0}):           NoneInt(),
		VectorOfInt([]int{0, 1, 2, 3}):  NoneInt(),
		VectorOfInt([]int{0, 1, -2, 3}): SomeInt(-2),
	}

	for iter, want := range testCases {
		got := iter.Find(func(item int) bool {
			return item < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorPosition(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForUint{
		VectorOfInt([]int{}):            NoneUint(),
		VectorOfInt([]int{0}):           NoneUint(),
		VectorOfInt([]int{0, 1, 2, 3}):  NoneUint(),
		VectorOfInt([]int{0, 1, -2, 3}): SomeUint(2),
	}

	for iter, want := range testCases {
		got := iter.Position(func(item int) bool {
			return item < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorSkipWhile(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		VectorOfInt([]int{}):            NoneInt(),
		VectorOfInt([]int{0}):           NoneInt(),
		VectorOfInt([]int{0, 1, 2, 3}):  NoneInt(),
		VectorOfInt([]int{0, 1, -2, 3}): SomeInt(3),
	}

	for iter, want := range testCases {
		got := iter.SkipWhile(func(item int) bool {
			return item < 0
		}).Next()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorSkip(t *testing.T) {
	testCases := map[*IteratorForInt]OptionForInt{
		VectorOfInt([]int{}):            NoneInt(),
		VectorOfInt([]int{0}):           NoneInt(),
		VectorOfInt([]int{0, 1, 2, 3}):  SomeInt(2),
		VectorOfInt([]int{0, 1, -2, 3}): SomeInt(-2),
	}

	for iter, want := range testCases {
		got := iter.Skip(2).Next()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorFilter(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		VectorOfInt([]int{}):            {},
		VectorOfInt([]int{0}):           {},
		VectorOfInt([]int{0, 1, 2, 3}):  {},
		VectorOfInt([]int{0, 1, -2, 3}): {-2},
	}

	for iter, want := range testCases {
		got := iter.Filter(func(item int) bool {
			return item < 0
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorTakeWhile(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		VectorOfInt([]int{}):            {},
		VectorOfInt([]int{0}):           {0},
		VectorOfInt([]int{0, 1, 2, 3}):  {0, 1, 2, 3},
		VectorOfInt([]int{0, 1, -2, 3}): {0, 1},
	}

	for iter, want := range testCases {
		got := iter.TakeWhile(func(item int) bool {
			return item >= 0
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorTake(t *testing.T) {
	testCases := map[*IteratorForInt][]int{
		VectorOfInt([]int{}):            {},
		VectorOfInt([]int{0}):           {0},
		VectorOfInt([]int{0, 1, 2, 3}):  {0, 1, 2},
		VectorOfInt([]int{0, 1, -2, 3}): {0, 1, -2},
	}

	for iter, want := range testCases {
		got := iter.Take(3).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorChain(t *testing.T) {
	base := []int{-1, -6}

	testCases := map[*IteratorForInt][]int{
		VectorOfInt([]int{}):            {-1, -6},
		VectorOfInt([]int{0}):           {-1, -6, 0},
		VectorOfInt([]int{0, 1, 2, 3}):  {-1, -6, 0, 1, 2, 3},
		VectorOfInt([]int{0, 1, -2, 3}): {-1, -6, 0, 1, -2, 3},
	}

	for iter, want := range testCases {
		got := VectorOfInt(base).Chain(iter).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func BenchmarkVectorStringSearch(b *testing.B) {
	text := []string{
		"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit.", "Ut", "tincidunt", "felis", "at", "purus", "congue,", "eu", "sollicitudin", "elit", "condimentum.", "Morbi", "efficitur", "egestas", "porta.", "Suspendisse", "quis", "tellus", "facilisis,", "ultricies", "dolor", "a,", "eleifend", "nisi.", "Suspendisse", "euismod", "metus", "mi,", "quis", "porttitor", "turpis", "auctor", "blandit.", "Cras", "ut", "lobortis", "massa.", "Donec", "dignissim", "pretium", "nisi,", "sed", "tincidunt", "urna", "porttitor", "nec.", "Phasellus", "vulputate", "tincidunt", "fermentum.", "Pellentesque", "at", "lobortis", "ante.", "Donec", "arcu", "ligula,", "pharetra", "a", "congue", "sed,", "ultricies", "vitae", "felis.", "Phasellus", "interdum", "quam", "sit", "amet", "libero", "elementum", "molestie.", "Integer", "porta", "felis", "vitae", "risus", "laoreet", "cursus.", "Mauris", "libero", "odio,", "eleifend", "eu", "mauris", "sit", "amet,", "laoreet", "volutpat", "quam.", "Etiam", "dictum", "diam", "vel", "laoreet", "feugiat.", "Vestibulum", "ante", "ipsum", "primis", "in", "faucibus", "orci", "luctus", "et", "ultrices", "posuere", "cubilia", "curae;", "Fusce", "suscipit", "posuere", "nunc", "id", "consequat.", "Etiam", "nulla", "nunc,", "tincidunt", "nec", "rhoncus", "vel,", "ultrices", "vel", "massa.", "In", "hac", "habitasse", "platea", "dictumst.", "Mauris", "quis", "dui", "a", "lacus", "varius", "molestie.", "Cras", "sollicitudin", "a", "orci", "eu", "feugiat.", "Integer", "cursus", "justo", "quis", "felis", "tincidunt", "iaculis.", "Phasellus", "feugiat", "vitae", "justo", "eu", "dignissim.", "Duis", "ut", "euismod", "metus.", "Fusce", "id", "justo", "ante.", "Mauris", "sit", "amet", "efficitur", "mauris.", "Mauris", "et", "enim", "at", "turpis", "volutpat", "semper.", "Donec", "fringilla", "nibh", "ante,", "lacinia", "condimentum", "velit", "viverra", "ut.", "Proin", "quis", "dolor", "vel", "tellus", "facilisis", "cursus", "non", "eu", "ligula.", "Maecenas", "malesuada", "lacus", "sit", "amet", "magna", "facilisis", "efficitur.", "Pellentesque", "a", "interdum", "purus.", "Pellentesque", "vulputate", "consequat", "enim,", "viverra", "fermentum", "augue", "pharetra", "vitae.", "Sed", "vitae", "nulla", "nec", "tortor", "molestie", "iaculis.", "Nunc", "pharetra", "feugiat", "odio,", "vitae", "tempus", "neque", "faucibus", "eget.", "Nulla", "rutrum", "suscipit", "tincidunt.", "Sed", "semper", "tellus", "at", "diam", "sodales", "tincidunt", "eget", "sed", "libero.", "In", "egestas", "mi", "in", "odio", "blandit", "laoreet.", "Praesent", "accumsan", "metus", "vitae", "facilisis", "lacinia.", "Vivamus", "sed", "enim", "a", "nisi", "varius", "venenatis", "id", "consectetur", "risus.", "Vestibulum", "non", "dolor", "feugiat,", "pellentesque", "est", "ac,", "maximus", "neque.", "Pellentesque", "consequat", "tellus", "a", "consectetur", "porta.", "Nunc", "iaculis", "et", "arcu", "nec", "dapibus.", "Maecenas", "id", "lacinia", "nisi,", "non", "tincidunt", "orci.", "Maecenas", "et", "dui", "in", "libero", "fermentum", "egestas.", "Etiam", "rutrum", "ligula", "ipsum,", "vel", "gravida", "lorem", "pellentesque", "sed.", "Quisque", "sit", "amet", "facilisis", "libero.", "Curabitur", "semper", "quam", "a", "leo", "fringilla", "maximus.", "Nullam", "eros", "leo,", "pretium", "non", "volutpat", "volutpat,", "feugiat", "porta", "diam.", "Quisque", "odio", "metus,", "varius", "et", "iaculis", "vitae,", "gravida", "eu", "diam.", "Etiam", "pellentesque", "faucibus", "lorem,", "quis", "iaculis", "metus.", "Nullam", "vehicula", "consectetur", "lacus,", "id", "sodales", "diam", "auctor", "luctus.", "Proin", "sit", "amet", "ante", "nisi.", "Donec", "varius", "egestas", "consectetur.", "Ut", "a", "lacus", "eros.", "Phasellus", "vestibulum", "enim", "sit", "amet", "purus", "scelerisque", "bibendum.", "Integer", "lacus", "sapien,", "tempus", "id", "commodo", "in,", "blandit", "vitae", "arcu.", "Aenean", "at", "ornare", "nunc.",
	}

	b.Run("with a loop", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			func() {
				for _, item := range text {
					if item == "nunc." {
						break
					}
				}
			}()
		}
	})

	b.Run("with a VectorOfInt", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			VectorOfString(text).Find(func(item string) bool {
				return item == "nunc."
			})
		}
	})
}
