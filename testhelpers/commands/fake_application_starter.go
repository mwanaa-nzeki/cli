// This file was generated by counterfeiter
package commands

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/commands/application"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/flags"
)

type FakeApplicationStarter struct {
	MetaDataStub        func() command_registry.CommandMetadata
	metaDataMutex       sync.RWMutex
	metaDataArgsForCall []struct{}
	metaDataReturns     struct {
		result1 command_registry.CommandMetadata
	}
	SetDependencyStub        func(deps command_registry.Dependency, pluginCall bool) command_registry.Command
	setDependencyMutex       sync.RWMutex
	setDependencyArgsForCall []struct {
		deps       command_registry.Dependency
		pluginCall bool
	}
	setDependencyReturns struct {
		result1 command_registry.Command
	}
	RequirementsStub        func(requirementsFactory requirements.Factory, context flags.FlagContext) (reqs []requirements.Requirement, err error)
	requirementsMutex       sync.RWMutex
	requirementsArgsForCall []struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}
	requirementsReturns struct {
		result1 []requirements.Requirement
		result2 error
	}
	ExecuteStub        func(context flags.FlagContext)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		context flags.FlagContext
	}
	SetStartTimeoutInSecondsStub        func(timeout int)
	setStartTimeoutInSecondsMutex       sync.RWMutex
	setStartTimeoutInSecondsArgsForCall []struct {
		timeout int
	}
	ApplicationStartStub        func(app models.Application, orgName string, spaceName string) (updatedApp models.Application, err error)
	applicationStartMutex       sync.RWMutex
	applicationStartArgsForCall []struct {
		app       models.Application
		orgName   string
		spaceName string
	}
	applicationStartReturns struct {
		result1 models.Application
		result2 error
	}
}

func (fake *FakeApplicationStarter) MetaData() command_registry.CommandMetadata {
	fake.metaDataMutex.Lock()
	fake.metaDataArgsForCall = append(fake.metaDataArgsForCall, struct{}{})
	fake.metaDataMutex.Unlock()
	if fake.MetaDataStub != nil {
		return fake.MetaDataStub()
	} else {
		return fake.metaDataReturns.result1
	}
}

func (fake *FakeApplicationStarter) MetaDataCallCount() int {
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	return len(fake.metaDataArgsForCall)
}

func (fake *FakeApplicationStarter) MetaDataReturns(result1 command_registry.CommandMetadata) {
	fake.MetaDataStub = nil
	fake.metaDataReturns = struct {
		result1 command_registry.CommandMetadata
	}{result1}
}

func (fake *FakeApplicationStarter) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	fake.setDependencyMutex.Lock()
	fake.setDependencyArgsForCall = append(fake.setDependencyArgsForCall, struct {
		deps       command_registry.Dependency
		pluginCall bool
	}{deps, pluginCall})
	fake.setDependencyMutex.Unlock()
	if fake.SetDependencyStub != nil {
		return fake.SetDependencyStub(deps, pluginCall)
	} else {
		return fake.setDependencyReturns.result1
	}
}

func (fake *FakeApplicationStarter) SetDependencyCallCount() int {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return len(fake.setDependencyArgsForCall)
}

func (fake *FakeApplicationStarter) SetDependencyArgsForCall(i int) (command_registry.Dependency, bool) {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return fake.setDependencyArgsForCall[i].deps, fake.setDependencyArgsForCall[i].pluginCall
}

func (fake *FakeApplicationStarter) SetDependencyReturns(result1 command_registry.Command) {
	fake.SetDependencyStub = nil
	fake.setDependencyReturns = struct {
		result1 command_registry.Command
	}{result1}
}

