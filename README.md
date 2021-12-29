# prometheus-wechatbot-webhook
wechatbot for prometheus alertmanager webhook  

Build the app from src:
powershell:
```
$env:GOOS="linux"
cd src
go build

```

Change the `src/wechatbot.tmpl` `{{ define "wechatbot.url.api" }}https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=YourKey{{end}}` the key=your wechat bot api key.

Run As docker :
```
docker run -d --name wechatbot --restart=always \
  -v /etc/localtime:/etc/localtime \
  -v src/wechatbot.tmpl:/apps/wechatbot.tmpl \
  -p 9080:9080 lckei/promethues-wechatbot-webhook

```

Configure the promethues alertmanager config.yml:
```
...
receivers:
- name: 'wechatbot.hook'
  webhook_configs:
  - url: 'http://yourip:9080/wechatbot'
    send_resolved: true

...
```

