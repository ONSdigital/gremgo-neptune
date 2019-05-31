// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package gremgo

import (
	"context"
	"sync"
)

var (
	lockdialerMockclose       sync.RWMutex
	lockdialerMockconnect     sync.RWMutex
	lockdialerMockconnectCtx  sync.RWMutex
	lockdialerMockgetAuth     sync.RWMutex
	lockdialerMockisConnected sync.RWMutex
	lockdialerMockisDisposed  sync.RWMutex
	lockdialerMockping        sync.RWMutex
	lockdialerMockpingCtx     sync.RWMutex
	lockdialerMockread        sync.RWMutex
	lockdialerMockreadCtx     sync.RWMutex
	lockdialerMockwrite       sync.RWMutex
)

// Ensure, that dialerMock does implement dialer.
// If this is not the case, regenerate this file with moq.
var _ dialer = &dialerMock{}

// dialerMock is a mock implementation of dialer.
//
//     func TestSomethingThatUsesdialer(t *testing.T) {
//
//         // make and configure a mocked dialer
//         mockeddialer := &dialerMock{
//             closeFunc: func() error {
// 	               panic("mock out the close method")
//             },
//             connectFunc: func() error {
// 	               panic("mock out the connect method")
//             },
//             connectCtxFunc: func(in1 context.Context) error {
// 	               panic("mock out the connectCtx method")
//             },
//             getAuthFunc: func() *auth {
// 	               panic("mock out the getAuth method")
//             },
//             isConnectedFunc: func() bool {
// 	               panic("mock out the isConnected method")
//             },
//             isDisposedFunc: func() bool {
// 	               panic("mock out the isDisposed method")
//             },
//             pingFunc: func(errs chan error)  {
// 	               panic("mock out the ping method")
//             },
//             pingCtxFunc: func(in1 context.Context, in2 chan error)  {
// 	               panic("mock out the pingCtx method")
//             },
//             readFunc: func() (int, []byte, error) {
// 	               panic("mock out the read method")
//             },
//             readCtxFunc: func(in1 context.Context, in2 chan message)  {
// 	               panic("mock out the readCtx method")
//             },
//             writeFunc: func(in1 []byte) error {
// 	               panic("mock out the write method")
//             },
//         }
//
//         // use mockeddialer in code that requires dialer
//         // and then make assertions.
//
//     }
type dialerMock struct {
	// closeFunc mocks the close method.
	closeFunc func() error

	// connectFunc mocks the connect method.
	connectFunc func() error

	// connectCtxFunc mocks the connectCtx method.
	connectCtxFunc func(in1 context.Context) error

	// getAuthFunc mocks the getAuth method.
	getAuthFunc func() *auth

	// isConnectedFunc mocks the isConnected method.
	isConnectedFunc func() bool

	// isDisposedFunc mocks the isDisposed method.
	isDisposedFunc func() bool

	// pingFunc mocks the ping method.
	pingFunc func(errs chan error)

	// pingCtxFunc mocks the pingCtx method.
	pingCtxFunc func(in1 context.Context, in2 chan error)

	// readFunc mocks the read method.
	readFunc func() (int, []byte, error)

	// readCtxFunc mocks the readCtx method.
	readCtxFunc func(in1 context.Context, in2 chan message)

	// writeFunc mocks the write method.
	writeFunc func(in1 []byte) error

	// calls tracks calls to the methods.
	calls struct {
		// close holds details about calls to the close method.
		close []struct {
		}
		// connect holds details about calls to the connect method.
		connect []struct {
		}
		// connectCtx holds details about calls to the connectCtx method.
		connectCtx []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// getAuth holds details about calls to the getAuth method.
		getAuth []struct {
		}
		// isConnected holds details about calls to the isConnected method.
		isConnected []struct {
		}
		// isDisposed holds details about calls to the isDisposed method.
		isDisposed []struct {
		}
		// ping holds details about calls to the ping method.
		ping []struct {
			// Errs is the errs argument value.
			Errs chan error
		}
		// pingCtx holds details about calls to the pingCtx method.
		pingCtx []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 chan error
		}
		// read holds details about calls to the read method.
		read []struct {
		}
		// readCtx holds details about calls to the readCtx method.
		readCtx []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 chan message
		}
		// write holds details about calls to the write method.
		write []struct {
			// In1 is the in1 argument value.
			In1 []byte
		}
	}
}

// close calls closeFunc.
func (mock *dialerMock) close() error {
	if mock.closeFunc == nil {
		panic("dialerMock.closeFunc: method is nil but dialer.close was just called")
	}
	callInfo := struct {
	}{}
	lockdialerMockclose.Lock()
	mock.calls.close = append(mock.calls.close, callInfo)
	lockdialerMockclose.Unlock()
	return mock.closeFunc()
}

