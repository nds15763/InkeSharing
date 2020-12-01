package main

type (
	Char = rune
	Trie map[Char]Trie
)

//添加
func (t Trie) AddString(s string) {
	for _, ch := range []Char(s) {
		if t[ch] != nil {
			t = t[ch]
		} else {
			prevT := t
			t = make(Trie)
			prevT[ch] = t
		}
	}
}

//查找
func (t Trie) ExistPrefix(prefix string) bool {
	for _, ch := range []Char(prefix) {
		if t[ch] == nil {
			return false
		}
		t = t[ch]
	}
	return true
}

//推荐
func (t Trie) Suggest(prefix string) []string {
	prefixBytes := []Char(prefix)
	for _, ch := range prefixBytes {
		t = t[ch]
	}
	return dfs(t, prefixBytes, nil)
}

func dfs(t Trie, prefix []Char, suggestions []string) []string {
	for ch, newT := range t {
		prefix := append(prefix, ch)
		if len(newT) == 0 {
			suggestions = append(suggestions, string(prefix))
		} else {
			suggestions = dfs(newT, prefix, suggestions)
		}
	}
	return suggestions
}

// func main() {
// 	dict := make(Trie)
// 	dict.AddString("go")
// 	dict.AddString("golang")
// 	dict.AddString("by")

// 	exsist := dict.ExistPrefix("golang")
// 	fmt.Println(exsist)

// 	relist := dict.Suggest("go")
// 	fmt.Println(relist)
// }
