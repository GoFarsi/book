---
title: "تاریخچه"
type: chapter
weight: 3
---

زبان گو (Go) به عنوان یک زبان کامپایلری و تایپ استاتیک در نوامبر سال ۲۰۰۹ بصورت عمومی معرفی شد که توسط شرکت گوگل توسعه داده شده است و بسیاری از اعضای تیم طراحی و توسعه زبان گو [[Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson), [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike), [Robert Griesemer](https://en.wikipedia.org/wiki/Robert_Griesemer)] سال ها تجربه در زمینه تحقیق و توسعه زبان برنامه نویسی دارند.


زبان گو دارای یک سری ویژگی های منحصر به فرد می باشد و برخی از ویژگی هایش از سایر زبان ها الهام گرفته شده است :

-   پشتیبانی از برنامه نویسی  {{< tooltip text="همزمانی note="con currency" >}} بصورت Built-in 
-    {{< tooltip text="گوروتین note="goroutine" >}} ها امکان اجرای همزمانی توابع را فراهم می کند
-   گوروتین ها واقعا خیلی سبک هستند به طوری که شما می توانید هزاران گوروتین را همزمان در سخت افزارهای بدون هیچ مشکل عملکردی اجرا کنید.
-   {{< tooltip text="کانال note="channel" >}}  ها برپایه مدل CSP می باشد که امکان همگام سازی داده ها بین گوروتین ها را فراهم می کند.
-   پشتیبانی از تایپ های map و slice
-   امکان پیاده سازی {{< tooltip text="پلی مورفیسم note="polymorphism" >}} با استفاده از  {{< tooltip text="اینترفیس note="interface" >}} ها
-   پوینترها
-    {{< tooltip text="کلوژر note="closures" >}} توابع (یک تابع معمولی که داخل یک تابع دیگر به عنوان  {{< tooltip text="بازگشت note="return" >}} تعریف می شود)
-   {{< tooltip text="متدها note="method" >}}
-   امکان defer برای تعویق فراخوانی یک تابع
-   قابلیت  {{< tooltip text="جاسازی note="embedding" >}} تایپ ها
-   ایمنی حافظه در زبان گو
-   قابلیت  {{< tooltip text="زباله جمع کن note="garbage collector" >}} خودکار
-   سازگاری کامل با انواع پلتفرم ها نظیر [linux, windows, mac, AIX, android, freeBSD] جهت توسعه و کامپایل
-   امکان Cross-compile با این امکان می توانید در هر پلتفرمی برای سایر پلتفرم ها کامپایل کنید
-   پشتیبانی از  {{< tooltip text="جنریک note="generics" >}} یا تایپ پارامتر (از نسخه ۱.۱۸)
-   تست نویسی آسان
-    {{< tooltip text="اینترفیس note="interface" >}} و  {{< tooltip text="رفلکشن note="reflection" >}}
-   زبان گو مثل سایر زبان ها نظیر c, cpp یا java نیاز به نقطه ویرگول (Semicolons) ندارد و به نسبت زبان هایی که معرفی کردیم پرانتز کمتری استفاده می کند و همچنین ظاهر سینتکس گو خیلی خواناتر و راحتر می باشد.


