package main

import "fmt"

type books struct {
	author string
	title string
	id int
	subject string
}

func print (book books){
	fmt.Printf("Book1 Author is %s with title %s for the subject %s and id is %d\n", book.author, book.title, book.subject, book.id)
}

func main()  {
	var book1 books
	var book2 books

	book1.author = "Mahantesh"
	book1.subject = "Data Analysis"
	book1.title = "Big Data"
	book1.id = 1234

	book2.author = "Akshay"
	book2.subject = "Real Time Data Analysis"
	book2.title = "Big Data V 2.0"
	book2.id = 9876

	print(book1)
	print(book2)
}
