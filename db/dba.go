package db

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

var Projects ProjectTable

func Init(path string) {
	dbPath = path
	if dbPath == "" {
		println("dbPath is nil")
		os.Exit(1)
	}

	// create dir if it's not exist.
	if _, err := os.Stat(filepath.Dir(dbPath)); os.IsNotExist(err) {
		// os.RemoveAll(path)
		if err := os.MkdirAll(filepath.Dir(dbPath), 0777); err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
	// create db file if it's not exist.
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file.Close()
	}
}

func RebuildDb() {
	if dbFile, err := os.OpenFile(dbPath, os.O_RDWR, 0644); err != nil {
		panic(err)
	} else {
		if fileInfo, err := dbFile.Stat(); err != nil {
			panic(err)
		} else {
			if fileInfo.Size() != 0 {
				enc := gob.NewDecoder(dbFile)
				if err := enc.Decode(&Projects); err != nil {
					panic(err)
				} else {
					fmt.Println("rebuild db completed")
				}
			} else {
				println("don't need to rebuild the data !")
			}
		}
		defer dbFile.Close()
	}
}

func SaveDb() {
	if size := Projects.Size(); size == 0 {
		println("don't need to save the data !")
	} else {
		if dbFile, err := os.OpenFile(dbPath, os.O_RDWR, 0644); err != nil {
			panic(err.Error())
		} else {
			enc := gob.NewEncoder(dbFile)
			if err := enc.Encode(Projects); err != nil {
				fmt.Println(err)
			}
			defer dbFile.Close()
		}
	}
}
