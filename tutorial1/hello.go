/*
學習重點:
1. 執行方式: go run hello.go 
2. 打包方式: go build hello.go
3. func的 "{" 不能放在單獨的行上，如:
==================================
func main
{
	//這樣子會報錯
}
==================================
4. Println若需要分行，可隔行編寫，如:
==================================
fmt.Println("Hello World! 您好!")
fmt.Println("第二行")
fmt.Println("第三行")
==================================
5. 與Javascript一樣的註解寫法
6. 自符串接用"+"號
7. 宣告變數的方法
8. 變數的結合及輸出方法
9. 用Scanf函式來監聽使用者輸入
*/

//包
package main

//引入
import "fmt"

//函式
func main() {
	//列印
	fmt.Println("Hello World! 您好!")
	fmt.Println("第二行")
	fmt.Println("第三行")
	fmt.Println("一起來學"+"GO"+"語言")

	//宣告
	var age int = 37 		//宣告數字
	var sex string = "男性" //宣告字串
	var birth = "1985-03" 	//宣告但不特別指定型別

	//透過Printf來組合變數輸出，並透過"\n"進行斷行
	fmt.Printf("直接輸出: 我是%d歲的%s,出生於%s的某一天\n", age, sex, birth)

	//另外也可以透過Sprintf來做格式化字符串
	var format = "格式化輸出: 我是%d歲的%s,出生於%s的某一天"
	var apply = fmt.Sprintf(format, age, sex, birth)
    fmt.Println(apply)

	//留著視窗等待關閉指令輸入
	fmt.Println("關閉式窗請鍵入[q]後，按下輸入")
	fmt.Scanf("q")
}
