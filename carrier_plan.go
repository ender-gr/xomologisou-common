package xomologisou

import (
	"github.com/go-pg/pg/orm"
	"time"
)

type PaymentType int

const (
	PaymentFree PaymentType = iota
)

type CarrierPlan struct {
	Id      string      `json:"id"`
	Carrier string      `json:"carrier"`
	Type    PaymentType `json:"type"`

	ExpiresAt      time.Time `json:"expires_at"`
	QuotaReset     time.Time `json:"quota_reset"`
	QuotaTotal     int       `json:"quota_total"`
	QuotaAvailable int       `json:"quota_available"`
}

func FindCarrierPlan(db orm.DB, id string) (*CarrierPlan, error) {
	c := &CarrierPlan{}
	err := db.Model(c).Where("id = ?", id).Select()
	return c, err
}

func FindCarrierPlanByCarrier(db orm.DB, carrier string) (*CarrierPlan, error) {
	c := &CarrierPlan{}
	err := db.Model(c).Where("carrier = ?", carrier).Select()
	return c, err
}

func InsertCarrierPlan(db orm.DB, CarrierPlan *CarrierPlan) error {
	CarrierPlan.Id = RandomId(32)
	return db.Insert(CarrierPlan)
}

func SaveCarrierPlan(db orm.DB, CarrierPlan *CarrierPlan) error {
	return db.Update(CarrierPlan)
}

func DeleteCarrierPlan(db orm.DB, CarrierPlan *CarrierPlan) error {
	return db.Delete(CarrierPlan)
}
