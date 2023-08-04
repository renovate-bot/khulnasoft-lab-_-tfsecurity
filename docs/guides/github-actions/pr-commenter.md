---
title: Github Action - PR Commenter
description: Github Action - PR Commenter
subtitle: Adding tfsecurity PR commenter to your GitHub project
description: Adding tfsecurity PR commenter to your GitHub project
author: tfsecurity
tags: [installation, github action, PR commenting]
---

## What is it?

The PR Commenter action will process a Pull request and add comments to any areas of the change which fail the `tfsecurity` checks.

## Adding the action


To add the action, add tfsecurity_pr_commenter.yml into the .github/workflows directory in the root of your Github project.

![Setup a new workflow](../../imgs/newworkflow.png)

The contents of tfsecurity_pr_commenter.yml should be;

```yaml
{% raw %}
name: tfsecurity-pr-commenter
on:
  pull_request:
jobs:
  tfsecurity:
    name: tfsecurity PR commenter
    runs-on: ubuntu-latest

    steps:
      - name: Clone repo
        uses: actions/checkout@master

      - name: tfsecurity
        uses: tfsecurity/tfsecurity-pr-commenter-action@main
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
{% endraw %}
```


On each pull request and subsequent commit, tfsecurity will run and add comments to the PR where tfsecurity has failed.

The comment will only be added once per transgression.

## Example PR Comment

The screenshot below demonstrates the comments that can be expected when using the action

![PR Commenter Example](../../imgs/pr_commenter.png)
