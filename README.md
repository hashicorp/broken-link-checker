# Broken Link Checker

A wrapper around [muffet](https://github.com/raviqqe/muffet) to automate broken link checking for HashiCorp properties.

## Usage

Create a Github workflow file in your repository located at `/.github/workflows/broken-link-checker.yaml`.

```yaml
name: Broken Link Checker
on: [push]
jobs:
  check_links:
    runs-on: ubuntu-16.04
    steps:
    - uses: SFDigitalServices/wait-for-deployment-action@v2
      id: deployment
      with:
        timeout: 600
        github-token: ${{ github.token }}
        environment: Preview
    - name: Install Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.15.x
    - name: Install Muffet
      run: GO111MODULE=on go get -u github.com/raviqqe/muffet/v2
    - name: Install Broken Link Checker
      run: GO111MODULE=on go get -u github.com/hashicorp/broken-link-checker
    - name: Run
      run: broken-link-checker ${{ steps.deployment.outputs.url }}
      env:
        VERBOSE: true
        MAX_CONNECTIONS: 5
        TIMEOUT_SECONDS: 10
        EXCLUSIONS: linkedin.com,facebook.com
```

## Configuration

At the bottom of the file above, there are four environment variables that can be adjusted.

#### VERBOSE

Having this enabled will echo out all errors that came from Muffet that we deem dubious as specified by the [filterErrors](https://github.com/hashicorp/broken-link-checker/blob/master/main.go#L60) function.

This filters out a bunch of things that would have triggered an error (failed Github workflow) that really aren't indicative of an actual problem.

#### MAX_CONNECTIONS

In order to limit the number of 429 errors this is exposed as an environment variable.

The higher the value the quicker this script will run, but the more likely that servers will hit you with 429s.

#### TIMEOUT_SECONDS

The max amount of time a request can take before canceling.

#### EXCLUSIONS

Allows the filtering of specific domains to not check links from.

This is helpful for sites like LinkedIn who successfully detect the request comes from bots and throws an error.
