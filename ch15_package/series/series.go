package series

import "errors"

var LessThanTwoError = errors.New("n 必须大于2！")

func GetFib(n int) ([]int, error){

	if n < 2 {
		return nil, LessThanTwoError
	}

	fbList := []int{1, 1}
	a := 1;
	b := 1;

	for i := 2; i < n; i++ {
		temp := a
		a = b;
		b = temp + a
		fbList = append(fbList, b)
	}
	return fbList, nil
}