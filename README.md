# [Akisa][akisa] (Container) ![][version-badge] [![][license-badge]](LICENSE)

> This is still in beta. For the stable release, expect a complete rewrite of the underlying implementation while keeping the API intact. So, if you use it, expect nothing to change.

Container is an implementation of an IoC (inversion of control) pattern which adheres to the dependency inversion principle of the [SOLID design principles][scotch-solid].

<hr/>

IoC facilitates dependency injection (DI) which allows us to remove dependencies from our code. It is a pattern that allows developers hot swap dependencies without breaking our application.

> &mdash; Built by [Samuel Oloruntoba][my-twitter] of [Rafdel][rafdel] as part of the [Akisa][akisa] project.

## Table of Contents <!-- omit in toc -->
<!-- TOC -->
1. [Install](#install)
2. [Primer](#primer)
3. [Usage](#usage)
4. [Questions or Issues](#questions-or-issues)
5. [Alternatives](#alternatives)
6. [License](#license)

## Install
We recommend locking to [SemVer](http://semver.org/) using Go's package manager [dep](https://golang.github.io/dep/)

```sh
dep ensure -add go.rafdel.co/akisa/container
```

## Primer
There's a short primer on dependency injection in the [wiki][wiki-primer]

## Usage

For complete documentation, head over to the [getting started][wiki-getting-started] page.

```go
package main

import (
    "fmt"
    "go.rafdel.co/akisa/container"
)

func main() {
    // new container instance
    c := container.New()

    // teaching the container to inject MarkdownFormat struct 
    // whenever the Format interface is requested
    c.Provide(new(Format), MarkdownFormat{}, false)
    
    // invoking a function with the interface and the container
    // automatically resolves the function arguments.
    // Since the invoked function returns a string, it's automatically returned
    text := c.Invoke(func(f Format) string {
        return f.Process("I can haz *ip*")
    })

    // alternatively

    // get the concretion bound to the interface
    formatter := c.Get(new(Format)).(Format)
    text := formatter.Process("I can haz *ip*")

    fmt.Printf("We got \"%s\" from the container", text)
}
```

For complete documentation, head over to the [getting started][wiki-getting-started] page.

## Questions or Issues
For questions and support feel free to send me a message on [Twitter][my-twitter] or create a [new issue][issue].

If you discovered any bug or would like to make a feature request, you can create a [new issue][issue] explaining the bug/feature.

## Alternatives
- [Uber's dig](https://github.com/uber-go/dig) is a good DI container

## License
This project is licensed under the [MIT](LICENSE) license.

[akisa]: https://github.com/GoAkisa
[scotch-solid]: https://bit.ly/1HJaKXW
[rafdel]: https://rafdel.co
[my-twitter]: https://twitter.com/kayandrae07
[wiki-primer]: ../../wiki/primer
[wiki-getting-started]: ../../wiki/getting-started
[issue]: https://github.com/GoAkisa/Container/issues
[version-badge]: https://img.shields.io/github/tag/GoAkisa/Container.svg
[license-badge]: https://img.shields.io/apm/l/vim-mode.svg?longCache=true
[badge-progress]: https://img.shields.io/badge/status-progress-blue.svg
[badge-planning]: https://img.shields.io/badge/status-planning-orange.svg