## Simple Blockchain example

This is the simplest blockchain example possible.
it implements only a chain of blocks with data.
And does not include any of high order concepts like mining of chain verification. 

### How to use it
**1)** First, download the repository
```bash
git clone https://github.com/Rufaim/Simple-golang-blockchain-example.git
```
OR
```bash
go get github.com/Rufaim/Simple-golang-blockchain-example
```

**2)** Second, run a server
```bash
go run . config.json
```
Message `Server is on http://localhost:8080` should appear is console for default config.

**3)** Then open visualization interface  at `http://localhost:8080`.
![interface](interface_screenshot.png)

Direct API requests are also supported. \
To add block use
```bash
curl -X POST http://localhost:8080/bc -d '{"data":"some_data"}'
```
To get all blocks use
```bash
curl -X GET http://localhost:8080/bc
```

Please note that a block added via API request is visible on visualization interface after page update.
