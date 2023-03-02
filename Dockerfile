FROM alpine:latest
LABEL maintainer="kei"
ENV VERSION 1.0
WORKDIR /apps 
ADD  src/app /apps/app
RUN  chmod +x /apps/app
ADD src/wechatbot.tmpl /apps/wechatbot.tmpl
EXPOSE 9080 
CMD ["/apps/app"]