go build -buildmode=plugin -o PollConfig-0-registry.so PollConfigRegistryPlugin.go
cp PollConfig-0-registry.so PollConfig-1-registry.so
cp PollConfig-0-registry.so PollConfig-2-registry.so
