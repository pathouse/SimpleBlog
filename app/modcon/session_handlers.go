package modcon

import (
	"net/http"
	"regexp"
	"simpleblog/app/support"
)

func RegistrationHandler(context *appContext, resp http.ResponseWriter, req *http.Request) (int, error) {
	p := &Page{
		Title:     "Register",
		Bodyclass: "registration",
	}
	if err := context.appTemplates.ExecuteTemplate(resp, "registration", p); err != nil {
		support.LogStacktrace(err)
		return http.StatusInternalServerError, err
	}
	return 200, nil
}

func CreateUserHandler(context *appContext, resp http.ResponseWriter, req *http.Request) {
	// TODO - validations
	user := &User{
		FirstName: req.FormValue("FirstName"),
		LastName:  req.FormValue("LastName"),
		Email:     req.FormValue("Email"),
	}

	// TODO - move this logic elsewhere?
	pass, pass_conf := req.FormValue("Password"), req.FormValue("PasswordConfirmation")
	pass_ok, err := regexp.MatchString(pass_conf, pass)
	if !pass_ok || err != nil {
		http.Redirect(resp, req, "/register", 301)
	}
	// TODO - return form values in case of validation failure

	if err := user.GenerateHashedPassword(pass); err != nil {
		http.Redirect(resp, req, "/register", 301)
	}

	if err := CreateRecord(context.db, user); err != nil {
		http.Redirect(resp, req, "/register", 301)
	}

	// TODO - make sure using correct redirect response code
	http.Redirect(resp, req, "/", 301)
}
