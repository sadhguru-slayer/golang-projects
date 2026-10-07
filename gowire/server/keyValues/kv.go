package keyvalues

type KeyValues struct {
	data map[string]string
}

func CreateKeyValue() *KeyValues {
	return &KeyValues{
		data: map[string]string{},
	}
}

func (kv *KeyValues) Set(key string, value string) {
	kv.data[key] = value
}

func (kv *KeyValues) Get(key string) (string, bool) {
	value, exists := kv.data[key]
	return value, exists
}

func (s *KeyValues) Del(key string) bool {
	_, exists := s.data[key]
	if exists {
		delete(s.data, key)
	}
	return exists
}

func (s *KeyValues) Exists(key string) bool {
	_, exists := s.data[key]
	return exists
}
