package module

import (
	"fmt"
	"github.com/moezakura/IveRouter/server/model"
	"log"
	"os"
	"time"
)

const (
	AUTO_SAVE_FILE = "data/devices_%s.dat"
)

type AutoSave struct {
	packetLive *PacketLive
}

func NewAutoSave(packetLive *PacketLive) *AutoSave {
	return &AutoSave{
		packetLive: packetLive,
	}
}

func (a *AutoSave) Run() {
	go func() {
		t := time.NewTicker(5 * time.Second)
		for {
			<-t.C
			a.do()
		}
	}()
}

func (a *AutoSave) do() {
	var d model.Devices = a.packetLive.GetDevices()
	data, err := d.Encode()

	if err != nil {
		log.Printf("devices data encode error: %+v\n", err)
		return
	}

	now := time.Now()
	filePath := fmt.Sprintf(AUTO_SAVE_FILE, now.Format("2006-01-02"))
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("devices data save file create error: %+v\n", err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("devices data close file error: %+v\n", err)
		}
	}()

	_, err = file.Write(data)
	if err != nil {
		log.Printf("devices data write(save) file error: %+v\n", err)
	}
}
