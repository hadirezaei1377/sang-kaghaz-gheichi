                                                    سنگ کاغذ قیچی

میتونی به تعداد دلخواهت با ترمینال بازی کنی ... فقط کافیه این برنامه رو روی سیستم لوکال خوت ران کنی...

go run main.go

-redis
-sqlite
-golang

یادت باشه ابتدا باید وابستگی‌های پروژه رو نصب کنی 

- **Go**:      (https://golang.org/doc/install) 
- **Redis**:   (https://redis.io/download) 
- **SQLite**:  (https://www.sqlite.org/download.html)

- go get github.com/go-redis/redis/v8
- go get gorm.io/gorm
- go get gorm.io/driver/sqlite


این دستور بازی را اجرا می‌کند و از شما می‌خواهد که یکی از گزینه‌های "rock"، "paper" یا "scissors" را وارد کنید. سیستم به طور رندوم یکی از گزینه‌ها را انتخاب می‌کند و نتیجه را نمایش می‌دهد. بازی ادامه می‌یابد تا یکی از بازیکنان به امتیاز ۳ برسد.

### مشاهده امتیازات مرحله قبل

برای مشاهده امتیازات مرحله قبل، می‌توانید از Redis استفاده کنید. در کد ذخیره‌سازی Redis، امتیازات با نام بازیکنان به عنوان کلید ذخیره می‌شود. برای مشاهده امتیازات، در ترمینال Redis CLI را باز کنید و دستورات زیر را اجرا کنید:

```sh
redis-cli
```

سپس برای مشاهده امتیاز بازیکن و کامپیوتر، از دستورات زیر استفاده کنید:

```sh
GET User
GET Computer
```

این دستورات امتیازات بازیکن و کامپیوتر را نمایش می‌دهد.

### مشاهده اطلاعات بازی‌های گذشته

اطلاعات بازی‌های گذشته در SQLite ذخیره می‌شود. برای مشاهده این اطلاعات، می‌توانید از SQLite CLI استفاده کنید:

```sh
sqlite3 games.db
```

سپس برای مشاهده تمامی رکوردهای بازی، دستور زیر را اجرا کنید:

```sh
SELECT * FROM game_records;
```

این دستور تمام بازی‌های ثبت شده در دیتابیس را نمایش می‌دهد.
