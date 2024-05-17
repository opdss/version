package version

import (
	"strconv"
	"strings"
	"time"
)

var (
	//下面变量由编译时注入
	buildTimestamp  string
	buildCommitHash string
	buildVersion    string
	buildRelease    string
	buildBranch     string

	Build build
)

func init() {
	Build = buildInfo()
}

// build info
type build struct {
	Version    string    `json:"version"`   // tag版本号
	CommitHash string    `json:"hash"`      //commit hash
	Timestamp  time.Time `json:"timestamp"` //编译时间
	Release    bool      `json:"release"`   //是否发行版本
	Branch     string    `json:"branch"`    //分支
}

func buildInfo() build {
	timestamp, err := strconv.ParseInt(buildTimestamp, 10, 64)
	if err != nil {
		timestamp = time.Now().Unix()
	}
	return build{
		Timestamp:  time.Unix(timestamp, 0),
		CommitHash: buildCommitHash,
		Version:    buildVersion,
		Branch:     buildBranch,
		Release:    strings.ToLower(buildRelease) == "true",
	}
}
