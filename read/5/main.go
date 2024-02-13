package main

// Создали нашу кастомную ошибку
type customError struct {
	msg string
}

// Имплементировали необходимый метод для удов. типу error
func (e *customError) Error() string {
	return e.msg
}

// Возвращаем nil, который, тем не менее, имеет тип customError, т.к является указателем на структуру customError
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
