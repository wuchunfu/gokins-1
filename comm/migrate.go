package comm

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _mysql_000001_gokins_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\xcd\xaa\xc2\x30\x10\x46\xf7\xf7\x29\xf2\x1e\x5d\x5d\xb1\x42\x41\x50\x6c\x17\xee\xf2\x63\x87\x1a\x9b\x64\xc2\x4c\xda\xe7\x17\x04\x97\xce\xac\xcf\x99\xf9\xce\xf1\x76\xb9\x9a\xe9\xff\x70\xee\xcd\x70\x32\xfd\x7d\x18\xa7\xd1\xb8\x66\xc3\x16\xd3\xec\xba\xbf\x5f\xfc\x91\x67\x9b\x62\x01\x41\x79\x22\xae\x02\x7e\x61\x10\x68\x06\x66\xbf\x48\xef\x6b\xac\xa0\x14\x7c\x15\xbb\x03\x71\xc4\x22\xa8\xdc\xe4\xb5\xdd\x53\xf4\x21\x89\x6b\x9e\x7c\x16\x38\xa1\x7c\x0e\x94\x59\xa9\xdc\x18\x48\xc1\x36\xf3\xa2\x29\x04\x15\x55\x47\xae\xfd\x38\x0d\x57\x28\xae\x7b\x07\x00\x00\xff\xff\x07\x1c\x99\xbc\x44\x02\x00\x00")

func mysql_000001_gokins_down_sql() ([]byte, error) {
	return bindata_read(
		_mysql_000001_gokins_down_sql,
		"mysql/000001_gokins.down.sql",
	)
}

