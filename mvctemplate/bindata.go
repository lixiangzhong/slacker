// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// controllers.tpl
// dao.tpl
// js.tpl
// models.tpl
// service.tpl
// sql.tpl
// vue.tpl
package mvctemplate

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _controllersTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x51\x6f\xdb\x36\x10\x7e\x16\x7f\xc5\x4d\x0f\x83\x14\x68\x52\x3a\x14\x7d\xf0\xe0\x87\xc6\x09\xda\x6e\x41\xea\x2d\xc9\x53\x51\x0c\x0c\x79\x76\x88\x48\x94\x47\x52\xe9\x0c\x4d\xff\x7d\x20\xa5\xb4\x96\x2c\xba\x5e\x6c\x2c\x7a\x92\x2d\xf2\xbb\xef\xbe\xef\xee\x28\x65\xd9\x3b\x94\xa8\xa8\x41\x0e\x77\x6b\xd0\x39\x65\x0f\xa8\x80\x64\xd9\x52\x48\x50\x65\x65\x50\x91\xec\x84\xd0\x95\x48\xdf\x5d\xdc\x44\x61\x56\xd7\xe9\x65\xf9\x05\xd5\x15\x2d\xb0\x69\xb2\x89\xe0\x61\xc2\x4a\x69\x54\x99\xe7\xa8\x74\x7a\x43\x1f\xb0\xae\xd3\x19\x2d\x30\x9f\x51\x8d\xed\xba\xd8\x8b\xd0\xdf\x7d\x29\xb4\xf1\xed\x9e\x7f\xbc\xfe\xde\xf6\x99\x42\x6a\xbc\xe1\xe7\xb7\xfb\x24\x70\xbb\xe2\xbb\x30\xde\xde\xcc\xde\xef\x81\x32\xa7\x86\xdd\xc7\xdb\x28\x0e\xe4\xfc\xe2\xf2\xe2\xe6\x62\x0f\x94\x73\xcc\x71\x94\x0b\xec\x20\xd3\x87\x38\xb3\x44\x1c\x1b\x5f\x4a\x1e\x36\x23\x30\x5e\x3a\x5e\x79\x47\x40\xbc\xfa\x02\x39\xc9\x08\x59\x51\xf6\x40\x97\x08\x1b\xfb\x08\x11\xc5\xaa\x54\x26\x22\x01\x84\x75\x1d\x2e\x4b\xad\x58\x56\x94\x1c\x73\x1d\xfe\x03\xe9\x07\xf7\xf4\x52\xdc\x29\xaa\xd6\x4d\x63\x29\x74\xd1\x49\x10\x2e\x85\xb9\xaf\xee\x52\x56\x16\xb6\xa4\x7f\x5a\x96\x52\x30\x7b\x67\x9f\x49\x34\xd9\xbd\x31\x2b\x7b\xaf\x8d\x62\xa5\x7c\x0c\x49\x4c\x80\x90\x45\x25\x19\x00\x8c\x97\x63\xc4\xe0\x64\x29\x64\x3a\x2b\xa5\xc1\xbf\x4d\x0c\x50\x13\x00\x80\x47\xaa\xa0\x9f\xfe\xe0\x67\x3a\x52\x0f\x41\xb9\x58\x68\x34\xc9\x9f\x93\x69\x47\x21\x9d\x53\xa5\xf1\x56\x48\x13\xb1\xf4\x1c\x17\xb4\xca\xcd\xef\x15\xaa\x75\x14\xb6\x6b\xc3\x24\x3c\x0d\xe3\x04\x5e\x9d\x26\xf0\xe6\xb5\xe5\x6b\xa3\xe7\xa2\x10\x7b\xc2\xb8\xa5\x1d\xca\xab\xd3\xc4\x61\x04\x2c\xbd\xbe\x2f\xab\x9c\x9f\x09\xc9\xdb\x75\x3f\xf6\xd9\xc7\x36\x0c\xa7\x86\x26\xa6\x34\x34\x4f\x50\xa9\xc9\x14\xae\x51\x3d\x0a\x86\x9e\xce\x8d\xba\xec\x5a\x72\x43\xbc\x20\x08\x7e\xbd\xfe\x78\x15\xb1\x04\xac\x9e\xef\x6b\x12\x04\xa1\x8d\x10\x4e\xba\x48\xf6\x0f\x17\x2d\x9c\x40\x1b\x95\x04\x8d\x8d\x1c\x93\xe6\xab\x4b\xe3\x23\xe7\xb8\x2e\xed\x29\x0f\x08\x9e\x00\x2a\x05\x93\x29\xf4\x7c\xf8\xe0\x6c\x98\x53\x45\x8b\x28\x14\xfc\x9b\xf0\x81\x58\xb8\x0d\x3f\x4c\x41\x8a\x1c\xac\x04\x4f\x9a\x48\xe1\x44\x66\x25\xc7\xf4\x0a\xbf\x44\x4f\xf7\x67\x94\xff\x81\x7f\x55\xa8\x8d\x53\xc2\xea\xa8\xd0\x54\x4a\x92\xa0\x69\x6b\xa1\xcf\xac\x65\xf4\xcd\x2a\x8f\x5e\x82\xc7\xb0\x69\xc9\x00\xa4\x2f\xba\x6f\xd2\x1e\x55\x75\x27\xe8\xe2\x49\xcf\x4d\x07\x1c\xc7\xa1\x01\xbf\xf4\x85\x04\x77\x3d\x43\x4e\xe8\xae\xaf\xa2\x8e\x69\xda\x91\x7a\xd2\xd4\xab\xc7\x76\xcd\xef\xa3\xaf\x6f\x48\x8e\xc9\x1b\x1c\xa5\xea\x0e\x29\xba\x03\xfa\xea\xf9\xf6\x1e\xd6\x27\x24\xd8\x26\x37\x57\xa2\xa0\x6a\xfd\x1b\xae\x67\x65\x5e\x15\x72\xc8\x16\xa6\x20\x38\x09\xfa\xcd\xe4\x75\xca\x6f\x7c\xc7\x77\xa3\x9b\x3c\xe7\xf3\xa8\xdb\x2f\xed\x75\x60\x8d\xae\x5c\xd6\x0b\x81\x39\xd7\x30\x85\x82\x3e\x60\x54\xd0\xd5\x27\x6d\x94\x90\xcb\xcf\x42\x1a\x54\x0b\xca\xb0\x6e\xe2\xdd\x1e\x6f\x02\x6d\x39\x7c\xd8\x20\x1c\x18\xe5\xd3\x58\xf0\xa4\x47\x62\x2b\xec\x86\x4f\xbe\x17\x20\xdf\xd4\x3b\xd8\x2b\x38\xf8\x38\x18\xa8\xe0\xcd\x40\x70\x5f\x89\x66\x27\x5d\xf6\x3b\xdf\x01\xb7\x24\x70\xa5\xaa\xbb\xd4\x85\x5c\xea\x9f\x85\x34\x6f\x5e\xeb\x88\xa5\xee\xfc\x7c\xab\x14\x5d\xb7\xd9\xc7\xed\xd2\x29\xd0\xd5\x0a\x25\x8f\x04\xd7\xc9\xf7\x76\x7d\xfa\x1c\xc6\x71\x9a\xa6\x71\x97\xe2\xc6\x2c\xde\xcd\x53\x78\x3d\x26\x9b\x79\xee\xdb\x92\xf5\xff\xd9\x10\x87\xce\xbc\x97\x34\xc4\xdf\x7f\x3a\x81\x41\x07\xee\x68\xc1\x9d\xdf\x10\x3e\x77\xfe\xeb\xb9\x04\x47\x3d\x98\x5e\xc2\x84\x71\x0f\xbc\xba\xd9\xa6\x18\x99\x35\x31\x69\xec\x57\xd9\xbf\x01\x00\x00\xff\xff\x67\x3d\xdb\xf5\x26\x10\x00\x00")

