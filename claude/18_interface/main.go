package main

import "fmt"

// Example 1: Basic Interface
type Speaker interface {
	speak() string
}

type Dog struct {
	name string
}

func (d Dog) speak() string {
	return "Woof!"
}

type Cat struct {
	name string
}

func (c Cat) speak() string {
	return "Meow!"
}

type People struct {
	name string
}

func (p People) speak() string {
	return "Hello!"
}

func makeItSpeak(s Speaker) {
	fmt.Println(s.speak())
}

// __ end __

// Example 2: Real-World Interface - Payment Processing
type Payment interface {
	pay(amount float64) string
}

type CreditCard struct {
	Name   string
	Number string
}

func (c CreditCard) pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card %s", amount, c.Number)
}

type PayPal struct {
	Email string
}

func (p PayPal) pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal %s", amount, p.Email)
}

func ProcessPayment(p Payment, amount float64) {
	fmt.Println(p.pay(amount))
}

// __ end __

// Example 3: Empty Interface (interface{}) and Type Assertions
func PrintAnything(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func DescribeValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float64: %f\n", v)
	default:
		fmt.Println("It's something else")
	}
}

func main() {
	dog := Dog{name: "Aung Nat"}
	cat := Cat{name: "Meow"}
	people := People{name: "Julian"}

	makeItSpeak(dog)
	makeItSpeak(cat)
	makeItSpeak(people)

	creditCard := CreditCard{
		Name:   "Julian",
		Number: "1234567890123456",
	}
	paypal := PayPal{
		Email: "[EMAIL_ADDRESS]",
	}

	ProcessPayment(creditCard, 100)
	ProcessPayment(paypal, 200)

	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(true)
	PrintAnything(3.14)

	fmt.Println()

	DescribeValue(100)
	DescribeValue("Go is awesome!")
	DescribeValue(false)
	DescribeValue(99.99)

	// Exercises 1
	circle := Circle{
		radius: 5,
	}
	rectangle := Rectangle{
		width:  10,
		height: 5,
	}
	triangle := Triangle{
		base:   10,
		height: 5,
	}

	shapes := []Shape{circle, rectangle, triangle}

	for _, shape := range shapes {
		fmt.Printf("Area: %f, Perimeter: %f\n", shape.Area(), shape.Perimeter())
	}

	// Exercises 2
	mysql := MySQL{
		name: "MySQL",
	}
	postgreSQL := PostgreSQL{
		name: "PostgreSQL",
	}

	databases := []Database{mysql, postgreSQL}

	for _, db := range databases {
		fmt.Println(db.Connect())
		fmt.Println(db.Query("SELECT * FROM users"))
		fmt.Println(db.Close())
	}

	// Exercises 3
	product := Product{
		name:  "Product 1",
		id:    1,
		price: 100,
	}

	jsonSerializer := JSONSerializer{
		name: "JSONSerializer",
	}
	xmlSerializer := XMLSerializer{
		name: "XMLSerializer",
	}

	serializers := []Serializer{jsonSerializer, xmlSerializer}

	for _, serializer := range serializers {
		fmt.Println(serializer.Serialize(product))
	}
}

// Exercises 1: Create a Shape interface with methods Area() and Perimeter(). Implement it for Circle, Rectangle, and Triangle structs.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.142 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.142 * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return t.base + t.height + t.base
}

func ShowArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

// Exercises 2: Create a Database interface with methods Connect(), Query(query string), and Close(). Implement it for MySQL and PostgreSQL types.

type Database interface {
	Connect() string
	Query(query string) string
	Close() string
}

type MySQL struct {
	name string
}

func (m MySQL) Connect() string {
	return "Connected to MySQL"
}

func (m MySQL) Query(query string) string {
	return "Querying MySQL: " + query
}

func (m MySQL) Close() string {
	return "Closed MySQL"
}

type PostgreSQL struct {
	name string
}

func (p PostgreSQL) Connect() string {
	return "Connected to PostgreSQL"
}

func (p PostgreSQL) Query(query string) string {
	return "Querying PostgreSQL: " + query
}

func (p PostgreSQL) Close() string {
	return "Closed PostgreSQL"
}

// Exercises 3: Create a Serializer interface with a Serialize() string method. Implement it for JSON and XML types that can serialize a Product struct.

type Product struct {
	name  string
	id    int
	price float64
}

type Serializer interface {
	Serialize(p Product) string
}

type JSONSerializer struct {
	name string
}

func (j JSONSerializer) Serialize(p Product) string {
	return fmt.Sprintf("JSON: %s", p.name)
}

type XMLSerializer struct {
	name string
}

func (x XMLSerializer) Serialize(p Product) string {
	return fmt.Sprintf("XML: %s", p.name)
}
