package pegomock

import "reflect"

const MaxNumParams = 20

type Invocation struct {
	Mock       interface{}
	MethodName string
	Params     [MaxNumParams]interface{}
}

var LastInvocation *Invocation
var Invocations = make(map[Invocation]int)
var Stubbings = make(map[Invocation][][]interface{})
var StubbingPointer = make(map[Invocation]int)

type Call struct {
	Invocation  Invocation
	ReturnVlaue interface{}
}

type OngoingStubbing struct {
	lastMockInvocation Invocation
	returnValues       []interface{}
}

func When(invocation ...interface{}) *OngoingStubbing {
	if LastInvocation == nil {
		panic("when() requires an argument which has to be 'a method call on a mock'.")
	}
	result := &OngoingStubbing{lastMockInvocation: *LastInvocation, returnValues: invocation}
	Stubbings[*LastInvocation] = make([][]interface{}, 0)

	Invocations[*LastInvocation] = 0

	LastInvocation = nil
	return result
}

func (stubbing *OngoingStubbing) ThenReturn(values ...interface{}) *OngoingStubbing {
	if len(values) != len(stubbing.returnValues) {
		panic("Different number of return values")
	}
	for i := range values {
		if reflect.TypeOf(values[i]) != reflect.TypeOf(stubbing.returnValues[i]) {
			panic("Return type doesn't match")
		}
	}
	Stubbings[stubbing.lastMockInvocation] = append(Stubbings[stubbing.lastMockInvocation], values)
	StubbingPointer[stubbing.lastMockInvocation] = 0
	return stubbing
}

type Stubber struct {
	returnValue interface{}
}

func DoPanic(value interface{}) *Stubber {
	return &Stubber{returnValue: value}
}

func (stubber *Stubber) When(mock interface{}) {

}
