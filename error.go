package potofgreed

// SentinelError is the error type to be used for creating sentinel errors
type SentinelError string

func (err SentinelError) Error() string {
	return string(err)
}
