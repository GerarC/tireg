package usecase_test

import (
	"errors"
	"testing"

	sharedException "github.com/gerarc/tireg/internal/shared/domain/exception"

	taskModel "github.com/gerarc/tireg/internal/task/domain/model"

	"github.com/gerarc/tireg/internal/time-registry/domain/model"
	"github.com/gerarc/tireg/internal/time-registry/domain/usecase"
)

type stubTimeEntryCommandRepository struct {
	insertFunc             func(model.TimeEntry) (model.TimeEntry, error)
	updateByIDAndOwnerFunc func(string, string, model.TimeEntry) (model.TimeEntry, error)
	deleteByIDAndOwnerFunc func(string, string) error
}

func (stub *stubTimeEntryCommandRepository) Insert(timeEntry model.TimeEntry) (model.TimeEntry, error) {
	return stub.insertFunc(timeEntry)
}

func (stub *stubTimeEntryCommandRepository) UpdateByIDAndOwner(id string, ownerID string, timeEntry model.TimeEntry) (model.TimeEntry, error) {
	return stub.updateByIDAndOwnerFunc(id, ownerID, timeEntry)
}

func (stub *stubTimeEntryCommandRepository) DeleteByIDAndOwner(id string, ownerID string) error {
	return stub.deleteByIDAndOwnerFunc(id, ownerID)
}

type stubTimeEntryQueryRepository struct {
	selectAllByOwnerFunc   func(string) ([]model.TimeEntry, error)
	selectByIDAndOwnerFunc func(string, string) (model.TimeEntry, error)
}

func (stub *stubTimeEntryQueryRepository) SelectAllByOwner(ownerID string) ([]model.TimeEntry, error) {
	return stub.selectAllByOwnerFunc(ownerID)
}

func (stub *stubTimeEntryQueryRepository) SelectByIDAndOwner(id string, ownerID string) (model.TimeEntry, error) {
	return stub.selectByIDAndOwnerFunc(id, ownerID)
}

type stubListTaskMappingsUseCase struct {
	listFunc func(string) ([]taskModel.TaskMapping, error)
}

func (stub *stubListTaskMappingsUseCase) List(ownerID string) ([]taskModel.TaskMapping, error) {
	return stub.listFunc(ownerID)
}

func validTimeEntry() model.TimeEntry {
	return model.TimeEntry{
		Date:         "2026-08-03",
		ProjectLabel: "Abastecar",
		Start:        "08:00",
		End:          "10:30",
		Hours:        2.5,
		Description:  "Desarrollo backend HU",
	}
}

func TestCreate_Succeeds_WhenTimeEntryValid(t *testing.T) {
	commandRepository := &stubTimeEntryCommandRepository{
		insertFunc: func(timeEntry model.TimeEntry) (model.TimeEntry, error) {
			timeEntry.ID = "generated-id"
			return timeEntry, nil
		},
	}
	listTaskMappingsUseCase := &stubListTaskMappingsUseCase{
		listFunc: func(ownerID string) ([]taskModel.TaskMapping, error) { return nil, nil },
	}

	createTimeEntryUseCase := usecase.NewCreateTimeEntryUseCase(commandRepository, listTaskMappingsUseCase)

	timeEntry, err := createTimeEntryUseCase.Create("owner-id", validTimeEntry())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if timeEntry.ID != "generated-id" || timeEntry.OwnerID != "owner-id" {
		t.Fatalf("unexpected time entry: %+v", timeEntry)
	}
}

func TestCreate_ReturnsValidationError_WhenRequiredFieldsMissing(t *testing.T) {
	listTaskMappingsUseCase := &stubListTaskMappingsUseCase{
		listFunc: func(ownerID string) ([]taskModel.TaskMapping, error) { return nil, nil },
	}
	createTimeEntryUseCase := usecase.NewCreateTimeEntryUseCase(&stubTimeEntryCommandRepository{}, listTaskMappingsUseCase)

	_, err := createTimeEntryUseCase.Create("owner-id", model.TimeEntry{})

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 400 {
		t.Fatalf("expected a 400 DomainError, got %v", err)
	}

	if len(domainError.Details) != 5 {
		t.Fatalf("expected 5 validation details (date, project label, start, end, hours), got %d: %v", len(domainError.Details), domainError.Details)
	}
}

