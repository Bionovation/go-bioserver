package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Bionovation/go-bioserver/cgo"
)

/*
内存回收机制，如果一张玻片30秒钟没有被访问，则释放该片子占用的内存
*/

type SlideGC struct {
	Map sync.Map //当前已经打开的玻片列表,记录最后一次访问的事件
}

var bioGC SlideGC // 默认实例

// 访问了某张片子
func (gc *SlideGC) Visit(path string) error {
	if gc == nil {
		return fmt.Errorf("gc is nil")
	}

	gc.Map.Store(path, time.Now())
	return nil
}

// 清理超时的片子,dropSecond 超时秒数
func (gc *SlideGC) Clear(dropSecond int) error {
	if gc == nil {
		return fmt.Errorf("gc is nil")
	}

	tNow := time.Now()
	f := func(k, v interface{}) bool {
		var ok bool
		var path string
		var t time.Time
		if path, ok = k.(string); ok != true {
			log.Println("path convert error.")
			return true
		}
		if t, ok = v.(time.Time); ok != true {
			log.Println("time convert error.")
			return true
		}
		//fmt.Println(path)
		if tNow.Sub(t).Seconds() > float64(dropSecond) {
			log.Println(path, "clear...")
			gc.Map.Delete(path)
			cgo.SlideClose(path)
		}
		return true
	}
	gc.Map.Range(f)

	return nil
}

// 清理线程,死循环
func clearRoutine(gc *SlideGC) error {
	if gc == nil {
		gc = &bioGC
	}

	for {
		gc.Clear(30)
		time.Sleep(time.Duration(11) * time.Second) // 每11秒执行一次
		log.Println("Im still alive~")
	}
}
