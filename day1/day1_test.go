package day1

import(
    "testing"
)

var ints = [][2]int{{1122,3},{1111,4},{1234,0},{91212129,9}}
//expected values from problem

func TestIsEqual(t *testing.T) {
    for _,n := range ints{
       if ! IsEqual(n[0],n[0]) {
            t.Errorf("Expected %d,%d to return true", n[0], n[0])
       }
       if IsEqual(n[0],n[1]) {
            t.Errorf("Expected %d,%d to return false", n[0], n[1])
        }
    }
}

func TestDigits(t *testing.T) {
    for _,m := range Digits(111111111){
        if  m != 1 {
            t.Errorf("Cannot split int into digits")
       }
   }
}

func TestCode(t *testing.T) {
    for _,n := range ints{
       if m := Code(n[0]); m != n[1] {
            t.Errorf("Expected %s to return %d", n[0], n[1])
        }
    }
}
