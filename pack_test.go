
package main

import (
	"testing"
)

func TestGetMaxSideLength(t *testing.T) {

	var tests = []struct {
		l uint64
		max uint64
		want uint64
	}{
		{50,580,550},
		{70,580,560},
		{60,500,480},
		{70,500,490},
	}
	
	for _,test:=range tests{
		if got := GetMaxSideLength(test.l,test.max); got!=test.want{
			t.Errorf("GetMaxSideLength(%d, %d) = %v", test.l, test.max, got)
		}
	}
	
}