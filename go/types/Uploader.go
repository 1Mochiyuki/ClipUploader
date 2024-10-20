package types

import (
	"strings"

	"github.com/charmbracelet/log"
)

type Uploader interface {
	Upload(absPath string)
}

type Pomf struct {
	Name, BaseUrl, UploadUrl string
}

type Lobfile struct {
	Name, BaseUrl, UploadUrl string
}

type Catbox struct {
	Name, BaseUrl, UploadUrl string
}

func NewPomf() *Pomf {
	return &Pomf{"Pomf", "http://pomf.localhost", "http://pomf.localhost/upload"}
}

func NewLobfile() *Lobfile {
	return &Lobfile{"Lobfile", "http://lobfile.localhost", "http://lobfile.localhost/upload.php"}
}

func NewCatbox() *Catbox {
	return &Catbox{"Catbox", "http://catbox.localhost", "http://catbox.localhost/upload.php"}
}

func (p *Pomf) Upload(absPath string) {
	log.Infof("[%s UPLOAD] name=%s baseUrl=%s uploadUrl=%s", strings.ToUpper(p.Name), p.Name, p.BaseUrl, p.UploadUrl)
}

func (l *Lobfile) Upload(absPath string) {
	log.Infof("[%s UPLOAD] name=%s baseUrl=%s uploadUrl=%s", strings.ToUpper(l.Name), l.Name, l.BaseUrl, l.UploadUrl)
	// fileInfo, err := os.Stat(absPath)
	// if err != nil {
	// 	panic(err)
	// }
	// if fileInfo.Size() > (1024 * 1024 * 1024) {
	// 	msg := fmt.Sprintf("File: %s is over the 1 GB limit.\nSize: %d", fileInfo.Name(), fileInfo.Size())
	// 	log.Print(msg)
	// 	return
	// }
	// files := map[string]string{"files[]:": absPath}
	// data, err := json.Marshal(files)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// resp, err := http.Post("", "application/json", bytes.NewBuffer(data))
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// defer resp.Body.Close()
}

func (c *Catbox) Upload(absPath string) {
	log.Infof("[%s UPLOAD] name=%s baseUrl=%s uploadUrl=%s", strings.ToUpper(c.Name), c.Name, c.BaseUrl, c.UploadUrl)
}
