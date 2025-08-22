package dto

type TorrentMap map[string]*Torrent

type UpdateUiTorrentResponse struct {
	Torrents TorrentMap `json:"torrents,omitempty"`
}

type Torrent struct {
	DistributedCopies   float64 `json:"distributed_copies,omitempty"`
	DownloadPayloadRate int     `json:"download_payload_rate,omitempty"`
	IsAutoManaged       bool    `json:"is_auto_managed,omitempty"`
	MaxDownloadSpeed    float64 `json:"max_download_speed,omitempty"`
	MaxUploadSpeed      float64 `json:"max_upload_speed,omitempty"`
	NumPeers            int     `json:"num_peers,omitempty"`
	NumSeeds            int     `json:"num_seeds,omitempty"`
	Progress            float64 `json:"progress,omitempty"`
	DownloadLocation    string  `json:"download_location,omitempty"`
	SeedsPeersRatio     float64 `json:"seeds_peers_ratio,omitempty"`
	State               string  `json:"state,omitempty"`
	TimeAdded           int     `json:"time_added,omitempty"`
	TotalDone           int     `json:"total_done,omitempty"`
	TotalPeers          int     `json:"total_peers,omitempty"`
	TotalSeeds          int     `json:"total_seeds,omitempty"`
	TotalUploaded       int     `json:"total_uploaded,omitempty"`
	TotalWanted         int64   `json:"total_wanted,omitempty"`
	TotalRemaining      int64   `json:"total_remaining,omitempty"`
	TrackerHost         string  `json:"tracker_host,omitempty"`
	UploadPayloadRate   int     `json:"upload_payload_rate,omitempty"`
	Eta                 int     `json:"eta,omitempty"`
	Queue               int     `json:"queue,omitempty"`
	Ratio               float64 `json:"ratio,omitempty"`
	CompletedTime       int     `json:"completed_time,omitempty"`
	LastSeenComplete    int     `json:"last_seen_complete,omitempty"`
	Name                string  `json:"name,omitempty"`
	TimeSinceTransfer   int     `json:"time_since_transfer,omitempty"`
}
