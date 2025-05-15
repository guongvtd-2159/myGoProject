package main

import (
    "bufio"
    "fmt"
    "strings"
    "sort"
    "os"
    "strconv"
    "encoding/json"
)

type School struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Class struct {
	Class_id   string `json:"id"`
	ClassName  string `json:"name"`
	School_id  string `json:"school_id"`
	Teacher_id string `json:"teacher_id"`
}

type Teacher struct {
	Teacher_id string `json:"id"`
	TeacherName       string `json:"name"`
	Address    string `json:"address"`
}

type Student struct {
	Student_id  string `json:"id"`
	StudentName string `json:"name"`
	Address     string `json:"address"`
}

type StudentClass struct {
	Student_id string `json:"student_id"`
	Class_id   string `json:"class_id"`
}

type Data struct {
	Schools        []School       `json:"schools"`
	Classes        []Class        `json:"classes"`
	Teachers       []Teacher      `json:"teachers"`
	Students       []Student      `json:"students"`
	StudentClasses []StudentClass `json:"student_classes"`
}

func loadData(filename string) (Data, error) {
	var data Data
	content, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(content, &data)
	return data, err
}

func Sort[T any](data []T, less func(a, b T) bool) {
    sort.Slice(data, func(i, j int) bool {
        return less(data[i], data[j])
    })
}

func getTeacher(id string, teachers []Teacher) Teacher {
	for _, t := range teachers {
		if t.Teacher_id == id {
			return t
		}
	}
	return Teacher{}
}

func getStudent(id string, students []Student) Student {
	for _, s := range students {
		if s.Student_id == id {
			return s
		}
	}
	return Student{}
}

func getTeacherName(id string, teachers []Teacher) string {
	for _, t := range teachers {
		if t.Teacher_id == id {
			return t.TeacherName
		}
	}
	return "(unknown)"
}

func isGVCN(id string, classes []Class) bool {
	for _, class := range classes {
		if class.Teacher_id == id {
			return true
		}
	}
	return false
}

func showMenu() {
    fmt.Println("\n===== MENU =====")
    fmt.Println("1. Danh sách lớp sort theo name ASC")
    fmt.Println("2. Danh sách học sinh sort theo name DESC")
    fmt.Println("3. Danh sách lớp học sinh X đã tham gia(X là student_id nhập từ terminal)")
    fmt.Println("4. Danh sách giáo viên với các filter như sau:")
    fmt.Println("5. Kết thúc")
    fmt.Print("Chọn (1-5): ")
}

func listClasses(classesData []Class, teacherData []Teacher) {
    if len(classesData) == 0 {
        fmt.Println("Chưa có dữ liệu")
        return
    }
    Sort(classesData, func(a, b Class) bool {
        return a.ClassName < b.ClassName
    })
    fmt.Println("\n--- Danh sách lớp học ---")
    for i, c := range classesData {
        GVCN := getTeacher(c.Teacher_id, teacherData)
        fmt.Printf("%d class_id: %s, Class: %s, GVCN: %s\n",
            i+1, c.Class_id, c.ClassName, GVCN.TeacherName)
    }
}

func listStudents(studentData []Student) {
    if len(studentData) == 0 {
        fmt.Println("Chưa có dữ liệu")
        return
    }
    Sort(studentData, func(a, b Student) bool {
        return a.StudentName > b.StudentName
    })
    fmt.Println("\n--- Danh sách học sinh ---")
    for i, s := range studentData {
        fmt.Printf("%d. %s, Student: %s, Address: %s\n",
            i+1, s.Student_id, s.StudentName, s.Address)
    }
}

func getClassOfStudent(studentID string, studentClasses []StudentClass) map[string]bool {
    result := make(map[string]bool)
    for _, sc := range studentClasses {
        if sc.Student_id == studentID {
            result[sc.Class_id] = true
        }
    }
    return result
}

func getListTeacherWithFilter(scanner *bufio.Scanner, data Data) {
    fmt.Println("Chọn filter:")
    fmt.Println("1. Tất cả giáo viên")
    fmt.Println("2. Giáo viên là GVCN")
    fmt.Println("3. GVCN lớp có trên X học sinh")
    fmt.Println("4. GVCN trên X lớp")
    fmt.Print("Chọn filter: ")
    scanner.Scan()
    option := strings.TrimSpace(scanner.Text())

    switch option {
    case "1":
        fmt.Println("Danh sách tất cả giáo viên:")
        for _, teacher := range data.Teachers {
            fmt.Println("-", teacher.TeacherName)
        }
    case "2":
        fmt.Println("Giáo viên là GVCN:")
        for _, teacher := range data.Teachers {
            if isGVCN (teacher.Teacher_id, data.Classes) {
                fmt.Println("-", teacher.TeacherName)
            }
        }
    case "3":
        fmt.Print("Nhập số  lượng học sinh X: ")
        scanner.Scan()
        x, _ := strconv.Atoi(scanner.Text())
        count := map[string]int{}
		for _, sc := range data.StudentClasses {
			count[sc.Class_id]++
		}
		classMap := map[string]bool{}
		for _, c := range data.Classes {
			if count[c.Class_id] > x {
				classMap[c.Teacher_id] = true
			}
		}
		for _, t := range data.Teachers {
			if classMap[t.Teacher_id] {
				fmt.Printf("- %s (%s)\n", t.TeacherName, t.Address)
			}
		}
    case "4":
        fmt.Print("Nhập số lượng lớp: ")
        scanner.Scan()
        x, _ := strconv.Atoi(scanner.Text())
        teacherCount := make(map[string]int)
        for _, class := range data.Classes {
            teacherCount[class.Teacher_id]++
        }
        for tid, c := range teacherCount {
            if c > x {
                t := getTeacher(tid, data.Teachers)
                fmt.Printf("- %s (%s): %d lớp\n", t.TeacherName, t.Address, c)
            }
        }

    default:
        fmt.Println("Lựa chọn không hợp lệ.")
    }
}

func main() {
    var data Data
	data, err := loadData("data.json")
	if err != nil {
		fmt.Println("Lỗi đọc file:", err)
		return
	}
    scanner := bufio.NewScanner(os.Stdin)

    for {
        showMenu()
        scanner.Scan()
        choice := strings.TrimSpace(scanner.Text())

        switch choice {
        case "1":
            listClasses(data.Classes, data.Teachers)
        case "2":
            listStudents(data.Students)
        case "3":
            var inputID string
            fmt.Print("Nhập student_id: ")
            fmt.Scanln(&inputID)

            result := getClassOfStudent(inputID, data.StudentClasses)

            fmt.Printf("Danh sách lớp học sinh có id: %s tham gia:\n", inputID)
            for _, class := range data.Classes {
                if result[class.Class_id] {
                    fmt.Printf("- %s\n", class.ClassName)
                }
            }
        case "4":
           getListTeacherWithFilter(scanner, data)
        case "5":
            fmt.Println("Kết thúc.")
            return
        default:
            fmt.Println("Lựa chọn không hợp lệ.")
        }
    }
}