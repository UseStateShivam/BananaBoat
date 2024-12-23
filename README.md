# BananaBoat 🍌⛵

**Bananaboat** is an ambitious project aimed at building a containerization system from scratch, similar to Docker, using Go. This project will give hands-on experience with containerization, operating system concepts, resource management, and networking, all while strengthening Go skills.

## 📂 Folder Structure



## 🚀 Project Roadmap

### 1. **Master Go Basics** 🧑‍💻
Before diving into building Bananaboat, it’s essential to get comfortable with Go:
- Go Syntax, Data Types, and Control Structures
- Concurrency with Goroutines and Channels
- Error Handling and File Management
- Network Programming & Building HTTP Servers
- Writing Unit Tests

### 2. **Operating System Concepts** 🖥️
Understanding operating system fundamentals is crucial for creating isolated environments:
- Processes & Threads Management
- Linux Namespaces (PID, Mount, Network, etc.)
- Control Groups (cgroups) for Resource Management
- Filesystem Concepts & Mounting
- Union Filesystems (OverlayFS)

### 3. **Understanding Docker Internals** 🐳
Studying Docker's architecture to understand the building blocks of containerization:
- Docker Daemon & Container Lifecycle
- Images and Layered Filesystems
- Networking & Volumes Management
- Docker CLI & Communication Mechanisms

### 4. **Building the Basic Container System** 🏗️
Creating a basic container system with:
- Process Isolation using Namespaces (PID, Mount, Network)
- Container Filesystem with OverlayFS or a Simple Approach
- Start/Stop/Remove Containers Lifecycle Management
- Building Container Images and Layering

### 5. **Networking in Containers** 🌐
Container interaction over the network:
- Setup Network Namespaces for Isolation
- Virtual Ethernet Bridges for Communication
- Implementing Port Mapping (Container-to-Host)
- Container-to-Container Networking & DNS Resolution

### 6. **Image Management** 🖼️
Implement image handling:
- Parsing Dockerfiles to Build Images
- Layering Filesystem for Images (Image Layers)
- Storing and Retrieving Images Locally
- Basic Image Registry Management

### 7. **Resource Management (cgroups)** ⚙️
Control resource limits for containers:
- CPU and Memory Limitations using cgroups
- Disk I/O Management
- Process Resource Isolation with cgroups

### 8. **Persistent Storage (Volumes)** 💾
Manage persistent data for containers:
- Volume Management (Container-to-Host)
- Mounting Volumes Inside Containers
- Sharing Data Between Containers

### 9. **CLI and API Development** 🖥️📱
Create a Command-Line Interface (CLI) for managing containers:
- Build Commands for Starting, Stopping, and Creating Containers
- (Optional) Develop a REST API to Manage Containers Remotely
- Tools like `cobra` or `urfave/cli` for CLI and `Gin` for API

### 10. **Testing and Debugging** 🧪
Ensuring everything works perfectly:
- Write Unit and Integration Tests for Container Creation, Networking, etc.
- Implement Logging and Container Health Checks
- Develop Debugging Tools for Container Interaction

## 📚 Technologies and Tools

- **Go** 🦶: The programming language used to build Bananaboat.
- **Docker Internals** 🐳: Understanding Docker’s inner workings to replicate the functionality.
- **cgroups, Namespaces** 🧑‍🏫: Core concepts in container isolation and resource management.
- **OverlayFS/UnionFS**: Filesystem technologies for image layers.

## 🚧 Development Setup

To get started with the Bananaboat project, follow these steps:

### 1. **Install Go** 🦆
If you don’t have Go installed, download it from the [official Go website](https://golang.org/dl/).

### 2. **Clone the Project** 👨‍💻
```bash
git clone https://github.com/UseStateShivam/BananaBoat.git
cd bananaboat
```

### 3. **Start Coding** ✨
Once the repository is cloned, you can start building the core features based on the roadmap above. 

---

## 📈 Progress & Contribution

This project is under active development, and contributions are welcome! Feel free to:
- Suggest new features 💡
- Report bugs 🐛
- Submit pull requests ✂️

---

## 🎓 Learning Resources

- **Go Documentation**: [Go Docs](https://golang.org/doc/)
- **The Docker Book**: [The Docker Book](https://www.dockerbook.com/)
- **Operating System Concepts**: [OS Concepts](https://pages.cs.wisc.edu/~remzi/OSTEP/)

---

Happy coding! 🍌⛵