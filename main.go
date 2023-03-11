package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Node.js", 32},
		{"Java", 30},
		{"Python", 20},
	})
	if err != nil {
		panic(err)
	}
}
