.PHONY: test wk

test:
	hugo server --disableKinds=taxonomy,term --baseUrl=http://localhost
	
wk:
	git submodule update --remote --recursive
	cd app/worker && wrangler publish
