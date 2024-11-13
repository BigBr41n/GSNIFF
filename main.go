package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func handlePackets(ps *gopacket.PacketSource){ // ps : packet source
	for packet := range ps.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)

			// Print TCP infos
			log.Printf("From src port: %d to dst port: %d\n", tcp.SrcPort, tcp.DstPort)
			log.Printf("Sequence number: %d\n", tcp.Seq)

			// Print the payload if available (1666 bytes)
			if len(tcp.Payload) > 0 {
				log.Printf("Payload: %s\n", string(tcp.Payload))
			}
		}
}
}


func main() {
	fmt.Println("your available interfaces :")
	devices, err := pcap.FindAllDevs()
    if err!= nil {
        log.Fatal(err)
    }
    for _, d := range devices {
        fmt.Println(d.Name)
    }

	// read the interface from the user 
	log.Println("Enter the interface name:")
    var interfaceName string
    fmt.Scanln(&interfaceName)

    // Open the device for capturing packets and without a timeout (BlockForever)
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var filter string = "tcp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	//process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// handle the packets 
	handlePackets(packetSource);
}