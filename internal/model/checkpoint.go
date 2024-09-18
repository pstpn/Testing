package model

import "time"

type CheckpointID int64

func ToCheckpointID(id int64) *CheckpointID {
	checkpointID := CheckpointID(id)
	return &checkpointID
}

func (c *CheckpointID) Int() int64 {
	return int64(*c)
}

type Checkpoint struct {
	ID          *CheckpointID
	PhoneNumber string
}

type PassageID int64

func ToPassageID(id int64) *PassageID {
	passageID := PassageID(id)
	return &passageID
}

func (p *PassageID) Int() int64 {
	return int64(*p)
}

type PassageType int64

const (
	Entrance PassageType = iota
	Exit
	UnknownPassageType
)

func ToPassageTypeFromInt(passage int64) *PassageType {
	passageType := PassageType(passage)
	return &passageType
}

func ToPassageTypeFromString(passage string) *PassageType {
	var passageType PassageType
	switch passage {
	case "Вход":
		passageType = Entrance
	case "Выход":
		passageType = Exit
	default:
		passageType = UnknownPassageType
	}

	return &passageType
}

func (p *PassageType) String() string {
	switch *p {
	case Entrance:
		return "Вход"
	case Exit:
		return "Выход"
	default:
		return "Неизвестно"
	}
}

type Passage struct {
	ID           *PassageID
	CheckpointID *CheckpointID
	DocumentID   *DocumentID
	Type         *PassageType
	Time         *time.Time
}

type ShortPassage struct {
	Type string `json:"type"`
	Time string `json:"time"`
}

func ModelToShortPassages(passages []*Passage) []*ShortPassage {
	p := make([]*ShortPassage, 0)
	for _, passage := range passages {
		p = append(p, &ShortPassage{
			Type: passage.Type.String(),
			Time: passage.Time.Add(3 * time.Hour).Format("15:04:05 (02.01.2006)"),
		})
	}
	return p
}
