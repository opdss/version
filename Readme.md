# Version 

golang编译时注入版本相关信息

For example:

```shell
   go "-X github.com/opdss/version.buildTimestamp=${TIMESTAMP}\
   -X github.com/opdss/version.buildCommitHash=${COMMIT}\
   -X github.com/opdss/version.buildVersion=${VERSION}\
   -X github.com/opdss/version.buildRelease=${RELEASE}" build  
```