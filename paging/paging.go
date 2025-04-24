package paging

// DoPaging 获取分页数据及原始数据总量. a 为原始数据, page 为页码(需 > 0), size 为每页大小(需 > 0).
// 当 page, size 参数无效时, 返回原始数据和总数据量.
func DoPaging[T any](a []T, page, size int) ([]T, int64) {
	if page < 1 || size < 1 {
		return a, int64(len(a))
	}
	start := (page - 1) * size
	end := min(len(a), page*size)
	total := len(a)
	if start > len(a) {
		return []T{}, int64(total)
	}

	return a[start:end], int64(total)
}
