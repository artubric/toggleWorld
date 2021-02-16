### toggleWorld
In order to run the project, you must have following installed on your system: <br />
- nodejs/npm <br />
- angular <br />
- go<br />
- mongoDB (running locally on default port 27017)<br />

Then clone/download the repo and run following commands:
```
cd fe
npm i
ng build --prod
cp ./dist/toggleWorld/* ../be/static
cd ../be
go build
./toggleWorld.exe
```
App will be available at http://localhost:5000
