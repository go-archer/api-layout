package router

import (
	"time"

	"github.com/go-zelus/zelus/core/limiter"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.BucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})
