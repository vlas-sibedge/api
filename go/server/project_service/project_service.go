// Package project_service provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package project_service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

const (
	ServiceKeyScopes  = "ServiceKey.Scopes"
	ServiceNameScopes = "ServiceName.Scopes"
)

// AuthMethod defines model for AuthMethod.
type AuthMethod struct {
	CreatedAt *time.Time              `json:"created_at,omitempty"`
	Id        *string                 `json:"id,omitempty"`
	Name      *string                 `json:"name,omitempty"`
	ProjectId *string                 `json:"project_id,omitempty"`
	Settings  *map[string]interface{} `json:"settings,omitempty"`
	Type      *string                 `json:"type,omitempty"`
	UpdatedAt *time.Time              `json:"updated_at,omitempty"`
}

// CreateAuthMethodParams defines model for CreateAuthMethodParams.
type CreateAuthMethodParams struct {
	Name      *string                 `json:"name"`
	ProjectId *string                 `json:"project_id"`
	Settings  *map[string]interface{} `json:"settings,omitempty"`
	Type      *string                 `json:"type"`
}

// CreateProjectParams defines model for CreateProjectParams.
type CreateProjectParams struct {
	Description *map[string]interface{} `json:"description"`
	Name        *string                 `json:"name"`
}

// FindAuthMethodParams defines model for FindAuthMethodParams.
type FindAuthMethodParams struct {
	CreatedAt *time.Time              `json:"created_at"`
	Name      *string                 `json:"name,omitempty"`
	ProjectId *string                 `json:"project_id,omitempty"`
	Settings  *map[string]interface{} `json:"settings,omitempty"`
	Type      *string                 `json:"type,omitempty"`
	UpdatedAt *time.Time              `json:"updated_at"`
}

// Location defines model for Location.
type Location struct {
	ClientSelector map[string]interface{} `json:"client_selector"`
	Id             string                 `json:"id"`
	Labels         map[string]interface{} `json:"labels"`
	Name           string                 `json:"name"`
	NodeSelector   map[string]interface{} `json:"node_selector"`
}

// LocationMapping defines model for LocationMapping.
type LocationMapping struct {
	Mapping *map[string]interface{} `json:"mapping,omitempty"`
}

// PatchAuthMethodParams defines model for PatchAuthMethodParams.
type PatchAuthMethodParams struct {
	Name      *string                 `json:"name,omitempty"`
	ProjectId *string                 `json:"project_id,omitempty"`
	Settings  *map[string]interface{} `json:"settings,omitempty"`
	Type      *string                 `json:"type,omitempty"`
	UpdatedAt *time.Time              `json:"updated_at"`
}

// PatchProjectParams defines model for PatchProjectParams.
type PatchProjectParams struct {
	Description *map[string]interface{} `json:"description,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	UpdatedAt   *time.Time              `json:"updated_at"`
}

// Project defines model for Project.
type Project struct {
	CreatedAt   *time.Time              `json:"created_at,omitempty"`
	Description *map[string]interface{} `json:"description,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	UpdatedAt   *time.Time              `json:"updated_at,omitempty"`
}

// ProjectConfig defines model for ProjectConfig.
type ProjectConfig struct {
	Id        string     `json:"id"`
	Locations []Location `json:"locations"`
	Name      string     `json:"name"`
}

// UpdateAuthMethodParams defines model for UpdateAuthMethodParams.
type UpdateAuthMethodParams struct {
	Name      *string                 `json:"name"`
	ProjectId *string                 `json:"project_id"`
	Settings  *map[string]interface{} `json:"settings"`
	Type      *string                 `json:"type"`
	UpdatedAt *time.Time              `json:"updated_at"`
}

// UpdateProjectParams defines model for UpdateProjectParams.
type UpdateProjectParams struct {
	Description *map[string]interface{} `json:"description"`
	Name        *string                 `json:"name"`
	UpdatedAt   *time.Time              `json:"updated_at"`
}

// ListAuthMethodParams defines parameters for ListAuthMethod.
type ListAuthMethodParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// CreateAuthMethodJSONBody defines parameters for CreateAuthMethod.
type CreateAuthMethodJSONBody CreateAuthMethodParams

// PatchAuthMethodJSONBody defines parameters for PatchAuthMethod.
type PatchAuthMethodJSONBody PatchAuthMethodParams

// UpdateAuthMethodJSONBody defines parameters for UpdateAuthMethod.
type UpdateAuthMethodJSONBody UpdateAuthMethodParams

// FindAuthMethodJSONBody defines parameters for FindAuthMethod.
type FindAuthMethodJSONBody FindAuthMethodParams

// ListProjectParams defines parameters for ListProject.
type ListProjectParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// CreateProjectJSONBody defines parameters for CreateProject.
type CreateProjectJSONBody CreateProjectParams

// PatchProjectJSONBody defines parameters for PatchProject.
type PatchProjectJSONBody PatchProjectParams

// UpdateProjectJSONBody defines parameters for UpdateProject.
type UpdateProjectJSONBody UpdateProjectParams

// CreateAuthMethodJSONRequestBody defines body for CreateAuthMethod for application/json ContentType.
type CreateAuthMethodJSONRequestBody CreateAuthMethodJSONBody

// PatchAuthMethodJSONRequestBody defines body for PatchAuthMethod for application/json ContentType.
type PatchAuthMethodJSONRequestBody PatchAuthMethodJSONBody

// UpdateAuthMethodJSONRequestBody defines body for UpdateAuthMethod for application/json ContentType.
type UpdateAuthMethodJSONRequestBody UpdateAuthMethodJSONBody

