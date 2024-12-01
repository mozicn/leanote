package tests

import (
	"github.com/mozicn/leanote/app/db"
	"testing"
	//	. "github.com/mozicn/leanote/app/lea"
	//	"github.com/mozicn/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}
