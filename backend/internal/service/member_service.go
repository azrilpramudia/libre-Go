package service

import (
	"context"

	"github.com/azrilpramudia/libre-go/internal/domain/member"
	"github.com/google/uuid"
)

type MemberService struct {
	repo member.Repository
}

func NewMemberService(repo member.Repository) *MemberService {
	return &MemberService{repo: repo}
}

func (s *MemberService) CreateMember(ctx context.Context, m *member.Member) (*member.Member, error) {
	return s.repo.Create(ctx, m)
}

func (s *MemberService) GetMember(ctx context.Context, id uuid.UUID) (*member.Member, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MemberService) ListMembers(ctx context.Context, limit, offset int32) ([]*member.Member, error) {
	return s.repo.List(ctx, limit, offset)
}
