# مقدمه

سلام متن فارسی جهت تست زبان گو فارسی

## لورم اپیسوم

لورم ایپسوم متن ساختگی با تولید سادگی نامفهوم از صنعت چاپ، و با استفاده از طراحان گرافیک است، چاپگرها و متون بلکه روزنامه و مجله در ستون و سطرآنچنان که لازم است، و برای شرایط فعلی تکنولوژی مورد نیاز، و کاربردهای متنوع با هدف بهبود ابزارهای کاربردی می باشد، کتابهای زیادی در شصت و سه درصد گذشته حال و آینده، شناخت فراوان جامعه و متخصصان را می طلبد، تا با نرم افزارها شناخت بیشتری را برای طراحان رایانه ای علی الخصوص طراحان خلاقی، و فرهنگ پیشرو در زبان فارسی ایجاد کرد، در این صورت می توان امید داشت که تمام و دشواری موجود در ارائه راهکارها، و شرایط سخت تایپ به پایان رسد و زمان مورد نیاز شامل حروفچینی دستاوردهای اصلی، و جوابگوی سوالات پیوسته اهل دنیای موجود طراحی اساسا مورد استفاده قرار گیرد.

```go

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Model interface {
	Imp()
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Address struct {
	Address string `json:"address"`
	ZipCode string `json:"zip_code"`
}

type Phone struct {
	Mobile string `json:"mobile"`
	Home   string `json:"home"`
	Work   string `json:"work"`
}

func (u *User) Imp()    {}
func (u *Address) Imp() {}
func (u *Phone) Imp()   {}

func main() {
	u := &User{Id: 1, Name: "Javad", Age: 29}
	a := &Address{Address: "street 1", ZipCode: "W12345"}
	p := &Phone{Home: "+1123456789", Mobile: "+12233445566", Work: "+3123454645"}

	Print(u)
	Print(a)
	Print(p)

	user, err := StructToMap(u)
	if err != nil {
		log.Fatalln(err)
	}

	address, err := StructToMap(a)
	if err != nil {
		log.Fatalln(err)
	}

	phone, err := StructToMap(p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user, address, phone)
}

func StructToMap(m Model) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &model); err != nil {
		return nil, err
	}
	return model, nil
}

func Print(m Model) {
	fmt.Println(m)
}

```