package handler

import (
	"assignment2/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetBreedByCountry(c *gin.Context) {
	response, err := http.Get("https://catfact.ninja/breeds")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	reponsedata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "failed to read the file",
			"error":   err.Error(),
		})
		return 
	}
	err = ioutil.WriteFile("new.txt", reponsedata, 0644)
	if err != nil {
		fmt.Println("failed to write file")
	}
	var DataResponse models.Response
	_ = json.Unmarshal(reponsedata, &DataResponse)

	//Printing total page ------------------------------------------------------------------------------------------------

	fmt.Println("Total Pages is :    ", DataResponse.LastPage) // there is also two options are available | totalpage/perpage |
//------------------------------------------------------------------------------------------------------------------------
	catBreedByCountry := make(models.CatBreedByCountry)
	fmt.Println(catBreedByCountry)
	check := make(map[string]bool) 				//checking the page number is   already used or not
									  			// we can iterate the page 1 to last page add the page in end point
	for _, v := range DataResponse.Links {
		if !check[v.Url] {
			check[v.Url] = true
			fmt.Println(v.Url)
			responseBreed, err := http.Get(string(v.Url))
			if err != nil {
				continue
			}
			defer responseBreed.Body.Close()
			breedBody, err := ioutil.ReadAll(responseBreed.Body)
			if err != nil {
				fmt.Println("failed to read the beeds page", err)
				continue
			}
			var breedPage models.Response
			err = json.Unmarshal(breedBody, &breedPage)
			if err!=nil{
				continue
			}
			for _, breed := range breedPage.Breeds {
				var res models.CatBreedResponseByCountry
				res.Breed, res.Origin, res.Coat, res.Pattern = breed.Breed, breed.Origin, breed.Coat, breed.Pattern
				country := breed.Country
				catBreedByCountry[country] = append(catBreedByCountry[country], res)
			}
		}
	}
	c.JSON(201, gin.H{
		"message": "sucess",
		"breeds":  catBreedByCountry,
	})
}
