package main

import (
	"calculator/lib"
	"fmt"
	"strconv"
	"strings"
)

func padEnd(s string, length int) string {
	builder := strings.Builder{};
	builder.WriteString(s);
	for i := 0; i < length - len(s); i++ {
		builder.WriteByte(' ');
	}
	return builder.String();
}

func main() {
	const t = "|()*/. 8.910 100/40 p 40 root 50 alice margatroid";
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
		fmt.Printf("Token {%v %v %v}\n",
			padEnd(fmt.Sprintf("type: %v,", v.Type.String()), 20),
			padEnd(fmt.Sprintf("value: %v,", string(v.Value)), 20),
			padEnd(fmt.Sprintf("literal: %v,", strconv.FormatFloat(deref_literal, 'f', -1, 64)), 20),
		);
	}
}

