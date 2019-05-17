package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/rs/xid"
	"time"
	"unsafe"
)

func GenRid() string {
	return fmt.Sprintf("%s_%d", xid.New().String(), time.Now().Unix())
}

func Md5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Bytes2string(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
