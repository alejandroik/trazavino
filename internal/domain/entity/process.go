package entity

import (
	"time"

	"github.com/pkg/errors"
)

type Process struct {
	uuid string

	startTime time.Time

	pType string

	endTime      time.Time
	previousUUID string

	hash        string
	transaction string
}

func NewProcess(
	uuid string,
	startTime time.Time,
	pType string,
) (*Process, error) {
	if uuid == "" {
		return nil, errors.New("empty process uuid")
	}
	if startTime.IsZero() {
		return nil, errors.New("zero process start time")
	}
	if pType == "" {
		return nil, errors.New("empty process type")
	}

	return &Process{
		uuid:      uuid,
		startTime: startTime,
		pType:     pType,
	}, nil
}

func (p Process) UUID() string {
	return p.uuid
}

func (p Process) StartTime() time.Time {
	return p.startTime
}

func (p Process) EndTime() time.Time {
	return p.endTime
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

func (p *Process) UpdateEndTime(date time.Time) error {
	p.endTime = date

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
	startTime time.Time,
	ptype string,
	endTime time.Time,
	previousUUID string,
	hash string,
	transaction string,
) (*Process, error) {
	p, err := NewProcess(uuid, startTime, ptype)
	if err != nil {
		return nil, err
	}

	p.endTime = endTime
	p.previousUUID = previousUUID
	p.hash = hash
	p.transaction = transaction

	return p, nil
}
