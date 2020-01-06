package goshopify

import (
	"fmt"
	"time"
)

// InventoryLevelService is an interface for interfacing with the inventory level endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/locations
type InventoryLevelService interface {
	ListPage(int64, interface{}) ([]InventoryLevel, *ResponseMeta, error)
}

type InventoryLevel struct {
	// The Inventory Item ID.
	InventoryItemID int64 `json:"inventory_item_id"`

	// The Location ID.
	LocationID int64 `json:"location_id"`

	// Inventory Available.
	Available int32 `json:"available"`

	// The date and time (ISO 8601 format) when the location was last updated.
	UpdatedAt time.Time `json:"updated_at"`

	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
}

// InventoryLevelServiceOp handles communication with the inventory level related methods of the
// Shopify API.
type InventoryLevelServiceOp struct {
	client *Client
}

// InventoryLevelsResource represents the result from the locations/X/inventory_levels.json endpoint
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

// ListPage Location Inventory Levels
func (s *InventoryLevelServiceOp) ListPage(location int64, options interface{}) ([]InventoryLevel, *ResponseMeta, error) {
	path := fmt.Sprintf("%s/%s/%d/inventory_levels.json", globalApiPathPrefix, locationsBasePath, location)
	resource := new(InventoryLevelsResource)
	meta, err := s.client.GetPaging(path, resource, options)
	return resource.InventoryLevels, meta, err
}
