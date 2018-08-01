# [Akisa][akisa] (Container) ![][version-badge] [![][license-badge]](LICENSE)
Container is an implementation of an IoC (inversion of control) pattern which adheres to the dependency inversion principle of the [SOLID design principles][scotch-solid].

<hr/>

An IoC container facilitates dependency injection (DI) which allows us to remove dependencies from our code. It is a pattern that allows developers hot swap dependencies without breaking our application.

> &mdash; Built by [Samuel Oloruntoba][my-twitter] of [Rafdel][rafdel] as part of the [Akisa][akisa] project.

## Table of Contents <!-- omit in toc -->
<!-- TOC -->
- [Installation](#installation)
- [Usage](#usage)
    - [Binding](#binding)
    - [Make & Get](#make--get)
    - [Advanced binding](#advanced-binding)
    - [Function binding](#function-binding)
    - [Shared bindings or singletons](#shared-bindings-or-singletons)
    - [Invoke](#invoke)
    - [Alias](#alias)
    - [Has](#has)
- [Questions or Issues](#questions-or-issues)
- [Alternatives](#alternatives)
- [License](#license)

## Installation
We recommend locking to [SemVer](http://semver.org/) using Go's package manager [dep](https://golang.github.io/dep/)

```sh
dep ensure -add go.rafdel.co/akisa/container
```

## Usage
First, create a new instance of the container.

```go
package main

import "go.rafdel.co/akisa/container"

func main() {
    c := container.New()
    // ...
}
```

### Binding
After creating a container instance, we can bind dependencies to the container.

```go
c.Bind("stuff", "nonsense")
```

The `Bind()` call above binds a dependency "nonsense" to our container using the abstract (or identifier) "stuff". So, when we try to get a dependency from the container using the abstract "stuff", it'll return "nonsense".

### Make & Get
To get the dependency from the container, we do:

```go
dependency, err := c.Make("stuff")
```

We can also use another helper function `Get()`:

```go
dependency := c.Get("stuff")
```

*Note: `Get()` will panic if the dependency isn't found*

### Advanced binding
One of the most important bindings we can do is binding an interface to a concrete. *If you use an interface as an abstract, the concrete must implement the interface*.

```go
c.Bind(new(Formatter), HTMLFormatter{})
```

Doing this tells the container that whenever you need a new instance of the `Formatter` interface, the `HTMLFormatter` struct should be returned as the concrete.

**structs**

Just as we can use an interface as the abstract, we can also use a struct. Note that if you use a struct as an abstract, the concrete must be `nil` otherwise, the container will `panic`.

```go
c.Bind(Formatter{}, nil)
```

*The abstract and concrete can be of any type — but, **avoid using maps, slices and arrays as abstracts** as they can be quite problematic to resolve from the container*.

### Function binding
Binding functions to the container takes a different route. 

```go
c.Bind("stuff", func() string {
    return "nonsense"
})
```

When we try to get the binding above, the container will automatically call the function and return the return value from the function call.

```go
c.Get("stuff") // nonsense
```

**passing arguments**
```go
c.Bind("sum", func(a, b int) int {
    return a + b
})
```

If we try to get `c.Get("sum")` above, the invocation will fail as the function binding requires parameters to be passed to it. We can pass parameters to function bindings using:

```go
c.Make("sum", 10, 20)
```

**automatic injection**

Taking advantage of the dependency injection of our application, we can also resolve dependencies from the container.

For example, we can do:

```go
c.Bind(new(Formatter), HTMLFormatter{})

c.Bind("formatter", func(f Formatter) {
    // ...
})
```

When we try to `c.Get("formatter")`, the container will read the function and inject the dependencies from the container into the function. It'll panic if it cannot find the dependency.

*We can also resolve the "formatter" dependency above using `Make()`. If we don't pass parameters to `Make()`, it'll try to auto-resolve binding dependencies from the container.*

### Shared bindings or singletons
Currently, when you get a binding from the container — a new instance is returned. So if we pass a function as the concrete, a new instance of that function is returned.

For times when we want the entire application to share only one instance of the concrete, we bind using singletons, or shared bindings.

### Invoke
With `Invoke()`, we pass in a function (or struct method) and any interface we pass to function as arguments will automatically get resolved.

```go
value := c.Invoke(func(f Formatter) string {
    return "one step at a time"
})
```

This is useful for times when we just want to resolve a couple of dependencies but do not want to bind the result.

### Alias
We can give dependencies different names using the `Alias()` method. This can serve as a means to shorten the name of the dependency.

```go
c.Bind(new(Formatter), HTMLFormatter{})
c.Alias(new(Formatter), "formatter")
```

Aliasing doesn't affect existing bindings, it will only create a pointer to the underlying binding.

After creating an alias, we can resolve it as we normally would:

```go
c.Get("formatter")
```

*Attempting to alias a non-existent abstract will cause a panic*.

### Has
To check if a binding is present in a container, you can use the `Has()` method.

```go
if c.Has(new(Formatter)) {
    fmt.Printf("All is right with the world")
}
```

*It also works for aliases*

## Questions or Issues
For questions and support feel free to send me a message on [Twitter][my-twitter] or create a [new issue][issue]. If you discovered a bug or would like to make a feature request, you can create a [new issue][issue] explaining the bug/feature.

## Alternatives
- [Uber's dig](https://github.com/uber-go/dig) is a good DI container

## License
This project is licensed under the [MIT](LICENSE) license.

[akisa]: https://github.com/GoAkisa
[scotch-solid]: https://bit.ly/1HJaKXW
[rafdel]: https://rafdel.co
[my-twitter]: https://twitter.com/kayandrae07
[issue]: https://github.com/GoAkisa/Container/issues
[version-badge]: https://img.shields.io/github/tag/GoAkisa/Container.svg
[license-badge]: https://img.shields.io/apm/l/vim-mode.svg?longCache=true
[badge-progress]: https://img.shields.io/badge/status-progress-blue.svg
[badge-planning]: https://img.shields.io/badge/status-planning-orange.svg