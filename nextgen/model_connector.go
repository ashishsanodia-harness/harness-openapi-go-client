/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type Connector struct {
	// Connector name
	Name string `json:"name"`
	// Connector slug
	Slug string `json:"slug"`
	// Connector description
	Description string `json:"description,omitempty"`
	// Organization slug for connector
	Org string `json:"org,omitempty"`
	// Project slug for connector
	Project string `json:"project,omitempty"`
	// Connector tags
	Tags map[string]string `json:"tags,omitempty"`
	Spec *ConnectorSpec `json:"spec"`
}
