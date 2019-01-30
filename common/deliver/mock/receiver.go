// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	sync "sync"

	deliver "github.com/oktfolio/hyperledger-fabric-gm/common/deliver"
	common "github.com/oktfolio/hyperledger-fabric-gm/protos/common"
)

type Receiver struct {
	RecvStub        func() (*common.Envelope, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct {
	}
	recvReturns struct {
		result1 *common.Envelope
		result2 error
	}
	recvReturnsOnCall map[int]struct {
		result1 *common.Envelope
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Receiver) Recv() (*common.Envelope, error) {
	fake.recvMutex.Lock()
	ret, specificReturn := fake.recvReturnsOnCall[len(fake.recvArgsForCall)]
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct {
	}{})
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if fake.RecvStub != nil {
		return fake.RecvStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.recvReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Receiver) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *Receiver) RecvCalls(stub func() (*common.Envelope, error)) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = stub
}

func (fake *Receiver) RecvReturns(result1 *common.Envelope, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *common.Envelope
		result2 error
	}{result1, result2}
}

func (fake *Receiver) RecvReturnsOnCall(i int, result1 *common.Envelope, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	if fake.recvReturnsOnCall == nil {
		fake.recvReturnsOnCall = make(map[int]struct {
			result1 *common.Envelope
			result2 error
		})
	}
	fake.recvReturnsOnCall[i] = struct {
		result1 *common.Envelope
		result2 error
	}{result1, result2}
}

func (fake *Receiver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Receiver) recordInvocation(key string, args []interface{}) {
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

var _ deliver.Receiver = new(Receiver)
