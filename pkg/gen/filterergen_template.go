// Package gen contains all the components for code generation.
package gen

const filtererTmpl = `
// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at {{.GenTime}}
package {{.PackageName}}

import (
    log "github.com/golang/glog"
    "fmt"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"

    "github.com/joincivil/civil-events-crawler/pkg/model"
{{if .ContractImportPath -}}
    "{{.ContractImportPath}}"
{{- end}}
{{if .AdditionalImports -}}
{{- range .AdditionalImports}}
    "{{.}}"
{{- end}}
{{- end}}
)

{{if .EventHandlers -}}
var eventTypes{{.ContractTypeName}} = []string{
    {{- range .EventHandlers}}
        "{{.EventMethod}}",
    {{- end}}
}
{{- end}}

func New{{.ContractTypeName}}Filterers(contractAddress common.Address) *{{.ContractTypeName}}Filterers {
    contractFilterers := &{{.ContractTypeName}}Filterers{
        contractAddress: contractAddress,
        eventTypes: eventTypes{{.ContractTypeName}},
        eventToStartBlock: make(map[string]uint64),
        lastEvents: make([]model.CivilEvent, 0),
    }
    for _, eventType := range contractFilterers.eventTypes {
        contractFilterers.eventToStartBlock[eventType] = 0
    }
    return contractFilterers
}

type {{.ContractTypeName}}Filterers struct {
    contractAddress common.Address
    contract *{{.ContractTypePackage}}.{{.ContractTypeName}}
    eventTypes []string
    eventToStartBlock map[string]uint64
    lastEvents  []model.CivilEvent
}

func (f *{{.ContractTypeName}}Filterers) ContractName() string {
    return "{{.ContractTypeName}}"
}

func (f *{{.ContractTypeName}}Filterers) ContractAddress() common.Address {
    return f.contractAddress
}

func (f *{{.ContractTypeName}}Filterers) StartFilterers(client bind.ContractBackend, pastEvents []model.CivilEvent) (error, []model.CivilEvent) {
    return f.Start{{.ContractTypeName}}Filterers(client, pastEvents)
}

func (f *{{.ContractTypeName}}Filterers) EventTypes() []string {
    return f.eventTypes
}

func (f *{{.ContractTypeName}}Filterers) UpdateStartBlock(eventType string, startBlock uint64) {
    f.eventToStartBlock[eventType] = startBlock
}

func (f *{{.ContractTypeName}}Filterers) LastEvents() []model.CivilEvent {
    return f.lastEvents
}

// Start{{.ContractTypeName}}Filterers retrieves events for {{.ContractTypeName}}
func (f *{{.ContractTypeName}}Filterers) Start{{.ContractTypeName}}Filterers(client bind.ContractBackend, pastEvents []model.CivilEvent) (error, []model.CivilEvent) {
    contract, err := {{.ContractTypePackage}}.New{{.ContractTypeName}}(f.contractAddress, client)
    if err != nil {
        log.Errorf("Error initializing Start{{.ContractTypeName}}: err: %v", err)
        return err, pastEvents
    }
    f.contract = contract
    var startBlock uint64
    prevEventsLength := len(pastEvents)


{{if .EventHandlers -}}
{{- range .EventHandlers}}

    startBlock = f.eventToStartBlock["{{.EventMethod}}"]
    err, pastEvents = f.startFilter{{.EventMethod}}(startBlock, pastEvents)
    if err != nil {
        return fmt.Errorf("Error retrieving {{.EventMethod}}: err: %v", err), pastEvents
    }
    if len(pastEvents) > prevEventsLength {
        f.lastEvents = append(f.lastEvents, pastEvents[len(pastEvents) - 1])
        prevEventsLength = len(pastEvents)
    }


{{- end}}
{{- end}}

    return nil, pastEvents
}

{{if .EventHandlers -}}
{{- range .EventHandlers}}

func (f *{{$.ContractTypeName}}Filterers) startFilter{{.EventMethod}}(startBlock uint64, pastEvents []model.CivilEvent) (error, []model.CivilEvent) {
    var opts = &bind.FilterOpts{
        Start: startBlock,
    }

	log.Infof("Filtering events for {{.EventMethod}} for contract %v", f.contractAddress.Hex())
    itr, err := f.contract.Filter{{.EventMethod}}(
        opts,
    {{- if .ParamValues -}}
    {{range .ParamValues}}
        []{{.Type}}{},
    {{- end}}
    {{end}}
    )
    if err != nil {
        log.Errorf("Error getting event {{.EventMethod}}: %v", err)
        return err, pastEvents
    }
    nextEvent := itr.Next()
    for nextEvent {
        civilEvent, err := model.NewCivilEvent("{{.EventMethod}}", f.contractAddress, itr.Event)
        if err != nil {
            log.Errorf("Error creating new civil event: event: %v, err: %v", itr.Event, err)
            continue
        }
        pastEvents = append(pastEvents, *civilEvent)
        nextEvent = itr.Next()
    }
    return nil, pastEvents
}

{{- end}}
{{- end}}
`
