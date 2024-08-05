FROM nginx:alpine

COPY  ./dist /usr/share/nginx/html
COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf

EXPOSE 8080


CMD ["nginx", "-g", "daemon off;"]
