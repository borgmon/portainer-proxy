package helper

import "strings"

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func ParseCSVSlice(str string) []string {
	m := []string{}
	p := strings.Split(str, ",")
	for _, e := range p {
		e = strings.TrimSpace(e)
		m = append(m, e)
	}
	if len(m) == 1 && m[0] == "" {
		return nil
	}
	return m
}

func GetPort(str string) string {
	if str == "" {
		return ":8090"
	} else {
		return ":" + str
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
