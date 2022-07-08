package schema

type Repository struct {
	AllowSSLDowngrade        *bool          `json:"allow_ssl_downgrade,omitempty"`        
	Canonical                *bool          `json:"canonical,omitempty"`                  
	Exclude                  []string       `json:"exclude,omitempty"`                    
	ForceLazyProviders       *bool          `json:"force-lazy-providers,omitempty"`       
	Only                     []string       `json:"only,omitempty"`                       
	Options                  *OptionsClass  `json:"options,omitempty"`                    
	Type                     TypeEnum       `json:"type"`                                 
	URL                      *string        `json:"url,omitempty"`                        
	Branch                   *string        `json:"branch,omitempty"`                     
	BranchesPath             *Abandoned     `json:"branches-path"`                        
	Depot                    *string        `json:"depot,omitempty"`                      
	NoAPI                    *bool          `json:"no-api,omitempty"`                     
	P4Password               *string        `json:"p4password,omitempty"`                 
	P4User                   *string        `json:"p4user,omitempty"`                     
	PackagePath              *string        `json:"package-path,omitempty"`               
	SecureHTTP               *bool          `json:"secure-http,omitempty"`                
	SvnCacheCredentials      *bool          `json:"svn-cache-credentials,omitempty"`      
	TagsPath                 *Abandoned     `json:"tags-path"`                            
	TrunkPath                *Abandoned     `json:"trunk-path"`                           
	UniquePerforceClientName *string        `json:"unique_perforce_client_name,omitempty"`
	VendorAlias              *string        `json:"vendor-alias,omitempty"`               
	Package                  *FluffyPackage `json:"package"`                              
}
