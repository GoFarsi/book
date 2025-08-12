---
title: نمونه سوالات مصاحبه
type: chapter
weight: 8
---

در این بخش ما یکسری سوالات مصاحبه به زبان فارسی ارائه دادیم که برای یافتن ضعف‌ها کاربردی می‌باشد.


{{< details title="1: چه تایپ‌هایی مقدار zero آن‌ها nil است؟" open=false >}}
- interfaces
- slices
- channels
- maps
- pointers
- functions
{{< /details >}}

{{< details title="2: تایپ‌های نوع Reference؟" open=false >}}
- Pointers
- slices
- maps
- functions
- channels
{{< /details >}}

{{< details title="3: تایپ‌های نوع Aggregate؟" open=false >}}
- Array
- structs
{{< /details >}}

{{< details title="4: چه وقت باید از پوینتر استفاده کنیم؟" open=false >}}
**1- تابعی که یکی از پارامترهای خود را تغییر می‌دهد**
-وقتی تابعی را فراخوانی می‌کنیم که یک پوینتر را به عنوان پارامتر می‌گیرد، انتظار داریم که متغیر ما تغییر داده شود. اگر شما متغیر را در تابع خود تغییر نمی‌دهید، پس احتمالا نباید از پوینتر استفاده کنید.

**2- عملکرد بهتر**
-اگر رشته‌ای داشته باشید که شامل یک رمان کامل در حافظه باشد، کپی کردن این متغیر هر بار که به یک تابع جدید ارسال می‌شود، کاری بسیار گران است. ممکن است ارزشمند باشد که به جای این کار یک پوینتر را ارسال کنید، که باعث صرفه‌جویی در پردازنده و حافظه می‌شود. با این حال انجام این کار به قیمت خوانا بودن است، بنابراین فقط در صورت لزوم این بهینه‌سازی را انجام دهید.

**3- به گزینه nil نیاز دارید**
-گاهی اوقات یک تابع باید بداند که مقدار یک چیزی چیست، همچنین باید وجود یا عدم وجود آن را بداند. معمولا هنگام خواندن JSON از این استفاده می‌کنیم تا بدانیم فیلدی وجود دارد یا خیر.
{{< /details >}}

{{< details title="5: زبان گولنگ از موارد زیر پشتیبانی نمی‌کند؟" open=false >}}
- type inheritance
- operator overloading
- method overloading
- pointer arithmetic
- struct type in consts
{{< /details >}}

{{< details title="6: برای گوروتین ها چه مواقعی از channel و چه مواقعی از mutex استفاده می شود؟" open=false >}}
معمولاً در مواقعی که گوروتین ها نیاز به برقراری ارتباط با یکدیگر دارند، از channels استفاده می کنیم. درصورتی که قسمتی از کد ما(برای مثال مقدار متغیری را تغییر می دهیم) که در آن واحد فقط باید یک گوروتین به آن دسترسی داشته باشد، از یک قفل مانند mutext استفاده می کنیم.
{{< /details >}}

{{< details title="7: تفاوت بین goroutine و thread را توضیح دهید." open=false >}}
Goroutines سبک وزن هستند و دارای یک استک اولیه کوچک‌تر که به صورت پویا گسترش می‌یابد هستند، این در حالی است که threads استک ثابت دارند. Goroutines هم‌زمانی را در سطح زبان با استفاده از channelها مدیریت می‌کنند، در حالی که threads ممکن است نیاز به lockهای صریح داشته باشند. همچنین، سوئیچینگ بین goroutines کارایی بیشتری نسبت به thread switching دارد.
{{< /details >}}

{{< details title="8: توضیح دهید که interface در Golang چیست و چگونه می‌تواند مورد استفاده قرار گیرد." open=false >}}
یک interface در Go یک نوع خاص است که مجموعه‌ای از method signatures را تعریف می‌کند. هر نوع داده‌ای که این متدها را پیاده‌سازی کند، می‌تواند به عنوان آن interface مورد استفاده قرار گیرد. این بدون نیاز به ارث بری صورت می‌گیرد و امکان داکتایپینگ را فراهم می‌کند.
{{< /details >}}

{{< details title="9: چطور می‌توانید memory leak در برنامه‌های نوشته شده به وسیله Golang را شناسایی و مدیریت کنید؟" open=false >}}
برای شناسایی memory leaks در Golang، می‌توان از ابزارهایی مانند pprof به همراه نمودارهای ساخته شده بر اساس heap dumps استفاده کرد. برای پیشگیری از memory leaks، باید دقت کرد که از داده‌ها به درستی استفاده شود، حافظه رزرو شده آزاد گردد و منابع بسته شوند هنگامی که دیگر نیازی به آنها نیست.
{{< /details >}}

{{< details title="10: در Golang چگونه می‌توانیم dependency management را انجام دهیم؟" open=false >}}
Golang از Go Modules برای مدیریت وابستگی‌ها استفاده می‌کند که به توسعه دهندگان امکان می‌دهد پروژه‌ها را به صورت مستقل از GOPATH راحت‌تر مدیریت کنند. با استفاده از دستوراتی مانند `go mod init`, `go mod tidy`, و `go mod vendor` می‌توان وابستگی‌های لازم برای پروژه را مدیریت کرد.
{{< /details >}}

{{< details title="11: چه زمانی یک channel در Golang باید با buffer مورد استفاده قرار گیرد؟" open=false >}}
یک channel با buffer زمانی مورد استفاده قرار گیرد که شما می‌خواهید ارتباطات بین goroutines را بدون ایجاد blocking فوری داشته باشید. این امر می‌تواند بازده کدها را در مواقعی که عملیات‌ها از لحاظ عملکرد اندکی نابرابر هستند، بهبود بخشد.
{{< /details >}}


{{< details title="12: توضیح دهید که defer statement چیست و چرا ممکن است از آن استفاده کنیم." open=false >}}
Defer statement برای تضمین اجرای یک تابع مشخص، درست قبل از خارج شدن از تابع فعلی استفاده می‌شود. این برای راحتی در مدیریت منابع مثل بستن فایل‌ها و ارتباطات شبکه استفاده می‌شود که می‌خواهیم اطمینان حاصل کنیم که به‌درستی بسته خواهند شد.
{{< /details >}}

{{< details title="13: در Golang چگونه می‌توانید اطمینان حاصل کنید که یک goroutine نتیجه‌ای تولید می‌کند قبل از اینکه برنامه کار خود را به طور کامل متوقف کند؟" open=false >}}
برای اطمینان از اینکه یک goroutine کار خود را به اتمام برساند، معمولا از sync.WaitGroup برای همچین مدیریتی استفاده می‌کنیم. ساختار WaitGroup اجازه می‌دهد تا اصلی‌ترین goroutine صبر کند تا یک یا چند goroutines دیگر کار خود را تمام کنند.
{{< /details >}}

