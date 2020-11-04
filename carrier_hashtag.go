package xomologisou

import "github.com/go-pg/pg/orm"

type CarrierHashtag struct {
	Carrier   string `json:"carrier"`
	Hashtag   string `json:"hashtag" sql:",pk"`
	MinNumber int    `json:"min_number"`
}

func FindCarrierHashtag(db orm.DB, carrier, hashtag string) (*CarrierHashtag, error) {
	c := &CarrierHashtag{}
	err := db.Model(c).Where("carrier = ?", carrier).Where("hashtag = ?", hashtag).Select()
	return c, err
}

func InsertCarrierHashtag(db orm.DB, CarrierHashtag *CarrierHashtag) error {
	return db.Insert(CarrierHashtag)
}

func SaveCarrierHashtag(db orm.DB, CarrierHashtag *CarrierHashtag) error {
	return db.Update(CarrierHashtag)
}

func DeleteCarrierHashtag(db orm.DB, CarrierHashtag *CarrierHashtag) error {
	return db.Delete(CarrierHashtag)
}
