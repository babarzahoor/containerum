// Code generated by fileb0x at "2018-09-17 15:58:28.397374915 +0300 MSK m=+0.012016018" from config file "b0x.toml" DO NOT EDIT.
// modification hash(3b92493472b25398905220a9bb44590e.4ef37f784b9c47297535a22bfedd9837)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileHelmignore is "/.helmignore"
var FileHelmignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\x8e\xc1\x6a\x23\x31\x0c\x86\xef\x7a\x8a\x7f\x99\xcb\xee\xb0\x78\x1e\x22\xd9\xc3\x9e\x5a\x48\xc9\xb5\x78\x66\x14\x5b\x89\x47\x36\xb6\x26\x69\x7b\xe8\xb3\x97\x24\x84\xf6\xf2\x81\x3e\x24\xf1\x75\x78\xf6\x66\x5c\xb5\xc1\x32\x24\x68\xae\x8c\x4b\x64\xc5\xb8\x4a\x9a\x45\x03\x8a\x9f\x4e\x3e\x70\x73\xd4\xe1\x25\x4a\x43\x5b\x4b\xc9\xd5\x1a\x5a\xe4\x94\x10\x52\x1e\xb1\x78\x9b\xa2\x68\xf8\x8b\xca\xc9\x9b\x9c\x19\xc5\x5b\xfc\xe1\xbd\xce\xd4\x41\x39\x78\x93\xac\xf8\x5d\x2a\x1f\xe4\x8d\x67\x5c\xc4\x22\x7e\xfd\x71\x78\xd2\xf4\x8e\xac\xb7\xcb\x6b\x12\x0a\x57\x24\x51\x76\xe4\xb6\xbb\xd7\x9d\xe5\xca\xd4\x61\x93\x97\x25\x2b\xf6\x9b\x1d\x66\xa9\x8d\x5c\x10\x1b\x6e\xbc\xe7\x93\x1b\x3f\xea\x70\xe3\x43\xc4\x30\x5c\xf1\x18\xdb\x59\x87\xef\x47\xa3\x9f\x4e\x6b\xc1\x41\x12\x37\xea\x5d\xbb\x14\xea\xdd\xe8\x4f\xd4\x3b\x5b\x0a\xf5\x9f\xd4\x61\xef\xab\xe4\xb5\xe1\xff\xf6\x5f\x23\x57\x6a\x3e\xf2\x64\xe4\x64\x66\x3f\xdc\xf7\x6a\x3e\xd2\x57\x00\x00\x00\xff\xff\xbc\x5b\x94\x77\x4d\x01\x00\x00")

// FileChartYaml is "/Chart.yaml"
var FileChartYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4a\x2c\xc8\x0c\x4b\x2d\x2a\xce\xcc\xcf\xb3\x52\x28\x33\xe4\x4a\x2c\x28\x80\x73\x95\xca\x0c\xf5\x0c\xf4\x0c\x94\xb8\x52\x52\x8b\x93\x8b\x32\x0b\x4a\xc0\xa2\x8e\x0a\x1e\xa9\x39\xb9\x0a\xc9\x19\x89\x45\x25\x0a\x69\xf9\x45\x0a\xce\xf9\x79\x25\x89\x99\x79\xa9\x45\xa5\xb9\x5c\x79\x89\xb9\xa9\x56\x0a\xc9\x48\x22\x65\x18\xa6\x01\x02\x00\x00\xff\xff\x63\xa0\xa4\x4b\x72\x00\x00\x00")

