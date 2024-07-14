
package main;

import "fmt";
import "calculator/lib";

func main() {
	const t = "|()*/. 8.910 100/40 p 40 root 50 alice margatroid;";
	fmt.Printf("To scan: \" %v \"\n", t);
	tokenizer := lib.NewTokenizer(t);
	tkns, err := tokenizer.Parse();
	if err != nil {
		fmt.Println(err);
		return;
	}
	for _, v := range *tkns {
		var deref_literal float64;
		if v.Literal != nil {
			deref_literal = *v.Literal;
		}
		fmt.Printf("Token {type: %v, value: %v, literal: %v}\n", v.Type, v.Value, deref_literal);
	}
}
