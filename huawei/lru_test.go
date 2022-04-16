package huawei

import "testing"

func TestLRU(t *testing.T) {
	s := Constructor(2)
	s.Set(1, 1)
	s.Set(2, 2)
	if n := s.Get(1); n != 1 {
		t.Fatalf("excepted %d, but got %d", 1, n)
	}
	s.Set(3, 3)
	if n := s.Get(2); n != -1 {
		t.Fatalf("excepted %d, but got %d", -1, n)
	}
	s.Set(4, 4)
	if n := s.Get(1); n != -1 {
		t.Fatalf("excepted %d, but got %d", -1, n)
	}
	if n := s.Get(3); n != 3 {
		t.Fatalf("excepted %d, but got %d", 3, n)
	}
	if n := s.Get(4); n != 4 {
		t.Fatalf("excepted %d, but got %d", 4, n)
	}

}