// FileRequirementsLock is "/requirements.lock"
var FileRequirementsLock = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\xd2\x4f\x6e\xf3\x20\x10\x05\xf0\xbd\x4f\xc1\xfe\x93\x09\xe0\x7f\x84\x73\x7c\x17\x18\x60\x6a\xa3\xda\xe0\x32\x90\x2a\xb7\xaf\xaa\x4a\x4e\xbb\x8d\xb3\x1e\xe9\xa7\x37\x7a\xcf\xe3\x8e\xd1\x63\x74\x01\xc9\x34\x2d\x8b\xb0\xa1\x61\x7b\xa2\x32\x67\xa4\x8f\xb5\x61\x2c\xe3\x9e\x28\x94\x94\xef\x86\x2d\xa5\xec\x64\x2e\x17\xb7\x40\x2e\xc4\x5d\x8a\x05\x42\xc4\x5c\x37\x1e\x52\xc3\xd8\x0d\x33\x85\x14\x0d\x13\x5c\x0e\x5c\x1c\xe0\x96\xe2\x9c\xbc\x7d\x56\xeb\xb8\xe0\xfd\x81\xc1\x1e\xda\x19\x0a\x7e\xc2\xfd\x59\xf0\x26\xb9\xe0\xdd\x21\x56\xc2\xdc\x6e\x10\x61\xc6\x7c\x8a\x7c\x84\xcc\x48\xa9\x66\x87\x2f\xe2\x76\xcc\x5b\xa0\xef\x13\xbd\x48\x7c\xaf\xf6\x5c\xb8\xf1\x51\x48\x2d\xcb\x29\x4a\x3d\x86\x02\x61\x7d\x11\x55\xc3\xb9\xf7\xda\xec\x7e\x69\x31\x79\xdc\xb0\xe4\xe0\x9e\x2e\xe0\x6f\xbc\x5b\x5a\xeb\x86\x67\x2c\x79\x58\x94\xd6\x5a\xce\x6c\xe3\x87\xf3\x61\x46\x2a\x86\xd1\x02\x6a\x18\x4d\x67\xb5\x77\xe3\xa4\xc0\x69\xd5\x79\xeb\x46\xe5\x85\xd4\x30\xf5\x93\xb5\x12\xa4\x95\x7e\xf2\x0a\xbc\x83\x69\x90\xbd\xd5\xca\xa1\xd4\x4a\x8f\x6f\x57\x40\xe5\x40\xa9\x66\xc6\x88\x19\x0a\x7a\xc3\x94\x90\xba\x15\xba\x55\xf2\xbf\xec\xcd\xa0\x4d\x77\xe5\xa2\xbf\xea\x5e\x0e\xc3\xf4\x4f\x74\x46\x88\xe6\x2b\x00\x00\xff\xff\x07\x1f\x35\x60\x8c\x04\x00\x00")

// FileRequirementsYaml is "/requirements.yaml"
var FileRequirementsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xbc\xd3\x41\x6e\xea\x30\x10\x80\xe1\x3d\xa7\xb0\xd8\xc7\x84\x07\xaf\x8b\xdc\xc6\x38\x43\xb0\x88\x3d\xe9\xcc\x98\x8a\xdb\x57\x89\x2a\x17\xa2\x98\x45\xeb\x86\x15\x58\x24\xff\x27\xdb\xd3\xc2\x00\xa1\x85\x60\x1d\x70\xb3\x51\xaa\x52\xc1\x78\x68\xd4\x80\x2c\x1d\x01\xbf\xf7\x1b\xa5\x94\x22\x18\x90\x9d\x20\xdd\x1b\x75\x11\x19\xb8\xd9\xed\xec\xc5\x90\xb0\xb6\x18\xc4\xb8\x00\x14\xbd\x76\x38\xfd\xfb\x06\xc4\x0e\x43\xa3\x6a\xbd\xff\xaf\xeb\x69\x4d\x4c\x37\x05\xc6\x4f\xa5\xda\x53\xfa\x3a\x2b\x4d\x4b\xbd\x91\x33\x92\x7f\xf0\x78\x0c\x1d\x7e\x3d\x95\x5e\xbf\x3d\xe8\x5a\x1f\xb7\x3f\x10\x66\x35\x53\xe7\x15\xc4\x0c\xae\xea\x8c\xc0\x87\xb9\xcf\x30\xb7\xbd\xae\xf5\xe1\xf7\x9a\x93\xb1\xd7\xf4\xe3\x31\x95\x11\x45\x06\xaa\xbc\x09\xa6\x03\x5a\x22\x1d\x0b\x93\xc6\xde\x2b\x0f\x01\x63\x24\x0b\x6b\x58\x9e\x5a\x19\xcf\x00\xe4\x1d\x8f\x0a\x5e\x83\x34\xcf\x65\x54\xd7\x78\x5a\xdc\xa1\xb7\xc2\x9c\xd4\xc9\xdd\xe7\x28\x97\x25\xc7\xbf\xc2\x8e\xd4\xc9\x0d\xb8\x71\xfd\x1a\x8e\xd4\xc9\x4d\x93\x5b\x3c\x95\x8a\x6c\x09\xca\x99\x30\xc8\xf7\x24\xb9\x57\x92\x80\x2d\x78\x10\x72\x76\x7e\x6f\x0b\xed\x8b\xc7\x30\x3e\xe9\x42\x97\x96\xe6\xcd\x0c\xed\x86\x7d\xf4\xb0\xa0\xda\x17\x3e\xad\x87\x50\x86\xc2\xd8\x47\x59\x98\xed\xbf\xd0\x3c\xb7\x9e\x40\x9f\x01\x00\x00\xff\xff\xc7\x17\x8e\x44\x48\x07\x00\x00")

