type DynamicArray struct {
    array []int
    capacity int 
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        array: make([]int, 0, capacity), 
        capacity: capacity, 
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.array[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if len(da.array) == da.capacity{ 
        da.resize()
    }
    da.array = append(da.array, n)
}

func (da *DynamicArray) Popback() int {
	lastIdx := len(da.array) - 1
	val := da.array[lastIdx]

	da.array = da.array[:lastIdx]

	return val
}

func (da *DynamicArray) resize() {
    newCapacity := len(da.array) * 2 
    newArray := make([]int, len(da.array), newCapacity)
    copy(newArray, da.array) 

    da.array = newArray
    da.capacity = newCapacity 
}

func (da *DynamicArray) GetSize() int {
    return len(da.array)
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity 
}
