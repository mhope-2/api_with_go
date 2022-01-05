package main

import "net/http"

// a handler which writes plain-text response with information
// about the application status, os environment and version

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// a map to hold the information to be sent in the response
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	// pass the map to writeJSON function
	err := app.writeJSON(w, http.StatusOK, env, nil)

	if err != nil {
		// log a server error
		app.serverErrorResponse(w, r, err)
		return
	}
}
