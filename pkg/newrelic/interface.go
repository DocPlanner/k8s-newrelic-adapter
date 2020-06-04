package newrelic

// Client represents a client for Amazon CloudWatch.
type Client interface {
	// Query sends a list of queries to Cloudwatch for metric results.
	Query() (float64, error)
}
