// Code generated by counterfeiter. DO NOT EDIT.
package pubtktfakes

import (
	"net/http"
	"sync"

	pubtkt "github.com/orange-cloudfoundry/go-auth-pubtkt"
)

type FakeAuthPubTkt struct {
	RawToTicketStub        func(string) (*pubtkt.Ticket, error)
	rawToTicketMutex       sync.RWMutex
	rawToTicketArgsForCall []struct {
		arg1 string
	}
	rawToTicketReturns struct {
		result1 *pubtkt.Ticket
		result2 error
	}
	rawToTicketReturnsOnCall map[int]struct {
		result1 *pubtkt.Ticket
		result2 error
	}
	RequestToTicketStub        func(*http.Request) (*pubtkt.Ticket, error)
	requestToTicketMutex       sync.RWMutex
	requestToTicketArgsForCall []struct {
		arg1 *http.Request
	}
	requestToTicketReturns struct {
		result1 *pubtkt.Ticket
		result2 error
	}
	requestToTicketReturnsOnCall map[int]struct {
		result1 *pubtkt.Ticket
		result2 error
	}
	SignTicketStub        func(*pubtkt.Ticket) error
	signTicketMutex       sync.RWMutex
	signTicketArgsForCall []struct {
		arg1 *pubtkt.Ticket
	}
	signTicketReturns struct {
		result1 error
	}
	signTicketReturnsOnCall map[int]struct {
		result1 error
	}
	TicketInHeaderStub        func(http.Header, *pubtkt.Ticket) error
	ticketInHeaderMutex       sync.RWMutex
	ticketInHeaderArgsForCall []struct {
		arg1 http.Header
		arg2 *pubtkt.Ticket
	}
	ticketInHeaderReturns struct {
		result1 error
	}
	ticketInHeaderReturnsOnCall map[int]struct {
		result1 error
	}
	TicketInRequestStub        func(*http.Request, *pubtkt.Ticket) error
	ticketInRequestMutex       sync.RWMutex
	ticketInRequestArgsForCall []struct {
		arg1 *http.Request
		arg2 *pubtkt.Ticket
	}
	ticketInRequestReturns struct {
		result1 error
	}
	ticketInRequestReturnsOnCall map[int]struct {
		result1 error
	}
	TicketInResponseStub        func(http.ResponseWriter, *pubtkt.Ticket) error
	ticketInResponseMutex       sync.RWMutex
	ticketInResponseArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *pubtkt.Ticket
	}
	ticketInResponseReturns struct {
		result1 error
	}
	ticketInResponseReturnsOnCall map[int]struct {
		result1 error
	}
	TicketToRawStub        func(*pubtkt.Ticket) (string, error)
	ticketToRawMutex       sync.RWMutex
	ticketToRawArgsForCall []struct {
		arg1 *pubtkt.Ticket
	}
	ticketToRawReturns struct {
		result1 string
		result2 error
	}
	ticketToRawReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	VerifyFromRequestStub        func(*http.Request) (*pubtkt.Ticket, error)
	verifyFromRequestMutex       sync.RWMutex
	verifyFromRequestArgsForCall []struct {
		arg1 *http.Request
	}
	verifyFromRequestReturns struct {
		result1 *pubtkt.Ticket
		result2 error
	}
	verifyFromRequestReturnsOnCall map[int]struct {
		result1 *pubtkt.Ticket
		result2 error
	}
	VerifyTicketStub        func(*pubtkt.Ticket, string) error
	verifyTicketMutex       sync.RWMutex
	verifyTicketArgsForCall []struct {
		arg1 *pubtkt.Ticket
		arg2 string
	}
	verifyTicketReturns struct {
		result1 error
	}
	verifyTicketReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthPubTkt) RawToTicket(arg1 string) (*pubtkt.Ticket, error) {
	fake.rawToTicketMutex.Lock()
	ret, specificReturn := fake.rawToTicketReturnsOnCall[len(fake.rawToTicketArgsForCall)]
	fake.rawToTicketArgsForCall = append(fake.rawToTicketArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RawToTicket", []interface{}{arg1})
	fake.rawToTicketMutex.Unlock()
	if fake.RawToTicketStub != nil {
		return fake.RawToTicketStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.rawToTicketReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthPubTkt) RawToTicketCallCount() int {
	fake.rawToTicketMutex.RLock()
	defer fake.rawToTicketMutex.RUnlock()
	return len(fake.rawToTicketArgsForCall)
}

func (fake *FakeAuthPubTkt) RawToTicketCalls(stub func(string) (*pubtkt.Ticket, error)) {
	fake.rawToTicketMutex.Lock()
	defer fake.rawToTicketMutex.Unlock()
	fake.RawToTicketStub = stub
}

func (fake *FakeAuthPubTkt) RawToTicketArgsForCall(i int) string {
	fake.rawToTicketMutex.RLock()
	defer fake.rawToTicketMutex.RUnlock()
	argsForCall := fake.rawToTicketArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthPubTkt) RawToTicketReturns(result1 *pubtkt.Ticket, result2 error) {
	fake.rawToTicketMutex.Lock()
	defer fake.rawToTicketMutex.Unlock()
	fake.RawToTicketStub = nil
	fake.rawToTicketReturns = struct {
		result1 *pubtkt.Ticket
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthPubTkt) RawToTicketReturnsOnCall(i int, result1 *pubtkt.Ticket, result2 error) {
	fake.rawToTicketMutex.Lock()
	defer fake.rawToTicketMutex.Unlock()
	fake.RawToTicketStub = nil
	if fake.rawToTicketReturnsOnCall == nil {
		fake.rawToTicketReturnsOnCall = make(map[int]struct {
			result1 *pubtkt.Ticket
			result2 error
		})
	}
	fake.rawToTicketReturnsOnCall[i] = struct {
		result1 *pubtkt.Ticket
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthPubTkt) RequestToTicket(arg1 *http.Request) (*pubtkt.Ticket, error) {
	fake.requestToTicketMutex.Lock()
	ret, specificReturn := fake.requestToTicketReturnsOnCall[len(fake.requestToTicketArgsForCall)]
	fake.requestToTicketArgsForCall = append(fake.requestToTicketArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("RequestToTicket", []interface{}{arg1})
	fake.requestToTicketMutex.Unlock()
	if fake.RequestToTicketStub != nil {
		return fake.RequestToTicketStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requestToTicketReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthPubTkt) RequestToTicketCallCount() int {
	fake.requestToTicketMutex.RLock()
	defer fake.requestToTicketMutex.RUnlock()
	return len(fake.requestToTicketArgsForCall)
}

func (fake *FakeAuthPubTkt) RequestToTicketCalls(stub func(*http.Request) (*pubtkt.Ticket, error)) {
	fake.requestToTicketMutex.Lock()
	defer fake.requestToTicketMutex.Unlock()
	fake.RequestToTicketStub = stub
}

func (fake *FakeAuthPubTkt) RequestToTicketArgsForCall(i int) *http.Request {
	fake.requestToTicketMutex.RLock()
	defer fake.requestToTicketMutex.RUnlock()
	argsForCall := fake.requestToTicketArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthPubTkt) RequestToTicketReturns(result1 *pubtkt.Ticket, result2 error) {
	fake.requestToTicketMutex.Lock()
	defer fake.requestToTicketMutex.Unlock()
	fake.RequestToTicketStub = nil
	fake.requestToTicketReturns = struct {
		result1 *pubtkt.Ticket
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthPubTkt) RequestToTicketReturnsOnCall(i int, result1 *pubtkt.Ticket, result2 error) {
	fake.requestToTicketMutex.Lock()
	defer fake.requestToTicketMutex.Unlock()
	fake.RequestToTicketStub = nil
	if fake.requestToTicketReturnsOnCall == nil {
		fake.requestToTicketReturnsOnCall = make(map[int]struct {
			result1 *pubtkt.Ticket
			result2 error
		})
	}
	fake.requestToTicketReturnsOnCall[i] = struct {
		result1 *pubtkt.Ticket
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthPubTkt) SignTicket(arg1 *pubtkt.Ticket) error {
	fake.signTicketMutex.Lock()
	ret, specificReturn := fake.signTicketReturnsOnCall[len(fake.signTicketArgsForCall)]
	fake.signTicketArgsForCall = append(fake.signTicketArgsForCall, struct {
		arg1 *pubtkt.Ticket
	}{arg1})
	fake.recordInvocation("SignTicket", []interface{}{arg1})
	fake.signTicketMutex.Unlock()
	if fake.SignTicketStub != nil {
		return fake.SignTicketStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.signTicketReturns
	return fakeReturns.result1
}

func (fake *FakeAuthPubTkt) SignTicketCallCount() int {
	fake.signTicketMutex.RLock()
	defer fake.signTicketMutex.RUnlock()
	return len(fake.signTicketArgsForCall)
}

func (fake *FakeAuthPubTkt) SignTicketCalls(stub func(*pubtkt.Ticket) error) {
	fake.signTicketMutex.Lock()
	defer fake.signTicketMutex.Unlock()
	fake.SignTicketStub = stub
}

func (fake *FakeAuthPubTkt) SignTicketArgsForCall(i int) *pubtkt.Ticket {
	fake.signTicketMutex.RLock()
	defer fake.signTicketMutex.RUnlock()
	argsForCall := fake.signTicketArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthPubTkt) SignTicketReturns(result1 error) {
	fake.signTicketMutex.Lock()
	defer fake.signTicketMutex.Unlock()
	fake.SignTicketStub = nil
	fake.signTicketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) SignTicketReturnsOnCall(i int, result1 error) {
	fake.signTicketMutex.Lock()
	defer fake.signTicketMutex.Unlock()
	fake.SignTicketStub = nil
	if fake.signTicketReturnsOnCall == nil {
		fake.signTicketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.signTicketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) TicketInHeader(arg1 http.Header, arg2 *pubtkt.Ticket) error {
	fake.ticketInHeaderMutex.Lock()
	ret, specificReturn := fake.ticketInHeaderReturnsOnCall[len(fake.ticketInHeaderArgsForCall)]
	fake.ticketInHeaderArgsForCall = append(fake.ticketInHeaderArgsForCall, struct {
		arg1 http.Header
		arg2 *pubtkt.Ticket
	}{arg1, arg2})
	fake.recordInvocation("TicketInHeader", []interface{}{arg1, arg2})
	fake.ticketInHeaderMutex.Unlock()
	if fake.TicketInHeaderStub != nil {
		return fake.TicketInHeaderStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.ticketInHeaderReturns
	return fakeReturns.result1
}

func (fake *FakeAuthPubTkt) TicketInHeaderCallCount() int {
	fake.ticketInHeaderMutex.RLock()
	defer fake.ticketInHeaderMutex.RUnlock()
	return len(fake.ticketInHeaderArgsForCall)
}

func (fake *FakeAuthPubTkt) TicketInHeaderCalls(stub func(http.Header, *pubtkt.Ticket) error) {
	fake.ticketInHeaderMutex.Lock()
	defer fake.ticketInHeaderMutex.Unlock()
	fake.TicketInHeaderStub = stub
}

func (fake *FakeAuthPubTkt) TicketInHeaderArgsForCall(i int) (http.Header, *pubtkt.Ticket) {
	fake.ticketInHeaderMutex.RLock()
	defer fake.ticketInHeaderMutex.RUnlock()
	argsForCall := fake.ticketInHeaderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthPubTkt) TicketInHeaderReturns(result1 error) {
	fake.ticketInHeaderMutex.Lock()
	defer fake.ticketInHeaderMutex.Unlock()
	fake.TicketInHeaderStub = nil
	fake.ticketInHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) TicketInHeaderReturnsOnCall(i int, result1 error) {
	fake.ticketInHeaderMutex.Lock()
	defer fake.ticketInHeaderMutex.Unlock()
	fake.TicketInHeaderStub = nil
	if fake.ticketInHeaderReturnsOnCall == nil {
		fake.ticketInHeaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.ticketInHeaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) TicketInRequest(arg1 *http.Request, arg2 *pubtkt.Ticket) error {
	fake.ticketInRequestMutex.Lock()
	ret, specificReturn := fake.ticketInRequestReturnsOnCall[len(fake.ticketInRequestArgsForCall)]
	fake.ticketInRequestArgsForCall = append(fake.ticketInRequestArgsForCall, struct {
		arg1 *http.Request
		arg2 *pubtkt.Ticket
	}{arg1, arg2})
	fake.recordInvocation("TicketInRequest", []interface{}{arg1, arg2})
	fake.ticketInRequestMutex.Unlock()
	if fake.TicketInRequestStub != nil {
		return fake.TicketInRequestStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.ticketInRequestReturns
	return fakeReturns.result1
}

func (fake *FakeAuthPubTkt) TicketInRequestCallCount() int {
	fake.ticketInRequestMutex.RLock()
	defer fake.ticketInRequestMutex.RUnlock()
	return len(fake.ticketInRequestArgsForCall)
}

func (fake *FakeAuthPubTkt) TicketInRequestCalls(stub func(*http.Request, *pubtkt.Ticket) error) {
	fake.ticketInRequestMutex.Lock()
	defer fake.ticketInRequestMutex.Unlock()
	fake.TicketInRequestStub = stub
}

func (fake *FakeAuthPubTkt) TicketInRequestArgsForCall(i int) (*http.Request, *pubtkt.Ticket) {
	fake.ticketInRequestMutex.RLock()
	defer fake.ticketInRequestMutex.RUnlock()
	argsForCall := fake.ticketInRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthPubTkt) TicketInRequestReturns(result1 error) {
	fake.ticketInRequestMutex.Lock()
	defer fake.ticketInRequestMutex.Unlock()
	fake.TicketInRequestStub = nil
	fake.ticketInRequestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) TicketInRequestReturnsOnCall(i int, result1 error) {
	fake.ticketInRequestMutex.Lock()
	defer fake.ticketInRequestMutex.Unlock()
	fake.TicketInRequestStub = nil
	if fake.ticketInRequestReturnsOnCall == nil {
		fake.ticketInRequestReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.ticketInRequestReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) TicketInResponse(arg1 http.ResponseWriter, arg2 *pubtkt.Ticket) error {
	fake.ticketInResponseMutex.Lock()
	ret, specificReturn := fake.ticketInResponseReturnsOnCall[len(fake.ticketInResponseArgsForCall)]
	fake.ticketInResponseArgsForCall = append(fake.ticketInResponseArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *pubtkt.Ticket
	}{arg1, arg2})
	fake.recordInvocation("TicketInResponse", []interface{}{arg1, arg2})
	fake.ticketInResponseMutex.Unlock()
	if fake.TicketInResponseStub != nil {
		return fake.TicketInResponseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.ticketInResponseReturns
	return fakeReturns.result1
}

func (fake *FakeAuthPubTkt) TicketInResponseCallCount() int {
	fake.ticketInResponseMutex.RLock()
	defer fake.ticketInResponseMutex.RUnlock()
	return len(fake.ticketInResponseArgsForCall)
}

func (fake *FakeAuthPubTkt) TicketInResponseCalls(stub func(http.ResponseWriter, *pubtkt.Ticket) error) {
	fake.ticketInResponseMutex.Lock()
	defer fake.ticketInResponseMutex.Unlock()
	fake.TicketInResponseStub = stub
}

func (fake *FakeAuthPubTkt) TicketInResponseArgsForCall(i int) (http.ResponseWriter, *pubtkt.Ticket) {
	fake.ticketInResponseMutex.RLock()
	defer fake.ticketInResponseMutex.RUnlock()
	argsForCall := fake.ticketInResponseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthPubTkt) TicketInResponseReturns(result1 error) {
	fake.ticketInResponseMutex.Lock()
	defer fake.ticketInResponseMutex.Unlock()
	fake.TicketInResponseStub = nil
	fake.ticketInResponseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) TicketInResponseReturnsOnCall(i int, result1 error) {
	fake.ticketInResponseMutex.Lock()
	defer fake.ticketInResponseMutex.Unlock()
	fake.TicketInResponseStub = nil
	if fake.ticketInResponseReturnsOnCall == nil {
		fake.ticketInResponseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.ticketInResponseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) TicketToRaw(arg1 *pubtkt.Ticket) (string, error) {
	fake.ticketToRawMutex.Lock()
	ret, specificReturn := fake.ticketToRawReturnsOnCall[len(fake.ticketToRawArgsForCall)]
	fake.ticketToRawArgsForCall = append(fake.ticketToRawArgsForCall, struct {
		arg1 *pubtkt.Ticket
	}{arg1})
	fake.recordInvocation("TicketToRaw", []interface{}{arg1})
	fake.ticketToRawMutex.Unlock()
	if fake.TicketToRawStub != nil {
		return fake.TicketToRawStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.ticketToRawReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthPubTkt) TicketToRawCallCount() int {
	fake.ticketToRawMutex.RLock()
	defer fake.ticketToRawMutex.RUnlock()
	return len(fake.ticketToRawArgsForCall)
}

func (fake *FakeAuthPubTkt) TicketToRawCalls(stub func(*pubtkt.Ticket) (string, error)) {
	fake.ticketToRawMutex.Lock()
	defer fake.ticketToRawMutex.Unlock()
	fake.TicketToRawStub = stub
}

func (fake *FakeAuthPubTkt) TicketToRawArgsForCall(i int) *pubtkt.Ticket {
	fake.ticketToRawMutex.RLock()
	defer fake.ticketToRawMutex.RUnlock()
	argsForCall := fake.ticketToRawArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthPubTkt) TicketToRawReturns(result1 string, result2 error) {
	fake.ticketToRawMutex.Lock()
	defer fake.ticketToRawMutex.Unlock()
	fake.TicketToRawStub = nil
	fake.ticketToRawReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthPubTkt) TicketToRawReturnsOnCall(i int, result1 string, result2 error) {
	fake.ticketToRawMutex.Lock()
	defer fake.ticketToRawMutex.Unlock()
	fake.TicketToRawStub = nil
	if fake.ticketToRawReturnsOnCall == nil {
		fake.ticketToRawReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.ticketToRawReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthPubTkt) VerifyFromRequest(arg1 *http.Request) (*pubtkt.Ticket, error) {
	fake.verifyFromRequestMutex.Lock()
	ret, specificReturn := fake.verifyFromRequestReturnsOnCall[len(fake.verifyFromRequestArgsForCall)]
	fake.verifyFromRequestArgsForCall = append(fake.verifyFromRequestArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("VerifyFromRequest", []interface{}{arg1})
	fake.verifyFromRequestMutex.Unlock()
	if fake.VerifyFromRequestStub != nil {
		return fake.VerifyFromRequestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.verifyFromRequestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthPubTkt) VerifyFromRequestCallCount() int {
	fake.verifyFromRequestMutex.RLock()
	defer fake.verifyFromRequestMutex.RUnlock()
	return len(fake.verifyFromRequestArgsForCall)
}

func (fake *FakeAuthPubTkt) VerifyFromRequestCalls(stub func(*http.Request) (*pubtkt.Ticket, error)) {
	fake.verifyFromRequestMutex.Lock()
	defer fake.verifyFromRequestMutex.Unlock()
	fake.VerifyFromRequestStub = stub
}

func (fake *FakeAuthPubTkt) VerifyFromRequestArgsForCall(i int) *http.Request {
	fake.verifyFromRequestMutex.RLock()
	defer fake.verifyFromRequestMutex.RUnlock()
	argsForCall := fake.verifyFromRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthPubTkt) VerifyFromRequestReturns(result1 *pubtkt.Ticket, result2 error) {
	fake.verifyFromRequestMutex.Lock()
	defer fake.verifyFromRequestMutex.Unlock()
	fake.VerifyFromRequestStub = nil
	fake.verifyFromRequestReturns = struct {
		result1 *pubtkt.Ticket
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthPubTkt) VerifyFromRequestReturnsOnCall(i int, result1 *pubtkt.Ticket, result2 error) {
	fake.verifyFromRequestMutex.Lock()
	defer fake.verifyFromRequestMutex.Unlock()
	fake.VerifyFromRequestStub = nil
	if fake.verifyFromRequestReturnsOnCall == nil {
		fake.verifyFromRequestReturnsOnCall = make(map[int]struct {
			result1 *pubtkt.Ticket
			result2 error
		})
	}
	fake.verifyFromRequestReturnsOnCall[i] = struct {
		result1 *pubtkt.Ticket
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthPubTkt) VerifyTicket(arg1 *pubtkt.Ticket, arg2 string) error {
	fake.verifyTicketMutex.Lock()
	ret, specificReturn := fake.verifyTicketReturnsOnCall[len(fake.verifyTicketArgsForCall)]
	fake.verifyTicketArgsForCall = append(fake.verifyTicketArgsForCall, struct {
		arg1 *pubtkt.Ticket
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("VerifyTicket", []interface{}{arg1, arg2})
	fake.verifyTicketMutex.Unlock()
	if fake.VerifyTicketStub != nil {
		return fake.VerifyTicketStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.verifyTicketReturns
	return fakeReturns.result1
}

func (fake *FakeAuthPubTkt) VerifyTicketCallCount() int {
	fake.verifyTicketMutex.RLock()
	defer fake.verifyTicketMutex.RUnlock()
	return len(fake.verifyTicketArgsForCall)
}

func (fake *FakeAuthPubTkt) VerifyTicketCalls(stub func(*pubtkt.Ticket, string) error) {
	fake.verifyTicketMutex.Lock()
	defer fake.verifyTicketMutex.Unlock()
	fake.VerifyTicketStub = stub
}

func (fake *FakeAuthPubTkt) VerifyTicketArgsForCall(i int) (*pubtkt.Ticket, string) {
	fake.verifyTicketMutex.RLock()
	defer fake.verifyTicketMutex.RUnlock()
	argsForCall := fake.verifyTicketArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthPubTkt) VerifyTicketReturns(result1 error) {
	fake.verifyTicketMutex.Lock()
	defer fake.verifyTicketMutex.Unlock()
	fake.VerifyTicketStub = nil
	fake.verifyTicketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) VerifyTicketReturnsOnCall(i int, result1 error) {
	fake.verifyTicketMutex.Lock()
	defer fake.verifyTicketMutex.Unlock()
	fake.VerifyTicketStub = nil
	if fake.verifyTicketReturnsOnCall == nil {
		fake.verifyTicketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyTicketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthPubTkt) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.rawToTicketMutex.RLock()
	defer fake.rawToTicketMutex.RUnlock()
	fake.requestToTicketMutex.RLock()
	defer fake.requestToTicketMutex.RUnlock()
	fake.signTicketMutex.RLock()
	defer fake.signTicketMutex.RUnlock()
	fake.ticketInHeaderMutex.RLock()
	defer fake.ticketInHeaderMutex.RUnlock()
	fake.ticketInRequestMutex.RLock()
	defer fake.ticketInRequestMutex.RUnlock()
	fake.ticketInResponseMutex.RLock()
	defer fake.ticketInResponseMutex.RUnlock()
	fake.ticketToRawMutex.RLock()
	defer fake.ticketToRawMutex.RUnlock()
	fake.verifyFromRequestMutex.RLock()
	defer fake.verifyFromRequestMutex.RUnlock()
	fake.verifyTicketMutex.RLock()
	defer fake.verifyTicketMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthPubTkt) recordInvocation(key string, args []interface{}) {
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

var _ pubtkt.AuthPubTkt = new(FakeAuthPubTkt)
