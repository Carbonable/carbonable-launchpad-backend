// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"github.com/carbonable/carbonable-launchpad-backend/ent"
)

// Aggregated project data.
type ProjectDetails struct {
	Project   *ent.Project   `json:"project"`
	Mint      *ent.Mint      `json:"mint"`
	Launchpad *ent.Launchpad `json:"launchpad"`
}
