package hipchat

import (
	"fmt"
	"net/http"
)

// RoomService gives access to the room related methods of the API.
type RoomService struct {
	client *Client
}

// Rooms represents a HipChat room list.
type Rooms struct {
	Items      []Room    `json:"items"`
	StartIndex int       `json:"startIndex"`
	MaxResults int       `json:"maxResults"`
	Links      PageLinks `json:"links"`
}

// Room represents a HipChat room.
type Room struct {
	ID                int            `json:"id"`
	Links             RoomLinks      `json:"links"`
	Name              string         `json:"name"`
	XmppJid           string         `json:"xmpp_jid"`
	Statistics        RoomStatistics `json:"statistics"`
	Created           string         `json:"created"`
	IsArchived        bool           `json:"is_archived"`
	Privacy           string         `json:"privacy"`
	IsGuestAccessible bool           `json:"is_guess_accessible"`
	Topic             string         `json:"topic"`
	Participants      []User         `json:"participants"`
	Owner             User           `json:"owner"`
	GuestAccessURL    string         `json:"guest_access_url"`
}

// RoomStatistics represents the HipChat room statistics.
type RoomStatistics struct {
	Links Links `json:"links"`
}

// CreateRoomRequest represents a HipChat room creation request.
type CreateRoomRequest struct {
	Topic       string `json:"topic,omitempty"`
	GuestAccess bool   `json:"guest_access,omitempty"`
	Name        string `json:"name,omitempty"`
	OwnerUserID string `json:"owner_user_id,omitempty"`
	Privacy     string `json:"privacy,omitempty"`
}

// UpdateRoomRequest represents a HipChat room update request.
type UpdateRoomRequest struct {
	Name          string `json:"name"`
	Topic         string `json:"topic"`
	IsGuestAccess bool   `json:"is_guest_access"`
	IsArchived    bool   `json:"is_archived"`
	Privacy       string `json:"privacy"`
	Owner         ID     `json:"owner"`
}

// RoomLinks represents the HipChat room links.
type RoomLinks struct {
	Links
	Webhooks     string `json:"webhooks"`
	Members      string `json:"members"`
	Participants string `json:"participants"`
}

// NotificationRequest represents a HipChat room notification request.
type NotificationRequest struct {
	Color         string `json:"color,omitempty"`
	Message       string `json:"message,omitempty"`
	Notify        bool   `json:"notify,omitempty"`
	MessageFormat string `json:"message_format,omitempty"`
}

// ShareFileRequest represents a HipChat room file share request.
type ShareFileRequest struct {
	Path     string `json:"path"`
	Filename string `json:"filename,omitempty"`
	Message  string `json:"message,omitempty"`
}

// History represents a HipChat room chat history.
type History struct {
	Items      []Message `json:"items"`
	StartIndex int       `json:"startIndex"`
	MaxResults int       `json:"maxResults"`
	Links      PageLinks `json:"links"`
}

// Message represents a HipChat message.
type Message struct {
	Date          string      `json:"date"`
	From          interface{} `json:"from"` // string | obj <- weak
	ID            string      `json:"id"`
	Mentions      []User      `json:"mentions"`
	Message       string      `json:"message"`
	MessageFormat string      `json:"message_format"`
	Type          string      `json:"type"`
}

// SetTopicRequest represents a hipchat update topic request
type SetTopicRequest struct {
	Topic string `json:"topic"`
}

// InviteRequest represents a hipchat invite to room request
type InviteRequest struct {
	Reason string `json:"reason"`
}

