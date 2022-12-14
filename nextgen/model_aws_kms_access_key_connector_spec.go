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

// This contains details of the AWS and needs AWS encrypted access and secret keys for the AWS KMS.
type AwsKmsAccessKeyConnectorSpec struct {
	// This specifies the type of connector
	Type_ string `json:"type"`
	// Amazon Resource Name (ARN)
	KmsArn string `json:"kms_arn"`
	// AWS Region for kms
	Region string `json:"region"`
	// Boolean value to indicate if the Secret Manager is your default Secret Manager
	Default_ bool `json:"default,omitempty"`
	// Access Key for AWS authentication
	AccessKey string `json:"access_key"`
	// Secret Key for AWS authentication
	SecretKey string `json:"secret_key"`
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager
	DelegateSelectors []string `json:"delegate_selectors,omitempty"`
}
