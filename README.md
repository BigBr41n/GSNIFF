## GSNIFF

A simple TCP packet sniffer using `gopacket` package.

## Introduction

The Packet GSNIFF is a basic tool for network analysis, similar in concept to Wireshark, but much simpler. It captures TCP packets.

## Prerequisites

Before you can run the packet sniffer, you need to have the following installed:

- Go (Golang) - [Installation guide](https://golang.org/doc/install)
- libpcap - On Unix-based systems, this can usually be installed via your package manager (e.g., `apt` for Ubuntu, `yum` for CentOS).
  - For Debian/Ubuntu: `sudo apt-get install libpcap-dev`
  - For RedHat/CentOS: `sudo yum install libpcap-devel`
  - For macOS: `brew install libpcap` (using Homebrew)
- WinPcap or Npcap for Windows users - [Npcap](https://nmap.org/npcap/)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/BigBr41n/GSNIFF.git
   ```
2. Navigate to the cloned repository:
   ```bash
   cd GSNIFF
   ```

## Usage

1. To run the packet sniffer, use the following command:
   ```bash
   make
   sudo ./GSNIFF
   ```

## Disclaimer

This tool is for educational purposes only. Unauthorized packet sniffing or network analysis can be illegal and unethical. Always ensure you have permission to capture packets on the network you're monitoring.

## Contributing

Contributions to the Go Packet Sniffer are welcome!

## License

This project is licensed under the [GPL License](LICENSE) - see the [LICENSE](LICENSE) file for details.
