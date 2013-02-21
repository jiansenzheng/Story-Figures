package mongotxt

import (
	"bufio"
	// mongoDB驱动
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	// "flag"
	"log"
	"os"
	"path/filepath"
	// "strconv"
	"strings"
	"time"
)

const (
	IsDirectory = iota
	IsRegular
	IsSymlink
)

type sysFile struct {
	fType      int
	fName      string
	fLink      string
	fSize      int64
	fMtime     time.Time
	fPerm      os.FileMode
	fShortName string
}

type F struct {
	files []*sysFile
}

type Story struct {
	Book    string //书名称
	Line    int    //行数，第几行的内容
	Content string //行的内容
	words   int    //行的字数
}

type Book struct {
	Name string    //书名称
	Size int64     //大小
	Time time.Time //创建时间
	Line int       //总行数
}

func (self *F) visit(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	//如果是txt文本
	if strings.HasSuffix(f.Name(), "txt") {
		var tp int
		if f.IsDir() {
			tp = IsDirectory
		} else if (f.Mode() & os.ModeSymlink) > 0 {
			tp = IsSymlink
		} else {
			tp = IsRegular
		}
		inoFile := &sysFile{
			fName:      path,
			fType:      tp,
			fPerm:      f.Mode(),
			fMtime:     f.ModTime(),
			fSize:      f.Size(),
			fShortName: f.Name(),
		}
		self.files = append(self.files, inoFile)
	}
	return nil
}

func ScanImportMongo(root string, mongoHost string) {

	self := F{
		files: make([]*sysFile, 0),
	}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		return self.visit(path, f, err)
	})

	if err != nil {
		log.Fatalln("filepath.Walk() returned %v\n", err)
	}

	session, err := mgo.Dial(mongoHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	story := session.DB("gomboss").C("story")
	book := session.DB("gomboss").C("book")

	for _, v := range self.files {
		log.Println("正在从文本" + v.fName + "中导入数据...")
		TxtImport(story, book, v)
		log.Println("从文本" + v.fName + "中导入数据完成")
	}

}

// 开始导入文本小说内容
func TxtImport(story *mgo.Collection, book *mgo.Collection, sysFile *sysFile) {

	f, _ := os.OpenFile(sysFile.fName, os.O_RDONLY, 0666)
	defer f.Close()
	m := bufio.NewReader(f)
	// char := 0
	words := 0
	lines := 0
	for {
		s, ok := m.ReadString('\n')
		words = len(strings.Fields(s))
		if words > 0 {
			story.Insert(&Story{sysFile.fShortName, lines, s, words})
		}

		lines++
		if ok != nil {
			break
		}
	}
	book.Insert(&Book{f.Name(), sysFile.fSize, sysFile.fMtime, lines})
}
