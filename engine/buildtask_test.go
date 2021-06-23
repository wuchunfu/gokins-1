package engine

import (
	"context"
	"fmt"
	"github.com/gokins-main/core/runtime"
	"testing"
)

func TestGitClone(t *testing.T) {
	task := BuildTask{
		ctx: context.TODO(),
		build: &runtime.Build{
			Id: "1231",
			Repo: &runtime.Repository{
				Name:     "303600370@qq.com",
				Token:    "jimbir6666memory",
				Sha:      "c202ee042db1fc8b8c16c6c968195cec6185d7db",
				CloneURL: "https://gitee.com/SuperHeroJim/gokins-test",
			},
		},
	}
	err := task.getRepo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(task.repoPath)
	fmt.Println(task.isClone)
}
