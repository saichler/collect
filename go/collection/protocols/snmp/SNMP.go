package snmp

import (
	"github.com/gosnmp/gosnmp"
	"github.com/saichler/collect/go/collection/poll"
	"github.com/saichler/collect/go/collection/protocols"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/shared/go/share/interfaces"
	strings2 "github.com/saichler/shared/go/share/strings"
	"strconv"
	"time"
)

type SNMPCollector struct {
	resources interfaces.IResources
	config    *types.Config
	agent     *gosnmp.GoSNMP
	connected bool
}

func (this *SNMPCollector) Protocol() types.Protocol {
	return types.Protocol_SNMPV2
}

func (this *SNMPCollector) Init(conf *types.Config, resources interfaces.IResources) error {
	this.config = conf
	this.resources = resources
	this.agent = &gosnmp.GoSNMP{}
	this.agent.Version = gosnmp.Version2c
	this.agent.Timeout = time.Second * time.Duration(this.config.Timeout)
	this.agent.Target = this.config.Addr
	this.agent.Port = uint16(this.config.Port)
	this.agent.Community = this.config.ReadCommunity
	this.agent.Retries = 1
	return nil
}

func (this *SNMPCollector) Connect() error {
	err := this.agent.Connect()
	if err != nil {
		return err
	}
	this.connected = true
	return nil
}

func (this *SNMPCollector) Disconnect() error {
	this.resources.Logger().Info("SNMP Collector for ", this.config.Addr, " is closed.")
	this.agent = nil
	this.connected = false
	return nil
}

func (this *SNMPCollector) Exec(job *types.Job) {
	if !this.connected {
		err := this.Connect()
		if err != nil {
			job.Error = err.Error()
			return
		}
	}
	pollCenter := poll.Poll(this.resources)
	pll := pollCenter.PollByUuid(job.PollUuid)
	if pll == nil {
		this.resources.Logger().Error("cannot find poll for uuid ", job.PollUuid)
		return
	}
	if pll.Operation == types.Operation__Map {
		this.walk(job, pll, true)
	} else if pll.Operation == types.Operation__Table {
		this.table(job, pll)
	}
}

func (this *SNMPCollector) walk(job *types.Job, pll *types.Poll, encodeMap bool) *types.Map {
	if job.Timeout != 0 {
		this.agent.Timeout = time.Second * time.Duration(job.Timeout)
		defer func() { this.agent.Timeout = time.Second * time.Duration(this.config.Timeout) }()
	}
	pdus, e := this.agent.WalkAll(pll.What)
	if e != nil {
		job.Error = strings2.New("SNMP Error Walk Host:", this.config.Addr, "/",
			strconv.Itoa(int(this.config.Port)), " Oid:", pll.What, e.Error()).String()
		return nil
	}
	m := &types.Map{}
	m.Data = make(map[string][]byte)
	for _, pdu := range pdus {
		enc := object.New([]byte{}, 0, "SnmpPDU", this.resources.Registry())
		err := enc.Add(pdu.Value)
		if err != nil {
			this.resources.Logger().Error("Object Value Error: ", err.Error())
		}
		m.Data[pdu.Name] = enc.Data()
	}
	if encodeMap {
		enc := object.New([]byte{}, 0, "Map", this.resources.Registry())
		err := enc.Add(m)
		if err != nil {
			this.resources.Logger().Error("Object Table Error: ", err)
		}
		job.Result = enc.Data()
	}
	return m
}

func (this *SNMPCollector) table(job *types.Job, pll *types.Poll) {
	m := this.walk(job, pll, false)
	if job.Error != "" {
		return
	}
	tbl := &types.Table{}
	lastRowIndex := -1
	keys := protocols.Keys(m)
	for _, key := range keys {
		rowIndex, colIndex := getRowAndCol(key)
		if rowIndex > lastRowIndex {
			lastRowIndex = rowIndex
		}
		protocols.SetValue(rowIndex, colIndex, m.Data[key], tbl)
	}

	enc := object.New([]byte{}, 0, "Table", this.resources.Registry())
	err := enc.Add(tbl)
	if err != nil {
		this.resources.Logger().Error("Object Table Error: ", err)
		return
	}
	job.Result = enc.Data()
}
