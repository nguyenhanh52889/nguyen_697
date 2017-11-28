package main
import ( 
	"fmt"
	"strings"
	"regexp"
)

func main() {
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println("==========================")
	fmt.Println(result) // Должно напечататься [3 5]
	fmt.Println("==========================")
	gen := PowerGenerator(3)
	fmt.Println(gen()) // Должно напечатать 3
	fmt.Println(gen()) // Должно напечатать 9
	fmt.Println(gen()) // Должно напечатать 27
	fmt.Println(gen()) // Должно напечатать 81
	fmt.Println(gen()) // Должно напечатать 243
	fmt.Println("==========================")
	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
	// Должно напечатать 2

}
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
