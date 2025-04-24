package paging

// DoPaging 获取分页数据及原始数据总量. a 为原始数据, page 为页码(需 > 0), size 为每页大小(需 > 0).
// 当 page, size 参数无效时, 返回原始数据和总数据量.
func DoPaging[T any](a []T, page, size int) ([]T, int) {
	total := len(a)
	if page < 1 || size < 1 {
		return a, total
	}
	start := (page - 1) * size // 数组下标.
	end := min(total, page*size)
	if start >= total {
		return []T{}, total
	}

	return a[start:end], total
}
