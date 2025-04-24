package paging

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDoPaging(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 正常分页
	page := 2
	size := 2
	result, total := DoPaging(a, page, size)
	if diff := cmp.Diff(total, int64(len(a))); diff != "" {
		t.Errorf("total mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(result, a[2:4]); diff != "" {
		t.Errorf("result mismatch (-want +got):\n%s", diff)
	}

	// 参数异常
	page = 0
	size = 10
	result, total = DoPaging(a, page, size)
	if diff := cmp.Diff(total, int64(len(a))); diff != "" {
		t.Errorf("total mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(result, a); diff != "" {
		t.Errorf("result mismatch (-want +got):\n%s", diff)
	}
	page = 1
	size = 0
	result, total = DoPaging(a, page, size)
	if diff := cmp.Diff(total, int64(len(a))); diff != "" {
		t.Errorf("total mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(result, a); diff != "" {
		t.Errorf("result mismatch (-want +got):\n%s", diff)
	}

	// 超出范围
	page = 100
	size = 10
	result, total = DoPaging(a, page, size)
	if diff := cmp.Diff(total, int64(len(a))); diff != "" {
		t.Errorf("total mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(result, []int{}); diff != "" {
		t.Errorf("result mismatch (-want +got):\n%s", diff)
	}
	page = 2
	size = 8
	result, total = DoPaging(a, page, size)
	if diff := cmp.Diff(total, int64(len(a))); diff != "" {
		t.Errorf("total mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(result, a[8:]); diff != "" {
		t.Errorf("result mismatch (-want +got):\n%s", diff)
	}
}