{{< details title="14: توضیح دهید که واحد ایزوله برای کد نویسی در Golang چیست (table-driven tests) و چرا مفید است." open=false >}}
Table-driven tests شیوه‌ای برای نوشتن تست‌ها است که از جداول داده برای تعریف چندین case تست بهره می‌برد. این شیوه مفید است زیرا کد تست را می‌توان برای بسیاری از داده‌ها به راحتی توسعه داد و به خوبی سازماندهی می‌شود.
{{< /details >}}

{{< details title="15: چرا Go از ارث بری (inheritance) پشتیبانی نمی‌کند و از composition به عنوان جایگزین استفاده می‌کند؟" open=false >}}
Go ارث بری را پیاده‌سازی نمی‌کند زیرا می‌تواند پیچیده شود و معماری نرم‌افزار را سخت‌تر مدیریت کند. در عوض، از composition استفاده می‌کند که می‌تواند code reuse را تشویق کند و طراحی سیستم را ساده‌تر و ماژولارتر می‌کند.
{{< /details >}}

{{< details title="16: آیا در Golang می‌توان از polymorphism استفاده کرد؟ اگر بله، چگونه؟" open=false >}}
بله، در Go می‌توان از polymorphism استفاده کرد از طریق استفاده از interfaces. یک interface می‌تواند برای تعریف یک مجموعه از روش‌ها به کار رود و هر نوع که این روش‌ها را پیاده‌سازی کند به عنوان آن نوع interface شناخته شود.
{{< /details >}}

{{< details title="17: چه تفاوتی میان make و new در Golang وجود دارد؟" open=false >}}
`make` در Go برای ایجاد sliceها، maps و channels استفاده می‌شود و یک ابجکت از نوع مورد نظر را با مقدار اولیه مشخصی برمی‌گرداند. از طرفی `new` یک pointer به یک ابجکت از یک نوع داده‌ای تعریف شده توسط کاربر را برمی‌گرداند که صفر اولیه شده است.
{{< /details >}}

{{< details title="18: متود (method) receivers در Golang چگونه کار می‌کند و تفاوت بین استفاده از pointer receiver و value receiver چیست؟" open=false >}}
Method receivers در Go اجازه می‌دهند تا روی نوع معینی از مقادیر عملیات انجام دهیم. استفاده از pointer receiver به ما اجازه می‌دهد تا تغییراتی که در method روی receiver اعمال می‌شوند را بر روی خود آبجکت اصلی اعمال کنیم، در حالیکه استفاده از value receiver یک کپی از مقدار را می‌گیرد و تغییرات او روی کپی صورت می‌گیرد و بر آبجکت اصلی اثر نمی‌گذارد.
{{< /details >}}

{{< details title="19: چگونه می‌توان در Golang یک پکیج اختصاصی ایجاد کرد و چگونه می‌توان آن را در دیگر فایل‌های Go مورد استفاده قرار داد؟" open=false >}}
برای ایجاد پکیج اختصاصی در Go، کد مربوطه باید در یک دایرکتوری قرار داده شود و بالای فایل‌های Go باید `package mypackage` تعریف شود. برای استفاده از پکیج، `import "path/to/mypackage"` باید در دیگر فایل‌ها قرار داده شود.
{{< /details >}}

{{< details title="20: در Golang، چگونه می‌توانید error handling را اجرا کنید و چه روش‌هایی برای پیاده‌سازی custom error types وجود دارد؟" open=false >}}
Error handling در Go اغلب از طریق بازگرداندن ارور از توابع و بررسی آنها انجام می‌شود. برای ایجاد custom error types، می‌توانید از `errors.New()` برای ایجاد یک ارور ساده استفاده کنید یا یک تایپ که ارور را پیاده‌سازی می‌کند با متدهای اضافی برای داده‌های اضافی مرتبط با ارور ایجاد کرد.
{{< /details >}}

{{< details title="21: کامپایلر گولنگ از نوع AOT است یا JIT؟ تفاوت بین AOT و JIT را بگو." open=false >}}
کامپایلر گو یک Ahead Of Time compilation است. تفاوت AOT با JIT در این است که کامپایلر های AOT مستقیم کد ما را تبدیل به machine code می کنند اما در کامپایتر های JIT کد ما تبدیل به یک کد میانی Bytecode می شود و در زمان اجرا توسط runtime engine هر قسمت از برنامه که مورد استفاده قرار می گیرد، تفسیر می شود و تبدیل به machine code می شود.
{{< /details >}}

{{< details title="22: تفاوت بین nil و مقدار صفر در انواع داده چیست؟" open=false >}}
در Go مقدار **nil** مخصوص انواع reference است (مانند slice, map, channel, function, pointer, interface) و به معنای عدم وجود داده یا ارجاع است.
مقدار صفر (zero value) برای انواع مقداری مانند int، bool، struct و غیره تعریف می‌شود (مثلاً 0 برای int یا false برای bool) و همیشه یک مقدار معتبر است.
{{< /details >}}

{{< details title="23: تفاوت بین close(channel) و خواندن از یک کانال بسته چیست؟" open=false >}}

* `close(channel)` تنها برای ارسال‌کننده استفاده می‌شود تا به دریافت‌کننده‌ها اطلاع دهد داده‌ای دیگر ارسال نخواهد شد.
* خواندن از یک کانال بسته مقدار صفر نوع داده را برمی‌گرداند بدون این که بلاک شود.
* ارسال روی کانال بسته **panic** ایجاد می‌کند.
  {{< /details >}}

{{< details title="24: چه زمانی باید از context استفاده کنیم؟" open=false >}}
زمانی که نیاز داریم پردازش‌ها یا درخواست‌های طولانی را لغو کنیم یا timeout بگذاریم، از `context.Context` استفاده می‌کنیم. این ابزار برای همگام‌سازی لغو عملیات بین goroutineها و جلوگیری از resource leak بسیار مهم است.
{{< /details >}}

{{< details title="25: تفاوت sync.Mutex و sync.RWMutex چیست؟" open=false >}}

* `sync.Mutex` قفل ساده‌ای است که در یک زمان فقط اجازه دسترسی به یک goroutine را می‌دهد.
* `sync.RWMutex` دو حالت دارد: قفل خواندن (می‌تواند همزمان توسط چند goroutine گرفته شود) و قفل نوشتن (انحصاری). این باعث بهبود عملکرد در سناریوهایی با خواندن زیاد و نوشتن کم می‌شود.
  {{< /details >}}

{{< details title="26: data race چیست و چگونه می‌توان از آن جلوگیری کرد؟" open=false >}}
**Data race** زمانی رخ می‌دهد که چند goroutine همزمان به یک متغیر مشترک دسترسی پیدا کرده و حداقل یکی از آنها عملیات نوشتن انجام دهد، بدون هماهنگ‌سازی مناسب.
جلوگیری: استفاده از `sync.Mutex`، `sync.RWMutex`، channelها یا اجتناب از اشتراک داده‌ها. ابزار `go run -race` می‌تواند این مشکلات را پیدا کند.
{{< /details >}}

