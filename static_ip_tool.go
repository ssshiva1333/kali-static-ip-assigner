package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	iface_file  = "/etc/network/interfaces"
	user_manual = "* Parameters -> iface_name inet(4/6) ip_address netmask gateway dns_server\n"
)

func main() {
	fmt.Println(user_manual)

	var iface_name, inet, ip_address, netmask, gateway, dns_server string
	fmt.Print("> ")
	fmt.Scanln(&iface_name, &inet, &ip_address, &gateway, &netmask, &dns_server)

	write_ip_addr(iface_name, inet, ip_address, netmask, gateway, dns_server)
}

func write_ip_addr(iface_name, inet, ip_addr, netmask, gateway, dns_servers string) {
	content := "\niface " + iface_name + " " + inet + "\n"
	content += "address " + ip_addr + "\n"
	content += "netmask " + netmask + "\n"
	content += "gateway " + gateway + "\n"
	content += "dns_nameservers " + dns_servers

	file, file_error := os.OpenFile(iface_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if file_error != nil {
		fmt.Println("* ", file_error.Error())
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, writing_error := writer.WriteString(content)
	if writing_error != nil {
		fmt.Println("* ", writing_error.Error())
	}

	flushing_error := writer.Flush()
	if flushing_error != nil {
		fmt.Println("* ", flushing_error.Error())
	}
}
