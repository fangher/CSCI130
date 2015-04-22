package main

import ("fmt")

type Animal struct{
	Name string
	Age uint
}

type str string

func (animal Animal) Printanimal() {
	fmt.Printf("%s is %d years old\n", animal.Name, animal.Age);
}

func (s str) Printstr(){
	fmt.Println(s);
}

func main() {
	words := []string{"Cat", "Dog", "Rabbit", "Bird"}
	for _,word := range words {
		fmt.Println(word)
	}
	fmt.Println();

	words = append(words, "House");
	for _,word := range words {
		fmt.Println(word)
	}
	fmt.Println();
	
	words = append(words[:3], words[4:]...);
	for _,word := range words {
		fmt.Println(word)
	}
	fmt.Println();
		
	animal := Animal{"John", 10};
	animal.Printanimal();
	
	var message str = "Hello";
	message.Printstr();
}
