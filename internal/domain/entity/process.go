package entity

import (
	"time"
)

type Process struct {
	uuid string

	startDate time.Time
	endDate   time.Time

	hash        string
	transaction string

	pType string

	previousUUID string
}

func NewProcess(
	uuid string,
	startDate time.Time,
	endDate time.Time,
	hash string,
	transaction string,
	pType string,
	previousUUID string) (*Process, error) {
	return &Process{
		uuid:         uuid,
		startDate:    startDate,
		endDate:      endDate,
		hash:         hash,
		transaction:  transaction,
		pType:        pType,
		previousUUID: previousUUID,
	}, nil
}

func (p Process) UUID() string {
	return p.uuid
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

func (p Process) PreviousUUID() string {
	return p.previousUUID
}

func (p *Process) UpdateEndDate(date time.Time) error {
	p.endDate = date

	return nil
}

func (p *Process) UpdatePreviousUUID(uuid string) error {
	p.previousUUID = uuid

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
	uuid string,
	startDate time.Time,
	endDate time.Time,
	hash string,
	transaction string,
	ptype string,
	previousUUID string,
) (*Process, error) {
	p, err := NewProcess(uuid, startDate, endDate, hash, transaction, ptype, previousUUID)
	if err != nil {
		return nil, err
	}

	return p, nil
}
