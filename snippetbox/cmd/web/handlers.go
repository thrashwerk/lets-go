package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/thrashwerk/lets-go/snippetbox/internal/models"
	"github.com/thrashwerk/lets-go/snippetbox/internal/validator"
)

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a form for signing up a new user...")
	// data := app.newTemplateData(r)

	// data.Form = userSignupForm{}

	// app.render(w, r, http.StatusOK, "signup.tmpl.html", data)
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
	// var form userSignupForm

	// err := app.decodePostForm(r, &form)
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	// form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	// form.CheckField(validator.MaxChars(form.Name, 30), "name", "This field cannot be more than 30 characters long")
	// form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	// form.CheckField(validator.MaxChars(form.Email, 30), "email", "This field cannot be more than 30 characters long")
	// form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	// form.CheckField(validator.MaxChars(form.Password, 100), "name", "This field cannot be more than 100 characters long")

	// if !form.Valid() {
	// 	data := app.newTemplateData(r)
	// 	data.Form = form
	// 	app.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl.html", data)
	// 	return
	// }

	// app.sessionManager.Put(r.Context(), "flash", "Account successfully created!")

	// http.Redirect(w, r, fmt.Sprintf("SOMETHING"), http.StatusSeeOther)
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a form for logging in a user...")
	// data := app.newTemplateData(r)

	// data.Form = userLoginForm{}

	// app.render(w, r, http.StatusOK, "login.tmpl.html", data)
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
	// var form userLoginForm

	// err := app.decodePostForm(r, &form)
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	// form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	// form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")

	// if !form.Valid() {
	// 	data := app.newTemplateData(r)
	// 	data.Form = form
	// 	app.render(w, r, http.StatusUnprocessableEntity, "login.tmpl.html", data)
	// 	return
	// }

	// app.sessionManager.Put(r.Context(), "SOMETHING", "SOMETHING")

	// http.Redirect(w, r, fmt.Sprintf("SOMETHING"), http.StatusSeeOther) // redirect to home?
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
