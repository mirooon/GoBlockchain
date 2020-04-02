		

		package main

		import (
			"fmt"
			"net/http"
			"encoding/json"
			"io/ioutil"
		)
		
		type Block struct {
			BlockNumber  int
			Timestamp    int64
			Transactions []Transaction
			Nonce        int
			PreviousHash string
		}

		type Transaction struct {
			SenderPublicKey    string  `json:"SenderPublicKey"`
			RecipientPublicKey string  `json:"RecipientPublicKey"`
			Signature          string  `json:"Signature"`
			Amount             float32 `json:"Amount"`
		}

		func main() {
			resp, err := http.Get("http://" + "localhost:5001" + "/chain")
			if err != nil {
				fmt.Printf("%v\n", "Problem with connection with ip")
				fmt.Printf("%v\n", "err")
				fmt.Printf("%v\n", err)
			}
			var responseObj struct {
				Chain  []Block
				Length int
			}
			body, err := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(body, &responseObj)
			if err != nil {
				panic(err)
			}
			length := responseObj.Length
			chain := responseObj.Chain
		}