func TestCreate_AutoFillsFromMatchingTaskMapping_WhenClassificationFieldsBlank(t *testing.T) {
	commandRepository := &stubTimeEntryCommandRepository{
		insertFunc: func(timeEntry model.TimeEntry) (model.TimeEntry, error) {
			return timeEntry, nil
		},
	}
	listTaskMappingsUseCase := &stubListTaskMappingsUseCase{
		listFunc: func(ownerID string) ([]taskModel.TaskMapping, error) {
			return []taskModel.TaskMapping{
				{
					ProjectLabel:  "Abastecar",
					TypeKey:       "dev",
					IssueKey:      "EA-22",
					MatchKeywords: []string{"desarrollo backend"},
				},
			}, nil
		},
	}

	createTimeEntryUseCase := usecase.NewCreateTimeEntryUseCase(commandRepository, listTaskMappingsUseCase)

	timeEntry := validTimeEntry()
	timeEntry.ProjectLabel = ""

	result, err := createTimeEntryUseCase.Create("owner-id", timeEntry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.ProjectLabel != "Abastecar" || result.TypeKey != "dev" || result.IssueKey != "EA-22" {
		t.Fatalf("expected auto-filled classification, got %+v", result)
	}
}

func TestCreate_ExplicitFieldsWinOverAutoFill(t *testing.T) {
	commandRepository := &stubTimeEntryCommandRepository{
		insertFunc: func(timeEntry model.TimeEntry) (model.TimeEntry, error) {
			return timeEntry, nil
		},
	}
	listTaskMappingsUseCase := &stubListTaskMappingsUseCase{
		listFunc: func(ownerID string) ([]taskModel.TaskMapping, error) {
			return []taskModel.TaskMapping{
				{ProjectLabel: "Portal", TypeKey: "support", IssueKey: "PORT01-1", MatchKeywords: []string{"desarrollo backend"}},
			}, nil
		},
	}

	createTimeEntryUseCase := usecase.NewCreateTimeEntryUseCase(commandRepository, listTaskMappingsUseCase)

	timeEntry := validTimeEntry()
	timeEntry.TypeKey = "explicit-type"

	result, err := createTimeEntryUseCase.Create("owner-id", timeEntry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.ProjectLabel != "Abastecar" {
		t.Fatalf("expected explicit project label to win, got %q", result.ProjectLabel)
	}

	if result.TypeKey != "explicit-type" {
		t.Fatalf("expected explicit type key to win, got %q", result.TypeKey)
	}
}

func TestCreate_NoAutoFill_WhenNoMappingMatches(t *testing.T) {
	listTaskMappingsUseCase := &stubListTaskMappingsUseCase{
		listFunc: func(ownerID string) ([]taskModel.TaskMapping, error) {
			return []taskModel.TaskMapping{
				{ProjectLabel: "Portal", TypeKey: "support", MatchKeywords: []string{"despliegue"}},
			}, nil
		},
	}
	createTimeEntryUseCase := usecase.NewCreateTimeEntryUseCase(&stubTimeEntryCommandRepository{}, listTaskMappingsUseCase)

	timeEntry := validTimeEntry()
	timeEntry.ProjectLabel = ""

	_, err := createTimeEntryUseCase.Create("owner-id", timeEntry)

	var domainError *sharedException.DomainError
	if !errors.As(err, &domainError) || domainError.Code != 400 {
		t.Fatalf("expected a 400 validation error since nothing matched and project label stayed blank, got %v", err)
	}
}
