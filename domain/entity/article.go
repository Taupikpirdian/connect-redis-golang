package entity

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Article struct {
	codeArticle   string // generate code for relation
	titleOriginal string
	textOriginal  string
	date          time.Time // generate from time now
	banner        string
	author        string
	thumbs        string // for image with small quality
	isHighlight   bool   // for flagging
	translation   []*Translation
}

type Translation struct {
	codeLanguage string
	title        string
	text         string
}

type DTONewCreateArticle struct {
	TitleOriginal string
	TextOriginal  string
	Banner        string
	Author        string
	Thumbs        string
	IsHighLight   bool
	Translation   []DTOTranslation
}

type DTOTranslation struct {
	CodeLanguage string
	Title        string
	Text         string
}

type DTOArticleFromDatabase struct {
	Id            int
	CodeArticle   string
	TitleOriginal string
	TextOriginal  string
	Date          time.Time
	Banner        string
	Author        string
	Thumbs        string
	IsHighLight   bool
}

func NewAddTranslation(translation DTOTranslation) *Translation {
	return &Translation{
		codeLanguage: translation.CodeLanguage,
		title:        translation.Title,
		text:         translation.Text,
	}
}

func (article *Article) AddTranslation(dataTranslations []*Translation) *Article {
	article.translation = dataTranslations
	return article
}

// func create data
func NewCreateArticle(dataCreate DTONewCreateArticle) (*Article, error) {
	dataArticle := &Article{
		titleOriginal: dataCreate.TitleOriginal,
		textOriginal:  dataCreate.TextOriginal,
		banner:        dataCreate.Banner,
		author:        dataCreate.Author,
		thumbs:        dataCreate.Thumbs,
		isHighlight:   dataCreate.IsHighLight,
	}

	// create slice
	translations := make([]*Translation, 0)
	// loop and append to DTO
	for _, tr := range dataCreate.Translation {
		translation := NewAddTranslation(DTOTranslation{
			CodeLanguage: tr.CodeLanguage,
			Title:        tr.Title,
			Text:         tr.Text,
		})

		translations = append(translations, translation)
	}
	// add to struct utama
	dataArticle.AddTranslation(translations)

	err := dataArticle.validate()
	if err != nil {
		return nil, err
	}

	dataArticle, errGenerate := dataArticle.generateCode()
	if errGenerate != nil {
		return nil, errors.New("GENERATE KODE ARTIKEL ERROR")
	}

	dataArticle, errGenerateTime := dataArticle.generateTime()
	if errGenerateTime != nil {
		return nil, errors.New("GENERATE TANGGAL ARTIKEL ERROR")
	}

	return dataArticle, nil
}

// func get data
func FetchDataArticleFromDB(dataDTO DTOArticleFromDatabase) *Article {
	dataArticle := &Article{
		codeArticle:   dataDTO.CodeArticle,
		titleOriginal: dataDTO.TitleOriginal,
		textOriginal:  dataDTO.TextOriginal,
		date:          dataDTO.Date,
		banner:        dataDTO.Banner,
		author:        dataDTO.Author,
		thumbs:        dataDTO.Thumbs,
		isHighlight:   dataDTO.IsHighLight,
	}

	return dataArticle
}

// validasi
func (data *Article) validate() error {
	// validasi setelah masuk struct utama
	if data.titleOriginal == "" {
		return errors.New("JUDUL ARTIKEL TIDAK BOLEH KOSONG")
	}

	if data.textOriginal == "" {
		return errors.New("ISI ARTIKEL TIDAK BOLEH KOSONG")
	}

	if data.banner == "" {
		return errors.New("FILE BANNER ARTIKEL TIDAK BOLEH KOSONG")
	}

	if data.author == "" {
		return errors.New("PENULIS ARTIKEL TIDAK BOLEH KOSONG")
	}

	return nil
}

// generate kode article
func (data *Article) generateCode() (*Article, error) {
	// hanya bisa digenerate 1x
	if data.codeArticle != "" {
		return nil, errors.New("KODE ARTIKEL SUDAH ADA")
	}

	rand.Seed(time.Now().UnixNano())
	min := 10000000
	max := 30000000
	valueString := strconv.Itoa(rand.Intn(max-min+1) + min)

	data.codeArticle = "XXXX-" + valueString
	return data, nil
}

// GET TIME NOW
func (data *Article) generateTime() (*Article, error) {
	// hanya bisa digenerate 1x
	if data.date.IsZero() {
		data.date = time.Now()
		return data, nil

	} else {
		return nil, errors.New("TANGGAL ARTIKEL SUDAH ADA")
	}
}

// Create func Getter
// supaya data struct utama bisa di akses dari luar
func (data *Article) GetCodeArtikel() string {
	return data.codeArticle
}

func (data *Article) GetTitleArtikel() string {
	return data.titleOriginal
}

func (data *Article) GetTextArtikel() string {
	return data.textOriginal
}

func (data *Article) GetTanggalArtikel() time.Time {
	return data.date
}

func (data *Article) GetBannerArtikel() string {
	return data.banner
}

func (data *Article) GetAuthorArtikel() string {
	return data.author
}

func (data *Article) GetThumbsArtikel() string {
	return data.thumbs
}

func (data *Article) GetIsHighLightArtikel() bool {
	return data.isHighlight
}

func (data *Article) GetTranslationAll() []*Translation {
	return data.translation
}

func (data *Translation) GetCodeLanguageTranslation() string {
	return data.codeLanguage
}

func (data *Translation) GetTitleTranslation() string {
	return data.title
}

func (data *Translation) GetTextTranslation() string {
	return data.text
}
