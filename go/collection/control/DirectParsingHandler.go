package control

import (
	"github.com/saichler/collect/go/collection/parsing"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

type DirectParsingHandler struct {
	resources common.IResources
	any       interface{}
}

func NewDirectParsingHandler(any interface{}, resources common.IResources) *DirectParsingHandler {
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
