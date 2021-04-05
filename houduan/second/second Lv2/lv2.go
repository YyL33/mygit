package main
import"fmt"

func Receiver(v interface{}){
switch v.(type){
case int64:
	fmt.Println("类型：byte")
case rune:
	fmt.Println("类型：rune")
case string:
	fmt.Println("类型：string")
case []interface{}:
fmt.Println("类型：[]interface{}")
default:
	fmt.Println("类型：未知类型")
}
}
func main(){
	var s1 int64=255
	var s2 rune
	var s3 string="golang"
	var s4 []interface{}

Receiver(s1)
Receiver(s2)
Receiver(s3)
Receiver(s4)
}
