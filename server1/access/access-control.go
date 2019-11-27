package access

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func CheckAccess(username string, res string) error {
	f, err := os.Open("atl")
	if err != nil {
		return errors.New("Can't check your access")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if parts[0] == username {
			for _, part := range parts[1:] {
				if part == res {
					return nil
				}
			}
		}
	}
	return errors.New("You don't have any access to this resource")
}

func GrantAccess(filename string, username string) error {
	f, err := os.Open("atl")
	if err != nil {
		return errors.New("Can't grant access")
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	i := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if strings.Contains(lines[i], username) {
			if strings.Contains(lines[i], filename) {
				return nil
			}
			lines[i] = lines[i] + " " + filename
		}
		i++
	}
	f.Close()
	f, err = os.OpenFile("atl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return errors.New("Can't grant access")
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}
