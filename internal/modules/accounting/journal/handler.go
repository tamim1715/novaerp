package journal

import (
	"errors"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/common/response"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateJournalEntry(c *gin.Context) {
	var req CreateJournalEntryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var postedBy *uuid.UUID
	if uidStr, exists := c.Get("userID"); exists {
		if uid, ok := uidStr.(string); ok {
			if parsed, err := uuid.Parse(uid); err == nil {
				postedBy = &parsed
			}
		}
	}

	entry, err := h.service.CreateJournalEntry(c.Request.Context(), req, postedBy)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Journal entry created successfully", ToJournalEntryResponse(entry))
}

func (h *Handler) FindAllJournalEntries(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	status := c.Query("status")
	sourceType := c.Query("sourceType")

	var startDate, endDate *time.Time
	if s := c.Query("startDate"); s != "" {
		if t, err := time.Parse("2006-01-02", s); err == nil {
			startDate = &t
		}
	}
	if e := c.Query("endDate"); e != "" {
		if t, err := time.Parse("2006-01-02", e); err == nil {
			endDate = &t
		}
	}

	entries, total, err := h.service.FindAllJournalEntries(c.Request.Context(), req, status, sourceType, startDate, endDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	req.Normalize()
	pageResponse := pagination.PageResponse{
		Page:       req.Page,
		Size:       req.Size,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(req.Size))),
		Data:       ToJournalEntryResponseList(entries),
	}

	response.Success(c, "Journal entries fetched successfully", pageResponse)
}

func (h *Handler) FindJournalEntryByID(c *gin.Context) {
	id := c.Param("id")
	entry, err := h.service.FindJournalEntryByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Journal entry not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Journal entry fetched successfully", ToJournalEntryResponse(entry))
}

func (h *Handler) PostJournalEntry(c *gin.Context) {
	id := c.Param("id")
	entry, err := h.service.PostJournalEntry(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Journal entry not found")
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Journal entry posted and locked successfully", ToJournalEntryResponse(entry))
}

func (h *Handler) VoidJournalEntry(c *gin.Context) {
	id := c.Param("id")
	var req VoidJournalEntryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var voidedBy *uuid.UUID
	if uidStr, exists := c.Get("userID"); exists {
		if uid, ok := uidStr.(string); ok {
			if parsed, err := uuid.Parse(uid); err == nil {
				voidedBy = &parsed
			}
		}
	}

	original, reversal, err := h.service.VoidJournalEntry(c.Request.Context(), id, req, voidedBy)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Journal entry not found")
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	payload := gin.H{
		"original": ToJournalEntryResponse(original),
		"reversal": ToJournalEntryResponse(reversal),
	}

	response.Success(c, "Journal entry voided and reversal entry posted successfully", payload)
}
