# INSTALL GO:
   first , we have to do is to run command to install homebrew(it has a dependency on xcode)
   1. xcode-select --install
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

   1. install go
   1. brew install go
   
## Go Environment:(Go modules)
  Modules aim to solve problems related to dependency management, version selection and reproducible builds;they also enble users to run go cod outside of gopath.

## GoPATH and GoROOT
   1. GOPATH as the root of your project
   1. GOROOT and GOPATH are environment variables that define this layout. 
   1. GOROOT is a variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use different Go versions. 
   1. GOPATH is a variable that defines the root of your workspace

## SDK
A software development kit (SDK) is a set of software development tools that allows 
the creation of applications for a certain software package, software framework

## Module
A go.mod file will be generated, containing the module path, a Go version, and its dependency requirements, which are the other modules needed for a successful build.
         go mod init <modulepath>
### go mod commands
    The built-in documentation provides an overview of all available go mod commands.
    ( go help mod / go help mode init)

## install Editor:
1. brew cask install visual-studio-code
1. code .   {installed correctly}
1. code --install-extension golang.go   {To add Go support you must install an extension, there are a variety available for VS Code,}

## Go debugger:
go get -u github.com/go-delve/delve/cmd/dlv

## Go Linting
Linting is the process of running a program that will analyse code for potential errors
brew install golangci/tap/golangci-lint
