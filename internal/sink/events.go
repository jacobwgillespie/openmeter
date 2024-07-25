package sink

import (
	"errors"

	"github.com/openmeterio/openmeter/internal/event/models"
	"github.com/openmeterio/openmeter/internal/event/spec"
)

const (
	EventSubsystem spec.EventSubsystem = "ingest"
)

const (
	EventIngestion spec.EventName = "ingestion"
)

type IngestEvent struct {
	Namespace  models.NamespaceID `json:"namespace"`
	SubjectKey string             `json:"subjectKey"`

	// MeterSlugs contain the list of slugs that are affected by the event. We
	// should not use meterIDs as they are not something present in the open source
	// version, thus any code that is in opensource should not rely on them.
	MeterSlugs []string `json:"meterSlugs"`
}

var ingestEventSpec = spec.EventTypeSpec{
	Subsystem:   EventSubsystem,
	Name:        EventIngestion,
	SpecVersion: "1.0",
	Version:     "v1",
}

func (i IngestEvent) Spec() *spec.EventTypeSpec {
	return &ingestEventSpec
}

func (i IngestEvent) Validate() error {
	if err := i.Namespace.Validate(); err != nil {
		return err
	}

	if i.SubjectKey == "" {
		return errors.New("subjectKey must be set")
	}

	if len(i.MeterSlugs) == 0 {
		return errors.New("meterSlugs must not be empty")
	}

	return nil
}
