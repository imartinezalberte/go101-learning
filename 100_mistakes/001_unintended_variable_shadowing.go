// Unintended variable shadowing
//
// Variable shadowing is when a variable name declared in a block may be redeclared in an inner block.
package mistakes

import "errors"

type (
	Client struct{ tracing bool }

	CreateClientFn func() (*Client, error)
)

// I'm creating this funcs for testing purposes...
var (
	ErrCreateClientWithTracing = errors.New("create client with tracing")
	ErrCreateDefaultClient     = errors.New("create default client")

	CreateClientWithTracing = func() (*Client, error) {
		return &Client{}, nil
	}
	CreateClientWithTracingWithErr = func() (*Client, error) {
		return nil, ErrCreateClientWithTracing
	}

	CreateDefaultClient = func() (*Client, error) {
		return &Client{}, nil
	}
	CreateDefaultClientWithErr = func() (*Client, error) {
		return nil, ErrCreateDefaultClient
	}

	BuildClient               = BuildClientGen(CreateClientWithTracing, CreateDefaultClient)        // tracing value can be false or true
	BuildClientWithTracingErr = BuildClientGen(CreateClientWithTracingWithErr, CreateDefaultClient) // tracing value must be true
	BuildClientDefaultErr     = BuildClientGen(CreateClientWithTracing, CreateDefaultClientWithErr) // tracing value must be false
)

// Func that contains the error of 
func BuildClientGen(createClientWithTracing CreateClientFn, createDefaultClient CreateClientFn) func(bool) (*Client, error) {
	return func(tracing bool) (*Client, error) {
		var client *Client
		if tracing {
			client, err := createClientWithTracing()
			if err != nil {
				return client, err
			}
		} else {
			client, err := createDefaultClient()
			if err != nil {
				return client, err
			}
		}

		return client, nil
	}
}

func BuildClientGen1(createClientWithTracing CreateClientFn, createDefaultClient CreateClientFn) func(bool) (*Client, error) {
	return func(tracing bool) (client *Client, err error) {
		if tracing {
			client, err = createClientWithTracing()
			if err != nil {
				return nil, err
			}
		} else {
			client, err = createDefaultClient()
			if err != nil {
				return nil, err
			}
		}

		return client, nil
	}
}


func BuildClientGen2(createClientWithTracing CreateClientFn, createDefaultClient CreateClientFn) func(bool) (*Client, error) {
	return func(tracing bool) (*Client, error) {
		var (
			client *Client
			err error
		)

		if tracing {
			client, err = createClientWithTracing()
			if err != nil {
				return nil, err
			}
		} else {
			client, err = createDefaultClient()
			if err != nil {
				return nil, err
			}
		}

		return client, nil
	}
}

func BuildClientGen3(createClientWithTracing CreateClientFn, createDefaultClient CreateClientFn) func(bool) (*Client, error) {
	return func(tracing bool) (*Client, error) {
		var client *Client

		if tracing {
			cl, err := createClientWithTracing()
			if err != nil {
				return nil, err
			}
			client = cl
		} else {
			cl, err := createDefaultClient()
			if err != nil {
				return nil, err
			}
			client = cl
		}

		return client, nil
	}
}

func BuildClientGen4(createClientWithTracing CreateClientFn, createDefaultClient CreateClientFn) func(bool) (*Client, error) {
	return func(tracing bool) (*Client, error) {
		var (
			client *Client
			err error
		)

		if tracing {
			client, err = createClientWithTracing()
		} else {
			client, err = createDefaultClient()
		}

		// We can return client and err and that's it...
		if err != nil {
			return nil, err
		}

		return client, nil
	}
}
