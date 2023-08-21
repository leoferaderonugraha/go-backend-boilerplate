package queue

type Queue[T any, U any] interface {
    Pop() (T, error)
    Push(T) error
    PeekFront() (T, error)
    PeekBack() (T, error)
    Items() U
    Clear() error
    Size() int64
}
