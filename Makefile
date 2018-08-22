all: consumer producer

.PHONY: consumer
consumer:
	cd consumer; go build -o ../bin/consumer; cd ..

.PHONY: producer
producer:
	cd producer; go build -o ../bin/producer; cd ..

clean:
	rm -rf bin
