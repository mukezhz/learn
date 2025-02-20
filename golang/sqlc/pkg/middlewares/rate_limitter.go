package middlewares

import (
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

const RateLimit = "RateLimit"
// Global store
// using in-memory store with goroutine which clears expired keys.
var store = memory.NewStore()

type RateLimitOption struct {
	period time.Duration
	limit  int64
}

const (
	RateLimitPeriod   = 15 * time.Minute
	RateLimitRequests = int64(200)
)

type Option func(*RateLimitOption)

type RateLimitMiddleware struct {
	logger framework.Logger
	option RateLimitOption
}

func NewRateLimitMiddleware(logger framework.Logger) RateLimitMiddleware {
	return RateLimitMiddleware{
		logger: logger,
		option: RateLimitOption{
			period: RateLimitPeriod,
			limit:  RateLimitRequests,
		},
	}
}

func (lm RateLimitMiddleware) Handle(options ...Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP() // Gets cient IP Address

		lm.logger.Info("Setting up rate limit middleware")

		// Setting up rate limit
		// Limit -> # of API Calls
		// Period -> in a given time frame
		// setting default values
		opt := RateLimitOption{
			period: lm.option.period,
			limit:  lm.option.limit,
		}

		for _, o := range options {
			o(&opt)
		}

		rate := limiter.Rate{
			Limit:  opt.limit,
			Period: opt.period,
		}

		// Limiter instance
		instance := limiter.New(store, rate)

		// Returns the rate limit details for given identifier.
		// FullPath is appended with IP address. `/api/users&&127.0.0.1` as key
		context, err := instance.Get(c, c.FullPath()+"&&"+key)

		if err != nil {
			lm.logger.Panic(err.Error())
		}

		c.Set(RateLimit, instance)

		// Setting custom headers
		c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

		// Limit exceeded
		if context.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate Limit Exceed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func WithOptions(period time.Duration, limit int64) Option {
	return func(o *RateLimitOption) {
		o.period = period
		o.limit = limit
	}
}
