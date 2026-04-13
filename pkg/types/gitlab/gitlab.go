// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    gitlabCI, err := UnmarshalGitlabCI(bytes)
//    bytes, err = gitlabCI.Marshal()

package gitlab

import "bytes"
import "errors"

import "encoding/json"

type GitlabCI map[string]*GitlabCIValue

func UnmarshalGitlabCI(data []byte) (GitlabCI, error) {
	var r GitlabCI
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GitlabCI) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CacheItem struct {
	FallbackKeys                                                                                 []string                  `json:"fallback_keys,omitempty"`
	Key                                                                                          *KeyUnion                 `json:"key"`
	Paths                                                                                        []string                  `json:"paths,omitempty"`
	Policy                                                                                       *string                   `json:"policy,omitempty"`
	Unprotect                                                                                    *bool                     `json:"unprotect,omitempty"`
	Untracked                                                                                    *bool                     `json:"untracked,omitempty"`
	When                                                                                         *CacheItemWhen            `json:"when,omitempty"`
	Inputs                                                                                       *InputsClass              `json:"inputs,omitempty"`
	// Relative path from local repository root (`/`) to the `yaml`/`yml` file template. The                               
	// file must be on the same branch, and does not work across git submodules.                                           
	Local                                                                                        *string                   `json:"local,omitempty"`
	Rules                                                                                        []IncludeRuleElement      `json:"rules"`
	File                                                                                         *File                     `json:"file"`
	// Path to the project, e.g. `group/project`, or `group/sub-group/project` [Learn                                      
	// more](https://docs.gitlab.com/ci/yaml/#includeproject).                                                             
	Project                                                                                      *string                   `json:"project,omitempty"`
	// Branch/Tag/Commit-hash for the target project.                                                                      
	Ref                                                                                          *string                   `json:"ref,omitempty"`
	// Use a `.gitlab-ci.yml` template as a base, e.g. `Nodejs.gitlab-ci.yml`.                                             
	Template                                                                                     *string                   `json:"template,omitempty"`
	// Local path to component directory or full path to external component directory.                                     
	Component                                                                                    *string                   `json:"component,omitempty"`
	// SHA256 integrity hash of the remote file content.                                                                   
	Integrity                                                                                    *string                   `json:"integrity,omitempty"`
	// URL to a `yaml`/`yml` template file using HTTP/HTTPS.                                                               
	Remote                                                                                       *string                   `json:"remote,omitempty"`
	Alias                                                                                        *string                   `json:"alias,omitempty"`
	Command                                                                                      *Script                   `json:"command"`
	Docker                                                                                       *Docker                   `json:"docker,omitempty"`
	Entrypoint                                                                                   []string                  `json:"entrypoint,omitempty"`
	Kubernetes                                                                                   *Kubernetes               `json:"kubernetes,omitempty"`
	// Full name of the image that should be used. It should contain the Registry part if needed.                          
	Name                                                                                         *string                   `json:"name,omitempty"`
	PullPolicy                                                                                   *PullPolicyUnion          `json:"pull_policy"`
	Variables                                                                                    map[string]*VariableValue `json:"variables,omitempty"`
}

type Docker struct {
	// Image architecture to pull.                      
	Platform                                    *string `json:"platform,omitempty"`
	// Username or UID to use for the container.        
	User                                        *string `json:"user,omitempty"`
}

type InputsClass struct {
}

type KeyClass struct {
	Files        []string `json:"files,omitempty"`
	FilesCommits []string `json:"files_commits,omitempty"`
	Prefix       *string  `json:"prefix,omitempty"`
}

type Kubernetes struct {
	// Username or UID to use for the container. It also supports the UID:GID format.      
	User                                                                             *User `json:"user"`
}

type IncludeRuleClass struct {
	Changes *Changes         `json:"changes"`
	Exists  *Exists          `json:"exists"`
	If      *string          `json:"if,omitempty"`
	When    *IncludeRuleWhen `json:"when"`
}

type ChangesClass struct {
	// Ref for comparing changes.         
	CompareTo                    *string  `json:"compare_to,omitempty"`
	// List of file paths.                
	Paths                        []string `json:"paths"`
}