// FileTemplatesHelpersTpl is "/templates/_helpers.tpl"
var FileTemplatesHelpersTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x93\x41\x6b\xdc\x30\x10\x85\xef\xfe\x15\x0f\xd1\x40\x9b\xb2\xce\xa1\xd0\xc3\x42\x4e\x69\x0f\xa5\x90\x42\x03\xe9\xb1\xc8\xf6\xa8\x1e\x90\x65\x57\x23\x6d\xb3\x24\xf9\xef\x45\x92\xe3\xdd\x2d\x6c\xb3\x7b\x1b\xe4\xa7\x37\x6f\xbe\x91\x1f\x1f\xaf\x2e\xb1\xe1\x61\x0d\xa1\x00\xc3\x96\xc2\x76\xa2\xeb\x21\x4a\xd0\x6d\x4f\x6b\x5c\x5e\x3d\x3f\x57\x49\x55\x7d\x7e\x98\xb4\xeb\x10\x7a\x82\xd3\x03\x61\x34\xb9\x6e\x7b\xed\x43\x5d\xcd\xba\x15\x3a\x32\xec\x08\xaa\x1d\x5d\xd0\xec\xc8\xc7\xa1\x4e\x7a\x85\xd5\x4e\xa1\xa3\x0d\xa8\x6f\xf2\xd5\xdb\x64\x56\xdf\x6b\x1b\x49\xb2\xf2\xdb\x86\xbc\xe7\x8e\xf0\x84\xe0\xa3\x6b\xf1\xf1\x43\x2e\x79\xb8\x8b\xc6\xf0\x03\xd4\x6a\x67\x46\xae\xcb\x75\xc9\x78\xe3\x49\x07\x82\x5e\x7a\x98\x68\xed\x16\xbf\xa3\xb6\x6c\x98\x3a\xe8\x69\xca\xe9\xeb\xea\x07\x15\xf7\xac\x0f\xa9\x47\x9a\x44\xd0\x50\xab\xa3\x10\x64\x1c\x08\x5f\x63\x43\xde\x51\x20\x29\x33\x1b\x26\xdb\x09\xb4\x27\x58\x1e\x38\x50\x87\x30\x22\xf4\x2c\x78\xdb\x6c\x33\x8f\x4f\xb7\x77\x49\xcb\xee\x17\x64\xa2\xf6\x5d\x5d\x7d\x31\xf0\x64\x49\xcb\x0c\x6e\x26\x23\x05\x5d\x39\xe3\x80\x3f\x6c\x2d\x1a\x42\x94\x94\x53\xa0\x73\xf8\x39\xed\x7f\xf1\x26\xdd\x21\x62\x36\x0b\xd1\x97\x8f\x0b\xd5\x17\xcd\x51\xc1\x49\xd8\xad\xec\x9c\xde\xe4\x11\xd6\xd7\xa7\x6f\x76\x2f\xe7\x42\xa3\xb8\xd4\xdf\x0b\xaa\x72\x79\xc9\x7a\x70\x7a\x76\xc0\xc9\xb3\x0b\x06\xea\x42\x56\x17\xa2\xfe\x71\x2b\x7d\xcf\x79\x6b\xc7\xea\x83\x37\xb8\xb7\xdc\xf4\xdb\x6c\xc8\x0b\x8f\x2e\x2d\x36\x2f\x78\x7e\x2d\x45\x65\x75\x43\xf6\x95\x25\x67\xa5\x3a\x3a\xd2\x3e\xf1\x52\xdf\xcf\x1d\x9f\xe0\x69\xb2\xba\x25\xa8\xf7\x0a\xea\xa7\x3a\x6b\xd6\xbf\x01\x00\x00\xff\xff\xb1\x49\x54\x02\x21\x04\x00\x00")

