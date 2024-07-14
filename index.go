
package main;

import "fmt";
import "calculator/lib";

func main() {
	const t = "|()*/. 8.910 100/40";
	fmt.Printf("To scan: \" %v \"\n", t);
	tokenizer := lib.NewTokenizer(t);
	tkns, err := tokenizer.Parse();
	if err != nil {
		fmt.Println(err);
		return;
	}
	for _, v := range *tkns {
		fmt.Printf("Token {type: %v, value: %v, literal: %v}\n", v.Type, v.Value, v.Literal);
	}
}
