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

	
func get_words_from(text string) []string{
    words:= regexp.MustCompile(`\pL+('\pL+)*`)
    return words.FindAllString(text, -1)
}
func count_words (words []string) int{
    word_counts := make(map[string]int)
    for _, word :=range words{
        word_counts[word]++
    }
    return len(word_counts);
}
func DifferentWordsCount(x string) int{
	x =strings.ToLower(x)
	return count_words(get_words_from(x))
}
