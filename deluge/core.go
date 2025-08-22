package deluge

import (
	"delugerpc/dto"
	"fmt"
)

// modifyTorrentsState is a helper function to modify the state of multiple torrents
func (c *Client) modifyTorrentsState(method string, torrentIds []string) error {
	_, err := SendRequest[bool](c, method, torrentIds)
	return err
}

// CoreAddTorrentFile adds a torrent file to the session
// Returns The torrent_id or None
func (c *Client) CoreAddTorrentFile(filename string) (string, error) {
	resp, err := SendRequest[string](c, "core.add_torrent_file", filename)
	if err != nil {
		return "", err
	}

	if resp.Result == nil {
		return "", nil
	} else {
		return *resp.Result, nil
	}
}

// CoreAddTorrentMagnet adds a torrent from a magnet link
// Returns The torrent_id or None
func (c *Client) CoreAddTorrentMagnet(magnetLink string) (string, error) {
	resp, err := SendRequest[string](c, "core.add_torrent_magnet", magnetLink)
	if err != nil {
		return "", err
	}

	if resp.Result == nil {
		return "", nil
	} else {
		return *resp.Result, nil
	}
}

// CoreAddTorrentUrl adds a torrent from a url. Deluge will attempt to fetch the torrent from url prior to adding it to the session.
// Returns a Deferred which returns the torrent_id as a str or None
func (c *Client) CoreAddTorrentUrl(torrentUrl string) (string, error) {
	resp, err := SendRequest[string](c, "core.add_torrent_url", torrentUrl)
	if err != nil {
		return "", err
	}

	if resp.Result == nil {
		return "", nil
	} else {
		return *resp.Result, nil
	}
}

// CoreDisablePlugin disables a plugin by name
func (c *Client) CoreDisablePlugin(pluginName string) error {
	resp, err := SendRequest[bool](c, "core.disable_plugin", pluginName)
	if err != nil {
		return err
	}

	if resp.Result == nil || *resp.Result != true {
		return fmt.Errorf("failed to disable plugin: %s", pluginName)
	}
	return nil
}

// CoreEnablePlugin enables a plugin by name
func (c *Client) CoreEnablePlugin(pluginName string) error {
	resp, err := SendRequest[bool](c, "core.enable_plugin", pluginName)
	if err != nil {
		return err
	}

	if resp.Result == nil || *resp.Result != true {
		return fmt.Errorf("failed to disable plugin: %s", pluginName)
	}
	return nil
}

// CoreForceReannounce forces a reannounce of the torrents
func (c *Client) CoreForceReannounce(torrentIds ...string) error {
	_, err := SendRequest[bool](c, "core.force_reannounce", torrentIds)
	if err != nil {
		return err
	}
	return nil
}

// CoreForceRecheck forces a recheck of the torrents
func (c *Client) CoreForceRecheck(torrentIds ...string) error {
	_, err := SendRequest[bool](c, "core.force_recheck", torrentIds)
	if err != nil {
		return err
	}
	return nil

}

// CoreGetAvailablePlugins returns a list of available plugins
func (c *Client) CoreGetAvailablePlugins() ([]string, error) {
	resp, err := SendRequest[[]string](c, "core.get_available_plugins")
	if err != nil {
		return nil, err
	}

	if resp.Result == nil {
		resp.Result = &[]string{}
	}

	return *resp.Result, nil
}

