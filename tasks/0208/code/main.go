package main

import "fmt"

type Trie struct {
	char  string
	isEnd bool
	next  []*Trie
}

func Constructor() Trie {
	return Trie{
		char: "",
		next: []*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	wd := ""
	curr := this

	for _, ch := range word {
		wd += string(ch)

		if len(curr.next) == 0 {
			curr.next = append(curr.next, &Trie{
				char:  wd,
				isEnd: wd == word,
				next:  []*Trie{},
			})
			curr = curr.next[0]
			continue
		}

		fd := false
		for _, n := range curr.next {
			if n.char == wd {
				curr, fd = n, true
				if wd == word {
					curr.isEnd = true
				}
			}
		}

		if !fd {
			newTrie := &Trie{
				char:  wd,
				isEnd: wd == word,
				next:  []*Trie{},
			}
			curr.next = append(curr.next, newTrie)
			curr = newTrie
		}

	}
}

func (this *Trie) Find(word string) (bool, bool) {
	wd := ""

	curr := this

	for _, ch := range word {
		fd := false
		wd += string(ch)
		for _, n := range curr.next {
			if n.char == wd {
				fd, curr = true, n
			}
		}

		if !fd {
			return false, false
		}

	}

	return true, curr.isEnd

}

func (this *Trie) Search(word string) bool {
	found, end := this.Find(word)
	return found == true && end == true
}

func (this *Trie) StartsWith(prefix string) bool {
	found, _ := this.Find(prefix)
	return found == true
}

func main() {
	t := Constructor()
	t.Insert("a")
	t.Search("a")
	fmt.Println(t.StartsWith("a"))

}
