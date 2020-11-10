/*
 * Provides a container compatible interface.
 *
 * This documentation describes the Podman v2.0 RESTful API. It replaces the Podman v1.0 API and was initially delivered along with Podman v2.0.  It consists of a Docker-compatible API and a Libpod API providing support for Podman’s unique features such as pods.  To start the service and keep it running for 5,000 seconds (-t 0 runs forever):  podman system service -t 5000 &  You can then use cURL on the socket using requests documented below.  NOTE: if you install the package podman-docker, it will create a symbolic link for /var/run/docker.sock to /run/podman/podman.sock  See podman-service(1) for more information.  Quick Examples:  'podman info'  curl --unix-socket /run/podman/podman.sock http://d/v1.0.0/libpod/info  'podman pull quay.io/containers/podman'  curl -XPOST --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/images/create?fromImage=quay.io%2Fcontainers%2Fpodman'  'podman list images'  curl --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/libpod/images/json' | jq
 *
 * API version: 0.0.1
 * Contact: podman@lists.podman.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// LinuxMemory for Linux cgroup 'memory' resource management
type LinuxMemory struct {
	// DisableOOMKiller disables the OOM killer for out of memory conditions
	DisableOOMKiller bool `json:"disableOOMKiller,omitempty"`
	// Kernel memory limit (in bytes).
	Kernel int64 `json:"kernel,omitempty"`
	// Kernel memory limit for tcp (in bytes)
	KernelTCP int64 `json:"kernelTCP,omitempty"`
	// Memory limit (in bytes).
	Limit int64 `json:"limit,omitempty"`
	// Memory reservation or soft_limit (in bytes).
	Reservation int64 `json:"reservation,omitempty"`
	// Total memory limit (memory + swap).
	Swap int64 `json:"swap,omitempty"`
	// How aggressive the kernel will swap memory pages.
	Swappiness int32 `json:"swappiness,omitempty"`
	// Enables hierarchical memory accounting
	UseHierarchy bool `json:"useHierarchy,omitempty"`
}