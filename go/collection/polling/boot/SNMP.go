package boot

import (
	"github.com/saichler/collect/go/collection/parsing/rules"
	"github.com/saichler/collect/go/types"
	"strconv"
)

var DEFAULT_CADENCE int64 = 900
var DEFAULT_TIMEOUT int64 = 30

const (
	BOOT_GROUP = "BOOT"
)

func CreateSNMPBootPolls() []*types.Poll {
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
	attr.PropertyId = "networkbox.info.vendor"
	attr.Rules = make([]*types.Rule, 0)
	attr.Rules = append(attr.Rules, createContainsRule("cisco", ".1.3.6.1.2.1.1.1.0", "Cisco"))
	attr.Rules = append(attr.Rules, createContainsRule("ubuntu", ".1.3.6.1.2.1.1.1.0", "Ubuntu Linux"))
	return attr
}

func createSysName() *types.Attribute {
	attr := &types.Attribute{}
	attr.PropertyId = "networkbox.info.sysname"
	attr.Rules = make([]*types.Rule, 0)
	attr.Rules = append(attr.Rules, createSetRule(".1.3.6.1.2.1.1.5.0"))
	return attr
}

func createVersion() *types.Attribute {
	attr := &types.Attribute{}
	attr.PropertyId = "networkbox.info.vendor"
	attr.Rules = make([]*types.Rule, 0)
	attr.Rules = append(attr.Rules, createContainsRule("cisco", ".1.3.6.1.2.1.1.1.0", "Cisco"))
	attr.Rules = append(attr.Rules, createContainsRule("ubuntu", ".1.3.6.1.2.1.1.1.0", "Ubuntu Linux"))
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

func createToTable(columns, keycolumn int) *types.Rule {
	rule := &types.Rule{}
	rule.Name = "ToTable"
	rule.Params = make(map[string]*types.Parameter)
	rule.Params[rules.Columns] = &types.Parameter{Name: rules.Columns, Value: strconv.Itoa(columns)}
	rule.Params[rules.KeyColumn] = &types.Parameter{Name: rules.KeyColumn, Value: strconv.Itoa(keycolumn)}
	return rule
}

func createTableToMap() *types.Rule {
	rule := &types.Rule{}
	rule.Name = "TableToMap"
	rule.Params = make(map[string]*types.Parameter)
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
