package api

const (
	RpcMaxPageSize  = 1024
	RpcMaxCountSize = 1024
)

func GetRange(index, count, listLen uint32) (uint32, uint32) {
	start := index * count
	if start >= listLen {
		return listLen, listLen
	}
	end := start + count
	if end >= listLen {
		return start, listLen
	}
	return start, end
}
