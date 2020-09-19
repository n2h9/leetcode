package solution

func equationsPossible(equations []string) bool {
	// 26 is number of letters in abc
	uf := newUF(26)
	for _, e := range equations {
		if e[1] == '=' {
			uf.union(e[0]-'a', e[3]-'a')
		}
	}
	for _, e := range equations {
		if e[1] == '!' && uf.find(e[0]-'a') == uf.find(e[3]-'a') {
			return false
		}
	}
	return true
}

type uf struct {
	storage []byte
}

func newUF(len int) *uf {
	storage := make([]byte, len)
	for i := range storage {
		storage[i] = byte(i)
	}
	return &uf{
		storage: storage,
	}
}

func (uf *uf) union(q, p byte) {
	qr := uf.find(q)
	pr := uf.find(p)
	if qr != pr {
		uf.storage[qr] = pr
	}
}

func (uf *uf) find(p byte) byte {
	children := make([]byte, 0)
	for ; uf.storage[p] != p; p = uf.storage[p] {
		children = append(children, p)
	}
	if len(children) > 1 {
		for _, i := range children {
			uf.storage[i] = p
		}
	}
	return p
}
