FROM node:14-alpine
WORKDIR /app/
ARG MOD

COPY package.json /app/
RUN npm install
COPY . .
RUN npm run-script build.${MOD}

FROM nginx:stable
COPY nginx.conf /etc/nginx/conf.d/default.conf
RUN rm -rf /usr/share/nginx/html/index.html
COPY --from=0 /app/dist/tides-ui /usr/share/nginx/html
EXPOSE 80
