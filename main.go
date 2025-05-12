package main

import (
    "bufio"
    "fmt"
    "strings"
    "sort"
    "os"
)

// Struct học sinh
type Class struct {
    Class_id   string
    ClassName  string
    GVCN       string
    Students []Student
}

type Student struct {
    Student_id  string
    StudentName string
    Address     string
    Phone       string
}

var classes []Class

var students []Student

func Sort[T any](data []T, less func(a, b T) bool) {
    sort.Slice(data, func(i, j int) bool {
        return less(data[i], data[j])
    })
}

func showMenu() {
    fmt.Println("\n===== MENU =====")
    fmt.Println("1. Danh sách lớp sort theo name ASC")
    fmt.Println("2. Danh sách học sinh sort theo name DESC")
    fmt.Println("3. Danh sách lớp học sinh X đã tham gia(X là student_id nhập từ terminal)")
    //case 4  Danh sách giáo viên với các filter như sau
    fmt.Println("4. Kết thúc")
    fmt.Print("Chọn (1-4): ")
}

func listClasses() {
    if len(classes) == 0 {
        fmt.Println("Chưa có dữ liệu")
        return
    }
    Sort(classes, func(a, b Class) bool {
        return a.ClassName < b.ClassName
    })
    fmt.Println("\n--- Danh sách lớp học ---")
    for i, c := range classes {
        fmt.Printf("%d class_id: %s, Class: %s, GVCN: %s\n",
            i+1, c.Class_id, c.ClassName, c.GVCN)
    }
}

func listStudents() {
    if len(students) == 0 {
        fmt.Println("Chưa có dữ liệu")
        return
    }
    Sort(students, func(a, b Student) bool {
        return a.StudentName < b.StudentName
    })
    fmt.Println("\n--- Danh sách học sinh ---")
    for i, s := range students {
        fmt.Printf("%d. %s, Student: %s, Address: %s, Phone: %s\n",
            i+1, s.Student_id, s.StudentName, s.Address, s.Phone)
    }
}

func getClassOfStudent(studentID string, classes []Class) []Class {
    var result []Class
    for _, class := range classes {
        for _, st := range class.Students {
            if st.Student_id == studentID {
                result = append(result, class)
                break // tránh trùng lặp
            }
        }
    }
    return result
}

func main() {
    classes = []Class {
        {
            Class_id: "1", ClassName: "Class_F", GVCN: "Nguyen Thi A",
            Students: []Student{
                {Student_id: "1"},
            },
        },
        {
            Class_id: "2", ClassName: "Class_G", GVCN: "Nguyen Thi B",
            Students: []Student{
                {Student_id: "1"},
                {Student_id: "2"},
                {Student_id: "3"},
            },
        },
        {
            Class_id: "3", ClassName: "Class_H", GVCN: "Nguyen Thi C",
            Students: []Student{
                {Student_id: "1"},
                {Student_id: "2"},
                {Student_id: "3"},
                {Student_id: "4"},
                {Student_id: "5"},
                {Student_id: "6"},
            },
        },
        {
            Class_id: "4", ClassName: "Class_K", GVCN: "Nguyen Thi D",
            Students: []Student{
                {Student_id: "5"},
                {Student_id: "6"},
            },
        },
        {
            Class_id: "5", ClassName: "Class_O", GVCN: "Nguyen Thi E",
            Students: []Student{
                {Student_id: "1"},
                {Student_id: "7"},
            },
        },
    }
    students = [] Student {
        {"1", "Trần Văn Học Sinh", "Thủ Đức", "0351212312",},
        {"2", "Trần Văn H", "Hồ Chí Minh", "0351212313",},
        {"3", "Trần Văn Học", "Thủ Đức", "0351212314",},
        {"4", "Trần Văn Sinh", "Thủ Đức", "0351212315",},
        {"5", "Trần Văn Sinh Học", "Thủ Đức", "0351212316",},
        {"6", "Trần Văn Sin", "Hồ Chí Minh", "0351212317",},
        {"7", "Trần Văn Ho", "Bến Tre", "0351212318",},
    }
    scanner := bufio.NewScanner(os.Stdin)

    for {
        showMenu()
        scanner.Scan()
        choice := strings.TrimSpace(scanner.Text())

        switch choice {
        case "1":
            listClasses()
        case "2":
            listStudents()
        case "3":
            var inputID string
            fmt.Print("Nhập student_id: ")
            fmt.Scanln(&inputID)

            result := getClassOfStudent(inputID, classes)

            fmt.Printf("Danh sách lớp học sinh có id: %s tham gia:\n", inputID)
            for _, class := range result {
                fmt.Printf("- %s (%s)\n", class.ClassName, class.GVCN)
            }
        case "4":
            fmt.Println("Kết thúc.")
            return
        default:
            fmt.Println("Lựa chọn không hợp lệ.")
        }
    }
}