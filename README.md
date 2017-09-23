# CruizinSolutions
Rim/tire shop inventory and invoice manager

To start as development:
  - clone project into your github.com folder
  - go run cruizinserver/server.go
  - cd into cruizinweb/
  - npm start
  
To start as production:
  - cd into cruizinserver/
  - go build
  - there will be a new exe file named cruizinserver.exe
  - cd back into cruizinweb
  - npm run build
  - there will be a new folder named build
  - cp -r build ../cruizinserver/public
  - run exe
