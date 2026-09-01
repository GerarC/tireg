package api

// DeleteTaskMappingUseCase exposes the operation to delete a task mapping owned by a user.
type DeleteTaskMappingUseCase interface {
	// Delete removes the task mapping matching the given id, owned by the given user.
	Delete(ownerID string, id string) error
}
