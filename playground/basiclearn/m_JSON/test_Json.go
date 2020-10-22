package m_JSON

import (
	"encoding/json"
	"fmt"
)

var jsonStr string = `{"server_name":"ihui","server_ip":"127.0.0.1","server_port":8080}`

// 这里必须首字母大写  不然 拿不到数据
type Server struct {
	ServerName string `json:"server_name"`
	ServerIp   string `json:"server_ip"`
	ServerPort int    `json:"server_port"`
}

func serialize() {
	s := Server{
		ServerName: "ihui",
		ServerIp:   "127.0.0.1",
		ServerPort: 8080,
	}
	json, err := json.Marshal(s) // 返回的 []byte error
	if err != nil {
		fmt.Println("serialize error", err)
		return
	}
	fmt.Println("json: ", string(json))
}

func serializeMap() {
	m := make(map[string]interface{})
	m["servername"] = "test"
	m["server_ip"] = "192.168.1.1"
	m["serverPort"] = 8080
	json, err := json.Marshal(m) // 返回的 []byte error
	if err != nil {
		fmt.Println("serialize error", err)
		return
	}
	fmt.Println("json: ", string(json))
}

func deSerialize() {
	s2 := &Server{}
	err := json.Unmarshal([]byte(jsonStr), s2) // 返回的 []byte error
	if err != nil {
		fmt.Println("serialize error", err)
		return
	}
	fmt.Println("server: ", s2)
}

func deSerializeMap() {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m) // 返回的 []byte error
	if err != nil {
		fmt.Println("serialize error", err)
		return
	}
	fmt.Println("server: ", m)
}

// 结构体的序列化  map的序列化
func TestJson() {
	//serialize()
	//serializeMap()
	//deSerialize()
	deSerializeMap()
}
