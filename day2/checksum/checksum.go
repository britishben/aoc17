package day2

import(
)

func MinInt(a int,b int) (int){
    if a<b {
        return a
    }
    return b
}

func MaxInt(a int,b int) (int){
    if a>b {
        return a
    }
    return b
}

func MinMax(a int[]) (min int, max int){
    min,max = a[0]
    for _,v := range a {
        if v<min{
            min=v
        }
        if v>max{
            max=v
        }
    }
}

func RowSum(r int[]) int{
    a,b := MinMax(r)
    return a+b
}
