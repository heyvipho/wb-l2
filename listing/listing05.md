Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil { // err != (*customError)(nil)
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Ситуация аналогична с заданием №3. [nil, nil] != [*customError, nil]
```
