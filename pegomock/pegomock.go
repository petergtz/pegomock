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
	"fmt"
	"reflect"
	"testing"

	"github.com/petergtz/pegomock/pegomock/internal/testingtsupport"
	"github.com/petergtz/pegomock/pegomock/types"
)

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
	Params      []Param
}

type MethodInvocation struct {
	MethodName MethodName
	Params     []Param
}

type InvocationMatcher struct {
	Mock        Mock
	genericMock *GenericMock
	MethodName  MethodName
	Params      []Matcher
}

type Stubbing struct {
	paramMatchers        Matchers
	returnValuesSequence []ReturnValues
	sequencePointer      int
}

func (stubbing *Stubbing) Invoke(params []Param) ReturnValues {
	result := stubbing.returnValuesSequence[stubbing.sequencePointer]
	if stubbing.sequencePointer < len(stubbing.returnValuesSequence)-1 {
		stubbing.sequencePointer++
	}
	return result
}

type Stubbings []*Stubbing

func (stubbings Stubbings) find(params []Param) *Stubbing {
	for _, stubbing := range stubbings {
		if stubbing.paramMatchers.matches(params) {
			return stubbing
		}
	}
	return nil
}

func (stubbings Stubbings) findByMatchers(paramMatchers Matchers) *Stubbing {
	for _, stubbing := range stubbings {
		if reflect.DeepEqual(stubbing.paramMatchers, paramMatchers) {
			return stubbing
		}
	}
	return nil
}

type MockedMethod struct {
	name        MethodName
	invocations [][]Param
	stubbings   Stubbings
}

func (method *MockedMethod) Invoke(params []Param) ReturnValues {
	method.invocations = append(method.invocations, params)
	stubbing := method.stubbings.find(params)
	if stubbing == nil {
		return ReturnValues{}
	}
	return stubbing.Invoke(params)
}

func (method *MockedMethod) Stub(paramMatchers Matchers, returnValues ReturnValues) {
	stubbing := method.stubbings.findByMatchers(paramMatchers)
	if stubbing != nil {
		stubbing.returnValuesSequence = append(stubbing.returnValuesSequence, returnValues)
	} else {
		method.stubbings = append(method.stubbings, &Stubbing{paramMatchers: paramMatchers, returnValuesSequence: []ReturnValues{returnValues}})
	}
}

type GenericMock struct {
	mock          Mock
	mockedMethods map[MethodName]*MockedMethod
}

func (genericMock *GenericMock) Invoke(methodName MethodName, params ...Param) ReturnValues {
	LastInvocation = &Invocation{Mock: genericMock.mock, genericMock: genericMock, MethodName: methodName, Params: params}

	if _, ok := genericMock.mockedMethods[methodName]; !ok {
		genericMock.mockedMethods[methodName] = &MockedMethod{name: methodName}
	}
	return genericMock.mockedMethods[methodName].Invoke(params)
}

func (genericMock *GenericMock) Stub(methodName MethodName, paramMatchers []Matcher, returnValues ReturnValues) {
	if _, ok := genericMock.mockedMethods[methodName]; !ok {
		genericMock.mockedMethods[methodName] = &MockedMethod{name: methodName}
	}
	genericMock.mockedMethods[methodName].Stub(paramMatchers, returnValues)
}

func (genericMock *GenericMock) Reset(methodName MethodName, params []Matcher) {
	// TODO: should be called from When
}

func (genericMock *GenericMock) NumMethodInvocationsUsingMatchers(methodName MethodName, paramMatchers Matchers) int {
	count := 0
	for _, invocation := range genericMock.mockedMethods[methodName].invocations {
		if paramMatchers.matches(invocation) {
			count++
		}
	}
	return count
}

func (genericMock *GenericMock) NumMethodInvocations(methodName MethodName, params ...Param) int {
	// TODO: if something is in matchers stack, pop and call NumMethodInvocationsUsingMatchers
	count := 0
	for _, invocation := range genericMock.mockedMethods[methodName].invocations {
		if reflect.DeepEqual(params, invocation) {
			count++
		}
	}
	return count
}

type Matchers []Matcher

func (paramMatchers Matchers) matches(params []Param) bool {
	if len(paramMatchers) != len(params) {
		panic("Number of params and matchers different: params: " + fmt.Sprint(params) + ", matchers: " + fmt.Sprint(paramMatchers))
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
		LastInvocationMatcher.Params = make([]Matcher, 0)
	} else {
		// TODO: Do proper translation into matchers
		LastInvocationMatcher.Params = make([]Matcher, 0)
	}
	// Remove last invocation:
	LastInvocation.genericMock.mockedMethods[LastInvocation.MethodName].invocations = LastInvocation.genericMock.mockedMethods[LastInvocation.MethodName].invocations[:len(LastInvocation.genericMock.mockedMethods[LastInvocation.MethodName].invocations)-1]
	result := &OngoingStubbing{lastMockInvocationMatcher: LastInvocationMatcher, returnValues: invocation}
	LastInvocation = nil
	return result
}

var genericMocks = make(map[Mock]*GenericMock)

func GetGenericMockFrom(mock Mock) *GenericMock {
	if genericMocks[mock] == nil {
		genericMocks[mock] = &GenericMock{mock: mock, mockedMethods: make(map[MethodName]*MockedMethod)}
	}
	return genericMocks[mock]
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
	stubbing.lastMockInvocationMatcher.genericMock.Stub(stubbing.lastMockInvocationMatcher.MethodName, stubbing.lastMockInvocationMatcher.Params, values)
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
