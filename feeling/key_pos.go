package feeling

type keyPos struct {
	LoR bool // Left:false | Right:true

	// 0-10, 1对应1qaz, 0对应0p;/, 56对应左右大拇指_+, 10表示错误的键
	Fin int
}

var keyPosMap = make(map[byte]keyPos)

func init() {
	// 左手空格"_" 右手空格"+"
	lk := "_12345qwertasdfgzxcvb"  // left keys
	rk := "+67890yuiophjkl;'nm,./" // right keys
	for i := range lk {
		keyPosMap[lk[i]] = keyPos{false, 0}
	}
	for i := range rk {
		keyPosMap[rk[i]] = keyPos{true, 0}
	}
	// 手指
	keys := "1qaz2wsx3edc4rfv5tgb_+6yhn7ujm8ik,9ol.0p;/'"
	fins := "1111222233334444444456777777778888999900000"
	for i := range keys {
		tmp := keyPosMap[keys[i]]
		tmp.Fin = int(fins[i] - 48)
		keyPosMap[keys[i]] = tmp
	}
}

// 获取按键左右手以及用哪根手指击键
//
// 左右大拇指按空格分别用 _ + 表示
func GetKeyPos(b byte) keyPos {
	return keyPosMap[b]
}
