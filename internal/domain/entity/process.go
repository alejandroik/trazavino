package entity

import (
	"time"
)

type Process struct {
	id          int
	startDate   time.Time
	endDate     time.Time
	hash        string
	transaction string
	pType       string
	previousId  int
}

func NewProcess(
	startDate time.Time,
	endDate time.Time,
	hash string,
	transaction string,
	pType string,
	previousId int) (*Process, error) {
	return newProcess(0, startDate, endDate, hash, transaction, pType, previousId)
}

func newProcess(
	id int,
	startDate time.Time,
	endDate time.Time,
	hash string,
	transaction string,
	pType string,
	previousId int) (*Process, error) {
	return &Process{
		id:          id,
		startDate:   startDate,
		endDate:     endDate,
		hash:        hash,
		transaction: transaction,
		pType:       pType,
		previousId:  previousId,
	}, nil
}

func (p Process) Id() int {
	return p.id
}

func (p Process) StartDate() time.Time {
	return p.startDate
}

func (p Process) EndDate() time.Time {
	return p.endDate
}

func (p Process) Hash() string {
	return p.hash
}

func (p Process) Transaction() string {
	return p.transaction
}

func (p Process) Ptype() string {
	return p.pType
}

func (p Process) PreviousId() int {
	return p.previousId
}

func UnmarshalProcessFromDatabase(
	id int,
	startDate time.Time,
	endDate time.Time,
	hash string,
	transaction string,
	ptype string,
	previousId int,
) (*Process, error) {
	p, err := newProcess(id, startDate, endDate, hash, transaction, ptype, previousId)
	if err != nil {
		return nil, err
	}

	return p, nil
}