var _mysql_000001_gokins_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdf\x6f\xdb\xd6\xf5\x7f\xf7\x5f\x71\xdf\x64\x03\x16\x40\x52\x52\x6c\xb3\x4f\x6e\xa2\xef\x17\xc6\x6c\xb9\x50\xe4\xad\x7d\xa2\xae\xc4\x2b\x99\x89\xf8\x03\xbc\x97\x86\xfd\x66\xa3\x5b\xeb\xac\x5d\xed\x6c\x4e\xd2\xac\xcb\x96\x14\xc5\xd2\x01\x5d\x1d\x0c\x99\x87\xca\xde\xf2\xc7\xd4\xa2\xec\xa7\xfe\x0b\x03\x2f\x49\x89\x94\x28\xf1\x5e\x4a\x7d\x5b\x1f\x52\x58\x3a\xe7\xf0\xc3\x73\x3e\xf7\xfc\xba\xca\xe7\x41\x7e\xca\x7f\x0b\xf9\x3c\xa8\xc1\x46\x07\x01\x4c\x6c\xa7\x49\x1c\x1b\x81\x96\x69\x03\xa2\x34\x1c\xad\xa3\x2e\xa4\xa9\xdf\xad\x96\xd7\x6b\x65\x50\x5b\x7f\x7f\xb3\x0c\xea\x81\x56\x1d\x2c\x2e\x00\x50\xd7\xd4\x3a\xd8\x83\x76\x73\x17\xda\x8b\x77\x8a\x4b\xa0\xb2\x5d\x03\x95\x9d\xcd\xcd\x65\xef\x4b\x4b\xb3\x50\x47\x33\x90\x32\x26\xb5\xb3\xb9\x09\xee\x95\xff\x6f\x7d\x67\x33\x49\x7c\x0f\xd9\x58\x33\x0d\x46\x35\x4c\x20\x71\xf0\x50\x52\x14\x84\x04\x51\x70\x77\x7b\x6b\xab\x5c\xa9\x81\x9c\xfb\xe7\x5f\xf7\x2e\xbb\xfd\xdf\x5e\xb8\x87\x47\x39\x6a\x01\xd9\xb6\x69\x0f\x0d\x94\x52\x0c\xdc\x9e\x3d\xbf\x39\x3f\xbf\x7e\xf7\xca\x3d\x3a\x0f\x0c\xec\x21\x83\xb0\x23\xb8\xee\x7e\x76\x7d\x79\xe1\xab\x12\x4d\x47\x0a\x26\x50\xb7\xea\x40\x85\x04\x79\x7f\x2f\xa6\xbc\xc0\xa3\xd7\x37\xaf\x3e\x77\x9f\x5d\xdc\x3e\x79\x17\x1a\x21\x1d\x34\x7c\xbe\x54\x2a\x4d\x37\xf0\xf2\xd3\xdb\xaf\xbf\xf4\x55\x75\x84\x31\x6c\xf3\x28\x53\xf7\x45\xdf\x1e\x13\x68\x13\xa4\x32\xe3\xef\x5d\x1d\xf6\x5e\x7f\xe6\xe1\x7f\xf6\xd6\xb7\xd0\xd2\x0c\x0d\xef\x72\x98\xe8\x5f\xfe\xc1\x7d\xf1\x97\xa8\x89\xa6\x8d\x20\x17\x88\xe3\xaf\x7a\x97\xdd\xa8\x05\xc7\x52\xb9\x2c\xb8\x5f\xbd\x75\x9f\xbe\x89\x5a\x08\x78\xcb\xee\xcb\xfe\xa3\x63\xf7\x4f\xdf\x51\xe5\x0f\xaa\x1b\x5b\xeb\xd5\x8f\xc0\x2f\xca\x1f\x81\x45\xef\x58\x2d\x81\x9d\xfb\x1b\x95\xff\x07\xef\xd7\xaa\xe5\xf2\xc2\xd2\x7b\x63\xc7\xb0\xa9\xab\x8a\x77\x5c\x58\x4e\x62\xdb\x36\x1d\x8b\xf1\x3c\xd1\xe3\xcd\x28\xfb\xc0\x6c\x64\x3c\xa5\xa5\x24\xe7\x52\x49\xc3\xd1\xeb\x40\x33\xc8\xa2\x28\x4e\x12\x69\x9a\x06\xa1\x27\x8e\xa0\x7d\x12\xf9\x98\x89\x03\x1c\x9c\xe5\x21\x67\xb6\x20\xee\x9a\xe6\x43\x96\x00\x92\x03\x2b\xf5\x8c\xfa\x2f\x66\x40\x0b\xef\x9a\xa4\x0e\x3a\xa6\xd1\x8e\xfb\x67\x34\x06\x53\xcc\xe8\xb8\xcd\x24\xe7\xe1\x57\x58\xc1\xa5\xbb\x87\xa1\x92\x55\x51\xd3\xb4\x55\x0c\xcc\x16\xf0\xdd\xc7\xa2\x94\xad\x3c\x3e\x30\x1b\xdc\xc5\xf1\x81\xd9\x60\x89\x27\xcf\x21\xc3\x04\xb6\xe3\x35\x34\xb5\xc4\xfd\xf3\xc8\x7d\xf3\xb6\xdf\x7d\xa7\xa9\x7e\x66\x52\x35\x6c\x75\xe0\x81\x62\x40\x9d\x8d\x46\xfc\xb5\x78\xca\xd3\xa9\x4f\x58\xb8\x3b\x7b\x25\xdf\xd7\x88\xd2\x34\x55\x54\x07\x0d\xad\xed\x25\x11\x29\xa5\x96\x1f\x1e\xf6\x3e\xed\xf6\x5f\xce\xab\x13\x88\x3b\x38\xed\x05\x7a\xa7\xbf\xeb\xfd\xfd\xd9\xff\xaa\xe8\xdc\xaa\xa8\x17\x40\xad\x6d\x98\x76\x24\x08\x13\xe9\x66\x38\x7a\x03\xd9\xd3\x99\x12\x94\x1b\x5d\x87\x86\x8a\x47\xeb\x8d\x8a\x2c\x64\xa8\x58\xf1\xc0\x3e\xc0\xa6\x31\xfc\x46\xd3\x19\x7a\x2a\x1f\xb0\xb1\xa7\xd9\xa6\xa1\x23\x83\xe0\x51\x2b\xd8\xb4\xc9\x00\x9e\x98\xbd\xe8\xf0\x66\x55\x96\xbc\x97\x39\xa9\x06\x0d\x27\x77\x62\x1d\x34\xaa\x83\xe4\x1a\x8d\x5b\x90\x5b\xc1\xfa\x4e\x6d\x5b\xd9\xa8\xdc\xad\x96\x3d\x6e\x50\x2f\xee\x4f\x4d\xc4\x0e\x47\x62\xeb\x9d\x3c\xbe\x3d\x3c\xba\x39\xfc\xcd\x4f\x57\xc7\xbd\x93\xf3\xfe\xdf\xba\x3f\x5d\x3d\xe2\xe8\xc0\xe3\xbd\xcb\x78\x7d\xf6\xea\x28\x53\x8b\xc4\xd1\xe8\x68\x46\xcb\x1c\x23\xae\x63\x77\xd2\x92\x5c\x22\xb3\x96\x7d\x6f\xce\x4c\x30\x56\x0e\x64\x26\x99\x05\x6d\xa8\x73\x53\x8c\x6a\xf1\x13\x8c\x25\xe7\x73\x92\x44\x85\x04\x8e\x06\xcd\x8b\x33\x9e\x4b\x13\xca\x1b\x2e\x36\x6f\x66\x0f\x16\xb2\x75\xec\xa5\x7c\xfe\x80\x85\x9a\xf3\xce\x0a\x16\xb4\x63\x13\xfd\xe4\xe6\x6c\x24\xa6\x53\xa2\xbf\x07\x3b\x0e\x2b\x4f\xa6\x47\xfa\xee\x4e\xb5\x5a\xae\xd4\x94\xda\xc6\x56\xf9\x7e\x6d\x7d\xeb\x83\x45\x61\x29\x52\x30\x12\x67\x27\x51\x60\x3e\xd1\x9e\xe0\x46\xe5\x5e\xf9\x43\x50\xdf\xb8\xf7\xa1\x82\x0f\xf0\xd0\xd1\x4a\xe0\x99\xc5\xd0\x45\xb3\x93\x8b\x39\xfa\x1b\x95\xfb\xe5\x6a\x0d\x6c\x54\x6a\xdb\xd4\x4b\x91\xf0\x2f\xfc\x72\x7d\x73\xa7\x7c\x7f\x01\x50\x1e\x00\x20\x2e\xd3\xff\xe5\xbc\xca\x6d\x1a\x39\xff\xaf\xd0\xbf\xc0\x6b\xfb\xfe\xd8\x3f\xfb\x36\x97\x24\x95\x93\x04\x71\x2d\x2f\x0a\x79\xa9\x00\x44\x51\x96\xd6\xe4\x82\x14\x7c\x25\x2c\x00\xb0\xf4\x1e\x0f\x0e\x29\xb0\xd9\x31\xdb\x9a\x91\xfc\xb8\xfe\xf3\xcb\xde\xbf\x9f\xe4\x92\x04\xc7\x90\x14\x85\xe0\x2b\x51\xe4\x86\x52\x08\x8c\x3a\x96\x05\x31\x4e\xc6\x72\xfd\xee\x7b\xf7\xec\x87\xde\xf9\x27\x61\x43\x1c\x48\xc8\x71\x25\x49\x90\x84\xbc\x50\xc8\x0b\x22\x90\x44\xb9\xb8\x22\x17\x43\x0f\x89\x12\x37\xae\x62\x60\x14\xaa\xba\x96\x10\xa9\xde\xe9\x17\xbd\x93\x37\xfd\x27\x9f\xdf\xbe\xf8\x3a\x97\x20\x3a\x74\xd2\x0a\x10\x04\x59\x14\xe5\x42\xe8\x24\x7e\x2c\xa5\xc0\xa6\x6d\x76\x10\x1e\xc7\x72\xf3\xfa\xf7\x37\x8f\xfe\xd1\xff\xfe\x55\xff\xf4\x93\x5c\x44\x54\xee\x68\x98\xc4\xf1\x88\x79\xa1\x00\xc4\x82\x5c\x5c\x95\x4b\xab\xc1\x57\x05\x6e\x3c\x77\x22\x0f\x11\x73\x09\xe0\x42\x48\x57\x4f\x6f\xfe\xf3\x38\x06\x09\xa9\xda\x24\x48\x6b\xb2\x24\x66\xe7\xd1\x4a\xe4\x29\xd2\x14\x4c\xbd\xe3\x97\xb7\xcf\xbf\x89\x61\x52\x51\x67\x32\xa4\x19\x28\xb4\x1a\x18\x6d\xdb\xd0\x20\x89\x90\xfa\x67\xdf\xba\xc7\xff\x72\xbf\x38\x76\x5f\x7c\x1c\x7e\xe6\x60\x64\xcb\x31\x95\x01\xa8\x22\x10\xd6\xe4\x92\x24\x17\x57\x42\x50\xfc\xb1\x5b\x8b\x3c\x26\x81\x4b\x3e\xa2\x38\x97\x28\xa2\x44\x2e\xf9\x80\x44\x59\x2c\x06\x5f\x15\xb9\xf1\xf8\x05\x20\x78\xca\x7e\x3b\x97\x00\x2f\x04\x15\x63\x13\x05\x95\xc8\xa6\x22\x90\x04\x59\x5a\x91\x85\x41\xe8\x82\x04\x99\x96\xc8\x01\x98\xdc\x0b\x04\x7b\x07\x5f\x26\xc5\xca\x58\x3b\x10\x28\xd7\x83\x57\x9e\xba\x82\x19\x6b\xdc\x26\xb7\x63\x00\xd4\x6d\x64\x99\x6c\xeb\x9a\x0c\xdb\x96\xe8\xbe\x85\x75\x97\x06\x40\xdd\x1b\x16\x95\xc9\xe3\x44\x6a\x27\x48\x63\x35\xc9\x87\xca\x60\x0c\x67\xf4\x25\xc3\x30\x3d\xba\xb9\x71\xbf\x7b\xe5\x3e\x79\x13\x90\xa7\x4e\x6c\xad\xdd\x46\x36\xfb\xf6\xe4\xe6\xf5\x5f\x7b\x27\x8f\xdd\xa7\x3f\xf4\xae\x4e\x42\x23\xf4\x26\x86\x63\x85\xe4\x5f\xc5\x58\x0e\xde\xfd\xf1\xf0\xc8\xb2\x7f\x3c\x3c\x32\x4c\x82\x72\x83\xa8\xb7\x18\x63\xd1\xb0\xa1\xd1\xdc\xfd\x39\xc8\x44\x85\x39\x98\xe4\xd5\x6b\x8d\x28\x78\x17\xf2\x29\x0c\xe6\xeb\x01\x89\x22\xe1\x3a\x39\xbd\xee\x7e\x13\xd9\x73\x45\x19\x9b\x85\xe5\xb3\x9c\x10\x56\xc7\x31\xee\x91\x02\xe9\x03\xbd\x33\xf5\x2c\x71\xcc\xdc\x03\x61\xc5\xcb\x9d\xac\x51\x18\x6a\xb0\xe7\x19\xd4\x41\x14\x12\xd1\x8c\x03\xda\xf5\x8f\xc8\x0a\xe1\xd1\x82\x76\x1b\x11\x85\x97\x4a\x51\x35\x76\x3e\x45\xb5\xd8\x0f\x50\x54\xab\xd9\x31\x0d\xa4\xc4\x96\x14\x53\x75\x67\xde\x1b\xcf\x61\xf3\xeb\xa5\x40\x93\x30\x7b\x36\x61\x13\x90\xc4\x79\x5b\x49\xc8\xab\x93\x92\x04\xa7\xdb\x58\xaa\xc3\x0c\x95\xdc\xeb\xbf\x32\x55\x71\x4f\x31\x56\x75\xd8\x26\xfa\xb4\x99\x9e\x67\x56\xf7\x5c\xef\x35\x52\x23\x2b\x98\x19\x46\x73\xd6\x65\x1a\xbb\xd7\x63\x43\x34\xa3\xb3\xc7\xda\x45\xea\x6b\xa6\xe1\x39\x18\x99\xdd\x17\x1f\xdf\x3e\x3f\x8d\x4f\x8f\xcb\x74\x68\x5d\x1e\x9d\x11\xa3\xb3\x6b\x41\x94\x0b\xc5\xdc\x84\x9e\x35\x11\x84\x94\x34\xeb\xf9\x4f\xf7\x9b\xe6\xde\xe3\x2f\x63\xf3\xe0\x32\x6d\xf8\xe9\xbf\x22\xfd\x57\x5a\xa6\xbd\xed\x72\x72\x8b\x1f\xce\x1d\x85\x52\x6e\x66\xaa\xd3\x0b\xba\x4c\x5c\xa7\x9a\xac\x2d\xd6\x3c\xef\xe4\xf8\xae\x20\xe7\x95\x63\x67\xbb\x9d\x9b\x53\x96\x66\xa8\x7f\x89\x37\x74\xd9\x1a\xfb\x99\xaf\xf5\xe6\x72\xb1\x37\x8f\xab\xbd\x79\x5c\xee\xcd\x78\xbd\x07\x40\xdd\x34\x14\xec\x34\x9b\x08\xe3\xf4\x1b\x3e\x5f\xbc\x05\xb5\x8e\xc3\x72\x21\xc8\x7c\xe3\x16\xde\xc9\xcf\xaf\xce\x8e\xe6\x85\x3d\x68\x6b\x5e\xaa\xf9\x19\x26\xd9\x91\x35\xf4\xbc\x91\xfb\x2d\x6f\x96\xea\xcd\x7a\x8d\x42\xd7\xf3\x98\xe9\x67\x2c\x9e\x55\xad\xf9\x90\xd1\x2a\xdc\x83\x04\xa6\x66\x97\x2c\xe3\x00\x2d\x90\x8a\x27\xc5\x24\x3f\xb7\xce\x2c\xd6\x23\x78\x91\xc9\xd2\x23\xd0\x88\x4e\xea\x11\x62\xe5\x19\x89\x02\x54\x9b\x85\xb5\xe2\x5a\x03\x96\xd6\x60\xa3\x81\x4a\x77\x90\x50\x5a\x69\x49\x42\x6b\x75\xb5\x80\x06\x6b\xa6\x91\x02\x3e\x7c\xf3\xca\xf6\xaf\x16\x97\x86\x1f\xce\x5c\x99\xe9\x38\xa5\xe3\x76\xa6\xe2\x1c\x2a\x67\xa3\xb3\xce\x5a\x59\x79\x6e\xa0\xdd\xb3\x8b\xeb\xcb\x8b\xeb\x6e\x97\x2f\xab\x0f\xa6\x04\xa8\x12\x9d\x4d\x36\xac\xf7\x89\x97\x49\xc2\xe8\x00\x3a\x6d\xf8\x1c\x08\xcd\x89\xf9\x89\x61\xf2\x06\x20\xd6\x64\x99\xbc\x7e\x89\xcb\x0c\xe6\xf0\x84\x78\xcf\x0a\x35\x3a\xdf\xa4\x3f\xa7\xee\xc9\xd3\xae\x29\x61\x1c\xe9\x68\xba\x46\xc6\x3f\x8f\x63\x0b\x9f\xc1\x0a\x90\x98\x0f\x91\x91\x8d\xf4\xce\xa8\xf4\xa4\xf9\x37\xb6\xe5\x9c\xf8\x43\x07\xaf\x80\x5b\xc8\x48\xff\xe9\x1b\x7f\x19\x8c\xd7\x85\xa9\xa2\x5c\x75\x01\xd2\xde\x24\xf4\xe1\x70\x97\x96\x78\x1c\x5b\x36\xc2\xbb\x4c\xb2\x68\xdf\xd2\x6c\x84\x15\xcd\x98\xec\x5f\x61\x44\x96\xb9\xde\x0c\x91\x30\x6b\x70\xe5\x1d\xfa\x82\x09\xec\x75\x92\x7e\x9c\xc2\x74\xb2\xfe\x1b\x00\x00\xff\xff\x5a\x9e\xaf\xd6\xf2\x30\x00\x00")

