package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	a := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>CORS跨域</title>
    <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
</head>
<body>
<script type="text/javascript">
        function sub(submitform) {
             var formData= new FormData(document.getElementById(submitform));
            $.ajax({
                url: "test",
                type: 'POST',
                data: formData,
                dataType: "json",        //返回数据形式为json
                contentType: false,
                cache: false,
                processData: false,
                success: function (msg) {
                    if (msg == 'success') {
                        alert("成功")
                    }
 
                },
                error: function (msg) {
                    alert(msg)
                }
            });
        }
    </script>
 
<form action="" method="post" id="submitform" enctype="multipart/form-data">
    <input type="text" name="fileName" id="fileName"/><p></p>
    <input type="file" name="fileUrl"/>
    <input type="button" onclick="sub('submitform')"/>
</form>
</body>
</html>`

	_, err := w.Write([]byte(a))
	if err != nil {
		fmt.Println(err)
		return
	}

}

func formTest(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bodyBytes))
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/test", formTest)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
