# action-go-template

[![deploy-image](https://github.com/hi-actions/action-go-template/actions/workflows/Deploy-image.yml/badge.svg)](https://github.com/hi-actions/action-go-template/actions/workflows/Deploy-image.yml)
[![Test-run](https://github.com/hi-actions/action-go-template/actions/workflows/Test-run.yml/badge.svg)](https://github.com/hi-actions/action-go-template/actions/workflows/Test-run.yml)

Template for github action with golang

## Init Template

- please replace `action-go-template` to your project name.
- please replace `YOUR_IMAGE_NAME` to your docker image name.

### Example Workflow

Create a workflow (eg: `.github/workflows/my-action.yml` see [Creating a Workflow file](https://help.github.com/en/articles/configuring-a-workflow#creating-a-workflow-file)) to utilize the labeler action with content:

```yml
name: "Pull Request Labeler"
on:
- pull_request
jobs:
  labels:
    runs-on: ubuntu-latest
    steps:
    - uses: hi-actions/action-go-template@master
      env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          LABEL_CONFIG: .github/labeler.yml # this is default
```

_Note: This grants access to the `GITHUB_TOKEN` so the action can make calls to GitHub's rest API_

