package model

import (
	"bytes"
	"encoding/gob"
)

type Devices []Device

func (d Devices) Len() int {
	return len(d)
}

func (d Devices) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Devices) Less(i, j int) bool {
	return d[i].Traffic < d[j].Traffic
}

func (d Devices) Encode() (data []byte, err error) {
	buf := bytes.NewBuffer(nil)
	err = gob.NewEncoder(buf).Encode(&d)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d Devices) Decode(data []byte) error {
	buf := bytes.NewBuffer(data)
	err := gob.NewDecoder(buf).Decode(&d)
	return err
}
