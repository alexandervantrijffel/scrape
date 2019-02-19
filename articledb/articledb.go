package articledb

import (
	"fmt"

	"github.com/alexandervantrijffel/scrape/config"
	sdb "github.com/streamsdb/driver/go/sdb"
)

var THECONNECTION sdb.Connection
var THEDB sdb.DB

func Connect() sdb.DB {
	THECONNECTION = sdb.MustOpen(fmt.Sprintf("sdb://%s@sdb03.streamsdb.io:443?tls=1", config.THECONFIG.STREAMSDBCREDENTIALS))
	THEDB = THECONNECTION.DB("news")
	return THEDB
}

func Close() {
	THECONNECTION.Close()
}
