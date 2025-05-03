package main

type Trie struct {
	c      [26]bool
	childs [26]*Trie
	isEnd  bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	p := this
	for i := range word {
		p.c[word[i]-'a'] = true
		if p.childs[word[i]-'a'] == nil {
			p.childs[word[i]-'a'] = &Trie{}
		}
		p = p.childs[word[i]-'a']
	}
	p.isEnd = true
}

func (this *Trie) Search(word string) bool {
	for i := range word {
		if this.c[word[i]-'a'] == false {
			return false
		}
		this = this.childs[word[i]-'a']
	}
	if this.isEnd == true {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := range prefix {
		if this.c[prefix[i]-'a'] == false {
			return false
		}
		this = this.childs[prefix[i]-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
