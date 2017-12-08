package day2

import(
    "testing"
)

const sheet = [][]int{
    {5,1,9,5},
    {7,5,3},
    {2,4,6,8}
}

const answer int = 18

//expected values from problem

func TestMinMax(t *testing.T) {
    if m,n := MinMax(sheet[1); m != 1 || n != 9{
        t.Errorf("Cannot get MinMax: %s, %s", m,n)
    }
}
