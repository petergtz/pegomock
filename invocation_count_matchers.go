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
	. "github.com/petergtz/pegomock/internal/matcher"
)

func Times(numDesiredInvocations int) *EqMatcher {
	return &EqMatcher{numDesiredInvocations}
}

func AtLeast(numDesiredInvocations int) *AtLeastIntMatcher {
	return &AtLeastIntMatcher{numDesiredInvocations}
}

func AtMost(numDesiredInvocations int) *AtMostIntMatcher {
	return &AtMostIntMatcher{numDesiredInvocations}
}

func Never() *EqMatcher {
	return &EqMatcher{0}
}

func Once() *EqMatcher {
	return &EqMatcher{1}
}

func Twice() *EqMatcher {
	return &EqMatcher{2}
}
