package mysql_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"try/connect-redis-golang/domain/entity"
	"try/connect-redis-golang/domain/repository"
)

// interaksi dengan DB
type ArticleRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

// build structnya, yang mengacu ke connection dan kontrak interface di repository
func NewArticleRepositoryMysqlInteractor(connectionDatabse *sql.DB) repository.ArticleRepository {
	return &ArticleRepositoryMysqlInteractor{dbConn: connectionDatabse}
}

// implementasi dari interface kontrak dalam bentuk method receiver
func (repo *ArticleRepositoryMysqlInteractor) Store(ctx context.Context, dataArticle *entity.Article) error {
	var (
		errMysql            error
		errMysqlTranslation error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// query insert to table article
	insertQuery := "INSERT INTO article(code_article, title_original, text_original, date, " +
		"banner, author, thumbs, is_highlight) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"

	_, errMysql = repo.dbConn.Exec(insertQuery, dataArticle.GetCodeArtikel(), dataArticle.GetTitleArtikel(), dataArticle.GetTextArtikel(),
		dataArticle.GetTanggalArtikel(), dataArticle.GetBannerArtikel(), dataArticle.GetAuthorArtikel(), dataArticle.GetThumbsArtikel(), dataArticle.GetIsHighLightArtikel())

	if errMysql != nil {
		return errMysql
	}

	translationData := dataArticle.GetTranslationAll()
	// query insert to table translation
	insertQueryTranslation := "INSERT INTO translation(code_article, code_language, title, text) VALUES(?, ?, ?, ?)"

	for _, td := range translationData {
		// save to DB
		_, errMysqlTranslation = repo.dbConn.Exec(insertQueryTranslation, dataArticle.GetCodeArtikel(), td.GetCodeLanguageTranslation(), td.GetTitleTranslation(), td.GetTextTranslation())
	}

	if errMysqlTranslation != nil {
		return errMysqlTranslation
	}

	return nil
}

func (repo *ArticleRepositoryMysqlInteractor) GetAllData(ctx context.Context) ([]*entity.Article, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT * FROM article"
	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery)
	if errMysql != nil {
		return nil, errMysql
	}

	dataArticleCollection := make([]*entity.Article, 0)
	for rows.Next() {
		var (
			id             int
			code_article   string
			title_original string
			text_original  string
			date           time.Time
			banner         string
			author         string
			thumbs         string
			is_highlight   bool
		)

		err := rows.Scan(&id, &code_article, &title_original, &text_original, &date, &banner, &author, &thumbs, &is_highlight)
		if err != nil {
			return nil, err
		}
		dataArticle := entity.FetchDataArticleFromDB(entity.DTOArticleFromDatabase{
			Id:            id,
			CodeArticle:   code_article,
			TitleOriginal: title_original,
			TextOriginal:  text_original,
			Date:          date,
			Banner:        banner,
			Author:        author,
			Thumbs:        thumbs,
			IsHighLight:   is_highlight,
		})

		dataArticleCollection = append(dataArticleCollection, dataArticle)
	}
	defer rows.Close()

	return dataArticleCollection, nil
}

func (repo *ArticleRepositoryMysqlInteractor) GetDataByCode(ctx context.Context, codeArticle string) (*entity.Article, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT * FROM article where code_article = ?"
	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery, codeArticle)
	if errMysql != nil {
		return nil, errMysql
	}

	for rows.Next() {
		var (
			id             int
			code_article   string
			title_original string
			text_original  string
			date           time.Time
			banner         string
			author         string
			thumbs         string
			is_highlight   bool
		)

		err := rows.Scan(&id, &code_article, &title_original, &text_original, &date, &banner, &author, &thumbs, &is_highlight)
		if err != nil {
			return nil, err
		}

		dataArticle, err := entity.NewCreateArticleSingle(code_article, title_original, text_original, date, banner, author, thumbs, is_highlight)
		if err != nil {
			return nil, errors.New("gagal mapping data redis")
		}

		// ini gmna supaya tidak loop
		return dataArticle, nil
	}
	return nil, nil

}
