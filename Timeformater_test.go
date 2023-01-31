package timeformater

import (
	"fmt"
	"testing"
)


func Test_Time(t *testing.T){
	x:="2020-13-1"
	timeFormater:=TimeFormater{}

	time,err:=timeFormater.ConvertDateToTime(x,"")
	if err==nil{
		fmt.Println(err)
		t.Fatalf("Expected a fail %v",x)
	}
	y:="2020-12-1"
	time,err=timeFormater.ConvertDateToTime(y,"")
	if err!=nil{
		fmt.Println(err)
		t.Fatalf("Expected a fail %v",y)
	}
	fmt.Println(time)
}