package tokenizeexample

import "fmt"

func main() {
	str := "Lorem ipsum dolor sit amet! consectetur adipiscing elit. Nunc viverra, quam sit amet varius accumsan, augue mi viverra lacus, sed hendrerit justo magna eu augue. Aliquam in pretium justo. Nulla pulvinar tempus tempus. Nulla luctus lacus sed gravida congue. Aliquam a est magna. Nullam condimentum dui ut tortor placerat accumsan. Nullam eu ligula ante. Quisque finibus est eu lorem gravida, sit amet hendrerit metus pellentesque. Fusce vitae arcu sem."

	var punctuation = []rune{'.', '!', '?', ',', ' '}

	words := Create(str, punctuation)
	for _, s := range words {
		fmt.Println(s)
	}
}
