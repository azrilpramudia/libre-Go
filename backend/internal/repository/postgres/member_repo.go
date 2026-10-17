package postgres

import (
	"context"

	"github.com/azrilpramudia/libre-go/internal/domain/member"
	"github.com/azrilpramudia/libre-go/internal/repository/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type memberRepo struct {
	q *sqlc.Queries
}

func NewMemberRepository(q *sqlc.Queries) member.Repository {
	return &memberRepo{q: q}
}

func (r *memberRepo) Create(ctx context.Context, m *member.Member) (*member.Member, error) {
	row, err := r.q.CreateMember(ctx, sqlc.CreateMemberParams{
		FullName: m.FullName,
		Email:    m.Email,
		Phone:    pgtype.Text{String: m.Phone, Valid: m.Phone != ""},
	})
	if err != nil {
		return nil, err
	}
	return toDomainMember(row), nil
}

func (r *memberRepo) GetByID(ctx context.Context, id uuid.UUID) (*member.Member, error) {
	row, err := r.q.GetMemberByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return nil, err
	}
	return toDomainMember(row), nil
}

func (r *memberRepo) List(ctx context.Context, limit, offset int32) ([]*member.Member, error) {
	rows, err := r.q.ListMembers(ctx, sqlc.ListMembersParams{Limit: limit, Offset: offset})
	if err != nil {
		return nil, err
	}
	result := make([]*member.Member, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainMember(row))
	}
	return result, nil
}

func toDomainMember(row sqlc.Member) *member.Member {
	return &member.Member{
		ID:        row.ID.Bytes,
		FullName:  row.FullName,
		Email:     row.Email,
		Phone:     row.Phone.String,
		CreatedAt: row.CreatedAt.Time,
	}
}
