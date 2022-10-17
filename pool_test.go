package workerpool

import (
	"testing"
)

func TestNewPool(t *testing.T) {
	p := New(5)
	if p.capacity != 5 {
		t.Errorf("capacity want 5, actual %d\n", p.capacity)
	}
	if p.preAlloc {
		t.Errorf("preAlloc want false, actual %t\n", p.preAlloc)
	}
	if !p.block {
		t.Errorf("block want true, actual %t\n", p.block)
	}
	if len(p.active) != 0 {
		t.Errorf("active len want 0, actual %d\n", len(p.active))
	}

	p.Free()
	if len(p.active) != 0 {
		t.Errorf("after free, active len want 0, actual %d\n", len(p.active))
	}

	p = New(5, WithBlock(false), WithPreAllocWorkers(true))
	if !p.preAlloc {
		t.Errorf("preAlloc want true, actual %t\n", p.preAlloc)
	}
	if p.block {
		t.Errorf("block want false, actual %t\n", p.block)
	}
	if len(p.active) != 5 {
		t.Errorf("active len want 5, actual %d\n", len(p.active))
	}

	p.Free()
	if len(p.active) != 0 {
		t.Errorf("after free, active len want 0, actual %d\n", len(p.active))
	}

	p = New(-1)
	if p.capacity != defaultCapacity {
		t.Errorf("capacity want %d, actual %d\n", defaultCapacity, p.capacity)
	}
	p.Free()
}

func TestSchedule(t *testing.T) {
	p := New(5)
	if p.capacity != 5 {
		t.Errorf("capacity want 5, actual %d\n", p.capacity)
	}
	err := p.Schedule(func() {})
	if err != nil {
		t.Errorf("want nil, actual %s\n", err)
	}

	p.Free()
	err = p.Schedule(func() {})
	if err == nil {
		t.Errorf("want non nil, actual nil\n")
	}
}