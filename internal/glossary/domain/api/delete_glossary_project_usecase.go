package api

// DeleteGlossaryProjectUseCase exposes the operation to delete a glossary project owned by a user.
type DeleteGlossaryProjectUseCase interface {
	// Delete removes the glossary project matching the given id, owned by the given user.
	Delete(ownerID string, id string) error
}
