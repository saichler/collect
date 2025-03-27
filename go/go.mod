module github.com/saichler/collect/go

go 1.24.1

require (
	github.com/google/uuid v1.6.0
	github.com/gosnmp/gosnmp v1.38.0
	github.com/saichler/k8s_observer v0.0.0-20250327191433-107b978228cd
	github.com/saichler/l8test/go v0.0.0-20250327180559-be26a2d56481
	github.com/saichler/layer8/go v0.0.0-20250327183324-ab554613aa33
	github.com/saichler/reflect/go v0.0.0-20250327160656-726672fb5ebf
	github.com/saichler/serializer/go v0.0.0-20250327183844-6ed7a8a5b3ae
	github.com/saichler/servicepoints/go v0.0.0-20250327183714-c026a96f20d6
	github.com/saichler/shared/go v0.0.0-20250327144546-dc40bb3ea146
	github.com/saichler/types/go v0.0.0-20250327162701-de3b6c266ee5
	golang.org/x/crypto v0.35.0
	google.golang.org/protobuf v1.36.6
)

require golang.org/x/sys v0.30.0 // indirect
