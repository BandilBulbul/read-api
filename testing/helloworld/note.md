#TDD
TDD is a skill that needs practice to develop but by being able to break problems down into smaller components that you can test you will have a much easier time writing software.

“Test-driven development” refers to a style of programming in which three activities are tightly interwoven: 
1.coding, 
1.testing (in the form of writing unit tests) and 
1.design (in the form of refactoring)

Test-driven development (TDD) is a development technique where you must first write a test that fails before you write new functional code.

HelloWorld Testing Doc:
1. write fail test ,its fail
1. add functionality code then again run , its fail
1. add functionality as a requirements 
1. try to run it then pass that one ...code is working
1. refactor the code  as a varible , constant, function should be having signle functionality then divide the functions into various acc to func perform a single task
1. refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with
1. refactor the testing code as a well crafted , should not be duplication code 
1. clean and clear 
1. each contains its own functionlity

In our case we've gone from Hello() to Hello("name"), to Hello("name", "French") in small, easy to understand steps.
With Some of Go's syntax around
1. Writing tests
1. Declaring functions, with arguments and return types
1. if, const and switch
1. Declaring variables and constants


# HelloWorld

## How it works:
 Executing initally start from main package with main , init function in golang
## How to test:
  On writing the test case , knew about input(string),output(string)
  file name should be name_test.go
  test function should be TestFunctionName
## Go Modules
  1.  go mod init hello
  1. This file tells the go tools essential information about your code
  1. download as well as information about dependencies
  1. To read more about modules, you can check out the reference in the Golang documentation.
## Testing:
 1. go test , go test -v , go test -cover
 
## Writing a test is just like writing a function, with a few rules
1. It needs to be in a file with a name like xxx_test.go
1. The test function must start with the word Test
1. The test function takes one argument only t *testing.T
1. In order to use the *testing.T type, you need to import "testing"

Note: you can do things like  t.Fail() when you want to fail.

### t.Errorf :
We are calling the Errorf method on our t which will print out a message and fail the test.The f stands for format which allows us to build a string with values inserted into the placeholder values %q

## Go Doc:
 You can launch the docs locally by running godoc -http :8000. If you go to localhost:8000/pkg you will see all the packages installed on your system.

 ## Note: Writing tests first,Let's start by capturing these requirements in a test. This is basic test driven development and allows us to make sure our test is actually testing what we want. When you retrospectively write tests there is the risk that your test may continue to pass even if the code doesn't work as intended.

 ## Steps:

 1. Wrote Hello()
 1. Add Argument Hello(string)
 1. Add Contants and Variable for refactoring like const englishHelloPrefix = "Hello, "

 1.  Add Requiremnets like when our function is called with an empty string it defaults to printing "Hello, World", rather than "Hello, ".
 1. Add another tool in our testing arsenal, subtests. ( it is useful to group tests around a "thing" and then have subtests describing different scenarios.)
(A benefit of this approach is you can set up shared code that can be used in the other tests.)
 Note:It is important that your tests are clear specifications of what the code needs to do.
 1. then refactor our  tests code.


## What have we done here:

1. We've refactored our assertion into a function. This reduces duplication and improves readability of our tests. 

1. For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy, so you can call helper functions from a test, or a benchmark.

1. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier.


## Note:  Discipline
Let's go over the cycle again
 1.  Write a test
 1.  Make the compiler pass
 1. Run the test, see that it fails and check the error message is meaningful
 1. Write enough code to make the test pass
 1. Refactor
 1. Add more requirements(We now need to support a second parameter, specifying the language of the greeting.)
 1.  if into switch(When you have lots of if statements checking a particular value it is common to use a switch statement instead. We can use switch to refactor the code to make it easier to read and more extensible)
 1. You could argue that maybe our function is getting a little big. The simplest refactor for this would be to extract out some functionality into another function.

 ## update new version of gopls v0.6.11 for Multiple Modules at same time.
## gopls
gopls (pronounced "Go please") is the official Go language server developed by the Go team. It provides IDE features to any LSP-compatible editor.

GO111MODULE=on go get golang.org/x/tools/gopls@latest

gopls supports both Go module and GOPATH modes, but if you are working with multiple modules or uncommon project layouts, you will need to specifically configure your workspace.