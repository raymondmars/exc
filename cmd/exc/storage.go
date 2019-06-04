package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
			Path:    fields[0],
			Alias:   fields[1],
			Command: fields[2],
			Args:    fields[3:],
		})
	}

}
func (s *Storage) Add(c Commander) {
	if s.rows == nil {
		s.rows = []Commander{}
	}
	if _, isExists := s.FindByPathAndAlias(c.Path, c.Alias); isExists {
		return
	}
	s.rows = append(s.rows, c)
	//save to disk
	writeToFile([]Commander{c}, os.O_APPEND|os.O_WRONLY|os.O_CREATE)

}
func (s *Storage) GetAll() []Commander {
	return s.rows
}

func (s *Storage) GetCurrentAll() []Commander {
	currentPath, crr := os.Getwd()
	if crr != nil {
		panic(crr)
	}
	return s.FindByPath(currentPath)
}
func (s *Storage) FindByPath(path string) []Commander {
	var cm = []Commander{}
	for _, c := range s.rows {
		if c.Path == path {
			cm = append(cm, c)
		}
	}
	return cm
}
func (s *Storage) FindByPathAndAlias(path string, alias string) (Commander, bool) {
	for _, c := range s.rows {
		if c.Path == path && c.Alias == alias {
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
		if c.Path == currentPath && c.Alias == alias {
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
		if c.Path == currentPath && c.Alias == alias {
			copy(s.rows[index:], s.rows[index+1:])
			//save to disk
			writeToFile(s.rows[:len(s.rows)-1], os.O_RDWR|os.O_TRUNC)
			return true
		}
	}
	return false
}

func getDbFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%s", home, dbFileName)

}
func writeToFile(cmds []Commander, flag int) {
	filePath := getDbFile()
	f, err := os.OpenFile(filePath, flag, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, c := range cmds {
		_, err := f.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", c.Path, c.Alias, c.Command, strings.Join(c.Args, "\t")))
		if err != nil {
			log.Fatal(err)
		}
	}

}
