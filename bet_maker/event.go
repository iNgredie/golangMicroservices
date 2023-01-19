package bet_maker

import "time"

const (
	Open      string = "open"
	FirstWin  string = "firstwin"
	SecondWin string = "secondwin"
)

type Event struct {
	Id          int       `json:"-" db:"id"`
	Coefficient int       `json:"coefficient" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	Status      string    `json:"status" binding:"required"`
}
