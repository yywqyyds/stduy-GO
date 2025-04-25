package main

import "fmt"

//图书
type Book struct {
	ID          int
	Title       string
	Author      string
	IsAvailable bool
}

//杂志
type Magazine struct {
	ID          int
	Title       string
	Issue       int
	IsAvailable bool
}

//图书馆
type Library struct {
	Books     []*Book
	Magazines []*Magazine
	Name      string
}

// 操作接口
type Manageable interface {
	Borrow() bool
	Retrun() bool
	GetInfo() string
}

func (b *Book) Borrow() bool {
	if !b.IsAvailable {
		return false
	}
	b.IsAvailable = false
	return true
}

func (b *Book) Retrun() bool {
	if b.IsAvailable {
		return false
	}
	b.IsAvailable = true
	return true
}

func (b *Book) GetInfo() string {
	status := "借出中"
	if b.IsAvailable {
		status = "可借"
	}
	return fmt.Sprintf("【图书】ID:%d 标题:%s 作者:%s 状态:%s", b.ID, b.Title, b.Author, status)
}

func (l *Library) AddBook(book *Book) {
	l.Books = append(l.Books, book)
	fmt.Printf("成功添加图书：《%s》 到图书馆 %s\n", book.Title, l.Name)
}

func (l *Library) FindBook(book *Book) {
	for _, i := range l.Books {
		if i.ID == book.ID {
			i.GetInfo()
		}
	}
}

func (l *Library) FindAllBook() {
	for _, i := range l.Books {
		if i.IsAvailable {
			fmt.Println(i.GetInfo)
		}
	}
}

func (l *Library) AddMagazine(magazine *Magazine) {
	l.Magazines = append(l.Magazines, magazine)
	fmt.Printf("成功添加杂志：《%s》 到图书馆 %s\n", magazine.Title, l.Name)
}
