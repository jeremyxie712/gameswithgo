package info

//The Information struct contains three key elements, the file path,
//the listening port and the template path
type Information struct {
	filePath string
	LisPort  string
	tmplPath string
}

//Return the file path
func (cfig *Information) GetFilePath() string {
	path := cfig.filePath
	return path
}

//Return the listening port
func (cfig *Information) GetPort() string {
	port := cfig.LisPort
	return port
}

//Return the template path
func (cfig *Information) GetTmplPath() string {
	tplPath := cfig.tmplPath
	return tplPath
}
