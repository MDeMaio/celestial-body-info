# celestial-body-info

## About:
This application was created in order to track data related to objects in space for educational and recreational reasons. <br/>
It utilizes golang & gRPC in the backend, vue.js(Bootstrap for styles) in the frontend, and a Mongodb for data storage. <br/>
Features are constantly being worked on and added each day. This will be hosted live soon as well on the web.

## Running the Application:
To run this application locally, make sure you have the below requirements on your computer: <br/>
1) [golang](https://golang.org/) I recommended the newest version to avoid issue with go modules.<br/>
2) [*OPTIONAL Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/) Follow the guide to download the correct version for your pc.<br/>
3) [NPM](https://www.npmjs.com/get-npm) NPM will be used to handle our packages in the frontend, it is downloaded with Node.js.<br/>
4) [vue.js](https://www.npmjs.com/package/vue) vue.js will be our frontend framework used to build the users interface.<br/>
5) [Mongodb](https://www.mongodb.com/) a local instance of mongo will be used right now to store data for testing.<br/>

Currently running the app requires 3 terminals, 1 for the gRPC server, 1 for the http server, and 1 for the vue frontend. <br/>
A local instance of mongodb is also required to be running for this application. Make sure you have mongodb up and running before moving to the next steps. <br/>

Follow the below commands to get the entire application up and running: <br/>
To start the gRPC server: `go run server.go` --> Run this in the planet/planet_server directory.<br/>
To start the API server: `go run main.go`  --> Run this in the root directory of the project.<br/>
To start the Frontend portion: `npm run serve`  --> Run this in the client directory.<br/>

To view the application in the web, visit [localhost](http://localhost:8081)<br/><br/>