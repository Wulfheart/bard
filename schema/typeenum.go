package schema

type TypeEnum string
const (
	Artifact TypeEnum = "artifact"
	Bitbucket TypeEnum = "bitbucket"
	Composer TypeEnum = "composer"
	Fossil TypeEnum = "fossil"
	GitBitbucket TypeEnum = "git-bitbucket"
	Github TypeEnum = "github"
	Gitlab TypeEnum = "gitlab"
	Hg TypeEnum = "hg"
	Package TypeEnum = "package"
	Path TypeEnum = "path"
	Pear TypeEnum = "pear"
	Perforce TypeEnum = "perforce"
	Svn TypeEnum = "svn"
	TypeGit TypeEnum = "git"
	Vcs TypeEnum = "vcs"
)
