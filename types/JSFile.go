package types

type JSFile struct {
	FileName, AbsPath string
}

func NewJSFile(fileName, absPath string) *JSFile {
	return &JSFile{
		FileName: fileName,
		AbsPath:  absPath,
	}
}
