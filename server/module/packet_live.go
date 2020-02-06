package module

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/moezakura/IveRouter/server/model"
	"log"
	"sync"
	"time"
)

type PacketLive struct {
	handle      *pcap.Handle
	totalLength float64

	devices sync.Map
}

const (
	snapshot_len int32         = 2147483647
	timeout      time.Duration = 30 * time.Second
)

func NewPacketLive(devices model.Devices) *PacketLive {
	devicesSyncMap := sync.Map{}
	for _, d := range devices {
		devicesSyncMap.Store(d.MacAddress, d)
	}

	return &PacketLive{
		devices: devicesSyncMap,
	}
}

func (p *PacketLive) Live(device string) {
	var err error = nil
	p.handle, err = pcap.OpenLive(device, snapshot_len, false, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer p.handle.Close()

	packetSource := gopacket.NewPacketSource(p.handle, p.handle.LinkType())
	for packet := range packetSource.Packets() {
		e := packet.Layer(layers.LayerTypeEthernet)
		if e != nil {
			ep, _ := e.(*layers.Ethernet)
			appLength := p.getPacketSize(packet)
			src := ep.SrcMAC.String()
			dst := ep.DstMAC.String()
			p.addTraffic(src, appLength, true)
			p.addTraffic(dst, appLength, false)
		}
	}
}

func (p *PacketLive) GetDevices() []model.Device {
	devices := make([]model.Device, 0)
	p.devices.Range(func(key, value interface{}) bool {
		devices = append(devices, value.(model.Device))
		return true
	})

	return devices
}

func (p *PacketLive) addTraffic(mac string, appLength uint64, isUpload bool) {
	if _device, ok := p.devices.Load(mac); ok {
		device := _device.(model.Device)
		if isUpload {
			device.UploadTraffic += appLength
		} else {
			device.DownloadTraffic += appLength
		}
		p.devices.Store(mac, device)
	} else {
		var upload uint64 = 0
		var download uint64 = 0

		if isUpload {
			upload = appLength
		} else {
			download = appLength
		}

		p.devices.Store(mac, model.Device{
			MacAddress:      mac,
			UploadTraffic:   upload,
			DownloadTraffic: download,
		})
	}
}

func (*PacketLive) getPacketSize(packet gopacket.Packet) uint64 {
	tcp := packet.Layer(layers.LayerTypeTCP)
	if tcp != nil {
		return uint64(len(tcp.LayerPayload()))
	}

	udp := packet.Layer(layers.LayerTypeUDP)
	if udp != nil {
		return uint64(len(udp.LayerPayload()))
	}

	app := packet.ApplicationLayer()
	if app != nil {
		return uint64(len(app.LayerPayload()))
	}

	return uint64(len(packet.Data()))
}
