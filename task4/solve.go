package main
import "unicode"

func RemoveEven(input []int) []int{
	x := make([]int, 0)

	for i :=0; i< len(input); i++ {
		if (input[i] %2 ==1) {
			x = append(x, input[i])
		}
	}
	return x
}	

func PowerGenerator(x int) func() int{
	i:= 1
	return func()  {
		i *= x
		return i
	}
}

	
func DifferentWordsCount(x string) int {
    word := ""
    set := make(map[string]bool)
    ans := 0
    for _, c := range (x + " ") {
        if unicode.IsLetter(c) {
            word += string(unicode.ToLower(c))
        } else if word != "" {
            if !set[word] {
                ans += 1
            }
            set[word] = true
            word = ""
        }
    }
    return ans
}
