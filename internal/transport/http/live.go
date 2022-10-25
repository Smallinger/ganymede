package http

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zibbp/ganymede/ent"
	"github.com/zibbp/ganymede/internal/live"
	"net/http"
)

type LiveService interface {
	GetLiveWatchedChannels(c echo.Context) ([]*ent.Live, error)
	AddLiveWatchedChannel(c echo.Context, liveDto live.Live) (*ent.Live, error)
	DeleteLiveWatchedChannel(c echo.Context, lID uuid.UUID) error
	UpdateLiveWatchedChannel(c echo.Context, liveDto live.Live) (*ent.Live, error)
	Check() error
	ConvertChat(c echo.Context, convertDto live.ConvertChat) error
}

type AddWatchedChannelRequest struct {
	ChannelID   string `json:"channel_id" validate:"required"`
	Resolution  string `json:"resolution" validate:"required,oneof=best source 720p60 480p30 360p30 160p30"`
	ArchiveChat bool   `json:"archive_chat"`
}

type UpdateWatchedChannelRequest struct {
	Resolution  string `json:"resolution" validate:"required,oneof=best source 720p60 480p30 360p30 160p30"`
	ArchiveChat bool   `json:"archive_chat"`
}

type ConvertChatRequest struct {
	FileName      string `json:"file_name" validate:"required"`
	ChannelName   string `json:"channel_name" validate:"required"`
	VodID         string `json:"vod_id" validate:"required"`
	ChannelID     int    `json:"channel_id" validate:"required"`
	VodExternalID string `json:"vod_external_id" validate:"required"`
	ChatStart     string `json:"chat_start" validate:"required"`
}

func (h *Handler) GetLiveWatchedChannels(c echo.Context) error {
	channels, err := h.Service.LiveService.GetLiveWatchedChannels(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, channels)
}

func (h *Handler) AddLiveWatchedChannel(c echo.Context) error {
	ccr := new(AddWatchedChannelRequest)
	if err := c.Bind(ccr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(ccr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cUUID, err := uuid.Parse(ccr.ChannelID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	liveDto := live.Live{
		ID:          cUUID,
		IsLive:      false,
		ArchiveChat: ccr.ArchiveChat,
		Resolution:  ccr.Resolution,
	}
	l, err := h.Service.LiveService.AddLiveWatchedChannel(c, liveDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, l)
}

func (h *Handler) UpdateLiveWatchedChannel(c echo.Context) error {
	id := c.Param("id")
	lID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ccr := new(UpdateWatchedChannelRequest)
	if err := c.Bind(ccr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(ccr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	liveDto := live.Live{
		ID:          lID,
		ArchiveChat: ccr.ArchiveChat,
		Resolution:  ccr.Resolution,
	}
	l, err := h.Service.LiveService.UpdateLiveWatchedChannel(c, liveDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, l)
}

func (h *Handler) DeleteLiveWatchedChannel(c echo.Context) error {
	id := c.Param("id")
	lID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.Service.LiveService.DeleteLiveWatchedChannel(c, lID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) Check(c echo.Context) error {
	err := h.Service.LiveService.Check()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "ok")
}

// ConvertChat Adhoc convert live chat to TD chat format
func (h *Handler) ConvertChat(c echo.Context) error {
	ccr := new(ConvertChatRequest)
	if err := c.Bind(ccr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(ccr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	convertDto := live.ConvertChat{
		FileName:      ccr.FileName,
		ChannelName:   ccr.ChannelName,
		VodID:         ccr.VodID,
		ChannelID:     ccr.ChannelID,
		VodExternalID: ccr.VodExternalID,
		ChatStart:     ccr.ChatStart,
	}
	err := h.Service.LiveService.ConvertChat(c, convertDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "ok - converted chat found in /tmp/")
}
