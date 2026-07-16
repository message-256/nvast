package nvast
import (
	"errors"
)
type Nvast struct {
	Flat []string
	Inner []Nvast

}
func compile(input string,delims [2]rune) (Nvast,int,error){
	var returned Nvast
	var collective error
	var n int
	var err error
	var end int
	if input[0] == '(' {
		returned.Inner = append(returned.Inner,Nvast{Flat:nil})
			returned.Flat = append(returned.Flat,string(delims[0]) + string(delims[1]))
			returned.Inner[len(returned.Inner)-1] ,n,err = compile(input[1:],delims);
			collective = errors.Join(collective,err)
			if (rune(input[n]) != delims[1]){
				collective = errors.Join(collective,errors.New("found " + string(delims[0]) + " but no" + string(delims[1]) + " with remaining " + input[n:] ))
			}
			input = input[n+1:]
			end+=n+1
	}
	for i := 0; i<len(input); i++ {
		end++;
		if rune(input[i]) == delims[0] {
			returned.Inner = append(returned.Inner,Nvast{Flat:nil})
			returned.Flat = append(returned.Flat,input[:i])
			returned.Flat = append(returned.Flat,string(delims[0]) + string(delims[1]))
			returned.Inner[len(returned.Inner)-1] ,n,err = compile(input[i+1:],delims);
			collective = errors.Join(collective,err)
			if (rune(input[n+i]) != delims[1]){
				collective = errors.Join(collective,errors.New("found " + string(delims[0]) + " but no" + string(delims[1]) + " with remaining " + input[i+n:] ))
			}
			input = input[i+n+1:]
			end+=n+1
			i = 0;
		} else if rune(input[i]) == delims[1] {
			returned.Flat = append(returned.Flat,input[:i])
			return returned,end,nil
		}
	}
	returned.Flat = append(returned.Flat,input)

	return returned,end,collective
}
// i dont care how long the string is 
func Compile(input string,delims [2]rune) (Nvast,error){
	returned,n,err := compile(input,delims)
	if n != len(input) {
		return returned,errors.Join(err,errors.New("found a stray " + string(delims[1])))
	}
	return returned,err
}
