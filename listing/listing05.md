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
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Выведет error, тк тест возвращает класс кастомной ошибки. 
Этот класс находится в интерфейсе ошибки в main из-за чего 
этот интерфейс уже не является nil.
```