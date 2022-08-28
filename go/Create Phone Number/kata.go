package kata

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	tmp := make([]interface{}, len(numbers))
	for i, v := range numbers {
		tmp[i] = v
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", tmp...)
}
