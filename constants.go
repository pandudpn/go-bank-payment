package bankpayment

// Environment is environment server
type Environment string

const (
	Production  Environment = "production"
	Development Environment = "development"
	Testing     Environment = "testing"
)
