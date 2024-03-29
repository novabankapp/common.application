package base

import (
	"context"
	domainBase "github.com/novabankapp/common.data/domain/base"
	"github.com/novabankapp/common.data/repositories/base"
)

type NoSqlService[E domainBase.NoSqlEntity] struct {
	repo base.NoSqlRepository[E]
}

func NewDocumentDatabaseService[E domainBase.NoSqlEntity](repo base.NoSqlRepository[E]) NoSqlService[E] {

	return NoSqlService[E]{
		repo: repo,
	}
}
func (s *NoSqlService[E]) Create(ctx context.Context, entity E) (bool, error) {
	return s.repo.Create(ctx, entity)

}
func (s *NoSqlService[E]) GetById(ctx context.Context, id string) (*E, error) {
	return s.repo.GetById(ctx, id)
}

func (s *NoSqlService[E]) Update(ctx context.Context, entity E, id string) (bool, error) {
	return s.repo.Update(ctx, entity, id)
}
func (s *NoSqlService[E]) Delete(ctx context.Context, id string) (bool, error) {
	return s.repo.Delete(ctx, id)
}
func (s *NoSqlService[E]) Get(ctx context.Context,
	page []byte, pageSize int, queries []map[string]string, orderBy string) (*[]E, []byte, error) {
	return s.repo.Get(ctx, page, pageSize, queries, orderBy)

}
func (s *NoSqlService[E]) GetByCondition(ctx context.Context,
	queries []map[string]string) (*E, error) {
	return s.repo.GetByCondition(ctx, queries)

}
