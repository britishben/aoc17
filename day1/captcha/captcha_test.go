package captcha

import(
    "testing"
    "strconv"
)

var ones = [][2]int{{1122,3},{1111,4},{1234,0},{91212129,9}}
var twos = [][2]int{{1212,6},{1221,0},{123425,4},{123123,12},{12131415,4}}

//expected values from problem

func TestIsEqual(t *testing.T) {
    for _,n := range ones{
       if ! IsEqual(n[0],n[0]) {
            t.Errorf("Expected %d,%d to return true", n[0], n[0])
       }
       if IsEqual(n[0],n[1]) {
            t.Errorf("Expected %d,%d to return false", n[0], n[1])
        }
    }
}

func TestDigits(t *testing.T) {
    for _,m := range Digits("111111111"){
        if  m != 1 {
            t.Errorf("Cannot split int into digits")
       }
   }
}

func TestCodes(t *testing.T) {
    for _,n := range ones{
        if m := CodeOne(strconv.Itoa(n[0])); m != n[1] {
            t.Errorf("One: Expected %s to return %d", n[0], n[1])
        }
    }
    for _,n := range twos{
        if m := CodeTwo(strconv.Itoa(n[0])); m != n[1] {
            t.Errorf("Two: Expected %s to return %d", n[0], n[1])
        }
    }
}
