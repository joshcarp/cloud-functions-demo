module github.com/joshcarp/cloudfunctionsdemo

go 1.14

replace github.com/joshcarp/rosterbot/command => ./vendor/github.com/joshcarp/rosterbot/command

require (
	cloud.google.com/go/storage v1.10.0
	github.com/joshcarp/gop v0.0.0-20200914071452-a78ba3fc78af // indirect
	github.com/joshcarp/rosterbot v0.0.0-20201030013327-690bb0c7bb1a
)
