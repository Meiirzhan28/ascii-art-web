package ascii

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Toascii(text string, banner string) (string, int) {
	f, err := os.Open("server/ascii/" + banner + ".txt")
	if err != nil {
		return "", http.StatusInternalServerError
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	ascii := []string{}
	for scanner.Scan() {
		s := strings.ReplaceAll(scanner.Text(), "/n", "")
		ascii = append(ascii, s)
	}
	if err := scanner.Err(); err != nil {
		return http.StatusText(http.StatusBadRequest), http.StatusBadRequest
	}
	if checker(banner + ".txt") {

		argument := text

		if !check(argument) {
			return http.StatusText(http.StatusBadRequest), http.StatusBadRequest
		}
		return For_Letters(ascii, argument)
	} else {
		return http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError
	}
}

func check(s string) bool {
	for _, i := range s {
		if (i < 32 || i > 126) && i != 10 && i != 13 {
			return false
		}
	}
	return true
}

func checker(a string) bool {
	switch a {
	case "standard.txt":
		if Hash(a) == "ac85e83127e49ec42487f272d9b9db8b" {
			return true
		}
	case "shadow.txt":
		if Hash(a) == "a49d5fcb0d5c59b2e77674aa3ab8bbb1" {
			return true
		}

	case "thinkertoy.txt":
		if Hash(a) == "eef471ad03be9d13027560644dda8359" {
			return true
		}
	}
	return false
}

func For_Letters(s []string, w string) (string, int) {
	e := map[rune][]string{}
	var b string
	var q rune = 32
	for i := 1; i < len(s); i += 9 {
		e[q] = s[i : i+8]
		q++
	}

	r := strings.ReplaceAll(w, "\r\n", "\n")
	k := strings.SplitAfter(r, "\n")

	for _, d := range k {
		if d == "\n" {
			b += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for t := 0; t < len(d); t++ {
					if d[t] >= 32 && d[t] <= 126 {
						b += e[rune(d[t])][i]
					}
				}
				b += "\n"

			}
		}
	}
	return b, 0
}

func Hash(s string) string {
	h := md5.New()
	f, err := os.Open("server/ascii/" + s)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	a := fmt.Sprintf("%x", h.Sum(nil))
	return a
}
