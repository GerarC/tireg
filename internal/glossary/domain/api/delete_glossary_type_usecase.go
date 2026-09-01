package api

// DeleteGlossaryTypeUseCase exposes the operation to delete a glossary type owned by a user.
type DeleteGlossaryTypeUseCase interface {
	// Delete removes the glossary type matching the given id, owned by the given user.
	Delete(ownerID string, id string) error
}
