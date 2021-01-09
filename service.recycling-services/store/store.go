package store

import (
	"context"

	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
)

type Store interface {
	ReadProperty(context.Context, string) (*recyclingservices.Property, error)
	WriteProperty(context.Context, string, []recyclingservices.Service) (*recyclingservices.Property, error)
}