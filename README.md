Searching for a stupidly simplewebserver for development useage.  I
want it to serve static content and require no configuration beyond
the command line.

`python -m simpleserver` easily overloads when using moderns browser
pages that do a lot of loads at once.

Tried making a better python one, but it still involves package management and could be faster.

thttpd basically nailed it, but hasn't been maintained for so long
that its build system has somewhat rotted, and it requires build tools
on any box that you want to use it for.

If the tool is to be compile with a language that isn't always
pre-installed, it should be statically compiled such that the binary
can be reliably copied around.  Golang excels at that, so this dresses
up the build in server with some command line arguments, then gets out
of the way.