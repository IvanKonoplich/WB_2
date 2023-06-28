Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
nil, false
Это происходит потому что в интерфейсе ошибки уже сохранен тип os.PathError 
хотя его конкретное значение и является nil, сам интерфейс ошибки nil 
уже не является, тк в нем указан тип данных.
```