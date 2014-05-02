package baseball

import (
	"math/rand"
	"time"
	"errors"
	"fmt"
)

type ComNumber struct {
	No1 int
	No2 int
	No3 int
}

func NewComNumber() *ComNumber {
	var com *ComNumber
	com = new (ComNumber)
	com.ComRandomNumber()
	return com
}

func (c *ComNumber) ComRandomNumber() {
	rand.Seed(time.Now().UnixNano())
	c.No1 = rand.Intn(8) + 1
	
	second:
	c.No2 = rand.Intn(8) + 1
	if c.No1 == c.No2 {
		goto second
	}
	
	third:
	c.No3 = rand.Intn(8) + 1
	if c.No1 == c.No3 || c.No2 == c.No3 {
		goto third
	}	
}

func (com *ComNumber)Playball(in1, in2, in3 int) (err error) {
	var strk_count, ball_count int
	if com.No1 == in1 {
		strk_count++
	} else {
		if com.No2 == in1 || com.No3 == in1 {
			ball_count++
		}
	}
	if com.No2 == in2 {
		strk_count++
	} else {
		if com.No1 == in2 || com.No3 == in2 {
			ball_count++
		}
	}
	if com.No3 == in3 {
		strk_count++
	} else {
		if com.No1 == in3 || com.No2 == in3 {
			ball_count++
		}
	}
	
	var strMsg string
	if strk_count == 3 {
		err = nil
	} else if strk_count > 0 || ball_count > 0 {
		if strk_count > 0 {
			strMsg = fmt.Sprintf("%d 스트라이크 ", strk_count)
		}
		
		if ball_count > 0 {
			strMsg += fmt.Sprintf("%d 볼", ball_count)
		}
		
		if len(strMsg) > 0 {
		err = errors.New(strMsg)
		} else {
			err = errors.New("아웃")
		}
	}
	
	
	return err
}

