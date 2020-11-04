package xomologisou

import (
	"github.com/go-pg/pg/orm"
	"golang.org/x/oauth2"
	"strconv"
	"strings"
)

type Carrier struct {
	Id   string `sql:",pk"`
	Name string

	Suspended bool `json:"suspended"`

	FacebookPage string               `json:"facebook_page"`
	FacebookInfo *CarrierFacebookInfo `json:"facebook_info"`
	Form         *CarrierForm         `json:"form"`
	Posting      *CarrierPosting      `json:"posting"`

	Statistics map[string]interface{} `json:"statistics"`
}

type CarrierForm struct {
	Enabled bool `json:"enabled"`

	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Terms    string `json:"terms"`

	SecretPrompt  string `json:"secret_prompt"`
	ImagePrompt   string `json:"image_prompt"`
	SentMessage   string `json:"sent_message"`
	BackgroundUrl string `json:"background_url"`

	Dark         bool   `json:"dark"`
	AcceptsImage bool   `json:"accepts_image"`
	CustomCss    string `json:"custom_css"`

	OptionSets map[string]CarrierOptions `json:"option_sets"`
}

type CarrierPosting struct {
	Format  string `json:"format"`
	Hashtag string `json:"hashtag"`
	Id      int    `json:"id"`
}

type CarrierFacebookInfo struct {
	PageName  string
	UserName  string
	UserToken *oauth2.Token
	PageToken *oauth2.Token
	HasToken  bool
}

type CarrierOptions struct {
	Name        string            `json:"name"`
	Options     map[string]string `json:"options"`
	AllowCustom bool              `json:"allow_custom"`
}

func (c *Carrier) NextHashtag() string {
	if c.Posting == nil {
		return ""
	}
	if strings.Contains(c.Posting.Hashtag, "{n}") {
		return strings.Replace(c.Posting.Hashtag, "{n}", strconv.Itoa(c.Posting.Id), -1)
	}
	return c.Posting.Hashtag + strconv.Itoa(c.Posting.Id)
}

func FindCarrier(db orm.DB, id string) (*Carrier, error) {
	c := &Carrier{}
	err := db.Model(c).Where("id = ?", id).Select()
	return c, err
}

func FindCarrierByFacebookPage(db orm.DB, page string) (*Carrier, error) {
	c := &Carrier{}
	err := db.Model(c).Where("facebook_page = ?", page).Select()
	return c, err
}

func FindCarriers(db orm.DB) ([]*Carrier, error) {
	var results []*Carrier
	err := db.Model(&results).OrderExpr("name ASC").Select()
	return results, err
}

func FindFrondCarriers(db orm.DB) ([]*Carrier, error) {
	var results []*Carrier
	err := db.Model(&results).
		Where("(statistics->>'likes')::int IS NOT NULL").
		OrderExpr("statistics->>'likes' DESC").
		Limit(9).Select()
	return results, err
}

func InsertCarrier(db orm.DB, carrier *Carrier) error {
	return db.Insert(carrier)
}

func SaveCarrier(db orm.DB, carrier *Carrier) error {
	return db.Update(carrier)
}

func DeleteCarrier(db orm.DB, carrier *Carrier) error {
	return db.Delete(carrier)
}