func mysql_000001_gokins_up_sql() ([]byte, error) {
	return bindata_read(
		_mysql_000001_gokins_up_sql,
		"mysql/000001_gokins.up.sql",
	)
}

var _sqlite_000001_gokins_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x28\x89\x4f\x2a\xcd\xcc\x49\x49\xb0\x06\x04\x00\x00\xff\xff\xaa\x8a\x92\x63\x1f\x00\x00\x00")

func sqlite_000001_gokins_down_sql() ([]byte, error) {
	return bindata_read(
		_sqlite_000001_gokins_down_sql,
		"sqlite/000001_gokins.down.sql",
	)
}

var _sqlite_000001_gokins_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd2\x4d\x8b\xda\x40\x18\xc0\xf1\x7b\x3e\xc5\x73\xb3\x85\xe6\x13\xf4\x94\xea\x14\xa4\x89\x16\x19\xa1\x9e\x4c\x34\x63\x3b\x10\x63\x98\x99\x78\x56\x28\x54\x6b\x6f\x7d\xb1\x8b\xfb\xe6\x5e\xbc\x2c\xb8\x07\xf1\x60\x16\xfc\x32\x66\xd4\x6f\xb1\x68\x74\xc9\xc2\x1a\x93\x6b\xfe\x3f\xe6\x99\x87\x51\x55\x50\x13\x3e\x45\x55\x01\x5b\x35\x87\x00\x17\xcc\xaf\x0b\x9f\x11\x68\xb4\x18\x88\x6a\xcd\xa7\x8e\xad\x9c\xe3\xd9\x12\xd2\x30\x02\xac\x7d\xd0\x11\x98\x07\x65\xc2\x1b\x05\xc0\xa4\xb6\x09\x18\x7d\xc1\x50\x28\x62\x28\x94\x75\x1d\x3e\x97\xf2\x86\x56\xaa\xc0\x27\x54\x79\xb7\x2b\x3c\xea\x11\x87\xba\xa4\xfa\x9c\xe6\xd0\x47\xad\xac\x47\xf9\xcb\xa4\x4d\x18\xa7\x2d\x37\x21\xe5\xc2\x12\x3e\x7f\xe5\x2f\x64\x8b\x86\x81\x0a\x18\x32\xf2\xfa\x7b\x18\x2c\xd6\x3f\xe7\xb2\xd3\xcd\xec\x11\x61\xac\xc5\x12\xcd\xf6\xcf\xc5\x66\x3a\x5d\x2d\xc7\xb2\x3b\x3d\x98\x36\x71\x45\xa2\x59\x2d\x06\xab\x60\x1e\xd5\x82\x36\x49\x95\x0b\xab\xe9\x99\x90\xd3\x30\xc2\x79\x03\x9d\x1a\xaf\x3f\xd9\x8c\x7f\xc9\xe1\x7c\xfb\x77\x79\xc4\xc2\x21\xc7\x35\xee\xda\x13\xf0\xf6\xc7\xf6\xee\x7f\x44\x9a\x84\x73\xeb\x6b\x1a\xb4\x5f\x46\xfc\x62\x5c\x58\x4c\x10\x3b\x36\x67\x02\x0f\x1f\x3b\xe1\x64\xb0\x1b\x76\x38\x8b\x78\x83\xba\x94\x7f\x4b\xeb\xd7\xc1\x6f\x79\x75\x13\xf7\x75\x46\xac\xf4\xc7\xf7\x46\x61\xb0\x88\x73\xdf\xb3\xd3\x73\x39\x9a\xc9\x7f\x0f\x71\x7e\x78\x60\xe7\xf7\xb6\xee\xf7\xe4\xe5\x7d\x46\x79\xfb\xfe\x29\x00\x00\xff\xff\x91\x8c\x12\xe6\x5d\x03\x00\x00")

func sqlite_000001_gokins_up_sql() ([]byte, error) {
	return bindata_read(
		_sqlite_000001_gokins_up_sql,
		"sqlite/000001_gokins.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"mysql/000001_gokins.down.sql": mysql_000001_gokins_down_sql,
	"mysql/000001_gokins.up.sql": mysql_000001_gokins_up_sql,
	"sqlite/000001_gokins.down.sql": sqlite_000001_gokins_down_sql,
	"sqlite/000001_gokins.up.sql": sqlite_000001_gokins_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"mysql": &_bintree_t{nil, map[string]*_bintree_t{
		"000001_gokins.down.sql": &_bintree_t{mysql_000001_gokins_down_sql, map[string]*_bintree_t{
		}},
		"000001_gokins.up.sql": &_bintree_t{mysql_000001_gokins_up_sql, map[string]*_bintree_t{
		}},
	}},
	"sqlite": &_bintree_t{nil, map[string]*_bintree_t{
		"000001_gokins.down.sql": &_bintree_t{sqlite_000001_gokins_down_sql, map[string]*_bintree_t{
		}},
		"000001_gokins.up.sql": &_bintree_t{sqlite_000001_gokins_up_sql, map[string]*_bintree_t{
		}},
	}},
}}
