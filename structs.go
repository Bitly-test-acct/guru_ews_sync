package main

import "time"

type GuruCardExtended struct {
	Content string `json:"content"`
	Owner   struct {
		Status        string `json:"status"`
		Email         string `json:"email"`
		LastName      string `json:"lastName"`
		FirstName     string `json:"firstName"`
		ProfilePicURL string `json:"profilePicUrl"`
	} `json:"owner"`
	LastModified string `json:"lastModified"`
	Version      int    `json:"version"`
	ID           string `json:"id"`
	Tags         []struct {
		ID           string `json:"id"`
		CategoryName string `json:"categoryName"`
		CategoryID   string `json:"categoryId"`
		Value        string `json:"value"`
	} `json:"tags"`
	Collection struct {
		Color              string `json:"color"`
		ID                 string `json:"id"`
		CollectionType     string `json:"collectionType"`
		Slug               string `json:"slug"`
		RoiEnabled         bool   `json:"roiEnabled"`
		PublishingEnabled  bool   `json:"publishingEnabled"`
		PublicCardsEnabled bool   `json:"publicCardsEnabled"`
		Name               string `json:"name"`
	} `json:"collection"`
	VerificationReasons        []interface{} `json:"verificationReasons"`
	VerificationInitiationDate string        `json:"verificationInitiationDate"`
	KnowledgeAlerts            []interface{} `json:"knowledgeAlerts"`
	HasDrafts                  bool          `json:"hasDrafts"`
	Boards                     []struct {
		Description string `json:"description"`
		Action      struct {
			BoardEntries []struct {
				ID string `json:"id"`
			} `json:"boardEntries"`
		} `json:"action"`
		Title      string `json:"title"`
		ID         string `json:"id"`
		Collection struct {
		} `json:"collection"`
		Items         []interface{} `json:"items"`
		Slug          string        `json:"slug"`
		NumberOfFacts int           `json:"numberOfFacts"`
	} `json:"boards"`
	ShareStatus          string `json:"shareStatus"`
	PreferredPhrase      string `json:"preferredPhrase"`
	VerificationInterval int    `json:"verificationInterval"`
	Verifiers            []struct {
		Type string `json:"type"`
		User struct {
			Status        string `json:"status"`
			Email         string `json:"email"`
			LastName      string `json:"lastName"`
			FirstName     string `json:"firstName"`
			ProfilePicURL string `json:"profilePicUrl"`
		} `json:"user"`
		ID          string `json:"id"`
		DateCreated string `json:"dateCreated"`
	} `json:"verifiers"`
	VerificationType string `json:"verificationType"`
	LastVerified     string `json:"lastVerified"`
	LastVerifiedBy   struct {
		Status        string `json:"status"`
		Email         string `json:"email"`
		LastName      string `json:"lastName"`
		FirstName     string `json:"firstName"`
		ProfilePicURL string `json:"profilePicUrl"`
	} `json:"lastVerifiedBy"`
	LastModifiedBy struct {
		Status        string `json:"status"`
		Email         string `json:"email"`
		LastName      string `json:"lastName"`
		FirstName     string `json:"firstName"`
		ProfilePicURL string `json:"profilePicUrl"`
	} `json:"lastModifiedBy"`
	Attachments []struct {
		ID        string `json:"id"`
		Link      string `json:"link"`
		Filename  string `json:"filename"`
		Mimetype  string `json:"mimetype"`
		Extension string `json:"extension"`
		Size      int    `json:"size"`
	} `json:"attachments"`
	HTMLContent       bool   `json:"htmlContent"`
	TeamID            string `json:"teamId"`
	DateCreated       string `json:"dateCreated"`
	Slug              string `json:"slug"`
	CardType          string `json:"cardType"`
	VerificationState string `json:"verificationState"`
	OriginalOwner     struct {
		Status        string `json:"status"`
		Email         string `json:"email"`
		LastName      string `json:"lastName"`
		FirstName     string `json:"firstName"`
		ProfilePicURL string `json:"profilePicUrl"`
	} `json:"originalOwner"`
	NextVerificationDate  string `json:"nextVerificationDate"`
	GuruSlateToolsVersion string `json:"guruSlateToolsVersion"`
}

