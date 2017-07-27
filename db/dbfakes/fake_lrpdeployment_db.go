// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeLRPDeploymentDB struct {
	CreateLRPDeploymentStub        func(logger lager.Logger, lrp *models.LRPDeploymentCreation) (string, error)
	createLRPDeploymentMutex       sync.RWMutex
	createLRPDeploymentArgsForCall []struct {
		logger lager.Logger
		lrp    *models.LRPDeploymentCreation
	}
	createLRPDeploymentReturns struct {
		result1 string
		result2 error
	}
	createLRPDeploymentReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UpdateLRPDeploymentStub        func(logger lager.Logger, id string, definition *models.LRPDeploymentUpdate) (string, error)
	updateLRPDeploymentMutex       sync.RWMutex
	updateLRPDeploymentArgsForCall []struct {
		logger     lager.Logger
		id         string
		definition *models.LRPDeploymentUpdate
	}
	updateLRPDeploymentReturns struct {
		result1 string
		result2 error
	}
	updateLRPDeploymentReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	SaveLRPDeploymentStub        func(logger lager.Logger, lrpDeployment *models.LRPDeployment) error
	saveLRPDeploymentMutex       sync.RWMutex
	saveLRPDeploymentArgsForCall []struct {
		logger        lager.Logger
		lrpDeployment *models.LRPDeployment
	}
	saveLRPDeploymentReturns struct {
		result1 error
	}
	saveLRPDeploymentReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteLRPDeploymentStub        func(logger lager.Logger, id string) error
	deleteLRPDeploymentMutex       sync.RWMutex
	deleteLRPDeploymentArgsForCall []struct {
		logger lager.Logger
		id     string
	}
	deleteLRPDeploymentReturns struct {
		result1 error
	}
	deleteLRPDeploymentReturnsOnCall map[int]struct {
		result1 error
	}
	ActivateLRPDeploymentDefinitionStub        func(logger lager.Logger, id string, definitionID string) error
	activateLRPDeploymentDefinitionMutex       sync.RWMutex
	activateLRPDeploymentDefinitionArgsForCall []struct {
		logger       lager.Logger
		id           string
		definitionID string
	}
	activateLRPDeploymentDefinitionReturns struct {
		result1 error
	}
	activateLRPDeploymentDefinitionReturnsOnCall map[int]struct {
		result1 error
	}
	LRPDeploymentByDefinitionGuidStub        func(logger lager.Logger, id string) (*models.LRPDeployment, error)
	lRPDeploymentByDefinitionGuidMutex       sync.RWMutex
	lRPDeploymentByDefinitionGuidArgsForCall []struct {
		logger lager.Logger
		id     string
	}
	lRPDeploymentByDefinitionGuidReturns struct {
		result1 *models.LRPDeployment
		result2 error
	}
	lRPDeploymentByDefinitionGuidReturnsOnCall map[int]struct {
		result1 *models.LRPDeployment
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLRPDeploymentDB) CreateLRPDeployment(logger lager.Logger, lrp *models.LRPDeploymentCreation) (string, error) {
	fake.createLRPDeploymentMutex.Lock()
	ret, specificReturn := fake.createLRPDeploymentReturnsOnCall[len(fake.createLRPDeploymentArgsForCall)]
	fake.createLRPDeploymentArgsForCall = append(fake.createLRPDeploymentArgsForCall, struct {
		logger lager.Logger
		lrp    *models.LRPDeploymentCreation
	}{logger, lrp})
	fake.recordInvocation("CreateLRPDeployment", []interface{}{logger, lrp})
	fake.createLRPDeploymentMutex.Unlock()
	if fake.CreateLRPDeploymentStub != nil {
		return fake.CreateLRPDeploymentStub(logger, lrp)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createLRPDeploymentReturns.result1, fake.createLRPDeploymentReturns.result2
}

func (fake *FakeLRPDeploymentDB) CreateLRPDeploymentCallCount() int {
	fake.createLRPDeploymentMutex.RLock()
	defer fake.createLRPDeploymentMutex.RUnlock()
	return len(fake.createLRPDeploymentArgsForCall)
}

func (fake *FakeLRPDeploymentDB) CreateLRPDeploymentArgsForCall(i int) (lager.Logger, *models.LRPDeploymentCreation) {
	fake.createLRPDeploymentMutex.RLock()
	defer fake.createLRPDeploymentMutex.RUnlock()
	return fake.createLRPDeploymentArgsForCall[i].logger, fake.createLRPDeploymentArgsForCall[i].lrp
}

func (fake *FakeLRPDeploymentDB) CreateLRPDeploymentReturns(result1 string, result2 error) {
	fake.CreateLRPDeploymentStub = nil
	fake.createLRPDeploymentReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDeploymentDB) CreateLRPDeploymentReturnsOnCall(i int, result1 string, result2 error) {
	fake.CreateLRPDeploymentStub = nil
	if fake.createLRPDeploymentReturnsOnCall == nil {
		fake.createLRPDeploymentReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createLRPDeploymentReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDeploymentDB) UpdateLRPDeployment(logger lager.Logger, id string, definition *models.LRPDeploymentUpdate) (string, error) {
	fake.updateLRPDeploymentMutex.Lock()
	ret, specificReturn := fake.updateLRPDeploymentReturnsOnCall[len(fake.updateLRPDeploymentArgsForCall)]
	fake.updateLRPDeploymentArgsForCall = append(fake.updateLRPDeploymentArgsForCall, struct {
		logger     lager.Logger
		id         string
		definition *models.LRPDeploymentUpdate
	}{logger, id, definition})
	fake.recordInvocation("UpdateLRPDeployment", []interface{}{logger, id, definition})
	fake.updateLRPDeploymentMutex.Unlock()
	if fake.UpdateLRPDeploymentStub != nil {
		return fake.UpdateLRPDeploymentStub(logger, id, definition)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateLRPDeploymentReturns.result1, fake.updateLRPDeploymentReturns.result2
}

func (fake *FakeLRPDeploymentDB) UpdateLRPDeploymentCallCount() int {
	fake.updateLRPDeploymentMutex.RLock()
	defer fake.updateLRPDeploymentMutex.RUnlock()
	return len(fake.updateLRPDeploymentArgsForCall)
}

func (fake *FakeLRPDeploymentDB) UpdateLRPDeploymentArgsForCall(i int) (lager.Logger, string, *models.LRPDeploymentUpdate) {
	fake.updateLRPDeploymentMutex.RLock()
	defer fake.updateLRPDeploymentMutex.RUnlock()
	return fake.updateLRPDeploymentArgsForCall[i].logger, fake.updateLRPDeploymentArgsForCall[i].id, fake.updateLRPDeploymentArgsForCall[i].definition
}

func (fake *FakeLRPDeploymentDB) UpdateLRPDeploymentReturns(result1 string, result2 error) {
	fake.UpdateLRPDeploymentStub = nil
	fake.updateLRPDeploymentReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDeploymentDB) UpdateLRPDeploymentReturnsOnCall(i int, result1 string, result2 error) {
	fake.UpdateLRPDeploymentStub = nil
	if fake.updateLRPDeploymentReturnsOnCall == nil {
		fake.updateLRPDeploymentReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.updateLRPDeploymentReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDeploymentDB) SaveLRPDeployment(logger lager.Logger, lrpDeployment *models.LRPDeployment) error {
	fake.saveLRPDeploymentMutex.Lock()
	ret, specificReturn := fake.saveLRPDeploymentReturnsOnCall[len(fake.saveLRPDeploymentArgsForCall)]
	fake.saveLRPDeploymentArgsForCall = append(fake.saveLRPDeploymentArgsForCall, struct {
		logger        lager.Logger
		lrpDeployment *models.LRPDeployment
	}{logger, lrpDeployment})
	fake.recordInvocation("SaveLRPDeployment", []interface{}{logger, lrpDeployment})
	fake.saveLRPDeploymentMutex.Unlock()
	if fake.SaveLRPDeploymentStub != nil {
		return fake.SaveLRPDeploymentStub(logger, lrpDeployment)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.saveLRPDeploymentReturns.result1
}

func (fake *FakeLRPDeploymentDB) SaveLRPDeploymentCallCount() int {
	fake.saveLRPDeploymentMutex.RLock()
	defer fake.saveLRPDeploymentMutex.RUnlock()
	return len(fake.saveLRPDeploymentArgsForCall)
}

func (fake *FakeLRPDeploymentDB) SaveLRPDeploymentArgsForCall(i int) (lager.Logger, *models.LRPDeployment) {
	fake.saveLRPDeploymentMutex.RLock()
	defer fake.saveLRPDeploymentMutex.RUnlock()
	return fake.saveLRPDeploymentArgsForCall[i].logger, fake.saveLRPDeploymentArgsForCall[i].lrpDeployment
}

func (fake *FakeLRPDeploymentDB) SaveLRPDeploymentReturns(result1 error) {
	fake.SaveLRPDeploymentStub = nil
	fake.saveLRPDeploymentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDeploymentDB) SaveLRPDeploymentReturnsOnCall(i int, result1 error) {
	fake.SaveLRPDeploymentStub = nil
	if fake.saveLRPDeploymentReturnsOnCall == nil {
		fake.saveLRPDeploymentReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveLRPDeploymentReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDeploymentDB) DeleteLRPDeployment(logger lager.Logger, id string) error {
	fake.deleteLRPDeploymentMutex.Lock()
	ret, specificReturn := fake.deleteLRPDeploymentReturnsOnCall[len(fake.deleteLRPDeploymentArgsForCall)]
	fake.deleteLRPDeploymentArgsForCall = append(fake.deleteLRPDeploymentArgsForCall, struct {
		logger lager.Logger
		id     string
	}{logger, id})
	fake.recordInvocation("DeleteLRPDeployment", []interface{}{logger, id})
	fake.deleteLRPDeploymentMutex.Unlock()
	if fake.DeleteLRPDeploymentStub != nil {
		return fake.DeleteLRPDeploymentStub(logger, id)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteLRPDeploymentReturns.result1
}

func (fake *FakeLRPDeploymentDB) DeleteLRPDeploymentCallCount() int {
	fake.deleteLRPDeploymentMutex.RLock()
	defer fake.deleteLRPDeploymentMutex.RUnlock()
	return len(fake.deleteLRPDeploymentArgsForCall)
}

func (fake *FakeLRPDeploymentDB) DeleteLRPDeploymentArgsForCall(i int) (lager.Logger, string) {
	fake.deleteLRPDeploymentMutex.RLock()
	defer fake.deleteLRPDeploymentMutex.RUnlock()
	return fake.deleteLRPDeploymentArgsForCall[i].logger, fake.deleteLRPDeploymentArgsForCall[i].id
}

func (fake *FakeLRPDeploymentDB) DeleteLRPDeploymentReturns(result1 error) {
	fake.DeleteLRPDeploymentStub = nil
	fake.deleteLRPDeploymentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDeploymentDB) DeleteLRPDeploymentReturnsOnCall(i int, result1 error) {
	fake.DeleteLRPDeploymentStub = nil
	if fake.deleteLRPDeploymentReturnsOnCall == nil {
		fake.deleteLRPDeploymentReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteLRPDeploymentReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDeploymentDB) ActivateLRPDeploymentDefinition(logger lager.Logger, id string, definitionID string) error {
	fake.activateLRPDeploymentDefinitionMutex.Lock()
	ret, specificReturn := fake.activateLRPDeploymentDefinitionReturnsOnCall[len(fake.activateLRPDeploymentDefinitionArgsForCall)]
	fake.activateLRPDeploymentDefinitionArgsForCall = append(fake.activateLRPDeploymentDefinitionArgsForCall, struct {
		logger       lager.Logger
		id           string
		definitionID string
	}{logger, id, definitionID})
	fake.recordInvocation("ActivateLRPDeploymentDefinition", []interface{}{logger, id, definitionID})
	fake.activateLRPDeploymentDefinitionMutex.Unlock()
	if fake.ActivateLRPDeploymentDefinitionStub != nil {
		return fake.ActivateLRPDeploymentDefinitionStub(logger, id, definitionID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.activateLRPDeploymentDefinitionReturns.result1
}

func (fake *FakeLRPDeploymentDB) ActivateLRPDeploymentDefinitionCallCount() int {
	fake.activateLRPDeploymentDefinitionMutex.RLock()
	defer fake.activateLRPDeploymentDefinitionMutex.RUnlock()
	return len(fake.activateLRPDeploymentDefinitionArgsForCall)
}

func (fake *FakeLRPDeploymentDB) ActivateLRPDeploymentDefinitionArgsForCall(i int) (lager.Logger, string, string) {
	fake.activateLRPDeploymentDefinitionMutex.RLock()
	defer fake.activateLRPDeploymentDefinitionMutex.RUnlock()
	return fake.activateLRPDeploymentDefinitionArgsForCall[i].logger, fake.activateLRPDeploymentDefinitionArgsForCall[i].id, fake.activateLRPDeploymentDefinitionArgsForCall[i].definitionID
}

func (fake *FakeLRPDeploymentDB) ActivateLRPDeploymentDefinitionReturns(result1 error) {
	fake.ActivateLRPDeploymentDefinitionStub = nil
	fake.activateLRPDeploymentDefinitionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDeploymentDB) ActivateLRPDeploymentDefinitionReturnsOnCall(i int, result1 error) {
	fake.ActivateLRPDeploymentDefinitionStub = nil
	if fake.activateLRPDeploymentDefinitionReturnsOnCall == nil {
		fake.activateLRPDeploymentDefinitionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.activateLRPDeploymentDefinitionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPDeploymentDB) LRPDeploymentByDefinitionGuid(logger lager.Logger, id string) (*models.LRPDeployment, error) {
	fake.lRPDeploymentByDefinitionGuidMutex.Lock()
	ret, specificReturn := fake.lRPDeploymentByDefinitionGuidReturnsOnCall[len(fake.lRPDeploymentByDefinitionGuidArgsForCall)]
	fake.lRPDeploymentByDefinitionGuidArgsForCall = append(fake.lRPDeploymentByDefinitionGuidArgsForCall, struct {
		logger lager.Logger
		id     string
	}{logger, id})
	fake.recordInvocation("LRPDeploymentByDefinitionGuid", []interface{}{logger, id})
	fake.lRPDeploymentByDefinitionGuidMutex.Unlock()
	if fake.LRPDeploymentByDefinitionGuidStub != nil {
		return fake.LRPDeploymentByDefinitionGuidStub(logger, id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lRPDeploymentByDefinitionGuidReturns.result1, fake.lRPDeploymentByDefinitionGuidReturns.result2
}

func (fake *FakeLRPDeploymentDB) LRPDeploymentByDefinitionGuidCallCount() int {
	fake.lRPDeploymentByDefinitionGuidMutex.RLock()
	defer fake.lRPDeploymentByDefinitionGuidMutex.RUnlock()
	return len(fake.lRPDeploymentByDefinitionGuidArgsForCall)
}

func (fake *FakeLRPDeploymentDB) LRPDeploymentByDefinitionGuidArgsForCall(i int) (lager.Logger, string) {
	fake.lRPDeploymentByDefinitionGuidMutex.RLock()
	defer fake.lRPDeploymentByDefinitionGuidMutex.RUnlock()
	return fake.lRPDeploymentByDefinitionGuidArgsForCall[i].logger, fake.lRPDeploymentByDefinitionGuidArgsForCall[i].id
}

func (fake *FakeLRPDeploymentDB) LRPDeploymentByDefinitionGuidReturns(result1 *models.LRPDeployment, result2 error) {
	fake.LRPDeploymentByDefinitionGuidStub = nil
	fake.lRPDeploymentByDefinitionGuidReturns = struct {
		result1 *models.LRPDeployment
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDeploymentDB) LRPDeploymentByDefinitionGuidReturnsOnCall(i int, result1 *models.LRPDeployment, result2 error) {
	fake.LRPDeploymentByDefinitionGuidStub = nil
	if fake.lRPDeploymentByDefinitionGuidReturnsOnCall == nil {
		fake.lRPDeploymentByDefinitionGuidReturnsOnCall = make(map[int]struct {
			result1 *models.LRPDeployment
			result2 error
		})
	}
	fake.lRPDeploymentByDefinitionGuidReturnsOnCall[i] = struct {
		result1 *models.LRPDeployment
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPDeploymentDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createLRPDeploymentMutex.RLock()
	defer fake.createLRPDeploymentMutex.RUnlock()
	fake.updateLRPDeploymentMutex.RLock()
	defer fake.updateLRPDeploymentMutex.RUnlock()
	fake.saveLRPDeploymentMutex.RLock()
	defer fake.saveLRPDeploymentMutex.RUnlock()
	fake.deleteLRPDeploymentMutex.RLock()
	defer fake.deleteLRPDeploymentMutex.RUnlock()
	fake.activateLRPDeploymentDefinitionMutex.RLock()
	defer fake.activateLRPDeploymentDefinitionMutex.RUnlock()
	fake.lRPDeploymentByDefinitionGuidMutex.RLock()
	defer fake.lRPDeploymentByDefinitionGuidMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLRPDeploymentDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.LRPDeploymentDB = new(FakeLRPDeploymentDB)
