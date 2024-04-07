package main

import "fmt"

//st := new(Studiante)

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
		name:   s,
		rating: r,
	}
	return
}

func main() {
	var s Student
	st := new(Student)
	fmt.Printf("%+v \n", s)
	fmt.Printf("%+v \n ", st)

	estudiante1 := Student{
		name:   "Joe",
		rollNo: 12,
		marks:  []int{},
		grades: map[string]int{},
	}

	type Estudiante struct { //it is not recommendable
		name   string
		rollNo int
	}
	fmt.Printf("%+v\n\n", estudiante1)

	estudianteMalapractica := Estudiante{"Joe", 12} //it is not recommendable
	fmt.Printf("%+v\n", estudianteMalapractica)

	fmt.Println(estudianteMalapractica.name)

	fmt.Println("Movie Struct with function")
	fmt.Printf("%+v", getMovie("xyz", 3.5))

}