{{< details title="27: در چه شرایطی slice باعث memory leak می‌شود؟" open=false >}}
زمانی که slice‌ای کوچک ایجاد می‌کنیم ولی همچنان به یک آرایه بزرگ‌تر در حافظه اشاره دارد، GC نمی‌تواند آرایه اصلی را آزاد کند. راه‌حل: ایجاد یک کپی slice جدید با `copy` یا ساختن slice از ابتدا.
{{< /details >}}

{{< details title="28: تفاوت len و cap در slice چیست؟" open=false >}}

* `len` تعداد عناصر موجود در slice را برمی‌گرداند.
* `cap` ظرفیت slice را برمی‌گرداند، یعنی تعداد عناصر از اندیس صفر تا انتهای آرایه زیرین که می‌توان بدون تخصیص حافظه جدید استفاده کرد.
  {{< /details >}}

{{< details title="29: آیا map در Go thread-safe است؟" open=false >}}
خیر. map به صورت پیش‌فرض thread-safe نیست و دسترسی همزمان بدون قفل یا sync.Map باعث panic می‌شود. برای ایمنی، باید از `sync.Mutex`، `sync.RWMutex` یا `sync.Map` استفاده کرد.
{{< /details >}}

{{< details title="30: چگونه می‌توان یک panic را مدیریت کرد بدون اینکه برنامه متوقف شود؟" open=false >}}
با استفاده از `recover` در داخل یک defer می‌توان panic را گرفتن و از توقف کامل برنامه جلوگیری کرد:

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

{{< /details >}}

{{< details title="31: تفاوت بین unbuffered و buffered channel چیست؟" open=false >}}

* **Unbuffered channel**: ارسال‌کننده تا زمانی که دریافت‌کننده آماده نباشد بلاک می‌شود.
* **Buffered channel**: ارسال‌کننده می‌تواند تا ظرفیت کانال داده بفرستد بدون بلاک شدن. دریافت‌کننده فقط زمانی بلاک می‌شود که کانال خالی باشد.
  {{< /details >}}

{{< details title="32: چرا نباید از pointer به متغیر loop variable در goroutine استفاده کرد؟" open=false >}}
چون همه goroutineها به همان آدرس حافظه اشاره می‌کنند که مقدار آن در هر iteration تغییر می‌کند و باعث نتایج غیرمنتظره می‌شود. باید از یک متغیر محلی کپی‌شده استفاده کرد.
{{< /details >}}

{{< details title="33: تفاوت := و var در تعریف متغیر چیست؟" open=false >}}

* `:=` فقط در داخل توابع و برای تعریف و مقداردهی اولیه استفاده می‌شود.
* `var` در هر جایی (داخل یا بیرون تابع) می‌تواند استفاده شود و امکان تعریف بدون مقدار اولیه (با مقدار صفر) را می‌دهد.
  {{< /details >}}

{{< details title="34: چه زمانی باید از sync.Once استفاده کنیم؟" open=false >}}
زمانی که می‌خواهیم یک قطعه کد فقط یک‌بار در کل طول عمر برنامه اجرا شود (مثل بارگذاری تنظیمات یا ایجاد اتصال اولیه).
{{< /details >}}

{{< details title="35: تفاوت nil interface و interface حاوی nil چیست؟" open=false >}}

* **nil interface**: مقدار و نوع هر دو nil هستند.
* **interface حاوی nil**: نوع مشخص شده است ولی مقدار nil است. این تفاوت باعث می‌شود مقایسه مستقیم با nil نتیجه متفاوتی بدهد.
  {{< /details >}}

{{< details title="36: آیا گوروتین‌ها به صورت موازی اجرا می‌شوند یا همزمانی (Concurrency) دارند؟" open=false >}}
گوروتین‌ها ذاتاً همزمانی دارند و Go scheduler آنها را روی چندین thread اجرا می‌کند. اگر سیستم چند هسته‌ای باشد و `GOMAXPROCS` متناسب تنظیم شده باشد، گوروتین‌ها می‌توانند موازی اجرا شوند.
{{< /details >}}

{{< details title="37: تفاوت rune و byte چیست؟" open=false >}}

* `byte` معادل `uint8` است و برای نمایش داده‌های خام یا کاراکترهای ASCII استفاده می‌شود.
* `rune` معادل `int32` است و برای نمایش یک کاراکتر یونیکد استفاده می‌شود.
  {{< /details >}}

{{< details title="38: چرا استفاده بیش از حد از defer می‌تواند بر عملکرد تأثیر بگذارد؟" open=false >}}
هر `defer` یک فراخوانی اضافه ایجاد می‌کند که در خروج از تابع اجرا می‌شود. در مسیرهای پر فراخوانی (hot path)، این می‌تواند باعث افت کارایی شود. در این موارد بهتر است منابع را به صورت دستی آزاد کرد.
{{< /details >}}

{{< details title="39: garbage collector در Go چه زمانی اجرا می‌شود؟" open=false >}}
GC در Go به صورت خودکار و همزمان با اجرای برنامه، بر اساس تخصیص حافظه و فشار heap اجرا می‌شود. این فرآیند stop-the-world بسیار کوتاهی دارد تا حافظه غیرقابل دسترسی را آزاد کند.
{{< /details >}}

{{< details title="40: آیا slice به صورت مقادیر (value) پاس داده می‌شود یا ارجاع (reference)؟" open=false >}}
Slice خود یک ساختار کوچک است (pointer به آرایه، طول و ظرفیت) که به صورت value پاس داده می‌شود، ولی اشاره‌گر داخل آن باعث می‌شود به آرایه زیرین ارجاع دهد.
{{< /details >}}

{{< details title="41: چه زمانی استفاده از map به جای slice بهینه‌تر است؟" open=false >}}
وقتی که نیاز به دسترسی سریع بر اساس کلیدهای غیر ترتیبی داریم یا جستجوی O(1) می‌خواهیم. در مقابل، slice برای داده‌های ترتیبی کوچک و قابل پیمایش سریع‌تر است.
{{< /details >}}

{{< details title="42: چگونه می‌توان ظرفیت اولیه یک map را تعیین کرد و چرا این کار مهم است؟" open=false >}}
با استفاده از `make(map[KeyType]ValueType, initialCapacity)` می‌توان ظرفیت اولیه تعیین کرد. این باعث کاهش عملیات rehash و بهبود عملکرد می‌شود.
{{< /details >}}

{{< details title="43: تفاوت select بدون case و select با default چیست؟" open=false >}}

