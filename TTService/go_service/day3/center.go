package main
import(
	"encoding/json"
	"fmt"
)

type User struct {
	Name string  `json:"name"`
	ID int `json:"id"`
	Age int `json:"age"`
}

func main() {
	user := User{ID: 1001, Name: "jo", Age: 24}
	data, _ := json.Marshal(user)
	fmt.Println(string(data))

	//json->结构体
	jsonStr := `{"name":"tom","age":29,"id":1002}`
	var u User
	json.Unmarshal([]byte(jsonStr), &u)
	fmt.Println(u)
}