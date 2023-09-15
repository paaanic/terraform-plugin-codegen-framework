// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

type AssocExtType struct {
	*schema.AssociatedExternalType
}

func NewAssocExtType(assocExtType *schema.AssociatedExternalType) *AssocExtType {
	if assocExtType == nil {
		return nil
	}

	return &AssocExtType{
		AssociatedExternalType: assocExtType,
	}
}

func (a *AssocExtType) Imports() *Imports {
	imports := NewImports()

	if a == nil {
		return imports
	}

	if a.AssociatedExternalType.Import == nil {
		return imports
	}

	if len(a.AssociatedExternalType.Import.Path) > 0 {
		imports.Add(*a.AssociatedExternalType.Import)

		imports.Add(code.Import{
			Path: BaseTypesImport,
		})
	}

	return imports
}

func (a *AssocExtType) Type() string {
	if a == nil {
		return ""
	}

	return a.AssociatedExternalType.Type
}

func (a *AssocExtType) TypeReference() string {
	if a == nil {
		return ""
	}

	tr, _ := strings.CutPrefix(a.AssociatedExternalType.Type, "*")

	return tr
}

func (a *AssocExtType) Equal(other *AssocExtType) bool {
	if a == nil && other == nil {
		return true
	}

	if a == nil || other == nil {
		return false
	}

	return a.AssociatedExternalType.Equal(other.AssociatedExternalType)
}