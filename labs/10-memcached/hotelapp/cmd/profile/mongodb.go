//go:build mongodb

package main

import (
	"time",
	"flag"
	"path/filepath"

	"github.com/ucy-coast/hotel-app/internal/profile"
)

var (
	memcache_addr = flag.String("memc_addr", "memcached-profile:11211", "Address of the memcache server")
	database_addr = flag.String("db_addr", "mongodb-profile:27017", "Address of the data base server")
)

func initializeProfileDatabase() *profile.DatabaseSession {
	dbsession := profile.NewDatabaseSession(*database_addr, *memcache_addr)	dbsession.LoadDataFromJsonFile(filepath.Join(*jsonDataDir, "hotels.json"))
	return dbsession
}
