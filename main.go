package main
import(
	"fmt"
	"io/ioutil"
    "log"
    "net/http"
	"os"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
    Name    string    `json:"name"`
    Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
    EntryNo int            `json:"entry_number"`
    Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
    Name string `json:"name"`
}

var responseObject Response

func main (){
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	//fmt.Println(string(responseData))
	json.Unmarshal(responseData, &responseObject)

	//fmt.Println(responseObject.Name)
	//fmt.Println(len(responseObject.Pokemon))
	
	db, err := sql.Open("mysql", "root:root987@tcp(127.0.0.1:3306)/pokemon")   
    if err != nil {
        panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Prepare("Insert into kanto (Name) values (?)")
	if err != nil {
        panic(err.Error())
	}
	for i := 0; i < len(responseObject.Pokemon); i++ {
		//fmt.Println(responseObject.Pokemon[i].Species.Name)
		insert.Exec(responseObject.Pokemon[i].Species.Name)
	}
	defer insert.Close()

}