// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Ageing defines model for Ageing.
type Ageing struct {
	Cask         string              `json:"cask"`
	CaskUuid     openapi_types.UUID  `json:"caskUuid"`
	EndTime      *time.Time          `json:"endTime,omitempty"`
	Hash         *string             `json:"hash,omitempty"`
	PreviousUuid *openapi_types.UUID `json:"previousUuid,omitempty"`
	StartTime    time.Time           `json:"startTime"`
	Tank         string              `json:"tank"`
	TankUuid     openapi_types.UUID  `json:"tankUuid"`
	Transaction  *string             `json:"transaction,omitempty"`
	Uuid         openapi_types.UUID  `json:"uuid"`
	Winery       string              `json:"winery"`
	WineryUuid   openapi_types.UUID  `json:"wineryUuid"`
}

// Ageings defines model for Ageings.
type Ageings struct {
	Receptions *[]Ageing `json:"receptions,omitempty"`
}

// Bottling defines model for Bottling.
type Bottling struct {
	BottleQty    int32               `json:"bottleQty"`
	Cask         string              `json:"cask"`
	CaskUuid     openapi_types.UUID  `json:"caskUuid"`
	EndTime      *time.Time          `json:"endTime,omitempty"`
	Hash         *string             `json:"hash,omitempty"`
	PreviousUuid *openapi_types.UUID `json:"previousUuid,omitempty"`
	StartTime    time.Time           `json:"startTime"`
	Transaction  *string             `json:"transaction,omitempty"`
	Uuid         openapi_types.UUID  `json:"uuid"`
	Wine         string              `json:"wine"`
	WineUuid     openapi_types.UUID  `json:"wineUuid"`
	Winery       string              `json:"winery"`
	WineryUuid   openapi_types.UUID  `json:"wineryUuid"`
}

