package model

import "encoding/json"

type DisabledFunctionality struct {
	DisableTransfers   bool `json:"disableTransfers"`
	DisableKYC         bool `json:"disableKYC"`
	DisableWallet      bool `json:"disableWallet"`
	DisableBalance     bool `json:"disableBalance"`
	DisableSpot        bool `json:"disableSpot"`
	DisablePerps       bool `json:"disablePerps"`
	DisableDarkPool    bool `json:"disableDarkPool"`
	DisableCross       bool `json:"disableCross"`
	DisableAPIKeys     bool `json:"disableAPIKeys"`
	DisableSubaccounts bool `json:"disableSubaccounts"`
	DisableNetworkLink bool `json:"disableNetworkLink"`
	DisableReferrals   bool `json:"disableReferrals"`
}

type Account struct {
	KYCStatus             string                 `json:"KYCStatus"`             // KYC status of the account
	AccountType           string                 `json:"AccountType"`           // Type of the account (e.g., "regular")
	Country               string                 `json:"Country"`               // Country of the account
	DisabledFunctionality *DisabledFunctionality `json:"DisabledFunctionality"` // Functionality disabled for the account
	TermsVersion          int                    `json:"TermsVersion"`          // Version of terms accepted
	TermsUnixSecs         int                    `json:"TermsUnixSecs"`         // Timestamp of terms acceptance
	PrivacyVersion        int                    `json:"PrivacyVersion"`        // Version of privacy policy accepted
	PrivacyUnixSecs       int                    `json:"PrivacyUnixSecs"`       // Timestamp of privacy policy acceptance
	MarketingConsent      string                 `json:"MarketingConsent"`      // Marketing consent status (e.g., "accepted")
	Subaccount            json.RawMessage        `json:"Subaccount,omitempty"`  // Optional subaccount information
}
