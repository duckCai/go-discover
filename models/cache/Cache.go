package cache
import (
    "sync"
)

type MemoryCache struct {
	lock  *sync.RWMutex	
	items map[interface{}]interface{}
}
func (mc *MemoryCache) set(key interface{}, value interface{}) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	mc.items[key] = value
	return nil
} 
func (mc *MemoryCache) get(key interface{}) interface{} {
	mc.lock.RLock()	
	defer mc.lock.RUnlock() 	
	if val, ok := mc.items[key]; ok {
		return val
		}
	return nil
}