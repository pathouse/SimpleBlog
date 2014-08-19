package modcon

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func PostsIndexHandler(context *appContext, resp http.ResponseWriter, req *http.Request) (int, error) {
	posts := &[]Post{}
	context.db.Find(posts)

	jsonResponse, err := json.Marshal(posts)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprint(resp, string(jsonResponse[:]))

	return http.StatusOK, nil
}

func PostsShowHandler(context *appContext, resp http.ResponseWriter, req *http.Request) (int, error) {
	postId := mux.Vars(req)["id"]

	post := &Post{}
	query := map[string]interface{}{"id": postId}
	err := FindByMap(context.db, query, post, true)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	jsonResponse, jsonErr := json.Marshal(post)
	if jsonErr != nil {
		return http.StatusInternalServerError, jsonErr
	}

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprint(resp, string(jsonResponse[:]))

	return http.StatusOK, nil
}

func PostsCreateHandler(context *appContext, resp http.ResponseWriter, req *http.Request) (int, error) {
	decoder := json.NewDecoder(req.Body)
	post := &Post{}
	err := decoder.Decode(post)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// remarshalling and sending back for testing purposes
	jsonResponse, jsonErr := json.Marshal(post)
	if jsonErr != nil {
		return http.StatusInternalServerError, jsonErr
	}

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprint(resp, string(jsonResponse[:]))

	return http.StatusOK, nil
}

func PostsUpdateHandler(context *appContext, resp http.ResponseWriter, req *http.Request) (int, error) {
	postId := mux.Vars(req)["id"]

	post := &Post{}
	query := map[string]interface{}{"id": postId}
	err := FindByMap(context.db, query, post, true)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	decoder := json.NewDecoder(req.Body)
	requestArgs := &Post{}
	decodeErr := decoder.Decode(requestArgs)
	if decodeErr != nil {
		return http.StatusInternalServerError, decodeErr
	}

	// remarshalling and returning both for testing purposes
	posts := []*Post{post, requestArgs}
	jsonResponse, jsonErr := json.Marshal(posts)
	if jsonErr != nil {
		return http.StatusInternalServerError, jsonErr
	}

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprint(resp, string(jsonResponse[:]))

	return http.StatusOK, nil
}

func PostsDeleteHandler(context *appContext, resp http.ResponseWriter, req *http.Request) (int, error) {
	postId := mux.Vars(req)["id"]

	post := &Post{}
	query := map[string]interface{}{"id": postId}
	err := FindByMap(context.db, query, post, true)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// marshal and return for testing
	jsonResponse, jsonErr := json.Marshal(post)
	if jsonErr != nil {
		return http.StatusInternalServerError, jsonErr
	}

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprint(resp, string(jsonResponse[:]))

	return http.StatusOK, nil
}
