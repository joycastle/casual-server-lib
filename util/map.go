package util

func GetMapKeysString(m map[interface{}]interface{}) []string {
	ks := make([]string, 0, len(m))
	for k, _ := range m {
		ks = append(ks, k.(string))
	}
	return ks
}
