package feeling

import (
	"strconv"
	"strings"

	_ "embed"
)

//go:embed assets/equivalent.txt
var equivalent string

//go:embed assets/fingering.txt
var fingering string

type comb struct {
	dl   map[string]int      // 当量*10
	xkp  map[string]struct{} // 小跨排
	dkp  map[string]struct{} // 大跨排
	cs   map[string]struct{} // 错手
	xzgr map[string]struct{} // 小指干扰
}

var c comb

func init() {
	c.dl = make(map[string]int)
	c.xkp = make(map[string]struct{})
	c.dkp = make(map[string]struct{})
	c.cs = make(map[string]struct{})
	c.xzgr = make(map[string]struct{})

	// equivalent = strings.ReplaceAll(equivalent, "\r\n", "\n")
	// fingering = strings.ReplaceAll(fingering, "\r\n", "\n")

	// 当量
	for _, v := range strings.Split(equivalent, "\n") {
		tmp := strings.Split(v, "\t")
		if len(tmp) != 2 {
			continue
		}
		dl, _ := strconv.Atoi(tmp[1])
		c.dl[tmp[0]] = dl
	}

	// 指法
	fg := strings.Split(fingering, "\n")
	// 小跨排
	xkp := strings.Split(fg[0], " ")
	for _, v := range xkp {
		c.xkp[v] = struct{}{}
	}
	// 大跨排
	dkp := strings.Split(fg[1], " ")
	for _, v := range dkp {
		c.dkp[v] = struct{}{}
	}
	// 错手
	cs := strings.Split(fg[2], " ")
	for _, v := range cs {
		c.cs[v] = struct{}{}
	}
	// 小指干扰
	xzgr := strings.Split(fg[3], " ")
	for _, v := range xzgr {
		c.xzgr[v] = struct{}{}
	}
}

// 获取 10 倍当量
func GetDL10(s string) int {
	return c.dl[s]
}

// 获取当量
func GetDL(s string) float32 {
	return float32(c.dl[s]) / 10
}

// 判断是否为小跨排
func IsXKP(s string) bool {
	_, ok := c.xkp[s]
	return ok
}

// 判断是否为大跨排
func IsDKP(s string) bool {
	_, ok := c.dkp[s]
	return ok
}

// 判断是否为错手
func IsCS(s string) bool {
	_, ok := c.cs[s]
	return ok
}

// 判断是否为小指干扰
func IsXZGR(s string) bool {
	_, ok := c.xzgr[s]
	return ok
}
