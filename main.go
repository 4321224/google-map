package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    // Replace with your Google Maps API key
    apiKey := "AIzaSyAfcOxFQ-eKxWz46VpJlUL7vH0cTn1VPvI"
    address := "Jl. Pisangan Raya No.20, Cireundeu, Kec. Ciputat Tim., Kota Tangerang Selatan, Banten 15419"

    // Encode the address for the URL
    encodedAddress := url.QueryEscape(address)

    // Build the URL for the Geocoding API request
    apiUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", encodedAddress, apiKey)

    // Send the HTTP request
    response, err := http.Get(apiUrl)
    if err != nil {
        fmt.Println("Error making the request:", err)
        return
    }
    defer response.Body.Close()

    // Read the response body
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading the response:", err)
        return
    }

    // Define a struct that matches the structure of the JSON response
    type GeocodingResponse struct {
        Results []struct {
            Geometry struct {
                Location struct {
                    Lat float64 `json:"lat"`
                    Lng float64 `json:"lng"`
                } `json:"location"`
            } `json:"geometry"`
        } `json:"results"`
    }

    // Parse the JSON response into the defined struct
    var geocodingResponse GeocodingResponse
    if err := json.Unmarshal(data, &geocodingResponse); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    // Access the parsed data
    if len(geocodingResponse.Results) > 0 {
        lat := geocodingResponse.Results[0].Geometry.Location.Lat
        lng := geocodingResponse.Results[0].Geometry.Location.Lng
        fmt.Printf("Latitude: %f, Longitude: %f\n", lat, lng)
    } else {
        fmt.Println("No results found.")
    }
}
