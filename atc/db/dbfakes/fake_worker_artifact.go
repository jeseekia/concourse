// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	sync "sync"
	time "time"

	db "github.com/concourse/concourse/atc/db"
)

type FakeWorkerArtifact struct {
	BuildIDStub        func() int
	buildIDMutex       sync.RWMutex
	buildIDArgsForCall []struct {
	}
	buildIDReturns struct {
		result1 int
	}
	buildIDReturnsOnCall map[int]struct {
		result1 int
	}
	CreatedAtStub        func() time.Time
	createdAtMutex       sync.RWMutex
	createdAtArgsForCall []struct {
	}
	createdAtReturns struct {
		result1 time.Time
	}
	createdAtReturnsOnCall map[int]struct {
		result1 time.Time
	}
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	VolumeStub        func(int) (db.CreatedVolume, bool, error)
	volumeMutex       sync.RWMutex
	volumeArgsForCall []struct {
		arg1 int
	}
	volumeReturns struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}
	volumeReturnsOnCall map[int]struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerArtifact) BuildID() int {
	fake.buildIDMutex.Lock()
	ret, specificReturn := fake.buildIDReturnsOnCall[len(fake.buildIDArgsForCall)]
	fake.buildIDArgsForCall = append(fake.buildIDArgsForCall, struct {
	}{})
	fake.recordInvocation("BuildID", []interface{}{})
	fake.buildIDMutex.Unlock()
	if fake.BuildIDStub != nil {
		return fake.BuildIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.buildIDReturns
	return fakeReturns.result1
}

func (fake *FakeWorkerArtifact) BuildIDCallCount() int {
	fake.buildIDMutex.RLock()
	defer fake.buildIDMutex.RUnlock()
	return len(fake.buildIDArgsForCall)
}

func (fake *FakeWorkerArtifact) BuildIDCalls(stub func() int) {
	fake.buildIDMutex.Lock()
	defer fake.buildIDMutex.Unlock()
	fake.BuildIDStub = stub
}

func (fake *FakeWorkerArtifact) BuildIDReturns(result1 int) {
	fake.buildIDMutex.Lock()
	defer fake.buildIDMutex.Unlock()
	fake.BuildIDStub = nil
	fake.buildIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorkerArtifact) BuildIDReturnsOnCall(i int, result1 int) {
	fake.buildIDMutex.Lock()
	defer fake.buildIDMutex.Unlock()
	fake.BuildIDStub = nil
	if fake.buildIDReturnsOnCall == nil {
		fake.buildIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.buildIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorkerArtifact) CreatedAt() time.Time {
	fake.createdAtMutex.Lock()
	ret, specificReturn := fake.createdAtReturnsOnCall[len(fake.createdAtArgsForCall)]
	fake.createdAtArgsForCall = append(fake.createdAtArgsForCall, struct {
	}{})
	fake.recordInvocation("CreatedAt", []interface{}{})
	fake.createdAtMutex.Unlock()
	if fake.CreatedAtStub != nil {
		return fake.CreatedAtStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createdAtReturns
	return fakeReturns.result1
}

func (fake *FakeWorkerArtifact) CreatedAtCallCount() int {
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	return len(fake.createdAtArgsForCall)
}

func (fake *FakeWorkerArtifact) CreatedAtCalls(stub func() time.Time) {
	fake.createdAtMutex.Lock()
	defer fake.createdAtMutex.Unlock()
	fake.CreatedAtStub = stub
}

func (fake *FakeWorkerArtifact) CreatedAtReturns(result1 time.Time) {
	fake.createdAtMutex.Lock()
	defer fake.createdAtMutex.Unlock()
	fake.CreatedAtStub = nil
	fake.createdAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeWorkerArtifact) CreatedAtReturnsOnCall(i int, result1 time.Time) {
	fake.createdAtMutex.Lock()
	defer fake.createdAtMutex.Unlock()
	fake.CreatedAtStub = nil
	if fake.createdAtReturnsOnCall == nil {
		fake.createdAtReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.createdAtReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeWorkerArtifact) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeWorkerArtifact) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeWorkerArtifact) IDCalls(stub func() int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeWorkerArtifact) IDReturns(result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorkerArtifact) IDReturnsOnCall(i int, result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorkerArtifact) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeWorkerArtifact) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeWorkerArtifact) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeWorkerArtifact) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorkerArtifact) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorkerArtifact) Volume(arg1 int) (db.CreatedVolume, bool, error) {
	fake.volumeMutex.Lock()
	ret, specificReturn := fake.volumeReturnsOnCall[len(fake.volumeArgsForCall)]
	fake.volumeArgsForCall = append(fake.volumeArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Volume", []interface{}{arg1})
	fake.volumeMutex.Unlock()
	if fake.VolumeStub != nil {
		return fake.VolumeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.volumeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeWorkerArtifact) VolumeCallCount() int {
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	return len(fake.volumeArgsForCall)
}

func (fake *FakeWorkerArtifact) VolumeCalls(stub func(int) (db.CreatedVolume, bool, error)) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = stub
}

func (fake *FakeWorkerArtifact) VolumeArgsForCall(i int) int {
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	argsForCall := fake.volumeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorkerArtifact) VolumeReturns(result1 db.CreatedVolume, result2 bool, result3 error) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = nil
	fake.volumeReturns = struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerArtifact) VolumeReturnsOnCall(i int, result1 db.CreatedVolume, result2 bool, result3 error) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = nil
	if fake.volumeReturnsOnCall == nil {
		fake.volumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatedVolume
			result2 bool
			result3 error
		})
	}
	fake.volumeReturnsOnCall[i] = struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerArtifact) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildIDMutex.RLock()
	defer fake.buildIDMutex.RUnlock()
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorkerArtifact) recordInvocation(key string, args []interface{}) {
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

var _ db.WorkerArtifact = new(FakeWorkerArtifact)