// Bottlings defines model for Bottlings.
type Bottlings struct {
	Receptions *[]Bottling `json:"receptions,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// Fermentation defines model for Fermentation.
type Fermentation struct {
	EndTime       *time.Time          `json:"endTime,omitempty"`
	Hash          *string             `json:"hash,omitempty"`
	PreviousUuid  *openapi_types.UUID `json:"previousUuid,omitempty"`
	StartTime     time.Time           `json:"startTime"`
	Tank          string              `json:"tank"`
	TankUuid      openapi_types.UUID  `json:"tankUuid"`
	Transaction   *string             `json:"transaction,omitempty"`
	Uuid          openapi_types.UUID  `json:"uuid"`
	Warehouse     string              `json:"warehouse"`
	WarehouseUuid openapi_types.UUID  `json:"warehouseUuid"`
	Winery        string              `json:"winery"`
	WineryUuid    openapi_types.UUID  `json:"wineryUuid"`
}

// Fermentations defines model for Fermentations.
type Fermentations struct {
	Receptions *[]Fermentation `json:"receptions,omitempty"`
}

// Maceration defines model for Maceration.
type Maceration struct {
	EndTime       *time.Time          `json:"endTime,omitempty"`
	Hash          *string             `json:"hash,omitempty"`
	PreviousUuid  *openapi_types.UUID `json:"previousUuid,omitempty"`
	Reception     time.Time           `json:"reception"`
	ReceptionUuid openapi_types.UUID  `json:"receptionUuid"`
	StartTime     time.Time           `json:"startTime"`
	Transaction   *string             `json:"transaction,omitempty"`
	Uuid          openapi_types.UUID  `json:"uuid"`
	Warehouse     string              `json:"warehouse"`
	WarehouseUuid openapi_types.UUID  `json:"warehouseUuid"`
	Winery        string              `json:"winery"`
	WineryUuid    openapi_types.UUID  `json:"wineryUuid"`
}

// Macerations defines model for Macerations.
type Macerations struct {
	Receptions *[]Maceration `json:"receptions,omitempty"`
}

// PostAgeing defines model for PostAgeing.
type PostAgeing struct {
	CaskUuid   openapi_types.UUID `json:"caskUuid"`
	TankUuid   openapi_types.UUID `json:"tankUuid"`
	WineryUuid openapi_types.UUID `json:"wineryUuid"`
}

// PostBottling defines model for PostBottling.
type PostBottling struct {
	BottleQty  int32              `json:"bottleQty"`
	CaskUuid   openapi_types.UUID `json:"caskUuid"`
	WineUuid   openapi_types.UUID `json:"wineUuid"`
	WineryUuid openapi_types.UUID `json:"wineryUuid"`
}

// PostFermentation defines model for PostFermentation.
type PostFermentation struct {
	TankUuid      openapi_types.UUID `json:"tankUuid"`
	WarehouseUuid openapi_types.UUID `json:"warehouseUuid"`
	WineryUuid    openapi_types.UUID `json:"wineryUuid"`
}

// PostMaceration defines model for PostMaceration.
type PostMaceration struct {
	ReceptionUuid openapi_types.UUID `json:"receptionUuid"`
	WarehouseUuid openapi_types.UUID `json:"warehouseUuid"`
	WineryUuid    openapi_types.UUID `json:"wineryUuid"`
}

// PostReception defines model for PostReception.
type PostReception struct {
	GrapeTypeUuid openapi_types.UUID `json:"grapeTypeUuid"`
	Sugar         int32              `json:"sugar"`
	TruckUuid     openapi_types.UUID `json:"truckUuid"`
	VineyardUuid  openapi_types.UUID `json:"vineyardUuid"`
	Weight        int32              `json:"weight"`
	WineryUuid    openapi_types.UUID `json:"wineryUuid"`
}

// Reception defines model for Reception.
type Reception struct {
	EndTime       *time.Time         `json:"endTime,omitempty"`
	GrapeType     string             `json:"grapeType"`
	GrapeTypeUuid openapi_types.UUID `json:"grapeTypeUuid"`
	Hash          *string            `json:"hash,omitempty"`
	StartTime     time.Time          `json:"startTime"`
	Sugar         int32              `json:"sugar"`
	Transaction   *string            `json:"transaction,omitempty"`
	Truck         string             `json:"truck"`
	TruckUuid     openapi_types.UUID `json:"truckUuid"`
	Uuid          openapi_types.UUID `json:"uuid"`
	Vineyard      string             `json:"vineyard"`
	VineyardUuid  openapi_types.UUID `json:"vineyardUuid"`
	Weight        int32              `json:"weight"`
	Winery        string             `json:"winery"`
	WineryUuid    openapi_types.UUID `json:"wineryUuid"`
}

// Receptions defines model for Receptions.
type Receptions struct {
	Receptions []Reception `json:"receptions"`
}

// Status defines model for Status.
type Status struct {
	Message string `json:"message"`
}

// RegisterAgeingJSONBody defines parameters for RegisterAgeing.
type RegisterAgeingJSONBody = PostAgeing

// RegisterBottlingJSONBody defines parameters for RegisterBottling.
type RegisterBottlingJSONBody = PostBottling

// RegisterFermentationJSONBody defines parameters for RegisterFermentation.
type RegisterFermentationJSONBody = PostFermentation

// RegisterMacerationJSONBody defines parameters for RegisterMaceration.
type RegisterMacerationJSONBody = PostMaceration

// RegisterReceptionJSONBody defines parameters for RegisterReception.
type RegisterReceptionJSONBody = PostReception

// RegisterAgeingJSONRequestBody defines body for RegisterAgeing for application/json ContentType.
type RegisterAgeingJSONRequestBody = RegisterAgeingJSONBody

// RegisterBottlingJSONRequestBody defines body for RegisterBottling for application/json ContentType.
type RegisterBottlingJSONRequestBody = RegisterBottlingJSONBody

// RegisterFermentationJSONRequestBody defines body for RegisterFermentation for application/json ContentType.
type RegisterFermentationJSONRequestBody = RegisterFermentationJSONBody

// RegisterMacerationJSONRequestBody defines body for RegisterMaceration for application/json ContentType.
type RegisterMacerationJSONRequestBody = RegisterMacerationJSONBody

// RegisterReceptionJSONRequestBody defines body for RegisterReception for application/json ContentType.
type RegisterReceptionJSONRequestBody = RegisterReceptionJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /ageings)
	GetAgeings(c *gin.Context)

	// (POST /ageings)
	RegisterAgeing(c *gin.Context)

	// (GET /ageings/{ageingUUID})
	GetAgeing(c *gin.Context, ageingUUID openapi_types.UUID)

	// (GET /bottlings)
	GetBottlings(c *gin.Context)

	// (POST /bottlings)
	RegisterBottling(c *gin.Context)

	// (GET /bottlings/{bottlingUUID})
	GetBottling(c *gin.Context, bottlingUUID openapi_types.UUID)

	// (GET /fermentations)
	GetFermentations(c *gin.Context)

	// (POST /fermentations)
	RegisterFermentation(c *gin.Context)

	// (GET /fermentations/{fermentationUUID})
	GetFermentation(c *gin.Context, fermentationUUID openapi_types.UUID)

	// (GET /macerations)
	GetMacerations(c *gin.Context)

	// (POST /macerations)
	RegisterMaceration(c *gin.Context)

	// (GET /macerations/{macerationUUID})
	GetMaceration(c *gin.Context, macerationUUID openapi_types.UUID)

	// (GET /receptions)
	GetReceptions(c *gin.Context)

	// (POST /receptions)
	RegisterReception(c *gin.Context)

	// (GET /receptions/{receptionUUID})
	GetReception(c *gin.Context, receptionUUID openapi_types.UUID)

	// (GET /status)
	Status(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// GetAgeings operation middleware
func (siw *ServerInterfaceWrapper) GetAgeings(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAgeings(c)
}

// RegisterAgeing operation middleware
func (siw *ServerInterfaceWrapper) RegisterAgeing(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.RegisterAgeing(c)
}

// GetAgeing operation middleware
func (siw *ServerInterfaceWrapper) GetAgeing(c *gin.Context) {

	var err error

	// ------------- Path parameter "ageingUUID" -------------
	var ageingUUID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "ageingUUID", c.Param("ageingUUID"), &ageingUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter ageingUUID: %s", err)})
		return
	}

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAgeing(c, ageingUUID)
}

// GetBottlings operation middleware
func (siw *ServerInterfaceWrapper) GetBottlings(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetBottlings(c)
}

// RegisterBottling operation middleware
func (siw *ServerInterfaceWrapper) RegisterBottling(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.RegisterBottling(c)
}

// GetBottling operation middleware
func (siw *ServerInterfaceWrapper) GetBottling(c *gin.Context) {

	var err error

	// ------------- Path parameter "bottlingUUID" -------------
	var bottlingUUID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "bottlingUUID", c.Param("bottlingUUID"), &bottlingUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter bottlingUUID: %s", err)})
		return
	}

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetBottling(c, bottlingUUID)
}

// GetFermentations operation middleware
func (siw *ServerInterfaceWrapper) GetFermentations(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetFermentations(c)
}

// RegisterFermentation operation middleware
func (siw *ServerInterfaceWrapper) RegisterFermentation(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.RegisterFermentation(c)
}

// GetFermentation operation middleware
func (siw *ServerInterfaceWrapper) GetFermentation(c *gin.Context) {

	var err error

	// ------------- Path parameter "fermentationUUID" -------------
	var fermentationUUID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "fermentationUUID", c.Param("fermentationUUID"), &fermentationUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter fermentationUUID: %s", err)})
		return
	}

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetFermentation(c, fermentationUUID)
}

// GetMacerations operation middleware
func (siw *ServerInterfaceWrapper) GetMacerations(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetMacerations(c)
}

// RegisterMaceration operation middleware
func (siw *ServerInterfaceWrapper) RegisterMaceration(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.RegisterMaceration(c)
}

// GetMaceration operation middleware
func (siw *ServerInterfaceWrapper) GetMaceration(c *gin.Context) {

	var err error

	// ------------- Path parameter "macerationUUID" -------------
	var macerationUUID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "macerationUUID", c.Param("macerationUUID"), &macerationUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter macerationUUID: %s", err)})
		return
	}

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetMaceration(c, macerationUUID)
}

// GetReceptions operation middleware
func (siw *ServerInterfaceWrapper) GetReceptions(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetReceptions(c)
}

// RegisterReception operation middleware
func (siw *ServerInterfaceWrapper) RegisterReception(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.RegisterReception(c)
}

// GetReception operation middleware
func (siw *ServerInterfaceWrapper) GetReception(c *gin.Context) {

	var err error

	// ------------- Path parameter "receptionUUID" -------------
	var receptionUUID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "receptionUUID", c.Param("receptionUUID"), &receptionUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter receptionUUID: %s", err)})
		return
	}

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetReception(c, receptionUUID)
}

// Status operation middleware
func (siw *ServerInterfaceWrapper) Status(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.Status(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/ageings", wrapper.GetAgeings)

	router.POST(options.BaseURL+"/ageings", wrapper.RegisterAgeing)

	router.GET(options.BaseURL+"/ageings/:ageingUUID", wrapper.GetAgeing)

	router.GET(options.BaseURL+"/bottlings", wrapper.GetBottlings)

	router.POST(options.BaseURL+"/bottlings", wrapper.RegisterBottling)

	router.GET(options.BaseURL+"/bottlings/:bottlingUUID", wrapper.GetBottling)

	router.GET(options.BaseURL+"/fermentations", wrapper.GetFermentations)

	router.POST(options.BaseURL+"/fermentations", wrapper.RegisterFermentation)

	router.GET(options.BaseURL+"/fermentations/:fermentationUUID", wrapper.GetFermentation)

	router.GET(options.BaseURL+"/macerations", wrapper.GetMacerations)

	router.POST(options.BaseURL+"/macerations", wrapper.RegisterMaceration)

	router.GET(options.BaseURL+"/macerations/:macerationUUID", wrapper.GetMaceration)

	router.GET(options.BaseURL+"/receptions", wrapper.GetReceptions)

	router.POST(options.BaseURL+"/receptions", wrapper.RegisterReception)

	router.GET(options.BaseURL+"/receptions/:receptionUUID", wrapper.GetReception)

	router.GET(options.BaseURL+"/status", wrapper.Status)

	return router
}
