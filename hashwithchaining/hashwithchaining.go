package hashwithchaining

//Hash for hashing
func Hash(x int) int {
	return x % 10
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

type HashTable struct {
	Store [10]LinkedList
}

//Set the hashtable
func (hashtable *HashTable) Set(key string, val int) {

	str := key
	sum := 0
	for _, elem := range str {
		sum += int(elem)
	}

	hashKey := Hash(sum)
	linkedValues := hashtable.Store[hashKey]
	newNode := NewNode(key, val)

	linkedValues.PushFront(newNode)
	hashtable.Store[hashKey] = linkedValues
}

//Get the hashtable
func (hashtable *HashTable) Get(key string) int {

	value := 0
	str := key
	sum := 0
	for _, elem := range str {
		sum += int(elem)
	}

	hashKey := Hash(sum)

	linkedValues := hashtable.Store[hashKey]
	temp := linkedValues.Head

	if temp == nil {
		return value
	}

	for temp != nil {
		if temp.key == key {
			value = temp.val
			break
		}
		temp = temp.Next
	}

	return value
}
