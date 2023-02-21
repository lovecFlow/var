package main

import "fmt"

/* Определение перемнной: var_имя_тип ключевые слова использовать нельзя. Простейшее определение переменной :*/ var hello string /* Есть возможность объявить сразу несколько перемнных. Это делается через запятую: */ var a, b, c string /* Определили 3 перемнных, типом которых является string. 

/* После определения можно присвоить какое-либо значение, соотвествующие её типу: */ 
func main () { 
	var hello string 
	hello = "Hello World"
	fmt.Println(hello)
}  

/* Переменная имеет тип string и мы можем записать строку. Тут она хранит значение Hello World. Go - регистрозависимый язык, а занчит, что "hello" и "Hello" представляют разные переменные. 

Инициализация представляет собой прием, при котором можно сразу объявить начальное значение переменной: */ 

func main () { 
	var hello string = "Hello World"
	fmt.Println(hello) 
}

/* Если мы хотим определить и инициализировать сразу несколько переменных, то оборачиваем в скобки: */

func main() { 
	var (
		name string = "Tom"
		age int = 27 
	)
	
	fmt.Println(name) //Tom
	fmt.Println(age) //27 
}

/* Переменные (var) могут многократно меняться: */ 

func main() { 
	var hello string = "Hello World"
	fmt.Println(hello) //Hello World 

	hello = "Hello Go"
	fmt.Println(hello) //Hello Go 

	hello = "GOGOGOGO"
	fmt.Println(hello) //GOGOGOGO 
}

/* Возможен краткий вариант определения переменной: */ 

func main() { 
	name := "Tom"
	fmt.Println(name)
} 

/* В этом случае тип данных не указывается, он выводится автоматически из присваемого значения. 