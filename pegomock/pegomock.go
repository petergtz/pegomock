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

type Mock interface{}
type Param interface{}
type ReturnValue interface{}
type ReturnValues []ReturnValue

var lastInvocation *invocation
var argMatchers Matchers

type invocation struct {
	genericMock *genericMockClass
	MethodName  string
	Params      []Param
}

type genericMockClass struct {
	mockedMethods map[string]*mockedMethod
}

func (genericMock *genericMockClass) Invoke(methodName string, params ...Param) ReturnValues {
	lastInvocation = &invocation{genericMock: genericMock, MethodName: methodName, Params: params}

	if _, ok := genericMock.mockedMethods[methodName]; !ok {
		genericMock.mockedMethods[methodName] = &mockedMethod{name: methodName}
	}
	return genericMock.mockedMethods[methodName].Invoke(params)
}

func (genericMock *genericMockClass) stub(methodName string, paramMatchers []Matcher, returnValues ReturnValues) {
	if _, ok := genericMock.mockedMethods[methodName]; !ok {
		genericMock.mockedMethods[methodName] = &mockedMethod{name: methodName}
	}
	genericMock.mockedMethods[methodName].stub(paramMatchers, returnValues)
}

func (genericMock *genericMockClass) Reset(methodName string, params []Matcher) {
	// TODO: should be called from When
}

func (genericMock *genericMockClass) NumMethodInvocationsUsingMatchers(methodName string, paramMatchers Matchers) int {
	count := 0
	for _, invocation := range genericMock.mockedMethods[methodName].invocations {
		if paramMatchers.matches(invocation) {
			count++
		}
	}
	return count
}

func (genericMock *genericMockClass) NumMethodInvocations(methodName string, params ...Param) int {
	// TODO: if something is in matchers stack, pop and call NumMethodInvocationsUsingMatchers
	count := 0
	for _, invocation := range genericMock.mockedMethods[methodName].invocations {
		if reflect.DeepEqual(params, invocation) {
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
	if stubbing != nil {
		stubbing.returnValuesSequence = append(stubbing.returnValuesSequence, returnValues)
	} else {
		method.stubbings = append(method.stubbings, &Stubbing{paramMatchers: paramMatchers, returnValuesSequence: []ReturnValues{returnValues}})
	}
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
	checkArgument(len(matchers) == len(params), "Number of params and matchers different: params: %v, matchers: %v", params, matchers)
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
	genericMock   *genericMockClass
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

	var paramMatchers []Matcher
	fmt.Println(argMatchers)
	if len(argMatchers) == 0 {
		paramMatchers = make([]Matcher, 0)
		for param := range lastInvocation.Params {
			paramMatchers = append(paramMatchers, &EqMatcher{param})
		}
	} else {
		checkArgument(len(argMatchers) == len(lastInvocation.Params),
			"You must use the same number of matchers as arguments")
		paramMatchers = argMatchers
	}
	lastInvocation.genericMock.mockedMethods[lastInvocation.MethodName].removeLastInvocation()
	return &ongoingStubbing{
		genericMock:   lastInvocation.genericMock,
		MethodName:    lastInvocation.MethodName,
		ParamMatchers: paramMatchers,
		returnValues:  invocation,
	}
}

var genericMocks = make(map[Mock]*genericMockClass)

func GetGenericMockFrom(mock Mock) *genericMockClass {
	if genericMocks[mock] == nil {
		genericMocks[mock] = &genericMockClass{mockedMethods: make(map[string]*mockedMethod)}
	}
	return genericMocks[mock]
}

func (stubbing *ongoingStubbing) ThenReturn(values ...ReturnValue) *ongoingStubbing {
	checkArgument(len(values) == len(stubbing.returnValues),
		"Different number of return values")
	for i := range values {
		checkArgument(reflect.TypeOf(values[i]) == reflect.TypeOf(stubbing.returnValues[i]),
			"Return type doesn't match")
	}
	stubbing.genericMock.stub(stubbing.MethodName, stubbing.ParamMatchers, values)
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
