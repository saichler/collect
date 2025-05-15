package main

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8types/go/ifs"
)

var Plugin ifs.IPlugin = &DeviceConfigRegistryPlugin{}

type DeviceConfigRegistryPlugin struct{}

func (this *DeviceConfigRegistryPlugin) Install(vnic ifs.IVNic) error {
	vnic.Resources().Registry().Register(&types.DeviceConfig{})
	return nil
}