* `select {}` بدون case باعث بلاک شدن دائمی گوروتین می‌شود.
* `select { default: ... }` بلافاصله default را اجرا می‌کند اگر هیچ کانالی آماده نباشد.
  {{< /details >}}

{{< details title="44: چرا باید منابع مانند فایل یا connection را در defer بلافاصله بعد از ایجاد آن ببندیم؟" open=false >}}
این کار تضمین می‌کند که حتی در صورت بروز خطا یا panic، منبع به درستی آزاد شود و از resource leak جلوگیری شود.
{{< /details >}}

{{< details title="45: تفاوت channel جهت‌دار (send-only/receive-only) با channel دوطرفه چیست؟" open=false >}}

* Send-only: `chan<- T` فقط اجازه ارسال دارد.
* Receive-only: `<-chan T` فقط اجازه دریافت دارد.
* دوطرفه: `chan T` می‌تواند هم ارسال و هم دریافت انجام دهد.
  {{< /details >}}

{{< details title="46: چگونه می‌توان benchmark در Go نوشت و اجرا کرد؟" open=false >}}
با نوشتن تابع‌هایی با امضای `func BenchmarkXxx(b *testing.B)` و اجرای `go test -bench=.` می‌توان benchmark گرفت. مقدار `b.N` مشخص می‌کند چند بار تست اجرا می‌شود.
{{< /details >}}

{{< details title="47: sync.Pool چیست و چه کاربردی دارد؟" open=false >}}
`sync.Pool` یک ساختار برای ذخیره و بازیابی موقت اشیاء قابل استفاده مجدد است که فشار روی GC را کاهش می‌دهد. برای داده‌های موقتی که ایجادشان هزینه‌بر است استفاده می‌شود.
{{< /details >}}

{{< details title="48: تفاوت testing.T و testing.B در پکیج testing چیست؟" open=false >}}

* `testing.T` برای تست‌های معمولی و مدیریت خطاها استفاده می‌شود.
* `testing.B` برای benchmark استفاده می‌شود و شامل حلقه اجرای تکرارها (`b.N`) است.
  {{< /details >}}

{{< details title="49: آیا Go از tail call optimization پشتیبانی می‌کند؟" open=false >}}
خیر، Go به صورت رسمی tail call optimization ندارد. بنابراین توابع بازگشتی عمیق می‌توانند منجر به مصرف زیاد stack شوند.
{{< /details >}}

{{< details title="50: چه تفاوتی بین make و append برای ایجاد slice وجود دارد؟" open=false >}}

* `make` یک slice با طول و ظرفیت مشخص ایجاد می‌کند.
* `append` برای افزودن عناصر به slice موجود استفاده می‌شود و در صورت پر بودن ظرفیت، یک آرایه جدید ایجاد می‌کند.
  {{< /details >}}

{{< details title="51: تفاوت بین زمان‌بندی گوروتین‌ها در Go و threadها در سیستم‌عامل چیست؟" open=false >}}
Scheduler گوروتین‌ها در Go به صورت **M\:N** کار می‌کند، یعنی تعداد زیادی گوروتین (M) روی تعداد محدودی thread سیستم‌عامل (N) نگاشت می‌شوند. این مدل سبک‌تر از زمان‌بندی مستقیم threadها توسط سیستم‌عامل است.
{{< /details >}}

{{< details title="52: آیا استفاده از global variable در Go توصیه می‌شود؟ چرا؟" open=false >}}
به طور کلی خیر، چون باعث کاهش تست‌پذیری، افزایش coupling و خطر data race می‌شود. در صورت نیاز، باید با قفل یا سایر روش‌های همگام‌سازی محافظت شود.
{{< /details >}}

{{< details title="53: چرا تغییر دادن slice در یک goroutine می‌تواند روی goroutineهای دیگر تأثیر بگذارد؟" open=false >}}
چون همه sliceها به یک آرایه زیرین اشاره می‌کنند. تغییر عناصر باعث تغییر در آرایه مشترک و در نتیجه مشاهده تغییرات توسط سایر goroutineها می‌شود.
{{< /details >}}

{{< details title="54: تفاوت string و \[]byte در Go چیست؟" open=false >}}
`string` غیرقابل تغییر (immutable) است و تغییر مستقیم آن ممکن نیست. `[]byte` قابل تغییر است و برای عملیات پردازشی روی داده‌ها بهینه‌تر است.
{{< /details >}}

{{< details title="55: چرا map در Go iteration order ثابت ندارد؟" open=false >}}
برای جلوگیری از وابستگی برنامه‌ها به ترتیب کلیدها و بهبود کارایی، Go عمداً ترتیب پیمایش map را تصادفی می‌کند.
{{< /details >}}

{{< details title="56: چه زمانی باید از atomic operations استفاده کرد؟" open=false >}}
زمانی که نیاز به عملیات خواندن/نوشتن thread-safe روی انواع عددی یا آدرس حافظه داریم، بدون استفاده از قفل‌های سنگین مانند Mutex. پکیج `sync/atomic` برای این منظور استفاده می‌شود.
{{< /details >}}

{{< details title="57: تفاوت nil slice با empty slice چیست؟" open=false >}}

* Nil slice: `nil` است، طول و ظرفیت صفر دارد.
* Empty slice: مقدار غیر nil با طول صفر است، ولی به یک آرایه صفر‌ظرفیت اشاره می‌کند.
  {{< /details >}}

{{< details title="58: چرا استفاده از pointer receiver برای struct بزرگ بهینه‌تر است؟" open=false >}}
چون از کپی شدن struct بزرگ جلوگیری می‌کند و تغییرات در method روی نمونه اصلی اعمال می‌شود.
{{< /details >}}

{{< details title="59: تفاوت fallthrough در switch چیست؟" open=false >}}
در Go، caseها به طور پیش‌فرض break می‌شوند. استفاده از `fallthrough` باعث می‌شود اجرای case بعدی بدون بررسی شرط ادامه یابد.
{{< /details >}}

{{< details title="60: چرا باید بعد از استفاده از ticker آن را متوقف کنیم؟" open=false >}}
عدم توقف (`ticker.Stop()`) باعث نشت منابع و ادامه کار goroutine داخلی آن می‌شود.
{{< /details >}}

{{< details title="61: تفاوت new و literal برای struct چیست؟" open=false >}}

* `new(T)` یک pointer به مقدار صفر نوع T برمی‌گرداند.
* `&T{}` نیز یک pointer می‌سازد ولی می‌توان فیلدها را مقداردهی کرد.
  {{< /details >}}

{{< details title="62: چرا حلقه for range روی map ممکن است نتیجه متفاوتی در هر اجرا بدهد؟" open=false >}}
چون ترتیب پیمایش map در Go عمداً تصادفی است تا از وابستگی برنامه به ترتیب کلیدها جلوگیری شود.
{{< /details >}}

{{< details title="63: تفاوت deep copy و shallow copy در Go چیست؟" open=false >}}

