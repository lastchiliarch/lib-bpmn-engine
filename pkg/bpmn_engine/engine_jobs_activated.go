package bpmn_engine

import (
	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine/variable_scope"
	"time"
)

// ActivatedJob is a struct to provide information for registered task handler
type activatedJob struct {
	processInstanceInfo      *ProcessInstanceInfo
	completeHandler          func()
	failHandler              func(reason string)
	key                      int64
	processInstanceKey       int64
	bpmnProcessId            string
	processDefinitionVersion int32
	processDefinitionKey     int64
	elementId                string
	createdAt                time.Time
	scope                    variable_scope.VariableContext
}

// ActivatedJob represents an abstraction for the activated job
// don't forget to call Fail or Complete when your task worker job is complete or not.
type ActivatedJob interface {
	// GetKey the key, a unique identifier for the job
	GetKey() int64


	// GetProcessInstanceKey the job's process instance key
	GetProcessInstanceKey() int64

	// GetBpmnProcessId Retrieve id of the job process definition
	GetBpmnProcessId() string

	// GetProcessDefinitionVersion Retrieve version of the job process definition
	GetProcessDefinitionVersion() int32

	// GetProcessDefinitionKey Retrieve key of the job process definition
	GetProcessDefinitionKey() int64

	// GetElementId Get element id of the job
	GetElementId() string

	// GetVariable from the process instance's variable context
	GetVariable(key string) interface{}

	// SetVariable in the variables context of the given process instance
	SetVariable(key string, value interface{})

	// GetCreatedAt when the job was created
	GetCreatedAt() time.Time

	// Fail does set the state the worker missed completing the job
	// Fail and Complete mutual exclude each other
	Fail(reason string)

	// Complete does set the state the worker successfully completing the job
	// Fail and Complete mutual exclude each other
	Complete()
}

// GetCreatedAt implements ActivatedJob
func (aj *activatedJob) GetCreatedAt() time.Time {
	return aj.createdAt
}

// GetElementId implements ActivatedJob
func (aj *activatedJob) GetElementId() string {
	return aj.elementId
}

// GetKey implements ActivatedJob
func (aj *activatedJob) GetKey() int64 {
	return aj.key
}

// GetBpmnProcessId implements ActivatedJob
func (aj *activatedJob) GetBpmnProcessId() string {
	return aj.bpmnProcessId
}

// GetProcessDefinitionKey implements ActivatedJob
func (aj *activatedJob) GetProcessDefinitionKey() int64 {
	return aj.processDefinitionKey
}

// GetProcessDefinitionVersion implements ActivatedJob
func (aj *activatedJob) GetProcessDefinitionVersion() int32 {
	return aj.processDefinitionVersion
}

// GetProcessInstanceKey implements ActivatedJob
func (aj *activatedJob) GetProcessInstanceKey() int64 {
	return aj.processInstanceKey
}

// GetVariable implements ActivatedJob
func (aj *activatedJob) GetVariable(key string) interface{} {
	return aj.scope.GetVariable(key)
}

// SetVariable implements ActivatedJob
func (aj *activatedJob) SetVariable(key string, value interface{}) {
	aj.scope.SetVariable(key, value)
}

// Fail implements ActivatedJob
func (aj *activatedJob) Fail(reason string) {
	aj.failHandler(reason)
}

// Complete implements ActivatedJob
func (aj *activatedJob) Complete() {
	aj.completeHandler()
	aj.scope.Propagation()
}
