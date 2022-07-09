package coverage

import (
	
	"os"
	"testing"
	"time"
	"errors"
	"strconv"
	//"github.com/stretchr/testify/assert"
	
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T) {
	people := People{{"Peter", "Petrosyan", time.Now()}, {"Hovo", "Hovoyan", time.Now()}}
	got := people.Len()
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSort (t *testing.T) {
	time1 := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	time2 := time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local)

	people  := People{{"Hovo", "Hovoyan", time2}, {"Hovo", "Hovoyan", time1}}
	people1 := People{{"Peter", "Petrosyan", time2}, {"Hovo", "Hovoyan", time2}}
	people3 := People{{"Peter", "Petrosyan", time2}, {"Peter", "Hovoyan", time2}}

	got := people.Less(0,1)
	want := false

	got2 := people1.Less(0,1)
	want2 := false

	got3 := people3.Less(0,1)
	want3 := false
	

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}else if got2 != want2 {
		t.Errorf("got %t want %t", got2, want2)
	}else if got3 != want3 {
		t.Errorf("got %t want %t", got3, want3)
	}
}

func TestSwap (t *testing.T) {
	time1 := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)

	people  := People{{"Petros", "Hovoyan", time1}, {"Hovo", "Hovoyan", time1}}
	people.Swap(0,1)
	got := people[0].firstName
	want := "Hovo"

	if  want != got {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestNew (t *testing.T) {

	_, err := New("1 2 3\n 4 5 6")
	if err != nil {
		t.Errorf("Wrong matrix")
	}

	_, err1 := New("1 2 3\n 4 5")
	if err1.Error() != "Rows need to be the same length"{
		t.Errorf("Wrong rows length")
	}

	_, err2 := New("test")
	if !errors.Is(err2, strconv.ErrSyntax) {
		t.Errorf("Wrong sytnax")
	}

}


func TestColsRols (t *testing.T){
	s := "4 5 \n 3 2"
	m, _ := New(s)

	r := m.Cols()

	if !(r[0][1] == 3 && r[1][0] == 5){
		t.Errorf("wrong columns")
	}

	f := m.Rows()
	if !(f[0][1] == 5 && f[1][0] == 3){
		t.Errorf("wrong rows")
	}

}

func TestSet (t *testing.T){
	m := Matrix{2,2,[]int{1,2,3,4}}
	m.Set(0,0,5)

	if m.data[0*m.cols+0] != 5 {
		t.Errorf("value doesn't set right")
	}

	if m.Set(-3, 0, 5) || m.Set(0, -7, 5){
		t.Errorf("row or column is negative")
	}

	if m.Set(6, 0, 5) || m.Set(0, 9, 5){
		t.Errorf("row or column is out of range")
	}
}