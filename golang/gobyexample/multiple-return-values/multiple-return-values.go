package main
import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main()  {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

// アンダースコア変数（ _ ）は，一度宣言するけど使わない変数
// https://qiita.com/penguin_dream/items/c1df36040b3fc6d42945
