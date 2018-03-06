
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

package logic

import (
	"testing"
	"github.com/Chiptopher/tag/data"
)


func TestTagIdentityMap_GetTagById(t *testing.T) {

	filePath := "path"

	tagData := &data.TagDataDTO{1, "name"}

	tagGateway := data.MockNewTagGateway()
	tagGateway.On("ReadTag", 1, filePath).Return(tagData)

	tagMapper := NewTagMapper(filePath)
	tagMapper.tagGateway = tagGateway

	tagMap := NewTagIdentityMap(tagMapper)

	tag, err := tagMap.GetTagById(1)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	tag2, err := tagMap.GetTagById(1)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if tag != tag2 {
		t.Fatalf("Doesn't appropriately save in map.")
	}


}

func TestTagIdentityMap_GetTagByName(t *testing.T) {
	
	filePath := "path"
	tagName := "name"

	tagData := &data.TagDataDTO{1, tagName}

	tagGateway := data.MockNewTagGateway()
	tagGateway.On("ReadTag", 1, filePath).Return(tagData)
	tagGateway.On("GetTagIdByName", "name", filePath).Return(1, nil)

	tagMapper := NewTagMapper(filePath)
	tagMapper.tagGateway = tagGateway

	tagMap := NewTagIdentityMap(tagMapper)

	tag, err := tagMap.GetTagByName(tagName)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	tag2, err := tagMap.GetTagByName(tagName)

	if err != nil {
		t.Fatalf("Error encountered: %s", err)
	}

	if tag != tag2 {
		t.Fatalf("Doesn't appropriately save in map.")
	}


}