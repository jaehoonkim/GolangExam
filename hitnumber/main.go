package main

import (
	"fmt"
	"strconv"
	"hitnumber/baseball"
	"errors"
)

func execute(input int, com *baseball.ComNumber) error {
	var err error
	var input_str string
	var in1, in2, in3 int
	
	input_str = strconv.Itoa(input)
	
	in_arr := []byte(input_str)
	in1, _ = strconv.Atoi((string)(in_arr[0]))
	in2, _ = strconv.Atoi((string)(in_arr[1]))
	in3, _ = strconv.Atoi((string)(in_arr[2]))
	
	if in1 == in2 || in1 == in3 || in2 == in3 {
		err = errors.New("입력한 숫자중에 같은 숫자가 있습니다.")
	} else {
		err = com.Playball(in1, in2, in3)
	}

	return err
}

func main() {
	//var userNo1, userNo2, userNo3 int
	
	var com *baseball.ComNumber
	com = baseball.NewComNumber()
	
	fmt.Println("+-------------------------------------------------------+")
	fmt.Println("|\t\t\t숫자야구 ver0.1\t\t\t|")
	fmt.Println("+-------------------------------------------------------+")	

	fmt.Println("컴퓨터의 숫자는 ", com.No1, com.No2, com.No3, " 입니다.")
	
	for {
		var input int
		fmt.Print("예상 숫자를 입력해 주세요(종료 : 0) =>")
		fmt.Scanf("%d",  &input)

		if input == 0 {
			break;
		}
		if input < 100 || input > 999 {
			fmt.Println("3자리의 숫자를 입력해 주세요.")
			continue;
		}
		
		err := execute(input, com)
		if err == nil {
			fmt.Println("홈~~런~~~!!!.")
			break;		
		} else {
			fmt.Println(err)
		}
		
	}
}
