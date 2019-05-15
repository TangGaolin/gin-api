package utils

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"unsafe"
)

func GenRid() string {
	return fmt.Sprintf("%d_%s", time.Now().Unix(), uuid.New().String())
}

func ByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}