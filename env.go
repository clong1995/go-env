package env

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var env map[string]string

func init() {
	exePath, err := os.Executable()
	if err != nil {
		log.Println(err)
		return
	}
	dir := filepath.Dir(exePath)
	envPath := path.Join(dir, "/.env")
	if _, err = os.Stat(envPath); err != nil {
		dir, err = os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}
		envPath, err = findEnv(dir)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	data, err := os.ReadFile(envPath)
	if err != nil {
		log.Fatalln(err)
		return
	}

	str := strings.ReplaceAll(string(data), "\r\n", "\n")
	row := strings.Split(str, "\n")
	env = make(map[string]string)
	for _, s := range row {
		if s == "" || strings.HasPrefix(s, "#") {
			continue
		}
		cell := strings.Split(s, " = ")
		if len(cell) != 2 {
			err = fmt.Errorf("env row error:%s", s)
			log.Println(err)
			return
		}
		key := strings.Trim(cell[0], " ")
		value := strings.Trim(cell[1], " ")
		env[key] = value
	}
}

func Env(key string) string {
	return env[key]
}

func findEnv(dir string) (envPath string, err error) {
	envPath = path.Join(dir + "/.env")
	if _, err = os.Stat(envPath); err != nil {
		dir = path.Join(dir, "..")
		if dir == "/" {
			err = fmt.Errorf(".env not found")
			return
		}
		return findEnv(dir)
	}
	return
}
