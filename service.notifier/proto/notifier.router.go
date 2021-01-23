// DO NOT EDIT: Router was autogenerated from 'notifier.proto'
package notifierproto

import (
	"context"

	"github.com/edstell/lambda/libraries/rpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

type validator interface {
	Validate() error
}

type Handler interface {
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
}

type Router struct {
	*rpc.Router
}

func NewRouter(handler Handler) *Router {
	router := rpc.NewRouter()
	router.Route("Notify", notify(handler.Notify))
	return &Router{
		Router: router,
	}
}

func notify(handler func(context.Context, *NotifyRequest) (*NotifyResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &NotifyRequest{}
		if err := protojson.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}
		var b interface{} = body
		if v, ok := b.(validator); ok {
			if err := v.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		rsp, err := handler(ctx, body)
		if err != nil {
			return nil, err
		}
		bytes, err := protojson.Marshal(rsp)
		if err != nil {
			return nil, err
		}
		return &rpc.Response{
			Body: bytes,
		}, nil
	}
}
