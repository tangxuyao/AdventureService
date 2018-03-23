package adventure

import (
	"golang.org/x/net/context"
	"proto/adventure"
	"../../MongoData"
)

type AdventureService struct {
	M *mongo.MongoDB
}

// 该方法不需要实现。
func (a *AdventureService) AdventurePing(
	ctx context.Context,
	req *adventrue_api.AdventurePingReq,
	rsp *adventrue_api.AdventurePingRsp) error {
	return nil
}

func (a *AdventureService) Start(ctx context.Context,
	req *adventrue_api.StartReq,
	rsp *adventrue_api.StartRsp) error {



	return nil
}
