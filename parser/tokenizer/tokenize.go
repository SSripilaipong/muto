package tokenizer

import (
	"fmt"
	"io"

	"phi-lang/common/rslt"
)

func Tokenize(reader io.RuneReader) func() rslt.Of[Token] {
	var carry []rune
	var err error

	return func() rslt.Of[Token] {
		if err != nil {
			return rslt.Error[Token](err)
		}

		builder := newTokenBuilder()
		for _, c := range carry {
			builder.Consume(c)
		}
		carry = carry[:0]

		for !builder.IsDone() {
			var r rune
			r, _, err = reader.ReadRune()
			if err == io.EOF {
				break
			} else if err != nil {
				return rslt.Error[Token](fmt.Errorf("cannot get next token: %w", err))
			}
			builder.Consume(r)
		}

		token, c := builder.Build().Return()
		carry = c
		return rslt.Value(token)
	}
}
