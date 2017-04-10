/*
	仿照:https://github.com/orcaman/concurrent-map的思路
*/
package concurrent_map

import "sync"

var SHARD_COUNT = 32

//ConcurrentMap 线程安全的map  存储数据类型为map[string]interface{}
//不能直接调用 需要通过New函数初始化
//为了避免锁的瓶颈,划分为一些小的map片,通过fnv32哈希去映射
type ConcurrentMap []*ConcurrentMapShared

//ConcurrentMapShared 一个线程安全的类型为map[string]interface{}的map
type ConcurrentMapShared struct {
	items map[string]interface{}
	//读写锁
	sync.RWMutex
}

//New 初始化一个线程安全的ConcurrentMap
//初始化了SHARD_COUNT个小的map片,为了避免锁的瓶颈
func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}

//GetShard 通过对key值进行hash散列,对散列值进行取余操作,从而获取一个小的分片
func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

//Set 给map设置单个键值对
func (m ConcurrentMap) Set(key string, value interface{}) {
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

//SetIfAbsent 如果key的值为空,则进行设置,否则不设置 返回是否设置
func (m ConcurrentMap) SetIfAbsent(key string, value interface{}) bool {
	shard := m.GetShard(key)
	shard.Lock()
	_, ok := shard.items[key]
	if !ok {
		shard.items[key] = value
	}
	shard.Unlock()
	return !ok
}

//MultiSet 给map设置多个键值对
func (m ConcurrentMap) MultiSet(data map[string]interface{}) {
	for key, value := range data {
		shard := m.GetShard(key)
		shard.Lock()
		shard.items[key] = value
		shard.Unlock()
	}
}

//Get 获取指定键的值  返回值以及是否存在
func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	shard := m.GetShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

//Count 获取map中存储的键值对数量
func (m ConcurrentMap) Count() int {
	count := 0
	for i := 0; i < SHARD_COUNT; i++ {
		shard := m[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

//Delete 删除map中的某个key以及对应的value
func (m ConcurrentMap) Delete(key string) {
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

//Pop 删除map中的某个key以及对应的value,并且返回value,以及是否存在
func (m ConcurrentMap) Pop(key string) (interface{}, bool) {
	shard := m.GetShard(key)
	shard.Lock()
	value, exist := shard.items[key]
	delete(shard.items, key)
	shard.Unlock()
	return value, exist
}

//IsEmpty 判断map是否为空
func (m *ConcurrentMap) IsEmpty() bool {
	return m.Count() == 0
}

//Tuple 包含键值对的元组 用于在下面的函数中通过channel包裹键值对
type Tuple struct {
	Key string
	Val interface{}
}

//Iter 返回一个有缓冲的迭代器可以用于range循环中使用,返回的是只读channel
func (m ConcurrentMap) IterBuffered() <-chan Tuple {
	ch := make(chan Tuple, m.Count())
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(SHARD_COUNT)
		for _, shard := range m {
			go func(shard *ConcurrentMapShared) {
				shard.RLock()
				for key, val := range shard.items {
					ch <- Tuple{key, val}
				}
				shard.RUnlock()
				wg.Done()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()
	return ch
}

//Items 获取map中存储的全部键值对
func (m ConcurrentMap) Items() map[string]interface{} {
	data := make(map[string]interface{})
	for item := range m.IterBuffered() {
		data[item.Key] = item.Val
	}
	return data
}

//fnv32 32位的FNV-1算法的实现(hash算法,对key进行散列)
//由于hash结果是按位异或和乘积的，如果任何一步出现0，则结果可能会造成冲突；
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
