---
title: "ابزارها, دستورات گو و معرفی IDE"
type: chapter
weight: 6
---

پس از اینکه زبان گو را برروی سیستم خود نصب کردید از طریق {{< tooltip text="ترمینال" note="Terminal" >}} و یا  {{< tooltip text="خط فرمان" note="Command prompt" >}}  ویندوز قابل دسترس است که می‌توانید با زدن کلمه `go` یکسری عملیات انجام دهید:

```shell
$ go
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	work        workspace maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	packages        package lists and patterns
	private         configuration for downloading non-public code
	testflag        testing flags
	testfunc        testing functions
	vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.
```

## جدول دستورات (Commands)

|  دستور      | توضیحات                                                     |
|---------|-----------------------------------------------------------------|
| build | با این دستور (`go build main.go`) می‌توانید فایل `go` را کامپایل  کنید.    |
| clean | با این دستور می‌توانید فایل‌های کش‌شده را پاکسازی کنید.    |
| doc | با این دستور (`go doc fmt.Println`) می‌توانید داکیومنت هر یک از توابع عمومی  را ببینید.    |
| env | با این دستور می‌توانید {{< tooltip text="متغیرهای محیطی" note="environment variable" >}} تنظیم شده زبان گو را ببینید و آن‌ها را با دستور `go env -w key=value` مقداردهی کنید.    |
| fmt | با این دستور می‌توانید کدهای خود را مرتب `go fmt ./...` کنید.    |
| install | با استفاده از این دستور می‌توانید یکسری پکیج‌ها را گرفته و کامپایل کنید.      |
| list | لیست پکیج‌ها و ماژول‌های دانلود شده.   |
| generate | با این دستور می‌توانید از قابلیت Generator زبان گو استفاده کنید و فایل generate کنید.    |
| mod | برای مدیریت و ایجاد فایل mod به ازای هر پروژه.    |
| get | با این دستور می‌توانید پکیجی را دانلود یا بروزرسانی کنید و همچنین برای استفاده از این دستور نیاز به `git` دارید که نصب باشد.   |
| work | با این دستور می‌توانید یک workspace ایجاد کنید و آن را مدیریت کنید.   |
| run | با این دستور (`go run main.go`) می‌توانید فایل‌های گو یا پروژه را اجرا کنید.   |
| test | با این دستور (`go test example_test.go`) می‌توانید فایل‌های تست را اجرا کنید.   |
| version | با این دستور می‌توانید نسخه نصب شده زبان گو را ببینید.  |
| vet | با این دستور می‌توانید اشتباهات کدهای خود را ببینید.  |


## اجرای کد ساده در زبان گو

در زیر یک نمونه کد ساده قرار دادیم که به سادگی می‌توانید با دستور `go run main.go` اجرا کنید و خروجی را ببینید.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!!!")
}
```

```shell
$ go run main.go
Hello world!!!
```

 در کد فوق، ما یکسری {{< tooltip text="کلید واژه" note="keyword" >}}  نظیر package , import , func استفاده کردیم که در بخش [1.2 کلید واژه و شناسه‌ها](https://book.gofarsi.ir/chapter-1/go-built-in-keywords-identifiers/) با آن‌ها آشنا می‌شوید.

## انواع نرم‌افزارهای محیط توسعه (IDE) و ویرایشگر کد (Code Editor)
در زبان برنامه‌نویسی گو نیز مثل سایر زبان‌ها، می‌توان از ابزارهای کدنویسی مانند  {{< tooltip text="نرم‌افزارهای محیط توسعه"note="IDE(s)" >}} و همچنین از  {{< tooltip text="ویرایشگرهای کد" note="Code Editor(s)" >}} استفاده کرد که محبوب‌ترین‌ها معرفی می‌شوند:

1. نرم‌افزار [Jetbrains Goland](https://www.jetbrains.com/go/) اگر با سایر محصولات جت‌برینز آشنا باشید به‌راحتی می‌توانید از این  {{< tooltip text="نرم‌افزار محیط توسعه"note="IDE" >}}  استفاده کنید.
2. نرم‌افزار [VsCode](https://code.visualstudio.com/) تقریباً همه برنامه‌نویس‌ها با آن آشنا هستند. یک  {{< tooltip text="ویرایشگر کد"note="Code Editor" >}} متن‌باز می‌باشد که بین برنامه‌نویس‌ها خیلی محبوب است و برای استفاده زبان گو در این نرم‌افزار می‌توان افزونه[ vscode-go](https://code.visualstudio.com/docs/languages/go) را نصب کنید.
3. نرم‌افزار [Vim](https://www.vim.org/) یک نرم‌افزار متن‌باز و معروف می‌باشد که داخل  {{< tooltip text="ترمینال" note="Terminal" >}} قابل استفاده است و از آن می‌توانید جهت توسعه با زبان گو استفاده کنید و البته نیازمند [نصب پلاگین](https://github.com/fatih/vim-go) می‌باشد.
