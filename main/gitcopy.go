package main

func GitCopy(filePath string, gitRepos []string) {
	//TODO check whether git is installed and working
	Info("Checking the local file path: %v", filePath)

	files := TraverseFilePath(filePath)
	Info("Found a total of %v files!", len(files))

	for _, repo := range gitRepos {
		Info("Checking access of git repo: %s", repo)
		CheckGitRepo(repo)
	}
}
