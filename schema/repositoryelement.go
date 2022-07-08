package schema

// A set of additional repositories where packages can be found.
type RepositoryElement struct {
	AllowSSLDowngrade        *bool          `json:"allow_ssl_downgrade,omitempty"`  
	Canonical                *bool          `json:"canonical,omitempty"`            
	Exclude                  *Exclude       `json:"exclude"`                        
	ForceLazyProviders       *bool          `json:"force-lazy-providers,omitempty"` 
	Only                     *Exclude       `json:"only"`                           
	Options                  *OptionsUnion  `json:"options"`                        
	Type                     *TypeUnion     `json:"type"`                           
	URL                      *Abandoned     `json:"url"`                            
	Branch                   *Abandoned     `json:"branch"`                         
	BranchesPath             *Abandoned     `json:"branches-path"`                  
	Depot                    *Abandoned     `json:"depot"`                          
	NoAPI                    *bool          `json:"no-api,omitempty"`               
	P4Password               *Abandoned     `json:"p4password"`                     
	P4User                   *Abandoned     `json:"p4user"`                         
	PackagePath              *Abandoned     `json:"package-path"`                   
	SecureHTTP               *bool          `json:"secure-http,omitempty"`          
	SvnCacheCredentials      *bool          `json:"svn-cache-credentials,omitempty"`
	TagsPath                 *Abandoned     `json:"tags-path"`                      
	TrunkPath                *Abandoned     `json:"trunk-path"`                     
	UniquePerforceClientName *Abandoned     `json:"unique_perforce_client_name"`    
	VendorAlias              *Abandoned     `json:"vendor-alias"`                   
	Package                  *PurplePackage `json:"package"`                        
}
