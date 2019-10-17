package main
import "fmt"

func StrLen(str string) int {
	count := 0
	for _, element := range str {
		count++
		_ = element
	}
	return count
}

func StrRev(s string) string{
	runes := []rune(s)
	var word string
	count:=StrLen(s) 
	for i:=count-1; i >= 0; i--{
		word=word+string(runes[i])
	}
	return word
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
