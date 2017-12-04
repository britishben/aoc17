package captcha

import(
    "strconv"
)

func IsEqual(a int,b int) (bool){
    return a==b
}

func Code(s string) (value int){
    a := Digits(s)
    c := 0 //what to compare to v

    for i, v := range a{

        if i == len(a)-1 {
            c=a[0]
        }else{
            c=a[i+1]
        }

        if IsEqual(v, c) {
            value += v
        }
    }
    return value
}

func Digits(s string) (r []int){
    //s := strconv.Itoa(i)
    for _,v := range s {
        n,_ := strconv.Atoi(string(v))
        r = append (r, n)
    }
    return r
}