FROM nginx:latest

# Remove default nginx config
RUN rm /etc/nginx/conf.d/*

# Copy our nginx configurations
COPY ./*.conf /etc/nginx/conf.d/

# Copy hub HTML file
COPY ./hub.html /usr/share/nginx/html/