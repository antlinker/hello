package hello

import "fmt"

// Hello 输出Hello,...
func Hello(v interface{}) string {
	return fmt.Sprintf("Hello,%v", v)
}
