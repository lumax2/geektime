package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Trigger struct {
	param string
}

type TriggerRepo interface {
	CreateTrigger(context.Context, *Trigger) error
	UpdateTrigger(context.Context, *Trigger) error
}

type TriggerUsecase struct {
	repo TriggerRepo
	log  *log.Helper
}

func NewTriggerUsecase(repo TriggerRepo, logger log.Logger) *TriggerUsecase {
	return &TriggerUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TriggerUsecase) Create(ctx context.Context, g *Trigger) error {
	return uc.repo.CreateTrigger(ctx, g)
}

func (uc *TriggerUsecase) Update(ctx context.Context, g *Trigger) error {
	return uc.repo.UpdateTrigger(ctx, g)
}
