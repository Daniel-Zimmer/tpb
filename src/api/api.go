package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Search(search string) []SearchRes {

	var response []SearchRes

	r, _ := http.Get("https://apibay.org/q.php?q=" + search + "&cat=0")
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &response)

	return response
}

func Page(id string) PageRes {
	var response PageRes

	r, _ := http.Get("https://apibay.org/t.php?id=" + id)
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &response)

	return response
}