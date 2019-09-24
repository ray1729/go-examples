package adder

import (
  "math/rand"
  "reflect"
  "testing"
  "testing/quick"
)

func BenchmarkAccumAtomic(b *testing.B) {
  for n := 0; n < b.N; n++ {
    AccumAtomic(10000)
  }
}

func BenchmarkAccumMutex(b *testing.B) {
  for n := 0; n < b.N; n++ {
    AccumMutex(10000)
  }
}

func TriangularNumber(n uint64) uint64 {
  return n*(n+1)/2
}

var config = &quick.Config{
  Values: func(args []reflect.Value, r *rand.Rand) {
    args[0] = reflect.ValueOf(r.Uint64() % 1001)
  },
}

func TestAccumAtomic(t *testing.T) {
  if err := quick.CheckEqual(AccumAtomic, TriangularNumber, config); err != nil {
    t.Error(err)
  }
}

// func TestAccumAtomicVariant(t *testing.T) {
//   err := quick.Check(func(n uint64) bool {
//     m := n % 1001
//     return AccumAtomic(m) == TriangularNumber(m)
//   }, nil)
//   if err != nil {
//     t.Error(err)
//   }
// }

func TestAccumMutex(t *testing.T) {
  if err := quick.CheckEqual(AccumMutex, TriangularNumber, config); err != nil {
    t.Error(err)
  }
}