type ExistsClass struct {
	// List of file paths.                       
	Paths                               []string `json:"paths"`
	// Path of the project to search in.         
	Project                             *string  `json:"project,omitempty"`
	// Ref of the project to search in.          
	Ref                                 *string  `json:"ref,omitempty"`
}

type VariableClass struct {
	Expand *bool   `json:"expand,omitempty"`
	Value  *string `json:"value,omitempty"`
}

// Used to configure the kubernetes deployment for this environment. This is currently not
// supported for kubernetes clusters that are managed by GitLab.
//
// Specify what to forward to the downstream pipeline.
type CacheItemClass struct {
	// Path to JSON file with accessibility report.                                                                       
	Accessibility                                                                                *string                  `json:"accessibility,omitempty"`
	// Path to JSON file with annotations report.                                                                         
	Annotations                                                                                  *string                  `json:"annotations,omitempty"`
	// Path to a single file with browser performance metric report(s).                                                   
	BrowserPerformance                                                                           *string                  `json:"browser_performance,omitempty"`
	// Path to file or list of files with code quality report(s) (such as Code Climate).                                  
	Codequality                                                                                  *StringFileList          `json:"codequality"`
	// Path to file or list of files with Container scanning vulnerabilities report(s).                                   
	ContainerScanning                                                                            *StringFileList          `json:"container_scanning"`
	// Used to collect coverage reports from the job.                                                                     
	CoverageReport                                                                               *CoverageReport          `json:"coverage_report"`
	Cyclonedx                                                                                    *StringFileList          `json:"cyclonedx"`
	// Path to file or list of files with DAST vulnerabilities report(s).                                                 
	Dast                                                                                         *StringFileList          `json:"dast"`
	// Path to file or list of files with Dependency scanning vulnerabilities report(s).                                  
	DependencyScanning                                                                           *StringFileList          `json:"dependency_scanning"`
	// Path to file or list of files containing runtime-created variables for this job.                                   
	Dotenv                                                                                       *StringFileList          `json:"dotenv"`
	// Path for file(s) that should be parsed as JUnit XML result                                                         
	Junit                                                                                        *StringFileList          `json:"junit"`
	// Deprecated in 12.8: Path to file or list of files with license report(s).                                          
	LicenseManagement                                                                            *StringFileList          `json:"license_management"`
	// Path to file or list of files with license report(s).                                                              
	LicenseScanning                                                                              *StringFileList          `json:"license_scanning"`
	LoadPerformance                                                                              *StringFileList          `json:"load_performance"`
	// Path to file or list of files containing code intelligence (Language Server Index Format).                         
	Lsif                                                                                         *StringFileList          `json:"lsif"`
	// Path to file or list of files with custom metrics report(s).                                                       
	Metrics                                                                                      *StringFileList          `json:"metrics"`
	// Path to file or list of files with Repository X-Ray report(s).                                                     
	RepositoryXray                                                                               *StringFileList          `json:"repository_xray"`
	// Path to file or list of files with requirements report(s).                                                         
	Requirements                                                                                 *StringFileList          `json:"requirements"`
	Sarif                                                                                        *StringFileList          `json:"sarif"`
	// Path to file or list of files with SAST vulnerabilities report(s).                                                 
	Sast                                                                                         *StringFileList          `json:"sast"`
	// Path to file or list of files with secret detection report(s).                                                     
	SecretDetection                                                                              *StringFileList          `json:"secret_detection"`
	// Path to file or list of files with terraform plan(s).                                                              
	Terraform                                                                                    *StringFileList          `json:"terraform"`
	Files                                                                                        []string                 `json:"files,omitempty"`
	FilesCommits                                                                                 []string                 `json:"files_commits,omitempty"`
	Prefix                                                                                       *string                  `json:"prefix,omitempty"`
	// Image architecture to pull.                                                                                        
	Platform                                                                                     *string                  `json:"platform,omitempty"`
	// Username or UID to use for the container.                                                                          
	//                                                                                                                    
	// Username or UID to use for the container. It also supports the UID:GID format.                                     
	User                                                                                         *User                    `json:"user"`
	// Specifies the GitLab Agent for Kubernetes. The format is                                                           
	// `path/to/agent/project:agent-name`.                                                                                
	Agent                                                                                        *string                  `json:"agent,omitempty"`
	// Used to configure the dashboard for this environment.                                                              
	Dashboard                                                                                    *Dashboard               `json:"dashboard,omitempty"`
	// Deprecated. Use `dashboard.flux_resource_path` instead. The Flux resource path to                                  
	// associate with this environment. This must be the full resource path. For example,                                 
	// 'helm.toolkit.fluxcd.io/v2/namespaces/gitlab-agent/helmreleases/gitlab-agent'.                                     
	FluxResourcePath                                                                             *string                  `json:"flux_resource_path,omitempty"`
	// Used to configure the managed resources for this environment.                                                      
	ManagedResources                                                                             *ManagedResources        `json:"managed_resources,omitempty"`
	// Deprecated. Use `dashboard.namespace` instead. The kubernetes namespace where this                                 
	// environment's dashboard should be deployed to.                                                                     
	Namespace                                                                                    *string                  `json:"namespace,omitempty"`
	// Include asset links in the release.                                                                                
	Links                                                                                        []Link                   `json:"links,omitempty"`
	// Variables added for manual pipeline runs and scheduled pipelines are passed to downstream                          
	// pipelines.                                                                                                         
	PipelineVariables                                                                            *bool                    `json:"pipeline_variables,omitempty"`
	// Variables defined in the trigger job are passed to downstream pipelines.                                           
	YAMLVariables                                                                                *bool                    `json:"yaml_variables,omitempty"`
	Aud                                                                                          *StringFileList          `json:"aud"`
	Default                                                                                      interface{}              `json:"default"`
	Description                                                                                  *string                  `json:"description,omitempty"`
	Options                                                                                      []Option                 `json:"options,omitempty"`
	Regex                                                                                        *string                  `json:"regex,omitempty"`
	Type                                                                                         *Type                    `json:"type,omitempty"`
	AwsSecretsManager                                                                            *AwsSecretsManagerUnion  `json:"aws_secrets_manager"`
	AzureKeyVault                                                                                *AzureKeyVault           `json:"azure_key_vault,omitempty"`
	File                                                                                         *bool                    `json:"file,omitempty"`
	GcpSecretManager                                                                             *GcpSecretManager        `json:"gcp_secret_manager,omitempty"`
	GitlabSecretsManager                                                                         *GitlabSecretsManager    `json:"gitlab_secrets_manager,omitempty"`
	// Specifies the JWT variable that should be used to authenticate with the secret provider.                           
	Token                                                                                        *string                  `json:"token,omitempty"`
	Vault                                                                                        *VaultUnion              `json:"vault"`
	Expand                                                                                       *bool                    `json:"expand,omitempty"`
	Value                                                                                        *string                  `json:"value,omitempty"`
	Rules                                                                                        []map[string]interface{} `json:"rules,omitempty"`
}

