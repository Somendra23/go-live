package st

type Person struct {
	firstname string
	lastname  string
}

func DemoStruct() Person {
	p1 := Person{
		firstname: "John",
		lastname:  "Kelly",
	}
	return p1

}
