# Version Checker for Go

go-update-checker is a go library for checking the version of a currently installed application or package against its latest release on github. It also enables caching and setting a minimum interval of days after which a updatecheck against the github API should be performed to prevent spamming the API.

Versions used with go-update-checker must follow [SemVer](http://semver.org/).

## Installation and Usage

Installation can be done with a normal `go get`:

```
$ go get github.com/Christian1984/go-update-checker
```

#### Update Check Example

```go
import (
    "fmt"

    updatechecker "github.com/Christian1984/go-update-checker"
)

func main() {
    uc := updatechecker.New("Christian1984", "go-update-checker", "Go Update Checker", "", 0, true, false)
    uc.CheckForUpdate("0.0.1")
    uc.PrintMessage()
    /*
    =============================================================
    === INFO: A new update is available for Go Update Checker ===

    Version: 0.0.2

    Title: Go Update Checker - 0.0.2

    Description:
    Changed receivers to pointer receivers


    Download the latest version here:
    https://github.com/Christian1984/go-update-checker/releases
    =============================================================
    */

    uc.CheckForUpdate("0.0.2")
    uc.PrintMessage()
    /*
    ========================================================================
    === INFO: You are running the latestest Version of Go Update Checker ===
    ========================================================================
    */

    /* alternatively use uc.Message (type string) in any other context */
}
```

## Issues and Contributing

If you find an issue with this library, please report an issue. If you'd like, we welcome any contributions. Fork this library and submit a pull request.
