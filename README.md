# gitest

[![Build Status](https://drone.yq16.fun/api/badges/fx/gitest/status.svg)](https://drone.yq16.fun/fx/gitest)


## runner pvc
* https://github.com/drone-runners/drone-runner-kube/blob/master/samples/volume_claim.yml


## drone admin
* https://docs.drone.io/server/user/admin/

## build test docker image
* by hand
  * ```
    cd go/
	make
	cd ..
	docker build -t gitest:latest .
	docker run --rm -p 8888:8080 -it gitest
    ```
* by ci drone
  * ```
    git tag -a v0.0.14   -m"alpine image"
	git push origin v0.0.14

	# wait drone ..............    

    docker pull harbor.yq16.fun/test/gitest
	docker run --rm -p 8888:8080 -it harbor.yq16.fun/test/gitest
    ```
