package base

type Module struct {
	files []File
}

func NewModule(files []File) Module {
	return Module{files: files}
}

func (p Module) Files() []File {
	return p.files
}
