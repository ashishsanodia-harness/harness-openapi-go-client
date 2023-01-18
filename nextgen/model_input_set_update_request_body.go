/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * The Harness Software Delivery Platform uses OpenAPI Specification v3.0. Harness constantly improves these APIs. Please be aware that some improvements could cause breaking changes. # Introduction       The Harness API allows you to integrate and use all the services and modules we provide on the Harness Platform. If you use client-side SDKs, Harness functionality can be integrated with your client-side automation, helping you reduce manual efforts and deploy code faster.      For more information about how Harness works, read our [documentation](https://developer.harness.io/docs/getting-started) or visit the [Harness Developer Hub](https://developer.harness.io/).    ## How it works      The Harness API is a RESTful API that uses standard HTTP verbs. You can send requests in JSON, YAML, or form-data format. The format of the response matches the format of your request. You must send a single request at a time and ensure that you include your authentication key. For more information about this, go to [Authentication](#section/Introduction/Authentication).  ## Get started    Before you start integrating, get to know our API better by reading the following topics:      * [Harness key concepts](https://developer.harness.io/docs/getting-started/learn-harness-key-concepts/)   * [Authentication](#section/Introduction/Authentication)   * [Requests and responses](#section/Introduction/Requests-and-Responses)   * [Errors](#tag/Error-Response)   * [Versioning](#section/Introduction/Versioning)   * [Pagination](/#section/Introduction/Pagination)      The methods you need to integrate with depend on the functionality you want to use. Work with  your Harness Solutions Engineer to determine which methods you need.  ## Authentication  To authenticate with the Harness API, you need to:   1. Generate an API token on the Harness Platform.   2. Send the API token you generate in the `x-api-key field` in each request.  ### Generate an API token  To generate an API token, complete the following steps:   1. Go to the [Harness Platform](app.harness.io).   2. On the left-hand navigation, click **My Profile**.   3. Click **+API Key**, enter a name for your key and then click **Save**.   4. Within the API Key tile, click **+Token**.   5. Enter a name for your token and click **Generate Token**. **Important**: Make sure to save your token securely. Harness does not store the API token for future reference, so make sure to save your token securely before you leave the page.  ### Send the API token in your requests  Send the token you created in the Harness Platform in the x-api-key parameter. For example:   `x-api-key: YOUR_API_KEY_HERE`  ## Requests and Responses      The structure for each request and response is outlined in the API documentation. We have examples in JSON and YAML for every request and response. You can use our online editor to test the examples.          ## Versioning      The current version of our API is **1.0**.      The version number represents the core API and does not change frequently. The version number changes only if there is a significant departure from the basic underpinnings of the existing API. For example, the version number would change with a system-wide refactoring of core concepts or resources.   This version number is represented within the API path. For example, the v1 in `api.harness.io/v1/` is the version number.    ## Pagination      We use pagination to place limits on the number of responses associated with list endpoints. Pagination is achieved by the use of limit query parameters. The limit defaults to 30. Its maximum value is 100.      Link headers are used, which are useful for moving between pages of the overall list.  For example:      ``` Link: </v1/roles?page=2&limit=30>; rel=\"next\", </v1/roles?page=1&limit=30>; rel=\"self\", </v1/roles?page=0&limit=30>; rel=\"previous\",   ```
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Input Set update request body
type InputSetUpdateRequestBody struct {
	// Input Set YAML (to be passed as a String).
	InputSetYaml string `json:"input_set_yaml"`
	// Input Set identifier
	Identifier string `json:"identifier"`
	// Input Set name
	Name string `json:"name"`
	// Input Set description
	Description string `json:"description,omitempty"`
	// Input Set tags
	Tags       map[string]string         `json:"tags,omitempty"`
	GitDetails *InputSetGitUpdateDetails `json:"git_details,omitempty"`
}
