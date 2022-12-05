.PHONY: test wk

test:
	hugo server --disableKinds=taxonomy,term --baseUrl=http://localhost
	
wk:
	wrangler publish
