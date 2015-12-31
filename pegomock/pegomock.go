// Copyright 2015 Peter Goetz
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pegomock

import (
	"reflect"
	"testing"

	"github.com/petergtz/pegomock/pegomock/internal/testingtsupport"
	"github.com/petergtz/pegomock/pegomock/types"
)

const MaxNumParams = 20

var GlobalFailHandler types.FailHandler

func RegisterMockFailHandler(handler types.FailHandler) {
	GlobalFailHandler = handler
}
func RegisterMockTestingT(t *testing.T) {
	RegisterMockFailHandler(testingtsupport.BuildTestingTGomegaFailHandler(t))
}

type Invocation struct {
	Mock        Mock
	genericMock *GenericMock
	MethodName  MethodName
	Params      [MaxNumParams]Param
}

type MethodInvocation struct {
	MethodName MethodName
	Params     [MaxNumParams]Param
}

type InvocationMatcher struct {
	Mock        Mock
	genericMock *GenericMock
	MethodName  MethodName
	Params      [MaxNumParams]Matcher
}

type Stubbing struct {
	paramMatchers        Matchers
	returnValuesSequence []ReturnValues
	sequencePointer      int
}

func (stubbing *Stubbing) Invoke(params ...Param) ReturnValues {
	result := stubbing.returnValuesSequence[stubbing.sequencePointer]
	stubbing.sequencePointer++
	return result
}

type Stubbings2 []Stubbing

func (stubbings Stubbings2) find(params ...Param) *Stubbing {
	for _, stubbing := range stubbings {
		if stubbing.paramMatchers.matches(params) {
			return &stubbing
		}
	}
	return nil
}

func (stubbings Stubbings2) findByMatchers(paramMatchers ...Matcher) *Stubbing {
	for _, stubbing := range stubbings {
		if reflect.DeepEqual(stubbing.paramMatchers, paramMatchers) {
			return &stubbing
		}
	}
	return nil
}

type MockedMethod struct {
	name        MethodName
	invocations [][]Param
	stubbings   Stubbings2
}

func (method *MockedMethod) Invoke(params ...Param) ReturnValues {
	method.invocations = append(method.invocations, params)
	stubbing := method.stubbings.find(params)
	if stubbing == nil {
		return ReturnValues{}
	}
	return stubbing.Invoke(params)
}

func (method *MockedMethod) Stub(paramMatchers Matchers, returnValues ReturnValues) {
	stubbing := method.stubbings.find(paramMatchers)
	if stubbing != nil {
		stubbing.returnValuesSequence = append(stubbing.returnValuesSequence, returnValues)
	}
	method.stubbings = append(method.stubbings, Stubbing{paramMatchers: paramMatchers, returnValuesSequence: []ReturnValues{returnValues}})
}

type GenericMock struct {
	mock          Mock
	mockedMethods map[MethodName]MockedMethod
}

func (genericMock *GenericMock) Invoke(methodName MethodName, params ...Param) ReturnValues {
	var p [MaxNumParams]Param
	copy(p[:], params)
	// TODO: store genericMock in Invocation
	LastInvocation = &Invocation{Mock: genericMock.mock, MethodName: methodName, Params: p}

	if _, ok := genericMock.mockedMethods[methodName]; !ok {
		genericMock.mockedMethods[methodName] = MockedMethod{name: methodName}
	}
	mockedMethod := genericMock.mockedMethods[methodName]
	return mockedMethod.Invoke(params)
}

func (genericMock *GenericMock) Stub(methodName MethodName, paramMatchers []Matcher, returnValues ReturnValues) {
	if _, ok := genericMock.mockedMethods[methodName]; !ok {
		genericMock.mockedMethods[methodName] = MockedMethod{name: methodName}
	}
	mockedMethod := genericMock.mockedMethods[methodName]
	mockedMethod.Stub(paramMatchers, returnValues)
}

func (genericMock *GenericMock) Reset(methodName MethodName, params []Matcher) {
	// TODO: should be called from When
}

func (genericMock *GenericMock) NumMethodInvocations(methodName MethodName, paramMatchers Matchers) int {
	count := 0
	for _, invocation := range genericMock.mockedMethods[methodName].invocations {
		if paramMatchers.matches(invocation) {
			count++
		}
	}
	return count
}

type Matchers []Matcher

func (paramMatchers Matchers) matches(params []Param) bool {
	if len(paramMatchers) != len(params) {
		panic("Number of params and matchers different: TODO: numbers")
	}
	for i := range params {
		if !paramMatchers[i].matches(params[i]) {
			return false
		}
	}
	return true
}

type Mock interface{}
type MethodName string
type Param interface{}
type ReturnValue interface{}
type ReturnValues []ReturnValue

var LastInvocation *Invocation
var Invocations = make(map[Mock]map[MethodName]map[[MaxNumParams]Param]int)
var Stubbings = make(map[Mock]map[MethodName]map[[MaxNumParams]Matcher][]ReturnValues)
var StubbingPointer = make(map[Mock]map[MethodName]map[[MaxNumParams]Matcher]int)

type OngoingStubbing struct {
	lastMockInvocationMatcher InvocationMatcher
	returnValues              []interface{}
}

type Matcher interface {
	matches(param Param) bool
}

type EqMatcher struct {
	value Param
}

func (matcher *EqMatcher) matches(param Param) bool {
	return matcher.value == param
}

type AnyMatcher struct{}

var argMatchers []Matcher

func EqInt(value int) int {
	argMatchers = append(argMatchers, &EqMatcher{value: value})
	return value
}

