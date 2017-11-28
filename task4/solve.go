package main
import "unicode"

func RemoveEven(input []int) []int{
	x := make([]int,len(input))
	var j int = 0

	for i :=0; i< len(input); i++ {
		if (input[i] %2 ==1) {
			x[j] =input[i]
			j =j+1
		}
	}
	return x[0:j]
}	

func PowerGenerator(x int) func() uint{
	i:= uint(x)
	return func() (ret uint) {
		ret =i
		i = i*uint(x)
	return 
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
