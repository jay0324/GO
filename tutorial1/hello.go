/*
說明:
1. 執行方式: go run hello.go 
2. 打包方式: go build hello.go
3. func的 "{" 不能放在單獨的行上，如:
==================================
func main
{
	//這樣子會報錯
}
==================================
*/

//包
package main

//引入
import "fmt"

//函式
func main() {
	//列印
	fmt.Println("Hello World! 您好!")
}
