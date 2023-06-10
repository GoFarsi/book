.PHONY: test wk

test:
	hugo server
	
wk:
	git submodule update --remote --recursive
	cd app/worker && wrangler publish
