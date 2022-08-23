// Code generated by goa v3.8.3, DO NOT EDIT.
//
// resource service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	resourceviews "github.com/tektoncd/hub/api/gen/resource/views"
	goa "goa.design/goa/v3/pkg"
)

// The resource service provides details about all kind of resources
type Service interface {
	// Find resources by a combination of name, kind, catalog, categories,
	// platforms and tags
	Query(context.Context, *QueryPayload) (res *QueryResult, err error)
	// List all resources sorted by rating and name
	List(context.Context) (res *Resources, err error)
	// Find all versions of a resource by its id
	VersionsByID(context.Context, *VersionsByIDPayload) (res *VersionsByIDResult, err error)
	// Find resource using name of catalog & name, kind and version of resource
	ByCatalogKindNameVersion(context.Context, *ByCatalogKindNameVersionPayload) (res *ByCatalogKindNameVersionResult, err error)
	// Find a resource using its version's id
	ByVersionID(context.Context, *ByVersionIDPayload) (res *ByVersionIDResult, err error)
	// Find resources using name of catalog, resource name and kind of resource
	ByCatalogKindName(context.Context, *ByCatalogKindNamePayload) (res *ByCatalogKindNameResult, err error)
	// Find a resource using it's id
	ByID(context.Context, *ByIDPayload) (res *ByIDResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "resource"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [7]string{"Query", "List", "VersionsByID", "ByCatalogKindNameVersion", "ByVersionId", "ByCatalogKindName", "ById"}

// ByCatalogKindNamePayload is the payload type of the resource service
// ByCatalogKindName method.
type ByCatalogKindNamePayload struct {
	// name of catalog
	Catalog string
	// kind of resource
	Kind string
	// Name of resource
	Name string
	// To find resource compatible with a Tekton pipelines version, use this param
	Pipelinesversion *string
}

// ByCatalogKindNameResult is the result type of the resource service
// ByCatalogKindName method.
type ByCatalogKindNameResult struct {
	// Redirect URL
	Location string
}

// ByCatalogKindNameVersionPayload is the payload type of the resource service
// ByCatalogKindNameVersion method.
type ByCatalogKindNameVersionPayload struct {
	// name of catalog
	Catalog string
	// kind of resource
	Kind string
	// name of resource
	Name string
	// version of resource
	Version string
}

// ByCatalogKindNameVersionResult is the result type of the resource service
// ByCatalogKindNameVersion method.
type ByCatalogKindNameVersionResult struct {
	// Redirect URL
	Location string
}

// ByIDPayload is the payload type of the resource service ById method.
type ByIDPayload struct {
	// ID of a resource
	ID uint
}

// ByIDResult is the result type of the resource service ById method.
type ByIDResult struct {
	// Redirect URL
	Location string
}

// ByVersionIDPayload is the payload type of the resource service ByVersionId
// method.
type ByVersionIDPayload struct {
	// Version ID of a resource's version
	VersionID uint
}

// ByVersionIDResult is the result type of the resource service ByVersionId
// method.
type ByVersionIDResult struct {
	// Redirect URL
	Location string
}

type Catalog struct {
	// ID is the unique id of the catalog
	ID uint
	// Name of catalog
	Name string
	// Type of catalog
	Type string
	// URL of catalog
	URL string
	// Provider of catalog
	Provider string
}

type Category struct {
	// ID is the unique id of the category
	ID uint
	// Name of category
	Name string
}

type Platform struct {
	// ID is the unique id of platform
	ID uint
	// Name of platform
	Name string
}

// QueryPayload is the payload type of the resource service Query method.
type QueryPayload struct {
	// Name of resource
	Name string
	// Catalogs of resource to filter by
	Catalogs []string
	// Kinds of resource to filter by
	Kinds []string
	// Category associated with a resource to filter by
	Categories []string
	// Tags associated with a resource to filter by
	Tags []string
	// Platforms associated with a resource to filter by
	Platforms []string
	// Maximum number of resources to be returned
	Limit uint
	// Strategy used to find matching resources
	Match string
}

// QueryResult is the result type of the resource service Query method.
type QueryResult struct {
	// Redirect URL
	Location string
}

// The resource type describes resource information.
type ResourceData struct {
	// ID is the unique id of the resource
	ID uint
	// Name of resource
	Name string
	// Type of catalog to which resource belongs
	Catalog *Catalog
	// Categories related to resource
	Categories []*Category
	// Kind of resource
	Kind string
	// Url path of the resource in hub
	HubURLPath string
	// Latest version of resource
	LatestVersion *ResourceVersionData
	// Tags related to resource
	Tags []*Tag
	// Platforms related to resource
	Platforms []*Platform
	// Rating of resource
	Rating float64
	// List of all versions of a resource
	Versions []*ResourceVersionData
}

type ResourceDataCollection []*ResourceData

// The Version result type describes resource's version information.
type ResourceVersionData struct {
	// ID is the unique id of resource's version
	ID uint
	// Version of resource
	Version string
	// Display name of version
	DisplayName string
	// Deprecation status of a version
	Deprecated *bool
	// Description of version
	Description string
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion string
	// Raw URL of resource's yaml file of the version
	RawURL string
	// Web URL of resource's yaml file of the version
	WebURL string
	// Timestamp when version was last updated
	UpdatedAt string
	// Platforms related to resource version
	Platforms []*Platform
	// Url path of the resource in hub
	HubURLPath string
	// Resource to which the version belongs
	Resource *ResourceData
}

// Resources is the result type of the resource service List method.
type Resources struct {
	Data ResourceDataCollection
}

type Tag struct {
	// ID is the unique id of tag
	ID uint
	// Name of tag
	Name string
}

// VersionsByIDPayload is the payload type of the resource service VersionsByID
// method.
type VersionsByIDPayload struct {
	// ID of a resource
	ID uint
}

// VersionsByIDResult is the result type of the resource service VersionsByID
// method.
type VersionsByIDResult struct {
	// Redirect URL
	Location string
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "internal-error", false, false, false)
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not-found", false, false, false)
}

// NewResources initializes result type Resources from viewed result type
// Resources.
func NewResources(vres *resourceviews.Resources) *Resources {
	return newResources(vres.Projected)
}

// NewViewedResources initializes viewed result type Resources from result type
// Resources using the given view.
func NewViewedResources(res *Resources, view string) *resourceviews.Resources {
	p := newResourcesView(res)
	return &resourceviews.Resources{Projected: p, View: "default"}
}

// newResources converts projected type Resources to service type Resources.
func newResources(vres *resourceviews.ResourcesView) *Resources {
	res := &Resources{}
	if vres.Data != nil {
		res.Data = newResourceDataCollectionWithoutVersion(vres.Data)
	}
	return res
}

// newResourcesView projects result type Resources to projected type
// ResourcesView using the "default" view.
func newResourcesView(res *Resources) *resourceviews.ResourcesView {
	vres := &resourceviews.ResourcesView{}
	if res.Data != nil {
		vres.Data = newResourceDataCollectionViewWithoutVersion(res.Data)
	}
	return vres
}

// newResourceDataCollectionInfo converts projected type ResourceDataCollection
// to service type ResourceDataCollection.
func newResourceDataCollectionInfo(vres resourceviews.ResourceDataCollectionView) ResourceDataCollection {
	res := make(ResourceDataCollection, len(vres))
	for i, n := range vres {
		res[i] = newResourceDataInfo(n)
	}
	return res
}

// newResourceDataCollectionWithoutVersion converts projected type
// ResourceDataCollection to service type ResourceDataCollection.
func newResourceDataCollectionWithoutVersion(vres resourceviews.ResourceDataCollectionView) ResourceDataCollection {
	res := make(ResourceDataCollection, len(vres))
	for i, n := range vres {
		res[i] = newResourceDataWithoutVersion(n)
	}
	return res
}

// newResourceDataCollection converts projected type ResourceDataCollection to
// service type ResourceDataCollection.
func newResourceDataCollection(vres resourceviews.ResourceDataCollectionView) ResourceDataCollection {
	res := make(ResourceDataCollection, len(vres))
	for i, n := range vres {
		res[i] = newResourceData(n)
	}
	return res
}

// newResourceDataCollectionViewInfo projects result type
// ResourceDataCollection to projected type ResourceDataCollectionView using
// the "info" view.
func newResourceDataCollectionViewInfo(res ResourceDataCollection) resourceviews.ResourceDataCollectionView {
	vres := make(resourceviews.ResourceDataCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceDataViewInfo(n)
	}
	return vres
}

// newResourceDataCollectionViewWithoutVersion projects result type
// ResourceDataCollection to projected type ResourceDataCollectionView using
// the "withoutVersion" view.
func newResourceDataCollectionViewWithoutVersion(res ResourceDataCollection) resourceviews.ResourceDataCollectionView {
	vres := make(resourceviews.ResourceDataCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceDataViewWithoutVersion(n)
	}
	return vres
}

// newResourceDataCollectionView projects result type ResourceDataCollection to
// projected type ResourceDataCollectionView using the "default" view.
func newResourceDataCollectionView(res ResourceDataCollection) resourceviews.ResourceDataCollectionView {
	vres := make(resourceviews.ResourceDataCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceDataView(n)
	}
	return vres
}

// newResourceDataInfo converts projected type ResourceData to service type
// ResourceData.
func newResourceDataInfo(vres *resourceviews.ResourceDataView) *ResourceData {
	res := &ResourceData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.HubURLPath != nil {
		res.HubURLPath = *vres.HubURLPath
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Categories != nil {
		res.Categories = make([]*Category, len(vres.Categories))
		for i, val := range vres.Categories {
			res.Categories[i] = transformResourceviewsCategoryViewToCategory(val)
		}
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if vres.Platforms != nil {
		res.Platforms = make([]*Platform, len(vres.Platforms))
		for i, val := range vres.Platforms {
			res.Platforms[i] = transformResourceviewsPlatformViewToPlatform(val)
		}
	}
	if vres.Catalog != nil {
		res.Catalog = newCatalogMin(vres.Catalog)
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = newResourceVersionData(vres.LatestVersion)
	}
	return res
}

// newResourceDataWithoutVersion converts projected type ResourceData to
// service type ResourceData.
func newResourceDataWithoutVersion(vres *resourceviews.ResourceDataView) *ResourceData {
	res := &ResourceData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.HubURLPath != nil {
		res.HubURLPath = *vres.HubURLPath
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Categories != nil {
		res.Categories = make([]*Category, len(vres.Categories))
		for i, val := range vres.Categories {
			res.Categories[i] = transformResourceviewsCategoryViewToCategory(val)
		}
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if vres.Platforms != nil {
		res.Platforms = make([]*Platform, len(vres.Platforms))
		for i, val := range vres.Platforms {
			res.Platforms[i] = transformResourceviewsPlatformViewToPlatform(val)
		}
	}
	if vres.Catalog != nil {
		res.Catalog = newCatalogMin(vres.Catalog)
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = newResourceVersionDataWithoutResource(vres.LatestVersion)
	}
	return res
}

// newResourceData converts projected type ResourceData to service type
// ResourceData.
func newResourceData(vres *resourceviews.ResourceDataView) *ResourceData {
	res := &ResourceData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Kind != nil {
		res.Kind = *vres.Kind
	}
	if vres.HubURLPath != nil {
		res.HubURLPath = *vres.HubURLPath
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Categories != nil {
		res.Categories = make([]*Category, len(vres.Categories))
		for i, val := range vres.Categories {
			res.Categories[i] = transformResourceviewsCategoryViewToCategory(val)
		}
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if vres.Platforms != nil {
		res.Platforms = make([]*Platform, len(vres.Platforms))
		for i, val := range vres.Platforms {
			res.Platforms[i] = transformResourceviewsPlatformViewToPlatform(val)
		}
	}
	if vres.Versions != nil {
		res.Versions = make([]*ResourceVersionData, len(vres.Versions))
		for i, val := range vres.Versions {
			res.Versions[i] = transformResourceviewsResourceVersionDataViewToResourceVersionData(val)
		}
	}
	if vres.Catalog != nil {
		res.Catalog = newCatalogMin(vres.Catalog)
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = newResourceVersionDataWithoutResource(vres.LatestVersion)
	}
	return res
}

// newResourceDataViewInfo projects result type ResourceData to projected type
// ResourceDataView using the "info" view.
func newResourceDataViewInfo(res *ResourceData) *resourceviews.ResourceDataView {
	vres := &resourceviews.ResourceDataView{
		ID:         &res.ID,
		Name:       &res.Name,
		Kind:       &res.Kind,
		HubURLPath: &res.HubURLPath,
		Rating:     &res.Rating,
	}
	if res.Categories != nil {
		vres.Categories = make([]*resourceviews.CategoryView, len(res.Categories))
		for i, val := range res.Categories {
			vres.Categories[i] = transformCategoryToResourceviewsCategoryView(val)
		}
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if res.Platforms != nil {
		vres.Platforms = make([]*resourceviews.PlatformView, len(res.Platforms))
		for i, val := range res.Platforms {
			vres.Platforms[i] = transformPlatformToResourceviewsPlatformView(val)
		}
	}
	if res.Catalog != nil {
		vres.Catalog = newCatalogViewMin(res.Catalog)
	}
	return vres
}

// newResourceDataViewWithoutVersion projects result type ResourceData to
// projected type ResourceDataView using the "withoutVersion" view.
func newResourceDataViewWithoutVersion(res *ResourceData) *resourceviews.ResourceDataView {
	vres := &resourceviews.ResourceDataView{
		ID:         &res.ID,
		Name:       &res.Name,
		Kind:       &res.Kind,
		HubURLPath: &res.HubURLPath,
		Rating:     &res.Rating,
	}
	if res.Categories != nil {
		vres.Categories = make([]*resourceviews.CategoryView, len(res.Categories))
		for i, val := range res.Categories {
			vres.Categories[i] = transformCategoryToResourceviewsCategoryView(val)
		}
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if res.Platforms != nil {
		vres.Platforms = make([]*resourceviews.PlatformView, len(res.Platforms))
		for i, val := range res.Platforms {
			vres.Platforms[i] = transformPlatformToResourceviewsPlatformView(val)
		}
	}
	if res.Catalog != nil {
		vres.Catalog = newCatalogViewMin(res.Catalog)
	}
	if res.LatestVersion != nil {
		vres.LatestVersion = newResourceVersionDataViewWithoutResource(res.LatestVersion)
	}
	return vres
}

// newResourceDataView projects result type ResourceData to projected type
// ResourceDataView using the "default" view.
func newResourceDataView(res *ResourceData) *resourceviews.ResourceDataView {
	vres := &resourceviews.ResourceDataView{
		ID:         &res.ID,
		Name:       &res.Name,
		Kind:       &res.Kind,
		HubURLPath: &res.HubURLPath,
		Rating:     &res.Rating,
	}
	if res.Categories != nil {
		vres.Categories = make([]*resourceviews.CategoryView, len(res.Categories))
		for i, val := range res.Categories {
			vres.Categories[i] = transformCategoryToResourceviewsCategoryView(val)
		}
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if res.Platforms != nil {
		vres.Platforms = make([]*resourceviews.PlatformView, len(res.Platforms))
		for i, val := range res.Platforms {
			vres.Platforms[i] = transformPlatformToResourceviewsPlatformView(val)
		}
	}
	if res.Versions != nil {
		vres.Versions = make([]*resourceviews.ResourceVersionDataView, len(res.Versions))
		for i, val := range res.Versions {
			vres.Versions[i] = transformResourceVersionDataToResourceviewsResourceVersionDataView(val)
		}
	}
	if res.Catalog != nil {
		vres.Catalog = newCatalogViewMin(res.Catalog)
	}
	if res.LatestVersion != nil {
		vres.LatestVersion = newResourceVersionDataViewWithoutResource(res.LatestVersion)
	}
	return vres
}

// newCatalogMin converts projected type Catalog to service type Catalog.
func newCatalogMin(vres *resourceviews.CatalogView) *Catalog {
	res := &Catalog{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	return res
}

// newCatalog converts projected type Catalog to service type Catalog.
func newCatalog(vres *resourceviews.CatalogView) *Catalog {
	res := &Catalog{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.URL != nil {
		res.URL = *vres.URL
	}
	if vres.Provider != nil {
		res.Provider = *vres.Provider
	}
	return res
}

// newCatalogViewMin projects result type Catalog to projected type CatalogView
// using the "min" view.
func newCatalogViewMin(res *Catalog) *resourceviews.CatalogView {
	vres := &resourceviews.CatalogView{
		ID:   &res.ID,
		Name: &res.Name,
		Type: &res.Type,
	}
	return vres
}

// newCatalogView projects result type Catalog to projected type CatalogView
// using the "default" view.
func newCatalogView(res *Catalog) *resourceviews.CatalogView {
	vres := &resourceviews.CatalogView{
		ID:       &res.ID,
		Name:     &res.Name,
		Type:     &res.Type,
		URL:      &res.URL,
		Provider: &res.Provider,
	}
	return vres
}

// newResourceVersionDataTiny converts projected type ResourceVersionData to
// service type ResourceVersionData.
func newResourceVersionDataTiny(vres *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	res := &ResourceVersionData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.Resource != nil {
		res.Resource = newResourceData(vres.Resource)
	}
	return res
}

// newResourceVersionDataMin converts projected type ResourceVersionData to
// service type ResourceVersionData.
func newResourceVersionDataMin(vres *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	res := &ResourceVersionData{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.RawURL != nil {
		res.RawURL = *vres.RawURL
	}
	if vres.WebURL != nil {
		res.WebURL = *vres.WebURL
	}
	if vres.HubURLPath != nil {
		res.HubURLPath = *vres.HubURLPath
	}
	if vres.Platforms != nil {
		res.Platforms = make([]*Platform, len(vres.Platforms))
		for i, val := range vres.Platforms {
			res.Platforms[i] = transformResourceviewsPlatformViewToPlatform(val)
		}
	}
	if vres.Resource != nil {
		res.Resource = newResourceData(vres.Resource)
	}
	return res
}

// newResourceVersionDataWithoutResource converts projected type
// ResourceVersionData to service type ResourceVersionData.
func newResourceVersionDataWithoutResource(vres *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	res := &ResourceVersionData{
		Deprecated: vres.Deprecated,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.DisplayName != nil {
		res.DisplayName = *vres.DisplayName
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.MinPipelinesVersion != nil {
		res.MinPipelinesVersion = *vres.MinPipelinesVersion
	}
	if vres.RawURL != nil {
		res.RawURL = *vres.RawURL
	}
	if vres.WebURL != nil {
		res.WebURL = *vres.WebURL
	}
	if vres.HubURLPath != nil {
		res.HubURLPath = *vres.HubURLPath
	}
	if vres.UpdatedAt != nil {
		res.UpdatedAt = *vres.UpdatedAt
	}
	if vres.Platforms != nil {
		res.Platforms = make([]*Platform, len(vres.Platforms))
		for i, val := range vres.Platforms {
			res.Platforms[i] = transformResourceviewsPlatformViewToPlatform(val)
		}
	}
	if vres.Resource != nil {
		res.Resource = newResourceData(vres.Resource)
	}
	return res
}

// newResourceVersionData converts projected type ResourceVersionData to
// service type ResourceVersionData.
func newResourceVersionData(vres *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	res := &ResourceVersionData{
		Deprecated: vres.Deprecated,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.DisplayName != nil {
		res.DisplayName = *vres.DisplayName
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.MinPipelinesVersion != nil {
		res.MinPipelinesVersion = *vres.MinPipelinesVersion
	}
	if vres.RawURL != nil {
		res.RawURL = *vres.RawURL
	}
	if vres.WebURL != nil {
		res.WebURL = *vres.WebURL
	}
	if vres.HubURLPath != nil {
		res.HubURLPath = *vres.HubURLPath
	}
	if vres.UpdatedAt != nil {
		res.UpdatedAt = *vres.UpdatedAt
	}
	if vres.Platforms != nil {
		res.Platforms = make([]*Platform, len(vres.Platforms))
		for i, val := range vres.Platforms {
			res.Platforms[i] = transformResourceviewsPlatformViewToPlatform(val)
		}
	}
	if vres.Resource != nil {
		res.Resource = newResourceDataInfo(vres.Resource)
	}
	return res
}

// newResourceVersionDataViewTiny projects result type ResourceVersionData to
// projected type ResourceVersionDataView using the "tiny" view.
func newResourceVersionDataViewTiny(res *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	vres := &resourceviews.ResourceVersionDataView{
		ID:      &res.ID,
		Version: &res.Version,
	}
	return vres
}

// newResourceVersionDataViewMin projects result type ResourceVersionData to
// projected type ResourceVersionDataView using the "min" view.
func newResourceVersionDataViewMin(res *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	vres := &resourceviews.ResourceVersionDataView{
		ID:         &res.ID,
		Version:    &res.Version,
		RawURL:     &res.RawURL,
		WebURL:     &res.WebURL,
		HubURLPath: &res.HubURLPath,
	}
	if res.Platforms != nil {
		vres.Platforms = make([]*resourceviews.PlatformView, len(res.Platforms))
		for i, val := range res.Platforms {
			vres.Platforms[i] = transformPlatformToResourceviewsPlatformView(val)
		}
	}
	return vres
}

// newResourceVersionDataViewWithoutResource projects result type
// ResourceVersionData to projected type ResourceVersionDataView using the
// "withoutResource" view.
func newResourceVersionDataViewWithoutResource(res *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	vres := &resourceviews.ResourceVersionDataView{
		ID:                  &res.ID,
		Version:             &res.Version,
		DisplayName:         &res.DisplayName,
		Deprecated:          res.Deprecated,
		Description:         &res.Description,
		MinPipelinesVersion: &res.MinPipelinesVersion,
		RawURL:              &res.RawURL,
		WebURL:              &res.WebURL,
		UpdatedAt:           &res.UpdatedAt,
		HubURLPath:          &res.HubURLPath,
	}
	if res.Platforms != nil {
		vres.Platforms = make([]*resourceviews.PlatformView, len(res.Platforms))
		for i, val := range res.Platforms {
			vres.Platforms[i] = transformPlatformToResourceviewsPlatformView(val)
		}
	}
	return vres
}

// newResourceVersionDataView projects result type ResourceVersionData to
// projected type ResourceVersionDataView using the "default" view.
func newResourceVersionDataView(res *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	vres := &resourceviews.ResourceVersionDataView{
		ID:                  &res.ID,
		Version:             &res.Version,
		DisplayName:         &res.DisplayName,
		Deprecated:          res.Deprecated,
		Description:         &res.Description,
		MinPipelinesVersion: &res.MinPipelinesVersion,
		RawURL:              &res.RawURL,
		WebURL:              &res.WebURL,
		UpdatedAt:           &res.UpdatedAt,
		HubURLPath:          &res.HubURLPath,
	}
	if res.Platforms != nil {
		vres.Platforms = make([]*resourceviews.PlatformView, len(res.Platforms))
		for i, val := range res.Platforms {
			vres.Platforms[i] = transformPlatformToResourceviewsPlatformView(val)
		}
	}
	if res.Resource != nil {
		vres.Resource = newResourceDataViewInfo(res.Resource)
	}
	return vres
}

// transformResourceviewsCategoryViewToCategory builds a value of type
// *Category from a value of type *resourceviews.CategoryView.
func transformResourceviewsCategoryViewToCategory(v *resourceviews.CategoryView) *Category {
	if v == nil {
		return nil
	}
	res := &Category{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformResourceviewsTagViewToTag builds a value of type *Tag from a value
// of type *resourceviews.TagView.
func transformResourceviewsTagViewToTag(v *resourceviews.TagView) *Tag {
	if v == nil {
		return nil
	}
	res := &Tag{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformResourceviewsPlatformViewToPlatform builds a value of type
// *Platform from a value of type *resourceviews.PlatformView.
func transformResourceviewsPlatformViewToPlatform(v *resourceviews.PlatformView) *Platform {
	if v == nil {
		return nil
	}
	res := &Platform{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformResourceviewsResourceVersionDataViewToResourceVersionData builds a
// value of type *ResourceVersionData from a value of type
// *resourceviews.ResourceVersionDataView.
func transformResourceviewsResourceVersionDataViewToResourceVersionData(v *resourceviews.ResourceVersionDataView) *ResourceVersionData {
	if v == nil {
		return nil
	}
	res := &ResourceVersionData{
		ID:                  *v.ID,
		Version:             *v.Version,
		DisplayName:         *v.DisplayName,
		Deprecated:          v.Deprecated,
		Description:         *v.Description,
		MinPipelinesVersion: *v.MinPipelinesVersion,
		RawURL:              *v.RawURL,
		WebURL:              *v.WebURL,
		UpdatedAt:           *v.UpdatedAt,
		HubURLPath:          *v.HubURLPath,
	}
	if v.Platforms != nil {
		res.Platforms = make([]*Platform, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = transformResourceviewsPlatformViewToPlatform(val)
		}
	}
	if v.Resource != nil {
		res.Resource = transformResourceviewsResourceDataViewToResourceData(v.Resource)
	}

	return res
}

// transformResourceviewsResourceDataViewToResourceData builds a value of type
// *ResourceData from a value of type *resourceviews.ResourceDataView.
func transformResourceviewsResourceDataViewToResourceData(v *resourceviews.ResourceDataView) *ResourceData {
	res := &ResourceData{}
	if v.ID != nil {
		res.ID = *v.ID
	}
	if v.Name != nil {
		res.Name = *v.Name
	}
	if v.Kind != nil {
		res.Kind = *v.Kind
	}
	if v.HubURLPath != nil {
		res.HubURLPath = *v.HubURLPath
	}
	if v.Rating != nil {
		res.Rating = *v.Rating
	}
	if v.Categories != nil {
		res.Categories = make([]*Category, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = transformResourceviewsCategoryViewToCategory(val)
		}
	}
	if v.Tags != nil {
		res.Tags = make([]*Tag, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	if v.Platforms != nil {
		res.Platforms = make([]*Platform, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = transformResourceviewsPlatformViewToPlatform(val)
		}
	}
	if v.Versions != nil {
		res.Versions = make([]*ResourceVersionData, len(v.Versions))
		for i, val := range v.Versions {
			res.Versions[i] = transformResourceviewsResourceVersionDataViewToResourceVersionData(val)
		}
	}

	return res
}

// transformCategoryToResourceviewsCategoryView builds a value of type
// *resourceviews.CategoryView from a value of type *Category.
func transformCategoryToResourceviewsCategoryView(v *Category) *resourceviews.CategoryView {
	res := &resourceviews.CategoryView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}

// transformTagToResourceviewsTagView builds a value of type
// *resourceviews.TagView from a value of type *Tag.
func transformTagToResourceviewsTagView(v *Tag) *resourceviews.TagView {
	res := &resourceviews.TagView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}

// transformPlatformToResourceviewsPlatformView builds a value of type
// *resourceviews.PlatformView from a value of type *Platform.
func transformPlatformToResourceviewsPlatformView(v *Platform) *resourceviews.PlatformView {
	res := &resourceviews.PlatformView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}

// transformResourceVersionDataToResourceviewsResourceVersionDataView builds a
// value of type *resourceviews.ResourceVersionDataView from a value of type
// *ResourceVersionData.
func transformResourceVersionDataToResourceviewsResourceVersionDataView(v *ResourceVersionData) *resourceviews.ResourceVersionDataView {
	res := &resourceviews.ResourceVersionDataView{
		ID:                  &v.ID,
		Version:             &v.Version,
		DisplayName:         &v.DisplayName,
		Deprecated:          v.Deprecated,
		Description:         &v.Description,
		MinPipelinesVersion: &v.MinPipelinesVersion,
		RawURL:              &v.RawURL,
		WebURL:              &v.WebURL,
		UpdatedAt:           &v.UpdatedAt,
		HubURLPath:          &v.HubURLPath,
	}
	if v.Platforms != nil {
		res.Platforms = make([]*resourceviews.PlatformView, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = transformPlatformToResourceviewsPlatformView(val)
		}
	}
	if v.Resource != nil {
		res.Resource = transformResourceDataToResourceviewsResourceDataView(v.Resource)
	}

	return res
}

// transformResourceDataToResourceviewsResourceDataView builds a value of type
// *resourceviews.ResourceDataView from a value of type *ResourceData.
func transformResourceDataToResourceviewsResourceDataView(v *ResourceData) *resourceviews.ResourceDataView {
	res := &resourceviews.ResourceDataView{
		ID:         &v.ID,
		Name:       &v.Name,
		Kind:       &v.Kind,
		HubURLPath: &v.HubURLPath,
		Rating:     &v.Rating,
	}
	if v.Categories != nil {
		res.Categories = make([]*resourceviews.CategoryView, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = transformCategoryToResourceviewsCategoryView(val)
		}
	}
	if v.Tags != nil {
		res.Tags = make([]*resourceviews.TagView, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	if v.Platforms != nil {
		res.Platforms = make([]*resourceviews.PlatformView, len(v.Platforms))
		for i, val := range v.Platforms {
			res.Platforms[i] = transformPlatformToResourceviewsPlatformView(val)
		}
	}
	if v.Versions != nil {
		res.Versions = make([]*resourceviews.ResourceVersionDataView, len(v.Versions))
		for i, val := range v.Versions {
			res.Versions[i] = transformResourceVersionDataToResourceviewsResourceVersionDataView(val)
		}
	}

	return res
}
