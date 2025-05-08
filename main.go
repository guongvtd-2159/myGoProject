package main

import "fmt"

// func main() {
//     var attendance, test, exam float64

//     fmt.Print("input attendance (0-10): ")
//     fmt.Scanln(&attendance)

//     fmt.Print("input test (0-10): ")
//     fmt.Scanln(&test)

//     fmt.Print("input exam (0-10): ")
//     fmt.Scanln(&exam)

//     average := 0.1*attendance + 0.3*test + 0.6*exam

//     fmt.Printf("average: %.2f\n", average)
// }

// interface

type AverageCalculator interface {
    Average() float64
}

type Student struct {
    Attendance float64
    Test float64
    Exam float64
}

// Interface AverageCalculator student
func (st Student) Average() float64 {
    return 0.1*st.Attendance + 0.3*st.Test + 0.6*st.Exam
}

func main() {
    var attendance, test, exam float64

    fmt.Print("input attendance (0-10): ")
    fmt.Scanln(&attendance)

    fmt.Print("input test (0-10): ")
    fmt.Scanln(&test)

    fmt.Print("input exam (0-10): ")
    fmt.Scanln(&exam)

    var calculator AverageCalculator = Student{
        Attendance: attendance,
        Test: test,
        Exam: exam,
    }

    fmt.Printf("average: %.2f\n", calculator.Average())
}