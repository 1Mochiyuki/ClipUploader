package types

import "fmt"

type Uploader interface {
	Upload()
}

type Host struct {
	Name, BaseUrl, UploadUrl string
}

type Pomf struct {
	Host
}

type Lobfile struct {
	Host
}

type Catbox struct {
	Host
}

func NewHost(name, baseUrl, uploadUrl string) *Host {
	return &Host{
		Name:      name,
		BaseUrl:   baseUrl,
		UploadUrl: uploadUrl,
	}
}

func (h *Host) Upload() {
	h.Upload()
}

func NewPomf() *Host {
	host := NewHost("Pomf", "http://pomf.localhost", "http://pomf.localhost/upload")
	return host
}

func NewLobfile() *Host {
	return NewHost("Lobfile", "http://lobfile.localhost", "http://lobfile.localhost/upload.php")
}

func NewCatbox() *Host {
	return NewHost("Catbox", "http://catbox.localhost", "http://catbox.localhost/upload.php")
}

func (p *Pomf) Upload() {
	fmt.Printf("host: %s baseUrl: %s uploadUrl: %s\n", p.Name, p.BaseUrl, p.UploadUrl)
}

func (l *Lobfile) Upload() {
	fmt.Printf("host: %s baseUrl: %s uploadUrl: %s\n", l.Name, l.BaseUrl, l.UploadUrl)
}

func (c *Catbox) Upload() {
	fmt.Printf("host: %s baseUrl: %s uploadUrl: %s\n", c.Name, c.BaseUrl, c.UploadUrl)
}
