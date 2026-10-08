package keyvalues

import "sync"

type KeyValues struct {
	mu   sync.RWMutex
	data map[string]string
}

func CreateKeyValue() *KeyValues {
	return &KeyValues{
		data: map[string]string{},
	}
}

func (kv *KeyValues) Set(key string, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[key] = value
}

func (kv *KeyValues) Get(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	value, exists := kv.data[key]
	return value, exists
}

func (kv *KeyValues) Del(key string) bool {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	_, exists := kv.data[key]
	if exists {
		delete(kv.data, key)
	}
	return exists
}

func (kv *KeyValues) Exists(key string) bool {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	_, exists := kv.data[key]
	return exists
}
