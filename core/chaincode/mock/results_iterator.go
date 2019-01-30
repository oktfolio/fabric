// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	sync "sync"

	ledger "github.com/oktfolio/hyperledger-fabric-gm/common/ledger"
)

type QueryResultsIterator struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	GetBookmarkAndCloseStub        func() string
	getBookmarkAndCloseMutex       sync.RWMutex
	getBookmarkAndCloseArgsForCall []struct {
	}
	getBookmarkAndCloseReturns struct {
		result1 string
	}
	getBookmarkAndCloseReturnsOnCall map[int]struct {
		result1 string
	}
	NextStub        func() (ledger.QueryResult, error)
	nextMutex       sync.RWMutex
	nextArgsForCall []struct {
	}
	nextReturns struct {
		result1 ledger.QueryResult
		result2 error
	}
	nextReturnsOnCall map[int]struct {
		result1 ledger.QueryResult
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *QueryResultsIterator) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *QueryResultsIterator) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *QueryResultsIterator) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *QueryResultsIterator) GetBookmarkAndClose() string {
	fake.getBookmarkAndCloseMutex.Lock()
	ret, specificReturn := fake.getBookmarkAndCloseReturnsOnCall[len(fake.getBookmarkAndCloseArgsForCall)]
	fake.getBookmarkAndCloseArgsForCall = append(fake.getBookmarkAndCloseArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBookmarkAndClose", []interface{}{})
	fake.getBookmarkAndCloseMutex.Unlock()
	if fake.GetBookmarkAndCloseStub != nil {
		return fake.GetBookmarkAndCloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getBookmarkAndCloseReturns
	return fakeReturns.result1
}

func (fake *QueryResultsIterator) GetBookmarkAndCloseCallCount() int {
	fake.getBookmarkAndCloseMutex.RLock()
	defer fake.getBookmarkAndCloseMutex.RUnlock()
	return len(fake.getBookmarkAndCloseArgsForCall)
}

func (fake *QueryResultsIterator) GetBookmarkAndCloseCalls(stub func() string) {
	fake.getBookmarkAndCloseMutex.Lock()
	defer fake.getBookmarkAndCloseMutex.Unlock()
	fake.GetBookmarkAndCloseStub = stub
}

func (fake *QueryResultsIterator) GetBookmarkAndCloseReturns(result1 string) {
	fake.getBookmarkAndCloseMutex.Lock()
	defer fake.getBookmarkAndCloseMutex.Unlock()
	fake.GetBookmarkAndCloseStub = nil
	fake.getBookmarkAndCloseReturns = struct {
		result1 string
	}{result1}
}

func (fake *QueryResultsIterator) GetBookmarkAndCloseReturnsOnCall(i int, result1 string) {
	fake.getBookmarkAndCloseMutex.Lock()
	defer fake.getBookmarkAndCloseMutex.Unlock()
	fake.GetBookmarkAndCloseStub = nil
	if fake.getBookmarkAndCloseReturnsOnCall == nil {
		fake.getBookmarkAndCloseReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getBookmarkAndCloseReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *QueryResultsIterator) Next() (ledger.QueryResult, error) {
	fake.nextMutex.Lock()
	ret, specificReturn := fake.nextReturnsOnCall[len(fake.nextArgsForCall)]
	fake.nextArgsForCall = append(fake.nextArgsForCall, struct {
	}{})
	fake.recordInvocation("Next", []interface{}{})
	fake.nextMutex.Unlock()
	if fake.NextStub != nil {
		return fake.NextStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.nextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *QueryResultsIterator) NextCallCount() int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	return len(fake.nextArgsForCall)
}

func (fake *QueryResultsIterator) NextCalls(stub func() (ledger.QueryResult, error)) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = stub
}

func (fake *QueryResultsIterator) NextReturns(result1 ledger.QueryResult, result2 error) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	fake.nextReturns = struct {
		result1 ledger.QueryResult
		result2 error
	}{result1, result2}
}

func (fake *QueryResultsIterator) NextReturnsOnCall(i int, result1 ledger.QueryResult, result2 error) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	if fake.nextReturnsOnCall == nil {
		fake.nextReturnsOnCall = make(map[int]struct {
			result1 ledger.QueryResult
			result2 error
		})
	}
	fake.nextReturnsOnCall[i] = struct {
		result1 ledger.QueryResult
		result2 error
	}{result1, result2}
}

func (fake *QueryResultsIterator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.getBookmarkAndCloseMutex.RLock()
	defer fake.getBookmarkAndCloseMutex.RUnlock()
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *QueryResultsIterator) recordInvocation(key string, args []interface{}) {
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