* Shallow copy فقط مقادیر سطح اول را کپی می‌کند و referenceها همچنان مشترک می‌مانند.
* Deep copy تمام داده‌ها را بازگشتی کپی می‌کند تا هیچ اشتراک حافظه وجود نداشته باشد.
  {{< /details >}}

{{< details title="64: چه زمانی باید از select با context.Done استفاده کنیم؟" open=false >}}
وقتی که می‌خواهیم عملیات منتظر روی channel را در صورت لغو context یا timeout متوقف کنیم.
{{< /details >}}

{{< details title="65: تفاوت بین runtime.GOMAXPROCS و تعداد گوروتین‌ها چیست؟" open=false >}}
`GOMAXPROCS` حداکثر تعداد threadهای همزمانی که می‌توانند کد Go اجرا کنند را مشخص می‌کند، ولی تعداد گوروتین‌ها می‌تواند بسیار بیشتر باشد و توسط scheduler مدیریت می‌شود.
{{< /details >}}

{{< details title="66: چرا string در Go غیرقابل تغییر است؟" open=false >}}
برای بهینه‌سازی عملکرد و امنیت، string به داده‌های فقط خواندنی اشاره می‌کند. این طراحی اجازه می‌دهد رشته‌ها را به‌طور ایمن بین goroutineها به اشتراک گذاشت بدون نیاز به قفل.
{{< /details >}}

{{< details title="67: تفاوت بین time.Sleep و استفاده از time.After چیست؟" open=false >}}

* `time.Sleep` اجرای گوروتین را برای مدت مشخص متوقف می‌کند.
* `time.After` یک channel برمی‌گرداند که بعد از گذشت مدت زمان مشخص سیگنال ارسال می‌کند و می‌تواند در `select` استفاده شود.
  {{< /details >}}

{{< details title="68: چرا نباید از panic برای کنترل جریان عادی برنامه استفاده کرد؟" open=false >}}
panic برای شرایط غیرمنتظره و خطاهای بحرانی طراحی شده است. استفاده از آن در منطق عادی باعث سختی در خواندن کد، مشکلات تست و مدیریت منابع می‌شود.
{{< /details >}}

{{< details title="69: sync.Cond چیست و چه زمانی استفاده می‌شود؟" open=false >}}
`sync.Cond` ابزاری برای هماهنگ‌سازی پیشرفته است که به goroutineها اجازه می‌دهد تا تا زمان برآورده شدن یک شرط منتظر بمانند و توسط دیگر goroutineها بیدار شوند.
{{< /details >}}

{{< details title="70: تفاوت بین context.Background و context.TODO چیست؟" open=false >}}

* `context.Background` برای شروع زنجیره context در برنامه‌های سطح بالا استفاده می‌شود.
* `context.TODO` زمانی استفاده می‌شود که هنوز مشخص نیست چه contextی باید استفاده شود یا در حال توسعه هستیم.
  {{< /details >}}

{{< details title="71: چرا mapهای بزرگ می‌توانند باعث فشار روی GC شوند؟" open=false >}}
چون map ممکن است مقادیر زیادی حافظه تخصیص دهد و GC باید همه کلیدها و مقادیر را بررسی کند. حذف مقادیر غیرضروری یا بازسازی map می‌تواند فشار را کاهش دهد.
{{< /details >}}

{{< details title="72: تفاوت بین کانال بسته و nil channel چیست؟" open=false >}}

* کانال بسته: خواندن از آن مقدار صفر و وضعیت بسته بودن را برمی‌گرداند. ارسال روی آن panic ایجاد می‌کند.
* nil channel: هر عملیات ارسال یا دریافت روی آن برای همیشه بلاک می‌شود.
  {{< /details >}}

{{< details title="73: چرا استفاده از range روی string با کاراکترهای یونیکد خاص می‌تواند کند باشد؟" open=false >}}
زیرا Go هر بار کاراکترها را به صورت rune دیکد می‌کند که شامل تبدیل UTF-8 به int32 است. برای داده‌های ASCII این تبدیل سریع‌تر است.
{{< /details >}}

{{< details title="74: چه زمانی باید از copy برای slice استفاده کنیم؟" open=false >}}
وقتی که می‌خواهیم داده‌ها را از یک slice به دیگری منتقل کنیم بدون اینکه به آرایه زیرین مشترک اشاره کنند، مخصوصاً برای جلوگیری از مشکلات memory leak یا تغییرات ناخواسته.
{{< /details >}}

{{< details title="75: تفاوت بین defer با تابع معمولی در مدیریت منابع چیست؟" open=false >}}
`defer` تضمین می‌کند که تابع در پایان محدوده اجرا شود، حتی در صورت panic یا بازگشت زودهنگام، ولی تابع معمولی بلافاصله اجرا می‌شود.
{{< /details >}}

{{< details title="76: چرا نباید روی کانال nil عملیات انجام داد؟" open=false >}}
هر ارسال یا دریافت روی کانال nil باعث بلاک شدن بی‌پایان goroutine می‌شود، که معمولاً به صورت ناخواسته deadlock ایجاد می‌کند.
{{< /details >}}

{{< details title="77: تفاوت بین append به slice و append به nil slice چیست؟" open=false >}}
هر دو معتبر هستند. append به nil slice باعث ایجاد slice جدید با ظرفیت مورد نیاز می‌شود.
{{< /details >}}

{{< details title="78: چرا حلقه‌های for بدون شرط در Go می‌توانند مشکل‌ساز شوند؟" open=false >}}
یک حلقه `for {}` بدون شرط توقف می‌تواند CPU را ۱۰۰٪ مشغول کند مگر اینکه شامل عملیات بلاک‌کننده یا sleep باشد.
{{< /details >}}

{{< details title="79: sync.Map چه مزیتی نسبت به map با Mutex دارد؟" open=false >}}
`sync.Map` برای سناریوهایی با خواندن زیاد و نوشتن کم بهینه‌سازی شده و نیازی به قفل دستی ندارد، ولی در همه موارد سریع‌تر از map+Mutex نیست.
{{< /details >}}

{{< details title="80: چرا تبدیل بین \[]byte و string در Go معمولاً باعث کپی می‌شود؟" open=false >}}
برای جلوگیری از تغییر داده‌های string (که immutable است)، Go معمولاً داده‌ها را کپی می‌کند. این رفتار هزینه زمانی و حافظه دارد.
{{< /details >}}

{{< details title="81: چرا استفاده از time.Tick بدون توقف می‌تواند memory leak ایجاد کند؟" open=false >}}
`time.Tick` یک channel بازگشتی ایجاد می‌کند که هرگز متوقف نمی‌شود. برای جلوگیری از نشت منابع، بهتر است از `time.NewTicker` استفاده کرده و در زمان مناسب `Stop()` را فراخوانی کنید.
{{< /details >}}

