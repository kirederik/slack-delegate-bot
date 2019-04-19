// Code generated by counterfeiter. DO NOT EDIT.
package delegatefakes

import (
	"sync"

	"github.com/dpb587/slack-delegate-bot/logic/delegate"
	"github.com/dpb587/slack-delegate-bot/message"
)

type FakeDelegator struct {
	DelegateStub        func(message.Message) ([]delegate.Delegate, error)
	delegateMutex       sync.RWMutex
	delegateArgsForCall []struct {
		arg1 message.Message
	}
	delegateReturns struct {
		result1 []delegate.Delegate
		result2 error
	}
	delegateReturnsOnCall map[int]struct {
		result1 []delegate.Delegate
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDelegator) Delegate(arg1 message.Message) ([]delegate.Delegate, error) {
	fake.delegateMutex.Lock()
	ret, specificReturn := fake.delegateReturnsOnCall[len(fake.delegateArgsForCall)]
	fake.delegateArgsForCall = append(fake.delegateArgsForCall, struct {
		arg1 message.Message
	}{arg1})
	fake.recordInvocation("Delegate", []interface{}{arg1})
	fake.delegateMutex.Unlock()
	if fake.DelegateStub != nil {
		return fake.DelegateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.delegateReturns.result1, fake.delegateReturns.result2
}

func (fake *FakeDelegator) DelegateCallCount() int {
	fake.delegateMutex.RLock()
	defer fake.delegateMutex.RUnlock()
	return len(fake.delegateArgsForCall)
}

func (fake *FakeDelegator) DelegateArgsForCall(i int) message.Message {
	fake.delegateMutex.RLock()
	defer fake.delegateMutex.RUnlock()
	return fake.delegateArgsForCall[i].arg1
}

func (fake *FakeDelegator) DelegateReturns(result1 []delegate.Delegate, result2 error) {
	fake.DelegateStub = nil
	fake.delegateReturns = struct {
		result1 []delegate.Delegate
		result2 error
	}{result1, result2}
}

func (fake *FakeDelegator) DelegateReturnsOnCall(i int, result1 []delegate.Delegate, result2 error) {
	fake.DelegateStub = nil
	if fake.delegateReturnsOnCall == nil {
		fake.delegateReturnsOnCall = make(map[int]struct {
			result1 []delegate.Delegate
			result2 error
		})
	}
	fake.delegateReturnsOnCall[i] = struct {
		result1 []delegate.Delegate
		result2 error
	}{result1, result2}
}

func (fake *FakeDelegator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.delegateMutex.RLock()
	defer fake.delegateMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDelegator) recordInvocation(key string, args []interface{}) {
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

var _ delegate.Delegator = new(FakeDelegator)
