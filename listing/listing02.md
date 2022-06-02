Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

в функции test мы вернули переменную неявно, поэтому defer смог повлиять на x перед возвратом
в функции anotherTest мы вернули переменную явно, передав значение до выполнения defer

вызовы defer выполняются по завершении функции в обратном порядке.

```
