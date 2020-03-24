package string

// 判断回文字符串
func IsHuiWen(text string) bool {
	if text == "" {
		return false
	}

	result := true

	middleIndex := len(text) / 2
	preIndex := 0
	afterIndex := 0
	if len(text)%2 == 0 {
		// 偶数
		preIndex = middleIndex - 1
		afterIndex = middleIndex
	} else {
		// 奇数
		preIndex = middleIndex - 1
		afterIndex = middleIndex + 1
	}

	for preIndex > 0 {
		if text[preIndex] != text[afterIndex] {
			result = false
			break
		}

		preIndex--
		afterIndex++
	}

	return result
}
