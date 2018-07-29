package hashmap

//Hash for hashing
func Hash(x int) int {
	return x % 10
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

type HashTable struct {
	Store [10]int
}

//Todo for collision
func (hashtable *HashTable) Set(key string, val int) {

	str := key
	sum := 0
	for _, elem := range str {
		sum += int(elem)
	}

	hashKey := Hash(sum)
	hashtable.Store[hashKey] = val
}

func (hashtable *HashTable) Get(key string) int {

	str := key
	sum := 0
	for _, elem := range str {
		sum += int(elem)
	}
	hashKey := Hash(sum)

	return hashtable.Store[hashKey]
}
