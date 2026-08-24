package domainmaster

import "time"

type Urusan struct {
	Id           string
	KodeUrusan   string
	NamaUrusan   string
	CreatedAt    time.Time
	BidangUrusan []BidangUrusan
}
