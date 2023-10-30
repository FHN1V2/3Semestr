package main

import (
	"fmt"
)

func main() {
//list1 := LinkedList{}
stack := Stack{}
queue := Queue{}
set:=NewSet()

hmap := HashMap{} // Создаем экземпляр хэштаблицы

for  {   
	var input string
	var index2 string
	fmt.Scanf("%s", &input)
	switch input {
	case "SPUSH":
		fmt.Scanf("%s", &input)
		stack.Spush(input)
	case "SPOP":
		fmt.Println(stack.Spop())
	case "QPUSH":
		fmt.Scanf("%s", &input)
		queue.Qadd(input)
	case "QPOP":
		fmt.Println(queue.Qdell())
	case "HADD":
		// Добавление элемента в хэштаблицу
		fmt.Scanf("%s", &input) // Считываем ключ
		fmt.Scanf("%s", &index2) // Считываем значение
		hmap.Hadd(input, index2)
	case "HGET":
		// Получение значения из хэштаблицы
		fmt.Scanf("%s", &input) // Считываем ключ
		fmt.Println(hmap.Hget(input))
	case "HDEL":
		// Удаление элемента из хэштаблицы
		fmt.Scanf("%s", &input) // Считываем ключ
		hmap.Hdel(input)
	case "SETPUSH":
		fmt.Scanf("%s", &input)
		set.SetAdd(input)  
	case "SETDEL":
		fmt.Scanf("%s", &input)
		set.SetRemove(input)
	case "SETPRINT":
		set.SetPrint()
	}
}
}
