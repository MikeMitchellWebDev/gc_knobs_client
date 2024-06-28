## What is it
The GC Knobs client allows you to make bulk requests to GC Knobs

## How to use 
clone and start https://github.com/MikeMitchellWebDev/gc_knobs.git

clone and run gc_knobs_client like this

./gc_knobs_client -g 200 -s 1 -r 10 -p "/Users/mm/rails"

This commmand will read a locally cloned copy of ruby on rails into memory
2000 times (spawning 200 go routines, each making the request 10 times with a sleep of 1 second between each request)


-g is the number of goroutines
-s is the sleep duration between requests
-r is the number of repeats
-p is the path to the locally cloned repository

You don't need the gc_knobs_client to use gc_knobs. You can also use gc_knobs with curl

curl -H 'Content-Type: application/json' -d '{"path":"/path/to/rails", "repeat":"50", "sleep":"2"}' -X POST  http://localhost:8000/git_repo

