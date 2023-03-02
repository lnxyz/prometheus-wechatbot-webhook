package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/golang/glog"
	alertmanager_template "github.com/prometheus/alertmanager/template"
)

func main() {
	http.HandleFunc("/wechatbot", func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var t alertmanager_template.Data

		err := decoder.Decode(&t)
		if err != nil {
			glog.Error(err)
		}

		tmpl := template.Must(template.ParseFiles("wechatbot.tmpl"))

		//@parma: tpl 获取模板字节内容
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, t); err != nil {
			glog.Error(err)
		}

		//@parma: wechatbotUrlBytes wechatbot api url with key
		var wechatbotUrlBytes bytes.Buffer
		if err := tmpl.ExecuteTemplate(&wechatbotUrlBytes, "wechatbot.url.api", "no data needed"); err != nil {
			glog.Error(err)
			return
		}

		//post the context to the wechat api
		postBody, _ := json.Marshal(map[string]interface{}{
			"msgtype": "markdown",
			"markdown": map[string]interface{}{
				"content": tpl.String(),
			},
		})
		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post(wechatbotUrlBytes.String(), "application/json", responseBody)
		if err != nil {
			glog.Error(err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != 200 {
			if err != nil {
				glog.Error(err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}
			glog.Error("Broken : ", string(body))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
		} else {

			w.WriteHeader(http.StatusOK)
			w.Write(body)
		}

	})

	http.ListenAndServe(":9080", nil)
}
