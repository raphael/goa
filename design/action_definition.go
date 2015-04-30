package design

import (
	"fmt"
	"regexp"
)

// A resource action
// Defines an HTTP endpoint and the shape of HTTP requests and responses made to
// that endpoint.
// The shape of requests is defined via "parameters", there are path parameters
// (i.e. portions of the URL that define parameter values), query string
// parameters and a payload parameter (request body).
type ActionDefinition struct {
	Name        string                // Action name, e.g. "create"
	Description string                // Action description, e.g. "Creates a task"
	Resource    *ResourceDefinition   // Resource containing action
	Routes      []*Route              // Action routes
	Responses   []*ResponseDefinition // Set of possible response definitions
	Params      ActionParams          // Path parameters if any
	Payload     *Member               // Payload blueprint (request body) if any
	Headers     []*Header             // Special request headers that need to be made available to action
}

// An action route
type Route struct {
	Verb string
	Path string
}

// Regular expression used to capture path parameters
var pathRegex = regexp.MustCompile("/:([^/]+)")

// Internal helper method that sets HTTP method, path and path params
func (a *Action) method(method, path string) *Action {
	a.HttpMethod = method
	a.Path = path
	var matches = pathRegex.FindAllStringSubmatch(path, -1)
	a.PathParams = make(map[string]*ActionParam, len(matches))
	for _, m := range matches {
		mem := Member{Type: String}
		a.PathParams[m[1]] = &ActionParam{Name: m[1], Member: &mem}
	}
	return a
}

// Validates that action definition is consistent: parameters have unique names, has at least one
// response.
func (a *Action) validate() error {
	if a.Name == "" {
		return fmt.Errorf("Action name cannot be empty")
	}
	for i, r := range a.Responses {
		for j, r2 := range a.Responses {
			if i != j && r.Status == r2.Status {
				return fmt.Errorf("Multiple response definitions with status code %d", r.Status)
			}
		}
	}
	for _, p := range a.PathParams {
		for _, q := range a.QueryParams {
			if p.Name == q.Name {
				return fmt.Errorf("Action has both path parameter and query parameter named %s",
					p.Name)
			}
		}
	}
	if err := a.validateParams(true); err != nil {
		return err
	}
	if err := a.validateParams(false); err != nil {
		return err
	}
	return nil
}

// Validate action parameters (make sure they have names, members and types)
func (a *Action) validateParams(isPath bool) error {
	var params ActionParams
	if isPath {
		params = a.PathParams
	} else {
		params = a.QueryParams
	}
	for n, p := range params {
		if n == "" {
			return fmt.Errorf("%s has parameter with no name", a.Name)
		} else if p.Member == nil {
			return fmt.Errorf("Member field of %s parameter :%s cannot be nil",
				a.Name, n)
		} else if p.Member.Type == nil {
			return fmt.Errorf("type of %s parameter :%s cannot be nil",
				a.Name, n)
		}
	}
	return nil
}
