package nvast

import (
	"errors"
)

type Nvast struct {
	Flat  []string
	Inner []Nvast
}
var ErrExprNoEnd error = errors.New("found delimiter[1] but no delimiter[2]")
var ErrExprKillEarly error = errors.New("found a stray delimiter[1]")
func compile(input string, delims [2]rune) (Nvast, int, error) {
	var returned Nvast
	var collective error
	var n int
	var err error
	var end int
	for i := 0; i < len(input); i++ {
		if rune(input[i]) == delims[0] {
			returned.Inner = append(returned.Inner,Nvast{})
			if input[:i] != "" {
				returned.Flat = append(returned.Flat,input[:i])
			}
			returned.Flat = append(returned.Flat,string(delims[0]) + string(delims[1]))
			returned.Inner[len(returned.Inner)-1],n,err = compile(input[i+1:],delims)
			collective = errors.Join(collective,err)
			if n+i+1 >= len(input){
				return returned,n,errors.Join(collective,ErrExprNoEnd)
			}
			input = input[n+i+2:]
			end+=n+2
			i = 0
			//not the prettiest solution
			//but it works (probably)
			if rune(input[i]) == delims[1] {
				return returned,end,collective
			}
		} else if rune(input[i]) == delims[1] {
			returned.Flat = append(returned.Flat,input[:i])
			return returned,end,collective
		}
		end++
	}
	returned.Flat = append(returned.Flat, input)

	return returned, end, collective
}

// i dont care how long the string is
func Compile(input string, delims [2]rune) (Nvast, error) {
	returned, n, err := compile(input, delims)
	if n != len(input) {
		return Nvast{},errors.Join(err,ErrExprKillEarly)
	}
	if err != nil {
		returned = Nvast{}
	}
	return returned, err
}