{{< details title="82: تفاوت بین interface خالی و interface با متد چیست؟" open=false >}}

* Interface خالی (`interface{}`) می‌تواند هر نوعی را نگه دارد.
* Interface با متدها فقط می‌تواند نوع‌هایی را نگه دارد که تمام متدهای تعریف‌شده را پیاده‌سازی کرده باشند.
  {{< /details >}}

{{< details title="83: چرا استفاده از goroutine در حلقه for بدون همگام‌سازی می‌تواند مشکل‌ساز باشد؟" open=false >}}
چون متغیرهای حلقه بین goroutineها به اشتراک گذاشته می‌شوند و مقدارشان در زمان اجرا ممکن است تغییر کند. باید متغیر را به صورت محلی کپی یا با آرگومان به goroutine پاس داد.
{{< /details >}}

{{< details title="84: تفاوت بین make برای map و literal map چیست؟" open=false >}}

* `make(map[Key]Value, cap)` ظرفیت اولیه را مشخص می‌کند.
* Literal map (`map[Key]Value{...}`) بلافاصله داده‌ها را مقداردهی می‌کند.
  {{< /details >}}

{{< details title="85: چرا نباید روی کانالی که چند تولیدکننده دارد بدون هماهنگی close انجام داد؟" open=false >}}
چون ممکن است چند goroutine همزمان تلاش به بستن کانال کنند که باعث panic می‌شود. معمولاً فقط یک تولیدکننده مسئول close است.
{{< /details >}}

{{< details title="86: تفاوت بین string literal با backtick و با quote چیست؟" open=false >}}

* با quote (`"..."`): رشته escape می‌شود و می‌توان از `\n` و غیره استفاده کرد.
* با backtick (`` `...` ``): رشته raw است و escape interpretation انجام نمی‌شود.
  {{< /details >}}

{{< details title="87: چرا garbage collector نمی‌تواند فایل‌های باز را ببندد؟" open=false >}}
GC فقط حافظه را آزاد می‌کند، ولی منابع سیستم‌عاملی مثل فایل‌ها یا socketها را باید به صورت صریح با `Close()` آزاد کرد.
{{< /details >}}

{{< details title="88: تفاوت بین untyped constant و typed constant چیست؟" open=false >}}

* Untyped constant می‌تواند با هر نوع سازگار استفاده شود تا زمانی که مقدارش در محدوده باشد.
* Typed constant نوع مشخص دارد و فقط با همان نوع یا نوع‌های سازگار استفاده می‌شود.
  {{< /details >}}

{{< details title="89: چرا استفاده از recover خارج از defer بی‌اثر است؟" open=false >}}
چون `recover` فقط زمانی panic را می‌گیرد که درون یک defer اجرا شود که در همان گوروتین panic اتفاق افتاده باشد.
{{< /details >}}

{{< details title="90: تفاوت بین range روی slice و روی array چیست؟" open=false >}}

* روی slice: مقدار و اندیس عناصر slice را برمی‌گرداند.
* روی array: مشابه slice، ولی آرایه به صورت کامل پاس داده می‌شود که می‌تواند هزینه‌بر باشد مگر اینکه با reference پاس شود.
  {{< /details >}}

{{< details title="91: چرا capacity slice بعد از append ممکن است چند برابر شود؟" open=false >}}
برای بهینه‌سازی تخصیص حافظه، Go هنگام نیاز به افزایش ظرفیت معمولاً آن را به صورت نمایی (دو برابر یا بیشتر) افزایش می‌دهد.
{{< /details >}}

{{< details title="92: تفاوت بین runtime.Gosched و time.Sleep چیست؟" open=false >}}

* `runtime.Gosched` به scheduler اجازه می‌دهد گوروتین‌های دیگر را اجرا کند بدون توقف زمان مشخص.
* `time.Sleep` گوروتین را برای مدت مشخص متوقف می‌کند.
  {{< /details >}}

{{< details title="93: چرا nil pointer در method receiver ممکن است باعث panic نشود؟" open=false >}}
اگر متد به فیلد یا داده‌ای از struct دسترسی نداشته باشد، حتی اگر receiver nil باشد، panic ایجاد نمی‌شود.
{{< /details >}}

{{< details title="94: تفاوت بین map\[string]struct{} و map\[string]bool چیست؟" open=false >}}

* `map[string]struct{}` حافظه کمتری مصرف می‌کند چون struct خالی صفر بایت است.
* `map[string]bool` نیاز به یک بایت برای مقدار دارد.
  {{< /details >}}

{{< details title="95: چرا استفاده از for-select بدون default می‌تواند کارایی را کاهش دهد؟" open=false >}}
چون گوروتین بلاک می‌ماند تا یکی از caseها آماده شود، که ممکن است باعث استفاده ناکارآمد از CPU شود.
{{< /details >}}

{{< details title="96: تفاوت بین var x = y و x := y چیست؟" open=false >}}

* `var x = y` می‌تواند در سطح package یا تابع استفاده شود.
* `x := y` فقط در داخل توابع مجاز است و همیشه تعریف جدید انجام می‌دهد.
  {{< /details >}}

{{< details title="97: چرا استفاده از init function باید محدود باشد؟" open=false >}}
چون باعث اجرای مخفیانه کد در هنگام بارگذاری پکیج می‌شود و می‌تواند تست و اشکال‌زدایی را سخت کند. بهتر است منطق راه‌اندازی به صراحت در main یا سازنده‌ها باشد.
{{< /details >}}

{{< details title="98: تفاوت بین context.WithCancel و context.WithTimeout چیست؟" open=false >}}

* `WithCancel` فقط با فراخوانی تابع cancel لغو می‌شود.
* `WithTimeout` به صورت خودکار بعد از زمان مشخص لغو می‌شود.
  {{< /details >}}

{{< details title="99: چرا حلقه‌های بازگشتی بدون شرط توقف می‌توانند باعث stack overflow شوند؟" open=false >}}
چون Go tail call optimization ندارد و هر فراخوانی بازگشتی stack را افزایش می‌دهد.
{{< /details >}}

{{< details title="100: تفاوت بین log.Fatal و panic چیست؟" open=false >}}

* `log.Fatal` پیام را چاپ کرده و برنامه را بلافاصله با `os.Exit` متوقف می‌کند.
* `panic` اجرای عادی را متوقف کرده و chain فراخوانی deferها را اجرا می‌کند قبل از توقف.
  {{< /details >}}

{{< details title="101: تفاوت بین unsafe.Pointer و uintptr چیست؟" open=false >}}

* `unsafe.Pointer` برای تبدیل بین انواع اشاره‌گر استفاده می‌شود.
* `uintptr` یک نوع عددی است که آدرس را به عنوان یک عدد ذخیره می‌کند. تبدیل بین آنها باید با احتیاط انجام شود چون ممکن است GC آدرس را جابجا کند.
  {{< /details >}}