func controllersTplBytes() ([]byte, error) {
	return bindataRead(
		_controllersTpl,
		"controllers.tpl",
	)
}

func controllersTpl() (*asset, error) {
	bytes, err := controllersTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controllers.tpl", size: 4134, mode: os.FileMode(420), modTime: time.Unix(1560927076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _daoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x4d\x6f\xdb\x38\x10\x3d\x93\xbf\x62\x2a\xa0\x81\xd4\x0a\x74\x0b\x2c\xf6\x60\x40\xbb\x68\xd3\x6c\x91\x5d\xf7\x23\x75\x8a\x3d\x04\x41\x4a\x4b\x23\x9b\xb0\x44\xda\x24\xb5\xa9\xa1\xea\xbf\x2f\x48\xca\x8e\xbf\x5a\xa7\x40\x7d\x11\x68\x0d\x87\x6f\xde\x7b\x33\xd4\x60\xf0\x16\x25\x6a\x6e\xb1\x80\xc9\x0a\x4c\xc5\xf3\x39\x6a\xba\xe0\xf9\x9c\x4f\x11\x0a\xae\x28\x15\xf5\x42\x69\x0b\x31\x35\x4b\x12\x4d\x85\x9d\x35\x13\x96\xab\x7a\xf0\x8e\x1b\x8b\xba\x16\xb2\x30\x03\xb3\x6c\x84\xd6\x58\x45\x34\x6a\xdb\x68\xaa\x8c\xce\x07\xb5\x2a\xb0\x32\xd1\x37\x60\x97\x3e\xc3\x48\x4c\x34\xd7\xab\xae\x1b\xb4\x2d\x7b\xcf\x6b\xec\xba\x88\x26\x94\x96\x8d\xcc\x21\x2e\xe0\xd9\x1b\xae\x12\x18\x09\x63\xdb\x96\x9d\xf3\x1a\xab\x73\x6e\x30\x04\xc6\xaa\x2c\x0d\xda\xb4\x12\xb5\xb0\xd0\x08\x69\x7f\xff\x2d\x35\xc8\x75\x3e\x83\xb6\x65\x23\x75\x8f\x3a\x44\xb2\xc3\xcd\x09\xc4\x37\xb7\x27\xa3\x52\x21\x6d\x8a\x5a\x2b\x9d\x40\x4b\x01\x00\xfe\xe3\x1a\x0a\x6e\x39\x64\x35\x9f\xe3\xa3\x72\xbc\x48\x36\x3b\xad\xb2\xbc\x02\x21\x2d\x25\x6e\x79\x3f\x43\x8d\x60\x96\xec\x95\x2c\x28\x69\x5b\x51\xc2\xb9\x92\x96\x0b\x69\x80\x8d\xef\x85\xcd\x67\x2e\x0f\x44\xc6\x72\x8b\x51\xd7\x51\x42\xc2\x9e\x0c\xf8\x62\x81\xb2\x88\xfd\x32\x75\x39\xde\x2b\x7b\xb1\x6c\xfb\xd0\xe1\x1e\xac\xb1\xfb\xf7\x0d\x56\xd0\x25\xee\x20\x94\x85\x4b\xb6\x4c\x81\xeb\xa9\x49\xe1\x0e\x86\x99\xcb\x31\xc6\x0a\x73\x1b\x47\xb9\x6a\xa4\x8d\xdb\x96\x7d\xd4\xa2\xe6\x7a\xf5\x0f\xae\xce\x55\xd5\xd4\x92\x85\x47\xcf\x61\x94\xb0\xbf\xb4\xaa\xe3\xe8\x41\xbc\x84\xfd\xeb\x10\x05\x5c\x09\xbb\x56\xe3\x65\x15\x07\x02\x50\xeb\x61\x56\xb0\x62\xc2\xde\xa2\x8d\xcf\x3c\x17\xe9\x32\x75\x08\x18\x63\x21\x46\x94\x2e\xec\x49\x26\x45\x15\xf8\x76\x3f\x8d\xb6\xd1\xd2\xd3\x9e\x86\x5d\xa8\xb5\x7f\xdb\x51\x32\x69\x44\x55\xa0\xde\xc5\xff\xec\x34\xb2\x0f\xba\x40\xfd\x7a\xe5\x23\x7e\x58\x25\x14\x68\xf2\x28\xa1\x44\x94\xe0\xad\xf6\xc7\x8b\x76\x73\x6c\xd6\x3f\xd9\xc8\xbd\x89\xfd\xfb\x84\x92\xce\x47\x07\x7f\x1e\x0d\xff\xe0\x5f\xf5\x0e\x0e\x1b\x02\x13\xe9\x1d\x64\xb0\x8e\xda\x63\x2f\x90\xd7\xd7\x78\xe6\xe9\xd8\xa3\xef\x38\x53\x1d\xa5\x6d\xcb\xde\xa1\x9d\xa9\xe2\x9a\xcf\xb1\xeb\xf6\x1b\xec\x5c\x23\xb7\x78\xa4\xc5\x76\x6d\xf4\x88\xb6\x8a\x4f\xf7\x03\xf8\x86\x82\xa4\xa5\x40\xb6\xa3\xe1\x1b\xb0\x57\x8d\x55\x35\xb7\x22\x0f\x88\x3e\x2f\x0a\x6e\xf1\xe2\xeb\x42\xa3\x31\x42\x49\x87\xbc\x2f\xd4\x34\x95\x6f\x4d\xe7\x5c\xcf\x8b\x4b\x51\x5c\x7c\xc5\x3c\x8e\x84\x34\xa8\xad\xeb\x33\x05\x1b\x03\x80\x41\xbb\x5e\x15\xe3\xab\x51\xd7\x45\xe9\x2e\xd8\xa0\xf1\x83\xff\x08\xe9\xf9\xdc\x0d\xf3\x9c\x3a\xc5\x0e\x4b\x3d\x34\xd2\x5e\xed\x4e\xc5\x80\x9d\x8d\xb8\xb1\x97\x1e\xe8\x65\x11\xef\xc8\xb7\x77\x1c\xf4\x1a\xee\x4a\x16\xa8\xf9\x15\x92\xf5\x7a\xb4\x94\xc0\x77\xe5\x38\x22\x84\x03\x7c\xf7\x3d\x05\x1a\x1f\xff\x23\xf2\xfb\xe1\x77\xaa\xfb\xb2\xe1\xa9\x88\x43\x15\xb7\xa8\x3c\x4a\xdd\x47\x6e\xf3\xd9\x11\xe6\x44\x01\xe1\x12\xe9\xd1\xd7\x7c\x71\x63\xac\x16\x72\x7a\x2b\xa4\x45\x5d\xf2\x1c\xdb\x2d\xbe\xc0\xa1\xdf\xa3\xe8\x1d\x5f\xec\xb0\xe4\xe7\xbc\x74\x55\xc3\xcd\x6d\x48\x46\x49\xa9\x34\xcc\x1d\x6d\x9a\xcb\x29\xc2\x9a\x2c\x4a\x88\xf1\x43\x1f\xe6\x5e\x0d\x80\xdc\x4d\xff\xb6\xdd\xba\x0b\xba\x6e\x48\x09\x21\x21\xe3\xe6\x16\xf0\xcb\x14\xca\xda\xb2\xf1\x42\x0b\x69\xcb\x38\xfa\xf2\xd4\x7c\xc9\x86\x4f\x4d\x94\xc2\x3c\x85\x79\x92\x50\xe2\x4c\x1b\x46\x53\x85\x32\x6c\x4a\x20\xcb\xe0\x05\x6c\x99\x5d\x8a\xca\x45\x51\x52\x0a\xac\x0a\xe3\x6f\x06\x8f\xdb\xb0\xbf\x95\x90\xeb\xb3\xa2\xd4\xcd\xc4\x00\xfd\xe6\xe4\x10\x8d\x6e\x21\x03\x51\x50\x72\xe7\xed\xfc\x68\xcb\x44\xcf\xa1\x87\xf1\x3c\xfa\x85\x8e\xe9\x29\x4f\x28\xd9\xf5\xc9\xc3\x9c\x7c\x83\x15\x5a\x3f\x29\xdd\x9f\xa2\x04\x76\x69\x3e\x1b\xd4\xd7\x7c\x52\x39\x70\xfb\x9e\x72\x73\xf5\xd0\x52\xaf\x57\xee\x19\x37\x06\xb5\xa3\xad\x27\x32\x81\x47\x4f\xc9\xed\xcf\x8e\x9f\x6d\x6c\x4a\x36\xe7\x3e\x48\x78\xad\xfc\xa6\x0d\xa4\x84\x92\xf1\xd5\xc8\xe9\x11\x19\x7f\xb1\xc0\x33\x28\xb5\xaa\xb7\x74\xd8\xd0\xfe\xb9\xdf\x73\x40\x29\x74\x5d\xf6\x67\xb8\x1c\xe1\x65\x44\xc9\xb6\xc2\xfe\xae\xdf\x9f\x68\xe3\xab\x51\x0a\x5b\x10\x7e\x66\xf0\x1d\x16\x7a\x69\x2e\xbe\x0a\x63\x0f\x69\x9e\x28\x55\x39\x6b\xef\x55\xf8\xf2\xe7\x2b\xdc\x2b\xd0\xc9\x81\xee\x4c\x7f\xc4\x6e\xc1\x57\x0d\xea\xd5\x27\x75\x1f\xbb\x22\x37\x35\xb2\x71\xce\x65\x7c\xe6\x37\x6d\xee\x19\x78\x92\xb9\x76\x83\xb3\xb3\xf5\xca\x2c\x2b\x76\xa1\xf5\x7b\xf5\x49\xdd\x9b\xed\xae\xb4\xba\x41\xdf\xbc\x6b\xc3\xba\x44\xe1\x66\xf7\x1f\x73\xff\x07\x00\x00\xff\xff\x9a\x4c\x5f\x5d\xb5\x0b\x00\x00")

