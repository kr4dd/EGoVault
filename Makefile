app = EGoVault
workPath = $(pwd)

sign: $(app)
	ego sign $<

$(app):
	ego-go build -o $(app) app.go

clean:
	rm -vf *.pem $(app)
	cd db ; ls | grep -vE "*.(go|json)" | xargs rm -f ; cd $(workPath)