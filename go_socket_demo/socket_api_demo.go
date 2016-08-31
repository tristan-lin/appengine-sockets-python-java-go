package socket-api-demo
import (
        "fmt"
        "net/http"
        "time"
        "strconv"
        "strings"
        "google.golang.org/appengine"
        "google.golang.org/appengine/socket"
)
func init() {
    http.HandleFunc("/", handler)
}
func handler(w http.ResponseWriter, r *http.Request) {
        //
        // socket API
        //
        fmt.Fprintf(w,"<h1>Socket API Demo - GOlang</h1>")
        fmt.Fprintf(w,"<br>This demo connects to a socket server on a GCE instance and display the file name/size it is providing.<br>")
        ctx := appengine.NewContext(r)
        //
        conn, err := socket.DialTimeout(ctx,"tcp", "server_ip:server:port",time.Second*3)
        if err != nil {
            fmt.Fprintf(w, "dial error:", err)
            return
        }
        defer conn.Close()
        bufferFileName := make([]byte, 64)
        bufferFileSize := make([]byte, 10)
        if _, err := conn.Read(bufferFileSize); err != nil {
                        return 
                }
        fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
        if _, err := conn.Read(bufferFileName); err != nil {
                        return 
                }
        fileName := strings.Trim(string(bufferFileName), ":")
        //
        // Output to HTTP response
        //
        fmt.Fprintf(w,"<br>file Name --> %s", fileName)
        fmt.Fprintf(w,"<br>file Size --> %d byte",fileSize)
        return
        }
                
