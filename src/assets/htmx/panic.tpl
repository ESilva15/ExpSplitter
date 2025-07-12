<code class="warning">
  {{ .msg }}
  <br>
  error: invalid resource
  /usr/lib/go/src/runtime/panic.go:261 (0x455737)
  	panicmem: panic(memoryError)
  /usr/lib/go/src/runtime/signal_unix.go:881 (0x455705)
  	sigpanic: panicmem()
  /home/esilva/Desktop/projetos/go_website/src/views/projects.go:55 (0xa25ab6)
  	ProjectsContent: if *req == e.Name() {
  /home/esilva/Desktop/projetos/go_website/src/views/projects.go:129 (0xa26244)
  	ProjectsPage: data, err := ProjectsContent(nil)
  /home/esilva/go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/context.go:174 (0x91d6b9)
  	(*Context).Next: c.handlers[c.index](c)
  /home/esilva/go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/recovery.go:102 (0x91d6a7)
  	CustomRecoveryWithWriter.func1: c.Next()
  /home/esilva/go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/context.go:174 (0x91c7fc)
  	(*Context).Next: c.handlers[c.index](c)
  /home/esilva/go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/logger.go:240 (0x91c7e3)
  	LoggerWithConfig.func1: c.Next()
  /home/esilva/go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/context.go:174 (0x91bced)
  	(*Context).Next: c.handlers[c.index](c)
  /home/esilva/go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/gin.go:620 (0x91b97c)
  	(*Engine).handleHTTPRequest: c.Next()
  /home/esilva/go/pkg/mod/github.com/gin-gonic/gin@v1.9.1/gin.go:576 (0x91b4b1)
  	(*Engine).ServeHTTP: engine.handleHTTPRequest(c)
  /usr/lib/go/src/net/http/server.go:3137 (0x6ee7ed)
  	serverHandler.ServeHTTP: handler.ServeHTTP(rw, req)
  /usr/lib/go/src/net/http/server.go:2039 (0x6e9ba7)
  	(*conn).serve: serverHandler{c.server}.ServeHTTP(w, w.req)
  /usr/lib/go/src/runtime/asm_amd64.s:1695 (0x472d80)
  	goexit: BYTE	$0x90	// NOP
</code>
