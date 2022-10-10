package mapper

import (
	"encoding/json"
	"errors"
	"time"
	"try/connect-redis-golang/domain/entity"
)

type AttributeArticle struct {
	Id            int       `json:"id"`
	CodeArticle   string    `json:"code_article"`
	TitleOriginal string    `json:"title_original"`
	TextOriginal  string    `json:"text_original"`
	Date          time.Time `json:"date"`
	Banner        string    `json:"banner"`
	Author        string    `json:"author"`
	Thumbs        string    `json:"thumbs"`
	IsHighLight   bool      `json:"is_highlight"`
}

func MapGetRedisField() string {
	return "data_article"
}

func MapFromJsonStringToDomainArticle(data string) (*entity.Article, error) {
	attributeArticle := new(AttributeArticle)
	err := json.Unmarshal([]byte(data), attributeArticle)
	if err != nil {
		return nil, err
	}

	article, err := entity.NewCreateArticleSingle(attributeArticle.CodeArticle, attributeArticle.TitleOriginal, attributeArticle.TextOriginal, attributeArticle.Date, attributeArticle.Banner, attributeArticle.Author, attributeArticle.Thumbs, attributeArticle.IsHighLight)
	if err != nil {
		return nil, errors.New("gagal mapping data redis")
	}

	return article, nil
}

func MapGetKeyValueRedis(data *entity.Article) string {
	return data.GetCodeArtikel()
}

func MapSetBukuToString(data *entity.Article) string {
	attrBuku := &AttributeArticle{
		CodeArticle:   data.GetCodeArtikel(),
		TitleOriginal: data.GetTitleArtikel(),
		TextOriginal:  data.GetTextArtikel(),
		Date:          data.GetTanggalArtikel(),
		Banner:        data.GetBannerArtikel(),
		Author:        data.GetAuthorArtikel(),
		Thumbs:        data.GetThumbsArtikel(),
		IsHighLight:   data.GetIsHighLightArtikel(),
	}

	attrJson, _ := json.Marshal(attrBuku)

	return string(attrJson)
}
