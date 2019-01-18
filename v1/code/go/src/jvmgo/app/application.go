package main

import (
	"fmt"
	//"github.com/cm-heclouds/GoSDK/oneNet"
	"bytes"
	"encoding/binary"
)

func main() {
	// on := oneNet.NewOneNet("2MGfqkx8yTuLA0n9lFBMZLNgGQwA")
	// datapoints := make(map[string]interface{})
	// datapoints["2014-09-01 15:11:01"] = 15
	// datapoints["2014-09-01 15:16:01"] = 20
	// ret, s := on.DatapointAdd("66114", "datastream_id1", datapoints)
	// if ret == true {
	// 	fmt.Println(ret)
	// 	fmt.Println(*s)
	// } else {
	// 	fmt.Println(ret)
	// 	if s != nil {
	// 		fmt.Println(*s)
	// 	} else {
	// 		fmt.Println(on.GetErrorNo())
	// 		fmt.Println(on.GetError())
	// 	}
	// }

	var i1 int64 = 511 // [00000000 00000000 ... 00000001 11111111] = [0 0 0 0 0 0 1 255]

	s1 := make([]byte, 0)
	buf := bytes.NewBuffer(s1)

	// 数字转 []byte, 网络字节序为大端字节序
	binary.Write(buf, binary.BigEndian, i1)
	fmt.Println(buf.Bytes())

	// 数字转 []byte, 小端字节序
	buf.Reset()
	binary.Write(buf, binary.LittleEndian, i1)
	fmt.Println(buf.Bytes())

	// []byte 转 数字
	s2 := []byte{6: 1, 7: 255} // [0 0 0 0 0 0 1 255]
	buf = bytes.NewBuffer(s2)
	var i2 int64
	binary.Read(buf, binary.BigEndian, &i2)
	fmt.Println(i2) // 511

	s3 := []byte{255, 1, 7: 0} // [255 1 0 0 0 0 0 0]
	buf = bytes.NewBuffer(s3)
	var i3 int64
	binary.Read(buf, binary.LittleEndian, &i3)
	fmt.Println(i3) // 511

}
