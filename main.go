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

var allcount int
var limit int

/******/

func (rows *Items) AddItem(row Item) []Item {
	rows.Data = append(rows.Data, row)
	return rows.Data
} // end func

/******/

func main() {
	var mode string
	var inputDir string
	var outputDir string

	// define option flag and parse
	flag.StringVar(&mode, "mode", "copy", "Mode : copy / move")
	flag.StringVar(&inputDir, "in", "FILES", "Directory to Scan")
	flag.StringVar(&outputDir, "out", "OUTPUT", "Directory to store result")
	flag.IntVar(&limit, "limit", 1, "Limit of max files to execute")
	flag.Parse()

	/****/

	dirname := "./" + inputDir
	newDir := outputDir

	// before begin operation, make sure output folder was created
	MakeOutputDir(outputDir)

	// main operation
	ProsesDir(dirname, newDir, mode)
} // end func

func ProsesDir(dirname, newDir, mode string) {
	if allcount > limit {
		log.Println("Limit Reached - only can process max", limit, "files, to bypass limit please add flag -limit NUMBER")
		return

	}

	result, count, err := scanDir(dirname)

	if err != nil {
		log.Println("ERROR detected -", err)
	}

	allcount = allcount + count

	// check if any file to proccess
	if allcount > 1 {
		log.Println("-- Files Detected :", allcount, "file(s) --")
	} else {
		log.Println("-- Empty Folder! Can't Find Files to Process Or File Already Moved --")
		return
	}

	// looping each file and folder from scanned result
	i := 1
	for _, source := range result {
		// check limit for first directory
		if allcount == 0 {
			if i >= limit {
				log.Println("Limit Reached - only can process max", limit, "files, to bypass limit please add flag -limit NUMBER")
				i++
				return
			}
		}

		// execute
		if source.Type == "file" {
			// if scanned result is file, decide next step based on opt flag to copy or move files
			if mode == "move" {
				err := MoveToDir(source.Name, source.Path, newDir)
				if err != nil {
					log.Println("ERROR on", source.Name, "-", err)
				}
			} else if mode == "copy" {
				err := CopyToDir(source.Name, source.Path, newDir)
				if err != nil {
					log.Println("ERROR on", source.Name, "-", err)
				}
			} else {
				log.Println("-- ERROR Invalid Mode! please use -mode flag with copy or paste", mode, "--")
				return
			}
		} else {
			// if scanned result was a directory, recursive scan that directory
			ProsesDir(dirname+"/"+source.Name, newDir, mode)
		}
	}
} // end func

/** we scan whats inside directory **/
func scanDir(dirname string) ([]Item, int, error) {
	items := []Item{}
	all := Items{items}
	path, _ := os.Getwd()
	count := 0

	// scanning directoruy here
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return items, count, err
	}

	for _, file := range files {
		tipe := "folder"

		// if detected as file
		if file.IsDir() == false {
			tipe = "file"
			count++
		}

		// store to array struct
		temp := Item{
			Name: file.Name(),
			Path: filepath.Join(path, dirname+"/"),
			Type: tipe,
		}
		all.AddItem(temp)
	}

	return all.Data, count, nil
} // end func

/** if output folder not found, create it **/
func MakeOutputDir(path string) {
	dirpath, _ := os.Getwd()
	fpath := filepath.Join(dirpath, "./"+path)
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		// create new directory if not exist
		os.Mkdir(fpath, os.ModePerm)
		log.Println("-- Creating", path, "directory --")
	}
} // end func

/** move file inside folder to output folder **/
func MoveToDir(filename, dir, newDir string) error {
	path, _ := os.Getwd()
	// define old location and new location to moving file
	oldLocation := filepath.Join(dir, filename)
	newLocation := filepath.Join(path, newDir+"/"+filename)

	// execute
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		// log.Println(err)
		return err
	}

	log.Println(filename, "Moved!")
	return nil
} // end func

/** copy file inside folder to output folder **/
func CopyToDir(filename, dir, newDir string) error {
	path, _ := os.Getwd()
	// define file fullpath and target to copying file
	sourceFile := filepath.Join(dir, filename)
	destinationFile := filepath.Join(path, newDir+"/"+filename)

	// read content of source file
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		// log.Println(err)
		return err
	}
	// write content to target destination file
	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		// log.Println(err)
		return err
	}

	log.Println(filename, "Copied!")
	return nil
} // end func
