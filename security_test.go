package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_Security(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "six910"
	var dbi db.Database = &mydb

	var sdb Six910Mysql
	var l lg.Logger
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	si := sdb.GetNew()
	var sec sdbi.Security
	sec.OauthOn = true
	suc, id := si.AddSecurity(&sec)
	fmt.Println("add success: ", suc)
	fmt.Println("add success id: ", id)
	if !suc || id == 0 {
		t.Fail()
	}

	secr := si.GetSecurity()
	fmt.Println("get : ", secr)
	if !secr.OauthOn {
		t.Fail()
	}
	sec.ID = id
	sec.OauthOn = false

	sucu := si.UpdateSecurity(&sec)
	if !sucu {
		t.Fail()
	}

	secr2 := si.GetSecurity()
	fmt.Println("get : ", secr2)
	if secr2.OauthOn {
		t.Fail()
	}

	dsuc := si.DeleteSecurity()
	if !dsuc {
		t.Fail()
	}
	dbi.Close()

}