type AwsSecretsManagerClass struct {
	// The name of the field to retrieve from the secret. If not specified, the entire secret is        
	// retrieved.                                                                                       
	Field                                                                                       *string `json:"field,omitempty"`
	// The AWS region where the secret is stored. Use this to override the region for a specific        
	// secret. Defaults to AWS_REGION variable.                                                         
	Region                                                                                      *string `json:"region,omitempty"`
	// The ARN of the IAM role to assume before retrieving the secret. Use this to override the         
	// ARN. Defaults to AWS_ROLE_ARN variable.                                                          
	RoleArn                                                                                     *string `json:"role_arn,omitempty"`
	// The name of the session to use when assuming the role. Use this to override the session          
	// name. Defaults to AWS_ROLE_SESSION_NAME variable.                                                
	RoleSessionName                                                                             *string `json:"role_session_name,omitempty"`
	// The ARN or name of the secret to retrieve. To retrieve a secret from another account, you        
	// must use an ARN.                                                                                 
	SecretID                                                                                    string  `json:"secret_id"`
	// The unique identifier of the version of the secret to retrieve. If you include both this         
	// parameter and VersionStage, the two parameters must refer to the same secret version. If         
	// you don't specify either a VersionStage or VersionId, Secrets Manager returns the                
	// AWSCURRENT version.                                                                              
	VersionID                                                                                   *string `json:"version_id,omitempty"`
	// The staging label of the version of the secret to retrieve. If you include both this             
	// parameter and VersionStage, the two parameters must refer to the same secret version. If         
	// you don't specify either a VersionStage or VersionId, Secrets Manager returns the                
	// AWSCURRENT version.                                                                              
	VersionStage                                                                                *string `json:"version_stage,omitempty"`
}

