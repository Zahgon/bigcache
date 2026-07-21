package queue

const (
	minimumHeaderSize = 17

	leftMarginIndex = 1
)

var (
	errEmptyQueue       = &queueError{"Empty queue"}
	errInvalidIndex     = &queueError{"Index must be greater than zero. Invalid index."}
	errIndexOutOfBounds = &queueError{"Index out of range"}
	errFullQueue        = &queueError{"Full queue. Maximum size limit reached."}
)

type BytesQueue struct {
	full         bool
	array        []byte
	capacity     int
	maxCapacity  int
	head         int
	tail         int
	count        int
	rightMargin  int
	headerBuffer []byte
	verbose      bool
}

type queueError struct {
	message string
}

func getNeededSize(length int) int { _ = "STUB: not implemented"; return 0 }

func NewBytesQueue(capacity int, maxCapacity int, verbose bool) *BytesQueue {
	_ = "STUB: not implemented"
	return nil
}

func (q *BytesQueue) Reset() { _ = "STUB: not implemented"; return }

func (q *BytesQueue) Push(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (q *BytesQueue) allocateAdditionalMemory(minimum int) { _ = "STUB: not implemented"; return }

func (q *BytesQueue) push(data []byte, len int) { _ = "STUB: not implemented"; return }

func (q *BytesQueue) copy(data []byte, len int) { _ = "STUB: not implemented"; return }

func (q *BytesQueue) Pop() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *BytesQueue) Peek() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *BytesQueue) Get(index int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *BytesQueue) CheckGet(index int) error { _ = "STUB: not implemented"; return nil }

func (q *BytesQueue) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (q *BytesQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (e *queueError) Error() string { _ = "STUB: not implemented"; return "" }

func (q *BytesQueue) peekCheckErr(index int) error { _ = "STUB: not implemented"; return nil }

func (q *BytesQueue) peek(index int) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (q *BytesQueue) canInsertAfterTail(need int) bool { _ = "STUB: not implemented"; return false }

func (q *BytesQueue) canInsertBeforeHead(need int) bool { _ = "STUB: not implemented"; return false }
