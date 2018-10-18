package main

//important packages
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Endpoint data structure to hold person data
type Endpoints struct {
	Endpoint string `json:"endpoint,omitempty"`
	Keys     *Keys  `json:"keys,omitempty"`
}
type Keys struct {
	Auth   string `json:"auth,omitempty"`
	P256dh string `json:"p256dh,omitempty"`
}

var endpoint []Endpoints

var (
	PushApi []byte
)

func JsonResponse(m string) []byte {
	data := make(map[string]interface{})
	data["msg"] = m
	jsonOut, _ := json.Marshal(data)
	return jsonOut
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(time.Now(), r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*req).Method == "OPTIONS" {
		return
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SetupResponse(&w, r)
	if r.Method != "POST" {
		w.Write(JsonResponse("method not allowed"))
		return
	}

	var endpoints Endpoints
	_ = json.NewDecoder(r.Body).Decode(&endpoints)
	fmt.Println(endpoints)
	endpoint = append(endpoint, endpoints)
	json.NewEncoder(w).Encode(endpoints)
	val, _ := json.Marshal(endpoints)
	PushApi = val
	//fmt.Println(string(val))
	fmt.Println(string(PushApi))

}

// func QuerySaveSQLite(w http.ResponseWriter, r *http.Request) {
// 	db, err := SQLiteConn()
// 	defer db.Close()
// 	var asset Endpoints
// 	var LastInsertId string
// 	log.Println(r.Body)
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // TODO.....
// 	log.Println("Request body", r.Body)              // TODO.....
// 	obj := map[string]interface{}{}
// 	if err := json.Unmarshal([]byte(body), &obj); err != nil {
// 		log.Fatal(err)
// 	}
// 	myOid, ok := obj["body"]
// 	var myJSON map[string]interface{}
// 	if ok {
// 		switch v := myOid.(type) {
// 		case map[string]interface{}:
// 			myJSON = v
// 		default:
// 			log.Println()
// 		}
// 	}
// 	log.Println(myJSON["siteOid"].(string))

// 	json.NewDecoder(r.Body).Decode(&asset)
// 	log.Println("Encoded Asset", &asset)
// 	stmt, err := db.Prepare(`insert into asset (
// 		oid, organizationOid, customerOid, siteOid,
// 		categoryOid, manufacturerOid,modelOid, assetName,
// 		productSerial, assetID, purchaseDate )values (
// 		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
// 	if err != nil {
// 		panic(err)
// 	}
// 	res, err := stmt.Exec(myJSON["oid"], myJSON["organizationOid"], myJSON["customerOid"],
// 		myJSON["siteOid"], myJSON["categoryOid"], myJSON["manufacturerOid"],
// 		myJSON["modelOid"], myJSON["assetName"], myJSON["productSerial"],
// 		myJSON["assetID"], myJSON["purchaseDate"])
// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(id)
// 	log.Println(LastInsertId)
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(asset)
// 	log.Println("After JSON")

// }

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/endpoint/post", CreatePerson)
	http.ListenAndServe(":8075", Logger(r))
}

func init() {
	fmt.Println(string(PushApi))
}
