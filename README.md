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

TODO

Verifying in Order
------------------

TODO

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

TODO
