package main

import (
	"fmt"
	"go-re-bootcamp/logger"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	// hello()
	// title := "The way to go"
	// author := "Some one"
	// var copies_sold = 10
	// printBook(title, author, copies_sold)
	// wayToGoBook := Book{Title: title, Author: author}
	// wayToGoBook.Copies.Store(int32(copies_sold))
	// wayToGoBook.showBook()
	// wayToGoBook.updateCopies(12)
	// wayToGoBook.showBook()
	// consoleLogger := logger.ConsoleLogger{}
	// consoleLogger.Log(logger.INFO, "info msg")
	// utils.Print()
	// utils.GIMax(23, 34)
	// var a = uint(23)
	// var b = uint(23)
	// utils.GIMax(a, b)
	// var events = []int{1, 2, 3, 4}
	// for _, e := range events {
	// 	println(e)
	// }
	// x := utils.Box[int]{V: 12}
	// xTimes2 := x.Map(func(i int) int { return i * 2 })
	// xTimes2AsStr := xTimes2.Map(func(i int) int { return i * 3 })
	// fmt.Printf("Generic Compute  %d %d \n", xTimes2.V, xTimes2AsStr.V)
	// conc.NewConc()
	// ch := make(chan int)
	// go conc.SendData(ch)
	// conc.ReceiveData(ch)

	logFile, err := os.OpenFile("bootcamp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	var fileLogger = logger.NewFileLogger(logFile)
	fileLogger.Info("bing")

	// conc.CoOrdinate(2)
}

func printBook(title string, author string, copies int) {
	fmt.Printf("Book Titled %s Author %s sold %d copies. \n", title, author, copies)
}

func printBookStruct(book *Book) {
	fmt.Printf("Book titled %s Author %s Sold %d Copies", book.Author, book.Title, book.Copies.Load())
}

func hello() {
	fmt.Println("Hello from the hello function!")
}

type Book struct {
	Title  string
	Author string
	Copies atomic.Int32
}

func (book *Book) showBook() {
	fmt.Printf("Book titled %s Author %s Sold %d Copies \n", book.Author, book.Title, book.Copies.Load())
}

func (book *Book) updateCopies(copies int) *Book {
	book.Copies.Add(int32(copies))
	return book
}

type Row struct {
	Id        int64
	Name      string
	Pin       string
	Age       byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RowStore struct {
	rowPtr atomic.Pointer[Row]
}

func NewRowStore(row Row) *RowStore {
	rowStore := &RowStore{}
	rowStore.rowPtr.Store(&row)
	return rowStore
}

func (rs *RowStore) showRow() *Row {
	return rs.rowPtr.Load()
}

func (rs *RowStore) updateAge(row Row) {
	old := rs.rowPtr.Load()
	new := row
	if rs.rowPtr.CompareAndSwap(old, &new) {
		return
	}
}
