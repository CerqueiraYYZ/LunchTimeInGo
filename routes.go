package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"LunchIndex",
		"GET",
		"/lunches",
		LunchIndex,
	},
	Route{
		"LunchShow",
		"GET",
		"/lunches/{lunchId}",
		LunchShow,
	},
	Route{
		"EmployeeFoodTime",
		"GET",
		"/EmployeeFoodTime/{name}",
		EmployeeFoodTime,
	},
	Route{
		"LunchCreate",
		"POST",
		"/lunches",
		LunchCreate,
	},
}
