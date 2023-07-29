module go.mau.fi/mautrix-gmessages

go 1.20

require (
	github.com/gabriel-vasile/mimetype v1.4.2
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/rs/zerolog v1.29.1
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	go.mau.fi/mautrix-gmessages/libgm v0.1.0
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1
	google.golang.org/protobuf v1.31.0
	maunium.net/go/maulogger/v2 v2.4.1
	maunium.net/go/mautrix v0.15.5-0.20230728182848-1ef656165098
)

require (
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/yuin/goldmark v1.5.4 // indirect
	go.mau.fi/zeroconfig v0.1.2 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	maunium.net/go/mauflag v1.0.0 // indirect
)

replace go.mau.fi/mautrix-gmessages/libgm => ./libgm
