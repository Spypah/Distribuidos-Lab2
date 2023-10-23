docker-ONU:
	sudo docker compose run --rm onu

docker-OMS:
	sudo docker compose up -d --build oms

docker-datanode:
	sudo docker compose up -d --build datanode

docker-continentes:
	sudo docker compose up -d --build continente
clean:
	sudo docker compose down