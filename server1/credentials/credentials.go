package credentials

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"sync"
)

var (
	users []Credentials
	once sync.Once
	mtx sync.RWMutex
)

func init(){
	once.Do(initializeUsers)
}

func initializeUsers(){
	f, err:= os.Open("resources/creds/users")
	if err!=nil{
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		line:=scanner.Text()
		parts:=strings.Split(line, " ")
		c:=Credentials{
			Password: parts[1],
			Username: parts[0],
		}
		mtx.Lock()
		users=append(users, c)
		mtx.Unlock()
	}
}

type Credentials struct {
	Password string `json: "password"`
	Username string `json: "username"`
}

func CheckUser(creds Credentials)error{
	mtx.RLock()
	defer mtx.RUnlock()
	for _,c:=range users {
		if creds.Username==c.Username{
			if c.Password==creds.Password{
				return nil
			} else {
				return errors.New("Wrong password!")
			}
		}
	}
	return errors.New("You are not registered yet, ask Alex to make you an account")
}