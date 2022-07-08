package schema

// Script listeners that will be executed before/after some events.
type Scripts struct {
	PostAutoloadDump       *Comment `json:"post-autoload-dump"`       // Occurs after the autoloader is dumped, contains one or more Class::method callables or; shell commands.
	PostCreateProjectCmd   *Comment `json:"post-create-project-cmd"`  // Occurs after the create-project command is executed, contains one or more Class::method; callables or shell commands.
	PostInstallCmd         *Comment `json:"post-install-cmd"`         // Occurs after the install command is executed, contains one or more Class::method; callables or shell commands.
	PostPackageInstall     *Comment `json:"post-package-install"`     // Occurs after a package is installed, contains one or more Class::method callables or; shell commands.
	PostPackageUninstall   *Comment `json:"post-package-uninstall"`   // Occurs after a package has been uninstalled, contains one or more Class::method callables; or shell commands.
	PostPackageUpdate      *Comment `json:"post-package-update"`      // Occurs after a package is updated, contains one or more Class::method callables or shell; commands.
	PostRootPackageInstall *Comment `json:"post-root-package-install"`// Occurs after the root-package is installed, contains one or more Class::method callables; or shell commands.
	PostStatusCmd          *Comment `json:"post-status-cmd"`          // Occurs after the status command is executed, contains one or more Class::method callables; or shell commands.
	PostUpdateCmd          *Comment `json:"post-update-cmd"`          // Occurs after the update command is executed, contains one or more Class::method callables; or shell commands.
	PreAutoloadDump        *Comment `json:"pre-autoload-dump"`        // Occurs before the autoloader is dumped, contains one or more Class::method callables or; shell commands.
	PreInstallCmd          *Comment `json:"pre-install-cmd"`          // Occurs before the install command is executed, contains one or more Class::method; callables or shell commands.
	PrePackageInstall      *Comment `json:"pre-package-install"`      // Occurs before a package is installed, contains one or more Class::method callables or; shell commands.
	PrePackageUninstall    *Comment `json:"pre-package-uninstall"`    // Occurs before a package has been uninstalled, contains one or more Class::method; callables or shell commands.
	PrePackageUpdate       *Comment `json:"pre-package-update"`       // Occurs before a package is updated, contains one or more Class::method callables or shell; commands.
	PreStatusCmd           *Comment `json:"pre-status-cmd"`           // Occurs before the status command is executed, contains one or more Class::method; callables or shell commands.
	PreUpdateCmd           *Comment `json:"pre-update-cmd"`           // Occurs before the update command is executed, contains one or more Class::method; callables or shell commands.
}
