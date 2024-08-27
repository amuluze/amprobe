// Package service
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package service

import (
	"context"

	"amprobe/service/audit/repository"
	"amprobe/service/schema"

	"github.com/google/wire"
)

var AuditServiceSet = wire.NewSet(NewAuditService, wire.Bind(new(IAuditService), new(*AuditService)))

type IAuditService interface {
	AuditQuery(ctx context.Context, args *schema.AuditQueryArgs) (*schema.AuditQueryReply, error)
}

type AuditService struct {
	AuditRepo repository.IAuditRepo
}

func NewAuditService(repo repository.IAuditRepo) *AuditService {
	return &AuditService{AuditRepo: repo}
}

func (a AuditService) AuditQuery(ctx context.Context, args *schema.AuditQueryArgs) (*schema.AuditQueryReply, error) {
	audits, err := a.AuditRepo.AuditQuery(ctx, args)
	if err != nil {
		return &schema.AuditQueryReply{}, err
	}

	var list []schema.Audit
	for _, audit := range audits {
		list = append(list, schema.Audit{
			ID:       audit.ID,
			Username: audit.Username,
			Operate:  audit.Operate,
			Created:  audit.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	total, err := a.AuditRepo.AuditCount(ctx)
	if err != nil {
		return &schema.AuditQueryReply{}, err
	}
	return &schema.AuditQueryReply{Data: list, Total: total, Page: args.Page, Size: args.Size}, nil
}
