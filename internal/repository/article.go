package repository

import (
	"context"

	"github.com/Ptt-official-app/go-bbs"
)

func (repo *repository) GetPopularArticles(ctx context.Context) ([]bbs.ArticleRecord, error) {
	// TODO: should implement popular articles there
	return []bbs.ArticleRecord{}, nil
}
