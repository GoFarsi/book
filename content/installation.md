---
title: "آموزش نصب"
type: chapter
weight: 4
---

جهت نصب زبان گو برروی سیستم عامل مورد نظر خود در ابتدا باید نسخه مناسب با پلت فرم خود را [دانلود](https://go.dev/dl/) کنید.

![Featured downloads golang](assets/img/content/installation/1.png)

```tpl
{{</* tabs "uniqueid" */>}}
{{</* tab "Linux" */>}} # لینوکس {{</* /tab */>}}
{{</* tab "Windows" */>}} # ویندوز {{</* /tab */>}}
{{</* tab "MacOS" */>}} # مکینتاش {{</* /tab */>}}
{{</* /tabs */>}}
```

# Linux

1. در [اینجا](https://go.dev/dl/) نسخه مرتبط با لینوکس خود را دانلود کنید.
```shell
$ wget -c https://go.dev/dl/go1.xx.x.linux-amd64.tar.gz
```
2. سپس دستور زیر را بزنید تا اگر نسخه قبلی وجود دارد پاک شود و نسخه جدید در مسیر `usr/local/go/` جایگزین شود.
```shell
$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.xx.x.linux-amd64.tar.gz
```
3. افزودن مسیر `usr/local/go/bin/` به  {{< tooltip text="متغیرهای محیطی" note="environment variable" >}}
```shell
export PATH=$PATH:/usr/local/go/bin
```

{{< hint=info >}}
بهتر است برای اینکه هر بار هر بار ترمینال را باز میکنید و فایل go توسط shell شناخته شود دستور زیر را به فایل `home/{user}/.profile/` اضافه کنید.

```shell
export PATH=$PATH:/usr/local/go/bin
```

{{< /hint >}}

4. جهت اطمینان از اینکه زبان گو بدرستی برروی سیستم عامل  شما نصب شده است دستور زیر را داخل ترمینال بزنید.
```shell
go version
```


# Windows

آموزش نصب در ویندوز...


# MacOS

آموزش نصب در مکینتاش...