package http

import (
	"context"

	"github.com/Ptt-official-app/go-bbs"
)

func (usecase *MockUsecase) GetPopularArticles(ctx context.Context) ([]bbs.ArticleRecord, error) {
	return []bbs.ArticleRecord{}, nil
}
