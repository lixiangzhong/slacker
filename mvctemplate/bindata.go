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

var _controllersTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\xdf\x6f\xdb\xb6\x13\x7f\x16\xff\x8a\xfb\xea\xe1\x0b\x29\xd0\xa4\x74\x28\xfa\xe0\xc2\x0f\x8d\x13\xb4\xdd\x82\xd4\x5b\x92\xa7\xa2\x18\x18\xf2\xec\x10\x91\x28\x8f\xa4\xd2\x19\x9a\xfe\xf7\x81\x94\xd2\x5a\xb6\x28\x7b\x75\xb0\xe8\x49\x96\x78\x77\x9f\x1f\xbc\xa3\x9c\x65\xef\x51\xa2\xa2\x06\x39\xdc\xad\x41\xe7\x94\x3d\xa0\x02\x92\x65\x4b\x21\x41\x95\x95\x41\x45\xb2\x13\x42\x57\x22\x7d\x7f\x71\x13\x85\x59\x5d\xa7\x97\xe5\x57\x54\x57\xb4\xc0\xa6\xc9\x26\x82\x87\x09\x2b\xa5\x51\x65\x9e\xa3\xd2\xe9\x0d\x7d\xc0\xba\x4e\x67\xb4\xc0\x7c\x46\x35\xb6\xeb\x62\x6f\x86\x7e\xf4\xa5\xd0\xc6\x17\x3d\xff\x74\xbd\x2f\x7c\xa6\x90\x1a\x6f\xf9\xf9\xed\x21\x04\x6e\x57\x7c\x2c\xc7\xbb\x9b\xd9\x87\x03\xb2\xcc\xa9\x61\xf7\xbe\x24\xe7\x17\x97\x17\x37\x17\x07\x64\x39\xc7\x1c\x07\xb1\xc0\x08\x98\x7e\x8a\x33\x0b\xe4\x47\xd0\x0c\xa4\xf1\xc2\xf1\xca\x3b\x90\xc4\xab\x2f\x90\x93\x8c\x90\x15\x65\x0f\x74\x89\xb0\x11\x47\x88\x28\x56\xa5\x32\x11\x09\x20\xac\xeb\x70\x59\x6a\xc5\xb2\xa2\xe4\x98\xeb\xf0\x6f\x48\x3f\xba\xb7\x97\xe2\x4e\x51\xb5\x6e\x1a\x0b\xa1\xab\xde\x0b\x40\xa5\x58\xc9\x71\x5f\x44\xb8\x14\xe6\xbe\xba\x4b\x59\x59\xd8\x26\xf8\x69\x59\x4a\xc1\xec\x9d\x7d\x27\xd1\x64\xf7\xc6\xac\xec\xbd\x36\x8a\x95\xf2\x31\x24\x31\x21\x8b\x4a\x32\x00\x18\xde\xbe\x11\x83\x93\xa5\x90\xe9\xac\x94\x06\xff\x32\x31\x40\x4d\x00\x00\x1e\xa9\x82\xbe\x5c\x5b\x3f\xd3\xdd\x5c\x24\x28\x17\x0b\x8d\x26\xf9\x63\x32\xed\x00\xa4\x73\xaa\x34\xde\x0a\x69\x22\x96\x9e\xe3\x82\x56\xb9\xf9\xad\x42\xb5\x8e\xc2\x76\x6d\x98\x84\xa7\x61\x9c\xc0\xab\xd3\x04\xde\xbc\x8e\x09\xb8\xea\xb9\x28\xc4\x81\x69\xdc\xd2\x2e\xcb\xab\xd3\xc4\xe5\x08\x58\x7a\x7d\x5f\x56\x39\x3f\x13\x92\xb7\xeb\xfe\xdf\x47\x1f\xdb\x32\x9c\x1a\x9a\x98\xd2\xd0\x3c\x41\xa5\x26\x53\xb8\x46\xf5\x28\x18\x7a\x3a\x3d\xea\xd8\xb5\xe0\xb6\xf3\x05\x41\xf0\xcb\xf5\xa7\xab\x88\x25\x60\xf5\xfc\x50\x93\x20\x08\x6d\x85\x70\xd2\x55\xb2\x0f\x5c\xb5\x70\x02\x6d\x55\x12\x34\xb6\x72\x4c\x9a\x6f\x2e\x0d\x8f\xa8\x31\x97\x6c\xee\x43\xbc\x19\x12\xc5\xc6\x3a\x29\x40\xf0\x04\x50\x29\x98\x4c\xa1\xa7\xf9\x47\x27\xf9\x9c\x2a\x5a\x44\xa1\xe0\xdf\x45\x0e\xc4\xc2\x05\xfc\x6f\x0a\x52\xe4\x60\xe9\x3e\xf1\x97\xc2\x09\x6a\x37\x74\x7a\x85\x5f\xa3\xa7\xfb\x33\xca\x7f\xc7\x3f\x2b\xd4\xc6\xb1\xb6\x9a\x29\x34\x95\x92\x24\x68\x80\x04\x16\x8c\x45\x3e\x57\xa2\xa0\x6a\xfd\x2b\xae\x67\x65\x5e\x15\x72\x9b\xca\x54\x70\x0b\xd9\x16\xff\x6e\x99\x47\xb7\x96\x21\x6c\xba\xe3\xbc\xe8\xab\xee\x1b\xcd\xcf\x20\xbb\xd3\x76\xf1\x24\xed\xa6\x05\x0e\x4f\x8b\xef\x6d\x5f\x49\x70\xd7\x0f\xe8\x09\xdd\xf5\x4d\x55\xf7\xa4\xab\xfd\xa4\x94\x97\x6c\xb7\x1b\xc6\x94\xf2\xcd\xc7\x21\xa1\x82\x67\xd9\x54\xc7\xec\xa9\x23\x06\xd9\x98\x65\x5b\x9d\xff\xf6\x39\xdb\x80\x04\xbb\xe0\xf6\xf5\x03\x4c\x41\x70\x12\xf4\xfb\xc1\xeb\xd4\xee\xdc\xda\xc2\xbb\xd1\x17\x9e\xa3\x79\xd0\xed\x97\xf6\x3a\xb0\x46\x57\x8e\xf5\x42\x60\xce\x35\x4c\xa1\xa0\x0f\x18\x15\x74\xf5\x59\x1b\x25\xe4\xf2\x8b\x90\x06\xd5\x82\x32\xac\x9b\x78\xdc\xe3\xcd\x44\x3b\x0e\x1f\x39\xe7\xfa\x46\xf9\x34\x16\x3c\xe9\x81\xd8\x29\xbb\xe1\x93\xef\xdb\xc7\x37\xbf\x8e\xf6\x0a\x8e\x9f\xf6\x7d\x15\xbc\x0c\x04\xf7\x6d\xd1\xec\xa4\x63\x3f\xfa\xf9\xb7\x23\x81\xdb\xaa\xba\xa3\x2e\xe4\x52\xff\x2c\xa4\x79\xf3\x5a\x47\x2c\x75\x87\xe2\x3b\xa5\xe8\xba\x65\x1f\xb7\x4b\xa7\x40\x57\x2b\x94\x3c\x12\x5c\x27\xfb\xa2\x3e\x7f\x09\xe3\x38\x4d\xd3\xb8\xa3\xb8\x31\x78\xc7\x71\x8a\x3d\x1e\x8f\x7d\x2b\x0f\x91\xfc\xef\xda\xe1\xd8\x89\xf7\x92\x76\xf8\xbb\x4f\x27\xb0\xd5\x7f\xfb\xcc\x39\xf8\x70\xec\xdc\xf9\xb7\xa7\x12\x3c\xeb\xb1\xf4\x12\x26\x0c\x7b\xe0\xd5\xcd\xb6\xc4\xc0\xa4\x89\x49\x63\xff\x8e\xfd\x13\x00\x00\xff\xff\x36\x9c\x2a\x37\x1f\x10\x00\x00")

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

	info := bindataFileInfo{name: "controllers.tpl", size: 4127, mode: os.FileMode(420), modTime: time.Unix(1575441929, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _daoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\x5f\x8b\xe3\x36\x14\xc5\x9f\xad\x4f\x71\xeb\x87\x20\x0f\x46\xee\x42\xe9\xc3\x42\x28\x6c\x66\x5a\x86\x66\x96\xcd\x66\x96\x3e\x2c\x4b\x51\xac\x9b\x44\xc4\xb6\x32\x92\xbc\xd9\xa0\xea\xbb\x17\xc9\xb1\x9b\x4c\x32\x24\x7d\xca\x78\xb8\x7f\xce\xf9\xe9\x48\x45\xf1\x07\x36\xa8\xb9\x45\x01\x8b\x3d\x98\x8a\x97\x1b\xd4\x64\xcb\xcb\x0d\x5f\x21\x08\xae\x08\x91\xf5\x56\x69\x0b\x94\x14\x85\x79\x49\xd2\x95\xb4\xeb\x76\xc1\x4a\x55\x17\x4f\xdc\x58\xd4\xb5\x6c\x84\x29\xcc\x4b\x2b\xb5\xc6\x2a\x25\xa9\x73\xe9\x4a\x19\x5d\x16\xb5\x12\x58\x99\xf4\x1f\x60\x8f\x71\xc6\x54\x2e\x34\xd7\x7b\xef\x0b\xe7\xd8\x47\x5e\xa3\xf7\x29\x90\x8c\x90\x65\xdb\x94\x40\x05\xdc\xdd\x73\x95\xc1\x54\x1a\xeb\x1c\x9b\xf0\x1a\xab\x09\x37\xd8\x55\xd2\xdd\x1a\x35\x02\x30\xc6\x42\x39\xbd\x5b\x29\x5d\xb3\xfb\x0f\xd9\xf0\x07\xd0\xaf\xdf\x9c\x63\x53\xb5\x43\xdd\xf5\xb0\xf3\x31\xb9\x6c\x6c\x8e\x5a\x2b\x9d\x81\x23\x00\x00\xdf\xb9\x06\xc1\x2d\x87\x71\xcd\x37\x78\xd3\x8c\x9f\xb3\xa1\xd3\x2a\xcb\x2b\x90\x8d\x25\x09\x6a\x0d\xef\xc7\x82\x45\x3d\x4f\xc1\x3b\xbd\x3a\xca\xf9\x8c\x4d\x54\xdb\x58\x3a\x8a\x93\x32\x36\x2f\xd5\x16\x4d\xe7\x96\x31\x96\xb1\xdf\x65\x23\xe8\x28\x28\xcc\xd8\x43\x10\x1e\x77\x6b\xb4\xad\x6e\xa2\xf0\x3c\x76\x06\x53\xc4\x13\xe2\x1c\x7b\x42\xbb\x56\xe2\x99\x6f\xd0\xfb\xd7\x70\x27\x1a\xb9\xc5\x0b\x78\x23\x82\xbb\xab\x82\x33\x88\xf0\x1c\x81\xe4\x44\x47\xe7\xba\x9b\x4e\x8f\xc5\x9e\x29\xf8\xb2\x15\x97\x15\x9c\xee\x86\x5b\xa5\x80\x23\x41\x49\x51\xc0\xdf\x79\x77\x04\x20\x98\x58\xc4\x80\x89\x87\x1f\x58\xd2\xb4\x8d\x2b\x61\x48\x1d\x18\xb4\xfd\x97\x98\xcf\xa6\xde\x43\x97\x2e\xe7\xd8\x27\x2d\x6b\xae\xf7\x7f\xe2\x7e\xa2\xaa\xb6\x6e\x58\xf7\xd3\x35\x8e\xdf\x5f\xab\x48\xf3\x53\xdd\x19\x1c\x63\xea\x39\xcd\xf9\x77\xa4\xa3\x57\x95\x6f\x01\xfb\xc4\x6d\xb9\xbe\xc0\x4b\x8a\x90\xbb\x5f\x7f\xc9\x0f\xf6\x6a\xbe\xfd\x6a\xac\x96\xcd\xea\x9b\x6c\x2c\xea\x25\x2f\xd1\x1d\x51\x82\x20\x25\x19\xe2\x7e\x95\x6f\xa8\x0e\x95\xec\xa2\xe7\xd3\xda\xb1\x14\x40\x92\xd3\x30\x74\x57\xa0\xcb\x42\x77\xe8\x86\x76\x52\x8f\xac\xfe\x97\xd7\x7b\xac\xd0\xc6\xc4\x86\x7f\xca\x25\xb0\x47\xf3\xc5\xa0\x7e\xe6\x8b\x2a\x9c\xd9\x6b\x2c\x21\xdf\xe7\xa2\x3f\xec\xc3\xef\xff\x4c\x33\x38\x20\x83\xd5\xb0\xb3\xe1\x35\x5e\x32\x0a\xde\xc3\x18\x3a\xc8\x86\x3d\xab\x38\x9d\xde\xd8\x99\x01\xf4\x8f\x44\x8f\xe8\xaf\x90\x3a\x9a\x5e\xe8\x1d\x22\x05\xde\x8f\x7f\x4b\xf3\xf4\xc6\x25\x69\x78\x30\xb4\xb1\x27\x77\xb0\x3f\x18\x38\xbc\x11\xa7\x20\xcf\xb1\x3c\x9a\x87\x1f\xd2\x58\xda\x1e\x76\x1d\x1c\x67\xb0\x50\xaa\x02\x47\x92\xf9\x6c\x1a\x5c\xa4\x06\x2b\x2c\x2d\xbc\x83\xa5\x56\xf5\xd1\xfd\x1a\x6e\xd3\x9b\xb6\x82\x2b\xa8\x64\x2d\x2d\xbc\x4b\x49\x4c\x25\x86\x9d\x71\xc5\x11\x26\xb1\x60\xb3\x16\xf5\xfe\xb3\xda\xd1\xf9\x6c\x9a\xf7\x92\xc2\x43\xc9\x1b\x3a\x8a\x4d\x19\x49\xe4\x32\x98\x83\x9f\xc6\xd0\xc8\x0a\x46\xa3\xfe\xcb\xbc\x54\x01\xc2\x47\xf5\x59\xed\x4c\x90\xde\xb3\xb0\xba\x45\x92\xf8\x81\x4d\x1c\xd4\x3d\xa0\xd8\x08\xef\xff\x0d\x00\x00\xff\xff\x11\x78\x7a\x02\x1b\x07\x00\x00")

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

	info := bindataFileInfo{name: "dao.tpl", size: 1819, mode: os.FileMode(420), modTime: time.Unix(1575441929, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x4f\x8f\xda\x3c\x10\xc6\xef\xf9\x14\x23\xeb\x3d\x38\x6f\x77\x93\x3b\xab\xa8\xed\xfe\x69\x55\x69\x55\xa1\x16\xda\x23\x0c\xc9\xa4\x58\x24\x71\x6a\x4f\x96\x22\x94\xef\x5e\xd9\x09\xdd\xed\x12\xa2\xc2\x05\x64\xcf\xcc\x33\x3f\x3f\x83\x4d\x1c\x03\x7c\xa4\x8a\x0c\x32\x65\xb0\xda\x81\x2d\x30\xdd\x90\x09\x54\x59\x6b\xc3\xb0\x66\xae\x21\x37\xba\x04\x11\xc5\x6e\x21\x6e\x82\x43\xcc\xb2\x36\xd4\x07\xdf\xc5\x7e\xf5\x22\xfa\x3f\xa0\x85\x54\xeb\x8d\x7a\xce\x69\x58\x15\x36\xee\x36\x5d\x2a\xfd\xf2\xa9\x79\x53\xa5\xac\x74\x05\x33\xdc\x90\x54\x59\x08\xfb\x00\xc0\x10\x37\xa6\xf2\x04\xd1\x0f\x62\xb9\x8c\xb1\x56\xf1\x7e\x1f\x3d\xea\x2d\x99\xcf\x58\x52\xdb\xc6\xff\xed\x55\xd6\x2e\xc3\x9b\xa0\x3d\x56\x7b\x54\x96\xe5\xcf\x86\xcc\x6e\x58\x50\x0c\x08\x8a\x2b\x9f\x0a\x50\xa3\xc1\xd2\x4e\xc0\xd7\x07\x00\xed\x70\x8f\x3b\x43\xc8\x24\x57\x3a\x1b\x68\x52\x6b\x7b\xb2\x8b\xaf\x18\x94\x9c\xd7\x99\x93\x54\x59\x9f\x74\x2c\xdb\x8c\x9a\x31\xa6\x3d\x45\x4e\xd7\x63\xd2\x3e\x7e\xa1\xf8\x3d\x15\xc4\xc3\xf3\xcb\xba\xd0\xd9\x23\xbc\x75\x3c\xbd\xee\x89\x49\x8e\x48\x2f\xcf\x1c\xa6\xef\xd6\x59\xe4\x33\xcf\x77\xa9\xf7\xe7\x92\xbe\xfd\xd8\xc7\x1b\x9f\x98\xfc\x65\x6d\x3f\xf9\x8b\x2a\x73\x55\x50\xd7\xea\x09\x0d\xe4\xda\x94\x90\x40\x45\x5b\xf8\xa0\x4d\x79\x8f\x8c\x32\xbc\x09\xc0\x07\x22\xac\x6b\xaa\x32\x29\x5c\x8d\xb8\x02\xf7\xd5\x7d\x46\x15\x96\xe4\xf3\x8e\x2e\x40\x07\xdc\xbd\x0a\xc7\xdc\x4e\xf6\xc0\xbd\x26\xcc\xc8\xd8\x49\xbf\x04\x10\x77\xba\x62\xaa\xf8\x7a\xb6\xab\x49\x4c\x40\x94\x4d\xc1\xaa\x46\xc3\xb1\xab\xbb\xce\x90\x51\xf8\xdc\xf6\xf4\x31\x1f\xfc\x5a\x3e\x1f\xb1\x31\x05\x24\x9d\x51\x46\xa7\x64\x6d\x44\xd5\x53\xf4\x6d\xfe\xb0\x78\x3f\x9d\x2e\x6e\xd1\xd2\xfc\xcb\x23\xbc\xf1\x19\xdd\xed\xed\x24\x5f\xc1\xbf\xc5\xd4\x15\x2f\x58\x6f\xa8\x4a\x44\x5f\xe0\x9f\xc1\xc8\x32\x32\x45\x3e\xe2\x3c\x29\x74\x8a\x8e\x25\x5a\x1b\xca\x21\x71\x04\x83\xa8\xdf\x69\xf5\x55\xa7\x1b\x3a\xd0\xaa\x1c\xe4\x18\x63\x92\x80\x10\x61\x6f\x57\x6f\xbc\x3c\x78\xb7\xb5\x93\x38\x3e\x60\xbd\x64\xd0\x96\xff\xec\x76\xe7\xdb\xd2\xca\xfa\xbe\xff\x74\x44\xe8\x9f\x75\xf7\x88\xce\x5c\x44\x86\x7e\xdf\x8f\xbf\x0d\x5e\x91\x8c\xf0\x47\x86\xea\x02\x53\x92\xc2\xff\xb3\x5c\x39\x66\x11\xfe\xe5\xfc\x79\x64\x43\x5c\xfe\x47\xf1\x3b\x00\x00\xff\xff\x12\x70\xb9\x58\xe7\x06\x00\x00")

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

	info := bindataFileInfo{name: "js.tpl", size: 1767, mode: os.FileMode(420), modTime: time.Unix(1569289381, 0)}
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

	info := bindataFileInfo{name: "models.tpl", size: 719, mode: os.FileMode(420), modTime: time.Unix(1563850859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serviceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x4d\x6f\xe3\x36\x10\x3d\x8b\xbf\x62\xaa\x93\xb4\x10\xe8\x4b\xd1\x43\x0b\xa3\x48\x9d\xb4\x0d\x9a\xb6\x01\x92\x6d\x0f\x41\xb0\x18\x4b\xe3\x98\x8d\x44\xaa\x24\x1d\xc7\xe0\xea\xbf\x2f\x48\x5a\xf1\x97\xec\x38\x27\x25\xd6\xcc\x7b\x6f\xe6\x3d\xd2\x1e\x8d\x7e\x23\x49\x1a\x2d\x55\x30\x5d\x81\xa9\xb1\x7c\x26\xcd\x5a\x2c\x9f\xf1\x89\xc0\x90\x7e\x11\x25\x31\x26\x9a\x56\x69\x0b\x19\x4b\xd2\x27\x61\xe7\x8b\x29\x2f\x55\x33\xfa\xaf\x51\x42\x2b\x39\x32\xff\xd7\xaf\x29\x4b\x52\xe7\xd2\x27\x65\x74\x39\x6a\x54\x45\xb5\x49\xbf\x02\xbf\x0e\x8d\x37\x62\xaa\x51\xaf\xba\x6e\xe4\x1c\xff\x0b\x1b\xea\xba\x9d\x7a\x53\xaa\x96\x06\xca\x53\x96\x33\xc6\x66\x0b\x59\x42\x66\xe0\xd3\x5d\x94\x93\xdf\xe3\x33\x39\xc7\x27\xd8\x50\x3d\x41\x43\x11\x31\xab\xd0\x22\x7c\x72\x8e\xdf\xa8\x25\xe9\xf8\x21\x3f\xac\xcb\x49\x6b\xa5\x1d\x30\x00\xd2\x1a\x7e\x1c\x83\xe1\x15\x2a\x7e\x02\x35\x07\x96\x68\xb2\x0b\x2d\x21\x36\xb1\x6e\x40\xd6\x8d\x30\x76\x00\x40\xcd\x66\x86\x6c\x51\x8b\x46\x58\x58\x08\x69\x7f\xf8\xbe\x30\x84\xba\x9c\xc3\xfb\x5a\xb3\x87\xc7\x77\x8b\x0a\x21\x6d\x11\x86\xca\x1d\xb0\xe4\x05\x35\x2c\xe7\xa4\x09\xc2\x5a\xf9\xbf\xfe\x6f\xc3\x92\xf0\xd9\x18\xdb\x96\x64\x95\x85\x7f\x8a\x58\xf0\x77\x10\x78\xe3\xf5\xed\x88\xcd\x73\x96\x38\x27\x66\x30\x51\xd2\xa2\x90\x06\xf8\xdd\x52\xd8\x72\xee\xa9\x21\x35\x16\x2d\xa5\x5d\x07\x2c\x39\x8e\x1d\xc8\xb3\x58\xfb\xdd\xf8\xe7\xb4\xd8\x9b\xe6\xce\xbf\xb8\xa4\x3a\xf7\x2b\x76\x8e\x64\xd5\x75\xa7\xa4\xea\x8a\x74\x96\x3a\xc7\x6f\xb5\x68\x50\xaf\xfe\xa0\xd5\x44\xd5\x8b\x46\xf2\xf8\x88\xb0\x50\x91\x29\xd3\x80\xe9\xed\x2b\xac\xb2\x58\x17\x3b\x6e\x1f\x31\x2b\x90\x71\xce\xb7\x1c\xdf\x42\x38\xee\xfd\x44\x13\xda\xa1\xf8\xec\xce\xfb\x91\x78\xc6\xdd\xf3\x6b\xf3\xd9\x90\xbe\xc7\x69\x4d\x7e\x35\xc9\x21\x80\x7f\x2f\xb1\xa1\x7e\x11\xbb\x78\x30\x06\x63\xb5\x90\x4f\x86\xdf\xab\xd0\x99\x7d\x14\x22\x1f\xa4\xbd\x45\x63\x96\x4a\x57\xc7\x69\xf9\x95\x2c\xf5\xaa\xb5\x7d\xe5\x00\xf1\x69\x10\x4f\x2c\x66\x6b\xc7\x0e\x57\x75\x6d\xae\x5e\x85\xb1\x1f\x9f\x07\x1c\x4b\x92\x9d\x13\x5d\xaa\x8a\xf8\x45\xad\x09\xab\x55\x00\x65\x49\xd2\x6d\x25\x72\x9b\x02\xbe\x02\xbf\x58\x58\xd5\xa0\x15\x65\x34\xfe\x73\x5b\xa1\xa5\xab\xd7\x56\x93\x31\x42\x49\x7f\x2a\xf6\xae\x97\x33\x13\x72\xde\x55\x13\xf9\xde\x8f\xdb\x07\xd2\x96\xc0\xd1\x19\x0f\xa7\x63\x09\xc0\x5a\x65\x9c\xee\x4c\x41\x7e\xba\xe1\x89\x6e\xd1\x96\xf3\x81\x7e\x51\x41\xbc\x32\x17\x81\x01\x1a\x6c\x1f\x62\x9e\x1f\x85\xb4\xa4\x67\x58\x92\x7b\xf7\xd0\x88\x19\xbc\xf8\xf3\xab\x9e\xbd\x21\x11\xea\x21\x1d\xc8\xdf\xd6\x25\x92\x3e\xfe\xe4\xeb\x43\x54\xda\x75\x59\x8f\xf0\x82\x35\xcf\xa2\x0c\x9f\x51\x8f\xdf\x97\xbe\xd5\x0e\x9e\x80\xfe\x65\xe8\x4a\xce\x15\x02\x63\xe8\x1b\x7d\x5f\xb7\x17\x4e\x9f\xce\x3d\xb3\xfe\xc4\x76\xcf\xaf\x1d\xbb\x8e\x6f\x7b\xbd\xe7\x7c\xd8\xa5\x4b\xaa\x69\xd0\xe6\xde\xa6\x8d\x11\x3b\x7c\x27\xfa\x72\x08\x54\x43\xbe\x1d\xd0\xc3\xf0\x77\xf4\x2f\x2b\xff\x3c\xfb\xfb\x3f\x1e\x2c\xa5\x01\xb6\x74\x26\xa7\x7e\x04\x6c\x11\xac\xe5\x46\x69\x5f\x36\xd2\x8e\x19\xbd\xbe\x7c\xf3\xf5\xd3\x67\x64\x8e\x66\x5e\x15\xf0\xc5\x07\x69\x1a\xba\x78\xff\x2b\xec\x57\xad\x9a\x37\x88\x87\xc7\xe9\xca\xd2\x26\x32\x45\x5f\x7d\x49\x33\x5c\xd4\x76\xa2\x8c\xcd\x37\x8b\x0e\xf8\x59\x00\xcf\x07\x25\xfe\x83\xb5\xa8\x0e\x04\x16\x14\x95\x53\xd5\x2e\x37\x6a\xa7\x4a\xd5\x5e\xeb\x1a\x5c\x8a\x1a\xc6\x6f\x6a\x27\xaa\x69\x51\xd3\xef\x68\xe6\x17\xb2\xda\xd7\xbb\x8d\x97\x17\xb0\x3f\x45\xce\x3a\x60\xeb\xe8\x7e\x0b\x00\x00\xff\xff\xe9\x1b\x7e\x94\x7f\x0a\x00\x00")

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

	info := bindataFileInfo{name: "service.tpl", size: 2687, mode: os.FileMode(420), modTime: time.Unix(1575441929, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vueTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xed\x6f\x1b\xc7\xd1\xff\xae\xbf\x62\xb2\x70\x02\x2a\xe1\x9d\xa4\x3c\x08\x1e\xe4\xc4\xe3\x13\x44\x79\x9e\x07\x41\x8b\x36\x40\x93\x4f\x86\x11\xac\xee\x96\xe4\xd6\x7b\x2f\xdd\x5d\x52\x52\x08\x02\x45\x9b\xc4\x76\xda\x44\x2e\x9a\x0a\x69\x5e\x60\xbb\x70\xfa\x06\x58\xf1\x87\xc2\xb5\x63\x25\xfd\x67\x44\xca\xf9\xd4\x7f\xa1\xd8\xdd\x23\x6f\xef\x85\x14\xe5\x28\x41\x0d\x14\x10\xc8\xe3\xee\xec\xcc\xec\xcc\x6f\x66\x67\xe7\x04\xd0\x7a\xca\x71\xe0\xff\x49\x4c\x38\x96\x24\x84\xed\x3d\x10\x0c\x07\x97\x09\x77\x9c\x36\xac\xb4\x24\x89\x52\x86\x25\x69\xaf\x00\xb4\x08\x73\x02\xcc\x43\xf5\x0c\xd0\x0a\xe9\x00\x04\x4b\xa4\x8f\x7a\x04\x87\x84\x23\x33\x6e\xe8\x3a\x09\x8f\xc0\xa3\x31\xa3\x31\xf1\x91\xe4\x7d\x82\x60\xe0\x44\x49\x48\x98\x8f\x04\xc1\x3c\xe8\x29\x12\x04\x2f\x89\xfe\x76\x44\xa5\x1b\x63\x49\x07\xc4\x4d\x39\x19\x90\x58\xfa\x28\x4e\x62\x32\xe3\x98\xf3\x74\xa8\x24\x51\x3e\x6c\x26\x68\x9c\xf6\x25\xa4\x9c\x74\xe8\xae\x43\x83\x24\xf6\x91\x1a\x0d\x92\xd8\x31\xa2\x10\x04\x8c\x60\x8e\xb7\x19\x99\x6a\xe1\x4a\x4e\x23\x5b\x15\x97\xa6\x08\x52\x86\x03\xd2\x4b\x58\x48\xb8\x8f\x26\x37\x3e\x9f\x5c\xfb\x7a\xfc\xee\x3b\xe3\xc3\x07\xa8\xdd\x5a\x9b\x4a\xb2\xb4\x5a\x2b\xa8\x05\xcb\xa9\xbb\xdd\x97\x32\x89\x41\xee\xa5\xc4\x47\x29\xa7\x11\xe6\x7b\x08\x5e\x0a\x18\x0d\x2e\xfb\xe8\x27\x99\xc6\xb5\xfb\x68\x4f\x6e\x7c\xfe\xe8\x8b\x3f\x68\xb9\x86\xcd\x1c\xe6\x53\x6e\x9c\x08\x22\x2d\x7b\xb7\x27\xf7\xdf\x39\xf9\xcb\x97\x86\xcd\xe4\xb3\x5b\xc7\x0f\xef\xd5\x32\x2b\xed\x6c\xa5\x34\x3a\x1b\x50\xe8\xa9\x11\x4b\xa3\x34\xe1\xb2\x43\x19\x71\x07\x54\xd0\x6d\x46\x7c\x83\x01\x9e\xf4\xe3\xb0\xbc\xf5\x80\x61\x21\x7c\xa4\xb0\x95\x31\x42\xed\x16\x9d\x0e\x4f\x4d\xd0\x4f\x59\x82\xc3\xe7\x95\x27\x68\x7b\xfc\xc5\xd1\xf8\x9d\xcf\x1f\xdd\xfa\xf3\xe4\xe6\x91\xbd\x01\x70\x9c\xd3\x74\xfb\xdf\x5d\xa5\x1b\x2a\x2b\x61\x34\x5b\x4e\x95\x30\xd9\x89\x95\x32\xb9\x2e\x57\xbe\x5c\xac\x4b\x55\x8d\x80\x13\x2c\x67\xe6\x01\x1f\x8c\x81\x96\xd5\xaa\x84\x8f\x94\xf5\x05\x6a\x4f\xfe\xfe\x70\xfc\xde\xcd\x1a\x87\x9a\x60\x95\x7b\x8c\xf8\x48\x87\x82\xb7\x9d\xc8\x9e\x52\x3f\xa4\x83\x2c\x9e\xad\x47\xc2\x1c\xa9\x83\xc5\x0b\xb1\xc4\x3e\xd2\x3f\x5c\xf5\x8c\xa6\x5c\x76\x68\x28\x7b\x1e\x6c\xac\xaf\x3f\xad\x82\x98\x30\x12\x48\x9a\xc4\x4e\xd0\xc3\x71\x97\x28\x18\x67\x23\x5b\x7a\x40\x45\xbf\xb2\x18\x8d\xbb\x53\x76\xd9\xcf\x42\xda\xd0\x33\x4e\x90\xb0\x7e\x34\x0d\x90\x19\x6b\x04\x5a\xa6\x8f\x5e\x78\x01\x15\x00\x69\x2f\xca\x11\x3c\x1c\x72\x25\x18\x2e\xd0\xe6\x85\x20\x61\x9e\xef\x6e\x69\x0a\x31\x1a\xcd\x13\xc8\xf0\xb6\x4a\x4f\xc3\xa1\x5a\xe0\x6e\xe1\x88\xb0\x2d\x2c\xc8\x8f\x70\x44\x46\x23\x3b\x19\x4d\xb3\xa2\xce\x7f\x8e\x08\x12\xad\xa8\xfa\x42\x85\x78\x14\x29\x8e\xdb\xd0\x6a\xb5\x40\x4f\xba\x3c\xd9\x71\xa7\xdc\xb5\x48\xc3\x1a\xda\xed\x36\xb4\xd6\x34\xb5\x15\x81\x76\xee\x3d\x65\xb3\x30\x1c\x92\x38\x3c\x75\x67\x93\xdf\xbe\x7f\xfc\xd5\xa7\x8f\xb3\x93\x1c\xbf\x82\xbe\xa5\x48\x22\xcc\x58\x05\xad\x29\xc3\x34\x47\x78\x3f\x0d\x15\xc2\xf5\x51\xe0\xc3\x8f\xb7\x7f\x4a\x02\xe9\x62\x21\x68\x37\x6e\x0c\x47\xcd\x99\x49\x56\x37\x33\xca\x72\x2c\x14\x21\x4e\x42\x2a\x51\x1b\x4e\xcf\x7d\x35\x1a\x86\x0a\x0a\x3c\x4f\xb2\xaf\x10\x46\x24\x69\xe4\x2a\x94\x85\x85\x9a\x20\xcb\xfa\xd5\xe4\xb8\xa4\x6b\xf2\x61\xfb\xcc\x34\x01\x14\x61\xde\xa5\xb1\xc3\x48\x47\x7a\xf0\xfc\x7a\xba\xbb\x99\x8d\xc8\x24\x35\x03\xa8\x26\x7b\x0c\x1c\xd1\x4b\x76\x7c\x14\xf5\x99\xa4\x29\x23\xb3\x30\x73\x19\x89\xbb\xb2\x57\x71\xc9\x74\xc7\x5b\xea\xeb\x65\x2c\x83\xde\x6b\xea\x03\xb5\x27\xd7\x1e\x7c\x73\x65\xff\xf8\x1f\x87\x93\x0f\x1f\xd4\x65\x8c\xb3\xcb\x2c\x1b\x39\x17\xf9\x4a\x66\x4d\x23\x73\x7c\xf5\xe6\x37\xbf\xbf\x5d\x91\x59\x4a\x40\x29\xee\x52\x55\x14\x24\x31\x78\x3d\x1a\x12\x47\x9d\x82\x34\xee\x32\xa2\xa6\x66\x45\x45\x96\x1a\x73\x6a\x04\xdb\x38\xb8\xdc\x35\x89\x93\xe1\xbd\xa4\x2f\x7d\xa4\x31\x21\x9a\x20\x13\x89\x59\x53\x95\x0a\x83\x26\x28\x36\xbc\x09\x31\xd9\x95\x08\x3c\x3d\x35\xcd\x4d\xfa\x07\x02\x4f\x91\x38\x06\x4f\x59\xd2\xa2\x11\x95\xf6\x84\xf0\xd1\xc5\x8d\xf5\x26\xbc\xb0\xde\x54\xc9\xb0\x09\x2f\xbe\xf8\xe2\x25\x95\x11\xe9\x5b\x24\x4f\x86\xf4\x2d\x32\xcd\x83\xc6\xbe\x5e\xd0\xe7\x9c\xc4\x72\xba\x17\xcd\x5c\x3d\x2b\xf3\x65\x53\xd3\xd5\xaf\xe1\xee\x74\xb5\x05\xab\x7c\xc7\xed\x95\x99\xd1\x42\x8a\x59\xd2\x05\x49\xa5\x82\x98\x39\x0e\x10\x78\x59\x60\xb9\x62\x2f\x0e\xca\x27\x4f\x9e\x56\xd7\xcb\x98\x33\x45\x5c\x56\xb4\x65\xab\x4c\xd5\xa6\xf3\x89\x93\x26\x82\x2a\x05\x7c\xc4\x69\xb7\x27\xa7\xc3\x19\xbf\x8d\x02\xbf\xe5\x52\x72\x46\x46\x3b\x10\x27\x12\x74\xa2\x7c\x55\xbc\x66\xb0\xfc\x03\xb2\x67\xd1\x15\x2a\xac\xa5\x33\xb7\x5d\x2f\x4e\x0b\x41\x2d\x6e\x2b\x89\x25\xa6\xb1\x30\x32\x5f\xdf\x4b\x09\x20\x1a\x4b\x34\x1a\xc5\xfd\x68\x9b\xf0\xe1\x90\x30\x41\x46\x23\x55\x35\x66\xb9\xb6\x60\x92\xba\xb4\x8e\x00\xf7\x65\xe2\x04\x49\x94\x2a\xf8\xfb\x28\xe9\x74\x96\xaa\x23\x6d\x63\x94\x32\x7b\x4d\xae\xaf\xd4\x64\x79\x61\xde\x49\x12\xa9\x22\x32\x0b\x12\x83\x0e\x27\x1b\x2d\x16\xd7\xa7\x55\x27\x1d\xcc\x04\x41\xed\xf1\xfe\x01\x4c\xee\x5d\xad\x4f\x8b\xa7\xd6\xb6\x5b\x9a\x2b\x6a\x9f\xdc\x3a\x84\xf1\xe1\xc7\x75\x69\xc7\x2e\x48\x66\x80\x9e\x8b\xf0\x93\xa3\x83\x47\x5f\xff\xa6\x82\xf0\xe2\x79\xb2\x2c\xc2\xad\xf3\xea\x71\x10\xbe\x1c\xba\x9f\x18\x70\x5b\xd6\xf8\xae\xc0\x5d\x04\xf2\xf7\x01\xeb\x4a\xa1\x71\x2e\xb0\x7e\x43\x73\x3d\x47\x58\xdb\x67\x73\x05\xdc\xdb\xea\x44\x4d\xd5\xc7\x59\x01\x6e\xad\x7c\xfc\x34\xbe\x64\x12\x7f\x72\x80\x5e\xb2\xca\xf7\x03\xf6\xef\x07\xee\x55\xa8\x9c\x17\xe4\xed\x4a\xf2\xec\xb0\xcf\x8a\xa0\x1a\xe0\xef\x5f\x3f\xb9\xfd\xa5\x77\x7c\xff\xbd\xe3\xa3\x9b\xe3\xeb\x1f\x8c\xef\xbe\x3b\xf9\xf0\xc1\x78\xff\xa3\xf1\x07\x37\x26\x9f\x5e\x9b\xfc\xee\xee\xe4\xfd\xc3\x4a\x44\x54\xbb\x0d\xb3\x88\xf8\xaf\xf5\xa7\x0b\xf1\x60\x9a\x08\xc0\x49\x47\x25\x03\x7d\x87\x07\xcf\xba\xb6\xfe\x1f\x65\xb3\x4a\xcd\x53\x1c\x1d\x46\x85\x2c\x88\x50\x1f\x6a\x10\x81\xd7\x93\x32\x75\x38\xf9\x59\x9f\x28\x9a\x57\x23\xd3\x58\xf0\x34\x58\x0c\x77\x1f\x19\x73\x43\xc8\x71\x17\x70\x60\xe2\xcc\xf6\xda\x9c\x3e\x87\x69\x2d\xe4\x64\x0a\x0e\x39\xa1\xa1\x79\xf3\x4d\xa9\x6a\xd7\xb6\xb2\xd2\xc1\x95\xe3\x87\xf7\x26\xbf\x3a\x18\x5f\xbd\x3b\xb9\x73\x7b\x7c\xfb\xed\x7f\x1e\xfd\x7a\x72\xf5\xa0\x10\x33\x51\xfb\xe4\x17\x0f\xc6\x57\x1e\x1a\xf3\xb6\xd6\x6c\x80\x5a\xae\x5a\x24\x8f\xa6\x28\x43\xa5\x7a\x6c\x8f\xf7\xff\xfa\xe8\x97\x5f\x19\x7e\xee\x2e\x13\xbb\x46\x8f\x02\x33\xed\x79\xc3\x60\x36\xa4\xae\xba\xdf\x1a\xdd\x55\xb7\x9f\x17\xba\x4d\x4f\xf2\x8d\xcc\x11\x8b\xf0\x9d\xdf\xd9\x6d\x80\x1b\x7c\xaf\x64\x83\xa6\x67\x6a\x5f\x1a\x57\x5a\x22\xe0\x34\xd5\x79\xc3\xec\x02\x9e\x05\x2c\x00\xa7\x14\x3a\x3c\x89\x00\xbd\xb4\x86\x53\xba\x36\x1c\xba\x59\x02\xda\x5c\x01\x20\xba\x71\x05\x21\xe9\xe0\x3e\x93\x30\xd4\x42\x42\x2c\x71\x63\x35\xfb\x01\xc0\x89\xec\xf3\x78\xf6\x13\xa0\x72\x69\xf3\xe0\xe2\xa5\xe6\x6c\x5a\x5f\x3c\x3c\x8b\x1e\x20\xeb\xcb\x78\xc6\x94\x4d\x6b\x46\xc9\x2a\x2e\x07\x73\xa3\xf2\x60\xdd\x1e\x53\xf7\x18\x0f\x36\xec\xa1\xa4\xd3\x11\x44\x96\xe8\xf4\x65\xca\x83\x8d\xf5\xd9\xd8\x28\x9f\xce\x9b\x96\x1e\x0c\xad\x71\x53\x99\x16\x55\xae\x52\x01\x64\x90\xc8\xb6\x51\x27\xc1\x14\x03\xe7\xc1\x29\xcf\xb3\xe7\xc1\x2d\xc7\x75\x89\x5b\x96\x78\xca\x2e\x98\xc7\x31\x7b\x1a\x6d\xae\x58\x02\x22\x22\x7b\x49\x28\x72\xce\x5d\x22\x4b\x18\x02\x90\x3d\x2a\xdc\x42\x97\x2e\xeb\xc9\x6c\xce\x48\x70\x4a\x2d\x15\xdc\x1f\x52\x21\x1b\x85\xf3\xbf\xd8\xef\x29\x4c\x01\x0c\x67\x80\xb0\x44\x99\xa1\xe6\x14\x16\xb6\x12\x6a\x04\x0a\x66\x9c\xa9\x99\xe3\xa4\x30\xbb\xba\x52\xff\xec\xca\x1e\x89\x1b\x5c\xa4\xe0\xb7\x0b\xd6\x2d\xec\x5a\xa3\x1a\x7c\xe0\x22\xd5\x6d\x4f\x33\xb0\x39\x8f\x5c\x91\xd8\xd4\xea\x63\x2e\x71\x6e\x51\xed\x2f\x9b\x6e\x54\xd0\x34\x50\x80\x6a\x10\xce\xb5\xaa\xa3\xd5\x29\xe5\xcc\x0e\x79\x63\xa0\x31\xc0\x6c\x8e\x03\x75\x83\x61\x80\x59\xdd\x9c\xb1\xb8\xf2\x6d\xd9\xd6\xcf\x82\x62\x09\x0e\x6c\xac\x6e\x16\x57\xce\xf0\x52\x51\x27\xef\x72\x2c\x50\xc7\xb0\xf7\x61\x60\x9b\xf3\x14\xc6\xfa\x25\x86\x89\x7b\xe5\xe7\x2a\x54\xeb\x6f\xab\x9b\xb5\x44\x59\x3b\x72\x38\xaa\x17\x93\xc3\xa9\x2a\x26\x9f\x2b\x30\xc8\x66\xcd\x8b\x9b\x3a\xbb\x64\x13\x0b\xfc\xb3\xb1\xd0\x3b\xeb\x4b\x5b\xca\x5c\xb0\x0b\xa2\x18\x91\x60\x05\xb1\x5a\x7f\x21\xfb\xdd\x28\xa6\xfd\xe0\xb2\xa7\x83\xbc\x90\xde\xc9\xae\xf4\x20\x6b\x1c\x1f\xdf\xbf\xe3\xba\x2e\xb2\xe7\x45\x4a\xe3\x98\x70\x0f\x66\xe5\xcb\xb4\xad\x6f\x53\xe5\x1d\x39\x0f\x10\xef\x6e\xe3\xc6\x7a\x13\xb2\x3f\xf7\xbf\x57\x51\x9e\xb3\x56\xe7\x66\x98\x6c\x6b\x65\x57\x9e\x25\xb6\x2b\x38\xaa\x89\xd1\x8a\x6d\xe1\xd4\xb8\xac\x68\xd0\x58\xad\x2a\x90\x99\xc5\x0d\x58\x22\x48\x99\x7b\xc5\x8f\x59\x7f\x9a\x27\x3b\x15\xd4\x5c\x08\x92\xb8\x43\x79\xd4\x40\x93\x3b\xb7\x8d\x5f\xc6\x77\xdf\x35\x3d\x55\x35\xf2\xd9\x2d\x53\x25\x37\x61\xf2\xd1\x17\xe3\xeb\x7f\x3c\x79\xf8\xa7\x93\x87\x77\xfe\x07\x35\x21\xab\xae\x51\xb3\xa0\x5a\xc6\xee\x65\x5d\xd8\xbc\x6e\xdc\x7d\x72\xeb\x70\x7c\xf8\x71\xc1\x85\x01\x8e\x03\xc2\x0a\x54\xe3\xfd\x83\xc9\xbd\xab\x05\x2a\x55\x4f\x79\x80\x76\x30\x8f\x15\x06\x56\xea\x0d\x38\xd7\x48\x45\x97\x6b\xda\xdc\x12\xee\x70\xe8\xe6\x97\x48\x73\x3d\x2b\xdc\xd2\x56\xcb\x6b\xe7\xc3\x21\xb7\x66\x9c\x48\xda\xd9\x73\x45\x3f\x08\x88\x10\x0d\x94\x19\xf2\xea\xf5\xf1\x7b\x37\x50\x09\x20\xb0\x10\x24\x95\x7d\xc2\xa2\x24\x5e\xb5\xca\xa9\xe9\xde\xb4\x19\x0a\xc1\x5d\x8a\x91\x8c\x42\xeb\x58\xec\xdf\x2c\xb6\x5c\x13\xca\x4b\xce\x12\x56\xf5\x4d\x95\x9a\xd8\xaa\x18\x7b\xf2\xc9\xdf\x26\x07\x77\xeb\x8d\xfd\xb8\xd1\x58\x93\x7e\x0b\x6f\x22\x6b\xcf\xa6\x4a\x95\x5c\x3c\xa1\xf2\xfc\x5a\x7a\xa5\x51\x4d\xea\x56\x78\x5a\x2f\x3b\xbe\xf9\xf9\xb5\xe3\xfb\x77\x4e\x3e\x7e\x7b\xf2\xd9\xad\x93\x4f\x9e\xc8\xd8\xd4\xbb\x2b\x6c\xfd\x34\x9f\x64\x5c\x6a\x5c\x32\xcf\x82\x4f\xec\x59\xa5\x14\xa7\xa1\x00\x1f\x2e\x5e\xca\x47\x3b\x09\x87\x86\x9e\x8a\x43\xb2\xab\x0f\xf2\xec\xb1\x35\x07\x76\xd9\x1b\xb5\x8c\xec\xb9\xe7\x56\xcb\x70\x10\x12\x74\xab\xcc\x9f\xc3\xe0\xa2\x5e\x78\xc9\x76\x0d\x0d\x85\x9b\xf6\x45\xaf\xa1\x16\x9e\x9e\x44\x37\x2b\x17\x88\x4a\x9a\xb1\xdd\x37\x84\xd3\x38\x7a\xca\x30\x35\x20\x9b\x9f\x50\xce\x94\x94\xff\x8d\x4e\xed\xd2\x3b\xd6\x6a\x72\x28\xb5\x1a\xeb\xaa\xc8\xda\x96\x9d\x7d\xf5\x2a\x06\x51\x55\xd0\x7f\x62\xe8\xc9\x89\x21\xe3\x3d\x58\x36\x86\x9a\xb5\x28\xfa\x56\x91\x65\xde\x31\x9c\x35\xb2\x16\x61\xf5\x8c\x97\xda\xef\x20\x0c\xcd\x3f\x78\x95\xeb\x24\x77\x3a\x5c\xa1\x37\x7d\x5b\xed\x5a\x7b\xcd\x00\x73\x7d\x66\x82\x6f\x89\x43\x38\x4d\x19\x0d\xf4\x7b\xf7\xb5\x41\x1c\xba\x49\x4a\xe2\xdd\x88\x29\x47\x60\x29\x9c\xa4\xd3\xa1\x01\x09\x93\xa0\x1f\x91\x58\xba\x22\xe5\x04\x87\xa2\x47\x88\x8c\x98\xab\xbf\x51\xae\x3d\xed\x80\xdd\x18\xd1\xd8\xd2\xed\x4c\x2d\xf6\x29\xdf\x37\xf2\x9f\x79\x66\x31\x51\x41\xa7\x9d\x74\xaa\x84\xee\xc7\xe6\xa1\x56\x0c\x03\x83\x86\x88\x08\x81\xbb\xc4\x25\x9c\x27\xbc\x81\xec\x6e\xae\x6e\xe6\xde\x3c\x1a\x1f\xed\x17\x61\x61\x1a\x8c\x75\x00\x5f\x26\xd5\x14\x02\x7f\x5e\x68\x58\xfe\xd0\x3b\xfd\x76\x47\x87\xfe\x47\xc1\x05\x00\x9f\xdf\x45\x3e\x67\x10\x17\xed\xb3\x5a\x07\xe6\xc5\xa7\x59\x05\xb8\x76\x9b\xba\xa6\x12\xe5\xa4\xa3\x8a\x73\x35\xeb\x1a\xd2\x1a\xf0\xe7\x2f\x3c\x1a\xca\x06\x4d\x60\x54\xc8\x0a\xaf\x9a\xd7\x1f\x2a\x4f\xab\x6f\x37\x4d\xd2\xc6\xea\xa5\x0a\xe3\x38\x89\x49\x63\x75\x68\xf0\x31\x6a\x9a\x96\xb8\xb9\x7c\xdb\xca\x96\x36\x6c\xc8\xd4\xa2\xd1\xe6\x4a\x6b\x6d\xda\x26\x5f\x81\x7f\x05\x00\x00\xff\xff\xcb\xd7\x86\xc3\xaa\x2c\x00\x00")

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

	info := bindataFileInfo{name: "vue.tpl", size: 11434, mode: os.FileMode(420), modTime: time.Unix(1575441929, 0)}
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
