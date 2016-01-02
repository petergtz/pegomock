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

type Mock interface{}
type Param interface{}
type ReturnValue interface{}
type ReturnValues []ReturnValue

var lastInvocation *invocation
var argMatchers Matchers

type invocation struct {
	genericMock *GenericMock
	MethodName  string
	Params      []Param
}

type GenericMock struct {
	mockedMethods map[string]*mockedMethod
}

func (genericMock *GenericMock) Invoke(methodName string, params ...Param) ReturnValues {
	lastInvocation = &invocation{
		genericMock: genericMock,
		MethodName:  methodName,
		Params:      params,
	}
	return genericMock.getOrCreateMockedMethod(methodName).Invoke(params)
}

func (genericMock *GenericMock) stub(methodName string, paramMatchers []Matcher, returnValues ReturnValues) {
	genericMock.getOrCreateMockedMethod(methodName).stub(paramMatchers, returnValues)
}

func (genericMock *GenericMock) getOrCreateMockedMethod(methodName string) *mockedMethod {
	if _, ok := genericMock.mockedMethods[methodName]; !ok {
		genericMock.mockedMethods[methodName] = &mockedMethod{name: methodName}
	}
	return genericMock.mockedMethods[methodName]
}

func (genericMock *GenericMock) Reset(methodName string, params []Matcher) {
	// TODO: should be called from When
}

func (genericMock *GenericMock) Verify(invocationCountMatcher Matcher, methodName string, params ...Param) {
	if !invocationCountMatcher.matches(genericMock.NumMethodInvocations(methodName, params...)) {
		GlobalFailHandler("Mock not called. TODO: better message")
	}
}

func (genericMock *GenericMock) NumMethodInvocations(methodName string, params ...Param) int {
	if len(argMatchers) != 0 {
		checkArgument(len(argMatchers) == len(params),
			"If you use matchers, you must use matchers for all parameters. Example: TODO")
		result := genericMock.numMethodInvocationsUsingMatchers(methodName, argMatchers)
		argMatchers = nil
		return result
	}

	count := 0
	for _, invocation := range genericMock.mockedMethods[methodName].invocations {
		if reflect.DeepEqual(params, invocation) {
			count++
		}
	}
	return count
}

func (genericMock *GenericMock) numMethodInvocationsUsingMatchers(methodName string, paramMatchers Matchers) int {
	count := 0
	for _, invocation := range genericMock.mockedMethods[methodName].invocations {
		if paramMatchers.matches(invocation) {
			count++
		}
	}
	return count
}

type mockedMethod struct {
	name        string
	invocations [][]Param
	stubbings   Stubbings
}

func (method *mockedMethod) Invoke(params []Param) ReturnValues {
	method.invocations = append(method.invocations, params)
	stubbing := method.stubbings.find(params)
	if stubbing == nil {
		return ReturnValues{}
	}
	return stubbing.Invoke(params)
}

func (method *mockedMethod) stub(paramMatchers Matchers, returnValues ReturnValues) {
	stubbing := method.stubbings.findByMatchers(paramMatchers)
	if stubbing == nil {
		stubbing = &Stubbing{paramMatchers: paramMatchers}
		method.stubbings = append(method.stubbings, stubbing)
	}
	stubbing.returnValuesSequence = append(stubbing.returnValuesSequence, returnValues)
}

func (method *mockedMethod) removeLastInvocation() {
	method.invocations = method.invocations[:len(method.invocations)-1]
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

type Matchers []Matcher

func (matchers Matchers) matches(params []Param) bool {
	checkArgument(len(matchers) == len(params),
		"Number of params and matchers different: params: %v, matchers: %v",
		params, matchers)
	for i := range params {
		if !matchers[i].matches(params[i]) {
			return false
		}
	}
	return true
}

func (matchers *Matchers) append(matcher Matcher) {
	*matchers = append(*matchers, matcher)
}

type ongoingStubbing struct {
	genericMock   *GenericMock
	MethodName    string
	ParamMatchers []Matcher
	returnValues  []interface{}
}

func When(invocation ...interface{}) *ongoingStubbing {
	checkNotNil(lastInvocation,
		"when() requires an argument which has to be 'a method call on a mock'.")
	defer func() {
		lastInvocation = nil
		argMatchers = nil
	}()
	lastInvocation.genericMock.mockedMethods[lastInvocation.MethodName].removeLastInvocation()
	return &ongoingStubbing{
		genericMock:   lastInvocation.genericMock,
		MethodName:    lastInvocation.MethodName,
		ParamMatchers: paramMatchersFromArgMatchersOrParams(argMatchers, lastInvocation.Params),
		returnValues:  invocation,
	}
}

func paramMatchersFromArgMatchersOrParams(argMatchers []Matcher, params []Param) []Matcher {
	if len(argMatchers) == 0 {
		return transformParamsIntoEqMatchers(params)
	} else {
		checkArgument(len(argMatchers) == len(lastInvocation.Params),
			"You must use the same number of matchers as arguments. Example: TODO")
		return argMatchers
	}
}

func transformParamsIntoEqMatchers(params []Param) []Matcher {
	paramMatchers := make([]Matcher, len(params))
	for param := range params {
		paramMatchers = append(paramMatchers, &EqMatcher{param})
	}
	return paramMatchers
}

var genericMocks = make(map[Mock]*GenericMock)

func GetGenericMockFrom(mock Mock) *GenericMock {
	if genericMocks[mock] == nil {
		genericMocks[mock] = &GenericMock{mockedMethods: make(map[string]*mockedMethod)}
	}
	return genericMocks[mock]
}

func (stubbing *ongoingStubbing) ThenReturn(values ...ReturnValue) *ongoingStubbing {
	checkArguments(values, stubbing.returnValues)
	stubbing.genericMock.stub(stubbing.MethodName, stubbing.ParamMatchers, values)
	return stubbing
}

func checkArguments(stubbedReturnValues []ReturnValue, pseudoReturnValues []interface{}) {
	checkArgument(len(stubbedReturnValues) == len(pseudoReturnValues),
		"Different number of return values")
	for i := range stubbedReturnValues {
		checkArgument(reflect.TypeOf(stubbedReturnValues[i]) == reflect.TypeOf(pseudoReturnValues[i]),
			"Return type doesn't match")
	}
}

type Stubber struct {
	returnValue interface{}
}

func DoPanic(value interface{}) *Stubber {
	return &Stubber{returnValue: value}
}

func (stubber *Stubber) When(mock interface{}) {

}