// FileTemplatesIngressYaml is "/templates/ingress.yaml"
var FileTemplatesIngressYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x95\x31\x6f\xdb\x30\x10\x85\x77\xfd\x8a\x43\xd0\x95\x6c\x53\x74\x08\x08\x74\xe8\x98\xc5\x30\x5a\x34\x40\xc7\xb3\x74\xb1\x88\x50\x14\x41\x9e\x92\x16\xae\xfe\x7b\x71\xa4\x5d\x4b\x6e\xe4\x66\xa8\x87\x02\xf1\x24\x1e\xbf\x77\x38\x3e\xf2\xc1\xbb\x9d\x82\x37\xf7\x83\x73\x2b\xec\x08\xcc\x47\xd0\x9f\xc9\x11\x26\xd2\xb9\xa0\xc6\xb1\xca\x88\xf5\xdb\x48\x29\x7d\xbd\x5d\x23\xb7\x99\xbb\x43\x37\x50\xd2\xfb\x0d\x3d\x58\x1d\x64\xeb\x54\xf1\x69\xbd\x28\xc1\x30\xd1\x64\x91\xbd\x7f\xae\x2f\x79\xdc\x38\x6a\x32\x86\xc1\xde\x51\x4c\xb6\xf7\x06\xe8\x3b\x93\x97\xcf\xf4\xf6\xf1\x7a\x43\x8c\xd7\xd5\x83\xf5\x8d\x81\xdb\xa2\xad\x3a\x62\x6c\x90\xd1\x54\x00\x1e\x3b\x32\xb0\xdb\x4d\x4e\x3b\x8e\x6a\xb0\x15\x80\xc3\x0d\xb9\x24\x10\x00\x86\x90\x29\xa6\x2e\x38\x64\x82\xab\xba\xf7\x8c\xd6\x53\x1c\x3a\x2d\x4d\xae\x40\xc3\x38\x66\xb8\x6e\x31\xf2\x19\x3c\xef\x4f\xf8\x58\xac\xcd\x8a\xb9\xcf\x7b\xa0\xa5\x68\x19\xb7\x27\xc4\x17\x8a\x8f\xb6\xce\x90\x98\xf4\x64\xb9\xfd\xd3\x4b\xef\x7b\x46\x16\x33\x4a\xb3\x49\xc1\x54\x32\x61\xff\x0d\x3b\x07\x1a\x7e\x82\xf5\x0d\x79\x86\x0f\x87\x86\xe4\x1b\xf9\x4c\x81\x6a\x73\xe6\x1a\xd8\xed\x5b\x73\x31\x4b\xc8\x88\x7e\x4b\x67\x61\x00\x05\x6d\x9f\x78\xef\xef\x4c\x96\xeb\x07\xac\xa0\x72\xee\x63\x65\x32\x5d\x29\x24\xaa\x23\xf1\xea\x70\x97\xfa\xb8\x2e\xcc\x44\x30\xd3\xc6\xc1\xd1\xdf\x87\x9e\xcd\x53\xc6\x36\x27\x13\xb5\xcc\xc1\xfc\x9e\x57\x5e\x6f\x3a\x2e\x45\x24\xa5\xf2\xce\xe6\x91\x99\x9c\x52\x7e\x1b\xac\x1f\xc8\x37\x66\x56\x94\xf3\xe5\xab\x5e\x2d\x3f\xd6\x67\xe8\x75\x2f\xaf\x50\x26\x5b\x74\xa0\x52\x4a\x2d\x46\x4c\x72\x78\xc1\x8c\x61\x78\x0d\xd9\xcb\x42\x26\x17\xf1\xf2\x94\xcd\xe8\xff\x26\x66\x32\xf5\x45\x72\x76\xf8\xa3\xf9\x07\x41\xc3\x60\xd5\x16\x99\x9e\xf0\xc7\xb9\xc4\xdd\xbc\xbb\x79\xbf\x68\xc6\xaf\x00\x00\x00\xff\xff\x89\x56\x9d\xb1\x59\x07\x00\x00")

