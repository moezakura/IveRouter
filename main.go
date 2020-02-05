package main

import (
	"bufio"
	"fmt"
	"github.com/moezakura/IveRouter/model"
	"github.com/moezakura/IveRouter/module"
	"os"
	"sort"
)

func main() {
	pl := module.NewPacketLive()
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
			trafficFormat := Util.toDataCast(float64(traffic))
			fmt.Printf("%s : %s (%d b)\n", d.MacAddress, trafficFormat, traffic)
		}
	}
}
