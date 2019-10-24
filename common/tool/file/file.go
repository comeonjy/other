package file

import "os"

func FileExists(filename string) bool {
	_, e := os.Stat(filename)
	if e != nil {
		if os.IsExist(e) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	info, e := os.Stat(path)
	if e != nil {
		return false
	}
	return info.IsDir()
}

func IsFile(filename string)  {

}

func Write(filename string, ) {
	os.Stat()
}
