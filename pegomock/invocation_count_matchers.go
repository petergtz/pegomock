package pegomock

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
