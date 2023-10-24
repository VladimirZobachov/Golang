package stack

type Stack interface {
	Put(item any) // Добавляет элемент наверх стека
	Pop() any     // Удаляет элемент сверху стека и возвращает его
	Len() int     // Возвращает количество элементов в стеке
}

// Реализовать стэк согласно интерфейса
type Simple struct {
}
