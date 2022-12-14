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
// ModuleType : Type of Modules
type ModuleType string

// List of ModuleType
const (
	CD_ModuleType ModuleType = "CD"
	CI_ModuleType ModuleType = "CI"
	CV_ModuleType ModuleType = "CV"
	CF_ModuleType ModuleType = "CF"
	CE_ModuleType ModuleType = "CE"
	STO_ModuleType ModuleType = "STO"
	CORE_ModuleType ModuleType = "CORE"
	PMS_ModuleType ModuleType = "PMS"
	TEMPLATESERVICE_ModuleType ModuleType = "TEMPLATESERVICE"
	GOVERNANCE_ModuleType ModuleType = "GOVERNANCE"
	CHAOS_ModuleType ModuleType = "CHAOS"
)
