package api

// DeleteTimeEntryUseCase exposes the operation to delete a time entry owned by a user.
type DeleteTimeEntryUseCase interface {
	// Delete removes the time entry matching the given id, owned by the given user.
	Delete(ownerID string, id string) error
}
