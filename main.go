package main

import (
	"fmt"
	"log"

	"github.com/terra-farm/go-virtualbox"
)
type VMInfo struct{
	Name string
	UUID string
	State string
	CPUCount int
	MemorySize int
	VRAM int
	OSType string
	Network []NIC
	NATNet []Natnet
}
type NIC struct{
	Network string
	Hardware string
	HostInterface string
	MAC string
}
type Natnet struct{
	Name string
	IPv4 string
	IPv6 string
}


func main(){
	vms,err:=virtualbox.ListMachines()
	if err!=nil{
		log.Printf("Error listing VMs:%v\n",err)
	}
	NAT,err:=virtualbox.NATNets()
		if err!=nil{
			log.Printf("Error listing NatNets:%v\n",err)
		}
	for _,vm:=range vms{

		nics:=vm.NICs
		var network []NIC
		var natnets []Natnet	
		for _,nic:=range(nics){
		network=append(network,NIC{
			Network: string(nic.Network),
			Hardware: string(nic.Hardware),
			HostInterface: nic.HostInterface,
			MAC:nic.MacAddr,
		})
	}
	
		for _,natnet:=range NAT{
			fmt.Print(natnet)
			natnets=append(natnets,Natnet{	
				Name: natnet.Name,
				IPv4: natnet.IPv4.String(),
				IPv6: natnet.IPv6.String(),
		})
	}
	
		vmInfo:=VMInfo{
			Name:vm.Name,
			UUID: vm.UUID,
			State: string(vm.State),
			CPUCount: int(vm.CPUs),
			MemorySize: int(vm.Memory),
			VRAM: int(vm.VRAM),
			OSType: vm.OSType,
			Network: network,
			NATNet: natnets,
		}
		displayVMInfo(vmInfo)
	}
}
func displayVMInfo(vm VMInfo){
	fmt.Printf("Name:%s\n",vm.Name)
	fmt.Printf("UUID:%s\n",vm.UUID)
	fmt.Printf("State:%s\n",vm.State)
	fmt.Printf("CPU Count:%d\n",vm.CPUCount)
	fmt.Printf("Memory Size:%dMb\n",vm.MemorySize)
	fmt.Printf("VRAM:%d\nMb",vm.VRAM)
	fmt.Printf("OS Type:%s\n",vm.OSType)
	for _,nic:=range(vm.Network){
		fmt.Println("Network Interface Card info:")
		fmt.Printf("*********\n")
		fmt.Printf("Network:%s,\nHardware:%s,\nHost-Interface:%s,\nMAC:%s\n",nic.Network,nic.Hardware,nic.HostInterface,nic.MAC)
	}
	fmt.Println("**********")
	for _,natnet:=range(vm.NATNet){
		fmt.Println("NAT netword Info:")
		fmt.Printf("*********\n")
		fmt.Printf("Name:%s,\nIPv4 address:%s,\nIPv6 Address:%s\n",natnet.Name,natnet.IPv4,natnet.IPv6)
	}

	fmt.Print("################\n")
}