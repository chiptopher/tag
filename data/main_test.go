
// MIT License
//
// Copyright (c) 2018 Christopher M. Boyer
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package data

import (
	"testing"
	"os"
	"path/filepath"
	"github.com/Chiptopher/tag/test"
)

func TestMain(m *testing.M) {

	os.RemoveAll(filepath.Join(test.GoPathTestFolder, DataFolderName))
	os.Remove(filepath.Join(test.GoPathTestFolder, DatabaseName))

	errorCode := m.Run()

	os.RemoveAll(filepath.Join(test.GoPathTestFolder, DataFolderName))
	os.Remove(filepath.Join(test.GoPathTestFolder, DatabaseName))

	os.RemoveAll(test.CreatePathInTestFolder(DataFolderName))
	os.Remove(test.CreatePathInTestFolder("sub"))

	os.Exit(errorCode)

}
