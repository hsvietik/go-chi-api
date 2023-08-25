# go-chi-api

1. I created a Handler function that creates a response containing whatever is passed in {mydata} used chi.URLParam to find {mydata} value.
2. I added myMiddleware function and set {mydata} value to the context, and changed Handler function to use {mydata} value from the context.
3. The [front-end part](https://github.com/hsvietik/test-react-front) was created with React to create the request with Roman numerals. To handle this request the function was created in separate package of back-end part.
4. Unit tests added to package for converting Roman numerals to integers
