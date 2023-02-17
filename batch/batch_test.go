package batch

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type JobImpl struct {
	db     *sql.DB
	count  int
	key    string
	value  int64
	values []int64
}

/*
group.ddl

CREATE TABLE `group` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT '',
  `description` varchar(200) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1011 DEFAULT CHARSET=utf8mb4
*/

func (j *JobImpl) Do() (err error) {
	if len(j.values) == 0 {
		j.values = []int64{j.value}
	}
	sql := "insert into `group` (`name`, `description`) values (?, ?)"
	if valueLen := len(j.values); valueLen > 1 {
		sql += strings.Repeat(",(?, ?)", len(j.values)-1)
	}
	columns := make([]interface{}, 0, len(j.values)*2)
	for _, value := range j.values {
		columns = append(columns, value)
		columns = append(columns, strconv.Itoa(int(j.value)))
	}
	_, err = j.db.Exec(sql, columns...)
	return err
}

func (j *JobImpl) Key() string {
	return j.key
}

func (j *JobImpl) Value() int64 {
	return j.value
}

func (j *JobImpl) UpdateValue(val int64) (ready bool) {
	if len(j.values) == 0 {
		j.values = append(j.values, j.value)
	}
	if val > j.value {
		j.count++
		j.value = val
		j.values = append(j.values, val)
	}
	return j.count >= 99
}

func (j *JobImpl) String() string {
	return fmt.Sprintf("key %s, %d, values %+v aggregate to %d\n", j.key, len(j.values), j.values, j.value)
}

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/sync_es")
	if err != nil {
		panic(err)
	}
	return db
}

func batchf(count int) {
	db := getDB()
	defer db.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	b := New[string, int64](ctx, 1, 80)
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		i := i
		promise, _ := b.AddAndWait(&JobImpl{
			db:    db,
			key:   "1",
			value: int64(i),
		})
		go func() {
			defer wg.Done()
			promise.Error()
			// fmt.Printf("goroutine %d, err %+v\n", i)
		}()
		// time.Sleep(time.Millisecond)
	}
	wg.Wait()
	// time.Sleep(time.Second)
	b.StopAndWait()
}

func single(count int) {
	db := getDB()
	defer db.Close()
	var wg sync.WaitGroup
	sema := newSema(80)
	wg.Add(count)
	for i := 0; i < count; i++ {
		i := i
		go func() {
			sema.acquire()
			defer sema.release()

			defer wg.Done()
			_, err := db.Exec("insert into `group` (name, description) values(?, ?)", i, i)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}

func BenchmarkBatch(b *testing.B) {
	single(10000)
}

func Test_btach(t *testing.T) {
	batchf(10000)
}
