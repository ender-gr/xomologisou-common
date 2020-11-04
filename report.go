package xomologisou

import (
	"github.com/go-pg/pg/orm"
	"time"
)

type ReportStatus int

const (
	ReportSubmitted ReportStatus = iota
	ReportRevoked
	ReportAccepted
	ReportRejected
	ReportCompleted
	ReportElevated
)

type ReportReason int

const (
	ReportDefamation ReportReason = iota
	ReportSuicide
	ReportOffensive
	ReportSexism
	ReportExtremism
	ReportOther
)

type Report struct {
	Id string `json:"id" sql:",pk"`

	ReportedId    string `json:"reported_id"`
	ReporterEmail string `json:"reporter_email"`

	ReportedCarrier string       `json:"reported_carrier"`
	ReportedSecret  string       `json:"reported_secret"`
	Reason          ReportReason `json:"reason"`
	Description     string       `json:"description"`

	Timestamp time.Time `json:"timestamp"`

	Status   ReportStatus `json:"status"`
	Resolver string       `json:"resolver"`
}

func FindReport(db orm.DB, id string) (*Report, error) {
	c := &Report{}
	err := db.Model(c).Where("id = ?", id).Select()
	return c, err
}

func InsertReport(db orm.DB, Report *Report) error {
	Report.Id = RandomId(16)
	return db.Insert(Report)
}

func SaveReport(db orm.DB, Report *Report) error {
	return db.Update(Report)
}

func DeleteReport(db orm.DB, Report *Report) error {
	return db.Delete(Report)
}
