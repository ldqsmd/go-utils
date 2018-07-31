package microwrapper

import (

	"context"
	"github.com/micro/go-micro/client"
	"github.com/xiaomeng79/go-log"
	"github.com/micro/go-micro/server"
)

// log wrapper logs every time a request is made
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.InfoO(ctx,req.Request(),"客户端记录 服务:"+req.Service()+" 方法:"+req.Method()+"请求信息")
	return l.Client.Call(ctx, req, rsp)
}

// Implements client.Wrapper as logWrapper
func LogClientWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

// logWrapper is a handler wrapper
func LogServerWrap(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := fn(ctx, req, rsp)
		log.InfoO(ctx,req.Request(),"服务端记录:"+req.Service()+" 方法:"+req.Method()+"请求信息")
		log.InfoO(ctx,rsp,"服务端记录:"+req.Service()+" 方法:"+req.Method()+"响应信息")
		if err != nil {
			log.InfoO(ctx,err,"服务端错误记录:"+req.Service()+" 方法:"+req.Method())
		}
		return err
	}
}
