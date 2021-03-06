package adcom1

import "encoding/json"

// Audio object provides additional detail about an ad specifically for audio ads.
type Audio struct {
	// Attribute:
	//   mime
	// Type:
	//   string array
	// Definition:
	//   Mime type(s) of the ad creative(s) (e.g., “audio/mp4”).
	MIME []string `json:"mime,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Definition:
	//   API required by the ad if applicable.
	//   Refer to List: API Frameworks.
	API []APIFramework `json:"api,omitempty"`

	// Attribute:
	//   ctype
	// Type:
	//   integer
	// Definition:
	//   Subtype of audio creative.
	//   Refer to List: Creative Subtypes - Audio/Video.
	CType MediaCreativeSubtype `json:"ctype,omitempty"`

	// Attribute:
	//   dur
	// Type:
	//   integer
	// Definition:
	//   Duration of the audio creative in seconds.
	Dur int64 `json:"dur,omitempty"`

	// Attribute:
	//   adm
	// Type:
	//   string
	// Definition:
	//   Audio markup (e.g., DAAST). Note that including both adm and curl is not recommended.
	AdM string `json:"adm,omitempty"`

	// Attribute:
	//   curl
	// Type:
	//   string
	// Definition:
	//   Optional means of retrieving markup by reference; a URL that returns audio markup (e.g., DAAST).
	//   If this ad is matched to a Placement specification, the Placement.curlx attribute indicates if this markup retrieval option is supported.
	//   Note that including both adm and curl is not recommended.
	CURL string `json:"curl,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
