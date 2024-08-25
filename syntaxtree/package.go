package syntaxtree

type Package struct {
	files []File
}

func NewPackage(files []File) Package {
	return Package{files: files}
}

func (p Package) Files() []File {
	return p.files
}
