package functionaloption

import "time"

const (
	DefaultTimeout = time.Duration(5) * time.Second
)

type Client struct {
	Timeout time.Duration
}

type Option func(*Client) error

func Timeout(d time.Duration) Option {
	return func(c *Client) error {
		c.Timeout = d
		return nil
	}
}

func NewClient(options ...Option) (*Client, error) {
	c := &Client{
		Timeout: DefaultTimeout,
	}

	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func Example() {
	// New Client with defaults
	_, _ = NewClient()

	// New Client with no Timeout
	_, _ = NewClient(Timeout(time.Duration(0)))
}
