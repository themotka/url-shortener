package _map

type MapUrl struct {
	Table map[string]string
}

func (m *MapUrl) Write(key string, url string) error {
	m.Table[key] = url
	return nil
}

func (m *MapUrl) IsPresentedByKey(key string) (bool, string) {
	if val, ok := m.Table[key]; ok {
		return true, val
	}
	return false, ""
}

func (m *MapUrl) IsPresentedByURL(url string) (bool, string) {
	for key, val := range m.Table {
		if val == url {
			return true, key
		}
	}
	return false, ""
}
