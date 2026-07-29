type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (this *LinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value
}

func (this *LinkedList) InsertHead(val int) {
	newNode := &Node{Value: val, Next: this.head}
	this.head = newNode

	if this.tail == nil {
		this.tail = newNode
	}

	this.size++
}

func (this *LinkedList) InsertTail(val int) {
	newNode := &Node{Value: val, Next: nil}

	if this.tail == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		this.tail.Next = newNode
		this.tail = newNode
	}

	this.size++
}

func (this *LinkedList) Remove(index int) bool {
	if index < 0 || index >= this.size {
		return false
	}

	if index == 0 {
		this.head = this.head.Next
		if this.head == nil {
			this.tail = nil
		}
		this.size--
		return true
	}

	current := this.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next

	if index == this.size-1 {
		this.tail = current
	}

	this.size--
	return true
}

func (this *LinkedList) GetValues() []int {
	arr := make([]int, 0, this.size)

	current := this.head
	for current != nil {
		arr = append(arr, current.Value)
		current = current.Next
	}

	return arr
}

