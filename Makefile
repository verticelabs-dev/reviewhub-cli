test-webhook:
	curl -X 'POST' 'http://localhost:7070/action/branch' -H 'connection: close' -H 'content-length: 67' -H 'content-type: application/json' -H 'accept: */*' -H 'user-agent: curl/7.81.0' -H 'host: webhook.site' -d '{"branch": "tes-branch", "repo": "verticelabs-dev/reviewhub-example-app"}'
clear-cache:
	rm -rf orchestrator/storage/repos/*
	rm -rf orchestrator/storage/temp/unzip/*