// List returns all the rooms authorized.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_all_rooms
func (r *RoomService) List() (*Rooms, *http.Response, error) {
	req, err := r.client.NewRequest("GET", "room", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	rooms := new(Rooms)
	resp, err := r.client.Do(req, rooms)
	if err != nil {
		return nil, resp, err
	}
	return rooms, resp, nil
}

// Get returns the room specified by the id.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_room
func (r *RoomService) Get(id string) (*Room, *http.Response, error) {
	req, err := r.client.NewRequest("GET", fmt.Sprintf("room/%s", id), nil, nil)
	if err != nil {
		return nil, nil, err
	}

	room := new(Room)
	resp, err := r.client.Do(req, room)
	if err != nil {
		return nil, resp, err
	}
	return room, resp, nil
}

// Notification sends a notification to the room specified by the id.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/send_room_notification
func (r *RoomService) Notification(id string, notifReq *NotificationRequest) (*http.Response, error) {
	req, err := r.client.NewRequest("POST", fmt.Sprintf("room/%s/notification", id), nil, notifReq)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}

// ShareFile sends a file to the room specified by the id.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/share_file_with_room
func (r *RoomService) ShareFile(id string, shareFileReq *ShareFileRequest) (*http.Response, error) {
	req, err := r.client.NewFileUploadRequest("POST", fmt.Sprintf("room/%s/share/file", id), shareFileReq)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}

// Create creates a new room.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/create_room
func (r *RoomService) Create(roomReq *CreateRoomRequest) (*Room, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "room", nil, roomReq)
	if err != nil {
		return nil, nil, err
	}

	room := new(Room)
	resp, err := r.client.Do(req, room)
	if err != nil {
		return nil, resp, err
	}
	return room, resp, nil
}

// Delete deletes an existing room.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/delete_room
func (r *RoomService) Delete(id string) (*http.Response, error) {
	req, err := r.client.NewRequest("DELETE", fmt.Sprintf("room/%s", id), nil, nil)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}

// Update updates an existing room.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/update_room
func (r *RoomService) Update(id string, roomReq *UpdateRoomRequest) (*http.Response, error) {
	req, err := r.client.NewRequest("PUT", fmt.Sprintf("room/%s", id), nil, roomReq)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}

// HistoryOptions represents a HipChat room chat history request.
type HistoryOptions struct {
	ListOptions

	// Either the latest date to fetch history for in ISO-8601 format, or 'recent' to fetch
	// the latest 75 messages. Paging isn't supported for 'recent', however they are real-time
	// values, whereas date queries may not include the most recent messages.
	Date string `url:"date,omitempty"`

	// Your timezone. Must be a supported timezone
	Timezone string `url:"timezone,omitempty"`

	// Reverse the output such that the oldest message is first.
	// For consistent paging, set to 'false'.
	Reverse bool `url:"reverse,omitempty"`
}

// History fetches a room's chat history.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/view_room_history
func (r *RoomService) History(id string, opt *HistoryOptions) (*History, *http.Response, error) {
	u := fmt.Sprintf("room/%s/history", id)
	req, err := r.client.NewRequest("GET", u, opt, nil)
	h := new(History)
	resp, err := r.client.Do(req, &h)
	if err != nil {
		return nil, resp, err
	}
	return h, resp, nil
}

// LatestHistoryOptions represents a HipChat room chat latest history request.
type LatestHistoryOptions struct {

	// The maximum number of messages to return.
	MaxResults int `url:"max-results,omitempty"`

	// Your timezone. Must be a supported timezone.
	Timezone string `url:"timezone,omitempty"`

	// The id of the message that is oldest in the set of messages to be returned.
	// The server will not return any messages that chronologically precede this message.
	NotBefore string `url:"not-before,omitempty"`
}

// Latest fetches a room's chat history.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/view_recent_room_history
func (r *RoomService) Latest(id string, opt *LatestHistoryOptions) (*History, *http.Response, error) {
	u := fmt.Sprintf("room/%s/history/latest", id)
	req, err := r.client.NewRequest("GET", u, opt, nil)
	h := new(History)
	resp, err := r.client.Do(req, &h)
	if err != nil {
		return nil, resp, err
	}
	return h, resp, nil
}

// SetTopic sets Room topic.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/set_topic
func (r *RoomService) SetTopic(id string, topic string) (*http.Response, error) {
	topicReq := &SetTopicRequest{Topic: topic}

	req, err := r.client.NewRequest("PUT", fmt.Sprintf("room/%s/topic", id), nil, topicReq)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}

// Invite someone to the Room.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/invite_user
func (r *RoomService) Invite(room string, user string, reason string) (*http.Response, error) {
	reasonReq := &InviteRequest{Reason: reason}

	req, err := r.client.NewRequest("POST", fmt.Sprintf("room/%s/invite/%s", room, user), nil, reasonReq)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}
