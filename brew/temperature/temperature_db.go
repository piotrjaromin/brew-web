package temperature

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/tidwall/buntdb"
)

type TemperatureDB struct {
	db *buntdb.DB
}

func CreateTempDb() (Temperatures, error) {
	db, err := buntdb.Open("temp-data.db")

	return &TemperatureDB{
		db: db,
	}, err
}

func (t TemperatureDB) Get() ([]TemperaturePoint, error) {
	temps := make([]TemperaturePoint, 0)

	err := t.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			t := TemperaturePoint{}
			err := json.Unmarshal([]byte(value), &t)
			if err != nil {
				log.Printf("unable to parse temp point from db %s", err.Error())
				return false
			}

			temps = append(temps, t)
			return true
		})
		return err
	})

	return temps, err
}

func (t *TemperatureDB) Add(s float64) error {
	point := TemperaturePoint{s, time.Now()}

	strPoint, err := json.Marshal(point)
	if err != nil {
		return fmt.Errorf("Could not parse temp point to JSON: %s", err)
	}

	return t.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(point.TimeStamp.String(), string(strPoint), &buntdb.SetOptions{
			Expires: true,
			TTL:     time.Hour * 24,
		})
		return err
	})

}