func When(invocation ...interface{}) *OngoingStubbing {
	if LastInvocation == nil {
		panic("when() requires an argument which has to be 'a method call on a mock'.")
	}
	var LastInvocationMatcher InvocationMatcher
	LastInvocationMatcher.Mock = LastInvocation.Mock
	LastInvocationMatcher.genericMock = LastInvocation.genericMock
	LastInvocationMatcher.MethodName = LastInvocation.MethodName
	if len(argMatchers) == 0 {
		// TODO: Do proper translation into matchers
		LastInvocationMatcher.Params = [MaxNumParams]Matcher{}
	} else {
		// TODO: Do proper translation into matchers
		LastInvocationMatcher.Params = [MaxNumParams]Matcher{}
	}
	result := &OngoingStubbing{lastMockInvocationMatcher: LastInvocationMatcher, returnValues: invocation}

	if _, ok := Stubbings[LastInvocationMatcher.Mock]; !ok {
		Stubbings[LastInvocationMatcher.Mock] = make(map[MethodName]map[[MaxNumParams]Matcher][]ReturnValues)
	}
	if _, ok := Stubbings[LastInvocationMatcher.Mock][LastInvocationMatcher.MethodName]; !ok {
		Stubbings[LastInvocationMatcher.Mock][LastInvocationMatcher.MethodName] = make(map[[MaxNumParams]Matcher][]ReturnValues)
	}
	Stubbings[LastInvocationMatcher.Mock][LastInvocationMatcher.MethodName][LastInvocationMatcher.Params] = make([]ReturnValues, 0)

	if _, ok := Invocations[LastInvocation.Mock]; !ok {
		Invocations[LastInvocation.Mock] = make(map[MethodName]map[[MaxNumParams]Param]int)
	}
	if _, ok := Invocations[LastInvocation.Mock][LastInvocationMatcher.MethodName]; !ok {
		Invocations[LastInvocation.Mock][LastInvocation.MethodName] = make(map[[MaxNumParams]Param]int)
	}
	Invocations[LastInvocation.Mock][LastInvocation.MethodName][LastInvocation.Params] = 0

	LastInvocation = nil
	return result
}

var genericMocks = make(map[Mock]*GenericMock)

func Invoke(mock Mock, methodName MethodName, params ...Param) {
	// TODO: make this nicer:
	var p [MaxNumParams]Param
	for i := 0; i < len(params); i++ {
		p[i] = params[i]
	}
	if genericMocks[mock] == nil {
		genericMocks[mock] = &GenericMock{mock: mock, mockedMethods: make(map[MethodName]MockedMethod)}
	}
	LastInvocation = &Invocation{Mock: mock, genericMock: genericMocks[mock], MethodName: methodName, Params: p}

	if _, ok := Invocations[mock]; !ok {
		Invocations[mock] = make(map[MethodName]map[[MaxNumParams]Param]int)
	}
	if _, ok := Invocations[mock][methodName]; !ok {
		Invocations[mock][methodName] = make(map[[MaxNumParams]Param]int)
	}

	Invocations[mock][methodName][p]++

}

func Match(m map[[MaxNumParams]Matcher][]ReturnValues, params [MaxNumParams]Param, numParams int) []ReturnValues {
	return m[[MaxNumParams]Matcher{}]
}

func MatchPointer(m map[[MaxNumParams]Matcher]int, params [MaxNumParams]Param, numParams int) int {
	return m[[MaxNumParams]Matcher{}]
}

func IncPointer(m map[[MaxNumParams]Matcher]int, params [MaxNumParams]Param, numParams int) {
	m[[MaxNumParams]Matcher{}]++
}

func (stubbing *OngoingStubbing) ThenReturn(values ...ReturnValue) *OngoingStubbing {
	if len(values) != len(stubbing.returnValues) {
		panic("Different number of return values")
	}
	for i := range values {
		if reflect.TypeOf(values[i]) != reflect.TypeOf(stubbing.returnValues[i]) {
			panic("Return type doesn't match")
		}
	}
	p := make([]Matcher, MaxNumParams)
	copy(p, stubbing.lastMockInvocationMatcher.Params[:])
	stubbing.lastMockInvocationMatcher.genericMock.Stub(stubbing.lastMockInvocationMatcher.MethodName, p, values)
	Stubbings[stubbing.lastMockInvocationMatcher.Mock][stubbing.lastMockInvocationMatcher.MethodName][stubbing.lastMockInvocationMatcher.Params] =
		append(Stubbings[stubbing.lastMockInvocationMatcher.Mock][stubbing.lastMockInvocationMatcher.MethodName][stubbing.lastMockInvocationMatcher.Params],
			values)
	if _, ok := StubbingPointer[stubbing.lastMockInvocationMatcher.Mock]; !ok {
		StubbingPointer[stubbing.lastMockInvocationMatcher.Mock] = make(map[MethodName]map[[MaxNumParams]Matcher]int)
	}
	if _, ok := StubbingPointer[stubbing.lastMockInvocationMatcher.Mock][stubbing.lastMockInvocationMatcher.MethodName]; !ok {
		StubbingPointer[stubbing.lastMockInvocationMatcher.Mock][stubbing.lastMockInvocationMatcher.MethodName] = make(map[[MaxNumParams]Matcher]int)
	}

	StubbingPointer[stubbing.lastMockInvocationMatcher.Mock][stubbing.lastMockInvocationMatcher.MethodName][stubbing.lastMockInvocationMatcher.Params] = 0
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
