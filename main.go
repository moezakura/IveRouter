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

		fmt.Printf("%17s : ", "MAC ADDRESS")
		fmt.Printf("%32s   ", "DOWNLOAD")
		fmt.Printf("%32s", "UPLOAD")
		fmt.Println()

		for _, d := range devices {
			uploadTraffic := d.UploadTraffic
			downloadTraffic := d.DownloadTraffic
			upload := util.ToDataCast(float64(uploadTraffic))
			download := util.ToDataCast(float64(downloadTraffic))
			fmt.Printf("%s : ", d.MacAddress)
			fmt.Printf("↑ %10s (% 15d b)", upload, uploadTraffic)
			fmt.Printf("   ")
			fmt.Printf("↓ %10s (% 15d b)", download, downloadTraffic)
			fmt.Printf("\n")
		}
	}
}
