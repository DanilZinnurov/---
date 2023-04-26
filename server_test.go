package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

type testMessage struct {
	Severity string
	Text     string
	Meta     map[string]string
}

func ReadLine(file io.Reader, lineNum int) (line string) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		if lineCount == lineNum {
			return string(scanner.Bytes())
		}
	}
	return ""
}

func Test_Server(t *testing.T) {
	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	data := testMessage{
		Severity: "debug",
		Text:     "test text",
		Meta:     map[string]string{"user_id": "12", "session_id": "156", "key": "value"},
	}
	payload, _ := json.Marshal(data)
	c, resp, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, payload)
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(resp.Status)
		fmt.Println(err)
	}
	file, _ = os.Open("logs.txt")
	fileStr := ReadLine(file, lineCount+1)
	require.Equal(t, "debug 12 156: test text map[key:value]", fileStr[20:])
}

func Test_Server_2(t *testing.T) {
	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	data := testMessage{
		Severity: "info",
		Text:     "Test_2",
		Meta:     map[string]string{"user_id": "4125", "session_id": "445", "1": "123"},
	}
	payload, _ := json.Marshal(data)
	c, resp, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, payload)
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(resp.Status)
		fmt.Println(err)
	}
	file, _ = os.Open("logs.txt")
	fileStr := ReadLine(file, lineCount+1)
	require.Equal(t, "info 4125 445: Test_2 map[1:123]", fileStr[20:])
}

func Test_Server_3(t *testing.T) {
	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	data := testMessage{
		Severity: "trace",
		Text:     "Test 3",
		Meta:     map[string]string{"user_id": "34", "session_id": "42", "42": "5325"},
	}
	payload, _ := json.Marshal(data)
	c, resp, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, payload)
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(resp.Status)
		fmt.Println(err)
	}
	file, _ = os.Open("logs.txt")
	fileStr := ReadLine(file, lineCount+1)
	require.Equal(t, "trace 34 42: Test 3 map[42:5325]", fileStr[20:])
}

func Test_Server_4(t *testing.T) {
	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	data := testMessage{
		Severity: "warning",
		Text:     "Test 4",
		Meta:     map[string]string{"user_id": "5", "session_id": "42", "6": "78"},
	}
	payload, _ := json.Marshal(data)
	c, resp, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, payload)
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(resp.Status)
		fmt.Println(err)
	}
	file, _ = os.Open("logs.txt")
	fileStr := ReadLine(file, lineCount+1)
	require.Equal(t, "warning 5 42: Test 4 map[6:78]", fileStr[20:])
}

func Test_Server_5(t *testing.T) {
	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	data := testMessage{
		Severity: "warning",
		Text:     "Test_5",
		Meta:     map[string]string{"user_id": "", "session_id": "", "6": "157"},
	}
	payload, _ := json.Marshal(data)
	c, resp, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, payload)
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(resp.Status)
		fmt.Println(err)
	}
	file, _ = os.Open("logs.txt")
	fileStr := ReadLine(file, lineCount+1)
	require.Equal(t, "warning  : Test_5 map[6:157]", fileStr[20:])
}

func Test_Server_6(t *testing.T) {
	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	data := "Sgdfg"
	c, resp, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, []byte(data))
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(resp.Status)
		fmt.Println(err)
	}
	file, _ = os.Open("logs.txt")
	fileStr := ReadLine(file, lineCount+1)
	require.Equal(t, "Error!!!", fileStr)
}
