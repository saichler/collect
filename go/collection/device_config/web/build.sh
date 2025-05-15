go build -buildmode=plugin -o DeviceConfig-0-registry.so DeviceConfigRegistryPlugin.go
cp DeviceConfig-0-registry.so DeviceConfig-1-registry.so
cp DeviceConfig-0-registry.so DeviceConfig-2-registry.so
