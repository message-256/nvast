package nvast_test

import (
	"fmt"
	"reflect"
	"testing"
	"errors"
	//"errors"
	"github.com/message-256/nvast"
)

type expectation struct {
	input  string
	delim  [2]rune
	output nvast.Nvast
	err    error
}

func TestCompile(t *testing.T) {
	var output = []expectation{
		{
			input: "1+2+3+(4+5)+5",
			delim: [2]rune{'(', ')'},
			output: nvast.Nvast{
				Flat: []string{"1+2+3+", "()", "+5"},
				Inner: []nvast.Nvast{
					nvast.Nvast{
						Flat: []string{"4+5"},
					},
				},
			},
			err: nil,
		},
		{
			input: "(4+5)1+2+3++5",
			delim: [2]rune{'(', ')'},
			output: nvast.Nvast{
				Flat: []string{"()", "1+2+3++5"},
				Inner: []nvast.Nvast{
					nvast.Nvast{
						Flat: []string{"4+5"},
					},
				},
			},
			err: nil,
		},
		{
			input: "(4+5)1+2+3++5",
			delim: [2]rune{'{', '}'},
			output: nvast.Nvast{
				Flat: []string{"(4+5)1+2+3++5"},
			},
			err: nil,
		},
		{
			input: "{4+5}1+2+3++5",
			delim: [2]rune{'{', '}'},
			output: nvast.Nvast{
				Flat: []string{"{}", "1+2+3++5"},
				Inner: []nvast.Nvast{
					nvast.Nvast{
						Flat: []string{"4+5"},
					},
				},
			},
			err: nil,
		},
		{
			input: "{4+51+2+3++5",
			delim: [2]rune{'{', '}'},
			err: nvast.ErrExprNoEnd,
		},
		{
			input: "}4+51+2+3++5",
			delim: [2]rune{'{', '}'},
			err: nvast.ErrExprKillEarly,
		},
		{
			input: "4+51+2+3++5{",
			delim: [2]rune{'{', '}'},
			err: nvast.ErrExprNoEnd,
		},
		{
			input: "4+51+2+3++5}",
			delim: [2]rune{'{', '}'},
			err: nvast.ErrExprKillEarly,
		},
	}
	for i := range output {
		returned, err := nvast.Compile(output[i].input, output[i].delim)
		if !errors.Is(err, output[i].err) || !reflect.DeepEqual(output[i].output, returned) {
			fmt.Print("with input ", output[i].input, " ")
			fmt.Printf("output = %+v,%v, expected = %+v,%v \n", returned, err, output[i].output,output[i].err)
			t.Errorf("test failed\n")
		}
	}

}
