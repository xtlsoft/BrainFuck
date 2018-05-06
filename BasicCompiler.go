package BrainFuck

import (
	"fmt"
	
)

type BasicCompiler struct{
	
}

func (this *BasicCompiler) Run(code string){

	rslt, _, _ := this.Calc(code)
	fmt.Println( rslt )

}

func (this *BasicCompiler) Calc(code string) (string, []int64, int) {

	rslt := ""

	bytes := []byte(code)

	var mem []int64
	mem = append(mem, 0)
	current := 0
	isSkip := false
	isSkipCalc := false
	skipped := ""

	for _, v := range bytes {

		if isSkip {
			if isSkipCalc {
				if v != ']'{
					skipped += string(v)
				}else{
					isSkip = false
					skippedRslt, skippedMem, skippedCurrent := this.Calc(skipped)
					rslt += skippedRslt
					isSkipCalc = false
					for skippedMem[skippedCurrent] != 0 {
						skippedRslt, skippedMem, skippedCurrent = this.Calc(skipped)
						rslt += skippedRslt
					}
					skipped = ""
				}
			}else{
				if v == ']'{
					isSkip = false
					continue
				}
			}
		}

		switch v{
			case '+':
				mem[current] += 1
			case '-':
				mem[current] -= 1
			case '>':
				current ++;
				if current >= len(mem) {
					mem = append(mem, 0)
				}
			case '<':
				if current != 0{
					current --;
				}else{
					panic("BrainFuck Runtime Error: Cannot allocate -1.")
				}
			case '.':
				rslt += string(byte(mem[current]))
			case ',':
				fmt.Scanf("%c", &mem[current])
			case '[':
				isSkip = true
				if current == 0 {
					isSkipCalc = false
				}else{
					isSkipCalc = true
				}
			default:
				
		}

	}

	return rslt, mem, current

}

var Basic BasicCompiler