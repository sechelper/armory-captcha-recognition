# armory-captcha-recognition

## Test

```bash
$ curl -X POST -H "Content-Type: multipart/form-data"   -F "captcha[]=@testdata/captcha/1577.png" -F "captcha[]=@testdata/captcha/1597.png" -F 'options={"psm":13,"languages":["eng"], "whitelist":"1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"}' http://127.0.0.1:8085/v1/captcha/image
$ curl -X POST -H "Content-Type: application/json" -d '{"psm":13,"languages":["eng"], "whitelist":"1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ","captcha":[{"id":"index.png", "text":"iVBORw0KGgoAAAANSUhEUgAAACwAAAASCAMAAAAXKszuAAACZ1BMVEX19fXPPeCcPI3dXcHmPoh0JMAeiwlvj2aashQWjtBCXk+MymZH/Yz9c8wcG1ueBUMh3qIxTwEF4CDdJKokktFcrzhR2YsA8hngXfyU663k+xE8kaW3lxp9lazPzX5vYPiWAKzzj/pBxldMdIdhp06LKJ4AG5BdJcdg6h4QjNxvOnpMXM/FChjqUjITXo8aZ1ZtRHebBdAdez80cBrRYBOFyBUw4zDPRMsTou4nmRit6A7S/02H5R+YABiNEV7UmhHUAA71OxnJKFj7n5JZyOFXL5bbZcEXRm2hR2SD9+/7sWvGB7aWhwyizp5gAaiAudnFACQg9s8RsC4v/FO/Gh9zk3yqzC9H/FhLMGz16TILmgshVtqxalRqbuRZapREYeqMrKYOI7Yv7P8AA6/HkmrxbShEtPWh1Ngb2h/q/Xh5VYiWIZrzoyp2wV1uiWiC1ZPogjZs23QMjlr6JzGEzwbMXp+CusS/6dDctnExQXn60E+Tm9zL2rw14XghW2M+DR2mVggXBL0nGo5kiR4fDdbIe+bA2AWXvyM7/pxoGgGh/LInxL/m0lkn0aXT8dqyWvXi01rDugL0hhGLiYOUxixhBfaXVp9g1p+PbBz2fdwRdl5M/6ezDePilBxWqq4CpLTYLl1Ux14onVK7oVKPXBGNsoS22/rDddn1AyQJdIgvMbDCMU+8iR438wxQnDTk/H8lZwHi4bAuYoqXpAR+molfMEudajQW6OEAfcyVd4JEPvYQuUbyxuPIJiJ2OJ2DR8/Jg28E35202JIJnu+rc21LraMtFmYjQwl6Z0gbAnx3NJ1JlMWIwlMTpk6jAAAACXBIWXMAAA7EAAAOxAGVKw4bAAABFklEQVQokWNgAAEeBhiYzqAMonYxMBTBxYQZvBiwgWQgTkRwS7EqQoDdBOSRwEzswtHIHEZGKIMJCICUATMEqDKYMqxnAQIUtVDFNkwg9UAGMzOQKAcJgRUiqWaEKWZi2sZQxgBSrAZUbIRkHEL1cQZRiOJ5TAxyINoTZDIz2HR0tSCztaEmM/SCXQFRCVMNcXMhiLkA5sFjDGDfgXAOA8MyBoRqNJOXAwktBiZdqPlAsA9ZMUj1QbSgAyqrhSiGKIOQXWhGgxVvhaiDORqqFqRsMlrQwc2GqGXokYaFRj48UrYwIINZDGigFUymAXF7AEzsJLoiBobDmEKYoJ/BHUw3n4YL+YtgKhNEMNdA6cx1UMZCCAUArYEe3879S5oAAAAASUVORK5CYII="}]}' http://127.0.0.1:8085/v1/captcha/image
```