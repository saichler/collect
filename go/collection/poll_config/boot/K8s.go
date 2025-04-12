package boot

import "github.com/saichler/collect/go/types"

func CreateK8sBootPolls() []*types.PollConfig {
	result := make([]*types.PollConfig, 0)
	result = append(result, createNodesPoll())
	result = append(result, createPodsPoll())
	return result
}

func createNodesPoll() *types.PollConfig {
	poll := createBaseK8sPoll("nodes", BOOT_GROUP)
	poll.What = "get nodes -o wide"
	poll.Operation = types.Operation__Table
	poll.Attributes = make([]*types.Attribute, 0)
	poll.Attributes = append(poll.Attributes, createNodesTable())
	return poll
}

func createPodsPoll() *types.PollConfig {
	poll := createBaseK8sPoll("pods", BOOT_GROUP)
	poll.What = "get pods -A -o wide"
	poll.Operation = types.Operation__Table
	poll.Attributes = make([]*types.Attribute, 0)
	poll.Attributes = append(poll.Attributes, createPodsTable())
	return poll
}

func createNodesTable() *types.Attribute {
	attr := &types.Attribute{}
	attr.PropertyId = "cluster.nodes"
	attr.Rules = make([]*types.Rule, 0)
	attr.Rules = append(attr.Rules, createToTable(10, 0))
	attr.Rules = append(attr.Rules, createTableToMap())
	return attr
}

func createPodsTable() *types.Attribute {
	attr := &types.Attribute{}
	attr.PropertyId = "cluster.pods"
	attr.Rules = make([]*types.Rule, 0)
	attr.Rules = append(attr.Rules, createToTable(10, 6))
	attr.Rules = append(attr.Rules, createTableToMap())
	return attr
}

func createBaseK8sPoll(name string, groups ...string) *types.PollConfig {
	poll := &types.PollConfig{}
	poll.Name = name
	poll.Groups = groups
	poll.DefaultTimeout = DEFAULT_TIMEOUT
	poll.DefaultCadence = DEFAULT_CADENCE
	poll.Protocol = types.Protocol_K8s
	return poll
}
