// Code generated by "esc "; DO NOT EDIT.

package js

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/helpers.js": {
		local:   "pkg/js/helpers.js",
		size:    16101,
		modtime: 0,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "pkg/js",
	},
}
