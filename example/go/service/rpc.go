package service

import (
	"context"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	//pb "github.com/nervatura/nervatura/v6/pkg/service/grpc/proto"

	pb "github.com/nervatura/nervatura/v6/protos/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type RpcClient struct {
	Config    cu.SM
	conn      *grpc.ClientConn
	permanent bool
}

func (rpc *RpcClient) Connect(keeping bool) (err error) {
	if rpc.conn == nil {
		rpc.conn, err = grpc.NewClient("localhost:"+rpc.Config["NT_GRPC_PORT"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if keeping {
		rpc.permanent = true
	}
	return err
}

func (rpc *RpcClient) Close(all bool) {
	if rpc.conn != nil && (!rpc.permanent || all) {
		_ = rpc.conn.Close()
		rpc.conn = nil
		rpc.permanent = false
	}
}

func (rpc *RpcClient) PermanentConnect() {

}

func (rpc *RpcClient) Metadata(token string) metadata.MD {
	md := metadata.New(map[string]string{})
	if token != "" {
		md.Set("Authorization", "Bearer "+token)
	} else {
		md.Set("x-api-key", rpc.Config["NT_API_KEY"])
	}
	return md
}

func (rpc *RpcClient) Database(options cu.IM) (result any, err error) {
	if err := rpc.Connect(false); err != nil {
		return nil, err
	}
	md := rpc.Metadata("")
	metaCtx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(metaCtx, time.Second*30)
	defer cancel()

	response, err := pb.NewAPIClient(rpc.conn).Database(ctx, &pb.RequestDatabase{
		Alias: cu.ToString(options["alias"], ""),
		Demo:  cu.ToBoolean(options["demo"], false),
	})
	rpc.Close(false)
	if err != nil {
		return nil, err
	}
	var results any
	err = cu.ConvertFromByte(response.Data, &results)
	return results, err
}

func (rpc *RpcClient) CustomerUpdate(token string, data *pb.Customer) (result *pb.Customer, err error) {
	if err := rpc.Connect(false); err != nil {
		return nil, err
	}
	md := rpc.Metadata(token)
	metaCtx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(metaCtx, time.Second*30)
	defer cancel()

	result, err = pb.NewAPIClient(rpc.conn).CustomerUpdate(ctx, data)
	rpc.Close(false)
	return result, err
}

func (rpc *RpcClient) CustomerGet(token string, options *pb.RequestGet) (result *pb.Customer, err error) {
	if err := rpc.Connect(false); err != nil {
		return nil, err
	}
	md := rpc.Metadata(token)
	metaCtx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(metaCtx, time.Second*30)
	defer cancel()

	result, err = pb.NewAPIClient(rpc.conn).CustomerGet(ctx, options)
	rpc.Close(false)
	return result, err
}

func (rpc *RpcClient) CustomerQuery(token string, options *pb.RequestQuery) (result *pb.Customers, err error) {
	if err := rpc.Connect(false); err != nil {
		return nil, err
	}
	md := rpc.Metadata(token)
	metaCtx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(metaCtx, time.Second*30)
	defer cancel()

	result, err = pb.NewAPIClient(rpc.conn).CustomerQuery(ctx, options)
	rpc.Close(false)
	return result, err
}

func (rpc *RpcClient) Delete(token string, options *pb.RequestDelete) (result *pb.ResponseStatus, err error) {
	if err := rpc.Connect(false); err != nil {
		return nil, err
	}
	md := rpc.Metadata(token)
	metaCtx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(metaCtx, time.Second*30)
	defer cancel()

	result, err = pb.NewAPIClient(rpc.conn).Delete(ctx, options)
	rpc.Close(false)
	return result, err
}

func (rpc *RpcClient) Function(token string, options *pb.RequestFunction) (result any, err error) {
	if err := rpc.Connect(false); err != nil {
		return nil, err
	}
	md := rpc.Metadata(token)
	metaCtx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(metaCtx, time.Second*30)
	defer cancel()

	var response *pb.JsonBytes
	if response, err = pb.NewAPIClient(rpc.conn).Function(ctx, options); err == nil {
		err = cu.ConvertFromByte(response.Data, &result)
	}
	rpc.Close(false)
	return result, err
}

func (rpc *RpcClient) View(token string, options *pb.RequestView) (result any, err error) {
	if err := rpc.Connect(false); err != nil {
		return nil, err
	}
	md := rpc.Metadata(token)
	metaCtx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(metaCtx, time.Second*30)
	defer cancel()

	var response *pb.JsonBytes
	if response, err = pb.NewAPIClient(rpc.conn).View(ctx, options); err == nil {
		err = cu.ConvertFromByte(response.Data, &result)
	}
	rpc.Close(false)
	return result, err
}