type AzureKeyVault struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

type CoverageReport struct {
	// Code coverage format used by the test framework.                       
	CoverageFormat                                            *CoverageFormat `json:"coverage_format,omitempty"`
	// Path to the coverage report file that should be parsed.                
	Path                                                      *string         `json:"path,omitempty"`
}

// Used to configure the dashboard for this environment.
type Dashboard struct {
	// The Flux resource path to associate with this environment. This must be the full resource        
	// path. For example,                                                                               
	// 'helm.toolkit.fluxcd.io/v2/namespaces/gitlab-agent/helmreleases/gitlab-agent'.                   
	FluxResourcePath                                                                            *string `json:"flux_resource_path,omitempty"`
	// The kubernetes namespace where the dashboard for this environment should be deployed to.         
	Namespace                                                                                   *string `json:"namespace,omitempty"`
}

type GcpSecretManager struct {
	Name    string   `json:"name"`
	Version *Version `json:"version"`
}

type GitlabSecretsManager struct {
	Name                                                                                        string  `json:"name"`
	// Source of the secret. Defaults to the current project if not given. For fetching a secret        
	// from a group, provide group/<full_path_of_the_group>                                             
	Source                                                                                      *string `json:"source,omitempty"`
}

type Link struct {
	// The redirect link to the url.                                 
	Filepath                                               *string   `json:"filepath,omitempty"`
	// The content kind of what users can download via url.          
	LinkType                                               *LinkType `json:"link_type,omitempty"`
	// The name of the link.                                         
	Name                                                   string    `json:"name"`
	// The URL to download a file.                                   
	URL                                                    string    `json:"url"`
}

// Used to configure the managed resources for this environment.
type ManagedResources struct {
	// Indicates whether the managed resources are enabled for this environment.      
	Enabled                                                                     *bool `json:"enabled,omitempty"`
}

type VaultClass struct {
	Engine Engine `json:"engine"`
	Field  string `json:"field"`
	Path   string `json:"path"`
}

type Engine struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type PullPolicyElement string

const (
	IfNotPresent     PullPolicyElement = "if-not-present"
	PullPolicyAlways PullPolicyElement = "always"
	PullPolicyNever  PullPolicyElement = "never"
)

type IncludeRuleWhen string

const (
	PurpleAlways IncludeRuleWhen = "always"
	WhenNever    IncludeRuleWhen = "never"
)

type CacheItemWhen string

const (
	FluffyAlways CacheItemWhen = "always"
	OnFailure    CacheItemWhen = "on_failure"
	OnSuccess    CacheItemWhen = "on_success"
)

// Code coverage format used by the test framework.
type CoverageFormat string

const (
	Cobertura CoverageFormat = "cobertura"
	Jacoco    CoverageFormat = "jacoco"
)

// The content kind of what users can download via url.
type LinkType string

const (
	Image   LinkType = "image"
	Other   LinkType = "other"
	Package LinkType = "package"
	Runbook LinkType = "runbook"
)

type Type string

const (
	Array   Type = "array"
	Boolean Type = "boolean"
	Number  Type = "number"
	String  Type = "string"
)

type GitlabCIValue struct {
	String     *string
	UnionArray []ReferenceElement
	UnionMap   map[string]*JobValue
}

func (x *GitlabCIValue) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	x.UnionMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, false, nil, true, &x.UnionMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *GitlabCIValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, false, nil, x.UnionMap != nil, x.UnionMap, false, nil, false)
}

type ReferenceElement struct {
	CacheItem   *CacheItem
	String      *string
	StringArray []string
}

