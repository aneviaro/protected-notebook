package credentials

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	users []Credentials
	once  sync.Once
	mtx   sync.RWMutex
)

func init() {
	once.Do(initializeUsers)
}

func initializeUsers() {
	f, err := os.Open("resources/creds/users")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		h := md5.New()
		h.Write([]byte(parts[1]))
		c := Credentials{
			Password: h.Sum(nil),
			Username: parts[0],
		}
		fmt.Println(c.Password)
		mtx.Lock()
		users = append(users, c)
		mtx.Unlock()
	}
}

type Credentials struct {
	Password []byte `json: "password"`
	Username string `json: "username"`
}

func CheckUser(creds Credentials) error {
	mtx.RLock()
	defer mtx.RUnlock()
	for _, c := range users {
		if creds.Username == c.Username {
			if string(creds.Password) == string(c.Password) {
				return nil
			} else {
				return errors.New("Wrong password!")
			}
		}
	}
	return errors.New("You are not registered yet, ask Alex to make you an account")
}
