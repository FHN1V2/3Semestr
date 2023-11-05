package main

import "fmt"

type MyArray struct {
    data   []string
    length int
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
        if arr.data[index] != "" {
            return arr.data[index]
        }
    }
    return "Error"
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
        for i := index; i < len(arr.data)-1; i++ {
            arr.data[i] = arr.data[i+1]
        }
        arr.data[len(arr.data)-1] = ""
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

func Subarrays(arr MyArray) [][]string {
    result := [][]string{}
    n := len(arr.data)

    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            subarray := []string{}
            for k := i; k <= j; k++ {
                if arr.data[k] != "" {
                    subarray = append(subarray, arr.data[k])
                }
            }
            result = append(result, subarray)
        }
    }

    return result
}

func main() {
    arr := NewMyArray(5)
    arr.ARadd("1")
    arr.ARadd("2")
    arr.ARadd("3")
	arr.ARadd("4")
	arr.ARadd("5")

    subarrays := Subarrays(arr)
    for _, subarray := range subarrays {
        fmt.Println(subarray)
    }
}
