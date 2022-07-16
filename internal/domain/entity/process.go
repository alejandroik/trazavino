package entity

import (
	"time"
)

type Process struct {
	id          int64
	startDate   time.Time
	endDate     time.Time
	hash        string
	transaction string
	pType       string
	previousId  int64
}

func NewProcess(
	id int64,
	startDate time.Time,
	endDate time.Time,
	hash string,
	transaction string,
	pType string,
	previousId int64) (*Process, error) {
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

func (p Process) ID() int64 {
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

func (p Process) PreviousID() int64 {
	return p.previousId
}

func (p *Process) UpdateEndDate(date time.Time) error {
	p.endDate = date

	return nil
}

func (p *Process) UpdatePreviousID(id int64) error {
	p.previousId = id

	return nil
}

func (p *Process) UpdateHash(hash string) error {
	p.hash = hash

	return nil
}

func (p *Process) UpdateTransaction(transaction string) error {
	p.transaction = transaction

	return nil
}

func UnmarshalProcessFromDatabase(
	id int64,
	startDate time.Time,
	endDate time.Time,
	hash string,
	transaction string,
	ptype string,
	previousId int64,
) (*Process, error) {
	p, err := NewProcess(id, startDate, endDate, hash, transaction, ptype, previousId)
	if err != nil {
		return nil, err
	}

	return p, nil
}