func daoTplBytes() ([]byte, error) {
	return bindataRead(
		_daoTpl,
		"dao.tpl",
	)
}

func daoTpl() (*asset, error) {
	bytes, err := daoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dao.tpl", size: 2997, mode: os.FileMode(420), modTime: time.Unix(1560927076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x51\x6b\xdb\x30\x10\xc7\xdf\xfd\x29\x0e\x33\xb0\xbd\xba\xd6\x7b\x86\x59\x69\xbb\x8d\x41\x19\x63\x6b\xd9\xe3\x72\xb1\xcf\x8b\x88\x6d\x69\xd2\xb9\x59\x09\xfe\xee\x43\xb2\xd3\xd0\xc6\x49\xc8\xd8\x4b\x8b\x74\x77\xff\xdf\x5f\x77\xb2\x22\x04\xc0\x27\x6a\xc9\x20\x53\x09\x8b\x27\xb0\x35\x16\x2b\x32\x81\x6c\xb4\x32\x0c\x4b\x66\x0d\x95\x51\x0d\x84\x99\x70\x8b\xf0\x5d\xb0\x8d\x59\x56\x86\xc6\xe0\x95\xf0\xab\xf0\x39\xf8\x16\xd0\x42\xa1\xd4\x4a\x8e\x29\xd1\x95\xe8\x58\xd6\x56\x0c\x9b\x51\x10\xd0\x1f\x9f\x59\x75\x6d\xc1\x52\xb5\x70\x8f\x2b\x8a\x65\x99\xc0\x26\x00\x30\xc4\x9d\x69\x3d\x3f\xfb\x45\x1c\xcf\x05\x6a\x29\x36\x9b\xec\x4e\xad\xc9\x7c\xc1\x86\xfa\x5e\xbc\xd9\xc8\xb2\x9f\x27\x41\xbf\x2f\x76\x27\x2d\xc7\xbf\x3b\x32\x4f\xd3\x7a\xd1\x84\x5e\x94\xfa\x54\x00\x8d\x06\x1b\x3b\xf3\xe5\x01\x40\xef\x09\x7b\x88\x1b\x43\xc8\x14\x2f\x54\x39\xc1\xd0\xca\x1e\x82\x78\x84\xab\x0a\x00\x26\xbd\x3f\xe8\xd2\x09\xcb\x32\x3d\xa0\xdd\x1d\xed\xc7\x69\xc0\x57\xe4\x62\x79\x44\xdf\x87\xcf\x22\xec\x21\x6e\xa9\x26\x9e\x1e\x67\x39\x84\xce\x9d\xe8\xb5\x73\x35\xca\x1e\x18\xec\x11\xe5\x79\xfa\x62\xb4\xf0\x72\xb6\x93\xac\xa1\x4b\x3e\x31\x85\x73\x3b\x35\x1f\x4a\x5e\xdd\x28\x38\x71\xa5\x3c\x77\x1c\xff\x71\xf0\x81\x2b\xf0\x6f\xd8\xcf\xfe\x9b\x8d\x2b\x59\xd3\x80\x7a\x44\x03\x95\x32\x0d\xe4\xd0\xd2\x1a\x3e\x2a\xd3\xdc\x22\x63\x9c\x04\xe0\xf7\x33\xd4\x9a\xda\x32\x8e\x5c\x49\x94\x82\xfb\x37\xfc\xcd\x5a\x6c\x28\x99\xfa\x18\x06\xbb\xc3\xf3\xb0\xef\xda\xa9\x6e\x5d\x2f\x09\x4b\x32\x76\x36\x2e\x01\xa2\x1b\xd5\x32\xb5\x7c\x79\xff\xa4\x29\x9a\x41\xd4\x74\x35\x4b\x8d\x86\x85\xab\xbb\x2c\x91\x31\xf2\xb9\xfd\xe1\x43\x7e\xf0\xeb\x78\x77\xc0\xce\xd4\x90\x83\x36\xaa\x20\x6b\x33\x6a\x1f\xb3\x6b\xb4\xf4\xf0\xed\x0e\x2e\x20\xf4\x66\x07\x89\x57\x66\xdf\x63\xe1\x0a\x7e\xb2\x5a\x51\x9b\x87\x70\x31\x3c\x84\x99\x65\x64\xca\xfc\x6e\x00\x50\xab\x02\x1d\x36\x5b\x1a\xaa\x20\x77\xb0\x49\x57\x3f\x68\xf1\x5d\x15\x2b\xda\x1a\x93\x15\xc4\x53\x96\xf2\x1c\xc2\x30\x19\x3b\x32\xf6\x36\x5c\xdb\x99\x10\xce\xc1\x8e\xa6\x2c\x3f\xdb\x5f\xd3\xc2\x7a\xed\x93\x27\x18\xde\x64\xf7\x30\xde\xbb\x3d\x3f\xe7\x7e\x37\xc4\x09\x43\x99\x21\x5d\x63\x41\x71\xe8\x7f\x15\x52\x67\x26\x4c\xfe\x0b\xba\x0f\xfe\x06\x00\x00\xff\xff\x69\x64\x92\x51\x96\x06\x00\x00")

