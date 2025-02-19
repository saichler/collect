package boot

import "github.com/saichler/collect/go/types"

var DEFAULT_CADENCE int64 = 900
var DEFAULT_TIMEOUT int64 = 30

const (
	BOOT_GROUP = "BOOT"
)

func CreateBootPolls() []*types.Poll {
	result := make([]*types.Poll, 0)
	result = append(result, createSystemMibPoll())
	return result
}

func createSystemMibPoll() *types.Poll {
	poll := createBaseSNMPPoll("systemMib", BOOT_GROUP)
	poll.What = ".1.3.6.1.2.1.1"
	poll.Operation = types.Operation__Map
	poll.Attributes = make([]*types.Attribute, 0)
	poll.Attributes = append(poll.Attributes, createVendor())
	poll.Attributes = append(poll.Attributes, createSysName())
	return poll
}

func createVendor() *types.Attribute {
	attr := &types.Attribute{}
	attr.InstancePath = "networkbox.info.vendor"
	attr.Rules = make([]*types.Rule, 0)
	attr.Rules = append(attr.Rules, createContainsRule("cisco", ".1.3.6.1.2.1.1.1.0", "Cisco"))
	attr.Rules = append(attr.Rules, createContainsRule("ubuntu", ".1.3.6.1.2.1.1.1.0", "Ubuntu Linux"))
	return attr
}

func createSysName() *types.Attribute {
	attr := &types.Attribute{}
	attr.InstancePath = "networkbox.info.sysname"
	attr.Rules = make([]*types.Rule, 0)
	attr.Rules = append(attr.Rules, createSetRule(".1.3.6.1.2.1.1.5.0"))
	return attr
}

func addParameter(name, value string, rule *types.Rule) {
	param := &types.Parameter{}
	param.Name = name
	param.Value = value
	rule.Params[name] = param
}

func createContainsRule(what, from, output string) *types.Rule {
	rule := &types.Rule{}
	rule.Name = "Contains"
	rule.Params = make(map[string]*types.Parameter)
	addParameter("what", what, rule)
	addParameter("from", from, rule)
	addParameter("output", output, rule)
	return rule
}

func createSetRule(from string) *types.Rule {
	rule := &types.Rule{}
	rule.Name = "Set"
	rule.Params = make(map[string]*types.Parameter)
	addParameter("from", from, rule)
	return rule
}

func createBaseSNMPPoll(name string, groups ...string) *types.Poll {
	poll := &types.Poll{}
	poll.Name = name
	poll.Groups = groups
	poll.DefaultTimeout = DEFAULT_TIMEOUT
	poll.DefaultCadence = DEFAULT_CADENCE
	poll.Protocol = types.Protocol_SNMPV2
	return poll
}
