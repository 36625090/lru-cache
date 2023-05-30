/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lru_cache

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestLruCache_SAdd(t *testing.T) {

	lru := NewLRUCache[int](4, time.Second*3)
	lru.SAdd("namekey", 111)
	lru.SAdd("namekey", 222)
	lru.SAdd("namekey", 333)
	for _, i := range lru.ARCCache.Keys() {
		val, ok := lru.Get(i)
		t.Log(i, val, ok)
	}
	//time.Sleep(time.Second * 5)
	for _, i := range lru.ARCCache.Keys() {
		val, ok := lru.Get(i)
		t.Log(i, val, ok)
	}

	//time.Sleep(time.Second * 2)
	t.Log(lru.Get("avl"))
	var ix uint64 = 0xffffffffffffffff
	t.Log(ix+1, ix, unsafe.Sizeof(ix), len(fmt.Sprintf("%d", ix)))
}
