package webapp

import (
	"iflytek.com/mongotxt"
	"testing"
)

func TestGetBook(t *testing.T) {
	result := mongotxt.Book{}
	result = GetBook("/Users/yfyang/Documents/武侠&小说/碧血剑.txt")
	if result.Name != "/Users/yfyang/Documents/武侠&小说/碧血剑.txt" {
		t.Error("查询结果不正确")
	}
}
