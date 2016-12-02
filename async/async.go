package async

import (
	"errors"

	"github.com/freshcn/async"
)

// AddFunc go func
type AddFunc struct {
	Logo    string
	Handler interface{}
	Params  []interface{}
}

// MultiAddFunc []go func
type MultiAddFunc []AddFunc

// GoAsyncRequest 异步调用
func GoAsyncRequest(addFunc []AddFunc, length int) (map[string][]interface{}, error) {
	async := async.New()

	for _, v := range addFunc {
		async.Add(v.Logo, v.Handler, v.Params...)
	}

	var res map[string][]interface{}
	var err error

	if chans, ok := async.Run(); ok {
		res = <-chans
		if len(res) != length {
			err = errors.New("async not execution all task")
		}
	}
	async.Clean()

	return res, err
}

//e.g
/*
func request1(str string) int {
	fmt.Println("1" + str)
	return 1
}

func request2() interface{} {
	fmt.Println("2")
	return nil
}

func main() {
	var addFunc MultiAddFunc
	addFunc = append(addFunc, AddFunc{Logo: "a", Handler: request1, Params: []interface{}{"a"}})
	addFunc = append(addFunc, AddFunc{Logo: "b", Handler: request2})

	res, err := GoAsyncRequest(addFunc, 2)
	fmt.Println(res, err)
	fmt.Println(res["a"][0])
	fmt.Println(res["b"][0])
}
*/
