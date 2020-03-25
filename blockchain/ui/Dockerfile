FROM node:10.16.0

RUN mkdir usr/src/app

COPY . /usr/src/app
WORKDIR /usr/src/app

ENV PATH /usr/src/app/node_modules/.bin:$PATH
ENV REACT_APP_NODEIP=$REACT_APP_NODEIP

RUN npm install --silent
# RUN npm install react-scripts@3.4.0 -g --silent

EXPOSE 3000
# start app
CMD ["npm", "start"]