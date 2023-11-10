package main

import (
	"fmt"
)

type Set struct {
    data map[string]bool
}

func NewSet() *Set {
    return &Set{data: make(map[string]bool)}
}

func (s *Set) SetAdd(item string) {
    s.data[item] = true
}

func (s *Set) SetRemove(item string) {
    // if a:=s.data[item];!a{
    //     fmt.Println("Not found")
    // }
    
    newData := make(map[string]bool)
    for key := range s.data {
        if key != item {
            newData[key] = false
        }
    }
    s.data = newData
}

func (s *Set) SetContains(item string) bool {
    return s.data[item]
}

func (s *Set) SetSize() int {
    return len(s.data)
}

func (s *Set) SetPrint() {
    fmt.Print("Set elements: ")
    for item := range s.data {
        fmt.Printf("%s ", item)
    }
    fmt.Println()
}

// func main() {
//     mySet := NewSet()
//     mySet.SetAdd("1")
//     mySet.SetAdd("2")
//     mySet.SetAdd("3")
//     mySet.SetAdd("4")
//     mySet.SetAdd("5")
//     mySet.SetAdd("1")
//     mySet.SetAdd("1")
//     mySet.SetAdd("3")
//     mySet.Remove("1")

//     mySet.SetPrint()
//     fmt.Println(mySet.SetSize())
    
// }
