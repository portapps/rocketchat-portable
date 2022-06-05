// Code generated by go-bindata. DO NOT EDIT.
// sources:
// res/Rocket.Chat.lnk (1.776kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _rocketChatLnk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\xc1\x4b\x14\x7f\x18\xc6\x3f\xfb\x73\x7f\x98\x10\x61\x20\x81\x50\x39\x44\x1b\x45\x38\x69\x91\xe0\x56\xa0\xcd\x6c\x4e\xac\xe6\xe0\x2a\xa8\x19\xb4\x4d\xdf\x72\x70\xc6\x99\x66\x26\x71\x8b\xba\x44\x17\x31\x08\x3c\x74\x29\x95\x08\x0a\xa1\x83\xe0\x39\x82\xfa\x03\x3c\x18\xd4\xa5\xb2\x2e\x21\x81\x87\x8e\x41\xc5\x8c\xa3\x99\xed\x2a\x75\x2b\x7a\x0e\xdf\x77\x77\x9e\x77\x9e\xe7\xf9\xbe\xbc\x4c\x2b\x90\xa8\xfa\x8f\x10\x4f\xa3\x93\x13\x6f\x9a\x40\x02\xbe\xcc\x8c\xe4\x7b\xec\xb9\x44\x72\xf4\xda\xb9\xb0\x32\x3e\x35\x7e\xd1\x9a\x4b\x70\xb4\x3c\x19\x36\x26\xf8\x11\x5f\xa9\xa2\x46\x7f\xdb\x3e\x2b\x2d\xa4\xcd\xca\xfb\x2f\x37\xb1\xbf\xae\xee\x5e\x35\x07\x94\x74\x1f\x3f\xe3\x3c\xf5\x51\xed\x6d\x3b\x39\x56\x49\x47\xbb\x92\xcd\x74\x5e\xaf\x97\x15\xad\x19\x34\x2a\x48\xb2\xf8\xa4\xb7\xad\x65\x2c\xe4\x65\xc0\x8f\xf2\xdd\x58\xa3\x22\x35\x3b\x74\xe0\x60\x30\x80\x20\x40\x46\xa1\x9f\x3c\x01\xb0\x1d\x8b\x83\x84\x71\x7b\xda\x02\x4b\xfa\xee\x91\xe9\xce\x80\xbe\xe2\xa1\xad\x78\x4c\x2f\x86\x9a\x5e\x91\xb4\xc5\x3d\x64\x04\xc3\x88\xc8\x0b\xba\xe2\x9a\x88\xeb\xe1\xf8\xdd\x4e\xa0\x1a\x28\x03\x1a\x5a\xde\x27\x2b\x01\xd5\xdc\x7d\xe9\xb2\x90\x14\x94\x74\x5f\x87\x63\x0c\x88\x40\x56\xfa\xf3\xc1\xea\xdf\xb2\x18\x16\xa0\x94\xb0\x96\x38\x15\x55\x93\x21\x04\x12\x0a\x1e\x0e\x3e\x3e\xb5\xe8\x58\x11\x77\x01\x07\x0f\x1b\x09\x15\x81\xcf\x00\x01\x0e\x2e\x12\xcd\xb8\xb8\x58\x98\x18\xb1\x8a\xc3\x20\x12\x43\x98\xe4\x91\xc8\x60\x21\x30\x08\x22\xd5\x41\x64\x0e\x21\x23\xd3\xf7\x0b\x67\xf1\xdc\xa5\x9e\x2f\x8f\x72\x0b\x0a\xe9\x92\x5d\x3b\xd7\x65\x37\xd2\xae\x2a\x83\x72\x98\x4c\xe5\x0a\x7e\x20\x6c\xd5\x33\x87\x44\x6a\xfd\xf1\xff\xb5\x48\x91\xa3\x80\x4f\x80\xc0\x46\xc5\x8b\x37\x29\xf5\xdb\xd3\xfd\x87\x3f\x11\x67\x97\x3e\x4c\x93\xdd\xf1\x7f\xc3\xcb\x5f\x29\xd4\xba\xc6\x32\x7f\xd7\xdf\xf5\xf9\xf9\xbb\x33\xd9\xa9\x3b\x0d\x7b\x8e\x2c\xe8\x13\xee\xeb\xf1\x67\xc7\x5f\x7d\xd8\x7a\xf3\xb6\x3a\x72\xda\x9e\xf9\xb8\x11\x3f\x02\x54\xc0\x64\x06\xa8\xcf\xe9\xb9\xae\xbd\xad\x13\x85\x89\xc6\xec\xa3\xd9\xf9\xb9\xda\xf9\x17\x9f\xf6\x01\xff\x87\x46\x35\xc0\xe6\x30\xc0\xca\x5a\x79\xab\xd6\x6d\x09\x8d\xb1\xc8\xf4\x36\x5b\x7d\x7c\xcb\xd5\x1e\x6a\x4d\xf2\x83\x63\xc3\xa3\x3b\x80\xfe\xb0\x41\x0b\x6f\x73\x35\xab\x2d\xb5\xeb\x35\x6b\x6f\xfb\x2d\x00\x00\xff\xff\x19\x47\xcb\x34\xf0\x06\x00\x00")

func rocketChatLnkBytes() ([]byte, error) {
	return bindataRead(
		_rocketChatLnk,
		"Rocket.Chat.lnk",
	)
}

func rocketChatLnk() (*asset, error) {
	bytes, err := rocketChatLnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Rocket.Chat.lnk", size: 1776, mode: os.FileMode(0666), modTime: time.Unix(1540578378, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb1, 0xac, 0x32, 0xcb, 0xe8, 0x67, 0x9a, 0x8e, 0xd7, 0x49, 0x29, 0x65, 0x32, 0xe4, 0xcf, 0x5e, 0x31, 0xc4, 0x4e, 0xb3, 0x98, 0x33, 0x81, 0xae, 0xa7, 0x43, 0x82, 0x80, 0xee, 0xde, 0x7c, 0x93}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"Rocket.Chat.lnk": rocketChatLnk,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"Rocket.Chat.lnk": {rocketChatLnk, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
