// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package index

import (
	"sync"
)

// Ensure, that GCPIndexFetchMock does implement GCPIndexFetch.
// If this is not the case, regenerate this file with moq.
var _ GCPIndexFetch = &GCPIndexFetchMock{}

// GCPIndexFetchMock is a mock implementation of GCPIndexFetch.
//
// 	func TestSomethingThatUsesGCPIndexFetch(t *testing.T) {
//
// 		// make and configure a mocked GCPIndexFetch
// 		mockedGCPIndexFetch := &GCPIndexFetchMock{
// 			GetImagesFunc: func() (GCPImages, error) {
// 				panic("mock out the GetImages method")
// 			},
// 		}
//
// 		// use mockedGCPIndexFetch in code that requires GCPIndexFetch
// 		// and then make assertions.
//
// 	}
type GCPIndexFetchMock struct {
	// GetImagesFunc mocks the GetImages method.
	GetImagesFunc func() (GCPImages, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetImages holds details about calls to the GetImages method.
		GetImages []struct {
		}
	}
	lockGetImages sync.RWMutex
}

// GetImages calls GetImagesFunc.
func (mock *GCPIndexFetchMock) GetImages() (GCPImages, error) {
	if mock.GetImagesFunc == nil {
		panic("GCPIndexFetchMock.GetImagesFunc: method is nil but GCPIndexFetch.GetImages was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetImages.Lock()
	mock.calls.GetImages = append(mock.calls.GetImages, callInfo)
	mock.lockGetImages.Unlock()
	return mock.GetImagesFunc()
}

// GetImagesCalls gets all the calls that were made to GetImages.
// Check the length with:
//     len(mockedGCPIndexFetch.GetImagesCalls())
func (mock *GCPIndexFetchMock) GetImagesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetImages.RLock()
	calls = mock.calls.GetImages
	mock.lockGetImages.RUnlock()
	return calls
}

// Ensure, that GCPIndexMock does implement GCPIndex.
// If this is not the case, regenerate this file with moq.
var _ GCPIndex = &GCPIndexMock{}

// GCPIndexMock is a mock implementation of GCPIndex.
//
// 	func TestSomethingThatUsesGCPIndex(t *testing.T) {
//
// 		// make and configure a mocked GCPIndex
// 		mockedGCPIndex := &GCPIndexMock{
// 			MatchFunc: func(version string) (string, error) {
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
	MatchFunc func(version string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Match holds details about calls to the Match method.
		Match []struct {
			// Version is the version argument value.
			Version string
		}
	}
	lockMatch sync.RWMutex
}

// Match calls MatchFunc.
func (mock *GCPIndexMock) Match(version string) (string, error) {
	if mock.MatchFunc == nil {
		panic("GCPIndexMock.MatchFunc: method is nil but GCPIndex.Match was just called")
	}
	callInfo := struct {
		Version string
	}{
		Version: version,
	}
	mock.lockMatch.Lock()
	mock.calls.Match = append(mock.calls.Match, callInfo)
	mock.lockMatch.Unlock()
	return mock.MatchFunc(version)
}

// MatchCalls gets all the calls that were made to Match.
// Check the length with:
//     len(mockedGCPIndex.MatchCalls())
func (mock *GCPIndexMock) MatchCalls() []struct {
	Version string
} {
	var calls []struct {
		Version string
	}
	mock.lockMatch.RLock()
	calls = mock.calls.Match
	mock.lockMatch.RUnlock()
	return calls
}
