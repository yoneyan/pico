package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var config Config
var ledState = false
var patternName = "default"
var job Job

func main() {
	err := GetConfig()
	if err != nil {
		panic(err)
	}
	go func() {
		// LEDの状態を監視する処理
		for {
			ExecuteJob()
			// 1秒ごとに状態を確認
			time.Sleep(1 * time.Second)
			checkOverwritePattern()
		}
	}()

	// 静的ファイルの提供
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// 各ボタンに対応する処理
	http.HandleFunc("/action1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Action 1 executed!")
		w.Write([]byte("Action 1 executed"))
		//pattern = "action1"

	})

	http.HandleFunc("/action2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Action 2 executed!")
		w.Write([]byte("Action 2 executed"))
	})

	http.HandleFunc("/action3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Action 3 executed!")
		w.Write([]byte("Action 3 executed"))
	})

	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
