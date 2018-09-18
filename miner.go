package main

import (
        "fmt"; 
        "encoding/json"; 
        "crypto/sha256"; 
        "strconv"; 
        "net/http";
        "encoding/hex";
        "strings";
        "io/ioutil";
        "bytes"
        )
type BlockBody struct {
  Id string   `json:"id"`
  Hash string `json:"prime"`
}

type Data struct {
  Data BlockBody `json:"data"`
}

type Block struct {
  Prime int
}

func isPrime(n int) bool {
  if (n % 2 == 0) { return false }
  for i := 2; i < n; i++ {
    if(n % i == 0) { return false }
  }
  return true
}

func getNextPrime(n int) int {
  status := false
  for status != true {
    if(isPrime(n)) {return n}
    n++
  }
  return n
}

func getLastPrime(url string) int {
  r, err := http.Get(url)
  var blocks []Block
  body, err := ioutil.ReadAll(r.Body)
  if(err != nil) { return 0 }
  json.Unmarshal(body, &blocks)
  return blocks[len(blocks) - 1].Prime
}

func postPrimeHash(number int, url string) {
  stringPrime := strconv.Itoa(number)
  h           := sha256.New()
  h.Write([]byte(stringPrime))
  byteHash := h.Sum(nil)
  input    :=  Data { Data: BlockBody{ Id: "aldous", Hash: strings.ToUpper(hex.EncodeToString(byteHash)) }} 

  var out []byte
  out, err := json.Marshal(input);
  if err != nil { fmt.Println(0)}
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(out))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
}

func main() {
  url := "https://node2-dot-crypto-challenge-fall-2018.appspot.com/blocks"
  lastPrime := getLastPrime(url)
  fmt.Println(lastPrime)

  n := 0
  prime := lastPrime + 1
  for n < 1 {
    prime = getNextPrime(prime)
    postPrimeHash(prime, url)
    prime++
  }
}
