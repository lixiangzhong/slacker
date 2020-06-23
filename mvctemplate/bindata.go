// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// component.tpl
// controllers.tpl
// dao.tpl
// js.tpl
// markdown.tpl
// models.tpl
// service.tpl
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

var _componentTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xcb\x6a\xdc\x4a\x10\xdd\xeb\x2b\x8a\xe6\x2e\x34\x17\x5b\xe3\xb5\x34\x1a\x4c\xb2\x4c\x30\x81\x2c\x8d\x17\x6d\xa9\x64\x35\xe9\x17\xdd\xa5\x49\x86\x41\x90\x6d\xc8\x97\x64\x95\x65\xfe\x28\x24\x7f\x11\xba\xf5\x18\x8d\x3d\xc6\x78\x23\x75\x55\x9f\x3a\xf5\x3a\xbd\x21\x54\x56\x72\xc2\x6d\x02\xb0\x41\x79\xe9\x51\x62\x45\x09\x00\xc0\xee\x52\x99\x1a\x65\x46\x4e\xa8\x92\x09\x6d\x3b\x62\xf1\xa2\x11\x92\xd0\xf1\x7b\x89\xd1\x74\xa8\x0c\x0d\x47\x2b\x79\x85\xad\x91\x35\xba\x92\xfd\xf9\xf9\xeb\xef\xd7\x6f\xbf\xbf\xff\x18\xa2\xae\x1b\x53\x75\xbe\x64\x0f\x48\x52\xf8\x91\x2a\x97\x86\xd7\x42\x3f\x94\x6c\x3c\x8c\xd8\xaa\xe5\xfa\x01\x4b\xd6\x72\x5d\x4b\xfc\x18\x8b\x1a\xae\x2a\x89\xfc\x98\x3b\xaf\x85\x0f\x46\x5d\xb2\xe9\x14\x60\xdb\x78\x19\xfa\x31\x96\x84\xd1\xd1\x0c\x1d\x35\xc6\x95\x4c\x10\x2a\x10\x1a\x8e\x65\x00\xe4\x9f\x70\x3f\xdc\x64\x87\x43\xf6\xc1\x09\xc5\xdd\xfe\x1d\xee\xdf\x1a\xd9\x29\x9d\x0d\xbf\x1b\xae\xb0\xef\xe7\x10\xc9\xef\x51\xbe\x36\x68\xc7\x65\x87\xaf\x0a\xda\x6e\xd6\x73\x23\x71\x4d\xeb\x79\x4f\xdb\x64\xb3\x3e\x6e\x30\xd9\xf8\xca\x09\x4b\x01\x24\x94\x35\x8e\xe0\x7f\xe0\x1e\xb8\x15\xd0\x38\xa3\x80\x5d\xaf\xb9\x15\xeb\xc3\x21\x1b\xf9\x8b\x04\x00\xbf\x44\x64\x8d\x0d\xef\x24\xc1\x21\xe6\xac\x39\xf1\x74\x35\x1a\x61\xc5\xd4\x39\x3d\x9b\x00\xe3\xb6\x72\x68\xb8\xf4\x78\x31\xfb\xa3\x4a\x72\x60\xec\xe8\x0a\x53\xce\xe1\xf6\x6e\x74\xf4\x45\x3c\xf4\x03\xc0\x3a\x63\x7d\x0e\x07\x88\x63\xc9\xe1\xa6\x53\xf7\xe8\x2e\x60\x5a\x66\x0e\x6f\x8c\x91\xc8\xf5\x14\xa0\x90\x5a\x53\x87\x90\x91\x6f\xa9\x90\x74\xb7\x5a\xd4\x28\x1a\x48\x69\x6f\xd1\x34\xb0\x83\xb2\x04\xa6\x23\x39\x5b\x62\x00\xa8\x15\x3e\xfb\x0f\x95\xa0\x74\xd4\xf8\x05\xec\x56\xc5\x33\x88\x41\x97\x8f\x20\x3d\xa0\xf4\xf8\x22\xeb\xd5\xcb\xac\x4b\x48\x3f\x0d\x6c\x1a\xe5\xf8\x70\xd2\x65\xfd\x91\x65\x5c\x06\x94\x40\xae\xc3\x23\x03\xb7\x62\x91\x30\x7b\x1f\x83\x97\x1e\x6a\x51\xa7\xce\x5b\x28\xb7\x27\xd5\x4f\xbc\xc2\x13\x94\xe0\xbc\xcd\x82\x1e\xe2\x67\xd9\x42\x7f\x42\x56\x71\xaa\xda\x14\x9d\x8b\x6c\xfd\xd3\x44\xe9\xea\xb9\x3c\x73\xfd\x51\x4d\xa7\x29\x8a\xc7\x63\x10\x5a\x9c\x9b\xc1\x50\xeb\xed\x5d\x71\xea\x8f\xc3\x0f\x93\x09\x46\x14\x59\x71\x2a\x90\xd9\x0f\x5b\xb8\x3a\x23\x8d\x79\xea\x67\x36\xb3\x54\xf2\xe7\xd0\xfe\x51\x96\xa3\x9e\x9b\x4e\x57\xe1\xdd\x3e\x2d\x78\x68\xa3\x38\x47\xa5\x4c\xa7\x09\xeb\x45\xcc\x93\x88\x80\xef\x8b\x64\xb3\x9e\x5e\x7c\xf2\x2f\x00\x00\xff\xff\x50\xb4\x2e\x04\xca\x05\x00\x00")

func componentTplBytes() ([]byte, error) {
	return bindataRead(
		_componentTpl,
		"component.tpl",
	)
}

