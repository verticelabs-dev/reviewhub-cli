package git_repo

type RepoInfo struct {
	Owner  string
	Name   string
	Branch string
}

type RepoStoredInfo struct {
	RepoInfo
	Url             string
	ImageName       string
	StorageFileHash string
	StorageFilePath string
}