// CoreGetConfig returns the current deluge config
func (c *Client) CoreGetConfig() (*dto.CoreConfigResponse, error) {
	resp, err := SendRequest[dto.CoreConfigResponse](c, "core.get_config")
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// CoreGetEnabledPlugins returns a list of enabled plugins
func (c *Client) CoreGetEnabledPlugins() ([]string, error) {
	resp, err := SendRequest[[]string](c, "core.get_enabled_plugins")
	if err != nil {
		return nil, err
	}

	if resp.Result == nil {
		resp.Result = &[]string{}
	}

	return *resp.Result, nil
}

// CoreGetExternalIp returns the external ip of the client
func (c *Client) CoreGetExternalIp() (string, error) {
	resp, err := SendRequest[string](c, "core.get_external_ip")
	if err != nil || resp.Result == nil {
		return "", err
	}

	return *resp.Result, nil
}

// CoreGetFreeSpace returns the free space in bytes
func (c *Client) CoreGetFreeSpace() (int64, error) {
	resp, err := SendRequest[int64](c, "core.get_free_space")
	if err != nil || resp.Result == nil {
		return 0, err
	}

	return *resp.Result, nil
}

// CoreGetLibtorrentVersion Returns the libtorrent version
func (c *Client) CoreGetLibtorrentVersion() (string, error) {
	resp, err := SendRequest[string](c, "core.get_config")
	if err != nil || resp.Result == nil {
		return "", err
	}

	return *resp.Result, nil
}

// CoreGetListenPort returns the listen port of the client
func (c *Client) CoreGetListenPort() (int, error) {
	resp, err := SendRequest[int](c, "core.get_listen_port")
	if err != nil || resp.Result == nil {
		return 0, err
	}

	return *resp.Result, nil
}

// CoreGetMagnetUri returns the magnet uri for the torrent
func (c *Client) CoreGetMagnetUri(torrentId string) (string, error) {
	resp, err := SendRequest[string](c, "core.get_magnet_uri", torrentId)
	if err != nil || resp.Result == nil {
		return "", err
	}

	return *resp.Result, nil
}

// CoreGetProxy returns the current proxy configuration
func (c *Client) CoreGetProxy() (*dto.CoreConfigProxy, error) {
	resp, err := SendRequest[dto.CoreConfigProxy](c, "core.get_proxy")
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// CoreGetSessionState returns the current session state
func (c *Client) CoreGetSessionState() ([]string, error) {
	resp, err := SendRequest[[]string](c, "core.get_session_state")
	if err != nil {
		return nil, err
	}

	if resp.Result == nil {
		resp.Result = &[]string{}
	}

	return *resp.Result, nil
}

// CoreListTorrents returns all torrents that match the search string
func (c *Client) CoreListTorrents(filterBy *dto.Torrent) (dto.TorrentMap, error) {
	valueFields := []string{"queue", "name", "total_wanted", "state", "progress", "num_seeds", "total_seeds", "num_peers", "total_peers", "download_payload_rate", "upload_payload_rate", "eta", "ratio", "distributed_copies", "is_auto_managed", "time_added", "tracker_host", "download_location", "last_seen_complete", "total_done", "total_uploaded", "max_download_speed", "max_upload_speed", "seeds_peers_ratio", "total_remaining", "completed_time", "time_since_transfer"}

	resp, err := SendRequest[dto.UpdateUiTorrentResponse](c, "core.add_torrent_url", valueFields, filterBy)
	if err != nil {
		return nil, err
	}

	if resp.Result == nil {
		return nil, nil
	}

	return resp.Result.Torrents, nil
}

// CorePauseTorrents pauses torrents by ID if they are not already paused
func (c *Client) CorePauseTorrents(torrentIds ...string) error {
	return c.modifyTorrentsState("core.pause_torrents", torrentIds)
}

// CoreQueueBottom moves torrents to the bottom of the queue
func (c *Client) CoreQueueBottom(torrentIds ...string) error {
	return c.modifyTorrentsState("core.queue_bottom", torrentIds)
}

// CoreQueueDown moves torrents down the queue by 1
func (c *Client) CoreQueueDown(torrentIds ...string) error {
	return c.modifyTorrentsState("core.queue_down", torrentIds)
}

// CoreQueueTop moves torrents to the top of the queue
func (c *Client) CoreQueueTop(torrentIds ...string) error {
	return c.modifyTorrentsState("core.queue_top", torrentIds)
}

// CoreQueueUp moves torrents up the queue by 1
func (c *Client) CoreQueueUp(torrentIds ...string) error {
	return c.modifyTorrentsState("core.queue_up", torrentIds)
}

// CoreRemoveTorrents removes torrents by ID
func (c *Client) CoreRemoveTorrents(removeData bool, torrentIds ...string) error {
	_, err := SendRequest[[]any](c, "core.remove_torrents", torrentIds, removeData)
	return err

}

// CoreResumeTorrents resumes torrents by ID if they have been paused
func (c *Client) CoreResumeTorrents(torrentIds ...string) error {
	return c.modifyTorrentsState("core.resume_torrents", torrentIds)
}

// CoreRescanPlugins rescans the plugin folders for new plugins
func (c *Client) CoreRescanPlugins() error {
	_, err := SendRequest[any](c, "core.rescan_plugins")
	return err
}

// CoreSetConfig sets the current deluge config, changing only what is provided in the input
func (c *Client) CoreSetConfig(config *dto.CoreConfigResponse) error {
	_, err := SendRequest[any](c, "core.set_config", config)
	return err
}

// CoreTestListenPort tests the listen port of the client. This operation can take a long time to complete.
func (c *Client) CoreTestListenPort() (bool, error) {
	resp, err := SendRequest[bool](c, "core.test_listen_port")

	if err != nil || resp.Result == nil {
		return false, err
	}

	return *resp.Result, nil
}