func componentTpl() (*asset, error) {
	bytes, err := componentTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "component.tpl", size: 1482, mode: os.FileMode(420), modTime: time.Unix(1569292443, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllersTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x4f\x6f\xdb\x36\x14\x3f\x8b\x9f\xe2\x8d\x87\x41\x0a\x34\x29\x1d\x8a\x1e\x3c\xf8\xd0\x38\x41\xdb\x2d\x48\xbd\x25\x39\x15\xc5\xc0\x90\xcf\x0e\x11\x89\x72\x1f\xa9\x74\x86\xe6\xef\x3e\x50\xb2\xd7\x38\x16\x15\x6f\x31\x90\xed\x66\x44\x7c\xbf\xf7\xfb\xf3\xf8\x27\x79\xfe\x0e\x0d\x92\x70\xa8\xe0\x66\x09\xb6\x10\xf2\x0e\x09\x58\x9e\xcf\xb5\x01\xaa\x6a\x87\xc4\xf2\x23\x26\x16\x3a\x7b\x77\x76\x15\xf3\xbc\x69\xb2\xf3\xea\x2b\xd2\x85\x28\x71\xb5\xca\x47\x5a\xf1\x54\x56\xc6\x51\x55\x14\x48\x36\xbb\x12\x77\xd8\x34\xd9\x44\x94\x58\x4c\x84\xc5\x6e\x5d\x12\x44\xd8\xae\x3e\xd7\xd6\x85\xaa\xa7\x1f\x2f\x9f\x2a\x9f\x10\x0a\x17\x6c\x3f\xbd\xde\x47\xc0\xf5\x42\x0d\x61\xbc\xbd\x9a\xbc\xdf\x03\x65\x2a\x9c\xbc\x0d\x81\x9c\x9e\x9d\x9f\x5d\x9d\xed\x81\x72\x8a\x05\xf6\x72\x81\x01\x32\xdb\x10\x27\x9e\xc8\xbf\x61\xd3\x03\x13\xa4\x13\xb4\xb7\x07\x24\xe8\x2f\xb0\xa3\x9c\xb1\x85\x90\x77\x62\x8e\xf0\xa0\x8e\x31\x5d\x2e\x2a\x72\x31\x8b\x80\x37\x0d\x9f\x57\x96\x64\x5e\x56\x0a\x0b\xcb\xff\x84\xec\x43\xfb\xf5\x5c\xdf\x90\xa0\xe5\x6a\xe5\x29\xac\xbb\x6f\x15\x20\x91\xac\x14\xf6\x54\x70\x16\xf1\xb9\x76\xb7\xf5\x4d\x26\xab\xd2\x8f\xfe\x0f\xf3\xca\x68\xe9\x7f\xf9\x6f\xd6\x91\xac\xcc\x3d\x67\x09\x63\xb3\xda\x48\xe8\x1f\xd3\x58\xc2\xd1\x5c\x9b\x6c\x52\x19\x87\x7f\xb8\x04\xa0\x61\xd1\xbd\x20\xf8\x52\x23\x2d\x61\xdb\x99\x6c\xb7\x9e\x45\xd5\x6c\x66\xd1\xa5\xf0\x3b\x8c\xc6\xb0\xee\x9a\x4d\x05\x59\xbc\xd6\xc6\xc5\x32\x3b\xc5\x99\xa8\x0b\xf7\xab\x07\x8c\x79\xb7\x9c\xa7\xfc\x98\x27\x29\xbc\x3a\x4e\xe1\xcd\xeb\x84\x45\x85\x2e\xf5\xfe\x20\xed\xea\x35\xc6\xab\xe3\xb4\x45\x90\xd9\xe5\x6d\x55\x17\xea\x44\x1b\xd5\x2d\xfb\xbe\xd5\x90\xb0\x48\x09\x27\x52\x57\x39\x51\xa4\x48\x34\x1a\xc3\x25\xd2\xbd\x96\x18\xd8\xba\xf1\x5a\x52\xc7\x69\x03\xf2\xf3\xe5\xc7\x8b\x58\xa6\xe0\xdd\x7a\xdf\xb0\x28\xe2\x1e\x96\x8f\x00\x5a\x78\xff\x87\xb6\x05\x1f\x41\xd7\x8a\x45\x2b\xdf\x2e\x61\xab\x75\x02\xfd\xc7\x4c\x30\x01\x0f\xbb\x4f\x00\x3d\xc2\x7d\x69\xc2\x22\xad\x52\x40\xa2\x1d\x4f\x3f\xb4\x96\x4e\x05\x89\x32\xe6\x5a\x3d\x30\x51\xcf\xda\x82\xef\xc6\x60\x74\xe1\x89\xfc\x2d\xdb\xe8\xd6\x3c\x3f\x8d\xd9\x05\x7e\x8d\x37\xbf\x4f\x84\xfa\x0d\xbf\xd4\x68\x5d\x2b\x36\x61\x51\x44\xe8\x6a\x32\x2c\x5a\x75\xce\x7b\xd2\x53\xd2\xa5\xa0\xe5\x2f\xb8\x9c\x54\x45\x5d\x9a\xc7\x2a\xc6\x5a\xb1\xc8\x77\xfe\x96\x4d\xc0\xad\x8d\xb6\x0d\xaf\xd6\xfc\x2d\x9b\x43\x07\xea\xb3\x8d\x5e\x9b\x33\x1a\xc3\x43\xcb\x5b\x22\x1d\xab\x9f\x0e\x6b\xde\xba\xd9\xc6\x90\xa0\xae\xa7\x2d\x09\x9d\x5f\xbd\x96\xbc\xfc\xd4\xf8\x50\xb6\xf3\x78\x6e\x3c\xdb\xe5\x87\x0e\x6a\x97\xdc\x53\xf3\x0e\x63\xd8\x9d\xf8\x60\x4a\x8f\xd8\xb3\xc7\x74\xbf\x05\x1d\xb8\x34\xff\xc3\x39\xd7\xad\xe8\x99\xc6\x42\x59\x18\x43\x29\xee\x30\x2e\xc5\xe2\x93\x75\xa4\xcd\xfc\xb3\x36\x0e\x69\x26\x24\x36\xab\x64\x38\xe2\x87\x40\x87\x0d\x18\x1e\x07\x15\x72\x59\xab\x74\x8b\xc5\x40\x50\xa1\x67\x49\x5f\x52\x00\x00\xcf\x0e\x0b\x0e\x6d\x42\x50\x81\x56\x01\xe1\xf9\x51\xa7\x1d\x06\x1f\x66\x3b\x0e\xb4\xa3\x6a\xd7\xca\xb5\x99\xdb\x1f\xb5\x71\x6f\x5e\xdb\x58\x66\xed\x75\xf7\x96\x48\x2c\x3b\xf1\x49\xb7\x74\x0c\x62\xb1\x40\xa3\x62\xad\x6c\xfa\x54\xd5\xa7\xcf\x3c\x49\xb2\x2c\x4b\x60\xe7\xc4\x1d\xe6\xa9\x87\x13\x86\xa1\x57\x6c\x9f\xc8\xff\xc9\x6e\x78\xd9\x34\xc2\x5b\xcf\xa6\xb0\xef\xe6\x83\xc1\x47\x7d\x28\x9b\x7f\x7a\x23\xc1\x41\xaf\xa4\x97\xc8\xa0\x3f\x82\xa0\x6f\x7e\x43\x40\x9f\xed\xfe\xbf\xa4\xbf\x02\x00\x00\xff\xff\x19\x5e\xf5\x5c\xb6\x0f\x00\x00")

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

	info := bindataFileInfo{name: "controllers.tpl", size: 4022, mode: os.FileMode(420), modTime: time.Unix(1587021651, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _daoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x95\x51\x6f\xfa\x36\x14\xc5\x9f\xe3\x4f\x71\x97\x07\x94\x54\xc8\xd9\x5f\x9a\xf6\x50\x09\x4d\x2a\x6d\xa7\x6a\xb4\x2a\xa5\xd5\x1e\xaa\x6a\x72\xe2\x0b\x78\x24\x31\xb5\x9d\x52\xe6\xe5\xbb\x4f\xb6\x09\x83\xc2\x06\xff\xa7\x10\x71\xaf\xef\x39\x3f\x1f\x3b\x59\xf6\x2b\xd6\xa8\x98\x41\x0e\xf9\x1a\x74\xc9\x8a\x05\x2a\xb2\x64\xc5\x82\xcd\x10\x38\x93\x84\x88\x6a\x29\x95\x81\x04\x48\x14\x5b\x1b\xcf\xa4\x56\x45\x56\x49\x8e\xa5\x8e\xff\x06\x7a\xe7\xff\x1e\x89\x5c\x31\xb5\x6e\xdb\xcc\x5a\xfa\xc0\x2a\x6c\xdb\x98\x44\xf1\x4c\x98\x79\x93\xd3\x42\x56\xd9\x9f\xa2\xfe\x6b\xde\x64\x33\xa9\xaa\x98\x58\x2b\xa6\x40\xef\xf4\x8b\x46\xf5\xcc\xf2\x12\xdb\x96\x44\xb1\x36\x4a\xd4\x33\xed\x1a\x39\x33\x2c\x67\x1a\x33\xfd\x5e\xba\x72\xac\xb9\x2b\x49\x09\x99\x36\x75\x01\x09\x87\x8b\x6b\x26\x53\x18\x09\x6d\xac\xa5\x43\x56\x61\x39\x64\x1a\xc3\xe8\x64\x35\x47\x85\x00\x94\x52\x57\x9e\x5c\xb8\xa9\xf4\xfa\x2a\xdd\xfe\x80\xe4\xf5\xcd\x5a\x3a\x92\x2b\x54\xa1\x87\x1e\x2e\xd3\x47\xa5\xa4\x4a\xc1\x12\x00\x80\x0f\xa6\xc0\xc9\x82\x41\xc5\x16\x78\x56\xff\x8f\x29\x89\x50\x29\xb8\x1c\x70\xea\x27\xdf\x3b\x6c\xc9\xc9\x46\xdb\xa6\x74\x52\xc8\x25\xea\xe0\x84\x52\x9a\xd2\x5b\x51\xf3\xa4\xe7\x14\xa4\xf4\xc6\x09\xf3\xaa\x14\x9a\x46\xd5\x5e\x98\x93\x4b\x5a\x42\xac\xa5\xf7\x68\xe6\x92\x3f\xb3\x85\x03\xfb\x05\xd9\x50\x21\x33\x78\x04\x9a\x37\x77\x71\x52\x5c\x0a\x1e\x8b\x25\x10\xed\x29\x08\x0e\xc3\xea\xc9\xae\xcc\x03\x05\x2f\x4b\xfe\x3f\x0a\xce\x15\x00\x96\xb8\xf9\x59\x06\x7f\xf4\x03\x64\xe0\x94\xe7\x3e\x7d\xfc\xe6\x13\x8b\x24\x6e\xfc\x20\xd8\x46\x12\x34\x9a\xee\x8d\x4f\xc6\xa3\xb6\x85\x90\x14\x6b\xe9\xa3\x12\x15\x53\xeb\xdf\x70\x3d\x94\x65\x53\xd5\x34\x3c\x42\xe3\xe0\xf2\x54\x45\xdc\xf7\x96\x77\x89\x74\x48\x26\xec\x03\xf7\x37\xee\x80\xc8\x23\x33\xc5\xfc\x08\x10\xc1\x41\xd4\xe6\xe7\x9f\xfa\x1b\x27\x15\x5b\xbe\x86\x53\xf2\x26\x6a\x83\x6a\xca\x0a\xb4\x3b\x40\x00\x08\x40\xb4\x4d\xea\x49\x94\xae\xda\x55\xd2\xa3\xf6\xf6\x6b\x07\x82\x03\x89\xf6\x77\x3b\xe4\x39\x58\x0b\xbb\x1a\x7a\x75\x12\x04\xef\x18\xfe\x37\x96\xd7\x58\xa2\xf1\xc1\x24\xc7\x2e\x02\xf8\x0a\xc7\xc5\xf8\x50\xfa\xd5\xda\x3d\xbf\x33\xb4\x60\x81\x6c\x0d\xbb\x99\x35\xab\xf0\x98\x5d\x68\x5b\x18\xc0\xe6\x42\xa2\xcf\xd2\xaf\x9e\x9c\xd9\x99\x02\x74\xe7\xbe\x03\xf5\xbb\x8b\x59\x12\x1f\xe9\xdd\x66\x08\xda\x76\xf0\x4b\x88\xd1\x39\x33\xe8\xad\x50\xda\xec\x1d\xb4\x6e\x73\x60\x73\x11\xec\x63\x3c\x84\x72\xa7\x6f\x3e\x85\x36\x49\xb3\x19\xb5\xf1\x9b\x42\x2e\x65\x09\x96\x44\x93\xf1\xc8\x79\x88\x35\x96\x58\x18\xf8\x06\x53\x25\xab\x9d\xe3\xb4\x3d\x3c\xff\x69\xca\x79\x82\x52\x54\xc2\xc0\xb7\x98\xf8\x64\xa2\x9b\xe9\x47\xec\x40\xe2\x39\x1d\x37\xa8\xd6\x4f\x72\x95\x4c\xc6\xa3\x7e\x27\xc9\xdd\x83\xac\x4e\x7a\xbe\x29\x25\x91\x98\x3a\x73\xf0\xc3\x00\x6a\x51\x42\xaf\xd7\xbd\xe9\xf7\xd2\x41\x78\x90\x4f\x72\xa5\x9d\xf4\x8e\x85\x51\x0d\x92\xa8\xdd\xb2\xf1\x0b\x85\x5b\xd2\x7f\x53\xfe\x09\x00\x00\xff\xff\xbf\xa5\x8a\x8a\x02\x07\x00\x00")

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

	info := bindataFileInfo{name: "dao.tpl", size: 1794, mode: os.FileMode(420), modTime: time.Unix(1592897881, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x4d\x6f\xd3\x40\x10\xbd\xfb\x57\x8c\x56\x39\x38\xd0\xda\xf7\x54\x16\xd0\x0f\x10\x52\x85\x2a\x68\xe1\xd8\x4c\xec\x09\xb1\x62\xef\x9a\xdd\x71\x43\x65\xf9\xbf\xa3\x9d\x75\x14\x20\x8e\xa5\xf6\x52\x6b\x77\xde\xbc\xf7\xf6\x4d\x77\x93\xa6\x00\x9f\x48\x93\x45\xa6\x02\x56\xcf\xe0\x2a\xcc\xb7\x64\xa3\xb2\x6e\x8c\x65\xd8\x30\x37\xb0\xb6\xa6\x06\x95\xa4\x7e\xa1\x2e\xa2\x7d\xcd\xb1\xb1\x34\x14\xdf\xa7\xb2\xfa\xab\xfa\x06\xd0\x41\x6e\xcc\xb6\x3c\x60\x5a\x2e\x2b\x97\x86\x4d\x0f\xa5\xdf\x02\x5d\xb7\x3a\xe7\xd2\x68\xb8\xc7\x2d\xc5\x65\x31\x87\x2e\x02\xb0\xc4\xad\xd5\xe2\x20\xf9\x49\x1c\x2f\x53\x6c\xca\xb4\xeb\x92\x5b\xb3\x23\xfb\x05\x6b\xea\xfb\x74\xd6\x95\x45\xbf\x9c\x5f\x44\xfd\x31\xdb\x6d\xe9\x38\xfe\xd5\x92\x7d\x1e\x27\x54\x23\x84\xea\x4c\xa0\x00\x0d\x5a\xac\xdd\x02\xa4\x3f\x02\xe8\xc7\x35\xae\x2c\x21\x53\xbc\x32\xc5\x88\x48\x63\xdc\x49\x15\xe9\x18\xa5\x7c\x68\x0a\x4f\x59\x16\x03\xe8\x98\xb6\x9d\x0c\x63\x8a\xfb\x0e\x39\xdf\x4c\x51\x4b\xfd\x95\xe4\xd7\x54\x11\x8f\xcf\xaf\x08\xa5\x17\x8f\xf0\xd2\xfb\x19\x78\x4f\x4c\x72\x82\x7a\xf9\xc2\x61\x8a\x5a\x88\x48\x90\x2f\x4f\x69\xc8\xe7\x35\xba\xc3\xd8\xa7\x85\x4f\x4c\xfe\x75\xb2\x9f\xe5\xa2\xc6\xeb\xb2\xa2\x20\xf5\x84\x16\xd6\xc6\xd6\x90\x81\xa6\x1d\x7c\x34\xb6\xbe\x46\xc6\x78\x7e\x11\x81\x14\x12\x6c\x1a\xd2\x45\xac\x7c\x8f\x3a\x03\xff\x09\x7f\x13\x8d\x35\x09\xee\xe8\x02\x04\xc3\xe1\x55\x38\xf6\xed\x69\xf7\xbe\x37\x84\x05\x59\xb7\x18\x96\x00\xea\xca\x68\x26\xcd\xe7\xf7\xcf\x0d\xa9\x05\xa8\xba\xad\xb8\x6c\xd0\x72\xea\xfb\xce\x0b\x64\x54\x82\xed\x4f\x1f\xf3\x46\xd6\xf1\xe1\x88\xad\xad\x20\x0b\x41\x59\x93\x93\x73\x09\xe9\xa7\xe4\xfb\xc3\xcd\xe3\x87\xbb\xbb\xc7\x4b\x74\xf4\xf0\xf5\x16\xde\x0a\x22\xdc\xde\x40\xf9\x9f\xf9\x77\x98\xfb\xe6\x47\x36\x5b\xd2\x99\x1a\x1a\xe4\x19\x4c\x1c\x23\x53\x22\x15\x9f\x49\x65\x72\xf4\x5e\x92\x8d\xa5\x35\x64\xde\xc1\xa8\xd5\x1f\xb4\xfa\x66\xf2\x2d\xed\xdd\xe6\x46\x3b\x06\xa1\x81\x6c\x78\x4c\xfd\xd3\x75\xef\x77\xc2\x58\x2a\x62\xd8\x18\xc7\x90\x81\x52\x7e\xa3\x5c\x43\x3c\x75\xae\xcc\x03\xe7\xfb\xc4\x43\xe7\x72\xe7\x16\x69\x3a\xeb\x0e\x3e\x8d\xe3\x7e\xe9\xe9\x7a\xa0\xca\xd1\xbf\xf0\x09\xfa\xc4\x52\x53\x61\x4e\xb1\x92\x1f\x8b\x33\x50\x3b\xa7\xc4\x68\x7f\xf8\xdf\x58\xce\x3a\x11\x90\x6c\x77\xb4\x72\x72\xe6\xc9\x78\x67\x9d\x7c\xbd\xa7\x3e\xfa\x13\x00\x00\xff\xff\xc7\x4e\x02\x78\xb5\x06\x00\x00")

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

	info := bindataFileInfo{name: "js.tpl", size: 1717, mode: os.FileMode(420), modTime: time.Unix(1578994132, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _markdownTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\xcd\x6e\xd3\x4c\x14\xdd\xcf\x53\x5c\xb9\xd5\xb7\x4a\xeb\xae\x2d\xd5\xdf\xa2\x0b\x36\x88\x15\x62\x01\x42\xaa\x89\xa7\x25\xd4\x3f\xc5\x76\x84\xa2\xc9\x48\xa1\x0a\x10\x14\x9a\x54\x22\x2d\x25\x25\x85\x40\x52\x75\x43\x1a\x24\x0a\x6d\x1c\xda\x97\xf1\x8c\x9d\xb7\x40\xf6\x38\x28\x81\x12\x58\x30\xab\xb9\x47\xe7\x9e\x73\xee\xcc\x95\xe5\x6b\xd8\xc2\x8e\xe6\x61\x1d\xee\x15\xc0\x35\xb4\xec\x06\x76\xd0\x1c\x10\xb2\x78\x43\x33\x31\xa5\x88\x7d\x7c\xc5\x7b\xa7\x6c\x67\x1b\x8a\xc0\x4a\xc3\xf0\x93\xcf\x0e\xab\x50\x84\xe8\xe4\x33\xdf\xaf\xa1\x85\xf1\x29\x2e\x5c\x75\x45\x84\x38\x9a\xb5\x8e\x61\x3e\x97\x99\xcf\xda\x86\xb2\xbc\xb8\x62\x1b\x79\xd3\x72\x29\x05\x42\x62\x28\x05\x84\x1d\x14\xc7\xe8\xcd\xc2\xe6\x54\xbd\x62\x9b\x26\xb6\x3c\x4a\x11\x21\xd8\xd2\x29\x45\x68\x6e\x0e\x1e\xb8\xb6\x15\x76\x06\xc1\x45\x15\xad\xae\xae\xc6\x15\x22\x33\x3d\x01\x40\xba\xc2\x57\x52\x52\xf0\x36\x76\xec\x5b\x9a\x91\xc7\x94\x66\x7e\x38\xd1\x58\x3c\xf1\xe3\xb5\x2e\xab\x7f\x88\x2e\x1b\xec\xe0\xf0\x17\xdf\x44\x3c\x6b\xeb\x58\x52\x96\x32\xa2\x32\xb1\xeb\x6a\xeb\x58\x52\x24\x7b\x43\x4a\x31\x5d\xf3\x34\x49\x11\xf4\x7f\x16\x55\x88\x89\xb8\xf1\x7d\x32\x34\xab\x1c\x30\x7f\x00\x48\x85\x4d\xdb\xf5\x92\x37\x03\x56\xe9\x83\xac\x6d\xe6\x64\x42\x16\xaf\xdb\x8f\xb0\x93\x7e\x77\xc2\x7f\x72\x2c\x7e\x3d\xb8\xec\xf1\xc6\x79\xdc\x97\x9f\xdd\x26\x2b\x39\x1d\xc4\x03\xbd\x78\xc6\x7a\xcd\xc9\xf6\xc4\x57\xf3\xb2\xf7\xff\x42\x41\x55\x21\xb8\xa8\xb2\xa3\x2d\x48\x5b\x88\x34\x7a\x53\x8a\x8e\x1e\x0b\xa9\xb0\x59\x16\xca\x92\xc2\x4a\x43\xfa\x87\x34\x62\xf4\x77\xa3\xd7\x1d\xa4\x82\x8e\x0d\xec\xe1\xdf\xb3\x21\x66\x47\xb5\xaf\xac\xbe\x17\x9c\x95\x78\xab\x8d\x54\x58\xc7\xde\xec\x61\x05\x9f\x75\x9a\x33\xf9\x48\x05\xfe\xb6\x1b\x9d\xbc\xe7\xad\x76\xe0\x7f\x09\xfc\xee\xc3\x3c\x76\x0a\xac\xbe\xc5\x77\xfb\x7c\xef\x9c\x0d\xeb\xa1\x5f\x66\x3b\x95\xcc\xc8\xdf\x8f\x7a\x9d\xff\xed\xb5\x35\x17\x7b\xcb\x4b\xff\x19\x39\x33\xe7\x2d\x2f\x21\x15\xa9\x20\x76\x0e\xa9\x30\xde\x37\x15\x08\x52\x61\x7a\xe7\xd2\xfa\xa7\xad\x4b\xd1\xf1\xde\x89\x32\x81\x3c\xdb\xd3\x8c\xb8\x13\x64\x99\x7d\x7b\xc9\x9e\x6f\x4f\x25\x3d\xab\x86\xcd\x32\x2f\xf9\xbc\xd5\xe6\xbb\xfd\x4c\xd8\x38\x0e\x06\x35\x56\x79\x3a\x6a\x9f\x4e\xaa\x08\xe1\x3b\x77\x53\x2c\x1e\x98\x8a\xa0\xdf\x03\x00\x00\xff\xff\x2c\x19\x42\x29\x60\x04\x00\x00")

func markdownTplBytes() ([]byte, error) {
	return bindataRead(
		_markdownTpl,
		"markdown.tpl",
	)
}

func markdownTpl() (*asset, error) {
	bytes, err := markdownTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "markdown.tpl", size: 1120, mode: os.FileMode(420), modTime: time.Unix(1563850859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelsTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x6f\xd4\x30\x10\xc5\xcf\x3b\x9f\x62\x14\xf5\x90\x48\xc5\x86\x2b\xd2\x22\xa1\x45\x42\x0b\xdb\x96\xd5\x96\x13\xe2\xe0\x4d\xa6\xa9\xdb\xd8\x0e\xb6\xa3\x12\x0d\xfe\xee\xc8\x49\x58\xf5\xdf\x81\x5b\x66\xc6\xcf\xcf\xbf\x37\x91\xf2\x33\x59\xf2\x2a\x52\x83\xc7\x11\x43\xa7\xea\x7b\xf2\x20\x65\x54\xc7\x8e\x90\x59\x5c\x2a\x43\x29\x81\x94\x87\xfd\x0e\xa4\x64\x16\x87\xd8\x1c\xf6\xbb\xa9\x97\x87\xb9\xc0\x37\x1f\xb0\xd5\xf1\x76\x38\x8a\xda\x19\x79\x67\x9c\xf6\xce\xca\xf0\xab\xfb\x3d\x6b\xfe\x1d\x9c\x54\x1b\xd7\x0d\xc6\x86\xe5\xb6\xfd\x6e\xa9\x53\x82\x5e\xd5\xf7\xaa\x9d\x7c\x77\xee\x81\xfc\x62\x0e\xda\xf4\xce\x47\x2c\xa1\x60\x2e\x5a\x17\x7c\x2d\x8d\x6b\xa8\x0b\xc5\x1f\x14\xdb\x69\xb8\xd3\x47\xaf\xfc\x98\x52\x01\xcc\xfa\x06\xc5\x36\x7c\x0f\xe4\xaf\x33\x47\x4a\xb0\x2a\xc8\xd6\xae\xd1\xb6\x95\x77\xc1\xd9\x7c\x88\x6c\x93\x12\x54\x50\x3b\x1b\x62\x09\xab\x43\x54\x91\xae\xbe\xae\xdf\x2e\x9f\x9f\xa8\x5b\xbf\x83\x0a\xe0\x66\xb0\x35\x6a\xab\x63\x59\x21\xe3\x6c\x2c\x3e\x0e\xd1\x5d\xe8\x36\x47\x57\x32\x8b\x8d\x32\xd4\x6d\x54\xa0\xf9\xc9\x9c\x2a\x4c\x80\x10\xc7\x7e\xa2\x79\x36\xc6\x10\xfd\x50\x47\x06\x44\x44\x66\xaf\x6c\x4b\x78\xa6\xcf\xcf\x6a\xd7\xbd\x5f\x8b\x53\x20\xc8\x9c\x5b\x2f\xe4\x4b\xfb\x7a\xec\x1f\x57\xaa\x7d\xa4\x70\xc6\x90\x8d\x29\x2d\x16\x33\x6c\x82\x57\xc3\x99\x00\x33\xc5\xd6\xea\xa8\x55\x37\x3b\x3f\x77\xad\xf0\x42\xf9\x70\xab\xba\x2f\x87\xab\xcb\xb2\xc2\xf2\xc7\xcf\xe3\x18\xe9\x1c\xc9\x7b\xe7\x2b\x64\x58\x4d\xb8\xd1\xf4\xaf\xa8\x61\xf5\xe4\x7e\xc1\x2c\xbe\xa9\x10\x1e\x9c\x6f\x66\xdc\x17\x90\x6b\x2c\x0a\x58\x79\x8a\x83\xb7\x98\x97\x26\x16\xff\x32\x9a\xfe\xe9\x6b\xab\x0a\xd2\x69\xa3\xff\x8d\x33\xe1\xe7\xa2\xac\xf2\x42\xb4\x6d\x91\xf1\xe4\x58\x9c\x7e\xfe\x22\xe7\x06\x08\x88\x7f\x03\x00\x00\xff\xff\x5b\xf7\x40\x23\x2f\x03\x00\x00")

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

	info := bindataFileInfo{name: "models.tpl", size: 815, mode: os.FileMode(420), modTime: time.Unix(1592897756, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serviceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x6f\xe3\x36\x10\x3d\x4b\xbf\x62\x22\x2c\x0a\x69\x21\x30\x7b\x68\x7b\x48\xa1\x16\xa9\x93\xb6\x41\xb3\xbb\x01\x92\x6d\x0f\x41\xb0\xa0\xa5\x71\x4c\x44\x22\x05\x92\x8e\xd7\x60\xf5\xdf\x0b\x92\x92\x25\x3b\x92\xed\xcb\xe6\x62\xc5\x9c\x8f\xf7\x86\xf3\x66\xe4\xf3\xf3\x3f\x91\xa3\xa4\x1a\x0b\x98\x6f\x40\x95\x34\x7f\x41\x19\xd6\x34\x7f\xa1\xcf\x08\x0a\xe5\x2b\xcb\x31\x0c\x59\x55\x0b\xa9\x21\x0e\x83\xc8\x98\xe8\x59\x28\x99\x9f\x57\xa2\xc0\x52\x45\xff\x01\xb9\x71\xa7\xb7\x6c\x2e\xa9\xdc\x34\xcd\xb9\x31\xe4\x13\xad\xb0\x69\xa2\xa1\xbd\xca\x45\x8d\x23\xe6\x51\x18\x18\xc3\x16\x40\x6e\xd4\x17\x85\xf2\x81\xce\x4b\x6c\x9a\xa1\x27\x4a\x99\x8b\x62\xd2\x17\x79\xd1\x34\x61\x12\x86\xe1\x62\xc5\x73\x88\x15\xbc\xbf\xf7\xc0\x93\x07\xfa\x82\xc6\x90\x19\xad\xb0\x9c\x51\x85\x1e\x56\x5c\x50\x4d\xe1\xbd\x31\xe4\x56\xac\x51\xfa\x2f\xc9\x5b\xbb\x04\xa5\x14\xd2\x84\x81\x44\xbd\x92\x1c\x14\x29\xa8\x20\x07\x62\x26\x61\x33\x02\xe2\x96\x29\x3d\xe2\x20\x16\x0b\x85\x3a\x2d\x59\xc5\x34\xac\x18\xd7\x3f\xff\x98\x2a\xa4\x32\x5f\xc2\x71\x64\xf1\xe3\xd3\x51\xa3\x94\x71\x9d\x3a\x0a\x89\x81\x30\x78\xa5\x12\xd6\x4b\x94\x08\xee\x26\xc8\xbf\xf6\x59\xb5\xc5\x9f\x09\xae\x29\xe3\x0a\xc8\xfd\x9a\xe9\x7c\x69\xa3\x40\xa4\x34\xd5\x18\x35\x0d\x84\x41\xe0\x5c\x33\x5a\xd7\xc8\x8b\xd8\xfd\x93\x0e\xe2\xc4\xde\xf6\x2c\xfb\x2d\x4a\xf7\x80\xdd\xdb\x83\x2b\x2c\x93\x04\xfa\xdb\x0a\x8c\x91\x94\x3f\x23\xbc\x63\xe9\xbb\x5c\x94\x70\x91\x91\x99\x28\x57\x15\x57\xf6\x74\x0f\x94\xb5\x68\x8f\x1f\x36\x35\x42\x94\x2f\xa9\x8c\x9c\x61\xc0\x16\xe0\xab\x46\x8c\xf1\x76\xbb\x55\x80\xb3\x0c\xa2\x08\x8c\xb5\x3d\xc6\xa1\x8b\xe0\x32\x79\x77\x4b\xe8\x60\xfc\x24\xb1\x91\x3d\xe6\x96\xdb\x31\xf4\x8c\x6b\x0b\x1e\x06\x7f\xce\xe3\x93\xd0\x63\x4e\x36\x4f\x7f\x17\x36\xdb\x71\xd2\xbf\xc2\x07\x4f\xf9\x7b\x92\x0e\xf6\x39\xf4\xfc\xb7\xb7\xec\x1f\xec\xb1\x16\x9a\x96\xb6\x1f\x2f\x32\x2f\xa4\x99\x58\x71\x1d\xfb\x3c\xa9\x03\x46\x08\x49\x9c\x2d\x5b\x00\x4a\x79\x96\x71\x56\x9a\x6d\x8e\x56\x85\x9c\x95\xe9\x07\x1b\xc7\x1d\x34\xe1\x34\xc1\xcf\x4e\x61\xb7\x56\x60\x9d\xda\xc0\xc9\xcd\xc2\x9f\xf6\x92\x05\x4a\x5b\x16\x72\x27\x59\x45\xe5\xe6\x6f\xdc\xf8\xe2\xec\xd4\x08\x0a\x54\x79\xe4\x9a\xda\x4a\xdf\x02\x82\x8b\xac\x9d\x11\x13\x92\xef\x49\x76\x23\xc5\xb9\xfa\xca\x58\xc6\x30\x3a\x40\x60\x26\x91\xea\x91\xa1\x03\x27\x4f\x32\x70\x73\xc0\xb6\xc4\xe8\xac\x75\x14\xac\x9f\xfd\x96\xd3\x0a\x3b\xc2\x7b\xf9\x32\x50\x5a\x32\xfe\xac\xc8\x83\x70\x19\xe3\xd3\x1c\x93\x41\x8a\x3b\xaa\xd4\x5a\xc8\x62\x3a\x05\xb9\xe6\xb9\xdc\xd4\xba\xb3\x8c\x4f\x73\xb5\x49\xac\x30\xdc\x0d\xbc\xad\xc1\x8d\xba\xfe\xc6\x94\x3e\x15\xb1\x97\x4f\x7b\x4d\xed\x06\x22\x97\xa5\x44\x5a\x6c\x5c\xa0\xd0\xa9\x7e\xd0\xea\xe4\x72\xa5\x45\x45\x35\xcb\xfd\x7d\x7d\xa9\x0b\xaa\xf1\xfa\x5b\x2d\x51\x29\x26\x38\x44\x36\x75\xa7\xfc\x9d\x7e\x99\xba\xe0\x76\xab\x0c\x60\x8c\x36\x88\xcf\x34\xb5\xe8\x4e\xdf\x73\x01\x0c\x59\x4c\xe3\x0f\xb6\x62\xf4\xf0\x0f\xe6\x4f\xc2\x60\xbc\xad\xef\xa8\xce\x97\x23\x4e\xac\x00\xbf\x0d\x57\x2e\x2c\x54\xb4\x7e\xf4\x6d\xf7\xc4\xb8\x46\xb9\xa0\x39\x9a\x0e\x34\x4c\x75\xb4\x5d\x77\xc6\x90\x1b\xce\x34\xa3\xa5\x6a\x9a\xe3\x75\x70\x93\x6b\xe0\x41\x46\x67\xc0\xae\x4f\xc6\x8a\x30\x08\x06\x53\x6d\xe2\xf5\xe0\x87\x9d\xc8\x6d\xab\xf6\x33\x6e\xa7\xd5\x7c\x67\x59\x83\x57\x3b\x17\xc4\x8b\x6d\x14\x5f\x8c\xc7\x68\x44\x06\x83\xa9\x14\x3d\xfd\x62\xed\x5d\xc0\xba\x35\xeb\x22\xbc\xd2\x92\xc4\xbe\x90\x49\xbb\x39\x3b\x53\xfb\x7c\xa6\xc8\x3f\xb4\x64\xc5\x56\x75\xbd\xff\xdb\xaa\x1c\x16\x62\xb7\x7a\xf6\xd4\x73\xc3\x73\x21\x25\xe6\xda\xde\xd4\x67\x79\xb7\x2e\xda\x2d\x12\xec\xa0\xfd\x0a\x3d\x5b\x8e\xeb\xee\x20\x7a\xda\x41\x1f\x9c\x5a\x90\xd1\x89\xd2\x05\xdd\x2e\xef\xa1\x94\x83\xb7\x2a\xf8\x48\xeb\x5e\x08\x4e\x01\x3b\x02\x98\x6e\xe5\xb6\x89\x93\x71\x09\x5c\x61\x89\xa3\xc2\xe9\x34\xd0\x77\xf9\x4e\xbe\x03\x7e\x09\xb8\x54\x63\xa2\x78\xbb\x58\xc6\x9b\xf5\xf7\x8d\xfd\x3c\x7d\xb7\x40\xbb\x5d\x60\x80\x33\x38\xa4\x86\x41\x82\x16\xae\x87\xf6\xb5\x87\x36\x75\x61\xed\x02\x4a\xda\x4f\xdb\x69\xdb\xa3\x0c\x3e\x5e\xfd\x34\xbc\xda\x25\x55\x4b\xd7\x50\x17\x19\xcc\x5d\x3c\xd2\xfd\xde\xf9\x43\x8a\x6a\x1b\xfc\xf1\x69\xbe\xd1\xd8\x7b\xa6\x9d\xf5\x15\x2e\xe8\xaa\xd4\x33\xa1\x74\x3f\x85\x7d\xe6\xd8\x05\x4f\x46\xc1\x4f\xe8\x08\x3d\x27\x2c\xea\x75\xcf\x63\x2e\x44\x79\x98\x45\xff\xde\x03\xd9\x96\xc7\x4c\x54\x35\x95\xf8\x17\x55\xcb\x4b\x5e\xec\x33\x19\x66\x4a\x52\xd8\xe7\x97\x84\x0d\x84\x6d\xbb\xff\x1f\x00\x00\xff\xff\x85\xb8\x58\x7d\x03\x0e\x00\x00")

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

	info := bindataFileInfo{name: "service.tpl", size: 3587, mode: os.FileMode(420), modTime: time.Unix(1591580659, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vueTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x6d\x6f\x1c\xb7\xf1\x7f\xaf\x4f\x31\x21\x9c\xe0\x2e\xb9\x5d\xc9\xf9\x23\xf8\x23\xab\xdb\x6b\x10\x25\x2d\x82\x16\x4d\x80\x26\xaf\x0c\x23\xa0\x76\x79\x77\xac\xb9\x0f\x25\x79\x27\x29\x87\x03\x8a\x36\x89\xed\xb4\x89\x5c\x34\x15\xd2\x3c\xc0\x76\xe1\xf4\x09\xb0\xe2\x17\x85\x6b\xc7\x4a\xfa\x65\x74\x27\xe7\x55\xbf\x42\x41\x72\xef\x96\xfb\x70\x27\xc9\x51\x82\x1a\x28\x60\x9c\x76\xc9\xe1\x70\x38\xfc\xcd\x8f\xc3\x59\x03\xb4\x9f\x70\x1c\xf8\x11\x89\x09\xc7\x92\x84\xb0\xb9\x03\x82\xe1\xe0\x12\xe1\x8e\xd3\x81\x95\xb6\x24\x51\xca\xb0\x24\x9d\x15\x80\x36\x61\x4e\x80\x79\xa8\x9e\x01\xda\x21\x1d\x82\x60\x89\xf4\x51\x9f\xe0\x90\x70\x64\xda\x8d\x5c\x37\xe1\x11\x78\x34\x66\x34\x26\x3e\x92\x7c\x40\x10\x0c\x9d\x28\x09\x09\xf3\x91\x20\x98\x07\x7d\x25\x82\xe0\x05\x31\xd8\x8c\xa8\x74\x63\x2c\xe9\x90\xb8\x29\x27\x43\x12\x4b\x1f\xc5\x49\x4c\xe6\x1a\x73\x9d\x0e\x95\x24\xca\x9b\x4d\x07\x8d\xd3\x81\x84\x94\x93\x2e\xdd\x76\x68\x90\xc4\x3e\x52\xad\x41\x12\x3b\x66\x2a\x04\x01\x23\x98\xe3\x4d\x46\x66\x56\xb8\x92\xd3\xc8\x36\xc5\xa5\x29\x82\x94\xe1\x80\xf4\x13\x16\x12\xee\xa3\xe9\xf5\xcf\xa7\x57\xbf\x9e\xbc\xfb\xce\x64\xff\x3e\xea\xb4\x57\x67\x33\x59\x56\xad\x16\xcc\x82\x93\x99\xbb\x39\x90\x32\x89\x41\xee\xa4\xc4\x47\x29\xa7\x11\xe6\x3b\x08\x5e\x08\x18\x0d\x2e\xf9\xe8\x67\x99\xc5\xb5\xeb\xe8\x4c\xaf\x7f\xfe\xf0\x8b\x3f\xe9\x79\x8d\x9a\x05\xca\x67\xda\x38\x11\x44\x5a\xfe\xee\x4c\xef\xbd\x73\xf4\xb7\x2f\x8d\x9a\xe9\x67\x37\x0f\x1f\xdc\xad\x55\x56\x5a\xd9\x4a\xa9\x75\xde\xa0\xd0\x53\x33\x2d\x8d\xd2\x84\xcb\x2e\x65\xc4\x1d\x52\x41\x37\x19\xf1\x0d\x06\x78\x32\x88\xc3\xf2\xd2\x03\x86\x85\xf0\x91\xc2\x56\xa6\x08\x75\xda\x74\xd6\x3c\x73\xc1\x20\x65\x09\x0e\x9f\x55\x3b\x41\x3b\x93\x2f\x0e\x26\xef\x7c\xfe\xf0\xe6\x5f\xa7\x37\x0e\xec\x05\x80\xe3\x1c\x67\xdb\xcb\xdb\xca\x36\x54\x36\xc2\x58\x76\x32\x53\xc2\x64\x2b\x56\xc6\xe4\xb6\x5c\xfe\x72\xb9\x2d\x55\x33\x02\x4e\xb0\x9c\xbb\x07\x7c\x30\x0e\x3a\xa9\x55\x25\x7c\xa4\x6c\x20\x50\x67\xfa\xcf\x07\x93\xf7\x6e\xd4\x6c\xa8\x09\x56\xb9\xc3\x88\x8f\x74\x28\x78\x9b\x89\xec\x2b\xf3\x43\x3a\xcc\xe2\xd9\x7a\x24\xcc\x91\x3a\x58\xbc\x10\x4b\xec\x23\xfd\xe2\xaa\x67\x34\xd3\xb2\x45\x43\xd9\xf7\xe0\xfc\xda\xda\x93\x2a\x88\x09\x23\x81\xa4\x49\xec\x04\x7d\x1c\xf7\x88\x82\x71\xd6\xb2\xa1\x1b\x54\xf4\x2b\x8f\xd1\xb8\x37\x53\x97\xbd\x16\x68\x43\xf7\x38\x41\xc2\x06\xd1\x2c\x40\xe6\xaa\x11\xe8\x39\x7d\xf4\xdc\x73\xa8\x00\x48\x7b\x50\x8e\xe0\xd1\x88\xab\x89\xe1\x1c\x6d\x9d\x0b\x12\xe6\xf9\xee\x86\x96\x10\xe3\xf1\xa2\x09\x19\xde\x54\xf4\x34\x1a\xa9\x01\xee\x06\x8e\x08\xdb\xc0\x82\xfc\x14\x47\x64\x3c\xb6\xc9\x68\xc6\x8a\x9a\xff\x1c\x11\x24\xda\x50\xf5\x07\x15\xe2\x51\xa4\x38\xee\x40\xbb\xdd\x06\xdd\xe9\xf2\x64\xcb\x9d\x69\xd7\x53\x1a\xd5\xd0\xe9\x74\xa0\xbd\xaa\xa5\xad\x08\xb4\xb9\xf7\x98\xc5\xc2\x68\x44\xe2\xf0\xd8\x95\x4d\x7f\xff\xfe\xe1\x57\x9f\x3e\xca\x4a\x72\xfc\x0a\xfa\x96\x12\x89\x30\x63\x15\xb4\xa6\x0c\xd3\x1c\xe1\xaf\xc6\x1b\xea\xe1\xe5\x90\xca\xc6\x7c\xfd\xcd\x32\x72\x49\x48\x25\xea\xc0\xf1\x94\x56\x33\x71\xa8\x76\x98\xe7\xdc\xf9\x12\x61\x44\x92\x25\x93\x85\x5a\x20\x23\xf3\x2a\xe7\x9d\xd0\xe3\x79\xb3\x7d\x14\x9a\xb8\x88\x30\xef\xd1\xd8\x61\xa4\x2b\x3d\x78\x76\x2d\xdd\x5e\xcf\x5a\x64\x92\x9a\x06\x54\x43\x0a\x43\x47\xf4\x93\x2d\x1f\x45\x03\x26\x69\xca\xc8\x3c\x7a\x5c\x46\xe2\x9e\xec\x57\x3c\x3d\x5b\xb1\xf6\xf0\x8b\x58\x06\xfd\xd7\xd4\x0f\xea\x4c\xaf\xde\xff\xe6\xf2\xee\xe1\xbf\xf6\xa7\x1f\xde\xaf\x23\x82\xd3\xcf\x59\x76\x72\x3e\xe5\x4b\x99\x37\xcd\x9c\x93\x2b\x37\xbe\xf9\xe3\xad\xca\x9c\x25\x5e\x49\x71\x8f\xaa\xb3\x3e\x89\xc1\xeb\xd3\x90\x38\xea\x70\xa3\x71\x8f\x11\xd5\x35\xcf\x15\x32\xc6\xcb\xa5\x11\x6c\xe2\xe0\x52\xcf\xf0\x21\xc3\x3b\xc9\x40\xfa\x48\x63\x42\xb4\x40\x26\x12\xb3\x96\xca\x00\x86\x2d\x50\x6a\x78\x0b\x62\xb2\x2d\x11\x78\xba\x6b\x46\x39\xfa\x05\x81\xa7\x44\x1c\x83\xa7\x8c\x8b\x68\x44\xa5\xdd\x21\x7c\x74\xe1\xfc\x5a\x0b\x9e\x5b\x6b\x29\x8e\x6b\xc1\xf3\xcf\x3f\x7f\x51\x11\x1d\x7d\x8b\xe4\x1c\x47\xdf\x22\x33\x7a\x33\xfe\xf5\x82\x01\xe7\x24\x96\xb3\xb5\x68\xe5\xea\x59\xb9\x2f\xeb\x9a\x8d\x7e\x0d\xf7\x66\xa3\x2d\x58\xe5\x2b\xee\xac\xcc\x9d\x16\x52\xcc\x92\x1e\x48\x2a\x15\xc4\x0c\xcb\x23\xf0\xb2\xb3\xc3\x15\x3b\x71\x50\x3e\x50\x72\xb6\x5c\x2b\x63\xce\xe4\x66\x59\x2e\x96\x8d\x32\xc9\x98\xa6\x09\x27\x4d\x04\x55\x06\xf8\x88\xd3\x5e\x5f\xce\x9a\x33\x7d\xe7\x0b\xfa\x4e\xc6\xb4\x99\x18\xed\x42\x9c\x48\xd0\xfc\xf7\x8a\x78\xcd\x60\xf9\xc7\x64\xc7\x92\x2b\x24\x4e\x27\x26\x64\x3b\x0d\x9c\xe5\x77\x7a\xba\x8d\x24\x96\x98\xc6\xc2\xcc\xf9\xfa\x4e\x4a\x00\xd1\x58\xa2\xf1\x38\x1e\x44\x9b\x84\x8f\x46\x84\x09\x32\x1e\xab\x64\x30\xa3\xd0\x82\x4b\xea\xd8\x1a\x01\x1e\xc8\xc4\x09\x92\x28\x55\xf0\xf7\x51\xd2\xed\x9e\x28\x3d\xb4\x9d\x51\x22\xec\x1a\x0a\xaf\xa4\x5a\x79\xbe\xdd\x4d\x12\xa9\x22\x32\x0b\x12\x83\x0e\x27\x6b\x2d\xe6\xcc\xc7\x25\x1d\x5d\xcc\x04\x41\x9d\xc9\xee\x1e\x4c\xef\x5e\xa9\xa7\xc5\x63\x53\xd6\x0d\xad\x15\x75\x8e\x6e\xee\xc3\x64\xff\xe3\x3a\xda\xb1\xf3\x8c\x39\xa0\x17\x22\xfc\xe8\x60\xef\xe1\xd7\xbf\xab\x20\x7c\x90\x86\x8f\x80\xf0\x6c\xd4\xa3\x22\xfc\x64\xe8\x7e\x6c\xc0\x6d\x79\xe3\xbb\x02\x77\x11\xc8\xdf\x07\xac\x8b\xc0\x38\x2b\x58\xbf\xa1\xb5\x9e\x21\xac\xed\xb3\xb9\x02\xee\x4d\x75\xa2\xa6\xea\xe7\xb4\x00\xb7\x46\x3e\x3a\x8d\x9f\x90\xc4\x1f\x1f\xa0\x97\xbc\xf2\xfd\x80\xfd\xfb\x81\x7b\x15\x2a\x67\x05\x79\x3b\x93\x3c\x3d\xec\xb3\x24\xa8\x06\xf8\xbb\xd7\x8e\x6e\x7d\xe9\x1d\xde\x7b\xef\xf0\xe0\xc6\xe4\xda\x07\x93\x3b\xef\x4e\x3f\xbc\x3f\xd9\xfd\x68\xf2\xc1\xf5\xe9\xa7\x57\xa7\x7f\xb8\x33\x7d\x7f\xbf\x12\x11\xd5\x22\xc2\x3c\x22\xfe\x6f\xed\xc9\x42\x3c\x98\xda\x00\x70\xd2\x55\x64\xa0\xaf\xe6\xe0\x59\xb7\xd1\x1f\x52\x36\xcf\xd4\x3c\xa5\xd1\x61\x54\xc8\xc2\x14\xea\x47\x35\x22\xf0\xfa\x52\xa6\x0e\x27\xbf\x18\x10\x25\xf3\x4a\x64\xea\x05\x9e\x06\x8b\xd1\xee\x23\xe3\x6e\x08\x39\xee\x01\x0e\x4c\x9c\xd9\xbb\xb6\xa0\x7c\x61\x2a\x06\xb9\x98\x82\x43\x2e\x68\x64\xde\x7c\x53\xaa\xdc\xb5\xa3\xbc\xb4\x77\xf9\xf0\xc1\xdd\xe9\x6f\xf6\x26\x57\xee\x4c\x6f\xdf\x9a\xdc\x7a\xfb\xdf\x07\xbf\x9d\x5e\xd9\x2b\xc4\x4c\xd4\x39\xfa\xd5\xfd\xc9\xe5\x07\xc6\xbd\xed\x55\x1b\xa0\xd6\x56\x2d\x9b\x8f\xa6\x28\x43\xa5\x7a\xec\x4c\x76\xff\xfe\xf0\xd7\x5f\x19\x7d\xee\x36\x13\xdb\xc6\x8e\x82\x32\xbd\xf3\x46\xc1\xbc\x49\xdd\x60\xbf\x35\xba\xab\xdb\x7e\x56\xe8\x36\xa5\xc6\x37\xb2\x8d\x58\x86\xef\xfc\x2a\x6e\x03\xdc\xe0\x7b\x25\x6b\x34\xa5\x50\xfb\xd2\xb8\xd2\x16\x01\xa7\xa9\xe6\x0d\xb3\x0a\x78\x1a\xb0\x00\x9c\x52\xe8\xf2\x24\x02\xf4\xc2\x2a\x4e\xe9\xea\x68\xe4\x66\x04\xb4\xbe\x02\x40\x74\x3d\x0a\x42\xd2\xc5\x03\x26\x61\xa4\x27\x09\xb1\xc4\x8d\x66\xf6\x02\xc0\x89\x1c\xf0\x78\xfe\x0a\x50\xb9\xb4\x79\x70\xe1\x62\x6b\xde\xad\x2f\x1e\x9e\x25\x0f\x90\x95\x5b\x3c\xe3\xca\x96\xd5\xa3\xe6\x2a\x0e\x07\x73\xa3\xf2\x60\xcd\x6e\x53\xf7\x18\x0f\xce\xdb\x4d\x49\xb7\x2b\x88\x2c\xc9\xe9\xcb\x94\x07\xe7\xd7\xe6\x6d\xe3\xbc\x3b\xaf\x45\x7a\x30\xb2\xda\x4d\x66\x5a\x34\xb9\x2a\x05\x90\x41\x22\x5b\x46\xdd\x0c\x26\x19\x38\x0b\x4d\x39\xcf\x9e\x85\xb6\x1c\xd7\x25\x6d\x19\xf1\x94\xb7\x60\x91\xc6\xec\x69\xbc\xbe\x62\x4d\x10\x11\xd9\x4f\x42\x91\x6b\xee\x11\x59\xc2\x10\x80\xec\x53\xe1\x16\x8a\x6f\x59\xd9\x71\x7d\x2e\x82\x53\x6a\x99\xe0\xfe\x84\x0a\xd9\x28\x9c\xff\xaf\x6e\xfe\x9c\x04\xd2\xc5\x42\xd0\x5e\x5c\xec\x02\x18\xcd\x01\x61\x4d\x65\x9a\x5a\x33\x58\xd8\x46\xa8\x16\x28\xb8\x71\x6e\x66\x8e\x93\x42\x6f\x73\xa5\xfe\xd9\x95\x7d\x12\x37\xb8\x48\xc1\xef\x14\xbc\x5b\x58\xb5\x46\x35\xf8\xc0\x45\xaa\xab\x99\xa6\x61\x7d\x91\xb8\x12\xb1\xa5\xd5\xcf\x42\xe1\xdc\xa3\x7a\xbf\x6c\xb9\x71\xc1\xd2\x40\x01\xaa\x41\x38\xd7\xa6\x8e\x9b\x33\xc9\xb9\x1f\xf2\xc2\x40\x63\x88\xd9\x82\x0d\xd4\x05\x86\x21\x66\x75\x7d\xc6\xe3\x6a\x6f\xcb\xbe\x7e\x1a\x94\x4a\x70\xe0\x7c\x73\xbd\x38\x72\x8e\x97\x8a\x39\x79\x95\x63\x89\x39\x46\xbd\x0f\x43\xdb\x9d\xc7\x28\xd6\xdf\x26\x4c\xdc\xab\x7d\xae\x42\xb5\xfe\xb6\xba\x5e\x2b\xa4\x33\x62\x1f\x46\xe3\xfa\x69\x72\x38\x55\xa7\xc9\xfb\x0a\x0a\xb2\x5e\xf3\x3d\xa6\xce\x2f\x59\xc7\x92\xfd\x39\xbf\x74\x77\xd6\x4e\xec\x29\x73\xc1\x2e\x4c\xc5\x88\x04\x2b\x88\xd5\xf8\x73\xd9\x7b\xa3\x48\xfb\xc1\x25\x4f\x07\x79\x81\xde\xc9\xb6\xf4\x20\xab\x07\x1f\xde\xbb\xed\xba\x2e\xb2\xfb\x45\x4a\xe3\x98\x70\x0f\xe6\xe9\xcb\xac\x5a\x6f\x4b\xe5\x15\x39\x0f\x10\xef\x6d\xe2\xc6\x5a\x0b\xb2\x7f\xee\xff\x37\x51\xce\x59\xcd\x85\x0c\x93\x2d\xad\xbc\x95\xa7\x89\xed\x0a\x8e\x6a\x62\xb4\xe2\x5b\x38\x36\x2e\x2b\x16\x34\x9a\x55\x03\x32\xb7\xb8\x01\x4b\x04\x29\x6b\xaf\xec\x63\x56\x9f\xe6\xc9\x56\x05\x35\xe7\x82\x24\xee\x52\x1e\x35\xd0\xf4\xf6\x2d\xb3\x2f\x93\x3b\xef\x9a\x9a\xaa\x6a\xf9\xec\xa6\xc9\x92\x5b\x30\xfd\xe8\x8b\xc9\xb5\x3f\x1f\x3d\xf8\xcb\xd1\x83\xdb\x3f\x40\x2d\xc8\xb2\x6b\xd4\x2a\x98\x96\xa9\x7b\x51\x27\x36\xaf\x9b\xed\x3e\xba\xb9\x3f\xd9\xff\xb8\xb0\x85\x01\x8e\x03\xc2\x0a\x52\x93\xdd\xbd\xe9\xdd\x2b\x05\x29\x95\x4f\x79\x80\xb6\x30\x8f\x15\x06\x56\xea\x1d\xb8\xd0\x49\xc5\x2d\xd7\xb2\xb9\x27\xdc\xd1\xc8\xcd\x2f\x91\xe6\x7a\x56\xb8\xa5\x35\xcb\x63\x17\xc3\x21\xf7\x66\x9c\x48\xda\xdd\x71\xc5\x20\x08\x88\x10\x0d\x94\x39\xf2\xca\xb5\xc9\x7b\xd7\x51\x09\x20\xb0\x14\x24\x95\x75\xc2\x32\x12\xaf\x7a\xe5\x58\xba\x37\x65\x86\x42\x70\x97\x62\x24\x93\xd0\x36\x16\xeb\x37\xcb\x3d\xd7\x82\xf2\x90\xd3\x84\x55\x7d\x51\xa5\x26\xb6\x2a\xce\x9e\x7e\xf2\x8f\xe9\xde\x9d\x7a\x67\x3f\x6a\x34\x56\xdc\x66\x7f\x5d\x52\x01\x65\xad\xa0\xbc\x6a\xf0\x4b\x69\xcb\x68\xdc\x02\x35\xa6\xc6\xb4\xca\xaa\xed\xfc\xc8\xe2\xfe\xc2\xd7\xcd\xda\x83\xb1\x92\xa2\x17\x8f\xc7\x9c\xdc\x4b\xdf\x53\xaa\x27\x8a\xc5\x0d\xd6\x97\x96\x6f\x7e\x79\xf5\xf0\xde\xed\xa3\x8f\xdf\x9e\x7e\x76\xf3\xe8\x93\xc7\x92\x18\xf4\xea\x0a\x4b\x3f\x0e\x10\x99\x96\x1a\x3c\x2c\xf2\xe0\x63\x7b\x50\x2a\xc3\x69\x28\xc0\x87\x0b\x17\xf3\xd6\x6e\xc2\xa1\xa1\xbb\xe2\x90\x6c\xeb\x2c\x22\x7b\x6c\x2f\x80\x5d\xf6\x39\x2f\x13\x7b\xe6\x99\x66\x19\x0e\x42\x82\xae\xd3\xf9\x0b\x14\x5c\xd0\x03\x2f\xda\x5b\x43\x43\xe1\xa6\x03\xd1\x6f\xa8\x81\xc7\x33\xf8\x7a\xe5\xf6\x52\xe1\x38\x7b\xfb\x46\x70\x9c\x46\x4f\x39\xa6\x06\x64\x8b\xd9\xec\x54\x27\xc2\x7f\x51\xca\x50\xfa\xc0\x5b\x25\x87\x52\x9d\xb3\x2e\x85\xad\xad\x17\xd6\xf2\xda\x82\x89\xfe\x17\x43\x8f\x4f\x0c\x99\xdd\x83\x93\xc6\x50\xab\x16\x45\xdf\x2a\xb2\xcc\x07\x8e\xd3\x46\xd6\x32\xac\x9e\xf2\x46\xfd\x1d\x84\xa1\xf9\x4f\x63\xe5\x24\xcd\x9d\x35\x57\xe4\x4d\xd1\x58\x6f\xad\x3d\x66\x88\xb9\x3e\x33\xc1\xb7\xa6\x43\x38\x4d\x19\x0d\xf4\x47\xff\xd5\x61\x1c\xba\x49\x4a\xe2\xed\x88\xa9\x8d\xc0\x52\x38\x49\xb7\x4b\x03\x12\x26\xc1\x20\x22\xb1\x74\x45\xca\x09\x0e\x45\x9f\x10\x19\x31\x57\xff\x45\xb9\xf5\xb4\x0b\x76\x55\x46\x63\x4b\xd7\x52\xf5\xb4\x4f\xf8\xbe\x99\xff\xa9\xa7\x96\x0b\x15\x6c\xda\x4a\x67\x46\xe8\x62\x70\x1e\x6a\xc5\x30\x30\x68\x88\x88\x10\xb8\x47\x5c\xc2\x79\xc2\x1b\xc8\x2e\x25\xeb\x4a\xf2\x8d\x83\xc9\xc1\x6e\x11\x16\xa6\xba\x59\x07\xf0\x93\x50\x4d\x21\xf0\x17\x85\x86\xb5\x1f\x7a\xa5\xdf\xee\xe8\xd0\xff\xf9\x70\x09\xc0\x17\x97\xb0\xcf\x18\xc4\x45\xff\x34\xeb\xc0\xbc\xfc\x34\xab\x00\xd7\xae\x91\xd7\x64\xa2\x9c\x74\x55\x8e\xac\x7a\x5d\x23\x5a\x03\xfe\xfc\x6b\x4b\x43\xf9\xa0\x05\x8c\x0a\x59\xd1\x55\xf3\xed\x45\xf1\xb4\xfa\xeb\xa6\x49\xda\x68\x5e\xac\x28\x8e\x93\x98\x34\x9a\x23\x83\x8f\x71\xcb\xd4\xe3\xcd\xcd\xdf\x36\xb6\xb4\x60\x23\xa6\x06\x8d\xd7\x57\xda\xab\xb3\x1a\xfd\x0a\xfc\x27\x00\x00\xff\xff\xea\x47\x7d\x3c\xfe\x2c\x00\x00")

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

	info := bindataFileInfo{name: "vue.tpl", size: 11518, mode: os.FileMode(420), modTime: time.Unix(1577779184, 0)}
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
	"component.tpl":   componentTpl,
	"controllers.tpl": controllersTpl,
	"dao.tpl":         daoTpl,
	"js.tpl":          jsTpl,
	"markdown.tpl":    markdownTpl,
	"models.tpl":      modelsTpl,
	"service.tpl":     serviceTpl,
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
	"component.tpl":   &bintree{componentTpl, map[string]*bintree{}},
	"controllers.tpl": &bintree{controllersTpl, map[string]*bintree{}},
	"dao.tpl":         &bintree{daoTpl, map[string]*bintree{}},
	"js.tpl":          &bintree{jsTpl, map[string]*bintree{}},
	"markdown.tpl":    &bintree{markdownTpl, map[string]*bintree{}},
	"models.tpl":      &bintree{modelsTpl, map[string]*bintree{}},
	"service.tpl":     &bintree{serviceTpl, map[string]*bintree{}},
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
