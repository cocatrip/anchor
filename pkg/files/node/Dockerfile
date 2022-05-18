FROM nginx:stable-alpine

ENV TZ=Asia/Jakarta

RUN apk add tzdata
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

COPY dist/ /usr/share/nginx/html
ADD ./nginx.conf /etc/nginx/nginx.conf
EXPOSE [[ .Docker.EXPOSE_PORT]]
CMD ["nginx", "-g", "daemon off;"]
