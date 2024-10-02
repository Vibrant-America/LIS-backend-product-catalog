package global

type RpcRequest struct {
	EndPoint string
	PeerIp   string
	Body     interface{}
}
