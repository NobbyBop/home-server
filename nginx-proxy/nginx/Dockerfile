FROM nginx:latest

RUN apt-get update && \
    apt-get install zip -y
# Remove default nginx config
RUN rm /etc/nginx/conf.d/*

# Copy our nginx configurations
COPY ./*.conf /etc/nginx/conf.d/

# Copy hub HTML file
COPY ./hub.html /usr/share/nginx/html/

COPY ./lib_letsencrypt.zip ./etc_letsencrypt.zip /tmp/

RUN unzip /tmp/lib_letsencrypt.zip -d /tmp/
RUN unzip /tmp/etc_letsencrypt.zip -d /tmp/

RUN cp /tmp/var/lib/letsencrypt /etc/ -r
RUN cp /tmp/etc/letsencrypt /etc/ -r