func (fake *FakeApplicationStarter) Requirements(requirementsFactory requirements.Factory, context flags.FlagContext) (reqs []requirements.Requirement, err error) {
	fake.requirementsMutex.Lock()
	fake.requirementsArgsForCall = append(fake.requirementsArgsForCall, struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}{requirementsFactory, context})
	fake.requirementsMutex.Unlock()
	if fake.RequirementsStub != nil {
		return fake.RequirementsStub(requirementsFactory, context)
	} else {
		return fake.requirementsReturns.result1, fake.requirementsReturns.result2
	}
}

func (fake *FakeApplicationStarter) RequirementsCallCount() int {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return len(fake.requirementsArgsForCall)
}

func (fake *FakeApplicationStarter) RequirementsArgsForCall(i int) (requirements.Factory, flags.FlagContext) {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return fake.requirementsArgsForCall[i].requirementsFactory, fake.requirementsArgsForCall[i].context
}

func (fake *FakeApplicationStarter) RequirementsReturns(result1 []requirements.Requirement, result2 error) {
	fake.RequirementsStub = nil
	fake.requirementsReturns = struct {
		result1 []requirements.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeApplicationStarter) Execute(context flags.FlagContext) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		context flags.FlagContext
	}{context})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		fake.ExecuteStub(context)
	}
}

func (fake *FakeApplicationStarter) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeApplicationStarter) ExecuteArgsForCall(i int) flags.FlagContext {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].context
}

func (fake *FakeApplicationStarter) SetStartTimeoutInSeconds(timeout int) {
	fake.setStartTimeoutInSecondsMutex.Lock()
	fake.setStartTimeoutInSecondsArgsForCall = append(fake.setStartTimeoutInSecondsArgsForCall, struct {
		timeout int
	}{timeout})
	fake.setStartTimeoutInSecondsMutex.Unlock()
	if fake.SetStartTimeoutInSecondsStub != nil {
		fake.SetStartTimeoutInSecondsStub(timeout)
	}
}

func (fake *FakeApplicationStarter) SetStartTimeoutInSecondsCallCount() int {
	fake.setStartTimeoutInSecondsMutex.RLock()
	defer fake.setStartTimeoutInSecondsMutex.RUnlock()
	return len(fake.setStartTimeoutInSecondsArgsForCall)
}

func (fake *FakeApplicationStarter) SetStartTimeoutInSecondsArgsForCall(i int) int {
	fake.setStartTimeoutInSecondsMutex.RLock()
	defer fake.setStartTimeoutInSecondsMutex.RUnlock()
	return fake.setStartTimeoutInSecondsArgsForCall[i].timeout
}

func (fake *FakeApplicationStarter) ApplicationStart(app models.Application, orgName string, spaceName string) (updatedApp models.Application, err error) {
	fake.applicationStartMutex.Lock()
	fake.applicationStartArgsForCall = append(fake.applicationStartArgsForCall, struct {
		app       models.Application
		orgName   string
		spaceName string
	}{app, orgName, spaceName})
	fake.applicationStartMutex.Unlock()
	if fake.ApplicationStartStub != nil {
		return fake.ApplicationStartStub(app, orgName, spaceName)
	} else {
		return fake.applicationStartReturns.result1, fake.applicationStartReturns.result2
	}
}

func (fake *FakeApplicationStarter) ApplicationStartCallCount() int {
	fake.applicationStartMutex.RLock()
	defer fake.applicationStartMutex.RUnlock()
	return len(fake.applicationStartArgsForCall)
}

func (fake *FakeApplicationStarter) ApplicationStartArgsForCall(i int) (models.Application, string, string) {
	fake.applicationStartMutex.RLock()
	defer fake.applicationStartMutex.RUnlock()
	return fake.applicationStartArgsForCall[i].app, fake.applicationStartArgsForCall[i].orgName, fake.applicationStartArgsForCall[i].spaceName
}

func (fake *FakeApplicationStarter) ApplicationStartReturns(result1 models.Application, result2 error) {
	fake.ApplicationStartStub = nil
	fake.applicationStartReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

var _ application.ApplicationStarter = new(FakeApplicationStarter)
