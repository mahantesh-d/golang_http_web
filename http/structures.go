package main

import "fmt"

type books struct {
	title string
	author string
	subject string
	id int
}

func main()  {
	var book1 books
	var book2 books

	book1.author = "Mahantesh"
	book1.title = "Big Data"
	book1.subject = "Data Science"
	book1.id = 10023

	book2.author = "Akshay"
	book2.title = "Data Scientist"
	book2.subject = "Data Science"
	book2.id = 100001

	fmt.Printf("Book1 Author is %s with title %s for the subject %s and id is %d", book1.author, book1.title, book1.subject, book1.id)
	fmt.Printf("\nBook2 Author is %s with title %s for the subject %s and id is %d", book2.author, book2.title, book2.subject, book2.id)

}
