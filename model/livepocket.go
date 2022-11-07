package model

type LivepocketApplicationData struct {
	GroupName                string                 `json:"group_name"`
	GroupID                  string                 `json:"group_id"`
	GroupPublishType         string                 `json:"group_publish_type"`
	GroupStarttime           string                 `json:"group_starttime"`
	GroupEndtimeType         string                 `json:"group_endtime_type"`
	GroupEndtime             string                 `json:"group_endtime"`
	Remarks                  string                 `json:"remarks"`
	OfficialFacebookURL      interface{}            `json:"official_facebook_url"`
	OfficialTwitterURL       interface{}            `json:"official_twitter_url"`
	GroupOrderLimited        string                 `json:"group_order_limited"`
	GroupEventOrderLimited   string                 `json:"group_event_order_limited"`
	GroupSmsLimited          string                 `json:"group_sms_limited"`
	LotteryAnnounce          string                 `json:"lottery_announce"`
	Publishtime              string                 `json:"publishtime"`
	FanClubName              interface{}            `json:"fan_club_name"`
	FanClubEntryDate         interface{}            `json:"fan_club_entry_date"`
	Cname                    interface{}            `json:"cname"`
	Referer                  interface{}            `json:"referer"`
	AccessLimited            bool                   `json:"access_limited"`
	IsGotoTarget             bool                   `json:"is_goto_target"`
	LimitedAccessAppName     interface{}            `json:"limited_access_app_name"`
	LimitedAccessAppType     interface{}            `json:"limited_access_app_type"`
	IsPurchaseLimitedByFcApp bool                   `json:"is_purchase_limited_by_fc_app"`
	TicketsInfo              []LivepocketTicketData `json:"tickets_info"`
}

type LivepocketTicketData struct {
	ID                         int         `json:"id"`
	Type                       string      `json:"type"`
	Name                       string      `json:"name"`
	Price                      int         `json:"price"`
	Starttime                  string      `json:"starttime"`
	Endtime                    string      `json:"endtime"`
	Publishtime                string      `json:"publishtime"`
	SoldOut                    bool        `json:"sold_out"`
	TicketStock                int         `json:"ticket_stock"`
	DisplayRemain              bool        `json:"display_remain"`
	LimitMin                   int         `json:"limit_min"`
	LimitMax                   int         `json:"limit_max"`
	PurchaseLimited            bool        `json:"purchase_limited"`
	PublishType                string      `json:"publish_type"`
	Referer                    interface{} `json:"referer"`
	Hash                       interface{} `json:"hash"`
	Free                       bool        `json:"free"`
	IsDisabled                 bool        `json:"is_disabled"`
	Pattern                    bool        `json:"pattern"`
	Detail                     string      `json:"detail"`
	EntryLimited               string      `json:"entry_limited"`
	OrderLimited               string      `json:"order_limited"`
	SalesStatus                int         `json:"sales_status"`
	EventTicketSendForbiddance string      `json:"event_ticket_send_forbiddance"`
}
