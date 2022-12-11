prompter
=======

[![Test Status](https://github.com/Songmu/prompter/workflows/test/badge.svg?branch=main)][actions]
[![Coverage Status](https://codecov.io/gh/Songmu/prompter/branch/main/graph/badge.svg)][codecov]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDev](https://pkg.go.dev/badge/github.com/Songmu/prompter)][godev]

[actions]: https://github.com/Songmu/prompter/actions?workflow=test
[codecov]: https://codecov.io/gh/Songmu/prompter
[license]: https://github.com/Songmu/prompter/blob/main/LICENSE
[godev]: http://pkg.go.dev/github.com/Songmu/prompter

## Description

utility for easy prompting in Golang

## Synopsis

```go
twitterID := prompter.Prompt("Enter your twitter ID", "")
lang := prompter.Choose("Which language do you like the most?", []string{"Perl", "Golang", "Scala", "Ruby"}, "Perl")
passwd := prompter.Password("Enter your password")
var likeSushi bool = prompter.YN("Do you like sushi?", true)
var likeBeer bool = prompter.YesNo("Do you like beer?", false)
```

## Features

- Easy to use
- Care non-interactive (not a tty) environment
  - `Default` is used and the process is not blocked
- No howeyc/gopass (which uses cgo) dependency
  - cross build friendly
- Customizable prompt setting by using `&prompter.Prompter{}` directly

## License

[MIT][license]

## Author

[Songmu](https://github.com/Songmu)
