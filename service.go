package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// GetKey handler shows the value at "key/id" as JSON.
func GetKey(key string) string {
	id, err := strconv.Atoi(strings.TrimSpace(string(key)))
	HandleError(err)
	value := FindKey(id)
	if value == "" {
		return "No value for this key"
	}
	return value
}

// KeyCreate creates a new post data
func KeyCreate(c net.Conn, key string, value string) string {
	// Save JSON to Post struct (should this be a pointer?)
	var kv KV
	id, err := strconv.Atoi(strings.TrimSpace(string(key)))
	HandleError(err)
	kv.Key = id
	kv.Value = value

	CreateKey(kv)
	return key + " " + value
}

func FindKey(id int) string {
	var keyval KV

	c := RedisConnect()
	defer c.Close()

	reply, err := c.Do("GET", "kv:"+strconv.Itoa(id))
	HandleError(err)

	if reply == nil {
		return "No value for this key"
	}

	if err = json.Unmarshal(reply.([]byte), &keyval); err != nil {
		panic(err)
	}
	return keyval.Value
}

// CreateKey creates a blog post.
func CreateKey(k KV) {

	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(k)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "kv:"+strconv.Itoa(k.Key), b)
	HandleError(err)

	fmt.Println("GET ", reply)
}
