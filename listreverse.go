package piscine

func ListReverse(l *List) {
	for i := l.Head; i != nil; i = i.Next {
		for j := i; j != nil; j = j.Next {
			i.Data, j.Data = j.Data, i.Data
		}
	}
}
