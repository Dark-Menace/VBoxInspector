
# VBoxInspector

VBoxInspector is a tool designed to manage and retrieve information about VMs and their network configurations using the go-virtualbox library. The primary objectives of this project are to provide detailed information about virtual machines, including their state, hardware configuration, and network interfaces, and to help administrators monitor and manage their virtual environments effectively.

## Features
#### VM Information Retrieval:
VBoxNetInspector retrieves comprehensive information about VMs, including name, UUID, state, CPU count, memory size, VRAM, OS type, and network configurations. This feature helps administrators get a clear overview of their virtual environments.

#### Network Interface Details:
The tool provides detailed information about each VM's network interface cards (NICs), including network type, hardware type, host interface, and MAC address. This ensures that administrators can monitor and manage network configurations effectively.

#### NAT Network Information:
VBoxNetInspector also retrieves and displays information about NAT networks associated with the VMs, including their names, IPv4, and IPv6 addresses. This helps in understanding and managing the network setups of VMs.
## Setting up and Running the Project


Clone the project. Enter the project directory and then install the necessary go dependencies. 

```bash
git clone https://github.com/Dark-Menace/VBoxInspector.git
cd VBoxInspector
go mod tidy
go run .
```

