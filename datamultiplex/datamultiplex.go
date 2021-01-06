package datamultiplex

// 向bitmap中添加一个属性
func AddAttribute(bitmap int64, attribute int64) int64 {
	return bitmap | attribute
}

// 判断该属性是否在该bitmap中
func IsAttribute(bitmap int64, attribute int64) bool {
	return (bitmap & attribute) == attribute
}

// 删除某个属性
func DelAttribute(bitmap int64, attribute int64) int64 {
	return bitmap & ^attribute
}

// 判断属性值是否是2的N次方
func CheckAttribute(attribute int64) bool {
	return (attribute & (attribute - 1)) == 0
}

// 根据bitmap 获取属性列表
func GetAttributeList(bitmap int64) []int64 {
	list := []int64{}
	for {
		if bitmap == 0 {
			break
		}
		list = append(list, bitmap-(bitmap&(bitmap-1)))
		bitmap = bitmap & (bitmap - 1)

	}
	return list
}
