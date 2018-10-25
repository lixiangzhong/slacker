package models

import (
	"dns.com/ini"
	"dns.com/log"
	"dns.com/utils"
	valid "github.com/asaskevich/govalidator"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net"
	"strings"
	"time"
)

const (
	StateOK  = 0
	StateDel = 1
)

var (
	db *sqlx.DB
)

func init() {
	initdb()
}

func initdb() {
	var err error
	var cfg = mysql.NewConfig()
	cfg.Addr = ini.String("mysql", "host")
	cfg.User = ini.String("mysql", "user")
	cfg.Passwd = ini.String("mysql", "password")
	cfg.DBName = ini.String("mysql", "db")
	cfg.ParseTime = true
	cfg.Loc = time.Local
	cfg.Net = "tcp"
	db, err = sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		log.Error(err)
	}
	if ini.Bool("mysql", "unsafe") {
		db = db.Unsafe()
	}
	go func() {
		tk := time.NewTicker(time.Second * 15)
		for range tk.C {
			db.Ping()
		}
	}()
	initValidator()
}

func ValidateStruct(s interface{}) error {
	ok, err := valid.ValidateStruct(s)
	if !ok || err != nil {
		errs, ok := (err).(valid.Errors)
		if ok {
			for _, e := range errs {
				return e
			}
		}
	}
	return err
}

func initValidator() {
	valid.TagMap["cidr"] = valid.Validator(func(s string) bool {
		ip, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			return false
		}
		return ipnet.IP.Equal(ip)
	})

	valid.TagMap["domain"] = valid.Validator(func(s string) bool {
		var err error
		s, err = utils.Domain.PunyCode(s)
		if err != nil {
			return false
		}
		s = strings.Trim(s, ".")
		if len(strings.Split(s, ".")) < 2 {
			return false
		}
		return utils.Domain.Valid(s)
	})
}
