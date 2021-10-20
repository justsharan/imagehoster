# imagehoster
This is a really simple utility you can use to serve images from your own domain. Think of it like your own, personal imgur.

All the other GitHub repos I found were awfully bloated for my use case, so I just decided to make my own. It's extremely fast, extensible, and written in less than 100 lines of Go.

### Setup
Once you build the utility, you just need to set these things:

 * `$USAGE_KEY` (env) : Authorization header that it looks for in POST requests
 * `-p` (flag) : Port that the utility listens to (Default: `4000`)
 * `-dir` (flag) : Directory where images will be stored (Default: `./uploads/`)

### Usage
The utility is used entirely through its REST endpoints:

| Request        | Description                           |
|----------------|---------------------------------------|
| `GET /:image`  | Gets an image identified by that tag. |
| `POST /upload` | Creates a new image.                  |

### Tips
1. This utility is ShareX-friendly. You can use the information above to add it as a private ShareX endpoint in your settings (I'm not on Windows so I wouldn't know much about actually configuring this, though).
2. If you want to include a landing page when the user hits `/`, just add an `index.html` file in that directory.

### License
Refer to the [LICENSE](LICENSE) file.
