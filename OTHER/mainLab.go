package main

import (
	"fmt"
)

func main() {

/*
	//Stack
					
	stack := Stack{}

	stack.Spush(1)
	stack.Spush(2)
	stack.Spush(3)
	//fmt.Println(stack.Peek()) // 3
	fmt.Println(stack.Spop()) // 3
	fmt.Println(stack.Spop()) // 2
	//fmt.Println(stack.IsEmpty()) // true
	//fmt.Println(stack.Peek()) // 1
	fmt.Println(stack.Spop()) // 1
	fmt.Println(stack.Spop())
	//fmt.Println(stack.IsEmpty()) // true


	//Queue
	queue := Queue{}

	queue.Qadd("1")
	queue.Qadd("ab")
	queue.Qadd("2")
	queue.Qadd("cd")
	

	fmt.Println("Размер:", queue.Qsize())

	fmt.Println("Извлечение элементов:")
	for !queue.IsEmpty() {
		fmt.Println(queue.Qdell())
	}

					//Array
    // Создаем новый MyArray с размером 5
    size:=5
    arr := NewMyArray(size)

    // Вызываем функцию Set для установки значений
    arr.Aset(0, "1")
    arr.Aset(1, "2")
    arr.Aset(2, "3")
    arr.Aset(3, "4")
    arr.Aset(4, "5")

    // Вызываем функцию Get для получения значений
    fmt.Println(arr.Aget(0)) //"1"
    fmt.Println(arr.Aget(1)) //"2"
    fmt.Println(arr.Aget(2)) //"3"
	fmt.Println(arr.Aget(3)) //"4"
	fmt.Println(arr.Aget(4)) //"5"
    // Вызываем функцию Del для удаления значений
    deletedValue := arr.Adel(1)
    fmt.Println(deletedValue)
    fmt.Println("")
    fmt.Println(arr.Aget(1)) // Вывод: ""

    // Создаем экземпляр HashMap
	hmap := HashMap{}

	// Вставляем пару ключ-значение
	err := hmap.Hadd("n1", "a")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Hadd("n2", "b")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Hadd("n3", "c")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	// Выводим все элементы через цикл
	for i := 0; i < len(hmap.table); i++ {
		if hmap.table[i] != nil {
			fmt.Printf("Ключ: %s, Значение: %s\n", hmap.table[i].key, hmap.table[i].value)
		}
	}


	//LsList
	list1 := NewLinkedList()

    list1.SLaddS("1")
    list1.SLaddS("2")
    list1.SLaddE("3")
    list1.SLaddE("4")

    fmt.Println("Список:")
    list1.SLdisplay()

    list1.SLdel("2")
    list1.SLdel("4")

    fmt.Println("Список после удаления элементов:")
    list1.SLdisplay()

	    // Создание нового двусвязного списка

	list2 := NewDoublyLinkedList()

		// Добавление элементов
	list2.Dladd("1")
	list2.Dladd("2")
	list2.Dladd("3")
	list2.Dladd("4")
	
		// Вывод содержимого
	fmt.Println("Содержимое списка:")
	list2.DLdisplay()
	
		// Удаление элемента
	list2.Dldel(2)
	fmt.Println("Содержимое списка после удаления элемента 2:")
	list2.DLdisplay()
*/
    list1 := LinkedList{}
    stack := Stack{}
    queue := Queue{}
    //set:=NewSet()
    //array := MyArray{}
    //var array *MyArray
    array:= NewMyArray(10)
    hmap := HashMap{} // Создаем экземпляр хэштаблицы
    list2 :=DoublyLinkedList{}
    //tree1:= BSTNode{}
    var tree *BSTNode
    // index:=0
    //tree = BSTadd(tree, index)
    for  {   
        var input string
        var index int
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
        case "APUSH":
            fmt.Scanf("%s", &input)
            fmt.Println(array.ARadd(input))
        case "AGET":
            fmt.Scanf("%d", &index)
            fmt.Println(array.Aget(index))
        case "ADEL":
            fmt.Scanf("%d", &index)
            array.Adel(index)
        case "ASET":
            fmt.Scanf("%d %s",&index,&input)
            array.Aset(index,input)
        case "APRINT":
            array.PrintArray()
        case "HADD":
            // Добавление элемента в хэштаблицу
            fmt.Scanf("%s", &input) // Считываем ключ
            fmt.Scanf("%s", &index2) // Считываем значение
            fmt.Println(hmap.Hadd(input, index2))

        case "HGET":
            // Получение значения из хэштаблицы
            fmt.Scanf("%s", &input) // Считываем ключ
            fmt.Println(hmap.Hget(input))

        case "HDEL":
            // Удаление элемента из хэштаблицы
            fmt.Scanf("%s", &input) // Считываем ключ
            hmap.Hdel(input)
        case "LSADD":
            fmt.Scanf("%s", &input)
            list1.SLaddS(input)
        case "LSEADD":
            fmt.Scanf("%s", &input)
            list1.SLaddE(input)
        case "LSDEL":
            fmt.Scanf("%s", &input)
            list1.SLdel(input)
        case "LSDISPLAY":
            list1.SLdisplay()
        case "LSFIND":
            fmt.Scanf("%s", &input)
            fmt.Println(list1.SLsearch(input))
        case "LSDELS":
            list1.SLremoveFront()
        case "LSDELE":
            list1.SLremoveEnd()
        case "DLADDE":
            fmt.Scanf("%s", &input)
            list2.DladdE(input)
        case "DLDEL":
            fmt.Scanf("%s", &input)
            list2.Dldel(input)
        case "DLADD":
            fmt.Scanf("%s", &input)
            list2.DlAdd(input)
        case "DLDISPLAY":
            list2.DLdisplay()
        case "DLDELH":
            list2.DlRemoveFirst()
        case "DLDELE":
            list2.DlRemoveLast()
        case "DLFIND":
            fmt.Scanf("%s", &input)
            fmt.Println(list2.DlFind(input))
        case "BSTADD":
            fmt.Scanf("%d", &index)
            tree =BSTadd(tree, index)
        case "BSTDISPLAY":
            PrintTree(tree, "", true)
        case "BSTFIND":
            fmt.Scanf("%d", &index)
            fmt.Println(BstFind(tree,index))
        case "BSTDEL":
            fmt.Scanf("%d",&index)
            tree = BStdel(tree,index)
    }
}
}