package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `app.js`,
		FileModTime: time.Unix(1572295140, 0),
		Content:     string([]byte{0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x29, 0x2e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x54, 0x4d, 0x4c, 0x20, 0x3d, 0x20, 0x22, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x20, 0x3c, 0x62, 0x72, 0x20, 0x2f, 0x3e, 0x3c, 0x70, 0x3e, 0x54, 0x6f, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x20, 0x62, 0x79, 0x20, 0x4e, 0x69, 0x63, 0x20, 0x52, 0x61, 0x62, 0x6f, 0x79, 0x3c, 0x2f, 0x70, 0x3e, 0x22, 0x3b}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `index.html`,
		FileModTime: time.Unix(1572294137, 0),
		Content:     string([]byte{0x3c, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x20, 0x54, 0x65, 0x73, 0x74, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x2f, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x70, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3e, 0x3c, 0x2f, 0x70, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x22, 0x61, 0x70, 0x70, 0x2e, 0x6a, 0x73, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x2f, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x3e}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `styles.css`,
		FileModTime: time.Unix(1572295076, 0),
		Content:     string([]byte{0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x34, 0x30, 0x30, 0x70, 0x78, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x70, 0x78, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x63, 0x79, 0x61, 0x6e, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x35, 0x30, 0x25, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x59, 0x28, 0x2d, 0x35, 0x30, 0x25, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x64, 0x61, 0x72, 0x6b, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x67, 0x72, 0x61, 0x79, 0x3b, 0xa, 0x7d}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1572321069, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // app.js
			file3, // index.html
			file4, // styles.css

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`website`, &embedded.EmbeddedBox{
		Name: `website`,
		Time: time.Unix(1572321069, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"index.html": file3,
			"styles.css": file4,
		},
	})
}
