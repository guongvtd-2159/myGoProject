package main

import "fmt"

// func main() { //for range
//     letters := []string{"a", "b", "c"}

//     //With index and value
//     fmt.Println("Both Index and Value")
//     for i, letter := range letters {
//         fmt.Printf("Index: %d Value:%s\n", i, letter)
//     }

//     //Only value
//     fmt.Println("\nOnly value")
//     for _, letter := range letters {
//         fmt.Printf("Value: %s\n", letter)
//     }

//     //Only index
//     fmt.Println("\nOnly Index")
//     for i := range letters {
//         fmt.Printf("Index: %d\n", i)
//     }

//     //Without index and value. Just print array values
//     fmt.Println("\nWithout Index and Value")
//     i := 0
//     for range letters {
//         fmt.Printf("Index: %d Value: %s\n", i, letters[i])
//         i++
//     }
// }

// func main() { // slice
//     numbers := [5]int{1, 2, 3, 4, 5}

//     //Both start and end
//     num1 := numbers[2:4]
//     fmt.Println("Both start and end")
//     fmt.Printf("num1=%v\n", num1)
//     fmt.Printf("length=%d\n", len(num1))
//     fmt.Printf("capacity=%d\n", cap(num1))

//     //Only start
//     num2 := numbers[2:]
//     fmt.Println("\nOnly start")
//     fmt.Printf("num1=%v\n", num2)
//     fmt.Printf("length=%d\n", len(num2))
//     fmt.Printf("capacity=%d\n", cap(num2))

//     //Only end
//     num3 := numbers[:3]
//     fmt.Println("\nOnly end")
//     fmt.Printf("num1=%v\n", num3)
//     fmt.Printf("length=%d\n", len(num3))
//     fmt.Printf("capacity=%d\n", cap(num3))

//     //None
//     num4 := numbers[:]
//     fmt.Println("\nOnly end")
//     fmt.Printf("num1=%v\n", num4)
//     fmt.Printf("length=%d\n", len(num4))
//     fmt.Printf("capacity=%d\n", cap(num4))
// }


import "time"

// func main() {
//     go start()
//     fmt.Println("Started")
//     // time.Sleep(1 * time.Second)
//     fmt.Println("Finished")
// }

// func start() {
//     fmt.Println("In Goroutine")
// }

func execute(id int) {
    fmt.Printf("id: %d\n", id)
}

func main() {
    fmt.Println("Started")
    for i := 0; i < 10; i++ {
        go execute(i)
    }
    time.Sleep(time.Second * 2)
    fmt.Println("Finished")
}