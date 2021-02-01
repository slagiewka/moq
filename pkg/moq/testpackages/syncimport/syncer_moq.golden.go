// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package syncimport

import (
	"github.com/matryer/moq/pkg/moq/testpackages/syncimport/sync"
	stdsync "sync"
)

// Ensure, that SyncerMock does implement Syncer.
// If this is not the case, regenerate this file with moq.
var _ Syncer = &SyncerMock{}

// SyncerMock is a mock implementation of Syncer.
//
// 	func TestSomethingThatUsesSyncer(t *testing.T) {
//
// 		// make and configure a mocked Syncer
// 		mockedSyncer := &SyncerMock{
// 			BlahFunc: func(s sync.Thing, wg *stdsync.WaitGroup)  {
// 				panic("mock out the Blah method")
// 			},
// 		}
//
// 		// use mockedSyncer in code that requires Syncer
// 		// and then make assertions.
//
// 	}
type SyncerMock struct {
	// BlahFunc mocks the Blah method.
	BlahFunc func(s sync.Thing, wg *stdsync.WaitGroup)

	// calls tracks calls to the methods.
	calls struct {
		// Blah holds details about calls to the Blah method.
		Blah []struct {
			// S is the s argument value.
			S sync.Thing
			// Wg is the wg argument value.
			Wg *stdsync.WaitGroup
		}
	}
	lockBlah stdsync.RWMutex
}

// Blah calls BlahFunc.
func (mock *SyncerMock) Blah(s sync.Thing, wg *stdsync.WaitGroup) {
	if mock.BlahFunc == nil {
		panic("SyncerMock.BlahFunc: method is nil but Syncer.Blah was just called")
	}
	callInfo := struct {
		S  sync.Thing
		Wg *stdsync.WaitGroup
	}{
		S:  s,
		Wg: wg,
	}
	mock.lockBlah.Lock()
	mock.calls.Blah = append(mock.calls.Blah, callInfo)
	mock.lockBlah.Unlock()
	mock.BlahFunc(s, wg)
}

// BlahCalls gets all the calls that were made to Blah.
// Check the length with:
//     len(mockedSyncer.BlahCalls())
func (mock *SyncerMock) BlahCalls() []struct {
	S  sync.Thing
	Wg *stdsync.WaitGroup
} {
	var calls []struct {
		S  sync.Thing
		Wg *stdsync.WaitGroup
	}
	mock.lockBlah.RLock()
	calls = mock.calls.Blah
	mock.lockBlah.RUnlock()
	return calls
}
