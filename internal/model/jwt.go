package model

import "strconv"

type Payload struct {
	InfoCardID string
}

func (p *Payload) ToString() string {
	return p.InfoCardID
}

func (p *Payload) ParseFromString(payloadString string) {
	p.InfoCardID = payloadString
}

func (p *Payload) GetInfoCardID() (int64, error) {
	return strconv.ParseInt(p.InfoCardID, 10, 64)
}
