# Golang WASM TurboFan Crash

chrome go boom

## How go boom?

Requirements:

- Chrome
- Python 3

Run either `run.sh` or `run.bat` depending on your platform. This will run a local webserver on `localhost:8989` so you can view a web page with the WASM file.

Open up Task Manager or Activity Monitor so you can keep an eye on Chrome's memory. Then go to `localhost:8989` in Chrome.

Observe that, although the page is responsive and the JS heap is not growing, after a short time Chrome starts using insane amounts of memory. On Windows, this seems to trip a memory limit and crashes the tab. On Mac (at least for some builds) it manages to proceed, but stays at high memory and CPU for a long time.

## How build?

Requirements:

- Go 1.16 (it may work with other versions, but I haven't tested it)

Run either `build.sh` or `build.bat` depending on your platform.

## Why go boom?

The WASM code contains a function with an extreme amount of nested branches. The TurboFan optimizer in Chrome takes a lot of CPU and a lot of memory attempting to optimize this function, causing at best extreme resource usage and at worst a tab crash. Ironically, this is the init function and it will never be called again, rendering this optimization useless.
