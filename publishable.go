package xomologisou

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"time"
)

type PublishableStatus int

const (
	PublishableSent PublishableStatus = iota
	PublishableQueued
	PublishablePublishing
	PublishablePublished
	PublishableDeleted
	PublishableDeclined
	PublishableHidden
	PublishableRemovedFB
	PublishableFailed
	PublishablePublishedManually
)

type ArchivedPublishable struct {
	Publishable
	ArchivedAt time.Time
}

type Publishable struct {
	Carrier string `json:"carrier"`
	Id      string `json:"id" sql:",pk"`

	//Who created the post
	Source int `json:"source" sql:",notnull"`
	//The parent post ID
	Parent string `json:"parent"`
	//Disables custom formatting
	DisableFormat bool `json:"disable_format"`

	Status            PublishableStatus `json:"status" sql:",notnull"`
	StatusDescription string            `json:"status_description,omitempty"`

	//Source data
	PublishableSource

	QueuedTime  time.Time `json:"queued_time,omitempty"`
	PublishTime time.Time `json:"publish_time,omitempty"`
	Publisher   string    `json:"publisher,omitempty"`
	PublishTag  string    `json:"publish_tag,omitempty"`

	//Facebook post details
	FacebookPost string `json:"facebook_post,omitempty"`

	ChecksData []PublishableSource `json:"checks_data,omitempty"`

	Content         string `json:"content"`
	OriginalContent string `json:"original_content"`

	ImageId    string                 `json:"image_id,omitempty"`
	Options    map[string]string      `json:"options,omitempty"`
	Statistics map[string]interface{} `json:"statistics,omitempty"`

	FinalForm  string `json:"final_form" sql:"-"`
	Properties string `json:"properties" sql:"-"`
}

type PublishableSource struct {
	Timestamp time.Time `json:"timestamp,omitempty"`

	IpAddress string `json:"ip_address,omitempty"`

	Hostname      string `json:"hostname,omitempty"`
	ShortHostname string `json:"short_hostname,omitempty"`
	Country       string `json:"country,omitempty"`

	UserAgent string `json:"user_agent,omitempty"`
}

func (p *PublishableSource) ClearSensitiveData() {
	p.IpAddress = ""
	p.Hostname = ""
	p.UserAgent = ""
}

func (p *Publishable) NoFormat(carrier *Carrier) bool {
	return p.DisableFormat || carrier.Posting == nil || carrier.Posting.Hashtag == ""
}

func (p *Publishable) IsChild() bool {
	return p.Parent != ""
}

func FindPublishable(db orm.DB, id string) (*Publishable, error) {
	p := &Publishable{}
	err := db.Model(p).Where("id = ?", id).Select()
	return p, err
}

func FindPublishableByHashtag(db orm.DB, hashtag string) (*Publishable, error) {
	p := &Publishable{PublishTag: hashtag}
	err := db.Model(p).Select()
	return p, err
}

func InsertPublishable(db orm.DB, publishable *Publishable) error {
	publishable.Id = RandomId(12)
	return db.Insert(publishable)
}

func SavePublishable(db orm.DB, publishable *Publishable) error {
	return db.Update(publishable)
}

func ArchivePublishable(db *pg.DB, p *Publishable) (*ArchivedPublishable, error) {
	a := &ArchivedPublishable{*p, time.Now()}
	err := db.RunInTransaction(func(tx *pg.Tx) error {
		if err := tx.Insert(a); err != nil {
			return err
		}
		if err := tx.Delete(p); err != nil {
			return err
		}
		return nil
	})
	return a, err
}
