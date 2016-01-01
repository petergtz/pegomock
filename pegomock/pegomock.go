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
	checkArgument(len(paramMatchers) == len(params), "Number of params and matchers different: params: %v, matchers: %v", params, paramMatchers)
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
	genericMock   *GenericMock
	MethodName    MethodName
	ParamMatchers []Matcher
	returnValues  []interface{}
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

func (matcher *AnyMatcher) matches(param Param) bool {
	return true
}

type AnyIntMatcher struct{}
type AnyStringMatcher struct{}
type AnyFloat32Matcher struct{}

func (matcher *AnyIntMatcher) matches(param Param) bool {
	return reflect.TypeOf(param).Kind() == reflect.Int
}

var argMatchers Matchers

func AppendEqMatcher(value Param) {
	argMatchers = append(argMatchers, &EqMatcher{value: value})
}

func AppendAnyMatcher() {
	argMatchers = append(argMatchers, &AnyMatcher{})
}

func EqInt(value int) int {
	AppendEqMatcher(value)
	return value
}

func EqString(value string) string {
	AppendEqMatcher(value)
	return value
}

func AnyString(value string) string {
	AppendAnyMatcher()
	return value
}

func When(invocation ...interface{}) *OngoingStubbing {
	checkNotNil(LastInvocation,
		"when() requires an argument which has to be 'a method call on a mock'.")
	defer func() {
		LastInvocation = nil
		argMatchers = nil
	}()

	var paramMatchers []Matcher
	if len(argMatchers) == 0 {
		paramMatchers = make([]Matcher, 0)
		for param := range LastInvocation.Params {
			paramMatchers = append(paramMatchers, &EqMatcher{param})
		}
	} else {
		checkArgument(len(argMatchers) == len(LastInvocation.Params),
			"You must use the same number of matchers as arguments")
		paramMatchers = argMatchers
	}
	// Remove last invocation:
	LastInvocation.genericMock.mockedMethods[LastInvocation.MethodName].invocations =
		LastInvocation.genericMock.mockedMethods[LastInvocation.MethodName].invocations[:len(LastInvocation.genericMock.mockedMethods[LastInvocation.MethodName].invocations)-1]
	return &OngoingStubbing{
		genericMock:   LastInvocation.genericMock,
		MethodName:    LastInvocation.MethodName,
		ParamMatchers: paramMatchers,
		returnValues:  invocation,
	}
}

var genericMocks = make(map[Mock]*GenericMock)

func GetGenericMockFrom(mock Mock) *GenericMock {
	if genericMocks[mock] == nil {
		genericMocks[mock] = &GenericMock{mock: mock, mockedMethods: make(map[MethodName]*MockedMethod)}
	}
	return genericMocks[mock]
}

func (stubbing *OngoingStubbing) ThenReturn(values ...ReturnValue) *OngoingStubbing {
	checkArgument(len(values) == len(stubbing.returnValues),
		"Different number of return values")
	for i := range values {
		checkArgument(reflect.TypeOf(values[i]) == reflect.TypeOf(stubbing.returnValues[i]),
			"Return type doesn't match")
	}
	stubbing.genericMock.Stub(stubbing.MethodName, stubbing.ParamMatchers, values)
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
