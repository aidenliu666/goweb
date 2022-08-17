package main

import (
	"context"
	"fmt"
	"goweb/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Minute))
	defer cancel()

	// mu := sync.Mutex{}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		fmt.Println("aaaaaa")
		time.Sleep(2 * time.Second)
		fmt.Println("aaaaaa")
		c.SetStatus(200).Json("ok")

		finish <- struct{}{}
	}()
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.SetStatus(500).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(500).Json("time out")
		c.SetHasTimeout()
	}
	fmt.Println("aaaaaa")
	return nil
}

func SubjectListController(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Minute))
	defer cancel()

	// mu := sync.Mutex{}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		fmt.Println("aaaaaa")
		time.Sleep(2 * time.Second)
		fmt.Println("aaaaaa")
		c.SetStatus(200).Json("ok")

		finish <- struct{}{}
	}()
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.SetStatus(500).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(500).Json("time out")
		c.SetHasTimeout()
	}
	fmt.Println("aaaaaa")
	return nil
}

// func Foo(request *http.Request, response http.ResponseWriter) {
// 	obj := map[string]interface{}{
// 		"errno":  50001,
// 		"errmsg": "inner error",
// 		"data":   nil,
// 	}

// 	response.Header().Set("Content-Type", "application/json")

// 	foo := request.PostFormValue("foo")
// 	if foo == "" {
// 		foo = "10"
// 	}
// 	fooInt, err := strconv.Atoi(foo)
// 	if err != nil {
// 		response.WriteHeader(500)
// 		return
// 	}
// 	obj["data"] = fooInt
// 	byt, err := json.Marshal(obj)
// 	if err != nil {
// 		response.WriteHeader(500)
// 		return
// 	}
// 	response.WriteHeader(200)
// 	response.Write(byt)
// 	return
// }

// func Foo2(ctx *framework.Context) error {
// 	obj := map[string]interface{}{
// 		"errno":  50001,
// 		"errmsg": "inner error",
// 		"data":   nil,
// 	}

// 	fooInt := ctx.FormInt("foo", 10)
// 	obj["data"] = fooInt
// 	return ctx.Json(http.StatusOK, obj)
// }

// func Foo3(ctx *framework.Context) error {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	return rdb.Set(ctx, "key", "value", 0).Err()
// }
