package info

//The Information struct contains three key elements, the file path,
//the listening port and the template path
type Information struct {
	FilePath string
	LisPort  string
	TemPath  string
}

//Return the file path
func (cfig *Information) GetFilePath() string {
	path := cfig.FilePath
	return path
}

//Return the listening port
func (cfig *Information) GetPort() string {
	port := cfig.LisPort
	return port
}

//GetTemplPath returns the template path
func (cfig *Information) GetTmplPath() string {
	tplPath := cfig.TemPath
	return tplPath
}
