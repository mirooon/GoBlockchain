FROM node:10.16.0-alpine

RUN mkdir /app

COPY . /app/
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

RUN npm install --production
# RUN npm install react-scripts@3.4.0 -g --silent

EXPOSE 3000
# start app
CMD ["npm", "start"]