/*lab 6

We are working on a program that allows users to store and retrieve information about
their favourite books. We have implemented a Book struct to store the information for the book.


You need to implement a function that will change the value of the Pages field for a given Book.

Also, make the required changes in the main function.

A Go file is located at /root/code/book directory.*/

package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	// your code goes here
	book.Pages = pages
}

func main() {

	/*
	   Create 3 Book Structs with the following data:

	   Book 1:
	   Title: "The Great Gatsby"
	   Author: "F. Scott Fitzgerald"
	   Pages: 180

	   Book 2
	   Title: "To Kill a Mockingbird"
	   Author: "Harper Lee"
	   Pages: 281

	   Book 3
	   Title: "Pride and Prejudice"
	   Author: "Jane Austen"
	   Pages: 279
	*/

	// your code for creating struct objects goes here
	Libro1 := Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Pages:  180,
	}

	Libro2 := Book{
		Title:  "To Kill a Mockingbird",
		Author: "Harper Lee",
		Pages:  281,
	}

	Libro3 := Book{
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Pages:  279,
	}
	/*
	   Update the information for Books as following:

	   Book 1: Updates Page Count to 210
	   Book 2: Updates Page Count to 250
	   Book 3: Updates Page Count to 295

	*/

	// your code for function calls to updatePages goes here
	updatePages(&Libro1, 210)
	updatePages(&Libro2, 250)
	updatePages(&Libro3, 295)
	/*
	   Print all the struct objects
	   fmt.Println(book)
	*/

	// your code for printing objects goes here
	fmt.Println(Libro1)
	fmt.Println(Libro2)
	fmt.Println(Libro3)
}

/*
Expected Output:


&{The Great Gatsby F. Scott Fitzgerald 210}
&{To Kill a Mockingbird Harper Lee 250}
&{Pride and Prejudice Jane Austen 295}*/
