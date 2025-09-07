package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return &person{}
		},
		//if New key of sync.Pool is not defined then in case when there is no instance
		//of person is available in pool, it will not create an empty instance and throw error.
		//When New key is defined then it will empty instance of person whenever pool.Get() is called.
		//Use these 2 things wisely, as per the usecase.
	}
	pool.Put(&person{name: "Raman", age: 25})
	//Get an object from the pool
	person1 := pool.Get().(*person)
	person1.name = "John"
	person1.age = 25
	fmt.Println("Got Person1:", person1)

	fmt.Printf("Person1 - Name: %s, Age : %d\n", person1.name, person1.age)
 
	pool.Put(person1)
	fmt.Println("Returned Person1 to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got Person2 :", person2)

	person3 := pool.Get().(*person)
	fmt.Println("Got Person3 :", person3)
	person3.name = "Jane"

	//Returning persons to pool again
	pool.Put(person2)
	pool.Put(person3)
	fmt.Println("Return Persons to Pool")

	person4 := pool.Get().(*person)
	fmt.Println("Got Person 4:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got Person 5 :", person5)
}
