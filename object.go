package main

import (
	"fmt"
	"database/sql"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"os"
	"log"
	"bytes"
	"time"
	"sort"
)

type ByAge []Person

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s->%d,", p.Name, p.Age)
}

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age > a[j].Age }

func main()  {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	/*	toInsert()
		obj()
		connRedis()*/
	//jsonT()
	//timex()
}


//定义结构体
type People struct {
	a int
	str string
}

//链接mysql
func toInsert()  {
	db, err := sql.Open("mysql", "root:ba199035@/test")
	result, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (?, ?)",
		"ddd",
		22,
	)
	fmt.Println(result)
	fmt.Println(err)

}


//调用结构体
func obj()  {
	p := new(People)
	p.a = 1
	p.str = "hello"
	fmt.Printf("%+v\n", *p) //  {a:1 str:hello}
}

//链接Redis
func connRedis()  {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	v, err := c.Do("SET", "name", "redxbbbd")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	//v, err = redis.String(c.Do("GET", "name"))
	v, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)
}

type Road struct {
	Name   string
	Number int
}

//json 输出
func jsonT()  {
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatalln(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}
	out.WriteTo(os.Stdout)

}

//时间处理
func timex()  {

	fmt.Println(time.Now())
	return

}