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

	drs := module.NewDailyReset(pl)
	drs.Run()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var devices model.Devices = pl.GetDevices()
		sort.Sort(sort.Reverse(devices))
		fmt.Println("==== DEVICE TRAFFIC ====")

		fmt.Printf("%13s%s%13s  ", "", "UPLOAD", "")
		fmt.Printf(" : %3s%s%3s : ", "", "MAC ADDRESS", "")
		fmt.Printf("  %12s%s%12s", "", "DOWNLOAD", "")
		fmt.Println()

		for _, d := range devices {
			uploadTraffic := d.UploadTraffic
			downloadTraffic := d.DownloadTraffic
			upload := util.ToDataCast(float64(uploadTraffic))
			download := util.ToDataCast(float64(downloadTraffic))
			fmt.Printf("↑ %10s (% 15d b)", upload, uploadTraffic)
			fmt.Printf("  ")
			fmt.Printf(" : %s : ", d.MacAddress)
			fmt.Printf("  ")
			fmt.Printf("↓ %10s (% 15d b)", download, downloadTraffic)
			fmt.Printf("\n")
		}
	}
}
