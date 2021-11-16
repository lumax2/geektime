package data

import (
	"HomeworkFour/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type triggerRepo struct {
	data *Data
	log  *log.Helper
}

//NewTriggerRepo
func NewTriggerRepo(data *Data, logger log.Logger) biz.TriggerRepo {
	return &triggerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *triggerRepo) CreateTrigger(ctx context.Context, g *biz.Trigger) error {
	return nil
}

func (r *triggerRepo) UpdateTrigger(ctx context.Context, g *biz.Trigger) error {
	return nil
}
