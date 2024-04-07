package main

import "fmt"

type Student struct {
	name   string
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func main() {
	s := Student{name: "Joe", grades: []int{90, 75, 80}}
	s.displayName()
	fmt.Printf("%.2f%%", s.calculatePercentage())
}

/*

package main
import "fmt"


type Rectangle struct {
	lenght int
	breadth int
}

func (r Rectangle) area() int {
	return r.lenght * r.breadth
}

func (r *Rectangle) incLength(n int) {
	for i := 0; i < n; i++ {
		r.lenght += i
	}
}

func main() {
	r := Rectangle{breadth: 10, lenght: 5}
	fmt.Println(r.area())
	fmt.Println(r)
	r.incLength(7)
	fmt.Println(r.area())
	fmt.Println(r)
}


-----------------------------


package main
import "fmt"


type Employee struct {
	eid int
	id int
}

func main(){
	employees := make([]Employee, 5)
	for i := range employees {
		employees[i] = Employee{i, i + 10}
		fmt.Println(employees[i])
	}
}


*/
