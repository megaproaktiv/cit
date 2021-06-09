
# Deploy

```bash
git add .
git commit -m "Changes v<version>"
git tag <version>
git push origin <version>
GOPROXY=proxy.golang.org go list -m github.com/megaproaktiv/cit@v<version>
```

# Generate

To generate moq files

```bash
go generate ./...
```