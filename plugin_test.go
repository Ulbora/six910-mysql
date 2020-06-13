package six910mysql

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Mysql_AddPlugin(t *testing.T) {
	var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "six910"
	var dbi db.Database = &mydb

	var sdb Six910Mysql
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	si := sdb.GetNew()

	var pi sdbi.Plugins

	pi.ActivateURL = "test"
	pi.Category = "catalog"
	pi.ContactPhone = "123-789-4555"
	pi.Developer = "Ulbora Labs LLC"
	pi.DeveloperAddress = "atlanta, GA"
	pi.Enabled = true
	pi.Fee = 0
	pi.IsPGW = false
	pi.OauthRedirectURL = "/redirect"
	pi.PluginName = "Catalog Easy"

	dbi.Close()
	pisuc, piid := si.AddPlugin(&pi)
	if !pisuc || piid == 0 {
		t.Fail()
	}

	pi.ID = piid
	pi.ActivateURL = "test2"
	pi.Category = "catalog2"
	pi.ContactPhone = "123-789-4552"
	pi.Developer = "Ulbora Labs LLC2"
	pi.DeveloperAddress = "atlanta, GA2"
	pi.Enabled = false
	pi.Fee = 100.00
	pi.IsPGW = true
	pi.OauthRedirectURL = "/redirec2t"
	pi.PluginName = "Catalog Easy2"

	dbi.Close()
	upisuc := si.UpdatePlugin(&pi)
	if !upisuc {
		t.Fail()
	}

	dbi.Close()
	fpi := si.GetPlugin(piid)
	fmt.Println("fpi: ", fpi)
	if fpi.DeveloperAddress != pi.DeveloperAddress {
		t.Fail()
	}

	dbi.Close()
	fpilist := si.GetPluginList(0, 10)
	fmt.Println("fpilist:", fpilist)
	fmt.Println("fpilist len:", len(*fpilist))
	if len(*fpilist) == 0 {
		t.Fail()
	}

	dbi.Close()
	dlpi := si.DeletePlugin(piid)
	if !dlpi {
		t.Fail()
	}

}
