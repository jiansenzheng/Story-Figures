package mongotxt

import (
	"log"
	"testing"
)

func TestScanImportMongo(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	ScanImportMongo("/Users/yfyang/Documents/小说/", "127.0.0.1:5565")
}