{{< details title="102: چرا استفاده از reflect می‌تواند کارایی را کاهش دهد؟" open=false >}}
پکیج `reflect` باعث عملیات‌های زمان اجرا و type checking پویا می‌شود که نسبت به کد معمولی کندتر است.
{{< /details >}}

{{< details title="103: تفاوت بین constant expression و runtime value چیست؟" open=false >}}
Constant expression در زمان کامپایل مشخص می‌شود و می‌تواند در تعریف ثابت‌ها استفاده شود. Runtime value فقط در زمان اجرای برنامه مشخص می‌شود.
{{< /details >}}

{{< details title="104: چرا nil channel برای همگام‌سازی گاهی مفید است؟" open=false >}}
با nil کردن یک channel می‌توانیم آن را از انتخاب در `select` حذف کنیم و رفتار برنامه را به صورت پویا کنترل کنیم.
{{< /details >}}

{{< details title="105: تفاوت بین stack و heap allocation در Go چیست؟" open=false >}}
Stack سریع‌تر و برای داده‌های کوتاه‌مدت استفاده می‌شود. Heap برای داده‌هایی که طول عمر نامعلوم دارند و بین goroutineها به اشتراک گذاشته می‌شوند.
{{< /details >}}

{{< details title="106: چرا slicing از یک array بزرگ می‌تواند باعث نگه‌داشتن حافظه اضافی شود؟" open=false >}}
چون slice به آرایه زیرین اشاره می‌کند و GC نمی‌تواند آرایه بزرگ را آزاد کند تا زمانی که slice زنده باشد.
{{< /details >}}

{{< details title="107: تفاوت بین break و continue در حلقه چیست؟" open=false >}}

* `break` حلقه را کامل متوقف می‌کند.
* `continue` فقط iteration جاری را رد کرده و iteration بعدی را شروع می‌کند.
  {{< /details >}}

{{< details title="108: چرا append ممکن است باعث تغییر آدرس حافظه slice شود؟" open=false >}}
وقتی ظرفیت پر شود، append یک آرایه جدید می‌سازد و داده‌ها را کپی می‌کند که آدرس حافظه متفاوت خواهد داشت.
{{< /details >}}

{{< details title="109: تفاوت بین go vet و go fmt چیست؟" open=false >}}

* `go fmt` کد را قالب‌بندی می‌کند.
* `go vet` مشکلات احتمالی کد را شناسایی می‌کند ولی تغییرات ظاهری ایجاد نمی‌کند.
  {{< /details >}}

{{< details title="110: چرا باید از io.Reader و io.Writer در طراحی API استفاده کنیم؟" open=false >}}
برای انعطاف‌پذیری بیشتر و امکان استفاده از منابع مختلف (فایل، شبکه، حافظه) بدون تغییر کد.
{{< /details >}}

{{< details title="111: تفاوت بین os.Exit و return در main چیست؟" open=false >}}
`os.Exit` فوراً برنامه را متوقف می‌کند و deferها اجرا نمی‌شوند، ولی `return` اجازه اجرای deferها را می‌دهد.
{{< /details >}}

{{< details title="112: چرا map نمی‌تواند به عنوان کلید slice داشته باشد؟" open=false >}}
چون slice قابل مقایسه نیست و hash آن ثابت نیست. فقط انواع قابل مقایسه (comparable) می‌توانند کلید باشند.
{{< /details >}}

{{< details title="113: تفاوت بین io.Pipe و channel چیست؟" open=false >}}
`io.Pipe` برای اتصال مستقیم بین io.Reader و io.Writer استفاده می‌شود، ولی channel برای ارسال هر نوع داده بین goroutineها استفاده می‌شود.
{{< /details >}}

{{< details title="114: چرا نباید از pointer به local variable بعد از پایان تابع استفاده کنیم؟" open=false >}}
چون متغیر ممکن است از بین برود یا آدرسش تغییر کند، که باعث رفتار غیرقابل پیش‌بینی می‌شود.
{{< /details >}}

{{< details title="115: تفاوت بین testing.Short و تست‌های عادی چیست؟" open=false >}}
`testing.Short()` می‌تواند در تست‌ها استفاده شود تا تست‌های طولانی در حالت کوتاه اجرا نشوند.
{{< /details >}}

{{< details title="116: چرا باید برای عملیات‌های سنگین I/O از buffered channel استفاده کرد؟" open=false >}}
برای جلوگیری از بلاک شدن تولیدکننده یا مصرف‌کننده و افزایش throughput.
{{< /details >}}

{{< details title="117: تفاوت بین go build و go install چیست؟" open=false >}}

* `go build` فایل اجرایی را در دایرکتوری جاری می‌سازد.
* `go install` آن را در مسیر bin در `$GOPATH` یا `$GOBIN` نصب می‌کند.
  {{< /details >}}

{{< details title="118: چرا استفاده از strings.Builder به جای + برای رشته‌ها بهتر است؟" open=false >}}
چون از تخصیص‌های متعدد جلوگیری می‌کند و حافظه را بهینه‌تر مدیریت می‌کند.
{{< /details >}}

{{< details title="119: تفاوت بین cap و len در array چیست؟" open=false >}}
برای array هر دو برابر طول آرایه هستند، ولی برای slice ممکن است متفاوت باشند.
{{< /details >}}

{{< details title="120: چرا باید بعد از استفاده از bufio.Writer حتماً Flush کرد؟" open=false >}}
چون داده‌ها در بافر ذخیره می‌شوند و تا زمانی که Flush نشوند به مقصد نهایی ارسال نمی‌شوند.
{{< /details >}}

{{< details title="121: تفاوت بین import \_ و import alias چیست؟" open=false >}}

* `import _` فقط برای اجرای init پکیج است.
* Import alias برای استفاده از نام متفاوت جهت ارجاع به پکیج.
  {{< /details >}}

{{< details title="122: چرا باید از context در عملیات‌های شبکه استفاده کنیم؟" open=false >}}
برای امکان لغو یا timeout در صورت طولانی شدن عملیات.
{{< /details >}}

{{< details title="123: تفاوت بین interface value و concrete value چیست؟" open=false >}}
Interface value شامل نوع و مقدار واقعی است، concrete value فقط مقدار واقعی است.
{{< /details >}}

{{< details title="124: چرا استفاده از constant برای magic number بهتر است؟" open=false >}}
باعث خوانایی بهتر، کاهش خطا و امکان تغییر راحت در آینده می‌شود.
{{< /details >}}

{{< details title="125: تفاوت بین log.Println و fmt.Println چیست؟" open=false >}}
`log.Println` علاوه بر چاپ پیام، timestamp هم اضافه می‌کند.
{{< /details >}}

