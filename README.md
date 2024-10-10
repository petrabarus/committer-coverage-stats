# Committer Code Coverage Stats
Applications to summarize committer coverage statistics

## Usage

```
committer-coverage-stats [options] <path-to-repo>

Options:
  -h, --help                    Show this help message and exit
  -v, --version                 Show version information and exit

  Output Options:
    --out-stdout                Enable standard output (default)
    --out-json [<file>]         Enable JSON output, optionally specify output file
    --out-github-pr <repo> <pr> Enable GitHub Pull Request comment output
    --out-curl <url>            Enable CURL output to specified URL

  Output Configurations:
    --out-stdout-format <format>    Set stdout format (default: plain, options: plain, colored, verbose)
    --out-json-pretty               Pretty print JSON output
    --out-github-pr-style <style>   Set GitHub PR comment style (default: simple, options: simple, detailed)
    --out-curl-method <method>      Set CURL HTTP method (default: POST, options: GET, POST, PUT, PATCH)
    --out-curl-headers <headers>    Set CURL headers (comma-separated key:value pairs)

  Filtering Options:
    --filter-min-timestamp <timestamp> Filter commits from this timestamp (ISO 8601 format)
    --filter-max-timestamp <timestamp> Filter commits until this timestamp (ISO 8601 format)
    --filter-branch <branch1> [<branch2>] Compare branches. If only one branch is specified, it's compared against HEAD

  Git Provider Options:
  This options required to get more information about the committer such as name, email, etc.

    --git-provider <provider>   Specify the Git provider for API integration (options: github, gitlab)
    --git-token <token>         Provide the API token for the specified Git provider
    --git-repo <repo>           Specify the repository in 'owner/repo' format

Examples:
  my-app --out-stdout --out-json output.json <command>
  my-app --out-stdout --out-stdout-format colored --out-github-pr myrepo 123 <command>
  my-app --out-json --out-json-pretty --out-curl https://api.example.com/endpoint --out-curl-headers "Content-Type:application/json,Authorization:Bearer token" <command>
  my-app --filter-min-timestamp 2023-01-01T00:00:00Z --filter-max-timestamp 2023-12-31T23:59:59Z --filter-branch develop main <command>
  my-app --out-stdout --filter-min-timestamp 2023-06-01T12:00:00+02:00 --filter-branch feature/new-api <command>
  my-app --filter-branch main --git-provider github --git-token YOUR_TOKEN --git-repo owner/repo <command>
```