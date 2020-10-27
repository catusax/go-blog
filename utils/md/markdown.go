package md

//Cut 用于拆分hexo post文件为MD和yaml配置两部分
func Cut(data []byte) ([]byte, []byte) {
	count := len(data)
	for i := 0; i < count; i++ {
		if data[i] == '-' && data[i+1] == '-' && data[i+2] == '-' {
			for i2 := i + 3; i2 < count; i2++ {
				if data[i2] == '-' && data[i2+1] == '-' && data[i2+2] == '-' {
					return data[4 : i2-1], data[i2+4:] //以---分割数据为两部分
				}
			}
		}

	}
	return nil, nil
}

//GetDescription 根据<!--more-->标签获取摘要
func GetDescription(md []byte) []byte {
	count := len(md)
	for i := 0; i < count; i++ {
		if md[i] == '<' && md[i+1] == '!' && md[i+2] == '-' && md[i+3] == '-' && md[i+4] == 'm' {
			return md[:i+11] //返回<!--more-->标签前的内容作为摘要
		}
	}
	return nil
}

// //MDParse 用于解析hexo post文件
// func MDParse(file *multipart.FileHeader) ([]byte, []byte) {
// 	openFile, _ := file.Open()
// 	var data []byte
// 	count, err := openFile.Read(data)
// 	for i = 0; i < count; i++ {
// 		if data[i] == '-' && data[i+1] == '-' && data[i+2] == '-' {
// 			return data[0 : i-1], data[i+3:]
// 		}
// 	}
// }
