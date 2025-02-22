package boot

import "github.com/saichler/collect/go/types"

func CreateK8sBootPolls() []*types.Poll {
	result := make([]*types.Poll, 0)
	result = append(result, createNodesPoll())
	return result
}

func createNodesPoll() *types.Poll {
	poll := createBaseK8sPoll("nodes", BOOT_GROUP)
	poll.What = "get nodes -o wide"
	poll.Operation = types.Operation__Table
	poll.Attributes = make([]*types.Attribute, 0)
	poll.Attributes = append(poll.Attributes, createNodesTable())
	return poll
}

func createNodesTable() *types.Attribute {
	attr := &types.Attribute{}
	attr.PropertyId = "cluster.nodes"
	attr.Rules = make([]*types.Rule, 0)
	attr.Rules = append(attr.Rules, createToTable(10))
	attr.Rules = append(attr.Rules, createTableToMap())
	return attr
}

func createBaseK8sPoll(name string, groups ...string) *types.Poll {
	poll := &types.Poll{}
	poll.Name = name
	poll.Groups = groups
	poll.DefaultTimeout = DEFAULT_TIMEOUT
	poll.DefaultCadence = DEFAULT_CADENCE
	poll.Protocol = types.Protocol_K8s
	return poll
}
