FROM node:lts-alpine

# make the 'app' folder the current working directory
WORKDIR /app

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .
# install project dependencies
RUN npm install



# build app for production with minification

EXPOSE 8081
CMD ["npm", "run", "serve"]