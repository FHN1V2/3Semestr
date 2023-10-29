package main


import "fmt"

// Структура для узла двусвязного списка
type Node struct {
    data     string
    next     *Node
    previous *Node
}

// Структура для самого двусвязного списка
type DoublyLinkedList struct {
    head *Node
    tail *Node
}

// Создание нового пустого двусвязного списка
func NewDoublyLinkedList() *DoublyLinkedList {
    return &DoublyLinkedList{}
}

// Добавление элемента в конец списка
func (list *DoublyLinkedList) DladdE(data string) {
    newNode := &Node{data: data}

    if list.head == nil {
        list.head = newNode
        list.tail = newNode
    } else {
        newNode.previous = list.tail
        list.tail.next = newNode
        list.tail = newNode
    }
}
// Добавление элемента в конец списка

func (list *DoublyLinkedList) DlAdd(data string) {
    newNode := &Node{data: data}

    if list.head == nil {
        // Если список пуст, новый узел становится и начальным, и конечным элементом
        list.head = newNode
        list.tail = newNode
    } else {
        newNode.next = list.head // Устанавливаем указатель следующего элемента нового узла на текущий начальный элемент
        list.head.previous = newNode // Устанавливаем указатель предыдущего элемента текущего начального элемента на новый узел
        list.head = newNode // Обновляем указатель начального элемента на новый узел
    }
}
// Удаление элемента из списка
func (list *DoublyLinkedList) Dldel(data interface{}) {
    current := list.head
    for current != nil {
        if current.data == data {
            if current.previous != nil {
                current.previous.next = current.next
            } else {
                list.head = current.next
            }
            if current.next != nil {
                current.next.previous = current.previous
            } else {
                list.tail = current.previous
            }
            return
        }
        current = current.next
    }
}

// Вывод содержимого списка
func (list *DoublyLinkedList) DLdisplay() {
    current := list.head
    for current != nil {
        fmt.Printf("%v <-> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}