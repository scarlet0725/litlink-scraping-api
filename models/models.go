package models

type ScrapingRequest struct {
	Url string `json:"url"`
}

type LitlinkProps struct {
	Props struct {
		PageProps struct {
			InitialState struct {
				Account struct {
					Email               string      `json:"email"`
					URL                 string      `json:"url"`
					EmailUpdateResponse interface{} `json:"emailUpdateResponse"`
					IsLoading           bool        `json:"isLoading"`
					EmailConnected      bool        `json:"emailConnected"`
					URLUpdateResponse   interface{} `json:"urlUpdateResponse"`
				} `json:"account"`
				CreatorDetailEdit struct {
					IsEdit                     bool          `json:"isEdit"`
					EditingProfile             interface{}   `json:"editingProfile"`
					EditingSnsIconLinkDetails  []interface{} `json:"editingSnsIconLinkDetails"`
					EditingProfileLinkDetails  []interface{} `json:"editingProfileLinkDetails"`
					EditingCreatorDetailLayout interface{}   `json:"editingCreatorDetailLayout"`
					SelectedBackgroundCategory string        `json:"selectedBackgroundCategory"`
					FontColor                  struct {
						R int `json:"r"`
						G int `json:"g"`
						B int `json:"b"`
						A int `json:"a"`
					} `json:"fontColor"`
					BackgroundColor struct {
						R int `json:"r"`
						G int `json:"g"`
						B int `json:"b"`
						A int `json:"a"`
					} `json:"backgroundColor"`
					BackgroundGradationStartColor struct {
						R int `json:"r"`
						G int `json:"g"`
						B int `json:"b"`
						A int `json:"a"`
					} `json:"backgroundGradationStartColor"`
					BackgroundGradationEndColor struct {
						R int `json:"r"`
						G int `json:"g"`
						B int `json:"b"`
						A int `json:"a"`
					} `json:"backgroundGradationEndColor"`
					BackgroundGradationColorPaletteIndex int  `json:"backgroundGradationColorPaletteIndex"`
					IsActiveURLPastingOnText             bool `json:"isActiveUrlPastingOnText"`
					LinkShapeColor                       struct {
						R int `json:"r"`
						G int `json:"g"`
						B int `json:"b"`
						A int `json:"a"`
					} `json:"linkShapeColor"`
					ProfileLinkWidth               int           `json:"profileLinkWidth"`
					IsLoading                      bool          `json:"isLoading"`
					SnsActivityGenres              []interface{} `json:"snsActivityGenres"`
					ImageUpLoading                 bool          `json:"imageUpLoading"`
					ShowSavedToast                 bool          `json:"showSavedToast"`
					ToastText                      string        `json:"toastText"`
					ProfileLinkURLType             interface{}   `json:"profileLinkUrlType"`
					ProfileLinkErrors              []interface{} `json:"profileLinkErrors"`
					ModalSnsType                   interface{}   `json:"modalSnsType"`
					SnsModalDefaultURL             string        `json:"snsModalDefaultUrl"`
					MultipleImageLinkIndex         int           `json:"multipleImageLinkIndex"`
					FourImageLinkIndex             int           `json:"fourImageLinkIndex"`
					SelectedIndexOnImageOrSnsModal int           `json:"selectedIndexOnImageOrSnsModal"`
					IsCheckedOpenCategory          bool          `json:"isCheckedOpenCategory"`
					CurrentOpenedGenreIndex        int           `json:"currentOpenedGenreIndex"`
					ShowIconQrCode                 bool          `json:"showIconQrCode"`
				} `json:"creatorDetailEdit"`
				GenreCategory struct {
					IsLoading              bool          `json:"isLoading"`
					SelectedMoreThanOne    bool          `json:"selectedMoreThanOne"`
					GenreCategoryList      []interface{} `json:"genreCategoryList"`
					OpenedGenreCategoryIds []interface{} `json:"openedGenreCategoryIds"`
				} `json:"genreCategory"`
				Profile struct {
					IsLoggedIn             bool        `json:"isLoggedIn"`
					ShowProfileQrcodeModal bool        `json:"showProfileQrcodeModal"`
					ShowCopiedMessage      bool        `json:"showCopiedMessage"`
					ShowIconQrCode         bool        `json:"showIconQrCode"`
					Profile                interface{} `json:"profile"`
				} `json:"profile"`
				LineLogin struct {
					LineLoginResponse interface{} `json:"lineLoginResponse"`
					IsLoading         bool        `json:"isLoading"`
				} `json:"lineLogin"`
				Login struct {
					LoginResponse interface{} `json:"loginResponse"`
					IsLoading     bool        `json:"isLoading"`
				} `json:"login"`
				Modal struct {
					ModalOpened              bool   `json:"modalOpened"`
					ModalComponentName       string `json:"modalComponentName"`
					MasterModalID            string `json:"masterModalId"`
					ConfirmationModalOptions struct {
						ModalText    string `json:"modalText"`
						PositiveText string `json:"positiveText"`
						NegativeText string `json:"negativeText"`
					} `json:"confirmationModalOptions"`
					SelectBackgroundImageModalOptions struct {
						IsButtonLinkDesignImage bool `json:"isButtonLinkDesignImage"`
					} `json:"selectBackgroundImageModalOptions"`
					SelectImageModalOptions struct {
						IsMultipleImageLink bool `json:"isMultipleImageLink"`
						IsButtonLink        bool `json:"isButtonLink"`
					} `json:"selectImageModalOptions"`
					BackgroundOverlayColorModalOptions interface{} `json:"backgroundOverlayColorModalOptions"`
				} `json:"modal"`
				LineMessaging struct {
					LineMessaging interface{} `json:"lineMessaging"`
					IsLoading     bool        `json:"isLoading"`
				} `json:"lineMessaging"`
				PasswordReminder struct {
					PasswordReminderResponse interface{} `json:"passwordReminderResponse"`
					IsLoading                bool        `json:"isLoading"`
					IsCompletedSendEmail     bool        `json:"isCompletedSendEmail"`
					HasErrorResponse         bool        `json:"hasErrorResponse"`
				} `json:"passwordReminder"`
				PasswordChange struct {
					PasswordChangeResponse interface{} `json:"passwordChangeResponse"`
					IsLoading              bool        `json:"isLoading"`
				} `json:"passwordChange"`
				SignUp struct {
					SingUpAuthResponse     interface{} `json:"singUpAuthResponse"`
					SignUpByLineResponse   interface{} `json:"signUpByLineResponse"`
					IsLoading              bool        `json:"isLoading"`
					RegisteredAlready      bool        `json:"registeredAlready"`
					HasAccountByEmailAuth  bool        `json:"hasAccountByEmailAuth"`
					DefaultEmail           string      `json:"defaultEmail"`
					HasErrorSignupResponse bool        `json:"hasErrorSignupResponse"`
				} `json:"signUp"`
				ResendEmailVerification struct {
					IsResendedEmailVerification     bool `json:"isResendedEmailVerification"`
					HasErrorResendEmailVerification bool `json:"hasErrorResendEmailVerification"`
				} `json:"resendEmailVerification"`
				FirebaseAuth struct {
					FirebaseUser  interface{} `json:"firebaseUser"`
					IsAuthLoading bool        `json:"isAuthLoading"`
				} `json:"firebaseAuth"`
				SignupDetail struct {
					IsInstagramConnected bool        `json:"isInstagramConnected"`
					IsTwitterConnected   bool        `json:"isTwitterConnected"`
					IsLoading            bool        `json:"isLoading"`
					IsVerifiedURL        interface{} `json:"isVerifiedUrl"`
				} `json:"signupDetail"`
				LineConnection struct {
					LineConnectionInit     bool        `json:"lineConnectionInit"`
					LineConnectionResponse interface{} `json:"lineConnectionResponse"`
					IsLoading              bool        `json:"isLoading"`
				} `json:"lineConnection"`
				LineSignup struct {
					LineSignUpResponse interface{} `json:"lineSignUpResponse"`
					IsLoading          bool        `json:"isLoading"`
				} `json:"lineSignup"`
				Analytics struct {
					DisplayPeriod            string      `json:"displayPeriod"`
					URLSortType              string      `json:"urlSortType"`
					TopSortType              string      `json:"topSortType"`
					IsURLSortAscendant       bool        `json:"isUrlSortAscendant"`
					IsTopSortAscendant       bool        `json:"isTopSortAscendant"`
					IsReferralSortAscendant  bool        `json:"isReferralSortAscendant"`
					IsDeviceSortAscendant    bool        `json:"isDeviceSortAscendant"`
					PvCounts                 int         `json:"pvCounts"`
					ClickCounts              int         `json:"clickCounts"`
					AccessTopTableSortType   string      `json:"accessTopTableSortType"`
					UserTodayAccessLog       interface{} `json:"userTodayAccessLog"`
					UserOneWeekAccessLog     interface{} `json:"userOneWeekAccessLog"`
					UserOneMonthAccessLog    interface{} `json:"userOneMonthAccessLog"`
					UserThreeMonthsAccessLog interface{} `json:"userThreeMonthsAccessLog"`
					UserSixMonthsAccessLog   interface{} `json:"userSixMonthsAccessLog"`
					UserOneYearAccessLog     interface{} `json:"userOneYearAccessLog"`
					UserAllAccessLog         interface{} `json:"userAllAccessLog"`
					UserGraphAccessLog       struct {
						Labels   []interface{} `json:"labels"`
						Datasets []struct {
							Label           string        `json:"label"`
							BackgroundColor string        `json:"backgroundColor"`
							BorderColor     string        `json:"borderColor"`
							Data            []interface{} `json:"data"`
						} `json:"datasets"`
					} `json:"userGraphAccessLog"`
					UserURLAccessLogs       []interface{} `json:"userUrlAccessLogs"`
					UserTopAccessLogs       []interface{} `json:"userTopAccessLogs"`
					UserReferralAccessLogs  []interface{} `json:"userReferralAccessLogs"`
					UserDeviceAccessLogs    []interface{} `json:"userDeviceAccessLogs"`
					IsAnalyticsStateLoading bool          `json:"isAnalyticsStateLoading"`
					URLAddedAreaHeight      int           `json:"urlAddedAreaHeight"`
					TopAddedAreaHeight      int           `json:"topAddedAreaHeight"`
					ReferralAddedAreaHeight int           `json:"referralAddedAreaHeight"`
					IsShowingMoreOnURL      bool          `json:"isShowingMoreOnUrl"`
					IsShowingMoreOnTop      bool          `json:"isShowingMoreOnTop"`
					IsShowingMoreOnReferral bool          `json:"isShowingMoreOnReferral"`
					IsAnalyticsAPIError     bool          `json:"isAnalyticsApiError"`
					IsShowingToast          bool          `json:"isShowingToast"`
					APIError                interface{}   `json:"apiError"`
				} `json:"analytics"`
				Notification struct {
					IsLoading            bool        `json:"isLoading"`
					SelectedNotification interface{} `json:"selectedNotification"`
				} `json:"notification"`
				CreatorDetailEditTutorial struct {
					EditingCreatorPreferance  interface{} `json:"editingCreatorPreferance"`
					TutorialCount             int         `json:"tutorialCount"`
					IsTutorialButtonEditDone  bool        `json:"isTutorialButtonEditDone"`
					IsTutorialLinkDraggerDone bool        `json:"isTutorialLinkDraggerDone"`
					IsTutorialLinkEditDone    bool        `json:"isTutorialLinkEditDone"`
				} `json:"creatorDetailEditTutorial"`
				AccountDelete struct {
					IsLoading         bool        `json:"isLoading"`
					IsSucceededDelete interface{} `json:"isSucceededDelete"`
				} `json:"accountDelete"`
				ProfileImageNFTModal struct {
				} `json:"profileImageNFTModal"`
				SignupGenre struct {
					IsSucceeded interface{} `json:"isSucceeded"`
					IsLoading   bool        `json:"isLoading"`
				} `json:"signupGenre"`
			} `json:"initialState"`
			ProfileString string `json:"profileString"`
			OgpImageURL   string `json:"ogpImageUrl"`
		} `json:"pageProps"`
		NSsp bool `json:"__N_SSP"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		CreatorURL string `json:"creatorUrl"`
	} `json:"query"`
	BuildID       string        `json:"buildId"`
	IsFallback    bool          `json:"isFallback"`
	Gssp          bool          `json:"gssp"`
	Locale        string        `json:"locale"`
	Locales       []string      `json:"locales"`
	DefaultLocale string        `json:"defaultLocale"`
	ScriptLoader  []interface{} `json:"scriptLoader"`
}

type LitlinkProfile struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Birthday struct {
		Seconds     int `json:"_seconds"`
		Nanoseconds int `json:"_nanoseconds"`
	} `json:"birthday"`
	Genre       string `json:"genre"`
	ProfileText string `json:"profileText"`
	URL         string `json:"url"`
	PictureURL  string `json:"pictureUrl"`
	PictureType string `json:"pictureType"`
	SnsIconLink struct {
		Details []struct {
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"details"`
	} `json:"snsIconLink"`
	ProfileLink struct {
		Details []struct {
			LinkType   string `json:"linkType"`
			ButtonLink struct {
				Title       string `json:"title"`
				URL         string `json:"url"`
				IconURL     string `json:"iconUrl"`
				Description string `json:"description"`
				URLType     string `json:"urlType"`
			} `json:"buttonLink,omitempty"`
			SingleImageLink struct {
				URL         string `json:"url"`
				Description string `json:"description"`
				URLType     string `json:"urlType"`
				Title       string `json:"title"`
				ImageURL    string `json:"imageUrl"`
			} `json:"singleImageLink,omitempty"`
			MusicLink struct {
				URL string `json:"url"`
			} `json:"musicLink,omitempty"`
		} `json:"details"`
	} `json:"profileLink"`
	CreatorDetailLayout struct {
		BackgroundImageURL     string `json:"backgroundImageUrl"`
		BackgroundOverlayColor string `json:"backgroundOverlayColor"`
		TextAlign              string `json:"textAlign"`
		LinkShapeColor         string `json:"linkShapeColor"`
		FontSize               string `json:"fontSize"`
		Template               string `json:"template"`
		BackgroundColor        string `json:"backgroundColor"`
		LinkShapeType          string `json:"linkShapeType"`
		FontFamily             string `json:"fontFamily"`
		FontColor              string `json:"fontColor"`
		BackgroundGradation    string `json:"backgroundGradation"`
	} `json:"creatorDetailLayout"`
	SnsActivitySetting struct {
		GenreSettings []struct {
			ID               string `json:"id"`
			Name             string `json:"name"`
			English          string `json:"english"`
			CategorySettings []struct {
				Name    string `json:"name"`
				English string `json:"english"`
				ID      string `json:"id"`
			} `json:"categorySettings"`
		} `json:"genreSettings"`
	} `json:"snsActivitySetting"`
}

type ApiResponse struct {
	Ok           bool                   `json:"ok"`
	Name         string                 `json:"name"`
	ProfileLinks []LitlinkProfileDetail `json:"profileLink"`
}

type LitlinkProfileDetail struct {
	Title       string `json:"title,"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type LivepocketEventData struct {
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
