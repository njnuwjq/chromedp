// Package domstorage provides the Chrome Debugging Protocol
// commands, types, and events for the DOMStorage domain.
//
// Query and modify DOM storage.
//
// Generated by the chromedp-gen command.
package domstorage

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
)

// EnableParams enables storage tracking, storage events will now be
// delivered to the client.
type EnableParams struct{}

// Enable enables storage tracking, storage events will now be delivered to
// the client.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes DOMStorage.enable against the provided context and
// target handler.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMStorageEnable, nil, nil)
}

// DisableParams disables storage tracking, prevents storage events from
// being sent to the client.
type DisableParams struct{}

// Disable disables storage tracking, prevents storage events from being sent
// to the client.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes DOMStorage.disable against the provided context and
// target handler.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMStorageDisable, nil, nil)
}

// ClearParams [no description].
type ClearParams struct {
	StorageID *StorageID `json:"storageId"`
}

// Clear [no description].
//
// parameters:
//   storageID
func Clear(storageID *StorageID) *ClearParams {
	return &ClearParams{
		StorageID: storageID,
	}
}

// Do executes DOMStorage.clear against the provided context and
// target handler.
func (p *ClearParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMStorageClear, p, nil)
}

// GetDOMStorageItemsParams [no description].
type GetDOMStorageItemsParams struct {
	StorageID *StorageID `json:"storageId"`
}

// GetDOMStorageItems [no description].
//
// parameters:
//   storageID
func GetDOMStorageItems(storageID *StorageID) *GetDOMStorageItemsParams {
	return &GetDOMStorageItemsParams{
		StorageID: storageID,
	}
}

// GetDOMStorageItemsReturns return values.
type GetDOMStorageItemsReturns struct {
	Entries []Item `json:"entries,omitempty"`
}

// Do executes DOMStorage.getDOMStorageItems against the provided context and
// target handler.
//
// returns:
//   entries
func (p *GetDOMStorageItemsParams) Do(ctxt context.Context, h cdp.Handler) (entries []Item, err error) {
	// execute
	var res GetDOMStorageItemsReturns
	err = h.Execute(ctxt, cdp.CommandDOMStorageGetDOMStorageItems, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Entries, nil
}

// SetDOMStorageItemParams [no description].
type SetDOMStorageItemParams struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`
	Value     string     `json:"value"`
}

// SetDOMStorageItem [no description].
//
// parameters:
//   storageID
//   key
//   value
func SetDOMStorageItem(storageID *StorageID, key string, value string) *SetDOMStorageItemParams {
	return &SetDOMStorageItemParams{
		StorageID: storageID,
		Key:       key,
		Value:     value,
	}
}

// Do executes DOMStorage.setDOMStorageItem against the provided context and
// target handler.
func (p *SetDOMStorageItemParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMStorageSetDOMStorageItem, p, nil)
}

// RemoveDOMStorageItemParams [no description].
type RemoveDOMStorageItemParams struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`
}

// RemoveDOMStorageItem [no description].
//
// parameters:
//   storageID
//   key
func RemoveDOMStorageItem(storageID *StorageID, key string) *RemoveDOMStorageItemParams {
	return &RemoveDOMStorageItemParams{
		StorageID: storageID,
		Key:       key,
	}
}

// Do executes DOMStorage.removeDOMStorageItem against the provided context and
// target handler.
func (p *RemoveDOMStorageItemParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDOMStorageRemoveDOMStorageItem, p, nil)
}