type GuruCardsQueryResponse []struct {
	Owner struct {
		Status        string `json:"status"`
		Email         string `json:"email"`
		LastName      string `json:"lastName"`
		FirstName     string `json:"firstName"`
		ProfilePicURL string `json:"profilePicUrl"`
	} `json:"owner"`
	LastModified string `json:"lastModified"`
	ID           string `json:"id"`
	Content      string `json:"content"`
	Collection   struct {
		Color              string `json:"color"`
		ID                 string `json:"id"`
		CollectionType     string `json:"collectionType"`
		RoiEnabled         bool   `json:"roiEnabled"`
		PublicCardsEnabled bool   `json:"publicCardsEnabled"`
		Name               string `json:"name"`
	} `json:"collection"`
	HTMLContent bool   `json:"htmlContent"`
	ShareStatus string `json:"shareStatus"`
	Boards      []struct {
		Title         string        `json:"title"`
		ID            string        `json:"id"`
		Items         []interface{} `json:"items"`
		Slug          string        `json:"slug"`
		NumberOfFacts int           `json:"numberOfFacts"`
	} `json:"boards"`
	PreferredPhrase string `json:"preferredPhrase"`
	Verifiers       []struct {
		Type string `json:"type"`
		User struct {
			Status        string `json:"status"`
			Email         string `json:"email"`
			LastName      string `json:"lastName"`
			FirstName     string `json:"firstName"`
			ProfilePicURL string `json:"profilePicUrl"`
		} `json:"user"`
		ID string `json:"id"`
	} `json:"verifiers"`
	LastVerified   string `json:"lastVerified"`
	LastVerifiedBy struct {
		Status        string `json:"status"`
		Email         string `json:"email"`
		LastName      string `json:"lastName"`
		FirstName     string `json:"firstName"`
		ProfilePicURL string `json:"profilePicUrl"`
	} `json:"lastVerifiedBy"`
	LastModifiedBy struct {
		Status        string `json:"status"`
		Email         string `json:"email"`
		LastName      string `json:"lastName"`
		FirstName     string `json:"firstName"`
		ProfilePicURL string `json:"profilePicUrl"`
	} `json:"lastModifiedBy"`
	VerificationInterval   int           `json:"verificationInterval"`
	VerificationType       string        `json:"verificationType"`
	DateCreated            string        `json:"dateCreated"`
	Slug                   string        `json:"slug"`
	CardType               string        `json:"cardType"`
	HighlightedAttachments []interface{} `json:"highlightedAttachments"`
	VerificationState      string        `json:"verificationState"`
	OriginalOwner          struct {
		Status        string `json:"status"`
		Email         string `json:"email"`
		LastName      string `json:"lastName"`
		FirstName     string `json:"firstName"`
		ProfilePicURL string `json:"profilePicUrl"`
	} `json:"originalOwner"`
	NextVerificationDate string `json:"nextVerificationDate"`
}

type GuruCollectionsResponse []struct {
	Color              string `json:"color"`
	Description        string `json:"description"`
	ID                 string `json:"id"`
	Contexts           int    `json:"contexts"`
	Boards             int    `json:"boards"`
	DateCreated        string `json:"dateCreated"`
	Cards              int    `json:"cards"`
	Slug               string `json:"slug"`
	RoiEnabled         bool   `json:"roiEnabled"`
	AssistEnabled      bool   `json:"assistEnabled"`
	PublishingEnabled  bool   `json:"publishingEnabled"`
	PublicCardsEnabled bool   `json:"publicCardsEnabled"`
	CollectionType     string `json:"collectionType"`
	CollectionStats    struct {
		Stats struct {
			CollectionTrustScore struct {
				Type                   string `json:"type"`
				TrustedCount           int    `json:"trustedCount"`
				NeedsVerificationCount int    `json:"needsVerificationCount"`
			} `json:"collection-trust-score"`
			CardCount struct {
				Type  string `json:"type"`
				Count int    `json:"count"`
			} `json:"card-count"`
		} `json:"stats"`
	} `json:"collectionStats"`
	PublicCards int    `json:"publicCards"`
	Name        string `json:"name"`
}

type EWSGuruDoc struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
}

type EWSGetDocsResults struct {
	LastUpdated     time.Time `json:"last_updated"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
	ContentSourceID string    `json:"content_source_id"`
	Source          string    `json:"source"`
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Body            string    `json:"body"`
	Type            string    `json:"type"`
	URL             string    `json:"url"`
}

type EWSGuruGetDocs struct {
	Results []*EWSGetDocsResults `json:"results"`
	Meta    struct {
		Page struct {
			Current      int `json:"current"`
			TotalPages   int `json:"total_pages"`
			TotalResults int `json:"total_results"`
			Size         int `json:"size"`
		} `json:"page"`
		Warnings []interface{} `json:"warnings"`
		Cursor   struct {
			Current interface{} `json:"current"`
			Next    string      `json:"next"`
		} `json:"cursor"`
	} `json:"meta"`
}

type CursorBody struct {
	Cursor string `json:"cursor"`
}
