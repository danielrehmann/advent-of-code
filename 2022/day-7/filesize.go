package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Calculate(scanner *bufio.Scanner) *System {
	basedirr, err := regexp.Compile(`\$ cd \/`)
	if err != nil {
		panic(err)
	}
	updirr, err := regexp.Compile(`\$ cd \.\.`)
	if err != nil {
		panic(err)
	}
	dirChanger, err := regexp.Compile(`\$ cd (\w+)`)
	if err != nil {
		panic(err)
	}
	lsr, err := regexp.Compile(`\$ ls`)
	if err != nil {
		panic(err)
	}
	lsDirResr, err := regexp.Compile(`dir (\w+)`)
	if err != nil {
		panic(err)
	}
	lsFileResr, err := regexp.Compile(`(\d+) (\w+)`)
	if err != nil {
		panic(err)
	}
	basedir := dir{name: "/", parentDir: nil}
	system := System{baseDir: &basedir, currentDir: &basedir}

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		switch {
		case basedirr.MatchString(text):
			system.moveToBaseDir()
		case updirr.MatchString(text):
			system.moveDirUp()
		case dirChanger.MatchString(text):
			dirName := dirChanger.FindStringSubmatch(text)[1]
			system.moveToDir(dirName)
		case lsr.MatchString(text):
			fmt.Println("ls now")
		case lsDirResr.MatchString(text):
			dirName := lsDirResr.FindStringSubmatch(text)[1]
			fmt.Printf("ls has dir %s\n", dirName)
			system.addDir(dirName)
		case lsFileResr.MatchString(text):
			fileRes := lsFileResr.FindStringSubmatch(text)
			fileSize, err := strconv.Atoi(fileRes[1])
			if err != nil {
				panic(err)
			}
			fileName := fileRes[2]
			fmt.Printf("ls has file %s with size %d \n", fileName, fileSize)
			system.addFile(fileName, fileSize)
		}

	}
	system.moveToBaseDir()
	return &system
}

//go:generate stringer -type=System

type System struct {
	baseDir    *dir
	currentDir *dir
	allDirs    []*dir
	allFiles   []*file
}

func (s *System) allDirsMoreThan(size int) []*dir {
	res := make([]*dir, 0)
	for _, v := range s.allDirs {
		if v.Size() >= size {
			res = append(res, v)
		}
	}
	return res
}

func (s *System) allDirsLessThan(maxSize int) ([]*dir, int) {
	res := make([]*dir, 0)
	sum := 0
	for _, v := range s.allDirs {
		if v.Size() < maxSize {
			res = append(res, v)
			sum += v.Size()
		}
	}
	return res, sum
}

func (s *System) addDir(newDirName string) {
	dir := dir{name: newDirName, parentDir: s.currentDir}
	s.allDirs = append(s.allDirs, &dir)
	s.currentDir.addDir(&dir)
}

func (s *System) addFile(fileName string, fileSize int) {
	file := file{name: fileName, size: fileSize}
	s.allFiles = append(s.allFiles, &file)
	s.currentDir.addFile(&file)
}

func (s *System) moveToBaseDir() {
	s.currentDir = s.baseDir
	fmt.Printf("Moved to %s\n", s.currentDir.name)
}

func (s *System) moveDirUp() {
	s.currentDir = s.currentDir.parentDir
	fmt.Printf("Moved to %s\n", s.currentDir.name)
}

func (s *System) moveToDir(dirName string) {
	for _, v := range s.currentDir.dirs {
		if v.name == dirName {
			s.currentDir = v
			fmt.Printf("Moved to %s\n", v.name)
			return
		}
	}
	panic("Could not find dir " + dirName)
}

type dir struct {
	name      string
	parentDir *dir
	files     []*file
	dirs      []*dir
}

func (d *dir) addDir(newDir *dir) {
	d.dirs = append(d.dirs, newDir)
}
func (d *dir) addFile(newFile *file) {
	d.files = append(d.files, newFile)
}

func (d *dir) Size() int {
	size := 0
	for _, v := range d.dirs {
		size = size + v.Size()
	}
	for _, v := range d.files {
		size += v.Size()
	}
	return size
}

func (d *file) Size() int {
	return d.size
}

type file struct {
	name string
	size int
}
type Sizer interface {
	Size() int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("file coul not be opened")
	}
	scanner := bufio.NewScanner(file)
	system := Calculate(scanner)

	fmt.Println(system.baseDir.Size())
	remaining := 70000000 - system.baseDir.Size()
	needed := 30000000 - remaining
	fmt.Println(system.allDirsLessThan(100000))
	min := system.baseDir
	for _, v := range system.allDirsMoreThan(needed) {
		if v.Size() < min.Size() {
			min = v
		}
	}
	fmt.Println(needed, min.name, min.Size())
}
