// The contents of this file are hand-written to add HandwrittenValidate to select message types

package v1

func (m *CheckPermissionRequest) HandwrittenValidate() error {
	if m.GetResource() != nil && m.GetResource().GetObjectId() == "*" {
		return ObjectReferenceValidationError{
			field:  "ObjectId",
			reason: "alphanumeric value is required",
		}
	}

	return nil
}

func (m *ExpandPermissionTreeRequest) HandwrittenValidate() error {
	if m.GetResource() != nil && m.GetResource().GetObjectId() == "*" {
		return ObjectReferenceValidationError{
			field:  "ObjectId",
			reason: "alphanumeric value is required",
		}
	}

	return nil
}

func (m *Precondition) HandwrittenValidate() error {
	if m.GetFilter() != nil {
		return m.GetFilter().HandwrittenValidate()
	}

	return nil
}

func (m *RelationshipFilter) HandwrittenValidate() error {
	if m.GetOptionalResourceId() == "*" {
		return RelationshipFilterValidationError{
			field:  "OptionalResourceId",
			reason: "alphanumeric value is required",
		}
	}
	return nil
}

func (m *RelationshipUpdate) HandwrittenValidate() error {
	if m.GetRelationship() != nil {
		return m.GetRelationship().HandwrittenValidate()
	}
	return nil
}

func (m *Relationship) HandwrittenValidate() error {
	if m.GetResource() != nil && m.GetResource().GetObjectId() == "*" {
		return ObjectReferenceValidationError{
			field:  "ObjectId",
			reason: "alphanumeric value is required",
		}
	}

	return nil
}

func (m *DeleteRelationshipsRequest) HandwrittenValidate() error {
	if m.GetOptionalPreconditions() != nil {
		for _, precondition := range m.GetOptionalPreconditions() {
			err := precondition.HandwrittenValidate()
			if err != nil {
				return err
			}
		}
	}

	if m.GetRelationshipFilter() != nil {
		return m.GetRelationshipFilter().HandwrittenValidate()
	}

	return nil
}

func (m *WriteRelationshipsRequest) HandwrittenValidate() error {
	if m.GetOptionalPreconditions() != nil {
		for _, precondition := range m.GetOptionalPreconditions() {
			err := precondition.HandwrittenValidate()
			if err != nil {
				return err
			}
		}
	}

	if m.GetUpdates() != nil {
		for _, update := range m.GetUpdates() {
			if update.GetRelationship() != nil {
				err := update.GetRelationship().HandwrittenValidate()
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
