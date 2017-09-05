package main

import (
	"testing"
)

func TestGetMaxSideLength(t *testing.T) {

	var tests = []struct {
		l    int64
		max  int64
		want int64
	}{
		{50, 580, 550},
		{70, 580, 560},
		{60, 500, 480},
		{70, 500, 490},
	}

	for _, test := range tests {
		if got := GetMaxSideLength(test.l, test.max); got != test.want {
			t.Errorf("GetMaxSideLength(%d, %d) = %v", test.l, test.max, got)
		}
	}

}

func TestGetPackSolutionImp(t *testing.T) {
	var tests = []struct {
		l     int64
		w     int64
		h     int64
		wantL int64
		wantW int64
		wantH int64
	}{
		{240, 80, 80, 560, 480, 480},
		{80, 240, 80, 560, 480, 480},
		{80, 80, 240, 560, 480, 480},
		{180, 250, 25, 540, 500, 500},
		{180, 25, 250, 540, 500, 500},
		{250, 25, 180, 540, 500, 500},
		{190, 140, 140, 570, 420, 420},
		{140, 190, 140, 570, 420, 420},
		{140, 140, 190, 570, 420, 420},
		{50, 50, 50, 500, 500, 500},
		//{23, 23, 23, 555, 555, 555},
		//{390, 50, 50, 555, 555, 555},
	}

	for _, test := range tests {
		l, w, h := test.l, test.w, test.h
		wantL, wantW, wantH := test.wantL, test.wantW, test.wantH
		s := GetPackSolutionImp(l, w, h)
		L := s.BoxLength
		W := s.BoxWidth
		H := s.BoxHeigth
		if L != wantL || W != wantW || H != wantH {
			t.Errorf("GetPackSolutionImp(%d, %d,%d) => %v %v %v want %v %v %v", l, w, h, L, W, H, wantL, wantW, wantH)
		}
	}
}
