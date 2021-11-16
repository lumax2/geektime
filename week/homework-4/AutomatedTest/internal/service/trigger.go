package service

import (
	v1 "HomeworkFour/api/helloworld/v1"
	"HomeworkFour/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "HomeworkFour/api/trigger/v1"
)

type TriggerService struct {
	pb.UnimplementedTriggerServer

	uc *biz.TriggerUsecase
	log *log.Helper
}

func NewTriggerService(uc *biz.TriggerUsecase,logger log.Logger) *TriggerService {
	return &TriggerService{uc: uc,log:log.NewHelper(logger)}
}

func (s *TriggerService) OnTrigger(ctx context.Context, req *pb.OnTriggerRequest) (*pb.OnTriggerReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", req.GetParam())

	if req.GetParam() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", req.GetParam())
	}
	return &pb.OnTriggerReply{Message: "StartTrigger " + req.GetParam()}, nil
}
