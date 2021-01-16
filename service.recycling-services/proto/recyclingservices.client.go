// DO NOT EDIT: Client was autogenerated from 'github.com/edstell/lambda/service.recycling-services/proto/recyclingservices.proto'
package recyclingservicesproto

import (
	"context"

	"github.com/edstell/lambda/libraries/rpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type Client struct {
	readproperty   rpc.Invoker
	syncproperty   rpc.Invoker
	notifyproperty   rpc.Invoker
}

func NewClient(i rpc.LambdaInvoker, arn string) *Client {
	return &Client{
		readproperty:   rpc.Client(i, arn, "ReadProperty"),
		syncproperty:   rpc.Client(i, arn, "SyncProperty"),
		notifyproperty:   rpc.Client(i, arn, "NotifyProperty"),
	}
}

func (c *Client) ReadProperty(ctx context.Context, req *ReadPropertyRequest) (*ReadPropertyResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.readproperty.Invoke(ctx, payload)
	if err != nil {
		return nil, err
	}

	out := &ReadPropertyResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) SyncProperty(ctx context.Context, req *SyncPropertyRequest) (*SyncPropertyResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.syncproperty.Invoke(ctx, payload)
	if err != nil {
		return nil, err
	}

	out := &SyncPropertyResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) NotifyProperty(ctx context.Context, req *NotifyPropertyRequest) (*NotifyPropertyResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.notifyproperty.Invoke(ctx, payload)
	if err != nil {
		return nil, err
	}

	out := &NotifyPropertyResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}

	return out, nil
}