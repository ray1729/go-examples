package adder

import (
  "sync"
  "sync/atomic"
)

func AccumAtomic(n uint64) uint64 {
  var accum uint64
  var wg sync.WaitGroup

  for i := uint64(1); i <= n; i++ {
    wg.Add(1)
    go func(delta uint64) {
      defer wg.Done()
      atomic.AddUint64(&accum, delta)
    }(i)
  }

  wg.Wait()
  return accum
}

func AccumMutex(n uint64) uint64 {
  var accum uint64
  var mu sync.Mutex
  var wg sync.WaitGroup

  for i := uint64(1); i <= n; i++ {
    wg.Add(1)
    go func(delta uint64) {
      defer wg.Done()
      mu.Lock()
      accum += delta
      mu.Unlock()
    }(i)
  }

  wg.Wait()
  return accum
}
