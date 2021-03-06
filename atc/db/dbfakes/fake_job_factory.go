// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	sync "sync"

	db "github.com/concourse/concourse/atc/db"
)

type FakeJobFactory struct {
	VisibleJobsStub        func([]string) (db.Dashboard, error)
	visibleJobsMutex       sync.RWMutex
	visibleJobsArgsForCall []struct {
		arg1 []string
	}
	visibleJobsReturns struct {
		result1 db.Dashboard
		result2 error
	}
	visibleJobsReturnsOnCall map[int]struct {
		result1 db.Dashboard
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeJobFactory) VisibleJobs(arg1 []string) (db.Dashboard, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.visibleJobsMutex.Lock()
	ret, specificReturn := fake.visibleJobsReturnsOnCall[len(fake.visibleJobsArgsForCall)]
	fake.visibleJobsArgsForCall = append(fake.visibleJobsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("VisibleJobs", []interface{}{arg1Copy})
	fake.visibleJobsMutex.Unlock()
	if fake.VisibleJobsStub != nil {
		return fake.VisibleJobsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.visibleJobsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeJobFactory) VisibleJobsCallCount() int {
	fake.visibleJobsMutex.RLock()
	defer fake.visibleJobsMutex.RUnlock()
	return len(fake.visibleJobsArgsForCall)
}

func (fake *FakeJobFactory) VisibleJobsCalls(stub func([]string) (db.Dashboard, error)) {
	fake.visibleJobsMutex.Lock()
	defer fake.visibleJobsMutex.Unlock()
	fake.VisibleJobsStub = stub
}

func (fake *FakeJobFactory) VisibleJobsArgsForCall(i int) []string {
	fake.visibleJobsMutex.RLock()
	defer fake.visibleJobsMutex.RUnlock()
	argsForCall := fake.visibleJobsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeJobFactory) VisibleJobsReturns(result1 db.Dashboard, result2 error) {
	fake.visibleJobsMutex.Lock()
	defer fake.visibleJobsMutex.Unlock()
	fake.VisibleJobsStub = nil
	fake.visibleJobsReturns = struct {
		result1 db.Dashboard
		result2 error
	}{result1, result2}
}

func (fake *FakeJobFactory) VisibleJobsReturnsOnCall(i int, result1 db.Dashboard, result2 error) {
	fake.visibleJobsMutex.Lock()
	defer fake.visibleJobsMutex.Unlock()
	fake.VisibleJobsStub = nil
	if fake.visibleJobsReturnsOnCall == nil {
		fake.visibleJobsReturnsOnCall = make(map[int]struct {
			result1 db.Dashboard
			result2 error
		})
	}
	fake.visibleJobsReturnsOnCall[i] = struct {
		result1 db.Dashboard
		result2 error
	}{result1, result2}
}

func (fake *FakeJobFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.visibleJobsMutex.RLock()
	defer fake.visibleJobsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeJobFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.JobFactory = new(FakeJobFactory)
