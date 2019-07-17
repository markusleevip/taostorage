// 天地不仁，以万物为刍狗；
// 圣人不仁，以百姓为刍狗。

package taoweb

import (
	"flag"
	"github.com/markusleevip/taostorage/config"
	"github.com/markusleevip/taostorage/db"
	"github.com/markusleevip/taostorage/log"
	"net/http"
)

var flags struct {
	addr, albumPath string
	dbAddr          string
	logto           string
	loglevel        string
}

func init() {
	flag.StringVar(&flags.addr, "addr", ":8000", "The TCP address to bind to")
	flag.StringVar(&flags.dbAddr, "dbAddr", "127.0.0.1:7398", "The TCP address to connect to taodb")
	flag.StringVar(&flags.albumPath, "albumPath", "/data/album", "album save path")
	flag.StringVar(&flags.logto, "log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	flag.StringVar(&flags.loglevel, "log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
}

func Main() {
	flag.Parse()
	proFile := config.Profile{}
	proFile.AlbumPath = flags.albumPath
	config.PFile = proFile
	log.LogTo(flags.logto, flags.loglevel)

	router := NewRouter()

	log.Info("start taodb ...")
	log.Info("AlbumPath=%s",config.PFile.AlbumPath)
	_, err := db.New(flags.dbAddr)

	if err != nil {
		log.Error("start taodb fail:", err)
		panic(err)
	}
	log.Info("start taodb success")

	if err := http.ListenAndServe(flags.addr, router); err != nil {
		log.Error("start fail:", err)
		panic(err)
	}

}
