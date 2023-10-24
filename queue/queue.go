package queue

type Queue interface {
	Put(item any) // Добавляет элемент в конец очереди
	Get() any     // Возвращает элемент из начала очереди и удаляет его
	Len() int     // Возвращает количество элементов в очереди
}

// Реализовать очередь согласно интерфейса
type Simple struct {
}
