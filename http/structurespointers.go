package main

import "fmt"

type books struct {
	author string
	title string
	subject string
	id int
}

func print1(book *books){
	fmt.Printf("Book1 Author is %s with title %s for the subject %s and id is %d\n", book.author, book.title, book.subject, book.id)
}

func main()  {
	var book1 books
	var book2 books

	book1.title = "Machine Learning"
	book1.subject = "Scala"
	book1.author = "Mahantesh"
	book1.id = 4567

	book2.title = "Data Analysis with Machine Learning"
	book2.subject = "Go Programing"
	book2.author = "Akshay"
	book2.id = 7894

	print1(&book1)
	print1(&book2)
}