func jsTplBytes() ([]byte, error) {
	return bindataRead(
		_jsTpl,
		"js.tpl",
	)
}

func jsTpl() (*asset, error) {
	bytes, err := jsTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js.tpl", size: 1686, mode: os.FileMode(420), modTime: time.Unix(1542595017, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelsTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x6f\x9b\x40\x10\xc5\xcf\xcc\xa7\x18\xa1\x1c\x40\x4a\x77\xdb\x6b\x25\xf7\xe2\x4a\x55\x5a\x27\xa9\x45\x7a\xaa\x7a\x58\xc3\x84\x6c\xb2\x7f\xf0\xee\xa0\xd4\x5a\xed\x77\xaf\xc0\xc4\x72\x9a\x1e\x2a\x38\xf0\x06\xde\xfc\xe6\x0d\x2b\xe5\x17\x72\x14\x14\x53\x87\xbb\x03\x46\xa3\xda\x27\x0a\x20\x25\xab\x9d\x21\x4c\x49\xdc\x28\x4b\x39\x83\x94\xcd\x76\x03\x52\xa6\x24\x1a\xee\x9a\xed\x66\xae\x4d\x2f\x27\x81\xef\x3e\x61\xaf\xf9\x61\xdc\x89\xd6\x5b\xf9\x68\xbd\x0e\xde\xc9\xb8\x37\xbf\x8f\x9e\x97\x0f\x67\xd7\xda\x9b\xd1\xba\xb8\x74\xdb\x6e\x16\x9d\x33\x0c\xaa\x7d\x52\xfd\xcc\xdd\xf8\x67\x0a\x0b\x1c\xb4\x1d\x7c\x60\xac\x20\xee\x8b\xf2\x0c\x74\xad\x22\x53\xb0\xda\x75\x51\xc6\xfd\xa8\x43\x20\x53\x42\x0d\xd0\x7a\x17\xb9\x82\xa2\x61\xc5\x74\xfb\x6d\xf5\x7e\x79\xfc\x4c\x66\xf5\x01\x6a\x40\xe0\xc3\x30\x73\xd6\xca\x92\x59\xab\x48\x47\x16\x46\x0e\x63\xcb\x09\x10\x11\x53\x0a\xca\xf5\x84\x17\xfa\xf2\xa2\xf5\xe6\xe3\x4a\x9c\x46\xc5\x94\xa6\xd2\x1b\xfb\x52\xbe\x3b\x0c\xe7\x4a\xf5\x67\x0e\x6f\x2d\x39\xce\x79\x41\x90\xeb\x72\x86\x0c\x90\x92\xbe\x47\x71\x15\x7f\x44\x0a\x77\xd3\xfa\x73\x86\xfb\xd1\xb5\x58\xa5\x24\xae\x9c\x66\xad\xcc\x91\xfc\x37\xb5\xc6\x6b\x15\xe2\x83\x32\x5f\x9b\xdb\x9b\xaa\xc6\xea\xe7\xaf\xdd\x81\xe9\x12\x29\x04\x1f\x6a\x4c\x50\xcc\x71\xd9\x0e\xff\x70\x43\xf1\xaa\xbf\x48\x49\x7c\x57\x31\x3e\xfb\xd0\x1d\xe3\xbe\x09\xb9\xc2\xb2\x84\x22\x10\x8f\xc1\xe1\x63\xf4\x4e\x2c\xfc\x8a\xed\xf0\x7a\xda\xba\x86\x0c\x2f\x21\xff\x3b\xce\x1c\x7f\x12\x55\x3d\xfd\x10\xed\x7a\x4c\x78\x22\x96\xa7\x63\x59\x4e\x7b\xc3\xe9\x9e\x2f\xfc\x13\x00\x00\xff\xff\xe0\x18\x7f\x11\xcf\x02\x00\x00")

func modelsTplBytes() ([]byte, error) {
	return bindataRead(
		_modelsTpl,
		"models.tpl",
	)
}

