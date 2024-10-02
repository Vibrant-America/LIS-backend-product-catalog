package middleware

import (
	"context"
	"encoding/json"
	"productCatalog/global"

	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
	"google.golang.org/grpc/peer"
)

var logBody = func(ctx context.Context, desc string, body interface{}) {
	reqJSON, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	//log.Info(desc, zap.String("body", string(reqJSON)))
	log.Info(desc, string(reqJSON))
}

var LogHandler = func(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		p, _ := peer.FromContext(ctx)
		requestIp := p.Addr.String()
		//log.Info(requestIp)
		logBody(ctx, "request(in)", global.RpcRequest{
			EndPoint: req.Endpoint(),
			PeerIp:   requestIp,
			Body:     req.Body(),
		})
		err := fn(ctx, req, rsp)
		if err != nil {
			log.Error()
			return err
		}

		logBody(ctx, "response(out)", rsp)
		return err
	}
}
