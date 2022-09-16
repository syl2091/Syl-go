package main

import "fmt"

//类型中嵌入功能
//聚合
type Log struct {
	msg string
}
type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "LEGE"
	c.log = new(Log)
	c.log.msg = "1- yes we can"
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) Log() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	fmt.Println(c)
	return c.log
}
