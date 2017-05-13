all:
	git add web-frontend/build && git commit -m "Initial dist subtree commit"
	git subtree push --prefix web-frontend/build origin gh-pages

