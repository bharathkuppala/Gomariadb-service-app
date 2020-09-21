endpoint : http://localhost:8080/api/v1/register
{
    "firstName" : "Bharath",
    "lastName" : "Kuppala",
    "email" : "b@gmail.com",
    "bloodGroup" : "B+",
    "password" : "pass",
    "userName": "Bharath Kuppala"
}

endpoint :  http://localhost:8080/api/v1/login
{
     "email" : "b@gmail.com",
     "password" : "pass"
}

<!-- Just for reference -->
DSN: bharath:password@tcp(localhost:3306)/mariadb_test?checkConnLiveness=false&parseTime=true&maxAllowedPacket=0

kafka commands (To run locally)
JMX_PORT=8004 bin/kafka-server-start.sh config/server.properties ---> to start kafka
bin/zookeeper-server-start.sh config/zookeeper.properties ---> to start zookeeper
bin/cmak -Dconfig.file=conf/application.conf -Dhttp.port=9000 ----> to start kafka manager(gui)

<!-- Commands need to be execute to run services inside container -->
docker-compose up -d ----> for background mode
docker-compose up ----> for foreground mode

<!-- To get inside container -->
docker exec -it kafka bash or sh
navigate to /opt/kafka_2.12-2.3.0/bin/
kafka-topics.sh --list --zookeeper zookeeper:2181 ---> list of topics available 

<!-- Output: -->
__consumer_offsets
mariadb-events

<!-- Command to see the messages which are published to topic -->
kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic mariadb-events --from-beginning
Bharath Kuppala successfully logged in at: 2020-09-21 04:31:57
bharath Kuppala successfully logged in at: 2020-09-21 05:21:40
bharath kuppala student successfully registered


