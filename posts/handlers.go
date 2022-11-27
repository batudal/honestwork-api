package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetPost(redis *redis.Client, address string) schema.User {
	var user schema.User
	// update with hash
	data, err := redis.Do(redis.Context(), "JSON.GET", address).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return user
}

func getUserFromAddress(address string) *schema.User {
	var user schema.User
	resp, err := http.Get("http://localhost:3002/users/" + address)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	userBody, err := io.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	err = json.Unmarshal([]byte(userBody), &user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return &user
}

func HandleNewPost(redis *redis.Client, address string, signature string, body []byte) string {
	result := crypto.VerifySignature("post", address, signature)
	if !result {
		return "Wrong signature."
	}

	user := getUserFromAddress(address)
	http.Post("http://localhost:3002/users/increment_post/"+address+"/"+signature,
		"application/json", bytes.NewBuffer([]byte("")))
	salt := "post" + strconv.Itoa(user.Posts+1)
	hash := crypto.GenerateID(salt, address)
	redis.Do(redis.Context(), "JSON.SET", hash, "$", body)

	return "success"
}
