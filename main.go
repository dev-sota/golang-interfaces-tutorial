package main

import "fmt"

type Student struct {
	Name   string
	Number int
	Grade  int
}

type Teacher struct {
	Name string
}

func main() {
	s := Student{
		Name:   "Yamada",
		Number: 999,
		Grade:  5,
	}
	t := Teacher{
		Name: "Tsubomi",
	}

	cxtStu := sendEmailOfStudent(s)
	fmt.Println(cxtStu)
	cxtTea := sendEmailOfTeacher(t)
	fmt.Println(cxtTea)
}

func (s Student) getEmail() string {
	return s.Name + "@student.co.jp"
}
func (t Teacher) getEmail() string {
	return t.Name + "@teacher.co.jp"
}

func sendEmailOfStudent(s Student) (context string) {
	from := s.getEmail()
	context = `
  送信元 : ` + from + `
  これはテスト用のメールです。
  よろしくお願いします。
  `
	return context
}

func sendEmailOfTeacher(t Teacher) (context string) {
	from := t.getEmail()
	context = `
  送信元 : ` + from + `
  これはテスト用のメールです。
  よろしくお願いします。
  `
	return context
}
