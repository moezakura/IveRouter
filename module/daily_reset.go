package module

import (
	"sync"
	"time"
)

type DailyReset struct {
	packetLive *PacketLive

	lastDate string
}

func NewDailyReset(packetLive *PacketLive) *DailyReset {
	now := time.Now()
	today := now.Format("2006-01-02")
	return &DailyReset{
		packetLive: packetLive,
		lastDate:   today,
	}
}

func (d *DailyReset) Run() {
	go func() {
		t := time.NewTicker(50 * time.Millisecond)
		for {
			<-t.C
			d.do()
		}
	}()
}

func (d *DailyReset) do() {
	now := time.Now()
	today := now.Format("2006-01-02")
	if d.lastDate == today {
		return
	}
	d.lastDate = today
	d.packetLive.devices = sync.Map{}
}
