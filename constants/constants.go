package constants

import "time"

const (
	HeaderAuthorization = "Authorization"
	HeaderXSecretToken  = "X-Secret-Token"
)

const (
	RequestTimeout = 30 * time.Second
)

const (
	RateLimit             = 60
	RateLimitWindowLength = time.Minute
)
