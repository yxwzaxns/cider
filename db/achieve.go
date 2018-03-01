package db

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func (p *ProjectTable) FindByID(id int) []Project {
	var res []Project
	if id <= p.Size() {
		res = make([]Project, 1)
		res[0] = *(*p)[id-1]
	}
	return res
}

func (p *ProjectTable) FindByName(projectName string) []Project {
	var res []Project
	if 0 < p.Size() {
		for _, _p := range *p {
			if _p.ProjectName == projectName {
				res = make([]Project, 1)
				res[0] = *_p
			}
		}
	}
	return res
}

func (p *ProjectTable) FindAll() []Project {
	res := make([]Project, p.Size())
	for index := 0; index < p.Size(); index++ {
		res[index] = *(*p)[index]
	}
	return res
}

func (p *ProjectTable) Add(project *Project) {
	*p = append(*p, project)
}

func (p *ProjectTable) Size() int {
	return len(*p)
}

func Init(path string) {
	path = "db/cider.db"
	dbPath = path
	if dbPath == "" {
		println("dbPath is nil")
		os.Exit(1)
	}

	// create dir if it's not exist.
	if _, err := os.Stat(filepath.Dir(dbPath)); os.IsNotExist(err) {
		// os.RemoveAll(path)
		println(filepath.Dir(dbPath))
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

func (p *ProjectTable) RebuildDb() {
	if dbFile, err := os.OpenFile(dbPath, os.O_RDWR, 0644); err != nil {
		panic(err)
	} else {
		if fileInfo, err := dbFile.Stat(); err != nil {
			panic(err)
		} else {
			if fileInfo.Size() != 0 {
				enc := gob.NewDecoder(dbFile)
				if err := enc.Decode(p); err != nil {
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

func (p *ProjectTable) SaveDb() {
	if size := p.Size(); size == 0 {
		println("don't need to save the data !")
	} else {
		if dbFile, err := os.OpenFile(dbPath, os.O_RDWR, 0644); err != nil {
			panic(err.Error())
		} else {
			enc := gob.NewEncoder(dbFile)
			if err := enc.Encode(p); err != nil {
				fmt.Println(err)
			}
			defer dbFile.Close()
		}
	}
}