// FindAuthMethodJSONRequestBody defines body for FindAuthMethod for application/json ContentType.
type FindAuthMethodJSONRequestBody FindAuthMethodJSONBody

// CreateProjectJSONRequestBody defines body for CreateProject for application/json ContentType.
type CreateProjectJSONRequestBody CreateProjectJSONBody

// PatchProjectJSONRequestBody defines body for PatchProject for application/json ContentType.
type PatchProjectJSONRequestBody PatchProjectJSONBody

// UpdateProjectJSONRequestBody defines body for UpdateProject for application/json ContentType.
type UpdateProjectJSONRequestBody UpdateProjectJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List auth methods
	// (GET /api/project-service/auth-method)
	ListAuthMethod(w http.ResponseWriter, r *http.Request, params ListAuthMethodParams)
	// Add auth method
	// (POST /api/project-service/auth-method)
	CreateAuthMethod(w http.ResponseWriter, r *http.Request)
	// Delete a auth method
	// (DELETE /api/project-service/auth-method/{id})
	DeleteAuthMethod(w http.ResponseWriter, r *http.Request, id string)
	// Get auth method
	// (GET /api/project-service/auth-method/{id})
	GetAuthMethod(w http.ResponseWriter, r *http.Request, id string)
	// Patch auth method
	// (PATCH /api/project-service/auth-method/{id})
	PatchAuthMethod(w http.ResponseWriter, r *http.Request, id string)
	// Update auth method
	// (PUT /api/project-service/auth-method/{id})
	UpdateAuthMethod(w http.ResponseWriter, r *http.Request, id string)
	// Find auth method
	// (GET /api/project-service/find-auth-method)
	FindAuthMethod(w http.ResponseWriter, r *http.Request)
	// Get location mapping
	// (GET /api/project-service/location-mapping)
	GetLocationMapping(w http.ResponseWriter, r *http.Request)
	// List projects
	// (GET /api/project-service/project)
	ListProject(w http.ResponseWriter, r *http.Request, params ListProjectParams)
	// Create project
	// (POST /api/project-service/project)
	CreateProject(w http.ResponseWriter, r *http.Request)
	// Project config list
	// (GET /api/project-service/project-config)
	ListProjectConfig(w http.ResponseWriter, r *http.Request)
	// Delete a project
	// (DELETE /api/project-service/project/{id})
	DeleteProject(w http.ResponseWriter, r *http.Request, id string)
	// Get project
	// (GET /api/project-service/project/{id})
	GetProject(w http.ResponseWriter, r *http.Request, id string)
	// Patch project
	// (PATCH /api/project-service/project/{id})
	PatchProject(w http.ResponseWriter, r *http.Request, id string)
	// Update project
	// (PUT /api/project-service/project/{id})
	UpdateProject(w http.ResponseWriter, r *http.Request, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// ListAuthMethod operation middleware
func (siw *ServerInterfaceWrapper) ListAuthMethod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListAuthMethodParams

	// ------------- Required query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------
	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListAuthMethod(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateAuthMethod operation middleware
func (siw *ServerInterfaceWrapper) CreateAuthMethod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateAuthMethod(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteAuthMethod operation middleware
func (siw *ServerInterfaceWrapper) DeleteAuthMethod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteAuthMethod(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetAuthMethod operation middleware
func (siw *ServerInterfaceWrapper) GetAuthMethod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAuthMethod(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PatchAuthMethod operation middleware
func (siw *ServerInterfaceWrapper) PatchAuthMethod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchAuthMethod(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateAuthMethod operation middleware
func (siw *ServerInterfaceWrapper) UpdateAuthMethod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateAuthMethod(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// FindAuthMethod operation middleware
func (siw *ServerInterfaceWrapper) FindAuthMethod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindAuthMethod(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetLocationMapping operation middleware
func (siw *ServerInterfaceWrapper) GetLocationMapping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetLocationMapping(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListProject operation middleware
func (siw *ServerInterfaceWrapper) ListProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListProjectParams

	// ------------- Required query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------
	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListProject(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateProject operation middleware
func (siw *ServerInterfaceWrapper) CreateProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateProject(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListProjectConfig operation middleware
func (siw *ServerInterfaceWrapper) ListProjectConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListProjectConfig(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteProject operation middleware
func (siw *ServerInterfaceWrapper) DeleteProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteProject(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetProject operation middleware
func (siw *ServerInterfaceWrapper) GetProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProject(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PatchProject operation middleware
func (siw *ServerInterfaceWrapper) PatchProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchProject(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateProject operation middleware
func (siw *ServerInterfaceWrapper) UpdateProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, ServiceKeyScopes, []string{""})

	ctx = context.WithValue(ctx, ServiceNameScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateProject(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/project-service/auth-method", wrapper.ListAuthMethod)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/project-service/auth-method", wrapper.CreateAuthMethod)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/project-service/auth-method/{id}", wrapper.DeleteAuthMethod)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/project-service/auth-method/{id}", wrapper.GetAuthMethod)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/api/project-service/auth-method/{id}", wrapper.PatchAuthMethod)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/project-service/auth-method/{id}", wrapper.UpdateAuthMethod)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/project-service/find-auth-method", wrapper.FindAuthMethod)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/project-service/location-mapping", wrapper.GetLocationMapping)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/project-service/project", wrapper.ListProject)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/project-service/project", wrapper.CreateProject)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/project-service/project-config", wrapper.ListProjectConfig)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/project-service/project/{id}", wrapper.DeleteProject)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/project-service/project/{id}", wrapper.GetProject)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/api/project-service/project/{id}", wrapper.PatchProject)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/project-service/project/{id}", wrapper.UpdateProject)
	})

	return r
}
