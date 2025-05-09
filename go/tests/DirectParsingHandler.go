package tests

import (
	"github.com/saichler/collect/go/collection/parsing"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8types/go/ifs"
)

type DirectParsingHandler struct {
	resources ifs.IResources
	any       interface{}
}

func NewDirectParsingHandler(any interface{}, resources ifs.IResources) *DirectParsingHandler {
	handler := &DirectParsingHandler{}
	handler.resources = resources
	handler.any = any
	return handler
}

func (this *DirectParsingHandler) HandleCollectNotification(job *types.Job) {
	err := parsing.Parser.Parse(job, this.any, this.resources)
	if err != nil {
		this.resources.Logger().Error(err.Error())
	}
}
