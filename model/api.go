package model

type APIResponse struct {
	Ok         bool                        `json:"ok"`
	Livepocket []LivepocketApplicationData `json:"livepocketData,omitempty"`
	Litlink    LitlinkData                 `json:"litlinkData,omitempty"`
}
