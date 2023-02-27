package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	temp := make(map[string]struct{})
	for i := 0; i < len(developers); i++ {
		_, ok := temp[developers[i].Name]
		if !ok {
			temp[developers[i].Name] = struct{}{}
		}
	}
	str := make([]string, 0)
	for k, _ := range temp {
		str = append(str, k)
	}
	return str
}

func main() {
	hello := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Alan"},
		Developer{Name: "Jennifer"},
		Developer{Name: "Graham"},
		Developer{Name: "Paul"},
		Developer{Name: "Alan"},
	}
	fmt.Println(FilterUnique(hello))
}
