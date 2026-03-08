package main
import "fmt"

// каждое число помещаем в массив и сравниваем элементы 
func is_palindrome ( x int ) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

    s := []int{}
	h := x

    for i := 0; h != 0; i++ {
        s = append(s, h%10)
		h /= 10
    }	
	//fmt.Println(s)
	
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s) - 1 - i]{
			return false
		}
	}
	return true
}

// создаём 2е число - первернутое первоначальное и сравниваем
func is_palindrome2 ( x int ) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	direct_x := x
	reverse_x := 0
	
	for direct_x > 0 {
		reverse_x = reverse_x*10 + direct_x%10
		direct_x /= 10
	}
	fmt.Println(reverse_x)
	return x == reverse_x
}

func main() {	
	fmt.Println("is it palindrome? - ", is_palindrome2(1213121))
}
