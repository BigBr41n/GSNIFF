package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)


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

	//TODO: process all packets
	//packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// handle the packets 
	//handlePackets(packetSource);
}