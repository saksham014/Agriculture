package main

import "fmt"
	



 type struct person{
	name string,
	age int,
 }
 type struct employee{
	person,

}

func(e employee)isMajor()(bool){
	if(e.age>=18){
		return true
	}else{
		return false;
	}
}

func (p *person) setAge(num)(int){
	p.age= num
}

func main() {
	countries := make(map[string]map[string]string)
	uttarpradesh := map[string]string{"Uttar Pradesh": "Lucknow"}
	countries["India"] = uttarpradesh

	for key, value := range countries {
		fmt.Println(key, "=>", value)
	}

	// value, ok := countries["India"]["Uttar Pradesh"]
	// fmt.Println(value, ok)

	p:=person("saksham")



}
