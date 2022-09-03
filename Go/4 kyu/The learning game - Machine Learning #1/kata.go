package kata

type Machine struct {
	last []int
	list map[int]int
}

func NewMachine() Machine {
	var m Machine
	m.last = []int{}
	m.list = map[int]int{}
	return m
}

func (m *Machine) Command(cmd int, num int) int {
	if _, ok := m.list[cmd]; !ok {
		m.list[cmd] = 0
	}
	m.last = append(m.last, cmd)
	return Actions()[m.list[cmd]](num)
}

func (m *Machine) Response(res bool) {
	if !res {
		l := len(m.last) - 1
		m.list[m.last[l]]++
		m.list[m.last[l]] = m.list[m.last[l]] % 5
	}
}
