package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var Bien *int64
type Config struct {
	//day la mot bien co kieu struct
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}
func readFile(cfg * Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder :=yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
func GanGiaTriChoBien() {
	var i int64 = 64
	Bien = &i
}
func main() {
	GanGiaTriChoBien()
	fmt.Println(Bien)
	var cfg Config
	readFile(&cfg)
	fmt.Printf("%+v", cfg)
	fmt.Println(cfg)
}