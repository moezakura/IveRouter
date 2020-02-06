package module

import (
	"fmt"
	"github.com/moezakura/IveRouter/server/model"
	"github.com/moezakura/IveRouter/server/util"
	"io/ioutil"
	"time"
)

type DevicesRestore struct {
}

func NewDevicesRestore() *DevicesRestore {
	return &DevicesRestore{}
}

func (d *DevicesRestore) Restore() (devices model.Devices, err error) {
	now := time.Now()
	filePath := fmt.Sprintf(AUTO_SAVE_FILE, now.Format("2006-01-02"))

	devices = model.Devices{}
	if !util.Exists(filePath) {
		return devices, nil
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	ds, err := devices.Decode(b)
	return ds, err
}
