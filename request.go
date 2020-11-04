package xomologisou

import "github.com/go-pg/pg/orm"

type RequestStatus int

const (
	RequestSubmitted RequestStatus = iota
	RequestRevoked
	RequestAccepted
	RequestRejected
)

type Request struct {
	Id string `json:"id" sql:",pk"`

	Status RequestStatus `json:"status"`
}

func FindRequest(db orm.DB, id string) (*Request, error) {
	c := &Request{}
	err := db.Model(c).Where("id = ?", id).Select()
	return c, err
}

func InsertRequest(db orm.DB, Request *Request) error {
	Request.Id = RandomId(32)
	return db.Insert(Request)
}

func SaveRequest(db orm.DB, Request *Request) error {
	return db.Update(Request)
}

func DeleteRequest(db orm.DB, Request *Request) error {
	return db.Delete(Request)
}
