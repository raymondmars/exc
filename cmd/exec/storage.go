package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const dbFileName = ".cmdsaver"

type Storage struct {
	rows []Commander
}

func (s *Storage) LoadAll() {

	filePath := getDbFile()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return
	}

	file, er := os.Open(filePath)
	if er != nil {
		panic(er)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	s.rows = []Commander{}
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "\t")
		s.rows = append(s.rows, Commander{
			path:    fields[0],
			alias:   fields[1],
			command: fields[2],
			args:    fields[3:],
		})
	}

}
func (s *Storage) Add(c Commander) {
	if s.rows == nil {
		s.rows = []Commander{}
	}
	if _, isExists := s.FindByPathAndAlias(c.path, c.alias); isExists {
		return
	}
	s.rows = append(s.rows, c)
	//save to disk
	writeToFile([]Commander{c}, os.O_APPEND|os.O_WRONLY|os.O_CREATE)

}
func (s *Storage) GetAll() []Commander {
	return s.rows
}

func (s *Storage) FindByPath(path string) []Commander {
	var cm = []Commander{}
	for _, c := range s.rows {
		if c.path == path {
			cm = append(cm, c)
		}
	}
	return cm
}
func (s *Storage) FindByPathAndAlias(path string, alias string) (Commander, bool) {
	for _, c := range s.rows {
		if c.path == path && c.alias == alias {
			return c, true
		}
	}
	return Commander{}, false
}

func (s *Storage) FindByAlias(alias string) (Commander, bool) {
	currentPath, crr := os.Getwd()
	if crr != nil {
		panic(crr)
	}
	for _, c := range s.rows {
		if c.path == currentPath && c.alias == alias {
			return c, true
		}
	}
	return Commander{}, false
}

func (s *Storage) Remove(alias string) bool {
	currentPath, crr := os.Getwd()
	if crr != nil {
		panic(crr)
	}
	for index, c := range s.rows {
		if c.path == currentPath && c.alias == alias {
			copy(s.rows[index:], s.rows[index+1:])
			//save to disk
			writeToFile(s.rows[:len(s.rows)-1], os.O_RDWR|os.O_TRUNC)
			return true
		}
	}
	return false
}

func getDbFile() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%s", dir, dbFileName)

}
func writeToFile(cmds []Commander, flag int) {
	filePath := getDbFile()
	f, err := os.OpenFile(filePath, flag, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, c := range cmds {
		f.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", c.path, c.alias, c.command, strings.Join(c.args, "\t")))
	}

}
