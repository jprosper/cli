// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"github.com/cloudfoundry/cli/cf/actors/service_builder"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeServiceBuilder struct {
	GetServicesForBrokerStub        func(string) ([]models.ServiceOffering, error)
	getServicesForBrokerMutex       sync.RWMutex
	getServicesForBrokerArgsForCall []struct {
		arg1 string
	}
	getServicesForBrokerReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServiceVisibleToOrgStub        func(string, string) ([]models.ServiceOffering, error)
	getServiceVisibleToOrgMutex       sync.RWMutex
	getServiceVisibleToOrgArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceVisibleToOrgReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServicesVisibleToOrgStub        func(string) ([]models.ServiceOffering, error)
	getServicesVisibleToOrgMutex       sync.RWMutex
	getServicesVisibleToOrgArgsForCall []struct {
		arg1 string
	}
	getServicesVisibleToOrgReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	AttachPlansToServiceStub        func(models.ServiceOffering) (models.ServiceOffering, error)
	attachPlansToServiceMutex       sync.RWMutex
	attachPlansToServiceArgsForCall []struct {
		arg1 models.ServiceOffering
	}
	attachPlansToServiceReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
}

func (fake *FakeServiceBuilder) GetServicesForBroker(arg1 string) ([]models.ServiceOffering, error) {
	fake.getServicesForBrokerMutex.Lock()
	defer fake.getServicesForBrokerMutex.Unlock()
	fake.getServicesForBrokerArgsForCall = append(fake.getServicesForBrokerArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.GetServicesForBrokerStub != nil {
		return fake.GetServicesForBrokerStub(arg1)
	} else {
		return fake.getServicesForBrokerReturns.result1, fake.getServicesForBrokerReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesForBrokerCallCount() int {
	fake.getServicesForBrokerMutex.RLock()
	defer fake.getServicesForBrokerMutex.RUnlock()
	return len(fake.getServicesForBrokerArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesForBrokerArgsForCall(i int) string {
	fake.getServicesForBrokerMutex.RLock()
	defer fake.getServicesForBrokerMutex.RUnlock()
	return fake.getServicesForBrokerArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesForBrokerReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetServicesForBrokerStub = nil
	fake.getServicesForBrokerReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrg(arg1 string, arg2 string) ([]models.ServiceOffering, error) {
	fake.getServiceVisibleToOrgMutex.Lock()
	defer fake.getServiceVisibleToOrgMutex.Unlock()
	fake.getServiceVisibleToOrgArgsForCall = append(fake.getServiceVisibleToOrgArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	if fake.GetServiceVisibleToOrgStub != nil {
		return fake.GetServiceVisibleToOrgStub(arg1, arg2)
	} else {
		return fake.getServiceVisibleToOrgReturns.result1, fake.getServiceVisibleToOrgReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgCallCount() int {
	fake.getServiceVisibleToOrgMutex.RLock()
	defer fake.getServiceVisibleToOrgMutex.RUnlock()
	return len(fake.getServiceVisibleToOrgArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgArgsForCall(i int) (string, string) {
	fake.getServiceVisibleToOrgMutex.RLock()
	defer fake.getServiceVisibleToOrgMutex.RUnlock()
	return fake.getServiceVisibleToOrgArgsForCall[i].arg1, fake.getServiceVisibleToOrgArgsForCall[i].arg2
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetServiceVisibleToOrgStub = nil
	fake.getServiceVisibleToOrgReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrg(arg1 string) ([]models.ServiceOffering, error) {
	fake.getServicesVisibleToOrgMutex.Lock()
	defer fake.getServicesVisibleToOrgMutex.Unlock()
	fake.getServicesVisibleToOrgArgsForCall = append(fake.getServicesVisibleToOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.GetServicesVisibleToOrgStub != nil {
		return fake.GetServicesVisibleToOrgStub(arg1)
	} else {
		return fake.getServicesVisibleToOrgReturns.result1, fake.getServicesVisibleToOrgReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgCallCount() int {
	fake.getServicesVisibleToOrgMutex.RLock()
	defer fake.getServicesVisibleToOrgMutex.RUnlock()
	return len(fake.getServicesVisibleToOrgArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgArgsForCall(i int) string {
	fake.getServicesVisibleToOrgMutex.RLock()
	defer fake.getServicesVisibleToOrgMutex.RUnlock()
	return fake.getServicesVisibleToOrgArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgReturns(result1 []models.ServiceOffering, result2 error) {
	fake.GetServicesVisibleToOrgStub = nil
	fake.getServicesVisibleToOrgReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) AttachPlansToService(arg1 models.ServiceOffering) (models.ServiceOffering, error) {
	fake.attachPlansToServiceMutex.Lock()
	defer fake.attachPlansToServiceMutex.Unlock()
	fake.attachPlansToServiceArgsForCall = append(fake.attachPlansToServiceArgsForCall, struct {
		arg1 models.ServiceOffering
	}{arg1})
	if fake.AttachPlansToServiceStub != nil {
		return fake.AttachPlansToServiceStub(arg1)
	} else {
		return fake.attachPlansToServiceReturns.result1, fake.attachPlansToServiceReturns.result2
	}
}

func (fake *FakeServiceBuilder) AttachPlansToServiceCallCount() int {
	fake.attachPlansToServiceMutex.RLock()
	defer fake.attachPlansToServiceMutex.RUnlock()
	return len(fake.attachPlansToServiceArgsForCall)
}

func (fake *FakeServiceBuilder) AttachPlansToServiceArgsForCall(i int) models.ServiceOffering {
	fake.attachPlansToServiceMutex.RLock()
	defer fake.attachPlansToServiceMutex.RUnlock()
	return fake.attachPlansToServiceArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) AttachPlansToServiceReturns(result1 models.ServiceOffering, result2 error) {
	fake.AttachPlansToServiceStub = nil
	fake.attachPlansToServiceReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

var _ service_builder.ServiceBuilder = new(FakeServiceBuilder)
