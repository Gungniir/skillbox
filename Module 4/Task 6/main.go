package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Данная программа поможет вам распределить людей по группам")

	var studentsCount int
	fmt.Println("Введите кол-во студентов")
	_, err := fmt.Scan(&studentsCount)
	checkError(err)

	var groupsCount int
	fmt.Println("Введите кол-во групп")
	_, err = fmt.Scan(&groupsCount)
	checkError(err)

	// Деление с округлением вверх без использования float64
	// Можно было написать studentsPerGroup := int(math.Ceil(float64(studentsCount) / groupsCount))
	studentsPerGroup := studentsCount / groupsCount
	if studentsCount%groupsCount != 0 {
		studentsPerGroup++
	}

	var studentID int
	fmt.Println("Введите номер студента")
	_, err = fmt.Scan(&studentID)
	checkError(err)

	// Деление с округлением вверх без использования float64
	studentGroup := studentID / studentsPerGroup
	if studentID%studentsPerGroup != 0 {
		studentGroup++
	}

	fmt.Printf("Студента под номером %d следует распределить в группу номер %d", studentID, studentGroup)
}

func checkError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Произошла ошибка... Попробуйте перезапустить программу")
	os.Exit(1)
}
