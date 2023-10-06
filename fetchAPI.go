package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/gorilla/mux"
)

func fetchData(apiURL string) ([]byte, error) {
    response, err := http.Get(apiURL)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    return data, nil
}

func main() {
    router := mux.NewRouter()

    apiURL := "https://jsonplaceholder.typicode.com/posts/1" // Example API URL

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data, err := fetchData(apiURL)
        if err != nil {
            http.Error(w, "Error fetching data", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(data)
    })

    port := 8080
    fmt.Printf("Server is running on localhost:%d...\n", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
