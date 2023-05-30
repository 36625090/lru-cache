/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lru_cache

import (
	"github.com/google/uuid"
	lru "github.com/hnlq715/golang-lru"
	"sync"
	"time"
)

type LRUCache[T any] struct {
	*lru.ARCCache
	sync.Mutex
	keys map[string]map[string]int64
}

func NewLRUCache[T any](size int, expire time.Duration) *LRUCache[T] {
	l, _ := lru.NewARCWithExpire(size, expire)
	return &LRUCache[T]{ARCCache: l, keys: map[string]map[string]int64{}}
}

func (m *LRUCache[T]) SAdd(key string, val T) {
	m.Lock()
	defer m.Unlock()
	eKey := uuid.New().String()
	if m.keys[key] == nil {
		m.keys[key] = map[string]int64{}
	}
	m.keys[key][eKey] = 0
	m.ARCCache.Add(eKey, val)
}

func (m *LRUCache[T]) SMembers(key string) []T {
	m.Lock()
	defer m.Unlock()

	var members []T
	if m.keys[key] == nil {
		return members
	}
	keys := m.keys[key]
	for eKey, _ := range keys {
		val, ok := m.ARCCache.Get(eKey)
		if !ok {
			delete(m.keys[key], eKey)
			continue
		}
		members = append(members, val.(T))
	}
	return members
}
func (m *LRUCache[T]) SClear(key string) {
	m.Lock()
	defer m.Unlock()

	if m.keys[key] == nil {
		return
	}
	for _, k := range m.keys[key] {
		m.ARCCache.Remove(k)
	}
}

func (m *LRUCache[T]) SLen(key string) int {
	m.Lock()
	defer m.Unlock()
	if m.keys[key] == nil || len(m.keys[key]) == 0 {
		return 0
	}

	total := 0
	keys := m.keys[key]
	for eKey, _ := range keys {
		_, ok := m.ARCCache.Get(eKey)
		if ok {
			total++
		}
	}
	return total
}

func (m *LRUCache[T]) Get(key interface{}) (T, bool) {
	m.Lock()
	defer m.Unlock()
	value, ok := m.ARCCache.Get(key)
	if ok {
		return value.(T), true
	}
	var v T
	return v, false
}
