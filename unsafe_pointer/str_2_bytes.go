package unsafe_pointer

import "unsafe"

func str2bytes(s string) []byte {
	sp := (*[2]uintptr)(unsafe.Pointer(&s))
	bp := [3]uintptr{sp[0], sp[1], sp[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

func bytes2str(b []byte) string {
	bp := (*[3]uintptr)(unsafe.Pointer(&b))
	sp := [2]uintptr{bp[0], bp[1]}
	return *(*string)(unsafe.Pointer(&sp))
}
