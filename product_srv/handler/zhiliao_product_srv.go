package handler

import (
	"context"
	"io"
	"product_srv/proto/product"
	"time"

	"go-micro.dev/v4/logger"
)

type Zhiliaoproductsrv struct{}

func (e *Zhiliaoproductsrv) Call(ctx context.Context, req *product.CallRequest, rsp *product.CallResponse) error {
	logger.Infof("Received Zhiliaoproductsrv.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Zhiliaoproductsrv) ClientStream(ctx context.Context, stream product.Zhiliaoproductsrv_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.Infof("Got %v pings total", count)
			return stream.SendMsg(&product.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Zhiliaoproductsrv) ServerStream(ctx context.Context, req *product.ServerStreamRequest, stream product.Zhiliaoproductsrv_ServerStreamStream) error {
	logger.Infof("Received Zhiliaoproductsrv.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Sending %d", i)
		if err := stream.Send(&product.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Zhiliaoproductsrv) BidiStream(ctx context.Context, stream product.Zhiliaoproductsrv_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&product.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