// FileValuesYaml is "/values.yaml"
var FileValuesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x56\x4b\xcf\xa3\x36\x14\xdd\xf3\x2b\xae\x26\x8b\xae\xc8\xec\xbd\x23\xc1\x4d\x50\x79\x89\xc7\xb4\xa3\xaa\x42\x0e\x71\x08\x2a\xd8\xd4\x36\x33\xf3\xa9\xea\x7f\xaf\x70\x80\x84\x00\x33\xd1\x74\xd1\xd9\x44\x91\xef\xb9\xaf\x73\x8f\xaf\x51\xa4\x90\xc8\x00\x38\x9f\x10\x5c\x48\x25\xa9\x01\xd0\x54\x44\x5d\xb8\xa8\x11\x28\xd1\x3e\x1e\x74\xa0\xfe\xa8\xe6\xac\x54\x5c\x94\xac\x18\xfc\x8c\x92\x15\x82\x4a\x1d\xad\x2d\xbb\x5f\x00\xca\xc8\xa9\xa2\xe7\xd1\x0b\x80\x30\xc6\x15\x51\x25\x67\x12\xc1\xdf\xff\xe8\x33\x80\x0d\xfc\xd9\x9e\xa8\x60\x54\x51\xb9\x2d\xf9\xfb\x3e\xd4\x36\xaf\x88\x94\x08\x58\x51\xb2\x2f\x2b\x50\x55\x49\x93\xe4\x35\x45\xf0\xae\x4b\xf2\x4e\xc3\x1a\xa2\xae\x08\xde\xeb\xff\x57\x2e\x95\x44\xbd\xb7\x09\x15\xcf\x49\xb5\xcd\x39\x53\xa4\x64\x54\xb4\xf5\xb6\xe4\xda\xa8\x2a\x89\xe0\xf7\x3f\x8c\x5b\x12\x30\x41\xd2\x5c\x50\xe5\x93\x2e\x76\x7e\x25\x42\x99\xf4\x0b\xa9\x9b\x8a\x9a\xaa\x92\x03\x6c\x12\x7f\x33\x24\x99\xc0\xb7\x3a\xa5\x01\x40\x9a\x1f\x97\x16\xd2\x94\xdb\xff\x8d\x1a\xc3\xb8\x09\x86\xb2\x4f\x37\xb4\x3e\x1e\x8a\xdb\x6c\xc0\xa6\x17\xd2\x56\x0a\x5a\x49\x05\xe4\xbc\x65\x4a\xbc\xf5\xd6\x7d\x90\xfa\x49\xf4\x11\x41\x1a\xdf\xf1\x17\x2e\xe0\x98\x24\x61\xdc\x79\x20\x90\xb2\xea\x6d\x56\xe8\x64\x61\x14\x24\xc1\x3e\x70\xb3\xe4\x63\x88\x11\x30\x6e\x3e\xd9\x83\x28\x41\xc0\xda\x6a\x38\x74\xad\x04\xc7\x49\x16\x61\x17\x5b\x31\x7e\x30\x15\x15\x3f\xdd\xeb\xec\x7c\x8f\x41\x9c\xa0\xaf\x93\x09\x10\xe1\xbd\x15\x26\xfb\xa3\xd5\x87\x32\x6a\x52\xea\x28\x0d\x15\xb2\x94\x8a\xb2\x9c\x3e\x09\x65\xb8\x9a\xcb\x0c\xed\x8f\x99\x67\x39\x6e\x96\xc6\x38\xca\x3c\xcb\xb7\x0e\x38\xca\xd2\xc8\x9d\x74\x31\x80\x62\x2f\x09\xb3\xd0\x8a\xe3\x5f\x83\xc8\x5e\x6d\x66\x82\xb6\x6c\x3b\x5a\x8f\xe5\x06\x07\xc7\x5f\x36\x63\xdf\xd6\x15\x75\xff\x1d\x2f\x74\xf1\xb7\x60\x03\x23\x06\x69\xd5\xf5\x15\x4a\x8c\x4e\x12\x66\x4d\x18\x29\xa8\xe8\x10\xc3\x42\x5b\xdc\x5a\x23\x7d\xb3\x66\x35\x75\xe1\x21\xb3\x77\xbe\xe5\x61\x04\x0f\x83\x9b\x83\xfa\x8e\x1b\x2e\x55\x77\x21\x17\x07\x32\x60\x97\xb8\xeb\xe7\xd4\x4d\x6c\x3e\x24\x6d\xb4\xd2\xe4\x98\x1d\xa2\x70\xbf\xee\xaf\x21\x9d\xc8\xd7\x21\x21\x8e\x3c\x27\x8e\x9d\xc0\x8f\x67\x99\x36\x1b\x28\x2f\xf7\x60\xb6\xe7\xf8\xa3\x2a\x34\x0e\xd4\x95\x32\x20\xe7\xba\x64\x70\xe6\x54\xb2\x9f\x14\xe4\x82\x12\xd5\x37\x0c\x2b\xce\x08\x3e\x51\xf1\x26\x95\xe0\xac\x68\x88\x94\x9f\xb9\x38\x1b\x86\xe4\x55\x7b\x5b\x6f\xdf\x3f\xa3\x38\x70\xd3\x44\x37\xf3\xad\x41\x4d\x90\xaf\x4c\x6b\xe2\xb0\xc4\xe7\x04\xb0\x70\x7b\x9e\x40\x11\x8e\x83\x34\xda\xe3\xa5\x01\xdf\x51\xbf\xa4\x3b\x9c\x75\x6b\xe3\x19\x65\x18\x0d\x15\x75\x29\xe5\x7f\x61\xcc\xde\x65\x3b\xbd\xaf\xe6\x14\xd9\x3b\x3d\xb9\xaf\x91\x62\xef\xfa\x65\xf6\x50\x96\x96\xdc\x8c\x9d\x5e\x03\x4f\xa7\x63\x6f\x33\xcb\xc8\x4d\x8c\xa3\x0f\xce\x1e\xcf\x11\x3b\xc7\x75\x1d\x7f\x61\x10\x77\xea\x66\xa6\x0f\x81\x9b\x7a\x78\xdc\x7e\x0f\x76\x43\x50\xc9\x5b\x71\x5b\x20\xdf\x29\xbd\xb1\x66\x2f\xf0\x0f\x41\x66\xef\x56\x84\xf7\x84\xeb\xa5\x57\x73\x56\xf0\x45\xdd\x3d\xe1\x97\x94\x37\x42\xd6\x19\x7d\x44\x3d\xde\xf9\x47\x12\x18\x3f\xd3\x9a\x2a\x51\xe6\x72\xfd\xa9\x0d\xa3\xc0\xc3\xc9\x11\xa7\x53\xdf\x4f\xbc\x6a\x6b\xfa\xe3\xe8\x70\x9d\x89\x05\xe5\x18\xa4\x29\xcd\x82\x28\xfa\x99\xbc\xad\x77\x3e\x68\xb1\xcb\x95\x85\x11\xfe\xd9\xf9\x0d\x81\xa0\x15\x25\x92\x9a\x8c\xd4\xb4\xc7\xdd\x76\xf2\x70\x0f\x70\x1c\x8f\x59\x5e\x7d\xb0\x86\x56\xff\x7a\xf1\xc9\x1f\xf0\x36\x51\xe4\x44\xba\xcf\x99\x29\x8f\x9b\x11\x91\x4a\x2a\x26\x54\xde\x4d\x61\xbf\x88\x11\xdc\x57\xb2\xa1\x65\x79\x3e\xbd\x56\x46\x0f\x5e\xab\xa2\x37\x77\x25\x30\xfd\x59\x38\x68\x7e\x33\x98\x16\x4a\xf8\x37\x00\x00\xff\xff\x0b\xf7\x6e\xb5\x7d\x0c\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	err = FS.Mkdir(CTX, "/templates/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileHelmignore)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "/.helmignore", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileChartYaml)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "/Chart.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRequirementsLock)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "/requirements.lock", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileRequirementsYaml)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "/requirements.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileTemplatesHelpersTpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "/templates/_helpers.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileTemplatesIngressYaml)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "/templates/ingress.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileValuesYaml)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "/values.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
