package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)


func main() {
	// TODO: list all devices for the user (interfaces)
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