{{< details title="126: چرا sync.WaitGroup باید با Add قبل از اجرای goroutine استفاده شود؟" open=false >}}
برای جلوگیری از شرایطی که goroutine قبل از افزایش شمارنده شروع شود و WaitGroup صفر بماند.
{{< /details >}}

{{< details title="127: تفاوت بین path و filepath در Go چیست؟" open=false >}}

* `path` برای مسیرهای URL و forward slash استفاده می‌شود.
* `filepath` برای مسیرهای سیستم فایل وابسته به سیستم‌عامل استفاده می‌شود.
  {{< /details >}}

{{< details title="128: چرا استفاده از http.Client پیش‌فرض می‌تواند مشکل‌ساز باشد؟" open=false >}}
چون connectionها را cache نمی‌کند و ممکن است باعث نشت اتصال شود مگر اینکه timeout یا Transport سفارشی تعریف شود.
{{< /details >}}

{{< details title="129: تفاوت بین sync.Mutex و sync.RWMutex در خواندن زیاد چیست؟" open=false >}}
`sync.RWMutex` اجازه می‌دهد چندین خواننده همزمان کار کنند، ولی Mutex فقط یک قفل کلی می‌دهد.
{{< /details >}}

{{< details title="130: چرا panic در گوروتین جداگانه باید به صورت جداگانه recover شود؟" open=false >}}
چون recover فقط panic در همان goroutine را می‌گیرد.
{{< /details >}}

{{< details title="131: تفاوت بین os.Create و os.OpenFile چیست؟" open=false >}}

* `os.Create` فایل را با truncate ایجاد یا باز می‌کند.
* `os.OpenFile` کنترل کامل روی mode و flagها می‌دهد.
  {{< /details >}}

{{< details title="132: چرا context باید به صورت اولین آرگومان به توابع پاس داده شود؟" open=false >}}
این یک قرارداد استاندارد Go است که خوانایی و یکپارچگی API را بهبود می‌دهد.
{{< /details >}}

{{< details title="133: تفاوت بین errors.Is و errors.As چیست؟" open=false >}}

* `errors.Is` برای بررسی تطابق با یک خطای خاص استفاده می‌شود.
* `errors.As` برای استخراج و استفاده از نوع خاص خطا.
  {{< /details >}}

{{< details title="134: چرا نباید از defer در حلقه‌های شدیداً پرتکرار استفاده کرد؟" open=false >}}
چون هر defer تا پایان تابع ذخیره می‌شود و باعث سربار می‌شود.
{{< /details >}}

{{< details title="135: تفاوت بین nil slice و empty slice در JSON چیست؟" open=false >}}
Nil slice به `null` سریالایز می‌شود، ولی empty slice به `[]`.
{{< /details >}}

{{< details title="136: چرا map در Go به صورت داخلی rehash انجام می‌دهد؟" open=false >}}
برای حفظ کارایی O(1) در دسترسی و جلوگیری از افزایش بیش از حد load factor.
{{< /details >}}

{{< details title="137: تفاوت بین encoding/json و jsoniter چیست؟" open=false >}}
`jsoniter` سریع‌تر و انعطاف‌پذیرتر است ولی encoding/json رسمی و پایدار است.
{{< /details >}}

{{< details title="138: چرا باید کانال را فقط توسط ارسال‌کننده ببندیم؟" open=false >}}
برای جلوگیری از panic ناشی از ارسال به کانال بسته.
{{< /details >}}

{{< details title="139: تفاوت بین context.WithValue و متغیر global چیست؟" open=false >}}
Context برای داده‌های کوتاه‌مدت مرتبط با درخواست استفاده می‌شود، global برای داده‌های بلندمدت.
{{< /details >}}

{{< details title="140: چرا time.AfterFunc می‌تواند به جای time.Sleep مفید باشد؟" open=false >}}
چون به صورت asynchronous عمل می‌کند و callback مشخصی را بعد از مدت معین اجرا می‌کند.
{{< /details >}}

{{< details title="141: تفاوت بین select با یک case و استفاده مستقیم از channel چیست؟" open=false >}}
`select` حتی با یک case هم امکان اضافه کردن default یا caseهای دیگر را در آینده می‌دهد و ساختار کد را منعطف‌تر می‌کند، ولی دریافت مستقیم ساده‌تر است.
{{< /details >}}

{{< details title="142: چرا استفاده از buffer بزرگ‌تر در bufio.Reader می‌تواند کارایی را بهبود دهد؟" open=false >}}
چون تعداد فراخوانی‌های سیستم‌عاملی (syscall) را کاهش می‌دهد و داده‌ها را یک‌جا می‌خواند.
{{< /details >}}

{{< details title="143: تفاوت بین defer با anonymous function و با نام تابع چیست؟" open=false >}}
در defer با anonymous function می‌توان پارامترها را در زمان اجرای defer ارزیابی کرد، ولی در تابع نام‌دار پارامترها در لحظه تعریف defer ارزیابی می‌شوند.
{{< /details >}}

{{< details title="144: چرا از sync/atomic برای شمارنده‌ها استفاده می‌شود؟" open=false >}}
چون عملیات‌های اتمیک بدون نیاز به Mutex انجام می‌شوند و سرعت بیشتری دارند.
{{< /details >}}

{{< details title="145: تفاوت بین byte buffer و byte slice چیست؟" open=false >}}
Byte buffer (مثل bytes.Buffer) امکانات بیشتری مثل نوشتن و خواندن با رشد خودکار دارد، ولی byte slice ساده‌تر است و امکانات اضافی ندارد.
{{< /details >}}

{{< details title="146: چرا استفاده از fallthrough باید محدود باشد؟" open=false >}}
چون می‌تواند باعث اجرای غیرمنتظره caseها شود و خوانایی کد را کاهش دهد.
{{< /details >}}

{{< details title="147: تفاوت بین log.Fatal و os.Exit چیست؟" open=false >}}
`log.Fatal` قبل از توقف پیام خطا را چاپ می‌کند، ولی `os.Exit` فقط برنامه را متوقف می‌کند.
{{< /details >}}

{{< details title="148: چرا map در Go به‌صورت همزمان توسط چند goroutine قابل استفاده نیست؟" open=false >}}
چون پیاده‌سازی داخلی map thread-safe نیست و همزمانی بدون قفل باعث panic می‌شود.
{{< /details >}}

{{< details title="149: تفاوت بین os.Getenv و os.LookupEnv چیست؟" open=false >}}

* `os.Getenv` رشته را برمی‌گرداند و اگر وجود نداشته باشد مقدار خالی می‌دهد.
* `os.LookupEnv` مقدار و وضعیت وجود داشتن را برمی‌گرداند.
  {{< /details >}}

{{< details title="150: چرا range روی channel تا زمان بسته شدن ادامه دارد؟" open=false >}}
چون range از channel تا زمانی که همه مقادیر خوانده و کانال بسته نشود، بلاک می‌ماند.
{{< /details >}}
