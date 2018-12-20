# Validator
Validator provides an interface to check incoming request's body with given signature and private key.
The key is hold by the client/server privately.
This is a [negroni](github.com/urfave/negroni) compatible middleware.

## Checker Interface
Sha1Checker provides a demo implementation of Checker interface, which uses Sha1Sum & Base64 algorithm to sign.
You can implement you own Checker with different algorithms.

## Usage
Demo code could be fount in test/test.go.
There is no difference compared to using other [negroni](github.com/urfave/negroni) compatible middlewares.
```go 
    func handle(w http.ResponseWriter, r *http.Request) {
	    w.Write([]byte("hello there"))
    }

    n := negroni.New()
	ch := validator.Sha1Checker{
		Key:"fill_you_key_here",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)

	v := validator.NewValidator(&ch)
	n.Use(v)
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)

	log.Fatal(http.ListenAndServe(":8088", n))
```