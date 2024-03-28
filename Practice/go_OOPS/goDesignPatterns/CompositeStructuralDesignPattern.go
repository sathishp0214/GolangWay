package main

import "fmt"

//Composite Design Pattern - When single object(EX: file, leaf) and Group of objects(Ex: folder or tree) should be handled in similar way.
//Use case Examples-- file and folder, Leaf and Tree, Hierachy of objects

type Component interface {
	Search(string)
}

type File struct {
	name string
}

func (f File) Search(filename string) {
	fmt.Printf("searching %v in folder--", filename)
	if f.name == filename {
		fmt.Println("file found in the folder")
	}
}

type Folder struct {
	name  string
	files []Component
}

func (f *Folder) Search(filename string) {
	for _, file := range f.files {
		fmt.Println("searching inside folder---", f.name)
		file.Search(filename) //Searching file in the folder
	}
}

func (f *Folder) Add(component Component) { //Adding component's interface implements "file" struct
	f.files = append(f.files, component)
}

func main() {
	file1 := File{"file1"}
	file2 := File{"file2"}
	file3 := File{"file3"}

	Folder := &Folder{name: "Folder1"}
	Folder.Add(file1)
	Folder.Add(file2)
	Folder.Add(file3)

	Folder.Search("file2") //searching file2 in folder1
}
