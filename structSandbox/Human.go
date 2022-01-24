package main



type Human struct{
	firstName string
	lastName string
	task string
	age uint32
}

func (human *Human) setFirstname(firstName string){
	human.firstName = firstName
}

func(human *Human) getFirstname()string{
	return human.firstName
}


