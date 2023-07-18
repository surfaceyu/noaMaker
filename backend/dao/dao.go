package dao

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
)

type boltDb struct {
	db *bolt.DB
}

var db *boltDb

func init() {
	initBlotDb()
	initBookSites()
}

func initBlotDb() {
	d, err := bolt.Open("cache.db", 0600, nil)
	if err != nil {
		log.Fatal("Db file open filed!!! Please Check.")
	}
	db = &boltDb{
		db: d,
	}
}

func (bo boltDb) View(key string, viewData any) {
	err := bo.db.View(func(tx *bolt.Tx) error {
		// 获取名为"bookSites"的Bucket
		bucket := tx.Bucket([]byte(key))
		if bucket == nil {
			// 如果Bucket不存在，则直接返回
			return nil
		}
		// 获取Bucket中名为"data"的键对应的值
		data := bucket.Get([]byte("data"))
		if data == nil {
			// 如果值为空，则直接返回
			return nil
		}
		// 将值解码为bookSites切片
		err := json.Unmarshal(data, viewData)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (bo boltDb) Update(key string, viewData any) error {
	err := bo.db.Update(func(tx *bolt.Tx) error {
		// 创建或获取名为"bookSites"的Bucket
		bucket, err := tx.CreateBucketIfNotExists([]byte(key))
		if err != nil {
			return err
		}

		// 将bookSites切片转换为JSON格式
		data, err := json.Marshal(viewData)
		if err != nil {
			return err
		}

		// 将JSON数据保存到Bucket中
		err = bucket.Put([]byte("data"), data)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
