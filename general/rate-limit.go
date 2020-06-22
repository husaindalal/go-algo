package general

import (
	"sync"
	"time"
)

//type TimestampCount struct {
//	LastUpdated  int
//	CurrentCount int
//}

type TokenBucket struct {
	RefillIntervalSec int // 10 / sec
	RefillCount       int
	MaxCount          int
	//Lookup map[string]TimestampCount

	LastUpdatedTs int
	CurrentCount  int

	sync.RWMutex
}

func NewTokenBucket(RefillInterval int, RefillCount int) *TokenBucket {
	return &TokenBucket{
		RefillIntervalSec: RefillInterval,
		RefillCount:       RefillCount,
		//Lookup: map[string]TimestampCount{},

		LastUpdatedTs: int(time.Now().Unix()),
		CurrentCount:  RefillCount,
	}
}

func (t *TokenBucket) IsRateLimited(timestamp int) bool {
	t.refill(timestamp)

	if t.CurrentCount == 0 {
		return false
	}

	t.CurrentCount--
	return true
}

func (t *TokenBucket) refill(timestamp int) {
	now := timestamp // int(time.Now().Unix())
	refillCount := (now - t.LastUpdatedTs) / t.RefillIntervalSec * t.RefillCount

	if t.CurrentCount+refillCount > t.RefillCount {
		// overflow
		t.CurrentCount = t.RefillCount
	} else {
		t.CurrentCount = t.CurrentCount + refillCount
		t.LastUpdatedTs = now
	}

}

/*

	tb := general.NewTokenBucket(1, 3)

	fmt.Printf("result %v %v\n", 1, tb.IsRateLimited(1))
	fmt.Printf("result %v %v\n", 2, tb.IsRateLimited(2))
	fmt.Printf("result %v %v\n", 4, tb.IsRateLimited(4))
	fmt.Printf("result %v %v\n", 7, tb.IsRateLimited(7))
	fmt.Printf("result %v %v\n", 10, tb.IsRateLimited(10))
	fmt.Printf("result %v %v\n", 12, tb.IsRateLimited(12))

*/

/*
type TokenBucket struct {
	RefillIntervalSec int // 10 / sec
	RefillCount int
	//Lookup map[string]TimestampCount

	LastUpdatedTs  int
	CurrentCount int

	sync.RWMutex
}

func NewTokenBucket(RefillInterval int, RefillCount int) *TokenBucket {
	return &TokenBucket{
		RefillIntervalSec: RefillInterval,
		RefillCount: RefillCount,
		//Lookup: map[string]TimestampCount{},

		LastUpdatedTs: int(time.Now().Unix()),
		CurrentCount: RefillCount,
	}
}

func (t *TokenBucket) IsRateLimited(key string) bool {
	t.refill()



	return false
}

func (t *TokenBucket) refill() {
	now := int(time.Now().Unix())
	tokensToAdd := t.RefillCount * (now - t.LastUpdatedTs) / t.RefillIntervalSec
	t.CurrentCount = min(t.RefillCount, t.CurrentCount + tokensToAdd)

	t.LastUpdatedTs = now
}
*/
