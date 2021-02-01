package iter

import (
	"reflect"
	"testing"
)

func TestVectorFold(t *testing.T) {
	testCases := map[*Iterator]int{
		Vector([]interface{}{}):            0,
		Vector([]interface{}{0}):           0,
		Vector([]interface{}{0, 1, 2, 3}):  6,
		Vector([]interface{}{0, 1, -2, 3}): 2,
	}

	for iter, want := range testCases {
		got := iter.Fold(0, func(acc, item interface{}) interface{} {
			return acc.(int) + item.(int)
		}).(int)

		if got != want {
			t.Errorf("case: %s; got:%d; expected: %d", iter, got, want)
		}
	}
}

func TestVectorFoldFirst(t *testing.T) {
	testCases := map[*Iterator]interface{}{
		Vector([]interface{}{}):            nil,
		Vector([]interface{}{0}):           0,
		Vector([]interface{}{0, 1, 2, 3}):  6,
		Vector([]interface{}{0, 1, -2, 3}): 2,
	}

	for iter, want := range testCases {
		got := iter.FoldFirst(func(acc, item interface{}) interface{} {
			return acc.(int) + item.(int)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorForEach(t *testing.T) {
	testCases := map[*Iterator]int{
		Vector([]interface{}{}):            0,
		Vector([]interface{}{0}):           0,
		Vector([]interface{}{0, 1, 2, 3}):  6,
		Vector([]interface{}{0, 1, -2, 3}): 2,
	}

	for iter, want := range testCases {
		got := 0
		iter.ForEach(func(item interface{}) {
			got += item.(int)
		})

		if got != want {
			t.Errorf("case: %s; got:%d; expected: %d", iter, got, want)
		}
	}
}

func TestVectorMapAndCollect(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Vector([]interface{}{}):            {},
		Vector([]interface{}{0}):           {0},
		Vector([]interface{}{0, 1, 2, 3}):  {0, 1, 4, 9},
		Vector([]interface{}{0, 1, -2, 3}): {0, 1, 4, 9},
	}

	for iter, want := range testCases {
		got := iter.Map(func(item interface{}) interface{} {
			return item.(int) * item.(int)
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorCount(t *testing.T) {
	testCases := map[*Iterator]uint{
		Vector([]interface{}{}):            0,
		Vector([]interface{}{0}):           1,
		Vector([]interface{}{0, 1, 2, 3}):  4,
		Vector([]interface{}{0, 1, -2, 3}): 4,
	}

	for iter, want := range testCases {
		got := iter.Count()

		if got != want {
			t.Errorf("case: %s; got:%d; expected: %d", iter, got, want)
		}
	}
}

func TestVectorLast(t *testing.T) {
	testCases := map[*Iterator]interface{}{
		Vector([]interface{}{}):            nil,
		Vector([]interface{}{0}):           0,
		Vector([]interface{}{0, 1, 2, 3}):  3,
		Vector([]interface{}{0, 1, -2, 3}): 3,
	}

	for iter, want := range testCases {
		got := iter.Last()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorNth(t *testing.T) {
	testCases := map[*Iterator]interface{}{
		Vector([]interface{}{}):            nil,
		Vector([]interface{}{0}):           nil,
		Vector([]interface{}{0, 1, 2, 3}):  3,
		Vector([]interface{}{0, 1, -2, 3}): 3,
	}

	for iter, want := range testCases {
		got := iter.Nth(3)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorAll(t *testing.T) {
	testCases := map[*Iterator]bool{
		Vector([]interface{}{}):            true,
		Vector([]interface{}{0}):           true,
		Vector([]interface{}{0, 1, 2, 3}):  true,
		Vector([]interface{}{0, 1, -2, 3}): false,
	}

	for iter, want := range testCases {
		got := iter.All(func(item interface{}) bool {
			return item.(int) >= 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorAny(t *testing.T) {
	testCases := map[*Iterator]bool{
		Vector([]interface{}{}):            false,
		Vector([]interface{}{0}):           false,
		Vector([]interface{}{0, 1, 2, 3}):  false,
		Vector([]interface{}{0, 1, -2, 3}): true,
	}

	for iter, want := range testCases {
		got := iter.Any(func(item interface{}) bool {
			return item.(int) < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorFind(t *testing.T) {
	testCases := map[*Iterator]interface{}{
		Vector([]interface{}{}):            nil,
		Vector([]interface{}{0}):           nil,
		Vector([]interface{}{0, 1, 2, 3}):  nil,
		Vector([]interface{}{0, 1, -2, 3}): -2,
	}

	for iter, want := range testCases {
		got := iter.Find(func(item interface{}) bool {
			return item.(int) < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorPosition(t *testing.T) {
	testCases := map[*Iterator]interface{}{
		Vector([]interface{}{}):            nil,
		Vector([]interface{}{0}):           nil,
		Vector([]interface{}{0, 1, 2, 3}):  nil,
		Vector([]interface{}{0, 1, -2, 3}): uint(2),
	}

	for iter, want := range testCases {
		got := iter.Position(func(item interface{}) bool {
			return item.(int) < 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorSkipWhile(t *testing.T) {
	testCases := map[*Iterator]interface{}{
		Vector([]interface{}{}):            nil,
		Vector([]interface{}{0}):           nil,
		Vector([]interface{}{0, 1, 2, 3}):  nil,
		Vector([]interface{}{0, 1, -2, 3}): 3,
	}

	for iter, want := range testCases {
		got := iter.SkipWhile(func(item interface{}) bool {
			return item.(int) < 0
		}).Next()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorSkip(t *testing.T) {
	testCases := map[*Iterator]interface{}{
		Vector([]interface{}{}):            nil,
		Vector([]interface{}{0}):           nil,
		Vector([]interface{}{0, 1, 2, 3}):  2,
		Vector([]interface{}{0, 1, -2, 3}): -2,
	}

	for iter, want := range testCases {
		got := iter.Skip(2).Next()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorFilter(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Vector([]interface{}{}):            {},
		Vector([]interface{}{0}):           {},
		Vector([]interface{}{0, 1, 2, 3}):  {},
		Vector([]interface{}{0, 1, -2, 3}): {-2},
	}

	for iter, want := range testCases {
		got := iter.Filter(func(item interface{}) bool {
			return item.(int) < 0
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorEnumerate(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Vector([]interface{}{}):            {},
		Vector([]interface{}{0}):           {Enumeration{0, 0}},
		Vector([]interface{}{0, 1, 2, 3}):  {Enumeration{0, 0}, Enumeration{1, 1}, Enumeration{2, 2}, Enumeration{3, 3}},
		Vector([]interface{}{0, 1, -2, 3}): {Enumeration{0, 0}, Enumeration{1, 1}, Enumeration{2, -2}, Enumeration{3, 3}},
	}

	for iter, want := range testCases {
		got := iter.Enumerate().Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorTakeWhile(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Vector([]interface{}{}):            {},
		Vector([]interface{}{0}):           {0},
		Vector([]interface{}{0, 1, 2, 3}):  {0, 1, 2, 3},
		Vector([]interface{}{0, 1, -2, 3}): {0, 1},
	}

	for iter, want := range testCases {
		got := iter.TakeWhile(func(item interface{}) bool {
			return item.(int) >= 0
		}).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorTake(t *testing.T) {
	testCases := map[*Iterator][]interface{}{
		Vector([]interface{}{}):            {},
		Vector([]interface{}{0}):           {0},
		Vector([]interface{}{0, 1, 2, 3}):  {0, 1, 2},
		Vector([]interface{}{0, 1, -2, 3}): {0, 1, -2},
	}

	for iter, want := range testCases {
		got := iter.Take(3).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorChain(t *testing.T) {
	base := []interface{}{-1, -6}

	testCases := map[*Iterator][]interface{}{
		Vector([]interface{}{}):            {-1, -6},
		Vector([]interface{}{0}):           {-1, -6, 0},
		Vector([]interface{}{0, 1, 2, 3}):  {-1, -6, 0, 1, 2, 3},
		Vector([]interface{}{0, 1, -2, 3}): {-1, -6, 0, 1, -2, 3},
	}

	for iter, want := range testCases {
		got := Vector(base).Chain(iter).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func TestVectorZip(t *testing.T) {
	base := []interface{}{-1, -6}

	testCases := map[*Iterator][]interface{}{
		Vector([]interface{}{}):            {},
		Vector([]interface{}{0}):           {Pair{-1, 0}},
		Vector([]interface{}{0, 1, 2, 3}):  {Pair{-1, 0}, Pair{-6, 1}},
		Vector([]interface{}{0, 1, -2, 3}): {Pair{-1, 0}, Pair{-6, 1}},
	}

	for iter, want := range testCases {
		got := Vector(base).Zip(iter).Collect()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("case: %s; got: %v; expected: %v", iter, got, want)
		}
	}
}

func BenchmarkVectorStringSearch(b *testing.B) {
	text := []interface{}{
		"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit.", "Ut", "tincidunt", "felis", "at", "purus", "congue,", "eu", "sollicitudin", "elit", "condimentum.", "Morbi", "efficitur", "egestas", "porta.", "Suspendisse", "quis", "tellus", "facilisis,", "ultricies", "dolor", "a,", "eleifend", "nisi.", "Suspendisse", "euismod", "metus", "mi,", "quis", "porttitor", "turpis", "auctor", "blandit.", "Cras", "ut", "lobortis", "massa.", "Donec", "dignissim", "pretium", "nisi,", "sed", "tincidunt", "urna", "porttitor", "nec.", "Phasellus", "vulputate", "tincidunt", "fermentum.", "Pellentesque", "at", "lobortis", "ante.", "Donec", "arcu", "ligula,", "pharetra", "a", "congue", "sed,", "ultricies", "vitae", "felis.", "Phasellus", "interdum", "quam", "sit", "amet", "libero", "elementum", "molestie.", "Integer", "porta", "felis", "vitae", "risus", "laoreet", "cursus.", "Mauris", "libero", "odio,", "eleifend", "eu", "mauris", "sit", "amet,", "laoreet", "volutpat", "quam.", "Etiam", "dictum", "diam", "vel", "laoreet", "feugiat.", "Vestibulum", "ante", "ipsum", "primis", "in", "faucibus", "orci", "luctus", "et", "ultrices", "posuere", "cubilia", "curae;", "Fusce", "suscipit", "posuere", "nunc", "id", "consequat.", "Etiam", "nulla", "nunc,", "tincidunt", "nec", "rhoncus", "vel,", "ultrices", "vel", "massa.", "In", "hac", "habitasse", "platea", "dictumst.", "Mauris", "quis", "dui", "a", "lacus", "varius", "molestie.", "Cras", "sollicitudin", "a", "orci", "eu", "feugiat.", "Integer", "cursus", "justo", "quis", "felis", "tincidunt", "iaculis.", "Phasellus", "feugiat", "vitae", "justo", "eu", "dignissim.", "Duis", "ut", "euismod", "metus.", "Fusce", "id", "justo", "ante.", "Mauris", "sit", "amet", "efficitur", "mauris.", "Mauris", "et", "enim", "at", "turpis", "volutpat", "semper.", "Donec", "fringilla", "nibh", "ante,", "lacinia", "condimentum", "velit", "viverra", "ut.", "Proin", "quis", "dolor", "vel", "tellus", "facilisis", "cursus", "non", "eu", "ligula.", "Maecenas", "malesuada", "lacus", "sit", "amet", "magna", "facilisis", "efficitur.", "Pellentesque", "a", "interdum", "purus.", "Pellentesque", "vulputate", "consequat", "enim,", "viverra", "fermentum", "augue", "pharetra", "vitae.", "Sed", "vitae", "nulla", "nec", "tortor", "molestie", "iaculis.", "Nunc", "pharetra", "feugiat", "odio,", "vitae", "tempus", "neque", "faucibus", "eget.", "Nulla", "rutrum", "suscipit", "tincidunt.", "Sed", "semper", "tellus", "at", "diam", "sodales", "tincidunt", "eget", "sed", "libero.", "In", "egestas", "mi", "in", "odio", "blandit", "laoreet.", "Praesent", "accumsan", "metus", "vitae", "facilisis", "lacinia.", "Vivamus", "sed", "enim", "a", "nisi", "varius", "venenatis", "id", "consectetur", "risus.", "Vestibulum", "non", "dolor", "feugiat,", "pellentesque", "est", "ac,", "maximus", "neque.", "Pellentesque", "consequat", "tellus", "a", "consectetur", "porta.", "Nunc", "iaculis", "et", "arcu", "nec", "dapibus.", "Maecenas", "id", "lacinia", "nisi,", "non", "tincidunt", "orci.", "Maecenas", "et", "dui", "in", "libero", "fermentum", "egestas.", "Etiam", "rutrum", "ligula", "ipsum,", "vel", "gravida", "lorem", "pellentesque", "sed.", "Quisque", "sit", "amet", "facilisis", "libero.", "Curabitur", "semper", "quam", "a", "leo", "fringilla", "maximus.", "Nullam", "eros", "leo,", "pretium", "non", "volutpat", "volutpat,", "feugiat", "porta", "diam.", "Quisque", "odio", "metus,", "varius", "et", "iaculis", "vitae,", "gravida", "eu", "diam.", "Etiam", "pellentesque", "faucibus", "lorem,", "quis", "iaculis", "metus.", "Nullam", "vehicula", "consectetur", "lacus,", "id", "sodales", "diam", "auctor", "luctus.", "Proin", "sit", "amet", "ante", "nisi.", "Donec", "varius", "egestas", "consectetur.", "Ut", "a", "lacus", "eros.", "Phasellus", "vestibulum", "enim", "sit", "amet", "purus", "scelerisque", "bibendum.", "Integer", "lacus", "sapien,", "tempus", "id", "commodo", "in,", "blandit", "vitae", "arcu.", "Aenean", "at", "ornare", "nunc.",
	}

	b.Run("with a loop", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			func() {
				for _, item := range text {
					if item.(string) == "nunc." {
						break
					}
				}
			}()
		}
	})

	b.Run("with a vector", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			Vector(text).Find(func(item interface{}) bool {
				return item.(string) == "nunc."
			})
		}
	})

}
