# Broken Link Checker

A wrapper around [muffet](https://github.com/raviqqe/muffet) to automate broken link checking for HashiCorp properties.

## Usage

Create a Github workflow file in your repository located at `/.github/workflows/broken-link-checker.yaml`.

```yaml
name: Broken Link Checker
on: [push, pull_request]
jobs:
  check_links:
    runs-on: ubuntu-latest
    steps:
    - name: Install Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.15.x

    - name: Install Muffet
      run: GO111MODULE=on go get -u github.com/raviqqe/muffet/v2

    - name: Install Broken Link Checker
      run: GO111MODULE=on go get -u github.com/hashicorp/broken-link-checker

    - name: Run
      run: ./broken-link-checker ${{ github.event.inputs.preview_url }} # TODO need to figure out how to pull this right value
      env:
        VERBOSE: true
        MAX_CONNECTIONS: 5
        TIMEOUT_SECONDS: 10
        EXCLUSIONS: linkedin.com,facebook.com
```
