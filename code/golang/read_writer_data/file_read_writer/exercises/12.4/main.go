package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

const filename = "write_structure_data"

type Page struct {
	Title string
	Body  []byte
}

// Save the structure data to a text file
func (p *Page) save() (err error) {
	// Check that the title is empty
	if p.Title == "" {
		return errors.New("filename can't be empty")
	}
	return ioutil.WriteFile(fullTitle(p.Title), p.Body, 0666)
}

// Gets filename
func fullTitle(title string) string {
	return title + ".txt"
}

// Read a file content
func (p *Page) load(title string) (err error) {
	p.Title = title
	p.Body, err = ioutil.ReadFile(fullTitle(p.Title))
	return err
}

func main() {
	page := new(Page)
	page.Title = filename
	page.Body = []byte("hello world!\nmy email address is: wu.shaohua@foxmail.com\n")

	// Written to the file
	err := page.save()
	if err != nil {
		fmt.Printf("save struct data to text file error: %s\n", err)
		return
	}
	fmt.Printf("save struct data to text file successfully!\n")

	// Read the file
	err = page.load(filename)
	if err != nil {
		fmt.Printf("read the file error: %s\n", err)
		return
	}
	fmt.Println("Read the file contents as following:")
	fmt.Println(string(page.Body))
}
