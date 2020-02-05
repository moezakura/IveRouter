package main

import (
	"bufio"
	"fmt"
	"github.com/moezakura/IveRouter/model"
	"github.com/moezakura/IveRouter/module"
	"github.com/moezakura/IveRouter/util"
	"os"
	"sort"
)

func main() {
	dr := module.NewDevicesRestore()
	devices, err := dr.Restore()
	if err != nil {
		panic(err)
	}

	pl := module.NewPacketLive(devices)
	go pl.Live("eth0")

	as := module.NewAutoSave(pl)
	as.Run()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var devices model.Devices = pl.GetDevices()
		sort.Sort(sort.Reverse(devices))
		fmt.Println("==== DEVICE TRAFFIC ====")
		for _, d := range devices {
			traffic := d.Traffic
			trafficFormat := util.ToDataCast(float64(traffic))
			fmt.Printf("%s : %s (%d b)\n", d.MacAddress, trafficFormat, traffic)
		}
	}
}
