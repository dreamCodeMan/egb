package egb

import (
	"time"
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"os"
	"path"
)

//FileBufferedReader return io.Reader or error by  the given file url.
func FileBufferedReader(filenameOrURL string) (io.Reader, error) {
	data, err := FileGetBytes(filenameOrURL)
	if err != nil {
		return nil, err
	}
	return BytesReader(data), nil
}

//FileGetBytes return []byte or error by the given url and timeout.
func FileGetBytes(filenameOrURL string, timeout ...time.Duration) ([]byte, error) {
	if strings.Contains(filenameOrURL, "://") {
		if strings.Index(filenameOrURL, "file://") == 0 {
			filenameOrURL = filenameOrURL[len("file://"):]
		} else {
			client := http.DefaultClient
			if len(timeout) > 0 {
				client = &http.Client{Timeout: timeout[0]}
			}
			r, err := client.Get(filenameOrURL)
			if err != nil {
				return nil, err
			}
			defer r.Body.Close()
			if r.StatusCode < 200 || r.StatusCode > 299 {
				return nil, fmt.Errorf("%d: %s", r.StatusCode, http.StatusText(r.StatusCode))
			}
			return ioutil.ReadAll(r.Body)
		}
	}
	return ioutil.ReadFile(filenameOrURL)
}

//FileSetBytes set bytes to the given file.
func FileSetBytes(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0660)
}

//FileAppendBytes append bytes to the given file.
func FileAppendBytes(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

//FileGetString get file string content by given filename or url and timeout.
func FileGetString(filenameOrURL string, timeout ...time.Duration) (string, error) {
	bytes, err := FileGetBytes(filenameOrURL, timeout...)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//FileSetString set string to the given file.
func FileSetString(filename string, data string) error {
	return FileSetBytes(filename, []byte(data))
}

//FileAppendString append string to the given file.
func FileAppendString(filename string, data string) error {
	return FileAppendBytes(filename, []byte(data))
}

// FileTimeModified returns the modified time of a file,
// or the zero time value in case of an error.
func FileTimeModified(filename string) time.Time {
	info, err := os.Stat(filename)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}

//FileExists return true if the given file exist.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

//FileIsDir return true if the given filename is a dir.
func FileIsDir(dirname string) bool {
	info, err := os.Stat(dirname)
	return err == nil && info.IsDir()
}

//FileFind find given file int the given dirs and return the result.
func FileFind(searchDirs []string, filenames ...string) (filePath string, found bool) {
	for _, dir := range searchDirs {
		for _, filename := range filenames {
			filePath = path.Join(dir, filename)
			if FileExists(filePath) {
				return filePath, true
			}
		}
	}
	return "", false
}

// FileSize returns the size of a file or zero in case of an error.
func FileSize(filename string) int64 {
	info, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return info.Size()
}

//FileGetPrefix return the prefix of filename.
func FileGetPrefix(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[0:i]
		}
	}
	return ""
}

//FileGetExt return the ext of filename.
func FileGetExt(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[i + 1:]
		}
	}
	return ""
}

// FileCopy copies file source to destination dest.
func FileCopy(source string, dest string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		si, err := os.Stat(source)
		if err == nil {
			err = os.Chmod(dest, si.Mode())
		}
	}
	return err
}

// FileCopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func FileCopyDir(source string, dest string) (err error) {
	// get properties of source dir
	fileInfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return &FileCopyError{"Source is not a directory"}
	}
	// ensure dest dir does not already exist
	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		return &FileCopyError{"Destination already exists"}
	}
	// create dest dir
	err = os.MkdirAll(dest, fileInfo.Mode())
	if err != nil {
		return err
	}
	entries, err := ioutil.ReadDir(source)
	for _, entry := range entries {
		sourcePath := path.Join(source, entry.Name())
		destinationPath := path.Join(dest, entry.Name())
		if entry.IsDir() {
			err = FileCopyDir(sourcePath, destinationPath)
		} else {
			// perform copy
			err = FileCopy(sourcePath, destinationPath)
		}
		if err != nil {
			return err
		}
	}
	return err
}

//ListDir return all the names from the directory in a single slice.
func ListDir(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Readdirnames(-1)
}

//ListDirFiles return all the file names from the directory.
func ListDirFiles(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileInfos, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(fileInfos))
	for i := range fileInfos {
		if !fileInfos[i].IsDir() {
			result = append(result, fileInfos[i].Name())
		}
	}
	return result, nil
}

////ListDirDirectories return all the directory names from the directory.
func ListDirDirectories(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileInfos, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(fileInfos))
	for i := range fileInfos {
		if fileInfos[i].IsDir() {
			result = append(result, fileInfos[i].Name())
		}
	}
	return result, nil
}

// A struct for returning file copy error messages
type FileCopyError struct {
	What string
}

func (e *FileCopyError) Error() string {
	return e.What
}



