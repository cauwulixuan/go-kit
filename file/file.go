/*
 * Copyright 2022 The Inspur AIStation Group Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Note: the example only works with the code within the same release/branch.

package file

import (
	"bufio"
	"fmt"
	"github.com/cauwulixuan/go-kit/log"
	"io"
	"io/ioutil"
	"os"
)

// CheckFileExist Check file exists or not, if file exist, return true, else return false.
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		log.Slogger.Warnf("File %s does not exist.\n", fileName)
		return false
	}
	return true
}

// CreateFile Create a file with a given fileName.
func CreateFile(fileName string) {
	if CheckFileExist(fileName) {
		log.Slogger.Warnf("File %s already exist.\n", fileName)
		return
	}
	fp, err := os.Create(fileName)
	if err != nil {
		log.Slogger.Errorf("Create file %s failed, error msg: %v.\n", fileName, err.Error())
		return
	}
	defer fp.Close()
}

// ReadFile Read all file content and return.
func ReadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Slogger.Errorf("Read file %s failed, error msg: %v.\n", fileName, err.Error())
		return ""
	}

	return string(content)
}

// ReadFileByChunks Read file with given size.
func ReadFileByChunks(fileName string, size int) string {
	fp, err := os.Open(fileName)
	if err != nil {
		log.Slogger.Errorf("Open file %s failed, error msg: %v.\n", fileName, err.Error())
		return ""
	}
	defer fp.Close()

	rp := bufio.NewReader(fp)
	var chunks []byte

	buf := make([]byte, size)
	for {
		num, err := rp.Read(buf)
		if err != nil && err != io.EOF {
			log.Slogger.Errorf("Error happened while reading chunk from file %s, error msg: %v.\n", fileName, err.Error())
			return ""
		}

		if 0 == num {
			break
		}

		chunks = append(chunks, buf...)
	}

	return string(chunks)
}

// ReadFileByLine Read file line by line.
func ReadFileByLine(fileName string, delim byte) string {
	fp, err := os.Open(fileName)
	if err != nil {
		log.Slogger.Errorf("Open file %s failed, error msg: %v.\n", fileName, err.Error())
		return ""
	}
	defer fp.Close()

	rp := bufio.NewScanner(fp)
	var lines []byte

	for rp.Scan() {
		line := rp.Text()
		// do somethine with each line.

		lines = append(lines, []byte(line)...)
	}
	if err := rp.Err(); err != nil {
		log.Slogger.Error(err)
	}

	return string(lines)
}

// GetFileInfo Get file infos like Name, Size, Mode...
func GetFileInfo(fileName string) (os.FileInfo, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Slogger.Errorf("Get file state failed, error msg: %s\n", err.Error())
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
	return os.Stat(fileName)
}

// RemoveFile Remove file with given fileName.
func RemoveFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Slogger.Errorf("Remove file failed, error msg: %s\n", err.Error())
	}
}

// WriteFile Write file with given fileName and data, return error if write failed.
func WriteFile(fileName string, data []byte) error {
	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Slogger.Errorf("Write file %s failed, error msg: %v.\n", fileName, err.Error())
		return err
	}

	return nil
}

// WriteFileWithAppend Write data to file with a given fileName
// If file exists, append data
// If not, create a new file and write data to the file.
func WriteFileWithAppend(fileName string, data string) error {
	var fp *os.File
	var err error
	if CheckFileExist(fileName) {
		log.Slogger.Infof("File %s exist.\n", fileName)
		fp, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		fp, err = os.Create(fileName)
	}
	defer fp.Close()
	if err != nil {
		log.Slogger.Errorf("Open file or create file %s failed. Error msg: %v\n", fileName, err.Error())
		return err
	}

	w := bufio.NewWriter(fp)
	n, _ := w.WriteString(data)
	log.Slogger.Debugf("Written %v bytes.\n", n)
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}
