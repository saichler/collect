package parsing

import (
	"github.com/saichler/collect/go/collection/parsing/rules"
	"github.com/saichler/collect/go/collection/poll_config"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/types/go/common"
)

type _Parser struct {
	rules map[string]rules.ParsingRule
}

var Parser = newParser()

func newParser() *_Parser {
	p := &_Parser{}
	p.rules = make(map[string]rules.ParsingRule)
	con := &rules.Contains{}
	p.rules[con.Name()] = con
	set := &rules.Set{}
	p.rules[set.Name()] = set
	totable := &rules.ToTable{}
	p.rules[totable.Name()] = totable
	tableToMap := &rules.TableToMap{}
	p.rules[tableToMap.Name()] = tableToMap
	return p
}

func (this *_Parser) Parse(job *types.Job, any interface{}, resources common.IResources) error {
	workSpace := make(map[string]interface{})
	enc := object.NewDecode(job.Result, 0, resources.Registry())
	data, err := enc.Get()
	if err != nil {
		return resources.Logger().Error(err)
	}
	pc := poll_config.PollConfig(resources)
	poll := pc.PollByName(job.PollName)
	if poll == nil {
		return resources.Logger().Error("cannot find poll for name ", job.PollName)
	}
	workSpace[rules.Input] = data
	if poll.Parsing == nil {
		panic("")
	}
	for _, attr := range poll.Parsing.Attributes {
		workSpace[rules.PropertyId] = attr.PropertyId
		for _, rData := range attr.Rules {
			if rData.Params != nil {
				for p, v := range rData.Params {
					workSpace[p] = v.Value
				}
			}
			ruleImpl, ok := this.rules[rData.Name]
			if !ok {
				return resources.Logger().Error("Cannot find parsing rule ", rData.Name)
			}
			err = ruleImpl.Parse(resources, workSpace, rData.Params, any)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
