/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// Watcher is an interface we use to watch changes in files
type Watcher interface {
	Close() error
}

// OSFileWatcher defines a watch over a file
type OSFileWatcher struct {
	file    string
	watcher *fsnotify.Watcher
	// onEvent callback to be invoked after the file being watched changes
	onEvent func()
}

// NewFileWatcher creates a new FileWatcher
func NewFileWatcher(file string, onEvent func()) (Watcher, error) {
	fw := OSFileWatcher{
		file:    file,
		onEvent: onEvent,
	}
	err := fw.watch()
	return fw, err
}

// Close ends the watch
func (f OSFileWatcher) Close() error {
	return f.watcher.Close()
}

// watch creates a fsnotify watcher for a file and create of write events
func (f *OSFileWatcher) watch() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	f.watcher = watcher

	realFile, err := filepath.EvalSymlinks(f.file)
	if err != nil {
		return err
	}

	dir, file := path.Split(f.file)
	go func(file string) {
		for {
			select {
			case event := <-watcher.Events:
				if event.Has(fsnotify.Create) ||
					event.Has(fsnotify.Write) {
					if finfo, err := os.Lstat(event.Name); err != nil {
						log.Printf("can not lstat file: %v\n", err)
					} else if finfo.Mode()&os.ModeSymlink != 0 {
						if currentRealFile, err := filepath.EvalSymlinks(f.file); err == nil &&
							currentRealFile != realFile {
							f.onEvent()
							realFile = currentRealFile
						}
						continue
					}
					if strings.HasSuffix(event.Name, file) {
						f.onEvent()
					}
				}
			case err := <-watcher.Errors:
				if err != nil {
					log.Printf("error watching file: %v\n", err)
				}
			}
		}
	}(file)
	return watcher.Add(dir)
}
