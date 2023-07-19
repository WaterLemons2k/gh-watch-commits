# watch-commits

✨ 用于订阅仓库提交的 GitHub（`gh`）CLI 扩展。

[English](README.md) | 简体中文

## 动机

GitHub 一直缺少订阅仓库提交的方法，但 https://stackoverflow.com/a/42600376 提供了另一种解决方案：通过拉取请求。

此扩展也是该解决方案的 CLI 版本。

## 安装

```sh
gh extension install WaterLemons2k/gh-watch-commits
```

## 使用

运行：

```sh
gh watch-commits -R <repository>
```

将 `<repository>` 替换为您要订阅提交的仓库。

如果您已成功创建拉取请求，请注意以下几点：

- 请**不要**合并此拉取请求
  - 如果您合并了，不要担心，只需再次运行此命令即可打开另一个拉取请求。
- 在 [Notifications](https://github.com/settings/notifications) 设置的 Email 部分中启用：
  - Comments on Issues and Pull Requests
  - Pull Request reviews
  - Pull Request pushes

就是这样。您将收到有关默认分支每次提交的电子邮件通知。

运行 `gh watch-commits -h` 获取更多帮助：

```
Usage:
  gh watch-commits [-R <repository>] [flags]
		
Flags:
  -R string
    	repository using the OWNER/REPO format
  -b string
    	Body for the pull request
  -d	Mark pull request as a Draft
  -default-branch-only
    	Only include the default branch in the fork
  -fork-name string
    	Rename the forked repository
  -org string
    	Create the fork in an organization
  -t string
    	Title for the pull request
```
