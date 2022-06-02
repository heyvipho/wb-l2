Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/tmp/sandbox477242829/prog.go:11 +0xa8


Горутина завершила свою работу и перестала отправлять значения в канал. А чтение из канала бесконечное и зависло навсегда, поэтому произошла данная ошибка.

```
