app = EGoVault

sign: $(app)
	ego sign $<

$(app):
	ego-go build -o $(app) app.go

clean:
	rm -vf *.pem $(app)
	find db/ -type f -not -name '*.go' -delete