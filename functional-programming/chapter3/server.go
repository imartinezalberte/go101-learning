package chapter3

import (
	"errors"
	"strings"
)

const (
	TCP TransportType = iota
	UDP

	UnknownTransportType TransportType = -1
)

// Example
var _ = NewServer(MaxConnectionsOption(10), TransportTypeOption(TCP), NameOption("The coolest server ever created"))

type (
	ServerOptionsFn func(options) options

	TransportType int

	ServerConfig struct {
		isAlive bool
		options
	}

	options struct {
		transportType  TransportType
		name           string
		maxConnections uint
	}
)

func (t *TransportType) Set(input int) error {
	if !t.validate(input) {
		return errors.New("transport type invalid")
	}
	return nil
}

func (t *TransportType) validate(input int) bool {
	switch input {
	case 0:
		*t = TCP
	case 1:
		*t = UDP
	default:
		*t = UnknownTransportType
		return false
	}
	return true
}

func MaxConnectionsOption(input uint) ServerOptionsFn {
	return func(o options) options {
		if input > 0 {
			o.maxConnections = input
		}
		return o
	}
}

func TransportTypeOption(input TransportType) ServerOptionsFn {
	return func(o options) options {
		if err := o.transportType.Set(int(input)); err != nil {
			o.transportType.Set(int(TCP))
		}
		return o
	}
}

func NameOption(input string) ServerOptionsFn {
	return func(o options) options {
		if strings.TrimSpace(input) != "" {
			o.name = input
		}
		return o
	}
}

func NewServer(option ...ServerOptionsFn) ServerConfig {
	o := options{}

	for _, op := range option {
		o = op(o)
	}

	return ServerConfig{
		isAlive: true,
		options: o,
	}
}
