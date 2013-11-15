/**
 * User: jackong
 * Date: 11/14/13
 * Time: 7:21 PM
 */
package service

import (
	"time"
	. "morning-dairy/global"
	"morning-dairy/model"
	"github.com/Jackong/db"
)

type dairy struct {

}

var (
	Dairy dairy
)

const (
	FROM_GLOBAL = "0"
	FROM_LEFT = "1"
	FROM_RIGHT = "2"
	FROM_UP = "3"
	FROM_DOWN = "4"
)

func init() {
	Dairy = dairy{}
}

//Dairies get dairies by date and operation
//THe date defines the date user want to view
//THe from defines the direction user switch from
//All dairies surrounding with date will be return when from is FROM_GLOBAL
//Else the dairies surrounding with date (not contain that date and the date switch from)
func (this dairy) Dairies(dateStr string, from string) (dairies map[string] db.Item) {
	dairies = make(map[string] db.Item)
	date, err := time.Parse(FORMAT_DATE, dateStr)
	if err != nil {
		return
	}
	for _, val := range []string{FROM_GLOBAL, FROM_LEFT, FROM_RIGHT, FROM_UP, FROM_DOWN} {
		if from != FROM_GLOBAL && (val == from || val == FROM_GLOBAL) {
			continue
		}
		toDate, dairy := this.GetDairy(date, val)
		dairies[toDate] = dairy
	}
	return
}

func (this dairy) GetDairy(date time.Time, from string) (dateStr string, dairy db.Item) {
	date = this.correspondDate(date, from)
	return date.Format(FORMAT_DATE), model.Dairy.FindByDate(date)
}


func (this dairy) correspondDate(date time.Time, from string) (time.Time) {
	switch from {
	case FROM_LEFT:
		date = date.AddDate(0, 0, -1)
	case FROM_RIGHT:
		date = date.AddDate(0, 0, 1)
	case FROM_UP:
		date = date.AddDate(0, -1, 0)
	case FROM_DOWN:
		date = date.AddDate(0, 1, 0)
	default:
	}
	return date
}