func (x *ReferenceElement) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.CacheItem = nil
	var c CacheItem
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CacheItem = &c
	}
	return nil
}

func (x *ReferenceElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, x.CacheItem != nil, x.CacheItem, false, nil, false, nil, false)
}

type Script struct {
	String     *string
	UnionArray []StringFileList
}

func (x *Script) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Script) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}

// Path to file or list of files with code quality report(s) (such as Code Climate).
//
// Path to file or list of files with Container scanning vulnerabilities report(s).
//
// Path to file or list of files with DAST vulnerabilities report(s).
//
// Path to file or list of files with Dependency scanning vulnerabilities report(s).
//
// Path to file or list of files containing runtime-created variables for this job.
//
// Deprecated in 12.8: Path to file or list of files with license report(s).
//
// Path to file or list of files with license report(s).
//
// Path to file or list of files containing code intelligence (Language Server Index
// Format).
//
// Path to file or list of files with custom metrics report(s).
//
// Path to file or list of files with Repository X-Ray report(s).
//
// Path to file or list of files with requirements report(s).
//
// Path to file or list of files with SAST vulnerabilities report(s).
//
// Path to file or list of files with secret detection report(s).
//
// Path to file or list of files with terraform plan(s).
//
// Path for file(s) that should be parsed as JUnit XML result
//
// The name of one or more jobs to inherit configuration from.
type StringFileList struct {
	String      *string
	StringArray []string
}

func (x *StringFileList) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *StringFileList) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type File struct {
	String      *string
	StringArray []string
}

func (x *File) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *File) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}

type KeyUnion struct {
	KeyClass *KeyClass
	String   *string
}

func (x *KeyUnion) UnmarshalJSON(data []byte) error {
	x.KeyClass = nil
	var c KeyClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.KeyClass = &c
	}
	return nil
}

func (x *KeyUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.KeyClass != nil, x.KeyClass, false, nil, false, nil, false)
}

// Username or UID to use for the container. It also supports the UID:GID format.
type User struct {
	Integer *int64
	String  *string
}

func (x *User) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *User) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type PullPolicyUnion struct {
	Enum      *PullPolicyElement
	EnumArray []PullPolicyElement
}

func (x *PullPolicyUnion) UnmarshalJSON(data []byte) error {
	x.EnumArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.EnumArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *PullPolicyUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.EnumArray != nil, x.EnumArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

type IncludeRuleElement struct {
	IncludeRuleClass *IncludeRuleClass
	String           *string
	StringArray      []string
}

func (x *IncludeRuleElement) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.IncludeRuleClass = nil
	var c IncludeRuleClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.IncludeRuleClass = &c
	}
	return nil
}

func (x *IncludeRuleElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, x.IncludeRuleClass != nil, x.IncludeRuleClass, false, nil, false, nil, false)
}

type Changes struct {
	ChangesClass *ChangesClass
	StringArray  []string
}

func (x *Changes) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.ChangesClass = nil
	var c ChangesClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ChangesClass = &c
	}
	return nil
}

func (x *Changes) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, x.ChangesClass != nil, x.ChangesClass, false, nil, false, nil, false)
}

type Exists struct {
	ExistsClass *ExistsClass
	StringArray []string
}

func (x *Exists) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.ExistsClass = nil
	var c ExistsClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ExistsClass = &c
	}
	return nil
}

func (x *Exists) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, x.ExistsClass != nil, x.ExistsClass, false, nil, false, nil, false)
}

type VariableValue struct {
	Bool          *bool
	Double        *float64
	String        *string
	VariableClass *VariableClass
}

func (x *VariableValue) UnmarshalJSON(data []byte) error {
	x.VariableClass = nil
	var c VariableClass
	object, err := unmarshalUnion(data, nil, &x.Double, &x.Bool, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.VariableClass = &c
	}
	return nil
}

func (x *VariableValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, x.Bool, x.String, false, nil, x.VariableClass != nil, x.VariableClass, false, nil, false, nil, false)
}

type JobValue struct {
	AnythingArray []interface{}
	Bool          *bool
	Double        *float64
	Integer       *int64
	String        *string
	UnionMap      map[string]*CacheItemCacheItem
}

