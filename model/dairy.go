/**
 * User: jackong
 * Date: 11/15/13
 * Time: 5:50 PM
 */
package model

import (
	"github.com/Jackong/db"
	_ "github.com/Jackong/db/mysql"
	. "morning-dairy/global"
	"time"
)

type dairy struct {
	db.Collection
}

var (
	Dairy dairy
)

func init() {
    Dairy = dairy{Collection("dairy")}
}

func (this dairy) FindByDate(date time.Time) db.Item {
	item, err := this.Collection.Find(db.Cond{
		"date": date,
	})
	if err != nil {
		Log.Warn(err)
	}
	return item
}
