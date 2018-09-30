package cache


type Cache interface {

	Get(key string) interface{}

	Put(key string, value interface{}, timeout time.Duration) error
    //定时检测，处理过期的key
	TimerCheck() error
}