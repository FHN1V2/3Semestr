package main
import( //импортирование 1 пакета: errors для создания и возврата ошибок
    "fmt"
    "errors"
)


//Массив
type Array struct{
    data []string //хранение данных массива
    length int //хранение текущей длины массива
}


//NewArray создаёт и возвращает новый пустой объект массива (создание нового пустого массива)
func NewArray(size int) Array {
    return Array{
        data: make([]string, size),
    }
}


//Append добавляет элемент в конец массива
func (arr *Array) ArAdd(value string) string{
    for i:=0; i<len(arr.data); i++{
        if arr.data[i] == ""{
            arr.data[i] = value
            return ""
        }
    }
    return "Error: array is full!"
}


//Get возвращает элемент по индексу
func (arr *Array) AGet(index int) (string, error){
    if index<0 || index>= len(arr.data){
        return "0", errors.New("Index out of range")
}
return arr.data[index], errors.New("")
}


//Set устанавливает значение элемента по индексу
func (arr *Array) ASet(index int, value string) error{
    if index < 0|| index >= len(arr.data){
        return errors.New("Index out of range")
}
arr.data[index] = value
return nil
}


//Length возвращает текущую длину массива
func (arr *Array) ALength() int{
    return arr.length
}


//Нахождение элемента, удаление и смещение
func (arr *Array) ADel(index int) string{
    if index>=0 && index<len(arr.data){
        for i:=index; i<len(arr.data)-1; i++{
            arr.data[i]=arr.data[i+1]   //смещение
        }
        arr.data[len(arr.data)-1]=""
        return ""
    }
    return "Error: Index out of range"
}


//Функция вывода
func (arr Array) ArrPrint(){
    for i:=0; i<len(arr.data); i++{
        if arr.data[i]!=""{
            fmt.Printf("%s ",arr.data[i])
        }
    }
}
