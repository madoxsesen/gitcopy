package main

func GitCopy(filePath string, gitRepos []string) {
	//TODO check whether git is installed and working

	TraverseFilePath(filePath)

	for _, repo := range gitRepos {
		Info("Checking access of git repo: %s", repo)
		CheckGitRepo(repo)
	}
}
