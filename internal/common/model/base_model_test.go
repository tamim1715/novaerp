package model

import (
	"testing"

	"github.com/google/uuid"
)

func TestBaseModel_BeforeCreate(t *testing.T) {
	t.Run("Assigns new UUID if ID is Nil", func(t *testing.T) {
		bm := &BaseModel{}
		err := bm.BeforeCreate(nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if bm.ID == uuid.Nil {
			t.Error("expected non-nil UUID, got uuid.Nil")
		}
	})

	t.Run("Preserves existing UUID if ID is not Nil", func(t *testing.T) {
		existingID := uuid.New()
		bm := &BaseModel{ID: existingID}
		err := bm.BeforeCreate(nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if bm.ID != existingID {
			t.Errorf("expected ID %s, got %s", existingID, bm.ID)
		}
	})
}
