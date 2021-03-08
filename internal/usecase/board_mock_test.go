package usecase

import (
	"context"
	"time"

	"github.com/Ptt-official-app/Ptt-backend/internal/repository"
	"github.com/Ptt-official-app/go-bbs"
)

func (repo *MockRepository) GetBoards(ctx context.Context) []bbs.BoardRecord {
	return []bbs.BoardRecord{}
}

func (repo *MockRepository) GetBoardArticle(ctx context.Context, boardID, filename string) ([]byte, error) {
	return []byte{}, nil
}
func (repo *MockRepository) GetBoardArticleRecords(ctx context.Context, boardID string) ([]bbs.ArticleRecord, error) {
	articleRecords := []MockArticleRecord{
		{
			filename:       "",
			modified:       time.Time{},
			recommendCount: 10,
			owner:          "SYSOP",
			date:           "",
			title:          "[討論] 偶爾要發個廢文",
			money:          0,
		},
		{
			filename:       "",
			modified:       time.Time{},
			recommendCount: -20,
			owner:          "XDXD",
			date:           "",
			title:          "[問題] UNICODE",
			money:          0,
		},
	}
	result := make([]bbs.ArticleRecord, len(articleRecords))
	for i, v := range articleRecords {
		result[i] = v
	}
	return result, nil
}

func (repo *MockRepository) GetPopularArticles(ctx context.Context) ([]repository.PopularArticleRecord, error) {
	articleRecords := []MockPopularArticle{
		{
			filename:       "",
			modified:       time.Time{},
			recommendCount: 10,
			owner:          "SYSOP",
			date:           "",
			title:          "Popular Article 1",
			money:          0,
			boardID:        "Gossiping",
		},
		{
			filename:       "",
			modified:       time.Time{},
			recommendCount: -20,
			owner:          "XDXD",
			date:           "",
			title:          "Popular Article 2",
			money:          0,
			boardID:        "Gossiping",
		},
		{
			filename:       "",
			modified:       time.Time{},
			recommendCount: 0,
			owner:          "TEST",
			date:           "",
			title:          "Popular Article 3",
			money:          0,
			boardID:        "Joke",
		},
	}
	result := make([]repository.PopularArticleRecord, len(articleRecords))
	for i, v := range articleRecords {
		result[i] = v
	}
	return result, nil
}

func (repo *MockRepository) GetBoardTreasureRecords(ctx context.Context, boardID string, treasureIDs []string) ([]bbs.ArticleRecord, error) {
	return []bbs.ArticleRecord{}, nil
}

type MockArticleRecord struct {
	filename       string
	modified       time.Time
	recommendCount int
	owner          string
	date           string
	title          string
	money          int
}

func (m MockArticleRecord) Filename() string    { return m.filename }
func (m MockArticleRecord) Modified() time.Time { return m.modified }
func (m MockArticleRecord) Recommend() int      { return m.recommendCount }
func (m MockArticleRecord) Owner() string       { return m.owner }
func (m MockArticleRecord) Date() string        { return m.date }
func (m MockArticleRecord) Title() string       { return m.title }
func (m MockArticleRecord) Money() int          { return m.money }

type MockPopularArticle struct {
	filename       string
	modified       time.Time
	recommendCount int
	owner          string
	date           string
	title          string
	money          int
	boardID        string
}

func (m MockPopularArticle) Filename() string    { return "" }
func (m MockPopularArticle) Modified() time.Time { return time.Unix(0, 0) }
func (m MockPopularArticle) Recommend() int      { return 0 }
func (m MockPopularArticle) Date() string        { return "" }
func (m MockPopularArticle) Title() string       { return m.title }
func (m MockPopularArticle) Money() int          { return 0 }
func (m MockPopularArticle) Owner() string       { return "" }
func (m MockPopularArticle) BoardID() string     { return m.boardID }
