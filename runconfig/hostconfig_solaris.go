package runconfig

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/sysinfo"
)

// DefaultDaemonNetworkMode returns the default network stack the daemon should
// use.
func DefaultDaemonNetworkMode() container.NetworkMode {
	return container.NetworkMode("bridge")
}

// IsPreDefinedNetwork indicates if a network is predefined by the daemon
func IsPreDefinedNetwork(network string) bool {
	return false
}

// ValidateNetMode ensures that the various combinations of requested
// network settings are valid.
func ValidateNetMode(c *container.Config, hc *container.HostConfig) error {
	// We may not be passed a host config, such as in the case of docker commit
	return nil
}

// ValidateIsolation performs platform specific validation of the
// isolation level in the hostconfig structure.
// This setting is currently discarded for Solaris so this is a no-op.
func ValidateIsolation(hc *container.HostConfig) error {
	return nil
}

// ValidateQoS performs platform specific validation of the QoS settings
func ValidateQoS(hc *container.HostConfig) error {
	return nil
}

// ValidateResources performs platform specific validation of the resource settings
func ValidateResources(hc *container.HostConfig, si *sysinfo.SysInfo) error {
	return nil
}
