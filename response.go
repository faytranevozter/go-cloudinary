package cloudinary

// UploadImageResponse model
type UploadImageResponse struct {
	PublicID          string        `json:"public_id"`
	Version           int64         `json:"version"`
	Signature         string        `json:"signature"`
	Width             int64         `json:"width"`
	Height            int64         `json:"height"`
	Format            string        `json:"format"`
	ResourceType      string        `json:"resource_type"`
	CreatedAt         string        `json:"created_at"`
	Tags              []interface{} `json:"tags"`
	Bytes             int64         `json:"bytes"`
	Type              string        `json:"type"`
	Etag              string        `json:"etag"`
	Placeholder       bool          `json:"placeholder"`
	URL               string        `json:"url"`
	SecureURL         string        `json:"secure_url"`
	AccessMode        string        `json:"access_mode"`
	OriginalFilename  string        `json:"original_filename"`
	OriginalExtension string        `json:"original_extension"`
}
