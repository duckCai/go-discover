package cache
import (
    "sync"
)

type MemoryCache struct {
    lock  *sync.RWMutex	
    items map[string]interface{}
}

// func(mc *MemoryCache) GetInstance()*MemoryCache{
//     mc.lock.Lock()
//     defer mc.lock.Unlock()
  
//     if MemoryCache==nil{
//         MemoryCache=&MemoryCache{}
//     }
//     return MemoryCache
// }

func (mc *MemoryCache) Set(key string, value interface{}) error {
    mc.lock.Lock()
    defer mc.lock.Unlock()
    mc.items[key] = value
    return nil
} 
func (mc *MemoryCache) Get(key string) interface{} {
    mc.lock.RLock()	
    defer mc.lock.RUnlock() 	
    if val, ok := mc.items[key]; ok {
        return val
    }
    return nil
}