func (x *JobValue) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.UnionMap = nil
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, &x.Bool, &x.String, true, &x.AnythingArray, false, nil, true, &x.UnionMap, false, nil, true)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *JobValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, false, nil, x.UnionMap != nil, x.UnionMap, false, nil, true)
}

type CacheItemCacheItem struct {
	AnythingArray  []interface{}
	Bool           *bool
	CacheItemClass *CacheItemClass
	Double         *float64
	Integer        *int64
	String         *string
}

func (x *CacheItemCacheItem) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.CacheItemClass = nil
	var c CacheItemClass
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, &x.Bool, &x.String, true, &x.AnythingArray, true, &c, false, nil, false, nil, true)
	if err != nil {
		return err
	}
	if object {
		x.CacheItemClass = &c
	}
	return nil
}

func (x *CacheItemCacheItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, x.CacheItemClass != nil, x.CacheItemClass, false, nil, false, nil, true)
}

type AwsSecretsManagerUnion struct {
	AwsSecretsManagerClass *AwsSecretsManagerClass
	String                 *string
}

func (x *AwsSecretsManagerUnion) UnmarshalJSON(data []byte) error {
	x.AwsSecretsManagerClass = nil
	var c AwsSecretsManagerClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.AwsSecretsManagerClass = &c
	}
	return nil
}

func (x *AwsSecretsManagerUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.AwsSecretsManagerClass != nil, x.AwsSecretsManagerClass, false, nil, false, nil, false)
}

type Version struct {
	Integer *int64
	String  *string
}

func (x *Version) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Version) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type Option struct {
	Bool   *bool
	Double *float64
	String *string
}

func (x *Option) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, &x.Double, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Option) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type VaultUnion struct {
	String     *string
	VaultClass *VaultClass
}

func (x *VaultUnion) UnmarshalJSON(data []byte) error {
	x.VaultClass = nil
	var c VaultClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.VaultClass = &c
	}
	return nil
}

func (x *VaultUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.VaultClass != nil, x.VaultClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
			*pi = nil
	}
	if pf != nil {
			*pf = nil
	}
	if pb != nil {
			*pb = nil
	}
	if ps != nil {
			*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
			return false, err
	}

	switch v := tok.(type) {
	case json.Number:
			if pi != nil {
					i, err := v.Int64()
					if err == nil {
							*pi = &i
							return false, nil
					}
			}
			if pf != nil {
					f, err := v.Float64()
					if err == nil {
							*pf = &f
							return false, nil
					}
					return false, errors.New("Unparsable number")
			}
			return false, errors.New("Union does not contain number")
	case float64:
			return false, errors.New("Decoder should not return float64")
	case bool:
			if pb != nil {
					*pb = &v
					return false, nil
			}
			return false, errors.New("Union does not contain bool")
	case string:
			if haveEnum {
					return false, json.Unmarshal(data, pe)
			}
			if ps != nil {
					*ps = &v
					return false, nil
			}
			return false, errors.New("Union does not contain string")
	case nil:
			if nullable {
					return false, nil
			}
			return false, errors.New("Union does not contain null")
	case json.Delim:
			if v == '{' {
					if haveObject {
							return true, json.Unmarshal(data, pc)
					}
					if haveMap {
							return false, json.Unmarshal(data, pm)
					}
					return false, errors.New("Union does not contain object")
			}
			if v == '[' {
					if haveArray {
							return false, json.Unmarshal(data, pa)
					}
					return false, errors.New("Union does not contain array")
			}
			return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")
}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
			return json.Marshal(*pi)
	}
	if pf != nil {
			return json.Marshal(*pf)
	}
	if pb != nil {
			return json.Marshal(*pb)
	}
	if ps != nil {
			return json.Marshal(*ps)
	}
	if haveArray {
			return json.Marshal(pa)
	}
	if haveObject {
			return json.Marshal(pc)
	}
	if haveMap {
			return json.Marshal(pm)
	}
	if haveEnum {
			return json.Marshal(pe)
	}
	if nullable {
			return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
