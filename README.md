## Tokenize

Takes any string as text, tokenization characters as runes `rune`, and returns results a slice of string tokens `[]string`. Where each item in the result set are the tokenized words followed by the runes to tokenize on, in order.

## Example
Print a set of all substrings tokenized by the following punctuation characters `['.', '!', '?', ',', ' ']`.

```go
func main() {
	str := "Lorem ipsum dolor sit amet! consectetur adipiscing elit. Nunc viverra, quam sit amet varius accumsan, augue mi viverra lacus, sed hendrerit justo magna eu augue. Aliquam in pretium justo. Nulla pulvinar tempus tempus. Nulla luctus lacus sed gravida congue. Aliquam a est magna. Nullam condimentum dui ut tortor placerat accumsan. Nullam eu ligula ante. Quisque finibus est eu lorem gravida, sit amet hendrerit metus pellentesque. Fusce vitae arcu sem."

	var punctuation = []rune{'.', '!', '?', ',', ' '}

	words := tokenize.Create(str, punctuation)
	for _, s := range words {
		fmt.Println(s)
	}
}
```

The above returns `all text without punctuation`.

## License
This project is released under the terms of the MIT LICENSE.
