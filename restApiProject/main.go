package main
/*
Consuming REST APIs in Go - HTTP GET, PUT, POST and DELETE
-->to consume an external api in go
-->write application that fetch information from external APIs
with using jsonplaceholder external APIs
*/
/*we will see how to get the API response as a string &
how to map the response to a struct */
import(
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*create todo struct representing the json response
which information returned in the  API response*/
type Todo struct{
	UserID int `json:"user_Id"`
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`

}

func get(){
	fmt.Println("1. Performing Http Get...")
	//get method takes the url as the argument and returns a response n error
	//api call successfully then err will be nil
	resp,err :=http.Get("https://jsonplaceholder.typicode.com/todos/23")
	if err !=nil{
		log.Fatalln(err)
	}  /*if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}*/
	defer resp.Body.Close() //executed at the end of method

	bodyBytes,_:=ioutil.ReadAll(resp.Body)//returns resp body as a slice of bytes

	//Convert response body to string
	bodyString :=string(bodyBytes)
	fmt.Println("API Response as String:\n"+bodyString)

	//Convert response body to TOdo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes,&todoStruct)
	fmt.Printf("API Response as struct %+v\n",todoStruct)

}
//Post Request ->create a new resource on the server
func  post(){
	fmt.Println("2.Performing Http Post...")
	todo :=Todo{1,40,"lorem ipsum dolor sit amet",true}
	jsonReq ,err:=json.Marshal(todo) //converting the struct to a slice of bytes
	resp,err :=http.Post("https://jsonplaceholder.typicode.com/todos","application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	/*In order to use the jsonReq with the post method, we need to convert that to an instance of io.Reader,
	and this is where the bytes.NewBuffer() method */
	if err !=nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes,_ :=ioutil.ReadAll(resp.Body) //returns resp body as a slice of bytes

	//Convert response body to string
	bodyString :=string(bodyBytes)
	fmt.Println(bodyString)

	//Convert  response body to ToDo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes,&todoStruct)
	fmt.Printf("%+v\n",todoStruct)

}
//modify the data a resource  on the server
func put() {
	fmt.Println("3. Performing Http Put...")
	todo := Todo{1, 40, "have you look on it", true}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/40", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
}
func delete() {
	fmt.Println("4. Performing Http Delete...")
	todo := Todo{1, 40, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/40", bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func main(){
	get()
	post()
	put()
	//delete()
}

/*The Post method takes 3 arguments - The Url of the API,
the content-type, which is application/json in our case ,
and an instance of io.Reader. In order to use the jsonReq with the post method, we need to convert that to an instance of io.Reader,
and this is where the bytes.NewBuffer()
method comes into picture

The bytes.NewBuffer() takes the jsonReq bytes slice as input and returns a Buffer initialized from the contents of our jsonReq. As we can see from the buffer.go source, buffer has a read method which means it implements the io.Reader interface and it can be passed as an argument to the http.Post() method call.*/