package main

type (
	FChar = rune
	State map[FChar]State
)

const EOS = FChar(-1)

func NewFinalState() State {
	return State{
		EOS: nil,
	}
}

func isFinalState(st State) bool {
	_, found := st[EOS]
	return found
}

func Match(s string, st State) bool {
	for _, ch := range []FChar(s) {
		st = st[ch]
	}
	return isFinalState(st)
}

func main() {

	St1 := NewFinalState()
	St2 := make(State)

	St1['a'] = St2
	St1['b'] = St2
	St1['c'] = St1
	St2['c'] = St1

}
