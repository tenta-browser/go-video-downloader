package runtime

import (
	"fmt"
	"testing"
)

func TestDictBasic(t *testing.T) {
	ht := newHashTable(3)
	if ht.get(StrLiteral("one")) != nil {
		t.Fatal("Getting missing element failed")
	}
	if ht.put(StrLiteral("one"), IntLiteral(0), false) != nil {
		t.Fatal("Putting nonexistent item failed")
	}
	if ht.get(StrLiteral("one")) != IntLiteral(0) {
		t.Fatal("Getting existing item failed")
	}
	if ht.put(StrLiteral("one"), IntLiteral(1), false) != IntLiteral(0) ||
		ht.get(StrLiteral("one")) != IntLiteral(0) {
		t.Fatal("Putting existing item without overwrite failed")
	}
	if ht.put(StrLiteral("one"), IntLiteral(1), true) != IntLiteral(0) ||
		ht.get(StrLiteral("one")) != IntLiteral(1) {
		t.Fatal("Putting existing item with overwrite failed")
	}
	if ht.len != 1 {
		t.Fatal("Invalid table length")
	}
	ht.put(StrLiteral("two"), IntLiteral(2), false)
	ht.put(StrLiteral("three"), IntLiteral(3), false)
	if len(ht.table) != 6 {
		t.Fatal("Table was not grown")
	}
	if ht.get(StrLiteral("one")) != IntLiteral(1) ||
		ht.get(StrLiteral("two")) != IntLiteral(2) ||
		ht.get(StrLiteral("three")) != IntLiteral(3) {
		t.Fatal("Item corruption after growth")
	}
	if ht.remove(StrLiteral("four")) != nil {
		t.Fatal("Removing nonexistent item failed")
	}
	if ht.remove(StrLiteral("two")) != IntLiteral(2) {
		t.Fatal("Removing existing item failed")
	}
	exp := map[string]int{
		"one":   1,
		"three": 3,
	}
	it := ht.iterator()
	for {
		entry := it()
		if entry == nil {
			break
		}
		key := ToStr(entry.Key).Value()
		gotValue := ToInt(entry.Value).Value()
		expValue, found := exp[key]
		if !found {
			t.Fatal("Iteration returned invalid key")
		}
		if expValue != gotValue {
			t.Fatal("Iteration returned invalid value")
		}
		delete(exp, key)
	}
	if len(exp) > 0 {
		t.Fatal("Iteration did not return all values")
	}
	if it() != nil || it() != nil || it() != nil {
		t.Fatal("Iterator is not stably finished")
	}
}

func TestDictOrdering(t *testing.T) {
	ht := newHashTable(3)
	mkKey := func(i int) Object {
		return StrLiteral(fmt.Sprintf("key%d", i))
	}
	for i := 0; i < 100; i++ {
		ht.put(mkKey(i), IntLiteral(i), false)
	}
	for i := 100 - 1; i >= 0; i-- {
		ht.put(mkKey(i), IntLiteral(i+10), true)
	}
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			ht.put(mkKey(i), IntLiteral(i+100), true)
		}
	}
	it := ht.iterator()
	for e, i := it(), 0; e != nil; e, i = it(), i+1 {
		if e.Key != mkKey(i) {
			t.Fatal("Iteration returned wrong key")
		}
		var expVal int
		if i%2 == 0 {
			expVal = i + 100
		} else {
			expVal = i + 10
		}
		if e.Value != IntLiteral(expVal) {
			t.Fatal("Iteration returned wrong value")
		}
	}
}
