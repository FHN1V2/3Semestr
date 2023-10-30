package main

import "fmt"

type MyArray struct {
	data []string
	lenght int
}

func NewMyArray(size int) MyArray {
	return MyArray{
		data: make([]string, size),
	}
}

func (arr *MyArray) Aset(index int, value string) {
	if index >= 0 && index < len(arr.data) {
		arr.data[index] = value
	}
}

func (arr *MyArray) Aget(index int) string {
	if index >= 0 && index < len(arr.data) {
		if arr.data[index]!=""{
		return arr.data[index]
	}
}
	//else{
	return "Error"
//}
}

func (arr *MyArray) ARadd(value string) string {
	for i := 0; i < len(arr.data); i++ {
		if arr.data[i] == "" {
			arr.data[i] = value
			return "Added"
		}
	}
	return "Array is full"
}

func (arr *MyArray) Adel(index int) string {
	if index >= 0 && index < len(arr.data) {
		// deletedValue := arr.data[index]
		// arr.data[index] = ""
		for i:=index;i<len(arr.data)-1;i++{
			arr.data[i]=arr.data[i+1]
		}
		arr.data[len(arr.data)-1]=""
		//return deletedValue
		return "Deleted"
	}

	return "Index not found"
}

func (arr MyArray) PrintArray() {
	for i := 0; i < len(arr.data); i++ {
		if arr.data[i] != "" {
			fmt.Printf("%s ", arr.data[i])
		}
	}
	fmt.Println()
}

// func main(){
// 	myArr:=NewMyArray(10)
// 	myArr.ARadd("0")
// 	myArr.ARadd("1")
// 	myArr.ARadd("2")
// 	myArr.ARadd("3")
// 	myArr.ARadd("4")
// 	myArr.ARadd("5")
// 	myArr.ARadd("6")
// 	myArr.ARadd("7")
// 	myArr.ARadd("8")
// 	myArr.ARadd("9")
// 	myArr.PrintArray()
// 	myArr.Adel(5)
// 	myArr.PrintArray()

// }




/*package main


import "fmt"

type MyArray struct {
	data   []string
	length int
}

func NewMyArray(size int) MyArray {
	return MyArray{
		data:   make([]string, size),
		length: size,
	}
}

func (arr *MyArray) Aset(index int, value string) {
	if index >= 0 && index < arr.length {
		arr.data[index] = value
	}
}

func (arr *MyArray) Aget(index int) string {
	if index >= 0 && index < arr.length {
		//fmt.Println(index)
		return arr.data[index]
	}
	return "Error"
}
	// arr.data = append(arr.data, value)
	// arr.length++
func (arr *MyArray) ARadd(value string) {
	newData:=make([]string,arr.length+1)
	for i:=0;i<arr.length;i++{
		newData[i]=arr.data[i]
	}
	newData[arr.length]=value
	arr.data=newData
	arr.length++
}


func (arr *MyArray) Adel(index int) string {
	if index >= 0 && index < arr.length {
		deletedValue := arr.data[index] // Сохраняем значение элемента
		copy(arr.data[index:], arr.data[index+1:]) // Смещаем элементы влево на одну позицию
		arr.data = arr.data[:len(arr.data)-1]
		arr.length-- // Уменьшаем длину массива
		return deletedValue // Возвращаем удаленное значение
	}
	return "index not found" 
}


func (arr MyArray) PrintArray() {
    for i := 0; i < arr.length; i++ {
        fmt.Printf("%s ", arr.data[i])
    }
    fmt.Println() // Перейти на следующую строку после вывода массива
}
*/

