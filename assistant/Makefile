.PHONY: deploy delete

deploy:
	gcloud functions deploy hello --env-vars-file .env.yaml --entry-point Hello --runtime go111 --trigger-http

delete:
	gcloud functions delete hello --entry-point Hello --runtime go111 --trigger-http
