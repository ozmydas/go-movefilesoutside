package main

import (
	_ "errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Item struct {
	Name, Path, Type string
}

type Items struct {
	Data []Item
}

/******/

func (rows *Items) AddItem(row Item) []Item {
	rows.Data = append(rows.Data, row)
	return rows.Data
}

/******/

func main() {
	dirname := "./FILES"
	newDir := "OUTPUT"
	path, _ := os.Getwd()

	/****/
	var mode string
	flag.StringVar(&mode, "opt", "copy", "Mode : copy / move")
	flag.Parse()
	/****/

	// log.Println(result)
	MakeOutputDir(filepath.Join(path, "./OUTPUT"))
	ProsesDir(dirname, newDir, mode)
}

func ProsesDir(dirname, newDir, mode string) {
	result, _ := scanDir(dirname)

	for _, source := range result {
		// log.Printf("%v - %v", source.Name, source.Type)
		if source.Type == "file" {
			if mode == "move" {
				MoveToDir(source.Name, source.Path, newDir)
			} else if mode == "copy" {
				CopyToDir(source.Name, source.Path, newDir)
			} else {
				log.Println("invalid option", mode)
			}
		} else {
			ProsesDir(dirname+"/"+source.Name, newDir, mode)
		}
	}
}

/** we scan whats inside directory **/
func scanDir(dirname string) ([]Item, error) {
	items := []Item{}
	all := Items{items}
	path, _ := os.Getwd()

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return items, nil
	}

	for _, file := range files {
		// log.Println(file.IsDir())

		tipe := "folder"

		if file.IsDir() == false {
			tipe = "file"
		}

		temp := Item{
			Name: file.Name(),
			Path: filepath.Join(path, dirname+"/"),
			Type: tipe,
		}
		all.AddItem(temp)
	}

	return all.Data, nil
} // end func

/** if output folder not found, create it **/
func MakeOutputDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
		log.Println("-- Creating OUTPUT directory --")
	}
}

/** move file inside folder to output folder **/

func MoveToDir(filename, dir, newDir string) error {
	path, _ := os.Getwd()
	oldLocation := filepath.Join(dir, filename)
	newLocation := filepath.Join(path, newDir+"/"+filename)

	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(filename, "Moved!")
	return nil
}

/** copy file inside folder to output folder **/
func CopyToDir(filename, dir, newDir string) {
	path, _ := os.Getwd()
	sourceFile := filepath.Join(dir, filename)
	destinationFile := filepath.Join(path, newDir+"/"+filename)

	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		log.Println("Error creating", destinationFile)
		log.Println(err)
		return
	}

	log.Println(filename, "Copied!")
}
