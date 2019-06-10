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

var _controllersTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x51\x6f\xdb\x36\x10\x7e\x16\x7f\xc5\x4d\x0f\x83\x14\x68\x52\x3a\x14\x7d\xf0\xe0\x87\xc6\x09\xda\x6e\x41\xea\x2d\xc9\x53\x51\x0c\x0c\x79\x76\x88\x48\xa4\x47\x52\xe9\x0c\x4d\xff\x7d\x20\xad\xb4\x96\x2d\xba\x5e\x6c\x2c\x7a\x92\x2d\xf2\xbb\xef\xbe\xef\xee\x28\x15\xc5\x3b\x94\xa8\xa9\x45\x0e\x77\x4b\x30\x25\x65\x0f\xa8\x81\x14\xc5\x5c\x48\xd0\xaa\xb6\xa8\x49\x71\x42\xe8\x42\xe4\xef\x2e\x6e\x92\xb8\x68\x9a\xfc\x52\x7d\x41\x7d\x45\x2b\x6c\xdb\x62\x24\x78\x9c\x31\x25\xad\x56\x65\x89\xda\xe4\x37\xf4\x01\x9b\x26\x9f\xd0\x0a\xcb\x09\x35\xb8\x5a\x97\x06\x11\xfa\xbb\x2f\x85\xb1\xa1\xdd\xd3\x8f\xd7\xdf\xdb\x3e\xd1\x48\x6d\x30\xfc\xf4\x76\x9f\x04\x6e\x17\x7c\x17\xc6\xdb\x9b\xc9\xfb\x3d\x50\xa6\xd4\xb2\xfb\x74\x1b\xc5\x83\x9c\x5f\x5c\x5e\xdc\x5c\xec\x81\x72\x8e\x25\x0e\x72\x81\x1d\x64\xfa\x10\x67\x8e\x88\x67\x13\x4a\x29\xc0\x66\x00\x26\x48\x27\x28\xef\x00\x48\x50\x5f\x20\x27\x05\x21\x0b\xca\x1e\xe8\x1c\x61\x6d\x1f\x21\xa2\x5a\x28\x6d\x13\x12\x41\xdc\x34\xf1\x5c\x19\xcd\x8a\x4a\x71\x2c\x4d\xfc\x0f\xe4\x1f\xfc\xd3\x4b\x71\xa7\xa9\x5e\xb6\xad\xa3\xd0\x45\x27\x51\x3c\x17\xf6\xbe\xbe\xcb\x99\xaa\x5c\x49\xff\x34\x57\x52\x30\x77\xe7\x9e\x49\xb4\xc5\xbd\xb5\x0b\x77\x6f\xac\x66\x4a\x3e\xc6\x24\x25\x40\xc8\xac\x96\x0c\x00\x86\xcb\x31\x61\x70\x32\x17\x32\x9f\x28\x69\xf1\x6f\x9b\x02\x34\x04\x00\xe0\x91\x6a\xe8\xa7\xbf\xf1\x33\x1f\xa8\x87\x48\xcd\x66\x06\x6d\xf6\xe7\x68\xdc\x51\xc8\xa7\x54\x1b\xbc\x15\xd2\x26\x2c\x3f\xc7\x19\xad\x4b\xfb\x7b\x8d\x7a\x99\xc4\xab\xb5\x71\x16\x9f\xc6\x69\x06\xaf\x4e\x33\x78\xf3\xda\xf1\x75\xd1\x4b\x51\x89\x3d\x61\xfc\xd2\x0e\xe5\xd5\x69\xe6\x31\x22\x96\x5f\xdf\xab\xba\xe4\x67\x42\xf2\xd5\xba\x1f\xfb\xec\x53\x17\x86\x53\x4b\x33\xab\x2c\x2d\x33\xd4\x7a\x34\x86\x6b\xd4\x8f\x82\x61\xa0\x73\x93\x2e\x3b\x1f\x31\x25\x51\x14\xfd\x7a\xfd\xf1\x2a\x61\x19\x38\x01\xdf\x37\x24\x8a\x62\x07\x19\x8f\x3a\x68\xf7\x87\x87\x8f\x47\xb0\x0a\x43\xa2\xd6\x85\x4a\x49\xfb\xd5\x96\xe1\x19\x73\x5c\x5b\xf6\xd4\x03\x04\xcf\x00\xb5\x86\xd1\x18\x7a\xc2\x7f\xf0\xba\x4f\xa9\xa6\x55\x12\x0b\xfe\x4d\xe9\x48\xcc\xfc\x86\x1f\xc6\x20\x45\x09\x4e\x82\x27\x4d\xa4\xf0\xaa\x32\xc5\x31\xbf\xc2\x2f\xc9\xd3\xfd\x19\xe5\x7f\xe0\x5f\x35\x1a\xeb\x95\x70\x3a\x6a\xb4\xb5\x96\x24\x6a\x57\xe6\xf7\x99\xad\x18\x7d\xf3\x26\xa0\x97\xe0\x29\xac\x5b\xb2\x01\xd2\x17\x3d\x34\x5a\x8f\xaa\xba\x17\x74\xf6\xa4\xe7\xba\x03\x9e\xe3\xa6\x01\xbf\xf4\x85\x04\x7f\x3d\x43\x4e\xe8\xae\xaf\xa2\x0e\x69\xda\x91\x7a\xd2\x34\xa8\xc7\x66\x91\xec\xa7\x6f\x68\x2a\x0e\xc9\x1b\x1d\xa5\xea\x0e\x29\xba\x03\xfa\xea\xf9\xf6\x1e\xd6\x27\x24\xda\x26\x37\xd5\xa2\xa2\x7a\xf9\x1b\x2e\x27\xaa\xac\x2b\xb9\xc9\x16\xc6\x20\x38\x89\xfa\xcd\x14\x74\x2a\x6c\x7c\xc7\x77\xad\x9b\x02\x07\xf2\xa0\xdb\x2f\xed\x75\xe4\x8c\xae\x7d\xd6\x33\x81\x25\x37\x30\x86\x8a\x3e\x60\x52\xd1\xc5\x27\x63\xb5\x90\xf3\xcf\x42\x5a\xd4\x33\xca\xb0\x69\xd3\xdd\x1e\xaf\x03\x6d\x39\x7c\xd8\x20\xdc\x30\x2a\xa4\xb1\xe0\x59\x8f\xc4\x56\xd8\x35\x9f\x42\x6f\x3c\xa1\xa9\x77\xb0\x57\x70\xf0\x71\xb0\xa1\x42\x30\x03\xc1\x43\x25\x5a\x9c\x74\xd9\xef\x7c\xe9\xdb\x92\xc0\x97\xaa\xe9\x52\x17\x72\x6e\x7e\x16\xd2\xbe\x79\x6d\x12\x96\xfb\xf3\xf3\xad\xd6\x74\xb9\xca\x3e\x5d\x2d\x1d\x03\x5d\x2c\x50\xf2\x44\x70\x93\x7d\x6f\xd7\xa7\xcf\x71\x9a\xe6\x79\x9e\x76\x29\xae\xcd\xe2\xdd\x3c\x45\xd0\x63\xb2\x9e\xe7\xbe\x2d\xd9\xfc\x9f\x0d\x71\xe8\xcc\x7b\x49\x43\xc2\xfd\x67\x32\xd8\xe8\xc0\x1d\x2d\xb8\xf3\xa3\x21\xe4\xce\x7f\x3d\x97\xe0\xa8\x07\xd3\x4b\x98\x30\xec\x41\x50\x37\xd7\x14\x03\xb3\x26\x25\xad\xfb\x0c\xfb\x37\x00\x00\xff\xff\xe5\x8f\xc6\x76\x17\x10\x00\x00")

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

	info := bindataFileInfo{name: "controllers.tpl", size: 4119, mode: os.FileMode(420), modTime: time.Unix(1560151022, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _daoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x4d\x6f\xdb\x38\x10\x3d\x93\xbf\x62\x2a\xa0\x81\xd4\x0a\x74\x0b\x2c\xf6\x60\x40\xbb\x68\xd3\x6c\x91\x5d\xf7\x23\x75\x8a\x3d\x04\x41\x4a\x4b\x23\x87\xb0\x44\xda\x24\xb5\x89\xa1\xea\xbf\x2f\x48\xca\x8e\xbf\x5a\x27\x40\x7d\x11\x68\x0d\x87\x6f\xde\x7b\x33\xd4\x60\xf0\x1e\x25\x6a\x6e\xb1\x80\xc9\x12\x4c\xc5\xf3\x19\x6a\x3a\xe7\xf9\x8c\x4f\x11\x0a\xae\x28\x15\xf5\x5c\x69\x0b\x31\x35\x0b\x12\x4d\x85\xbd\x6d\x26\x2c\x57\xf5\xe0\x03\x37\x16\x75\x2d\x64\x61\x06\x66\xd1\x08\xad\xb1\x8a\x68\xd4\xb6\xd1\x54\x19\x9d\x0f\x6a\x55\x60\x65\xa2\xef\xc0\xce\x7d\x86\x91\x98\x68\xae\x97\x5d\x37\x68\x5b\xf6\x91\xd7\xd8\x75\x11\x4d\x28\x2d\x1b\x99\x43\x5c\xc0\x8b\x77\x5c\x25\x30\x12\xc6\xb6\x2d\x3b\xe5\x35\x56\xa7\xdc\x60\x08\x8c\x55\x59\x1a\xb4\x69\x25\x6a\x61\xa1\x11\xd2\xfe\xfe\x5b\x02\xf1\xd5\x75\xdb\xb2\x91\xba\x43\x1d\xc2\xd8\xfe\xce\x54\x48\x9b\xa2\xd6\x4a\x27\xd0\x52\x00\x80\xff\xb8\x86\x82\x5b\x0e\x59\xcd\x67\xf8\xa8\x1c\xaf\x92\xf5\x4e\xab\x2c\xaf\x40\x48\x4b\x89\x5b\xde\xdd\xa2\x46\x30\x0b\xf6\x46\x16\x94\xb4\xad\x28\xe1\x54\x49\xcb\x85\x34\xc0\xc6\x77\xc2\xe6\xb7\x2e\x0f\x44\xc6\x72\x8b\x51\xd7\x51\x42\xc2\x9e\x0c\xf8\x7c\x8e\xb2\x88\xfd\x32\x75\x39\x3e\x2a\x7b\xb6\x68\xfb\xd0\xe1\x0e\xac\xb1\xfb\xf7\x1d\x56\xd0\x25\xee\x20\x94\x85\x4b\xb6\x48\x81\xeb\xa9\x49\xe1\x06\x86\x99\xcb\x31\xc6\x0a\x73\x1b\x47\xb9\x6a\xa4\x8d\xdb\x96\x7d\xd6\xa2\xe6\x7a\xf9\x0f\x2e\x4f\x55\xd5\xd4\x92\x85\x47\xc8\x9a\x44\x09\xfb\x4b\xab\x3a\x8e\x1e\x34\x49\xd8\xbf\x0e\x51\xc0\x95\xb0\x4b\x35\x5e\x54\x71\x20\x00\xb5\x1e\x66\x05\x2b\x26\xec\x3d\xda\xf8\xc4\x73\x91\x2e\x52\x87\x80\x31\x16\x62\x44\xe9\xc2\x9e\x65\x52\x54\x81\x6f\xf7\xd3\x68\x1b\x2d\x3d\xed\x69\xd8\x85\x5a\xfb\xb7\x1d\x25\x93\x46\x54\x05\xea\x6d\xfc\x2f\x8e\x23\xfb\xa4\x0b\xd4\x6f\x97\x3e\xe2\xa7\x55\x42\x81\x26\x8f\x12\x4a\x44\x09\xde\x41\x7f\xbc\x6a\xd7\xc7\x66\xfd\x93\x8d\xdc\x9b\xd8\xbf\x4f\x28\xe9\x7c\x74\xb0\xdd\xc1\xf0\x4f\xfe\x55\x6f\xcc\xb0\x21\x30\x91\xde\x40\x06\xab\xa8\x1d\xf6\x02\x79\x7d\x8d\x27\x9e\x8e\x1d\xfa\x0e\x33\xd5\x51\xda\xb6\xec\x03\xda\x5b\x55\x5c\xf2\x19\x76\xdd\x6e\xdf\x9c\x6a\xe4\x16\x0f\x74\xce\xb6\x8d\xe0\xa8\xd9\x93\xf8\x78\x3f\x80\x6f\x28\x48\x5a\x0a\x64\x33\x1a\xbe\x03\x7b\xd3\x58\x55\x73\x2b\xf2\x80\xe8\xeb\xbc\xe0\x16\xcf\xee\xe7\x1a\x8d\x11\x4a\x3a\xe4\x7d\xa1\xa6\xa9\x7c\x6b\x3a\xe7\x7a\x5e\x5c\x8a\xe2\xec\x1e\xf3\x38\x12\xd2\xa0\xb6\xae\xcf\x14\xac\x0d\x00\x06\xed\x6a\x55\x8c\x2f\x46\x5d\x17\xa5\xdb\x60\x83\xc6\x0f\xfe\x23\xa4\xe7\x73\x3b\xcc\x73\xea\x14\xdb\x2f\x75\xdf\x48\x3b\xb5\x3b\x15\x03\x76\x36\xe2\xc6\x9e\x7b\xa0\xe7\x45\xbc\x25\xdf\xce\x71\xd0\x6b\xb8\x2d\x59\xa0\xe6\x57\x48\xd6\xeb\xd1\x52\x02\x3f\x94\xe3\x80\x10\x0e\xf0\xcd\x8f\x14\x68\x7c\xfc\xcf\xc8\xef\x87\xdf\xb1\xee\xcb\x86\xc7\x22\xf6\x55\xdc\xa0\xf2\x20\x75\x9f\xb9\xcd\x6f\x0f\x30\x27\x0a\xf0\x77\x43\xda\xa3\xaf\xf9\xfc\xca\x58\x2d\xe4\xf4\x5a\x48\x8b\xba\xe4\x39\xb6\x1b\x7c\x81\x43\xbf\x43\xd1\x07\x3e\xdf\x62\xc9\xcf\x79\xe9\xaa\x86\xab\xeb\x90\x8c\x92\x52\x69\x98\x39\xda\x34\x97\x53\x84\x15\x59\x94\x10\xe3\x87\x3e\xcc\xbc\x1a\x00\xb9\x9b\xfe\x6d\xbb\x71\x17\x74\xdd\x90\x12\x42\x42\xc6\xf5\x2d\xe0\x97\x29\x94\xb5\x65\xe3\xb9\x16\xd2\x96\x71\xf4\xed\xb9\xf9\x96\x0d\x9f\x9b\x28\x85\x59\x0a\xb3\x24\xa1\xc4\x99\x36\x8c\xa6\x0a\x65\xd8\x94\x40\x96\xc1\x2b\xd8\x30\xbb\x14\x95\x8b\xa2\xa4\x14\x58\x15\xc6\xdf\x0c\x1e\xb7\x61\x7f\x2b\x21\x57\x67\x45\xa9\x9b\x89\x01\xfa\xd5\xd1\x21\x1a\x5d\x43\x06\xa2\xa0\xe4\xc6\xdb\xf9\xd1\x96\x89\x5e\x42\x0f\xe3\x65\xf4\x0b\x1d\xd3\x53\x9e\x50\xb2\xed\x93\x87\x39\xf9\x0e\x2b\xb4\x7e\x52\xba\x3f\x45\x09\xec\xdc\x7c\x35\xa8\x2f\xf9\xa4\x72\xe0\x76\x3d\xe5\xe6\xea\xbe\xa5\xde\x2e\xdd\x33\x6e\x0c\x6a\x47\x5b\x4f\x64\x02\x8f\x9e\x92\x9b\x9f\x1d\x4f\x6d\x6c\x4a\xd6\xe7\x3e\x48\x78\xa9\xfc\xa6\x35\xa4\x84\x92\xf1\xc5\xc8\xe9\x11\x19\x7f\xb1\xc0\x0b\x28\xb5\xaa\x37\x74\x58\xd3\xfe\xb5\xdf\xb3\x47\x29\x74\x5d\xf6\x67\xb8\x1c\xe1\x75\x44\xc9\xa6\xc2\xfe\xae\xdf\x9d\x68\xe3\x8b\x51\x0a\x1b\x10\x9e\x32\xf8\xf6\x0b\x3d\x37\x67\xf7\xc2\xd8\x7d\x9a\x27\x4a\x55\xce\xda\x3b\x15\xbe\x7e\x7a\x85\x3b\x05\x3a\x39\xd0\x9d\xe9\x8f\xd8\x2e\xf8\xa2\x41\xbd\xfc\xa2\xee\x62\x57\xe4\xba\x46\x36\xce\xb9\x8c\x4f\xfc\xa6\xf5\x3d\x03\xcf\x32\xd7\x6e\x70\x72\xb2\x5a\x99\x45\xc5\xce\xb4\xfe\xa8\xbe\xa8\x3b\xb3\xd9\x95\x56\x37\xe8\x9b\x77\x65\x58\x97\x28\xdc\xec\xfe\x63\xee\xff\x00\x00\x00\xff\xff\x18\xd0\x1d\x18\x8c\x0b\x00\x00")

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

	info := bindataFileInfo{name: "dao.tpl", size: 2956, mode: os.FileMode(420), modTime: time.Unix(1560150704, 0)}
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

	info := bindataFileInfo{name: "models.tpl", size: 719, mode: os.FileMode(420), modTime: time.Unix(1560127998, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serviceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x3d\x6f\xdb\x30\x10\x86\x67\xf1\x57\x1c\x3c\x59\x85\x40\x2f\x45\x87\x16\x1e\x5a\x27\x6d\x03\x04\x45\x80\x26\x5d\x0c\x23\xa0\xa5\x93\xcd\x9a\x22\x55\x92\x4a\x62\xa8\xfa\xef\x05\xf5\xe1\x0f\x89\x76\x1c\xa0\x93\x06\xdd\xbd\xf7\xdc\xdd\x4b\x72\x32\xf9\x86\x12\x35\xb3\x98\xc0\x72\x0b\x46\xb0\x78\x83\x9a\xe4\x2c\xde\xb0\x15\x82\x41\xfd\xc4\x63\x24\x84\x67\xb9\xd2\x16\xc6\x24\x18\xad\xb8\x5d\x17\x4b\x1a\xab\x6c\xf2\x3b\x53\x5c\x2b\x39\x31\x7f\xc4\xcb\x88\x04\xa3\xb2\x1c\xad\x94\xd1\xf1\x24\x53\x09\x0a\x33\xfa\x0b\xf4\xa6\x4e\xbc\xe5\x4b\xcd\xf4\xb6\xaa\x26\x65\x49\x7f\xb0\x0c\xab\x6a\x44\x42\x42\x48\x5a\xc8\x18\xc6\x06\xde\xfd\x6c\x2a\x85\xf7\x6c\x83\x65\x49\x67\x2c\x43\x31\x63\x06\x9b\xe0\x31\x4f\x80\x4b\xfb\xe1\x7d\x38\x2e\x4b\x7a\xab\x9e\x51\x37\x3f\xe8\x30\x36\x42\xad\x95\x0e\x4b\x20\x41\xc2\x2c\x8b\x00\xb5\x86\x8f\x53\x30\x34\x61\x8a\x9e\xd4\x0f\x81\x04\x1a\x6d\xa1\x25\xec\xd3\x48\xe5\x63\xbc\xe5\xc6\x7a\x34\x54\x9a\x1a\xb4\x91\xe0\x19\xb7\x50\xb4\xb8\xf3\xc5\xeb\xc0\x5c\xda\x3e\xb4\x55\x96\x89\xe8\x88\xfc\x82\xaa\xfd\x1e\x1a\x95\x5d\x27\x83\x46\x66\x1a\x99\xf5\x8d\xe3\x98\x19\x5e\x6d\xe1\x6d\x6b\x29\x4b\x9e\x02\xbd\x31\x0f\x06\xf5\x3d\x5b\x0a\xac\x2a\x12\x04\x43\x05\xf7\x5f\xb2\x0c\x67\x4a\x14\x99\xec\x0b\xc2\x14\x8c\xd5\x5c\xae\x0c\xbd\x57\x75\xa6\x07\xe2\xbc\x44\xe8\x2d\x7b\xc7\x8c\x79\x56\x3a\x39\x5d\x96\x5e\xcb\x58\x6f\x73\xdb\x45\x7a\x0a\x9f\x17\x71\x85\x79\xda\x2e\x76\x38\xab\x1b\x73\xfd\xc2\x8d\x7d\x7b\x3f\x50\x92\x20\xe8\x2c\x70\x9c\x5d\xdb\x20\x56\x09\xd2\x56\x9d\x04\x41\xe5\x76\x81\x32\x71\xf3\xf7\x9c\x95\x0b\xed\x71\xf2\xe4\x0c\xfc\xf6\x90\x27\xff\xc9\x6f\xb5\x99\x9c\x97\x00\xda\xd2\x0d\xf2\x85\x15\x1c\xb2\x1f\xf1\x8e\xd9\x78\x7d\xe6\x02\x8a\x8a\xba\x02\x64\x2c\x9f\x37\xf6\x5b\x70\x69\x51\xa7\x2c\xc6\xf2\x80\xcb\xef\x71\x9e\xc2\x93\x3b\x91\x6a\xe3\xa6\xdc\x48\xcd\x47\x1e\xbb\xd4\x9f\xf6\xa2\x5c\x7c\x72\xf1\xf5\x66\xf3\x36\xac\x53\x78\x62\x82\x8e\x1b\x0c\x67\x29\xa7\xdf\x85\xee\x62\xbd\x86\xed\x7e\xd6\x59\xc1\xa5\x20\x30\x85\x2e\xd1\xe5\x55\x3d\x0b\x1d\xaf\xe2\xf4\x24\xdb\x19\x86\xfe\x0d\x5c\xa1\x40\xef\x0a\x77\x6f\xc0\x6e\xc8\x47\xf5\xce\xe4\x85\x50\x97\xf2\xed\x64\x50\x1e\xfc\x2f\xc4\x97\xad\xfb\x8e\x8b\xf6\xfc\xb5\x57\x4f\x08\x17\x5c\x7c\xd0\xdc\x7c\x70\x40\x1c\x9c\x7b\x8c\x7a\xa5\x5a\xf8\x96\xf4\x71\x4f\x7a\x6a\xa7\x3b\xb6\xe6\xeb\xec\xb0\x66\x66\x9d\x44\xf0\xe8\x3c\xb3\xac\xb3\x68\xf7\xee\x7f\xd5\x2a\xdb\x49\xcc\x17\xcb\xad\xc5\xbd\x3b\xa2\x2e\xfa\x0a\x53\x56\x08\x3b\x53\xc6\x86\xfb\xb9\xd7\xfa\xe3\x5a\x3c\xdc\xef\xf2\x00\xf1\x17\x13\x3c\x19\x00\x46\xd8\x90\x63\x92\x3f\xef\x69\x97\x4a\x09\xc7\xda\x8a\x4b\x2e\x60\xba\xa3\x9d\xa9\x2c\x67\x1a\xbf\x33\xb3\xfe\x2c\x93\x3e\xef\xa1\x5e\x18\x41\xbf\x8b\x90\x54\x40\x5a\x97\xfe\x0b\x00\x00\xff\xff\x82\x98\x5b\xc5\xf1\x08\x00\x00")

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

	info := bindataFileInfo{name: "service.tpl", size: 2289, mode: os.FileMode(420), modTime: time.Unix(1560150711, 0)}
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

var _vueTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x6d\x6f\x1b\xc7\xf1\x7f\xaf\x4f\x31\x59\x38\x01\x99\xf0\x4e\x72\xfe\x08\xfe\xc8\x89\xc7\x1a\x51\xda\x22\x68\xd1\x06\x68\xf2\xca\x30\x82\xd5\xdd\x92\xdc\xfa\x9e\xba\xbb\xd4\x43\x08\x02\x45\x9b\xc4\x76\xda\x44\x2e\x9a\x0a\x69\x1e\x60\xbb\x70\xfa\x04\x58\xf1\x8b\xc2\xb5\x63\x25\xfd\x32\x22\xe5\xbc\xea\x57\x28\x76\xf7\xc8\xdb\xbd\x3b\x8a\x94\xa3\x04\x35\x50\x40\x20\x8f\xbb\xb3\x33\xb3\x33\xbf\x99\x9d\x9d\x13\x40\xfb\x29\xc7\x81\x1f\x92\x84\x30\x2c\x48\x08\x9b\xbb\xc0\x23\x1c\x5c\x26\xcc\x71\x3a\xb0\xd2\x16\x24\xce\x22\x2c\x48\x67\x05\xa0\x4d\x22\x27\xc0\x2c\x94\xcf\x00\xed\x90\x6e\x01\x8f\x52\xe1\xa3\x3e\xc1\x21\x61\x48\x8f\x6b\xba\x6e\xca\x62\xf0\x68\x12\xd1\x84\xf8\x48\xb0\x01\x41\xb0\xe5\xc4\x69\x48\x22\x1f\x71\x82\x59\xd0\x97\x24\x08\x2e\xf0\xc1\x66\x4c\x85\x9b\x60\x41\xb7\x88\x9b\x31\xb2\x45\x12\xe1\xa3\x24\x4d\xc8\x8c\x63\xc1\xd3\xa1\x82\xc4\xc5\xb0\x9e\xa0\x49\x36\x10\x90\x31\xd2\xa5\x3b\x0e\x0d\xd2\xc4\x47\x72\x34\x48\x13\x47\x8b\x42\x10\x44\x04\x33\xbc\x19\x91\xa9\x16\xae\x60\x34\x36\x55\x71\x69\x86\x20\x8b\x70\x40\xfa\x69\x14\x12\xe6\xa3\xc9\x8d\xcf\x26\xd7\xbe\x1a\xbf\xf3\xf6\xf8\xe0\x01\xea\xb4\x57\xa7\x92\x0c\xad\x56\x2d\xb5\x60\x39\x75\x37\x07\x42\xa4\x09\x88\xdd\x8c\xf8\x28\x63\x34\xc6\x6c\x17\xc1\x85\x20\xa2\xc1\x65\x1f\xfd\x2c\xd7\xb8\x76\x1f\x9d\xc9\x8d\xcf\x1e\x7d\xfe\x27\x25\x57\xb3\x99\xc3\x7c\xca\x8d\x11\x4e\x84\x61\xef\xce\xe4\xfe\xdb\xc7\x7f\xfb\x42\xb3\x99\x7c\x7a\xeb\xe8\xe1\xbd\x5a\x66\xa5\x9d\xad\x94\x46\x67\x03\x12\x3d\x35\x62\x69\x9c\xa5\x4c\x74\x69\x44\xdc\x2d\xca\xe9\x66\x44\x7c\x8d\x01\x96\x0e\x92\xb0\xbc\xf5\x20\xc2\x9c\xfb\x48\x62\x2b\x67\x84\x3a\x6d\x3a\x1d\x9e\x9a\x60\x90\x45\x29\x0e\x9f\x97\x9e\xa0\x9d\xf1\xe7\x87\xe3\xb7\x3f\x7b\x74\xeb\xaf\x93\x9b\x87\xe6\x06\xc0\x71\x16\xe9\xf6\xfd\x1d\xa9\x1b\x2a\x2b\xa1\x35\x5b\x4e\x95\x30\xdd\x4e\xa4\x32\x85\x2e\x57\xbe\x38\x59\x97\xaa\x1a\x01\x23\x58\xcc\xcc\x03\x3e\x68\x03\x2d\xab\x55\x09\x1f\x59\x34\xe0\xa8\x33\xf9\xe7\xc3\xf1\xbb\x37\x6b\x1c\xaa\x83\x55\xec\x46\xc4\x47\x2a\x14\xbc\xcd\x54\xf4\xa5\xfa\x21\xdd\xca\xe3\xd9\x78\x24\x91\x23\x54\xb0\x78\x21\x16\xd8\x47\xea\x87\x2b\x9f\xd1\x94\xcb\x36\x0d\x45\xdf\x83\xf3\x6b\x6b\x4f\xcb\x20\x26\x11\x09\x04\x4d\x13\x27\xe8\xe3\xa4\x47\x24\x8c\xf3\x91\x0d\x35\x20\xa3\x5f\x5a\x8c\x26\xbd\x29\xbb\xfc\xa7\x95\x36\xd4\x8c\x13\xa4\xd1\x20\x9e\x06\xc8\x8c\x35\x02\x25\xd3\x47\x2f\xbc\x80\x2c\x40\x9a\x8b\x0a\x04\x0f\x87\x4c\x0a\x86\x73\xb4\x75\x2e\x48\x23\xcf\x77\x37\x14\x05\x1f\x8d\xe6\x09\x8c\xf0\xa6\x4c\x4f\xc3\xa1\x5c\xe0\x6e\xe0\x98\x44\x1b\x98\x93\x9f\xe0\x98\x8c\x46\x66\x32\x9a\x66\x45\x95\xff\x1c\x1e\xa4\x4a\x51\xf9\x85\xac\x78\xe4\x19\x4e\x3a\xd0\x6e\xb7\x41\x4d\xba\x2c\xdd\x76\xa7\xdc\x95\x48\xcd\x1a\x3a\x9d\x0e\xb4\x57\x15\xb5\x11\x81\x66\xee\x5d\xb0\x59\x18\x0e\x49\x12\x2e\xdc\xd9\xe4\xf7\xef\x1d\x7d\xf9\xc9\xe3\xec\xa4\xc0\x2f\xa7\x6f\x4a\x92\x18\x47\x51\x05\xad\x59\x84\x69\x81\xf0\x41\x16\x4a\x84\xab\xa3\xc0\x87\x9f\x6e\xfe\x9c\x04\xc2\xc5\x9c\xd3\x5e\xd2\x18\x8e\x5a\x33\x93\x34\xd7\x73\xca\x72\x2c\xd8\x10\x27\x21\x15\xa8\x03\x8b\x73\x5f\x8d\x86\xa1\x84\x02\x2b\x92\xec\xcb\x24\x22\x82\x34\x0a\x15\xca\xc2\x42\x45\x90\x67\xfd\x6a\x72\x5c\xd2\x35\xc5\xb0\x79\x66\xea\x00\x8a\x31\xeb\xd1\xc4\x89\x48\x57\x78\xf0\xfc\x5a\xb6\xb3\x9e\x8f\x88\x34\xd3\x03\xa8\x26\x7b\x6c\x39\xbc\x9f\x6e\xfb\x28\x1e\x44\x82\x66\x11\x99\x85\x99\x1b\x91\xa4\x27\xfa\x15\x97\x4c\x77\xbc\x21\xbf\x5e\xc2\x22\xe8\xbf\x2a\x3f\x50\x67\x72\xed\xc1\xd7\x57\xf6\x8e\xfe\x75\x30\xf9\xe0\x41\x5d\xc6\x38\xbd\xcc\xb2\x91\x0b\x91\x2f\xe7\xd6\xd4\x32\xc7\x57\x6f\x7e\xfd\xc7\xdb\x15\x99\xa5\x04\x94\xe1\x1e\x95\x45\x41\x9a\x4c\xd3\x5f\x31\x82\x60\x13\x07\x97\x7b\x3a\x39\x46\x78\x37\x1d\x08\x1f\x29\xbf\xf3\x16\x88\x54\xe0\xa8\x25\xcb\x81\xad\x16\x64\xb8\x47\x58\x0b\x12\xb2\x23\x10\x78\x6a\x6a\x9a\x7f\xd4\x0f\x04\x9e\x24\x71\x34\x66\xf2\xc4\x44\x63\x2a\xcc\x09\xee\xa3\x8b\xe7\xd7\x5a\xf0\xc2\x5a\x4b\x26\xbc\x16\xbc\xf8\xe2\x8b\x97\x64\xd6\xa3\x6f\x92\x22\xe1\xd1\x37\xc9\x34\xd7\x5d\x08\x06\x8c\x91\x44\xcc\x26\x5f\xc5\xbd\xe9\xa4\x81\x8c\x62\x43\x9d\x95\xd9\xbe\x43\x8a\xa3\xb4\x07\x82\x0a\x89\x12\x9d\xd1\x11\x78\x79\x6c\xb8\x7c\x37\x09\xca\x87\x47\x91\x19\xd7\xca\xb0\xd1\x75\x58\x5e\x77\xe5\xab\x74\xe1\xa5\x52\x82\x93\xa5\x9c\x4a\x05\x7c\xc4\x68\xaf\x2f\xa6\xc3\x39\xbf\xf3\x16\xbf\xe5\xb2\x6a\x4e\x46\xbb\x90\xa4\x02\x54\xae\x7b\x85\xbf\xaa\xe1\xf8\x23\xb2\x6b\xd0\x59\x45\xd2\xd2\xc9\xd7\x2c\xf9\xa6\xb5\x9c\x12\xb7\x91\x26\x02\xd3\x84\x6b\x99\xaf\xed\x66\x04\x10\x4d\x04\x1a\x8d\x92\x41\xbc\x49\xd8\x70\x48\x22\x4e\x46\x23\x59\xf8\xe5\xe9\xd2\x32\x49\x5d\x66\x46\x80\x07\x22\x75\x82\x34\xce\x24\x82\x7d\x94\x76\xbb\x4b\x95\x82\xa6\x31\x4a\xc9\xb9\x26\x5d\x57\xca\xaa\xa2\xb6\xee\xa6\xa9\x90\x41\x95\xc7\x80\x46\x87\x93\x8f\xda\xf5\xf1\xa2\x02\xa3\x8b\x23\x4e\x50\x67\xbc\xb7\x0f\x93\x7b\x57\xeb\x33\xdb\xc2\xf2\x74\x43\x71\x45\x9d\xe3\x5b\x07\x30\x3e\xf8\xa8\x2e\x73\x98\x35\xc5\x0c\xd0\x73\x11\x7e\x7c\xb8\xff\xe8\xab\xdf\x55\x10\x6e\x1f\x09\xcb\x22\xdc\x38\x72\x1e\x07\xe1\xcb\xa1\xfb\x89\x01\xb7\x61\x8d\x6f\x0b\xdc\x36\x90\xbf\x0b\x58\x57\x6a\x85\x33\x81\xf5\xeb\x8a\xeb\x19\xc2\xda\x3c\x5e\x2b\xe0\xde\x94\x87\x62\x26\x3f\x4e\x0b\x70\x63\xe5\xe3\xa7\xf1\x25\x93\xf8\x93\x03\xf4\x92\x55\xbe\x1b\xb0\x7f\x37\x70\xaf\x42\xe5\xac\x20\x6f\x16\x83\xa7\x87\x7d\xde\xeb\xa8\x01\xfe\xde\xf5\xe3\xdb\x5f\x78\x47\xf7\xdf\x3d\x3a\xbc\x39\xbe\xfe\xfe\xf8\xee\x3b\x93\x0f\x1e\x8c\xf7\x3e\x1c\xbf\x7f\x63\xf2\xc9\xb5\xc9\x1f\xee\x4e\xde\x3b\xa8\x44\x44\xb5\x61\x30\x8b\x88\xff\x5b\x7b\xda\x8a\x07\xdd\x07\x00\x46\xba\x32\x19\xa8\x6b\x38\x78\xc6\xcd\xf3\x07\x34\x9a\x15\x62\x9e\xe4\xe8\x44\x94\x0b\x4b\x84\xfc\x90\x83\x08\xbc\xbe\x10\x99\xc3\xc8\x2f\x06\x44\xd2\xbc\x12\xeb\xde\x80\xa7\xc0\xa2\xb9\xfb\x48\x9b\x1b\x42\x86\x7b\x80\x03\x1d\x67\xa6\xd7\xe6\xb4\x2a\x74\x77\xa0\x20\x93\x70\x28\x08\x35\xcd\x1b\x6f\x08\x59\x9a\x76\xa4\x95\xf6\xaf\x1c\x3d\xbc\x37\xf9\xcd\xfe\xf8\xea\xdd\xc9\x9d\xdb\xe3\xdb\x6f\xfd\xfb\xf0\xb7\x93\xab\xfb\x56\xcc\xc4\x9d\xe3\x5f\x3d\x18\x5f\x79\xa8\xcd\xdb\x5e\x35\x01\x6a\xb8\xea\x24\x79\x34\x43\x39\x2a\xe5\x63\x67\xbc\xf7\xf7\x47\xbf\xfe\x52\xf3\x73\x77\x22\xbe\xa3\xf5\xb0\x98\x29\xcf\x6b\x06\xb3\x21\x79\x5b\xfd\xc6\xe8\xae\xba\xfd\xac\xd0\xad\xdb\x8a\xaf\xe7\x8e\x38\x09\xdf\xc5\xb5\xdb\x04\xb8\xc6\xf7\x4a\x3e\xa8\xdb\x9e\xe6\xbd\x6f\xa5\xcd\x03\x46\x33\x95\x37\xf4\x2e\xe0\x59\xc0\x1c\x70\x46\xa1\xcb\xd2\x18\xd0\x85\x55\x9c\xd1\xd5\xe1\xd0\xcd\x13\xd0\xfa\x0a\x00\x51\xbd\x27\x08\x49\x17\x0f\x22\x01\x43\x25\x24\xc4\x02\x37\x9a\xf9\x0f\x00\x46\xc4\x80\x25\xb3\x9f\x00\x95\x7b\x97\x07\x17\x2f\xb5\x66\xd3\xea\xd2\xe2\x19\xf4\x00\x79\x6b\xc5\xd3\xa6\x6c\x19\x33\x52\x96\xbd\x1c\xf4\x85\xc9\x83\x35\x73\x2c\xed\x76\x39\x11\xa5\x41\x75\x31\xf2\xe0\xfc\xda\x6c\x6c\x54\x4c\x17\x4d\x46\x0f\x86\xc6\xb8\x2e\x43\x6d\xfd\xaa\x54\x00\xb9\xff\x73\x9d\xeb\x24\xe8\x93\xff\x2c\x38\x15\x49\xf5\x2c\xb8\x15\x20\x2e\x71\xcb\xb3\x4c\xd9\xde\xf3\x38\xe6\x4f\xa3\xf5\x15\x43\x40\x4c\x44\x3f\x0d\x79\xc1\xb9\x47\x44\x09\x30\x00\xa2\x4f\xb9\x6b\x75\xd5\xf2\x1e\xca\xfa\x8c\x04\x67\xd4\x50\xc1\xfd\x31\xe5\xa2\x61\x1d\xf6\x76\x7f\xc6\x9a\x02\x18\xce\x00\x61\x88\xd2\x43\xad\x29\x2c\x4c\x25\xe4\x08\x58\x66\x9c\xa9\x59\xe0\xc4\x9a\x6d\xae\xd4\x3f\xbb\xa2\x4f\x92\x06\xe3\x19\xf8\x1d\xcb\xba\xd6\xae\x15\x84\xc1\x07\xc6\x33\xd5\xa6\xd4\x03\xeb\xf3\xc8\x25\x89\x49\x2d\x3f\xe6\x12\x17\x16\x55\xfe\x32\xe9\x46\x96\xa6\x81\x04\x54\x83\x30\xa6\x54\x1d\x35\xa7\x94\x33\x3b\x14\x5d\x80\xc6\x16\x8e\xe6\x38\x50\x5b\x55\xfa\xaf\x6c\xcf\x67\x41\x2e\x03\x07\xce\x37\xd7\xed\x95\x33\x4c\x54\x44\x16\x5d\x89\x13\x44\x6a\xf6\x3e\x6c\x99\x26\x5b\xc0\x58\xbd\x58\xd0\xb1\x2d\x7d\x59\x85\x63\xfd\xf5\x73\xbd\x96\x28\x6f\x11\x0e\x47\xf5\x62\x0a\xc8\x54\xc5\x14\x73\x16\x83\x7c\x56\xbf\x4c\xa9\xb3\x4b\x3e\xb1\xc0\x07\x6b\x4b\xdb\x43\xdf\x8b\x2d\x86\x11\x11\x60\x84\xa3\x5c\x7f\x2e\xff\xdd\xb0\xb3\x75\x70\xd9\x53\xe1\x6a\x65\x65\xb2\x23\x3c\xc8\x5b\xb6\x47\xf7\xef\xb8\xae\x8b\xcc\x79\x9e\xd1\x24\x21\xcc\x83\x59\xd5\x31\x6d\xa8\x9b\x54\x45\x9f\xcc\x03\xc4\x7a\x9b\xb8\xb1\xd6\x82\xfc\xcf\xfd\xff\x26\x2a\xb2\x4f\x73\x6e\xae\xc8\xb7\x56\x76\xd8\x69\xa2\xb4\x82\x96\x9a\x68\xab\xd8\x16\x16\x46\x58\x45\x83\x46\xb3\xaa\x40\x6e\x16\x37\x88\x52\x4e\xca\xdc\x2b\x7e\xcc\x3b\xc3\x2c\xdd\xae\x60\xe3\x5c\x90\x26\x5d\xca\xe2\x06\x9a\xdc\xb9\xad\xfd\x32\xbe\xfb\x8e\xee\x66\xca\x91\x4f\x6f\xe9\xe2\xb6\x05\x93\x0f\x3f\x1f\x5f\xff\xf3\xf1\xc3\xbf\x1c\x3f\xbc\xf3\x3d\xd4\x82\xbc\x28\x46\x2d\x4b\xb5\x9c\xdd\x4b\xaa\x1e\x79\x4d\xbb\xfb\xf8\xd6\xc1\xf8\xe0\x23\xcb\x85\x01\x4e\x02\x12\x59\x54\xe3\xbd\xfd\xc9\xbd\xab\x16\x95\x2c\x83\x3c\x40\xdb\x98\x25\x12\x03\x2b\xf5\x06\x9c\x6b\x24\xdb\xe5\x8a\xb6\xb0\x84\x3b\x1c\xba\xc5\xdd\x4f\xdf\xaa\xac\xcb\x55\xb3\xbc\x76\x3e\x1c\x0a\x6b\x26\xa9\xa0\xdd\x5d\x97\x0f\x82\x80\x70\xde\x40\xb9\x21\xaf\x5e\x1f\xbf\x7b\x03\x95\x00\x02\x27\x82\xa4\xb2\x4f\x38\x29\x1d\x57\xad\xb2\x30\x71\xeb\xee\x80\x15\xdc\xa5\x18\xc9\x29\x94\x8e\x76\xdb\xe5\x64\xcb\xb5\xa0\xbc\xe4\x34\x61\x55\xdf\x0b\xa9\x89\xad\x8a\xb1\x27\x1f\xff\x63\xb2\x7f\xb7\xde\xd8\x8f\x1b\x8d\x35\x49\xd6\x7a\x07\x58\x7b\x02\x55\x8a\x5b\xfb\x1c\x2a\xf2\x6b\xe9\x65\x42\x35\x75\x1b\xe1\x69\xbc\x66\xf8\xfa\x97\xd7\x8e\xee\xdf\x39\xfe\xe8\xad\xc9\xa7\xb7\x8e\x3f\x7e\x22\x63\x53\xed\xce\xda\xfa\x22\x9f\xe4\x5c\x6a\x5c\x32\xcf\x82\x4f\xec\x59\x25\x15\xa7\x21\x07\x1f\x2e\x5e\x2a\x46\xbb\x29\x83\x86\x9a\x4a\x42\xb2\xa3\x0e\xf2\xfc\xb1\x3d\x07\x76\xf9\xbb\xac\x9c\xec\xb9\xe7\x9a\x65\x38\x70\x01\xaa\xc3\xe5\xcf\x61\x70\x51\x2d\xbc\x64\xba\x86\x86\xdc\xcd\x06\xbc\xdf\x90\x0b\x17\x27\xd1\xf5\xca\x55\xa0\x92\x66\x4c\xf7\x0d\x61\x11\x47\x4f\x1a\xa6\x06\x64\xf3\x13\xca\xa9\x92\xf2\x7f\xd1\xa9\x5d\x7a\xbb\x59\x4d\x0e\xa5\x0e\x61\x5d\xad\x58\xdb\x69\x33\x2f\x51\x76\x10\x55\x05\xfd\x2f\x86\x9e\x9c\x18\xd2\xde\x83\x65\x63\xa8\x55\x8b\xa2\x6f\x14\x59\xfa\xd5\xc0\x69\x23\xeb\x24\xac\x9e\xf2\x7a\xfa\x2d\x84\xa1\xfe\xd7\xaa\x72\x9d\xe4\x4e\x87\x2b\xf4\xba\xdd\xaa\x5c\x6b\xae\xd9\xc2\x4c\x9d\x99\xe0\x1b\xe2\x10\xce\xb2\x88\x06\xea\x75\xf9\xea\x56\x12\xba\x69\x46\x92\x9d\x38\x92\x8e\xc0\x82\x3b\x69\xb7\x4b\x03\x12\xa6\xc1\x20\x26\x89\x70\x79\xc6\x08\x0e\x79\x9f\x10\x11\x47\xae\xfa\x46\x85\xf6\xb4\x0b\x66\x8b\x43\x61\x4b\x75\x21\x95\xd8\xa7\x7c\x5f\xcb\x7f\xe6\x99\x93\x89\x2c\x9d\xb6\xb3\xa9\x12\xaa\x8d\x5a\x84\x9a\x1d\x06\x1a\x0d\x31\xe1\x1c\xf7\x88\x4b\x18\x4b\x59\x03\x99\x4d\x58\xd5\x83\xbd\x79\x38\x3e\xdc\xb3\x61\xa1\xfb\x82\x75\x00\x5f\x26\xd5\x58\x81\x3f\x2f\x34\x0c\x7f\xa8\x9d\x7e\xb3\xa3\x43\xfd\x8b\xde\x09\x00\x9f\xdf\xfc\x3d\x63\x10\xdb\xf6\x69\xd6\x81\xf9\xe4\xd3\xac\x02\x5c\xb3\xbb\x5c\x53\x89\x32\xd2\x95\xc5\xb9\x9c\x75\x35\x69\x0d\xf8\x8b\xf7\x14\x0d\x69\x83\x16\x44\x94\x8b\x0a\xaf\x9a\xb7\x16\x32\x4f\xcb\x6f\x37\x4b\xb3\x46\xf3\x52\x85\x71\x92\x26\xa4\xd1\x1c\x6a\x7c\x8c\x5a\xba\x93\xad\x2f\xdf\xa6\xb2\xa5\x0d\x6b\x32\xb9\x68\xb4\xbe\xd2\x5e\x9d\x76\xb7\x57\xe0\x3f\x01\x00\x00\xff\xff\xd4\xba\x6b\x20\x24\x2c\x00\x00")

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

	info := bindataFileInfo{name: "vue.tpl", size: 11300, mode: os.FileMode(420), modTime: time.Unix(1560134799, 0)}
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
