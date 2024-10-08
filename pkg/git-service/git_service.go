package gitservice

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-logr/logr"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GitService struct {
	GitURL      string
	reference   string
	secretValue string
	gitType     GitProvider
	owner       string
	repo        string
	logger      logr.Logger
	status      metav1.ConditionStatus
	reason      GitConditionReason
}

func New(gitURL, branch, secretValue string, logger logr.Logger) *GitService {
	var status metav1.ConditionStatus = metav1.ConditionUnknown
	var reason GitConditionReason = ReasonProcessing
	gitType := identifyGitType(gitURL)
	if gitType == Unknown {
		logger.Error(errors.New(ReasonUnsupportedGitType.String()), "Unsupported Git Type")
		return &GitService{
			status: metav1.ConditionFalse,
			reason: ReasonUnsupportedGitType,
		}
	}
	owner, repo, err := getOwnerAndRepo(gitURL)
	if err != nil {
		logger.Error(err, "Cannot get owner and repo from gitURL")
		return &GitService{
			status: metav1.ConditionFalse,
			reason: ReasonInvalidGitURL,
		}
	}

	return &GitService{
		GitURL:      gitURL,
		reference:   branch,
		gitType:     gitType,
		owner:       owner,
		repo:        repo,
		secretValue: secretValue,
		logger:      logger,
		status:      status,
		reason:      reason,
	}
}

func (g *GitService) IsRepoReachable() (metav1.ConditionStatus, GitConditionReason) {
	if g.status != metav1.ConditionUnknown {
		return g.status, g.reason
	}
	switch g.gitType {
	case Github:
		isGHRepoReachable(g)
	case Gitlab:
		isGLRepoReachable(g)
	}
	return g.status, g.reason
}

func identifyGitType(gitURL string) GitProvider {
	commonRegex := regexp.
		MustCompile(`^(https?://)?(www\.)?(github\.com|gitlab\.com)(:[0-9]{1,5})?\/([^\/]+)\/([^\/]+)(\.git)?\/?$`)
	if !commonRegex.MatchString(gitURL) {
		return Unknown
	}

	switch {
	case strings.Contains(gitURL, "github.com"):
		return Github
	case strings.Contains(gitURL, "gitlab.com"):
		return Gitlab
	default:
		return Unknown
	}
}

func getOwnerAndRepo(gitURL string) (string, string, error) {
	// This regular expression matches URLs of the form "https://github.com/username/repo"
	// or "https://gitlab.com/username/repo", with or without the "https://" prefix,
	// and optionally with ".git" at the end.
	re := regexp.
		MustCompile(`^(https?://)?(www\.)?(github\.com|gitlab\.com)(:[0-9]{1,5})?\/([^\/]+)\/([^\/]+)(\.git)?\/?$`)
	matches := re.FindStringSubmatch(gitURL)

	if len(matches) < 7 {
		return "", "", errors.New(ReasonInvalidGitURL.String())
	}

	username, repo := matches[5], matches[6]
	return username, strings.TrimSuffix(repo, ".git"), nil
}

/** Run this main function to test this package

func main() {
	logger := zap.New(zap.UseDevMode(true), zap.StacktraceLevel(zapcore.DPanicLevel))
	log.SetLogger(logger)

	ctx := context.Background()
	logger = log.FromContext(ctx)
	// g := New("https://github.com/openshift-console/console-application-operator", "main", "", logger)
	g := New("https://gitlab.com/avikkundu/oc-pipe", "main", "<PAT>", logger)
	fmt.Println(g.isRepoReachable())

}

*/
