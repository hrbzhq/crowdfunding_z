# Publishing the docs (GitHub Pages)

This repository includes a small documentation site under the `docs/` folder. The repository contains a GitHub Actions workflow that publishes `docs/` to the `gh-pages` branch and powers GitHub Pages.

Where the site is published

- By default the workflow publishes to the `gh-pages` branch. After the first successful publish, open the repository Settings → Pages and confirm the source is set to the `gh-pages` branch.
- The public URL will normally be:

```
https://<owner>.github.io/<repo>/
```

Replace `<owner>` and `<repo>` with your repository owner and name (`hrbzhq/crowdfunding_z`), so the URL will be:

```
https://hrbzhq.github.io/crowdfunding_z/
```

Triggering a publish

- Push changes to any files under `docs/` on the `master` branch. The workflow is configured to run on pushes that affect `docs/**` and will automatically publish the updated content.
- Or run the `Publish Docs to GitHub Pages` workflow manually from the Actions tab (select the workflow and click "Run workflow").

Permissions and tokens

- The workflow uses the built-in Actions token available as `secrets.GITHUB_TOKEN`. In most repositories this is sufficient to publish to `gh-pages`.
- If your organization restricts the built-in token, create a PAT with the least privileges needed (workflow/repo scope for pages) and add it to the repository secrets (for example `PUBLISH_PAT`) and update the workflow to use that secret.

Troubleshooting

- If the workflow fails with permission errors, check repository Settings → Actions → General for any restrictions on `GITHUB_TOKEN` or required approvals.
- If the site doesn't appear after a successful deploy, verify the Settings → Pages configuration and that the `gh-pages` branch exists and contains the published HTML.
- If the workflow never triggers after pushing `docs/` changes, confirm you pushed to the `master` branch and that the `paths:` filter in `.github/workflows/publish_docs.yml` matches the changed files.

If you'd like, I can also:
- Trigger the workflow manually once to create the `gh-pages` branch and confirm the site is live.
- Update the workflow to publish to a different branch or change the trigger rules.