func modelsTpl() (*asset, error) {
	bytes, err := modelsTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "models.tpl", size: 719, mode: os.FileMode(420), modTime: time.Unix(1560927076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serviceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\x4d\x6f\x9b\x4c\x10\xc7\xcf\xec\xa7\x18\xf9\x04\x8f\x10\xbe\x3c\xea\xa1\x95\x0f\xa9\x93\xb6\x91\xa2\x2a\x52\x93\x5e\x2c\x2b\x1a\xc3\x60\x6f\xbd\xb0\x74\x77\x49\x62\x51\xbe\x7b\xb5\xbc\xf8\x8d\xb5\xe3\x48\x3d\x71\x60\xe6\x3f\xbf\x99\xf9\xef\xc2\x78\xfc\x95\x72\x52\x68\x28\x81\xc5\x06\xb4\xc0\x78\x4d\x8a\x15\x18\xaf\x71\x49\xa0\x49\x3d\xf3\x98\x18\xe3\x59\x21\x95\x01\x9f\x79\xa3\x25\x37\xab\x72\x11\xc5\x32\x1b\xff\xca\x24\x57\x32\x1f\xeb\xdf\xe2\x75\xc4\xbc\x51\x55\x8d\x96\x52\xab\x78\x9c\xc9\x84\x84\x1e\xfd\x81\xe8\xb6\x49\xbc\xe3\x0b\x85\x6a\x53\xd7\xe3\xaa\x8a\xbe\x63\x46\x75\x3d\x62\x01\x63\x2c\x2d\xf3\x18\x7c\x0d\xff\xfd\x68\x2b\x05\x0f\xb8\xa6\xaa\x8a\xa6\x98\x91\x98\xa2\xa6\x36\xd8\xe7\x09\xf0\xdc\x7c\xf8\x3f\xf0\xab\x2a\xba\x93\x2f\xa4\xda\x17\xd1\x30\x36\x24\xa5\xa4\x0a\x2a\x60\x5e\x82\x06\x43\x20\xa5\xe0\xe3\x04\x74\x94\xa0\x8c\x4e\xea\x07\xc0\x3c\x45\xa6\x54\x39\xec\xd2\x58\xed\x62\xbc\xe3\xda\x38\x34\x64\x9a\x6a\x32\xa1\xe0\x19\x37\x50\x36\xb8\xa1\x26\x54\xf1\x0a\xde\x84\x0e\xfc\xd9\xfc\xed\xce\x78\x6e\x8e\xbb\x33\xd2\xa0\x08\x0f\x5a\xbc\x00\xaf\xe3\x3a\xee\xb9\x15\xdb\x76\x3e\x68\x7c\xaa\x08\x8d\x6b\x7c\x87\xe8\x97\xb4\xfb\x9e\x35\x56\x15\x4f\x21\xba\xd5\x8f\x9a\xd4\x03\x2e\x04\xd5\x35\xf3\xbc\xa1\x82\x7d\x9f\x63\x46\x53\x29\xca\x2c\x3f\x16\x84\x09\x68\xa3\x78\xbe\xd4\xd1\x83\x6c\x32\x1d\x10\xe7\x25\x02\x67\xd9\x7b\xd4\xfa\x45\xaa\xe4\x74\xd9\xe8\x26\x8f\xd5\xa6\x30\x7d\xa4\xa3\xf0\x79\x11\x5b\x98\xa7\xdd\x7e\x87\xb3\xba\xd5\x37\xaf\x5c\x9b\xf7\xf7\x03\x15\xf3\xbc\xde\x02\x87\xd9\x8d\x0d\x62\x99\x50\x74\x25\x14\x61\xb2\x69\x4a\x30\xcf\xab\xed\x42\x28\x4f\xec\x12\x1c\x07\xec\x42\x8f\x9c\x3c\x6e\x03\xd3\x3d\x16\xc9\x3f\x32\x5d\xe3\x28\x6b\x28\x80\xae\x74\x8b\x7c\x61\x05\x8b\xec\x46\xbc\x47\x13\xaf\xce\xdc\x5a\x61\xd9\x54\x80\x0c\x8b\x59\xeb\xc1\x39\xcf\x0d\xa9\x14\x63\xaa\xf6\xb8\xdc\x46\xe7\x29\x3c\xdb\x63\x29\xd7\x76\xca\xad\xd4\x6c\xe4\xf0\x4c\xf3\xe8\x6e\xd7\xf9\x27\x1b\xdf\xac\xb7\xe8\xc2\x7a\x85\x67\x14\x91\xdf\x62\x58\x5f\x59\xfd\x3e\x74\x1b\xeb\x74\x6d\xff\xb2\xc9\xf2\x2e\x05\x81\x09\xf4\x89\x36\xaf\x3e\xb2\xd0\xe1\x2a\x4e\x4f\xb2\x9b\x61\xe0\xde\xc0\x35\x09\x72\xae\x70\xfb\xe1\xd8\x0e\xf9\xa0\xde\x99\xbc\x00\x9a\x52\xae\x9d\x0c\xca\x83\xfb\xb3\xf2\x79\x63\x9f\x7e\xd9\x1d\xc2\xee\xfe\x09\xe0\x82\xdb\x0f\xda\xeb\x0f\xf6\x88\xbd\x73\x5f\xb0\xa3\x52\x1d\x7c\x47\xfa\xb4\x23\x3d\xb5\xd3\x2d\x5b\xfb\xb4\x76\x58\xa1\x5e\x25\x21\x3c\x59\xcf\x2c\x9a\xac\xa8\xff\x59\xf8\xa2\x64\xb6\x95\x98\xcd\x17\x1b\x43\x3b\x77\x84\x7d\xf4\x35\xa5\x58\x0a\x33\x95\xda\x04\xbb\xb9\x37\xfa\x7e\x23\x1e\xec\x76\xb9\x87\xf8\x13\x05\x4f\x06\x80\x21\xb5\xe4\x94\x14\x2f\x3b\xda\x85\x94\xc2\xb2\x76\xe2\x39\x17\x30\xd9\xd2\x4e\x65\x56\xa0\xa2\x6f\xa8\x57\x57\x79\x72\xcc\xbb\xaf\x17\x84\x70\xdc\x45\xc0\x6a\x60\x9d\x4b\xff\x06\x00\x00\xff\xff\x5d\xed\x47\xb7\x26\x09\x00\x00")

func serviceTplBytes() ([]byte, error) {
	return bindataRead(
		_serviceTpl,
		"service.tpl",
	)
}

func serviceTpl() (*asset, error) {
	bytes, err := serviceTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service.tpl", size: 2342, mode: os.FileMode(420), modTime: time.Unix(1560927076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xd6\x73\x2e\x4a\x4d\x2c\x49\x0d\x49\x4c\xca\x49\x0d\x0e\xf4\xa9\xad\x05\x04\x00\x00\xff\xff\x4d\x9a\x10\x9b\x13\x00\x00\x00")

func sqlTplBytes() ([]byte, error) {
	return bindataRead(
		_sqlTpl,
		"sql.tpl",
	)
}

func sqlTpl() (*asset, error) {
	bytes, err := sqlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql.tpl", size: 19, mode: os.FileMode(420), modTime: time.Unix(1542595017, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vueTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x6d\x6f\x1b\xc7\xf1\x7f\xaf\x4f\x31\x59\x38\x01\x99\xf0\x4e\x72\xfe\x08\xfe\xc8\x89\xc7\x1a\x51\xda\x22\x68\xd1\x06\x68\xf2\xca\x30\x82\xd5\xdd\x92\xdc\x7a\xef\xa1\xbb\x4b\x3d\x84\x20\x50\xb4\x49\x6c\xa7\x4d\xe4\xa2\xa9\x90\xe6\x01\xb6\x0b\xa7\x4f\x80\x15\xbf\x28\x5c\x3b\x56\xd2\x2f\x23\x52\xce\xab\x7e\x85\x62\x77\x8f\xbc\xbd\x07\x8a\x94\xa3\x04\x35\x50\x40\x20\x8f\xbb\xb3\x33\xb3\x33\xbf\x99\x9d\x9d\x13\x40\xfb\x29\xc7\x81\x1f\x92\x98\x70\x2c\x49\x08\x9b\xbb\x20\x18\x0e\x2e\x13\xee\x38\x1d\x58\x69\x4b\x12\xa5\x0c\x4b\xd2\x59\x01\x68\x13\xe6\x04\x98\x87\xea\x19\xa0\x1d\xd2\x2d\x10\x2c\x91\x3e\xea\x13\x1c\x12\x8e\xcc\xb8\xa1\xeb\x26\x3c\x02\x8f\xc6\x8c\xc6\xc4\x47\x92\x0f\x08\x82\x2d\x27\x4a\x42\xc2\x7c\x24\x08\xe6\x41\x5f\x91\x20\xb8\x20\x06\x9b\x11\x95\x6e\x8c\x25\xdd\x22\x6e\xca\xc9\x16\x89\xa5\x8f\xe2\x24\x26\x33\x8e\x39\x4f\x87\x4a\x12\xe5\xc3\x66\x82\xc6\xe9\x40\x42\xca\x49\x97\xee\x38\x34\x48\x62\x1f\xa9\xd1\x20\x89\x1d\x23\x0a\x41\xc0\x08\xe6\x78\x93\x91\xa9\x16\xae\xe4\x34\xb2\x55\x71\x69\x8a\x20\x65\x38\x20\xfd\x84\x85\x84\xfb\x68\x72\xe3\xb3\xc9\xb5\xaf\xc6\xef\xbc\x3d\x3e\x78\x80\x3a\xed\xd5\xa9\x24\x4b\xab\xd5\x82\x5a\xb0\x9c\xba\x9b\x03\x29\x93\x18\xe4\x6e\x4a\x7c\x94\x72\x1a\x61\xbe\x8b\xe0\x42\xc0\x68\x70\xd9\x47\x3f\xcb\x34\xae\xdd\x47\x67\x72\xe3\xb3\x47\x9f\xff\x49\xcb\x35\x6c\xe6\x30\x9f\x72\xe3\x44\x10\x69\xd9\xbb\x33\xb9\xff\xf6\xf1\xdf\xbe\x30\x6c\x26\x9f\xde\x3a\x7a\x78\xaf\x96\x59\x69\x67\x2b\xa5\xd1\xd9\x80\x42\x4f\x8d\x58\x1a\xa5\x09\x97\x5d\xca\x88\xbb\x45\x05\xdd\x64\xc4\x37\x18\xe0\xc9\x20\x0e\xcb\x5b\x0f\x18\x16\xc2\x47\x0a\x5b\x19\x23\xd4\x69\xd3\xe9\xf0\xd4\x04\x83\x94\x25\x38\x7c\x5e\x79\x82\x76\xc6\x9f\x1f\x8e\xdf\xfe\xec\xd1\xad\xbf\x4e\x6e\x1e\xda\x1b\x00\xc7\x59\xa4\xdb\xf7\x77\x94\x6e\xa8\xac\x84\xd1\x6c\x39\x55\xc2\x64\x3b\x56\xca\xe4\xba\x5c\xf9\xe2\x64\x5d\xaa\x6a\x04\x9c\x60\x39\x33\x0f\xf8\x60\x0c\xb4\xac\x56\x25\x7c\xa4\x6c\x20\x50\x67\xf2\xcf\x87\xe3\x77\x6f\xd6\x38\xd4\x04\xab\xdc\x65\xc4\x47\x3a\x14\xbc\xcd\x44\xf6\x95\xfa\x21\xdd\xca\xe2\xd9\x7a\x24\xcc\x91\x3a\x58\xbc\x10\x4b\xec\x23\xfd\xc3\x55\xcf\x68\xca\x65\x9b\x86\xb2\xef\xc1\xf9\xb5\xb5\xa7\x55\x10\x13\x46\x02\x49\x93\xd8\x09\xfa\x38\xee\x11\x05\xe3\x6c\x64\x43\x0f\xa8\xe8\x57\x16\xa3\x71\x6f\xca\x2e\xfb\x59\x48\x1b\x7a\xc6\x09\x12\x36\x88\xa6\x01\x32\x63\x8d\x40\xcb\xf4\xd1\x0b\x2f\xa0\x02\x20\xed\x45\x39\x82\x87\x43\xae\x04\xc3\x39\xda\x3a\x17\x24\xcc\xf3\xdd\x0d\x4d\x21\x46\xa3\x79\x02\x19\xde\x54\xe9\x69\x38\x54\x0b\xdc\x0d\x1c\x11\xb6\x81\x05\xf9\x09\x8e\xc8\x68\x64\x27\xa3\x69\x56\xd4\xf9\xcf\x11\x41\xa2\x15\x55\x5f\xa8\x10\x8f\x22\xc5\x71\x07\xda\xed\x36\xe8\x49\x97\x27\xdb\xee\x94\xbb\x16\x69\x58\x43\xa7\xd3\x81\xf6\xaa\xa6\xb6\x22\xd0\xce\xbd\x0b\x36\x0b\xc3\x21\x89\xc3\x85\x3b\x9b\xfc\xfe\xbd\xa3\x2f\x3f\x79\x9c\x9d\xe4\xf8\x15\xf4\x4d\x45\x12\x61\xc6\x2a\x68\x4d\x19\xa6\x39\xc2\x07\x69\xa8\x10\xae\x8f\x02\x1f\x7e\xba\xf9\x73\x12\x48\x17\x0b\x41\x7b\x71\x63\x38\x6a\xcd\x4c\xd2\x5c\xcf\x28\xcb\xb1\x50\x84\x38\x09\xa9\x44\x1d\x58\x9c\xfb\x6a\x34\x0c\x15\x14\x78\x9e\x64\x5f\x26\x8c\x48\xd2\xc8\x55\x28\x0b\x0b\x35\x41\x96\xf5\xab\xc9\x71\x49\xd7\xe4\xc3\xf6\x99\x69\x02\x28\xc2\xbc\x47\x63\x87\x91\xae\xf4\xe0\xf9\xb5\x74\x67\x3d\x1b\x91\x49\x6a\x06\x50\x4d\xf6\xd8\x72\x44\x3f\xd9\xf6\x51\x34\x60\x92\xa6\x8c\xcc\xc2\xcc\x65\x24\xee\xc9\x7e\xc5\x25\xd3\x1d\x6f\xa8\xaf\x97\xb0\x0c\xfa\xaf\xaa\x0f\xd4\x99\x5c\x7b\xf0\xf5\x95\xbd\xa3\x7f\x1d\x4c\x3e\x78\x50\x97\x31\x4e\x2f\xb3\x6c\xe4\x5c\xe4\xcb\x99\x35\x8d\xcc\xf1\xd5\x9b\x5f\xff\xf1\x76\x45\x66\x29\x01\xa5\xb8\x47\x55\x51\x90\xc4\xe0\xf5\x69\x48\x1c\x75\x0a\xd2\xb8\xc7\x88\x9a\x9a\x15\x15\x59\x6a\xcc\xa9\x11\x6c\xe2\xe0\x72\xcf\x24\x4e\x86\x77\x93\x81\xf4\x91\xc6\x84\x68\x81\x4c\x24\x66\x2d\x55\x2a\x6c\xb5\x40\xb1\xe1\x2d\x88\xc9\x8e\x44\xe0\xe9\xa9\x69\x6e\xd2\x3f\x10\x78\x8a\xc4\x31\x78\xca\x92\x16\x8d\xa8\xb4\x27\x84\x8f\x2e\x9e\x5f\x6b\xc1\x0b\x6b\x2d\x95\x0c\x5b\xf0\xe2\x8b\x2f\x5e\x52\x19\x91\xbe\x49\xf2\x64\x48\xdf\x24\xd3\x3c\x78\x21\x18\x70\x4e\x62\x39\x9b\x7c\x15\xf7\xa6\x93\x16\x6a\xf2\x0d\x75\x56\x66\x36\x09\x29\x66\x49\x0f\x24\x95\x0a\x41\x26\xdb\x23\xf0\xb2\xb8\x71\xc5\x6e\x1c\x94\x0f\x96\x3c\x6b\xae\x95\x21\x65\x6a\xb4\xac\x26\xcb\x56\x99\xa2\x4c\xa7\x0b\x27\x4d\x04\x55\x0a\xf8\x88\xd3\x5e\x5f\x4e\x87\x33\x7e\xe7\x0b\xfc\x96\xcb\xb8\x19\x19\xed\x42\x9c\x48\xd0\x79\xf0\x15\xf1\xaa\x81\xea\x8f\xc8\xae\x45\x57\x28\xa0\x96\x4e\xcc\x76\x39\x38\xad\xf3\xb4\xb8\x8d\x24\x96\x98\xc6\xc2\xc8\x7c\x6d\x37\x25\x80\x68\x2c\xd1\x68\x14\x0f\xa2\x4d\xc2\x87\x43\xc2\x04\x19\x8d\x54\x51\x98\xa5\xd2\x82\x49\xea\xb2\x36\x02\x3c\x90\x89\x13\x24\x51\xaa\xd0\xed\xa3\xa4\xdb\x5d\xaa\x4c\xb4\x8d\x51\x4a\xdc\x35\xa9\xbc\x52\x72\xe5\x75\x77\x37\x49\xa4\x0a\xb8\x2c\x06\x0c\x3a\x9c\x6c\xb4\x58\x3b\x2f\x2a\x3e\xba\x98\x09\x82\x3a\xe3\xbd\x7d\x98\xdc\xbb\x5a\x9f\xf5\x16\x96\xae\x1b\x9a\x2b\xea\x1c\xdf\x3a\x80\xf1\xc1\x47\x75\x59\xc5\xae\x37\x66\x80\x9e\x8b\xf0\xe3\xc3\xfd\x47\x5f\xfd\xae\x82\xf0\xe2\x71\xb1\x2c\xc2\xad\xe3\xe8\x71\x10\xbe\x1c\xba\x9f\x18\x70\x5b\xd6\xf8\xb6\xc0\x5d\x04\xf2\x77\x01\xeb\x4a\x1d\x71\x26\xb0\x7e\x5d\x73\x3d\x43\x58\xdb\x47\x6f\x05\xdc\x9b\xea\xc0\x4c\xd5\xc7\x69\x01\x6e\xad\x7c\xfc\x34\xbe\x64\x12\x7f\x72\x80\x5e\xb2\xca\x77\x03\xf6\xef\x06\xee\x55\xa8\x9c\x15\xe4\xed\x42\xf1\xf4\xb0\xcf\xfa\x20\x35\xc0\xdf\xbb\x7e\x7c\xfb\x0b\xef\xe8\xfe\xbb\x47\x87\x37\xc7\xd7\xdf\x1f\xdf\x7d\x67\xf2\xc1\x83\xf1\xde\x87\xe3\xf7\x6f\x4c\x3e\xb9\x36\xf9\xc3\xdd\xc9\x7b\x07\x95\x88\xa8\x36\x13\x66\x11\xf1\x7f\x6b\x4f\x17\xe2\xc1\xf4\x08\x80\x93\xae\x4a\x06\xfa\x8a\x0e\x9e\x75\x2b\xfd\x01\x65\xb3\x42\xcc\x53\x1c\x1d\x46\x85\x2c\x88\x50\x1f\x6a\x10\x81\xd7\x97\x32\x75\x38\xf9\xc5\x80\x28\x9a\x57\x22\xd3\x37\xf0\x34\x58\x0c\x77\x1f\x19\x73\x43\xc8\x71\x0f\x70\x60\xe2\xcc\xf6\xda\x9c\x36\x86\xe9\x1c\xe4\x64\x0a\x0e\x39\xa1\xa1\x79\xe3\x0d\xa9\x4a\xd3\x8e\xb2\xd2\xfe\x95\xa3\x87\xf7\x26\xbf\xd9\x1f\x5f\xbd\x3b\xb9\x73\x7b\x7c\xfb\xad\x7f\x1f\xfe\x76\x72\x75\xbf\x10\x33\x51\xe7\xf8\x57\x0f\xc6\x57\x1e\x1a\xf3\xb6\x57\x6d\x80\x5a\xae\x3a\x49\x1e\x4d\x51\x86\x4a\xf5\xd8\x19\xef\xfd\xfd\xd1\xaf\xbf\x34\xfc\xdc\x1d\x26\x76\x8c\x1e\x05\x66\xda\xf3\x86\xc1\x6c\x48\xdd\x64\xbf\x31\xba\xab\x6e\x3f\x2b\x74\x9b\x96\xe3\xeb\x99\x23\x4e\xc2\x77\x7e\x25\xb7\x01\x6e\xf0\xbd\x92\x0d\x9a\x96\xa8\x7d\x27\x5c\x69\x8b\x80\xd3\x54\xe7\x0d\xb3\x0b\x78\x16\xb0\x00\x9c\x52\xe8\xf2\x24\x02\x74\x61\x15\xa7\x74\x75\x38\x74\xb3\x04\xb4\xbe\x02\x40\x74\x5f\x0a\x42\xd2\xc5\x03\x26\x61\xa8\x85\x84\x58\xe2\x46\x33\xfb\x01\xc0\x89\x1c\xf0\x78\xf6\x13\xa0\x72\x27\xf3\xe0\xe2\xa5\xd6\x6c\x5a\x5f\x5a\x3c\x8b\x1e\x20\x6b\xbb\x78\xc6\x94\x2d\x6b\x46\xc9\x2a\x2e\x07\x73\x61\xf2\x60\xcd\x1e\x4b\xba\x5d\x41\x64\x69\x50\x5f\x8c\x3c\x38\xbf\x36\x1b\x1b\xe5\xd3\x79\x03\xd2\x83\xa1\x35\x6e\xca\xd0\xa2\x7e\x55\x2a\x80\xcc\xff\x99\xce\x75\x12\xcc\xc9\x7f\x16\x9c\xf2\xa4\x7a\x16\xdc\x72\x10\x97\xb8\x65\x59\xa6\x6c\xef\x79\x1c\xb3\xa7\xd1\xfa\x8a\x25\x20\x22\xb2\x9f\x84\x22\xe7\xdc\x23\xb2\x04\x18\x00\xd9\xa7\xc2\x2d\x74\xdc\xb2\xfe\xca\xfa\x8c\x04\xa7\xd4\x52\xc1\xfd\x31\x15\xb2\x51\x38\xec\x8b\xbd\x9b\xc2\x14\xc0\x70\x06\x08\x4b\x94\x19\x6a\x4d\x61\x61\x2b\xa1\x46\xa0\x60\xc6\x99\x9a\x39\x4e\x0a\xb3\xcd\x95\xfa\x67\x57\xf6\x49\xdc\xe0\x22\x05\xbf\x53\xb0\x6e\x61\xd7\x1a\xc2\xe0\x03\x17\xa9\x6e\x61\x9a\x81\xf5\x79\xe4\x8a\xc4\xa6\x56\x1f\x73\x89\x73\x8b\x6a\x7f\xd9\x74\xa3\x82\xa6\x81\x02\x54\x83\x70\xae\x55\x1d\x35\xa7\x94\x33\x3b\xe4\x5d\x80\xc6\x16\x66\x73\x1c\x68\xac\xaa\xfc\x57\xb6\xe7\xb3\xa0\x96\x81\x03\xe7\x9b\xeb\xc5\x95\x33\x4c\x54\x44\xe6\x5d\x89\x13\x44\x1a\xf6\x3e\x6c\xd9\x26\x5b\xc0\x58\xbf\x74\x30\xb1\xad\x7c\x59\x85\x63\xfd\xf5\x73\xbd\x96\x28\x6b\x1f\x0e\x47\xf5\x62\x72\xc8\x54\xc5\xe4\x73\x05\x06\xd9\xac\x79\xd1\x52\x67\x97\x6c\x62\x81\x0f\xd6\x96\xb6\x87\xb9\x17\x17\x18\x32\x22\xc1\x0a\x47\xb5\xfe\x5c\xf6\xbb\x51\xcc\xd6\xc1\x65\x4f\x87\x6b\x21\x2b\x93\x1d\xe9\x41\xd6\xce\x3d\xba\x7f\xc7\x75\x5d\x64\xcf\x8b\x94\xc6\x31\xe1\x1e\xcc\xaa\x8e\x69\xb3\xdd\xa6\xca\xfb\x64\x1e\x20\xde\xdb\xc4\x8d\xb5\x16\x64\x7f\xee\xff\x37\x51\x9e\x7d\x9a\x73\x73\x45\xb6\xb5\xb2\xc3\x4e\x13\xa5\x15\xb4\xd4\x44\x5b\xc5\xb6\xb0\x30\xc2\x2a\x1a\x34\x9a\x55\x05\x32\xb3\xb8\x01\x4b\x04\x29\x73\xaf\xf8\x31\xeb\x1a\xf3\x64\xbb\x82\x8d\x73\x41\x12\x77\x29\x8f\x1a\x68\x72\xe7\xb6\xf1\xcb\xf8\xee\x3b\xa6\xd3\xa9\x46\x3e\xbd\x65\x8a\xdb\x16\x4c\x3e\xfc\x7c\x7c\xfd\xcf\xc7\x0f\xff\x72\xfc\xf0\xce\xf7\x50\x0b\xb2\xa2\x18\xb5\x0a\xaa\x65\xec\x5e\xd2\xf5\xc8\x6b\xc6\xdd\xc7\xb7\x0e\xc6\x07\x1f\x15\x5c\x18\xe0\x38\x20\xac\x40\x35\xde\xdb\x9f\xdc\xbb\x5a\xa0\x52\x65\x90\x07\x68\x1b\xf3\x58\x61\x60\xa5\xde\x80\x73\x8d\x54\x74\xb9\xa6\xcd\x2d\xe1\x0e\x87\x6e\x7e\xf7\x33\xb7\xaa\xc2\xe5\xaa\x59\x5e\x3b\x1f\x0e\xb9\x35\xe3\x44\xd2\xee\xae\x2b\x06\x41\x40\x84\x68\xa0\xcc\x90\x57\xaf\x8f\xdf\xbd\x81\x4a\x00\x81\x13\x41\x52\xd9\x27\x9c\x94\x8e\xab\x56\x59\x98\xb8\x4d\x77\xa0\x10\xdc\xa5\x18\xc9\x28\xb4\x8e\xc5\xb6\xcb\xc9\x96\x6b\x41\x79\xc9\x69\xc2\xaa\xbe\x17\x52\x13\x5b\x15\x63\x4f\x3e\xfe\xc7\x64\xff\x6e\xbd\xb1\x1f\x37\x1a\x6b\x92\x6c\xe1\xfd\x60\xed\x09\x54\x29\x6e\x8b\xe7\x50\x9e\x5f\x4b\x2f\x1a\xaa\xa9\xdb\x0a\x4f\xeb\x15\xc4\xd7\xbf\xbc\x76\x74\xff\xce\xf1\x47\x6f\x4d\x3e\xbd\x75\xfc\xf1\x13\x19\x9b\x7a\x77\x85\xad\x2f\xf2\x49\xc6\xa5\xc6\x25\xf3\x2c\xf8\xc4\x9e\x55\x4a\x71\x1a\x0a\xf0\xe1\xe2\xa5\x7c\xb4\x9b\x70\x68\xe8\xa9\x38\x24\x3b\xfa\x20\xcf\x1e\xdb\x73\x60\x97\xbd\xe7\xca\xc8\x9e\x7b\xae\x59\x86\x83\x90\xa0\x3b\x5c\xfe\x1c\x06\x17\xf5\xc2\x4b\xb6\x6b\x68\x28\xdc\x74\x20\xfa\x0d\xb5\x70\x71\x12\x5d\xaf\x5c\x05\x2a\x69\xc6\x76\xdf\x10\x16\x71\xf4\x94\x61\x6a\x40\x36\x3f\xa1\x9c\x2a\x29\xff\x17\x9d\xda\xa5\x37\x9f\xd5\xe4\x50\xea\x10\xd6\xd5\x8a\xb5\x9d\x36\xfb\x12\x55\x0c\xa2\xaa\xa0\xff\xc5\xd0\x93\x13\x43\xc6\x7b\xb0\x6c\x0c\xb5\x6a\x51\xf4\x8d\x22\xcb\xbc\x1a\x38\x6d\x64\x9d\x84\xd5\x53\x5e\x4f\xbf\x85\x30\x34\xff\x76\x55\xae\x93\xdc\xe9\x70\x85\xde\xb4\x5b\xb5\x6b\xed\x35\x5b\x98\xeb\x33\x13\x7c\x4b\x1c\xc2\x69\xca\x68\xa0\x5f\x97\xaf\x6e\xc5\xa1\x9b\xa4\x24\xde\x89\x98\x72\x04\x96\xc2\x49\xba\x5d\x1a\x90\x30\x09\x06\x11\x89\xa5\x2b\x52\x4e\x70\x28\xfa\x84\xc8\x88\xb9\xfa\x1b\xe5\xda\xd3\x2e\xd8\x2d\x0e\x8d\x2d\xdd\x85\xd4\x62\x9f\xf2\x7d\x23\xff\x99\x67\x4e\x26\x2a\xe8\xb4\x9d\x4e\x95\xd0\x6d\xd4\x3c\xd4\x8a\x61\x60\xd0\x10\x11\x21\x70\x8f\xb8\x84\xf3\x84\x37\x90\xdd\x84\xd5\x3d\xd8\x9b\x87\xe3\xc3\xbd\x22\x2c\x4c\x5f\xb0\x0e\xe0\xcb\xa4\x9a\x42\xe0\xcf\x0b\x0d\xcb\x1f\x7a\xa7\xdf\xec\xe8\xd0\xff\xbe\x77\x02\xc0\xe7\x37\x7f\xcf\x18\xc4\x45\xfb\x34\xeb\xc0\x7c\xf2\x69\x56\x01\xae\xdd\x5d\xae\xa9\x44\x39\xe9\xaa\xe2\x5c\xcd\xba\x86\xb4\x06\xfc\xf9\x7b\x8a\x86\xb2\x41\x0b\x18\x15\xb2\xc2\xab\xe6\xad\x85\xca\xd3\xea\xdb\x4d\x93\xb4\xd1\xbc\x54\x61\x1c\x27\x31\x69\x34\x87\x06\x1f\xa3\x96\xe9\x64\x9b\xcb\xb7\xad\x6c\x69\xc3\x86\x4c\x2d\x1a\xad\xaf\xb4\x57\xa7\xdd\xed\x15\xf8\x4f\x00\x00\x00\xff\xff\xdc\x78\x49\x6f\x40\x2c\x00\x00")

func vueTplBytes() ([]byte, error) {
	return bindataRead(
		_vueTpl,
		"vue.tpl",
	)
}

func vueTpl() (*asset, error) {
	bytes, err := vueTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vue.tpl", size: 11328, mode: os.FileMode(420), modTime: time.Unix(1560927076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"controllers.tpl": controllersTpl,
	"dao.tpl":         daoTpl,
	"js.tpl":          jsTpl,
	"models.tpl":      modelsTpl,
	"service.tpl":     serviceTpl,
	"sql.tpl":         sqlTpl,
	"vue.tpl":         vueTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"controllers.tpl": &bintree{controllersTpl, map[string]*bintree{}},
	"dao.tpl":         &bintree{daoTpl, map[string]*bintree{}},
	"js.tpl":          &bintree{jsTpl, map[string]*bintree{}},
	"models.tpl":      &bintree{modelsTpl, map[string]*bintree{}},
	"service.tpl":     &bintree{serviceTpl, map[string]*bintree{}},
	"sql.tpl":         &bintree{sqlTpl, map[string]*bintree{}},
	"vue.tpl":         &bintree{vueTpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
