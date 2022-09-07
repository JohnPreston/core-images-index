// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package index

import (
	"context"
	"sync"
)

// Ensure, that GCPIndexMock does implement GCPIndex.
// If this is not the case, regenerate this file with moq.
var _ GCPIndex = &GCPIndexMock{}

// GCPIndexMock is a mock implementation of GCPIndex.
//
// 	func TestSomethingThatUsesGCPIndex(t *testing.T) {
//
// 		// make and configure a mocked GCPIndex
// 		mockedGCPIndex := &GCPIndexMock{
// 			MatchFunc: func(ctx context.Context, opts FilterOpts) (string, error) {
// 				panic("mock out the Match method")
// 			},
// 		}
//
// 		// use mockedGCPIndex in code that requires GCPIndex
// 		// and then make assertions.
//
// 	}
type GCPIndexMock struct {
	// MatchFunc mocks the Match method.
	MatchFunc func(ctx context.Context, opts FilterOpts) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Match holds details about calls to the Match method.
		Match []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts FilterOpts
		}
	}
	lockMatch sync.RWMutex
}

// Match calls MatchFunc.
func (mock *GCPIndexMock) Match(ctx context.Context, opts FilterOpts) (string, error) {
	if mock.MatchFunc == nil {
		panic("GCPIndexMock.MatchFunc: method is nil but GCPIndex.Match was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts FilterOpts
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockMatch.Lock()
	mock.calls.Match = append(mock.calls.Match, callInfo)
	mock.lockMatch.Unlock()
	return mock.MatchFunc(ctx, opts)
}

// MatchCalls gets all the calls that were made to Match.
// Check the length with:
//     len(mockedGCPIndex.MatchCalls())
func (mock *GCPIndexMock) MatchCalls() []struct {
	Ctx  context.Context
	Opts FilterOpts
} {
	var calls []struct {
		Ctx  context.Context
		Opts FilterOpts
	}
	mock.lockMatch.RLock()
	calls = mock.calls.Match
	mock.lockMatch.RUnlock()
	return calls
}