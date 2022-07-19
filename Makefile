app = EGoVault
workPath = $(pwd)

getid: sign
	ego uniqueid $(app)

sign: $(app)
	ego sign $<

$(app):
	ego-go build -o $(app) app.go

clean:
	rm -vf *.pem $(app)
	cd db ; ls | grep -vE "*.(go)" | xargs rm -f ; cd $(workPath)