// closeCalls gets all the calls that were made to close.
// Check the length with:
//     len(mockeddialer.closeCalls())
func (mock *dialerMock) closeCalls() []struct {
} {
	var calls []struct {
	}
	lockdialerMockclose.RLock()
	calls = mock.calls.close
	lockdialerMockclose.RUnlock()
	return calls
}

// connect calls connectFunc.
func (mock *dialerMock) connect() error {
	if mock.connectFunc == nil {
		panic("dialerMock.connectFunc: method is nil but dialer.connect was just called")
	}
	callInfo := struct {
	}{}
	lockdialerMockconnect.Lock()
	mock.calls.connect = append(mock.calls.connect, callInfo)
	lockdialerMockconnect.Unlock()
	return mock.connectFunc()
}

// connectCalls gets all the calls that were made to connect.
// Check the length with:
//     len(mockeddialer.connectCalls())
func (mock *dialerMock) connectCalls() []struct {
} {
	var calls []struct {
	}
	lockdialerMockconnect.RLock()
	calls = mock.calls.connect
	lockdialerMockconnect.RUnlock()
	return calls
}

// connectCtx calls connectCtxFunc.
func (mock *dialerMock) connectCtx(in1 context.Context) error {
	if mock.connectCtxFunc == nil {
		panic("dialerMock.connectCtxFunc: method is nil but dialer.connectCtx was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	lockdialerMockconnectCtx.Lock()
	mock.calls.connectCtx = append(mock.calls.connectCtx, callInfo)
	lockdialerMockconnectCtx.Unlock()
	return mock.connectCtxFunc(in1)
}

// connectCtxCalls gets all the calls that were made to connectCtx.
// Check the length with:
//     len(mockeddialer.connectCtxCalls())
func (mock *dialerMock) connectCtxCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	lockdialerMockconnectCtx.RLock()
	calls = mock.calls.connectCtx
	lockdialerMockconnectCtx.RUnlock()
	return calls
}

// getAuth calls getAuthFunc.
func (mock *dialerMock) getAuth() *auth {
	if mock.getAuthFunc == nil {
		panic("dialerMock.getAuthFunc: method is nil but dialer.getAuth was just called")
	}
	callInfo := struct {
	}{}
	lockdialerMockgetAuth.Lock()
	mock.calls.getAuth = append(mock.calls.getAuth, callInfo)
	lockdialerMockgetAuth.Unlock()
	return mock.getAuthFunc()
}

// getAuthCalls gets all the calls that were made to getAuth.
// Check the length with:
//     len(mockeddialer.getAuthCalls())
func (mock *dialerMock) getAuthCalls() []struct {
} {
	var calls []struct {
	}
	lockdialerMockgetAuth.RLock()
	calls = mock.calls.getAuth
	lockdialerMockgetAuth.RUnlock()
	return calls
}

// isConnected calls isConnectedFunc.
func (mock *dialerMock) isConnected() bool {
	if mock.isConnectedFunc == nil {
		panic("dialerMock.isConnectedFunc: method is nil but dialer.isConnected was just called")
	}
	callInfo := struct {
	}{}
	lockdialerMockisConnected.Lock()
	mock.calls.isConnected = append(mock.calls.isConnected, callInfo)
	lockdialerMockisConnected.Unlock()
	return mock.isConnectedFunc()
}

// isConnectedCalls gets all the calls that were made to isConnected.
// Check the length with:
//     len(mockeddialer.isConnectedCalls())
func (mock *dialerMock) isConnectedCalls() []struct {
} {
	var calls []struct {
	}
	lockdialerMockisConnected.RLock()
	calls = mock.calls.isConnected
	lockdialerMockisConnected.RUnlock()
	return calls
}

// isDisposed calls isDisposedFunc.
func (mock *dialerMock) isDisposed() bool {
	if mock.isDisposedFunc == nil {
		panic("dialerMock.isDisposedFunc: method is nil but dialer.isDisposed was just called")
	}
	callInfo := struct {
	}{}
	lockdialerMockisDisposed.Lock()
	mock.calls.isDisposed = append(mock.calls.isDisposed, callInfo)
	lockdialerMockisDisposed.Unlock()
	return mock.isDisposedFunc()
}

// isDisposedCalls gets all the calls that were made to isDisposed.
// Check the length with:
//     len(mockeddialer.isDisposedCalls())
func (mock *dialerMock) isDisposedCalls() []struct {
} {
	var calls []struct {
	}
	lockdialerMockisDisposed.RLock()
	calls = mock.calls.isDisposed
	lockdialerMockisDisposed.RUnlock()
	return calls
}

// ping calls pingFunc.
func (mock *dialerMock) ping(errs chan error) {
	if mock.pingFunc == nil {
		panic("dialerMock.pingFunc: method is nil but dialer.ping was just called")
	}
	callInfo := struct {
		Errs chan error
	}{
		Errs: errs,
	}
	lockdialerMockping.Lock()
	mock.calls.ping = append(mock.calls.ping, callInfo)
	lockdialerMockping.Unlock()
	mock.pingFunc(errs)
}

// pingCalls gets all the calls that were made to ping.
// Check the length with:
//     len(mockeddialer.pingCalls())
func (mock *dialerMock) pingCalls() []struct {
	Errs chan error
} {
	var calls []struct {
		Errs chan error
	}
	lockdialerMockping.RLock()
	calls = mock.calls.ping
	lockdialerMockping.RUnlock()
	return calls
}

// pingCtx calls pingCtxFunc.
func (mock *dialerMock) pingCtx(in1 context.Context, in2 chan error) {
	if mock.pingCtxFunc == nil {
		panic("dialerMock.pingCtxFunc: method is nil but dialer.pingCtx was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 chan error
	}{
		In1: in1,
		In2: in2,
	}
	lockdialerMockpingCtx.Lock()
	mock.calls.pingCtx = append(mock.calls.pingCtx, callInfo)
	lockdialerMockpingCtx.Unlock()
	mock.pingCtxFunc(in1, in2)
}

// pingCtxCalls gets all the calls that were made to pingCtx.
// Check the length with:
//     len(mockeddialer.pingCtxCalls())
func (mock *dialerMock) pingCtxCalls() []struct {
	In1 context.Context
	In2 chan error
} {
	var calls []struct {
		In1 context.Context
		In2 chan error
	}
	lockdialerMockpingCtx.RLock()
	calls = mock.calls.pingCtx
	lockdialerMockpingCtx.RUnlock()
	return calls
}

// read calls readFunc.
func (mock *dialerMock) read() (int, []byte, error) {
	if mock.readFunc == nil {
		panic("dialerMock.readFunc: method is nil but dialer.read was just called")
	}
	callInfo := struct {
	}{}
	lockdialerMockread.Lock()
	mock.calls.read = append(mock.calls.read, callInfo)
	lockdialerMockread.Unlock()
	return mock.readFunc()
}

// readCalls gets all the calls that were made to read.
// Check the length with:
//     len(mockeddialer.readCalls())
func (mock *dialerMock) readCalls() []struct {
} {
	var calls []struct {
	}
	lockdialerMockread.RLock()
	calls = mock.calls.read
	lockdialerMockread.RUnlock()
	return calls
}

// readCtx calls readCtxFunc.
func (mock *dialerMock) readCtx(in1 context.Context, in2 chan message) {
	if mock.readCtxFunc == nil {
		panic("dialerMock.readCtxFunc: method is nil but dialer.readCtx was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 chan message
	}{
		In1: in1,
		In2: in2,
	}
	lockdialerMockreadCtx.Lock()
	mock.calls.readCtx = append(mock.calls.readCtx, callInfo)
	lockdialerMockreadCtx.Unlock()
	mock.readCtxFunc(in1, in2)
}

// readCtxCalls gets all the calls that were made to readCtx.
// Check the length with:
//     len(mockeddialer.readCtxCalls())
func (mock *dialerMock) readCtxCalls() []struct {
	In1 context.Context
	In2 chan message
} {
	var calls []struct {
		In1 context.Context
		In2 chan message
	}
	lockdialerMockreadCtx.RLock()
	calls = mock.calls.readCtx
	lockdialerMockreadCtx.RUnlock()
	return calls
}

// write calls writeFunc.
func (mock *dialerMock) write(in1 []byte) error {
	if mock.writeFunc == nil {
		panic("dialerMock.writeFunc: method is nil but dialer.write was just called")
	}
	callInfo := struct {
		In1 []byte
	}{
		In1: in1,
	}
	lockdialerMockwrite.Lock()
	mock.calls.write = append(mock.calls.write, callInfo)
	lockdialerMockwrite.Unlock()
	return mock.writeFunc(in1)
}

// writeCalls gets all the calls that were made to write.
// Check the length with:
//     len(mockeddialer.writeCalls())
func (mock *dialerMock) writeCalls() []struct {
	In1 []byte
} {
	var calls []struct {
		In1 []byte
	}
	lockdialerMockwrite.RLock()
	calls = mock.calls.write
	lockdialerMockwrite.RUnlock()
	return calls
}
