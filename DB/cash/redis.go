package cash

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

func VisitHappened() (err error) {
	conn := Pool.Get()
	defer conn.Close()
	var key = "visit"

	_, err = redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	var data []byte
	data, err = redis.Bytes(conn.Do("GET", key))
	fmt.Println(err)

	if err != nil && data!=nil{
		return
	}
	number, _ := strconv.Atoi(string(data))
	number=number+1
	_, err = conn.Do("SET", key, number)
	if err != nil {
		v := string(number)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}

	return
}
func VisitGetNumber() (number string,err error) {
	conn := Pool.Get()
	defer conn.Close()
	var key = "visit"

	_, err = redis.String(conn.Do("PING"))

	if err != nil {
		return "0",fmt.Errorf("cannot 'PING' db: %v", err)
	}
	var data []byte
	data, err = redis.Bytes(conn.Do("GET", key))
	if err != nil && data!=nil{
		return
	}
	return string(data),nil
}