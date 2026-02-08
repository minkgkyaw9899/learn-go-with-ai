package main

import "fmt"

// Example 1: Defining and Using Structs
type Person struct {
	firstName string
	lastName  string
	age       int
	email     string
}

// Example 2: Nested Structs

type Address struct {
	street  string
	city    string
	zipCode string
}

type Employee struct {
	name    string
	id      int
	salary  float64
	address Address // Nested struct
}

// Example 3: Struct Methods

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

type Book struct {
	title  string
	author string
	pages  int
	price  float64
}

func (b Book) display() {
	fmt.Println("Title:", b.title)
	fmt.Println("Author:", b.author)
	fmt.Println("Pages:", b.pages)
	fmt.Println("Price:", b.price)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) circumference() float64 {
	return 2 * 3.14 * c.radius
}

type Student struct {
	name   string
	id     int
	grades []float64
}

func (s Student) average() float64 {
	total := 0.0
	for _, grade := range s.grades {
		total += grade
	}
	return total / float64(len(s.grades))
}

func main() {
	var person1 Person
	person1.firstName = "John"
	person1.lastName = "Doe"
	person1.age = 30
	person1.email = "[EMAIL_ADDRESS]"

	fmt.Println("First Name:", person1.firstName)
	fmt.Println("Last Name:", person1.lastName)
	fmt.Println("Age:", person1.age)
	fmt.Println("Email:", person1.email)

	person2 := Person{
		firstName: "Jane",
		lastName:  "Smith",
		age:       25,
		email:     "jane@example.com",
	}

	person3 := Person{"Bob", "Johnson", 35, "bob@example.com"}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)

	emp := Employee{
		name:   "Alice",
		id:     101,
		salary: 75000,
		address: Address{
			street:  "123 Main St",
			city:    "New York",
			zipCode: "10001",
		},
	}

	fmt.Printf("%s lives in %s\n", emp.name, emp.address.city)

	rect := Rectangle{
		length: 10,
		width:  5,
	}

	fmt.Println("Area:", rect.area())
	fmt.Println("Perimeter:", rect.perimeter())

	rect.scale(2)
	fmt.Println("Scaled Area:", rect.area())
	fmt.Println("Scaled Perimeter:", rect.perimeter())

	// Exercises 1:Create a Book struct with title, author, pages, and price. Create 3 books and display them.
	book1 := Book{
		title:  "Book 1",
		author: "Author 1",
		pages:  100,
		price:  10,
	}
	book2 := Book{
		title:  "Book 2",
		author: "Author 2",
		pages:  200,
		price:  20,
	}
	book3 := Book{
		title:  "Book 3",
		author: "Author 3",
		pages:  300,
		price:  30,
	}

	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book3)

	book1.display()
	book2.display()
	book3.display()

	// Exercises 2: Create a Circle struct with radius. Add methods for area and circumference.

	circle := Circle{
		radius: 5,
	}

	fmt.Println("Area:", circle.area())
	fmt.Println("Circumference:", circle.circumference())

	// Exercises 3: Create a Student struct with name, ID, and grades (slice). Add a method to calculate average grade.

	student := Student{
		name:   "John",
		id:     101,
		grades: []float64{100, 90, 80},
	}

	fmt.Println("Average grade:", student.average())

}
