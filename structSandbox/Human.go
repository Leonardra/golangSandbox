package main



type Human struct{
	firstName string
	lastName string
	task string
	age uint32
}

// To demonstrate a has-a relationship

type Boy struct{
	Human
	accountNumber string

}

func (human *Human) setFirstname(firstName string){
	human.firstName = firstName
}

func(human *Human) getFirstname()string{
	return human.firstName
}


