package common

type Grid[T any] struct {
	first  *gridEntryRow[T]
	Width  int
	Height int
}

type gridEntryRow[T any] struct {
	grid  *Grid[T]
	first *gridEntry[T]
	next  *gridEntryRow[T]
}
type gridEntry[T any] struct {
	row   *gridEntryRow[T]
	next  *gridEntry[T]
	value T
}

func (grid *Grid[T]) Slice(x int, y int, width int, height int) [][]T {

	result := make([][]T, height)

	for h := range height {
		row := grid.getRowAt(y + h)
		resultRow := make([]T, width)
		result[h] = resultRow

		for w := range width {
			entry := row.getEntryAt(x + w)

			resultRow[x+w] = entry.value
		}
	}

	return result
}

func (grid *Grid[T]) getRowAt(y int) *gridEntryRow[T] {

	var lastRow *gridEntryRow[T]
	row := grid.first

	for range y + 1 {
		if row == nil {
			row = grid.createRowAfter(y, lastRow)
		}
		lastRow = row
		row = row.next
	}
	return lastRow
}

func (row *gridEntryRow[T]) getEntryAt(x int) *gridEntry[T] {

	var lastEntry *gridEntry[T]
	entry := row.first

	for range x + 1 {
		if entry == nil {
			entry = row.createEntryAfter(x, lastEntry)
		}
		lastEntry = entry
		entry = entry.next
	}
	return lastEntry
}

func (grid *Grid[T]) createRowAfter(index int, previousRow *gridEntryRow[T]) *gridEntryRow[T] {
	row := &gridEntryRow[T]{
		grid: grid,
	}
	if grid.first == nil {
		grid.first = row
	}
	if previousRow != nil {
		previousRow.next = row
	}
	if grid.Height < index+1 {
		grid.Height = index + 1
	}
	return row
}
func (row *gridEntryRow[T]) createEntryAfter(index int, previousEntry *gridEntry[T]) *gridEntry[T] {

	entry := &gridEntry[T]{
		row: row,
	}
	if row.first == nil {
		row.first = entry
	}
	if previousEntry != nil {
		previousEntry.next = entry
	}
	if row.grid.Width < index+1 {
		row.grid.Width = index + 1
	}
	return entry
}

func (grid *Grid[T]) getEntryAt(x int, y int) *gridEntry[T] {
	row := grid.getRowAt(y)
	entry := row.getEntryAt(x)
	return entry
}

func (grid *Grid[T]) SetValue(x int, y int, value T) {
	entry := grid.getEntryAt(x, y)
	entry.value = value
}
func (grid *Grid[T]) SetValues(y int, values []T) {
	row := grid.getRowAt(y)

	var lastEntry *gridEntry[T]
	entry := row.first
	for i := range len(values) {

		if entry == nil {
			entry = row.createEntryAfter(i, lastEntry)
		}
		entry.value = values[i]
		lastEntry = entry
		entry = entry.next
	}
}

func (grid *Grid[T]) GetValue(x int, y int) T {

	entry := grid.getEntryAt(x, y)
	return entry.value
}
