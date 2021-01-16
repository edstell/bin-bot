// DO NOT EDIT: Router was autogenerated from 'twilio.proto'
package twilioproto

import (
	"context"

	"github.com/edstell/lambda/libraries/rpc"
	"github.com/edstell/lambda/libraries/validation"
	"google.golang.org/protobuf/encoding/protojson"
)

type Handler interface {
	SendSMS(context.Context, *SendSMSRequest) (*SendSMSResponse, error)
}

type Router struct {
	*rpc.Router
}

func NewRouter(handler Handler) *Router {
	router := rpc.NewRouter()
	router.Route("SendSMS", sendsms(handler.SendSMS))
	return &Router{
		Router: router,
	}
}

func sendsms(handler func(context.Context, *SendSMSRequest) (*SendSMSResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &SendSMSRequest{}
		if err := protojson.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}

		if err := validation.Validate(body); err != nil {
			return nil, err
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