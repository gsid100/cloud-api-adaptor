// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: podvminfo.proto
package __

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
)

type PodVMInfoService interface {
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

func RegisterPodVMInfoService(srv *ttrpc.Server, svc PodVMInfoService) {
	srv.RegisterService("podvminfo.PodVMInfo", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"GetInfo": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req GetInfoRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.GetInfo(ctx, &req)
			},
		},
	})
}

type podvminfoClient struct {
	client *ttrpc.Client
}

func NewPodVMInfoClient(client *ttrpc.Client) PodVMInfoService {
	return &podvminfoClient{
		client: client,
	}
}

func (c *podvminfoClient) GetInfo(ctx context.Context, req *GetInfoRequest) (*GetInfoResponse, error) {
	var resp GetInfoResponse
	if err := c.client.Call(ctx, "podvminfo.PodVMInfo", "GetInfo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}