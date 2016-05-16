[![Build Status](https://travis-ci.org/petergtz/pegomock.svg?branch=master)](https://travis-ci.org/petergtz/pegomock)

PegoMock is a mocking framework for the [Go programming language](http://golang.org/). It integrates well with Go's built-in `testing` package, but can be used in other contexts too. It is based on [golang/mock](https://github.com/golang/mock), but uses a DSL closely related to [Mockito](http://site.mockito.org/mockito/docs/current/org/mockito/Mockito.html).

Getting Pegomock
================

Just `go get` it:

```
go get github.com/petergtz/pegomock/pegomock
```

This will download the package and install an executable `pegomock` in your `$GOPATH/bin`.

Getting Started
===============

Using Pegomock with Ginkgo
--------------------------

When a Pegomock verification fails, it calls a `FailHandler`. This is a function that you must provide using `pegomock.RegisterMockFailHandler()`.

If you’re using [Ginkgo](http://onsi.github.io/ginkgo/), all you need to do is:

```go
pegomock.RegisterFailHandler(ginkgo.Fail)
```

before you start your test suite.

Using Pegomock with Golang’s XUnit-style Tests
----------------------------------------------

Use it like this:

```go
func TestUsingMocks(t *testing.T) {
	RegisterTestingT(t)

	// use Pegomock here
}
```

There are two caveats:

-	You must register the `t *testing.T` passed to your test with Pegomock before you make any verifications associated with that test. So every `Test...` function in your suite should have the `RegisterTestingT(t)` line.
-	Pegomock uses a global (singleton) fail handler. This has the benefit that you don’t need to pass the fail handler down to each test, but does mean that you cannot run your XUnit style tests in parallel with Pegomock.

Generating Your First Mock and Using It
---------------------------------------

Let's assume you have:

```go
type Display interface {
	Show(text string)
}
```

The simplest way is to call `pegomock` from within your go package specifying the interface by its name:

```
cd path/to/package
pegomock generate Display -o mock_display_test.go
```

You can now use it in your tests:

```go
// creating mock
display := NewMockDisplay()

// using the mock
display.Show("Hello World!")

// verifying
display.VerifyWasCalledOnce().Show("Hello World!")
```

Why yet Another Mocking Framework for Go?
=========================================

I've looked at some of the other frameworks, but found none of them satisfying:
- [GoMock](https://github.com/golang/mock) seemed overly complicated when setting up mocks and verifying them. The command line interface is also not quite intuitive. That said, Pegomock is based on the GoMock, reusing mostly the mockgen code.
- [Counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) uses a DSL that I didn't find expressive enough. It often seems to need more lines of code too. In one of its samples, it uses e.g.:

	```go
	fake.DoThings("stuff", 5)
	Expect(fake.DoThingsCallCount()).To(Equal(1))

	str, num := fake.DoThingsArgsForCall(0)
	Expect(str).To(Equal("stuff"))
	Expect(num).To(Equal(uint64(5)))
	```

	In Pegomock, this can be written as simple as:

	```go
	fake.DoThings("stuff", 5)
	fake.VerifyWasCalledOnce().DoThings("stuff", 5)
	```
- [Hel](https://github.com/nelsam/hel) uses a new and interesting approach to setting up and verifying mocks. However, I wonder how flexible it actually is. E.g. how about providing a callback function when stubbing? Can this be modeled with its current approach using channels?

In addition, Pegomock provides a "watch" command similar to [Ginkgo](http://onsi.github.io/ginkgo/), which constantly watches over changes in an interface and updates its mocks. It gives the framework a much more dynamic feel, similar to mocking frameworks in Ruby or Java.

Using Mocks In Your Tests
=========================

Verifying Behavior
------------------

Interface:

```go
type Display interface {
	Show(text string)
}
```

Test:

```go
// creating mock:
display := NewMockDisplay()

// using the mock:
display.Show("Hello World!")

// verifying:
display.VerifyWasCalledOnce().Show("Hello World!")
```

Stubbing
--------

Interface:

```go
type PhoneBook interface {
	GetPhoneNumber(name string) string
}
```

Test:

```go
// creating the mock
phoneBook := NewMockPhoneBook()

// stubbing:
When(phoneBook.GetPhoneNumber("Tom")).ThenReturn("345-123-789")
When(phoneBook.GetPhoneNumber("Invalid")).ThenPanic("Invalid Name")

// prints "345-123-789":
fmt.Println(phoneBook.GetPhoneNumber("Tom"))

// panics:
fmt.Println(phoneBook.GetPhoneNumber("Invalid"))

// prints "", because GetPhoneNumber("Dan") was not stubbed
fmt.Println(phoneBook.GetPhoneNumber("Dan"))

// Although it is possible to verify a stubbed invocation, usually it's redundant
// If your code cares what GetPhoneNumber("Tom") returns, then something else breaks (often even before a verification gets executed).
// If your code doesn't care what GetPhoneNumber("Tom") returns, then it should not be stubbed.

// Not convinced? See http://monkeyisland.pl/2008/04/26/asking-and-telling.
phoneBook.VerifyWasCalledOnce().GetPhoneNumber("Tom")
```

-	By default, for all methods that return a value, a mock will return zero values.
-	Once stubbed, the method will always return a stubbed value, regardless of how many times it is called.

Argument Matchers
-----------------

Pegomock provides matchers for stubbing and verification.

Verification:

```go
display := NewMockDisplay()

// Calling mock
display.Show("Hello again!")

// Verification:
display.VerifyWasCalledOnce().Show(AnyString())
```

Stubbing:

```go
phoneBook := NewMockPhoneBook()

// Stubbing:
phoneBook.GetPhoneNumber(AnyString()).ThenReturn("123-456-789")

// Prints "123-456-789":
fmt.Println(phoneBook.GetPhoneNumber("Dan"))
// Also prints "123-456-789":
fmt.Println(phoneBook.GetPhoneNumber("Tom"))
```

**Important**: When you use argument matchers, you must always use them for all arguments:

```go
// Incorrect, panics:
When(contactList.getContactByFullName("Dan", AnyString())).thenReturn(Contact{...})
// Correct:
When(contactList.getContactByFullName(EqString("Dan"), AnyString())).thenReturn(Contact{...})
```

Verifying the Number of Invocations
-----------------------------------

```go
display := NewMockDisplay()

// Calling mock
display.Show("Hello")
display.Show("Hello, again")
display.Show("And again")

// Verification:
display.VerifyWasCalled(Times(3)).Show(AnyString())
// or:
display.VerifyWasCalled(AtLeast(3)).Show(AnyString())
// or:
display.VerifyWasCalled(Never()).Show("This one was never called")
```

Verifying in Order
------------------

```go
display1 := NewMockDisplay()
display2 := NewMockDisplay()

// Calling mocks
display1.Show("One")
display1.Show("Two")
display2.Show("Another two")
display1.Show("Three")

// Verification:
inOrderContext := new(InOrderContext)
display1.VerifyWasCalledInOrder(Once(), inOrderContext).Show("One")
display2.VerifyWasCalledInOrder(Once(), inOrderContext).Show("Another two")
display1.VerifyWasCalledInOrder(Once(), inOrderContext).Show("Three")
```

Note that it's not necessary to verify the call for `display.Show("Two")` if that one is not of any interested. An `InOrderContext` only verifies that the verifications that are done, are in order.

The Pegomock CLI
================

Installation
------------

Install it via:

```
go install github.com/petergtz/pegomock/pegomock
```

Generating Mocks
----------------

Pegomock can generate mocks in two different ways:

1.	by parsing source code Go files

	```
	pegomock generate [<flags>] <gofile>
	```

2.	by building a Go package and using reflection

	```
	pegomock generate [<flags>] [<packagepath>] <interfacename>
	```

Flags can be any of the following:

-	`--output,-o`: A file to which to write the resulting source code. If you don't set this, the code is printed to standard output.

-	`--package`: The package to use for the resulting mock class source code. If you don't set this, the package name is `mock_` concatenated with the package of the input file.

For more flags, run:

```
pegomock --help
```

Continuously Generating Mocks
-----------------------------

The `watch` command lets Pegomock generate mocks continuously on every change to an interface:

```
pegomock watch
```

For this, Pegomock expects an `interfaces_to_mock` file in the package directory where the mocks should be generated. In fact, `pegomock watch` will create it for you if it doesn't exist yet. The contents of the file are similar to the ones of the `generate` command:

```
# Any line starting with a # is treated as comment.

# interface name without package specifies an Interface in the current package:
PhoneBook

 # generates a mock for SomeInterfacetaken from mypackage:
path/to/my/mypackage SomeInterface

# you can also specify a Go file:
display.go
```

Most of the options from the `generate` command can be used here too.
